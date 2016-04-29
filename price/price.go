package price

import (
    "github.com/wangchen1ren/stock-go/utils"
    "time"
    "github.com/doneland/yquotes"
    "github.com/zpatrick/go-config"
    "github.com/golang/glog"
    "github.com/jmoiron/sqlx"
)
import (
    _ "github.com/go-sql-driver/mysql"
    "sort"
)

type Price struct {
    Original yquotes.PriceH;
    Date     time.Time
    Open     float64
    High     float64
    Low      float64
    Close    float64
    Volume   float64
    AdjClose float64
    //Price    yquotes.PriceH;
}

func NewPrice(p yquotes.PriceH) Price {
    var price Price
    price.Original = p
    price.Date = p.Date
    if p.AdjClose != p.Close {
        rate := p.Close / p.AdjClose;
        price.Open = p.Open / rate;
        price.High = p.High / rate;
        price.Low = p.Low / rate;
        price.Close = p.Close / rate;
    } else {
        price.Open = p.Open;
        price.High = p.High;
        price.Low = p.Low;
        price.Close = p.Close;
    }
    price.Volume = p.Volume
    price.AdjClose = p.AdjClose
    return price
}

type Prices []Price;

func (prices Prices) Len() int {
    return len(prices)
}
func (prices Prices) Less(i, j int) bool {
    return prices[i].Date.Before(prices[j].Date)
}
func (prices Prices) Swap(i, j int) {
    prices[i], prices[j] = prices[j], prices[i]
}

func GetPrices(conf *config.Config, symbol string, from, to time.Time) (Prices, error) {
    origins, err := GetYquotesPrices(conf, symbol, from, to)
    if err != nil {
        return nil, err
    }
    return MakePrices(origins), nil;
}

func MakePrices(originals []yquotes.PriceH) Prices {
    prices := make(Prices, len(originals));
    for i, original := range originals {
        prices[i] = NewPrice(original);
    }
    sort.Sort(prices)
    return prices;
}

// Get original prices from yahoo
// cache the data in local database
func GetYquotesPrices(
conf *config.Config, symbol string, from, to time.Time) ([]yquotes.PriceH, error) {
    // normalize from and to in range 2000-01-01 to NOW()
    if from.Before(utils.StartDate()) {
        from = utils.StartDate()
    }
    if to.After(utils.Today()) {
        to = utils.Today()
    }
    glog.Infof("Get %s prices from %s to %s", symbol,
        utils.DateFormat(from, "Y-m-d H:i"), utils.DateFormat(to, "Y-m-d H:i"));

    // 1. Get fetch time from db
    start, end := getRangeFromDb(conf, symbol);
    glog.Infof("Current cache from %s to %s",
        utils.DateFormat(start, "Y-m-d H:i"), utils.DateFormat(end, "Y-m-d H:i"));

    // 2. local data not exists, fetch and save

    if from.Before(start) {
        if err := savePricesToDb(conf, symbol, from, start); err != nil {
            return nil, err
        }
    }
    if to.After(end) {
        if end.Before(from) {
            end = from;
        }
        if err := savePricesToDb(conf, symbol, end, to); err != nil {
            return nil, err
        }
    }

    // 3. Get fetched data
    prices, err := getPricesFromDb(conf, symbol, from, to);
    if err != nil {
        return nil, err
    }
    return convertPriceDbToPrice(prices), nil;
}

type priceDb struct {
    Id       int       `db:"id"`
    Symbol   string    `db:"symbol"`
    Date     time.Time `db:"date"`
    Open     float64   `db:"open"`
    High     float64   `db:"high"`
    Low      float64   `db:"low"`
    Close    float64   `db:"close"`
    AdjClose float64   `db:"adjclose"`
    Volume   float64   `db:"volumn"`
}

// Get local cache data range
func getRangeFromDb(conf *config.Config, symbol string) (time.Time, time.Time) {
    db, err := utils.GetDb(conf)
    if err != nil {
        return time.Unix(0, 0), time.Unix(0, 0);
    }
    defer db.Close();
    var start, end time.Time;
    db.Get(&start, "SELECT MIN(date) FROM price WHERE symbol = ?", symbol)
    db.Get(&end, "SELECT MAX(date) FROM price WHERE symbol = ?", symbol)
    return start, end;
}

// Fetch and save price
func savePricesToDb(conf *config.Config, symbol string, from, to time.Time) (error) {
    glog.Infof("Fetching %s prices from %s to %s", symbol,
        utils.DateFormat(from, "Y-m-d H:i"), utils.DateFormat(to, "Y-m-d H:i"));
    prices, err := yquotes.GetDailyHistory(symbol, from, to);
    if err != nil {
        return err;
    }

    db, err := utils.GetDb(conf)
    if err != nil {
        return err;
    }
    defer db.Close();

    tx := db.MustBegin()
    for _, price := range prices {
        if exists, _ := isPriceExists(db, symbol, price.Date); exists {
            glog.Infof("Price[%s @ %s] already saved.", symbol, utils.DateFormat(price.Date, "Y-m-d H:i"));
            continue
        }
        glog.Infof("Saving price[%s @ %s]", symbol, utils.DateFormat(price.Date, "Y-m-d H:i"))
        sql := "INSERT INTO price (symbol, date, open, high, low, close, adjclose, volumn) "
        sql += "VALUES (:symbol, :date, :open, :high, :low, :close, :adjclose, :volumn)"
        //glog.Infof("Insert sql: %s", sql)
        res, _ := tx.NamedExec(sql,
            map[string]interface{}{
                "symbol": symbol,
                "date": price.Date,
                "open": price.Open,
                "high":price.High,
                "low":price.Low,
                "close":price.Close,
                "adjclose":price.AdjClose,
                "volumn":price.Volume,
            });
        if rows, err := res.RowsAffected(); rows <= 0 && err != nil {
            glog.Error("Save error, rollback.")
            glog.Error(err)
            tx.Rollback()
            return err;
        }
    }
    return tx.Commit()
}

func isPriceExists(db *sqlx.DB, symbol string, date time.Time) (bool, error) {
    var id int
    err := db.Get(&id, "SELECT id FROM price WHERE symbol = ? AND date = ?", symbol, date)
    if err != nil {
        return false, err;
    }
    return id != 0, nil;
}

func getPricesFromDb(conf *config.Config, symbol string, from, to time.Time) ([]priceDb, error) {
    db, err := utils.GetDb(conf)
    if err != nil {
        return nil, err;
    }
    defer db.Close();

    prices := []priceDb{};
    err = db.Select(&prices, "SELECT * FROM price WHERE symbol = ? AND date >= ? AND date <= ?",
        symbol, from, to);
    //err = db.Select(&prices, "SELECT * FROM price")
    return prices, err;
}

func convertPriceDbToPrice(prices []priceDb) []yquotes.PriceH {
    list := make([]yquotes.PriceH, len(prices))
    for i, price := range prices {
        list[i] = yquotes.PriceH{
            Date: price.Date,
            Open: price.Open,
            High: price.High,
            Low : price.Low,
            Close :price.Close,
            Volume :price.Volume,
            AdjClose :price.AdjClose,
        };
    }
    return list
}

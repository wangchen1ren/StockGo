package stock

import (
    "github.com/doneland/yquotes"
    "fmt"
    "os"
    "github.com/zpatrick/go-config"
    "github.com/ethereum/go-ethereum/logger/glog"
    "time"
)

func Test() {
    conf, _ := LoadConfig(CONFIG_FILE)
    //TestGetFetch(conf)
    //TestSavePricesToDb(conf);
    //GetDb(conf)
    TestGetPrices(conf)

    //symbol := "002721.sz";
    //hist, _ := yquotes.HistoryForYears(symbol, 1, yquotes.Daily);
    //prices := NewPrices(hist);
    //prices.tr();
    //prices.n20();
    //fmt.Printf("%+v", prices);
}

func TestGetFetch(conf *config.Config) {
    //symbol := "002721.sz"
    symbol := "123"
    fetch, _ := getRangeFromDb(conf, symbol)
    fmt.Printf("%+v", fetch);
}

func TestGetPrices(conf *config.Config) {
    symbol := "002721.sz"
    from := time.Unix(1451577600, 0).AddDate(0, 0, 1)
    to := time.Now().AddDate(0, 0, -1)
    prices, err := GetPrices(conf, symbol, from, to);
    if err != nil {
        glog.Error(err.Error());
    }
    fmt.Println(prices)
}

func TestSavePricesToDb(conf *config.Config) {
    symbol := "002721.sz"
    //hist, _ := yquotes.HistoryForYears(symbol, 1, yquotes.Daily);
    from := time.Unix(1451577600, 0)
    to := time.Now()
    err := savePricesToDb(conf, symbol, from, to)
    if err != nil {
        glog.Error(err);
    }
    glog.Infof("success")
}

func TestPrice(symbol string) {
    price, err := yquotes.GetPrice(symbol)
    if err != nil {
        fmt.Fprintln(os.Stdout, err)
    }
    fmt.Fprintf(os.Stdout, "%v", price)
    fmt.Fprintln(os.Stdout)
    fmt.Fprintf(os.Stdout, "%+v", price)
    fmt.Fprintln(os.Stdout)
    fmt.Fprintf(os.Stdout, "%#v", price)
    fmt.Fprintln(os.Stdout)
}

func TestStock(symbol string) {
    ss, err := yquotes.NewStock(symbol, true)
    if err != nil {
        fmt.Fprintln(os.Stdout, err)
    }

    fmt.Fprintln(os.Stdout, ss.Name)
    fmt.Fprintln(os.Stdout, ss.Symbol)
    fmt.Fprintln(os.Stdout, ss.Price)
    for i := 0; i < len(ss.History); i++ {
        fmt.Fprintln(os.Stdout, ss.History[i])
    }
}

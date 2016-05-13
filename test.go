package stock

import (
    "github.com/doneland/yquotes"
    "fmt"
    "os"
    "github.com/zpatrick/go-config"
    "github.com/golang/glog"
    //"github.com/ethereum/go-ethereum/logger/glog"
    "time"
    //"github.com/wangchen1ren/stock-go/conf"
    "github.com/wangchen1ren/stock-go/utils"
    "github.com/wangchen1ren/stock-go/strategies"
    "github.com/wangchen1ren/stock-go/price"
    "github.com/wangchen1ren/stock-go/conf"
)

func Test() {
    conf, _ := conf.LoadConfig(conf.CONFIG_FILE)
    //TestGetFetch(conf)
    //TestSavePricesToDb(conf);
    //GetDb(conf)
    //TestGetPrices(conf)
    TestTurtle(conf)

    //symbol := "002721.sz";
    //hist, _ := yquotes.HistoryForYears(symbol, 1, yquotes.Daily);
    //prices := NewPrices(hist);
    //prices.tr();
    //prices.n20();
    //fmt.Printf("%+v", prices);
}

func TestTurtle(conf *config.Config) {
    symbol := "002721.sz"
    from, _ := utils.DateParse("20160101", "Ymd")
    to, _ := utils.DateParse("2016-01-10", "Y-m-d")
    prices, _ := price.GetPrices(conf, symbol, from, to)
    turtle := strategies.Turtle{}
    turtle.Eval(symbol, prices)
}

func TestGetPrices(conf *config.Config) {
    symbol := "002721.sz"
    from := time.Unix(1451577600, 0).AddDate(0, 0, 1)
    to := time.Now().AddDate(0, 0, -1)
    prices, err := price.GetPrices(conf, symbol, from, to);
    if err != nil {
        glog.Error(err.Error());
    }
    fmt.Println(prices)
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

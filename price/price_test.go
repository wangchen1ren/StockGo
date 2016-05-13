package price

import (
    "fmt"
    "testing"
    "github.com/golang/glog"
    "github.com/wangchen1ren/stock-go/conf"
    "github.com/wangchen1ren/stock-go/utils"
)

func TestGetPrices(t *testing.T) {
    conf, _ := conf.LoadConfig(conf.CONFIG_FILE);
    symbol := "002721.sz"
    from, _ := utils.DateParse("20160101", "Ymd")
    to, _ := utils.DateParse("2016-01-10", "Y-m-d")
    prices, err := GetPrices(conf, symbol, from, to);
    if err != nil {
        glog.Error(err.Error());
    }
    glog.Info(prices.Len())
}

func BenchmarkGetPrices(b *testing.B) {
    conf, _ := conf.LoadConfig(conf.CONFIG_FILE);
    symbol := "002721.sz"
    from, _ := utils.DateParse("20160101", "Ymd")
    to, _ := utils.DateParse("2016-01-10", "Y-m-d")
    prices, err := GetPrices(conf, symbol, from, to);
    if err != nil {
        glog.Error(err.Error());
    }
    fmt.Println(prices)
}

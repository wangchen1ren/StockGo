package stock

import (
    "fmt"
    "os"
    "github.com/doneland/yquotes"
    "github.com/zpatrick/go-config"
    "github.com/jmoiron/sqlx"
)

func DumpPrices(prices []yquotes.PriceH) {
    for i := 0; i < len(prices); i++ {
        fmt.Fprintf(os.Stdout, "%+v\n", prices[i])
    }
}

func GetDb(conf *config.Config) (*sqlx.DB, error) {
    host, _ := conf.String("db.host")
    port, _ := conf.IntOr("db.port", 3306)
    user, _ := conf.String("db.user")
    pass, _ := conf.String("db.pass")
    dbname, _ := conf.String("db.dbname")
    addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, pass, host, port, dbname);
    //glog.Info(addr);
    return sqlx.Open("mysql", addr)
}

package stock

import (
    "github.com/doneland/yquotes"
    "github.com/zpatrick/go-config"
    "time"
    "math"
)

type TurtleStock struct {
    Prices Prices
    TR     []float64
    N20    []float64
}

func NewTurtleStock(conf *config.Config, symbol string, from, to time.Time) (TurtleStock, error) {
    prices, err := GetPrices(conf, symbol, from, to)
    if err != nil {
        return TurtleStock{}, err
    }
    return NewTurtleStockByPrices(prices), nil
}

func NewTurtleStockByPrices(prices Prices) TurtleStock {
    var turtle TurtleStock
    turtle.Prices = prices
    turtle.calcTr()
    turtle.calcN20()
    return turtle
}

func (t *TurtleStock) calcTr() {
    t.TR = make([]float64, len(t.Prices))
    for i := 0; i < len(t.Prices); i++ {
        if i == 0 {
            t.TR[i] = t.Prices[i].High - t.Prices[i].Low
        } else {
            v1 := t.Prices[i].High - t.Prices[i].Low;
            v2 := t.Prices[i].High - t.Prices[i - 1].Close;
            v3 := t.Prices[i - 1].Close - t.Prices[i].Low;
            t.TR[i] = math.Max(math.Max(v1, v2), v3);
        }
    }
}

func (t *TurtleStock) calcN20() {
    t.N20 = make([]float64, len(t.Prices))
    for i := 0; i < len(t.Prices); i++ {
        if i == 0 {
            t.N20[i] = t.TR[i]
        } else {
            t.N20[i] = (19 * t.N20[i - 1] + t.TR[i]) / 20;
        }
    }
}

func SharpeRatio(prices []yquotes.PriceH) (float64, error) {
    return 0, nil;
}

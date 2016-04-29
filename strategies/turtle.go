package strategies

import (
    "math"
    "github.com/wangchen1ren/stock-go/price"
    "github.com/wangchen1ren/stock-go/utils"
    "fmt"
)

type Turtle struct {
}

func (t *Turtle) Eval(symbol string, prices price.Prices) Metrics {
    tr := calcTr(prices)
    n20 := calcN20(tr)

    fmt.Printf("TR: %v", tr)
    fmt.Printf("N20: %v", n20)

    var metrics Metrics
    return metrics
}

func calcTr(prices price.Prices) []float64 {
    tr := make([]float64, len(prices))
    for i := 0; i < len(prices); i++ {
        if i == 0 {
            tr[i] = utils.Round(prices[i].High - prices[i].Low, 2)
        } else {
            v1 := prices[i].High - prices[i].Low;
            v2 := prices[i].High - prices[i - 1].Close;
            v3 := prices[i - 1].Close - prices[i].Low;
            tr[i] = utils.Round(math.Max(math.Max(v1, v2), v3), 2);
        }
    }
    return tr
}

func calcN20(tr []float64) []float64 {
    n20 := make([]float64, len(tr))
    for i := 0; i < len(tr); i++ {
        if i == 0 {
            n20[i] = tr[i]
        } else {
            n20[i] = utils.Round((19 * n20[i - 1] + tr[i]) / 20, 2);
        }
    }
    return n20
}

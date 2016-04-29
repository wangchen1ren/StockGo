package strategies

import (
    "github.com/doneland/yquotes"
    "github.com/wangchen1ren/stock-go/price"
)

type Metrics struct {
    ART         float64 // 平均年化收益率

                        // Risk
    MaxDD       float64 // 最大挫跌
    MaxDDD      int     // 最大挫跌期
    RR          float64 // Return rate
    RRStd       float64

                        // Composite
    SharpeRatio float64
    RRR         float64 // Robust risk/reward ratio
    MAR         float64
}

type Strategy interface {
    Eval(symbol string, price price.Prices) Metrics
}

func SharpeRatio(prices []yquotes.PriceH) (float64, error) {
    return 0, nil;
}

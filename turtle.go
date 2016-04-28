package stock

import (
    "github.com/doneland/yquotes"
)

func NewPrices(yps []yquotes.PriceH) (Prices) {
    prices := make(Prices, len(yps));
    for i, yp := range yps {
        p := Price{}
        //fmt.Printf("%d : %v\n", i, yp.Date);
        p.Original = yp;
        p.init();
        prices[i] = p;
    }
    return prices;
}

/**
func (prices Prices) tr() {
	for i, p := range prices {
		if i == 0 {
			p.TR = p.Price.High - p.Price.Low;
		} else {
			pd := prices[i - 1];
			v1 := p.Price.High - p.Price.Low;
			v2 := p.Price.High - pd.Price.Close;
			v3 := pd.Price.Close - p.Price.Low;
			p.TR = math.Max(math.Max(v1, v2), v3);
		}
		prices[i] = p;
	}
}

func (prices Prices) n20() {
	for i, p := range prices {
		if i == 0 {
			p.N20 = p.TR;
		} else {
			pd := prices[i - 1];
			p.N20 = (19 * pd.N20 + p.TR) / 20;
		}
		prices[i] = p;
	}
}
*/

func SharpeRatio(prices []yquotes.PriceH) (float64, error) {
    return 0, nil;
}

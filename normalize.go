package stock

import "github.com/doneland/yquotes"

func NormalizePrice(prices []yquotes.PriceH) ([]yquotes.PriceH) {
	for i := 0; i < len(prices); i++ {
		if prices[i].AdjClose != prices[i].Close {
			rate := prices[i].Close / prices[i].AdjClose;
			prices[i].Open /= rate;
			prices[i].High /= rate;
			prices[i].Low /= rate;
			prices[i].Close /= rate;
		}
	}
	return prices;
}


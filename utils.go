package stock

import (
	"fmt"
	"os"
	"github.com/doneland/yquotes"
)

func DumpPrices(prices []yquotes.PriceH) {
	for i := 0; i < len(prices); i++ {
		fmt.Fprintf(os.Stdout, "%+v\n", prices[i])
	}
}

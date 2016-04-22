package stock

import (
	"github.com/doneland/yquotes"
	"sort"
)

const (
	ASC = "asc"
	DESC = "desc"
)

type Prices struct {
	prices []yquotes.PriceH
	by     func(p1, p2 *yquotes.PriceH) bool // Closure used in the Less method.
}

func (p *Prices) Len() int {
	return len(p.prices)
}
func (p *Prices) Less(i, j int) bool {
	return p.by(&p.prices[i], &p.prices[j])
}
func (p *Prices) Swap(i, j int) {
	p.prices[i], p.prices[j] = p.prices[j], p.prices[i]
}

// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(p1, p2 *yquotes.PriceH) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(prices []yquotes.PriceH) {
	ps := &Prices{
		prices: prices,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

func SortPriceByDate(prices []yquotes.PriceH, order string) ([]yquotes.PriceH) {
	asc := func(p1, p2 *yquotes.PriceH) bool {
		return p1.Date.Before(p2.Date)
	}
	desc := func(p1, p2 *yquotes.PriceH) bool {
		return !asc(p1, p2)
	}
	if order == DESC {
		By(desc).Sort(prices);
	} else {
		By(asc).Sort(prices);
	}
	return prices;
}

package main

import (
	"github.com/doneland/yquotes"
	"fmt"
)

func main() {
	id := "002721.sz"
	/*
	ss, err := yquotes.NewStock(id, true);
	if err != nil {
		fmt.Println(err);
	}

	//sharperatio := stock.CalcSharpeRatio(ss.History);
	fmt.Println(ss.Name);
	fmt.Println(ss.Symbol);
	fmt.Println(ss.Price);
	for i := 0; i < len(ss.History); i++ {
		fmt.Println(ss.History[i]);
	}
	*/

	price, err := yquotes.GetPrice(id);
	if err != nil {
		fmt.Println(err);
	}
	fmt.Printf("%v", price);
	fmt.Println();
	fmt.Printf("%+v", price);
	fmt.Println();
	fmt.Printf("%#v", price);
	fmt.Println();
}

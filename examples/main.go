package main

import (
	"fmt"

	"github.com/relusc/gonvista"
)

func main() {
	c := gonvista.NewClientDefault()

	// search by WKN (Apple)
	i, err := c.SearchInstruments("AAPL")
	if err != nil {
		panic(err)
	}

	fmt.Println(i[0])

	// Request by ISIN (IShares Core MSCI World)
	msci, err := c.GetInstrumentByISIN("IE00B4L5Y983")
	if err != nil {
		panic(err)
	}

	// Get Notations of instrument
	nList, err := c.GetInstrumentNotations(msci)
	if err != nil {
		panic(err)
	}

	for _, n := range nList {
		fmt.Println(n)
	}

	// List quotes for specific instrument and exchange
	q, err := c.GetInstrumentQuotesJSON(msci, nList[0].Id, gonvista.RangeOneYear, "2022-11-20")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(q))
}

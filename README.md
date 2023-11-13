# gonvista

Simple Go package to communicate with the non-public [onvista.de](https://onvista.de) API.

## Disclaimer

The API is non-public and not documented at all. There are only a few snippets and hints to this API that one can find when browsing forums and threads on various platforms around the internet.

There is no guarantee that all endpoints are up to date at all times. However, it's most likely guaranteed that this package does not implement all of the available endpoints by far. The package has been created out of [personal interests](#intention). Consider this package as experimental.

Also, as the API is not public it should be considered that using this package breaks the website user agreements. Treat with caution.

## Name

go + onvista = gonvista :exploding_head: :exploding_head: :exploding_head:

## Intention

This package was implemented out of personal interest. The main point was to be able to add assets to [Portfolio Performance](https://www.portfolio-performance.info/) for which the historical quotes cannot be found in Portfolio Report itself or other big financial sites (e.g. Yahoo Finance). The package can be used to find the correct API URL that can be submitted into Portfolio Performance as source for historical quotes. The quotes returned by the API can be used as JSON input when adding a new asset.

## Installation

```shell
go get -d github.com/relusc/gonvista
```

## Example usage

```go
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
 q, err := c.GetInstrumentQuotesJSON(msci, nList[0].Id, "Y1", "2022-11-20")
 if err != nil {
  panic(err)
 }

 fmt.Println(string(q))
}
```

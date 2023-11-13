package gonvista

import "net/http"

const (
	// Base URL of the onvista.de API
	// There are also v2 endpoints available, but not for financial data
	apiBase = "https://api.onvista.de/api/v1"

	// supported ranges to get historical quotes
	// NOTE I: not all of them are tested, no guarantee that all of them work for all instruments
	// NOTE II: there might also be more ranges, however these are returned when requesting an instrument "snapshot"
	RangeOneDay      = "D1"
	RangeOneWeek     = "W1"
	RangeOneMonth    = "M1"
	RangeThreeMonths = "M3"
	RangeSixMonths   = "M6"
	RangeOneYear     = "Y1"
	RangeThreeYears  = "Y3"
	RangeFiveYears   = "Y5"
	RangeTenYears    = "Y10"
	RangeMax         = "MAX"
)

var (
	// mapping types of different instruments
	instrument_type_map = map[string]string{
		"BOND":  "bonds",
		"FUND":  "funds",
		"STOCK": "stocks",
	}
)

// Client represents the client to talk to the onvista.de API
type Client struct {
	httpClient *http.Client
}

// -------------------------------------------------------------------------------------------------------- //
// API "Objects"
// -------------------------------------------------------------------------------------------------------- //

// Instrument represents a single stock, fund or bond
type Instrument struct {
	EntityType       string   `json:"entityType"`
	EntityAttributes []string `json:"entityAttributes"`
	EntityValue      string   `json:"entityValue"` // seems to be an onvista.de internal ID of the instrument
	Name             string   `json:"name"`
	URL              struct {
		Website string `json:"WEBSITE"`
	} `json:"urls"`
	InstrumentType string `json:"instrumentType"`
	ISIN           string `json:"isin"`
	WKN            string `json:"wkn"`
	DisplayType    string `json:"displayType"`
	URLName        string `json:"urlName"`
	TinyName       string `json:"tinyName"`
}

// Notation represents a stock exchange/market on which the instrument is listed
type Notation struct {
	Id      int    `json:"idNotation"`
	Name    string `json:"name"`
	Code    string `json:"codeExchange"`
	Country string `json:"isoCountry"`
}

// QuoteList represents the current quotes of an instrument on all exchanges it is listed on
type QuoteList struct {
	Quotes []Quote `json:"list"`
}

// Quote represents a the current quote of an instrument on a specific exchange
type Quote struct {
	Ask      float32  `json:"ask"`
	Bid      float32  `json:"bid"`
	Unit     string   `json:"unitType"`
	Volume   float32  `json:"volume"`
	Notation Notation `json:"market"`
}

// -------------------------------------------------------------------------------------------------------- //
// API Response types
// -------------------------------------------------------------------------------------------------------- //

// SearchInstrumentsResponse represents the API response when searching for instruments
type SearchInstrumentsResponse struct {
	Expires     int          `json:"expires"`
	SearchValue string       `json:"searchValue"`
	Instruments []Instrument `json:"list"`
}

// SnapshotInstrumentResponse represents the API response when requesting a single instrument
// This is referenced as "snapshot" by the onvista.de API
type SnapshotInstrumentResponse struct {
	Expires   int       `json:"expires"`
	Type      string    `json:"type"`
	QuoteList QuoteList `json:"quoteList"`
}

// Package gonvista provides a simple Go client to talk with the non-public onvista.de API
//
// This package can be considered as experimental as the API is non-public and not documented at all.
package gonvista

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

// SearchInstruments finds all instruments based on the submitted searchKey
// Search can be done by submitting a name, ISIN or WKN
// Also parts of a name, ISIN or WKN work, but those searches of course return more results
func (c *Client) SearchInstruments(searchKey string) ([]Instrument, error) {
	url := fmt.Sprintf("%s/instruments/search?searchValue=%s", apiBase, searchKey)

	// do request
	r, err := c.doHTTP(url, http.MethodGet)
	if err != nil {
		return nil, err
	}

	// parse response
	var resp SearchInstrumentsResponse

	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, fmt.Errorf("error while parsing search instruments response: %s", err.Error())
	}

	return resp.Instruments, nil
}

// GetInstrumentByISIN finds a single instrument based on the submitted ISIN
func (c *Client) GetInstrumentByISIN(searchKey string) (Instrument, error) {
	url := fmt.Sprintf("%s/instruments/search?searchValue=%s", apiBase, searchKey)

	// do request
	r, err := c.doHTTP(url, http.MethodGet)
	if err != nil {
		return Instrument{}, err
	}

	// parse response
	var resp SearchInstrumentsResponse

	err = json.Unmarshal(r, &resp)
	if err != nil {
		return Instrument{}, fmt.Errorf("error while parsing search instruments response: %s", err.Error())
	}

	if len(resp.Instruments) > 1 {
		return Instrument{}, fmt.Errorf("search returned %d instruments, only one is expected; please update search string", len(resp.Instruments))
	}

	return resp.Instruments[0], nil
}

// GetInstrumentNotations returns the notations of a single instrument
// Notations can be found when requesting a "snapshot" from onvista.de, hence the API URL
func (c *Client) GetInstrumentNotations(i Instrument) ([]Notation, error) {
	// Set type of instrument (fund, bond etc.)
	instrument_type := instrument_type_map[i.EntityType]

	// create API URL
	url := fmt.Sprintf("%s/%s/ISIN:%s/snapshot", apiBase, instrument_type, i.ISIN)

	// do request
	r, err := c.doHTTP(url, http.MethodGet)
	if err != nil {
		return nil, err
	}

	// parse response
	var resp SnapshotInstrumentResponse

	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, fmt.Errorf("error while parsing search instruments response: %s", err.Error())
	}

	var notation []Notation
	for _, quote := range resp.QuoteList.Quotes {
		notation = append(notation, quote.Notation)
	}

	return notation, nil
}

// Gets historical quotes for an instrument, a specific range and a specific exchange
// Uses the onvista.de API endpoint "/eod_history" (also used on website itself)
//
// The response will be returned plain (in JSON) and can be used in e.g. Portfolio Performance as input
func (c *Client) GetInstrumentQuotesJSON(i Instrument, notationID int, quoteRange, startDate string) ([]byte, error) {
	// Check if provided startDate has correct format
	// Expected format: yyyy-mm-dd
	ok := regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`).MatchString(startDate)
	if !ok {
		return nil, fmt.Errorf("provided startDate %s does not match format 'yyyy-mm-dd'", startDate)
	}

	// create API URL
	url := fmt.Sprintf("%s/instruments/%s/%s/eod_history?idNotation=%d&range=%s&startDate=%s", apiBase, i.EntityType, i.EntityValue, notationID, quoteRange, startDate)

	// do request
	r, err := c.doHTTP(url, http.MethodGet)
	if err != nil {
		return nil, err
	}

	return r, nil
}

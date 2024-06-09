package ecb

import (
	"encoding/xml"
	"net/http"
)

const daily = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

type Ecb struct {
	Data []struct {
		Date  string `xml:"time,attr"`
		Rates []struct {
			Currency string `xml:"currency,attr"`
			Rate     string `xml:"rate,attr"`
		} `xml:"Cube"`
	} `xml:"Cube>Cube"`
}

// FetchDaily fetches the "raw" xml from the ECB website
// and returns a pointer to an Ecb struct.
func FetchDaily() (*Ecb, error) {
	resp, err := http.Get(daily)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ecb := new(Ecb)
	if err := xml.NewDecoder(resp.Body).Decode(ecb); err != nil {
		return nil, err
	}

	return ecb, nil
}

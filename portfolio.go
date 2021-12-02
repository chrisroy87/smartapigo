package smartapigo

import (
	"net/http"
)

// Holding is an individual holdings response.
type Holding struct {
	Tradingsymbol      string  `json:"tradingsymbol"`
	Exchange           string  `json:"exchange"`
	ISIN               string  `json:"isin"`
	T1Quantity         uint64  `json:"t1quantity"`
	RealisedQuantity   uint64  `json:"realisedquantity"`
	Quantity           uint64  `json:"quantity"`
	AuthorisedQuantity uint64  `json:"authorisedquantity"`
	ProfitAndLoss      float64 `json:"profitandloss"`
	Product            string  `json:"product"`
	CollateralQuantity string  `json:"collateralquantity"`
	CollateralType     string  `json:"collateraltype"`
	Haircut            uint64  `json:"haircut"`
	AveragePrice       float64 `json:"averageprice"`
	Ltp                float64 `json:"ltp"`
	SymbolToken        string  `json:"symboltoken"`
	Close              float64 `json:"close"`
}

// Holdings is a list of holdings
type Holdings []Holding

// GetHoldings gets a list of holdings.
func (c *Client) GetHoldings() (Holdings, error) {
	var holdings Holdings
	err := c.doEnvelope(http.MethodGet, URIGetHoldings, nil, nil, &holdings, true)
	return holdings, err
}

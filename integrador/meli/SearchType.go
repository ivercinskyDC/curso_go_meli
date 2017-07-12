package meli

//Meli is the main type
type Meli struct {
	SiteID string
}

//SearchParams its the all the params for the search call
type SearchParams struct {
	MethodID string
	SearchID string
	SortID   string
	FilterID string
	Offset   string
	Limit    string
}

//SearchResult type of Search response
type SearchResult struct {
	SiteID      string       `json:"site_id"`
	Paging      SearchPage   `json:"paging"`
	SearchItems []SearchItem `json:"results"`
}

//SearchPage data from Search Response
type SearchPage struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

//SearchItem from Search Response
type SearchItem struct {
	ID                 string       `json:"id"`
	SiteID             string       `json:"site_id"`
	Title              string       `json:"title"`
	Seller             ItemSeller   `json:"seller"`
	Price              float64      `json:"price"`
	CurrencyID         string       `json:"currency_id"`
	AvailableQuantity  int          `json:"available_quantity"`
	SoldQuantity       int          `json:"sold_quantity"`
	BuyingMode         string       `json:"buying_mode"`
	ListingTypeID      string       `json:"listing_type_id"`
	Condition          string       `json:"condition"`
	AcceptsMercadoPago bool         `json:"accepts_mercadopago"`
	Shipping           ItemShipping `json:"shipping"`
}

//ItemSeller is struct of the seller info
type ItemSeller struct {
	ID                int    `json:"id"`
	PowerSellerStatus string `json:"power_seller_status"`
}

//ItemShipping is the struct of the shipping Info
type ItemShipping struct {
	FreeShipping bool   `json:"free_shipping"`
	Mode         string `json:"mode"`
}

//Suggestion is the type of the response
type Suggestion struct {
	Max       float64 `json:"max"`
	Suggested float64 `json:"suggested"`
	Min       float64 `json:"min"`
}

package model

type StockDetail struct {
	ID             string `json:"id"`
	StockSymbol    string `json:"ticker"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Sector         string `json:"sector"`
	MarketCap      int64  `json:"marketCap"`
	MarketCapGroup string `json:"marketCapGroup"`
	WKN            string `json:"wkn"`
	ISIN           string `json:"isin"`
	Nation         string `json:"nation"`
}

type StockTransaction struct {
	Date        string `json:"date"`
	Price       string `json:"price"`
	Currency    string `json:"currency"`
	StockDetail StockDetail
}

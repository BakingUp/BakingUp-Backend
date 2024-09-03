package domain

type StockItem struct {
	StockName    string  `json:"stock_name"`
	StockURL     string  `json:"stock_url"`
	Quantity     int     `json:"quantity"`
	LST          int     `json:"lst"`
	SellingPrice float64 `json:"selling_price"`
	LSTStatus    string  `json:"lst_status"`
}

type StockList struct {
	Stocks []StockItem `json:"stocks"`
}

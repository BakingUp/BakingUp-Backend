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

type StockDetail struct {
	CreatedAt    string  `json:"created_at"`
	LSTStatus    string  `json:"lst_status"`
	Quantity	 int     `json:"quantity"`
	SellByDate   string  `json:"sell_by_date"`
}

type StockItemDetail struct {
	StockName    	string  		`json:"stock_name"`
	StockURL     	[]string  		`json:"stock_url"`
	Quantity     	int     		`json:"quantity"`
	LST          	int     		`json:"lst"`
	SellingPrice 	float64 		`json:"selling_price"`
	StockLessThan   int  			`json:"stock_less_than"`
	StockDetails    []StockDetail 	`json:"stock_details"`
}
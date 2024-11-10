package domain

import "time"

type UnreadNotification struct {
	UnreadNotiAmount int `json:"unread_noti_amount"`
}

type FilterSellingRequest struct {
	UserID        string    `json:"user_id"`
	FilterType    string    `json:"filter_type"`
	StartDateTime time.Time `json:"start_date_time,omitempty"`
	EndDateTime   time.Time `json:"end_date_time,omitempty"`
	SalesChannel  []string  `json:"sales_channel,omitempty"`
	OrderTypes    []string  `json:"order_types,omitempty"`
	UnitType      string    `json:"unit_type,omitempty"`
	SortType      string    `json:"sort_type,omitempty"`
}

type FilterWastedRequest struct {
	UserID     string `json:"user_id"`
	FilterType string `json:"filter_type"`
}

type FilterProductResponse struct {
	Products []FilterItemResponse `json:"products"`
}

type FilterItemResponse struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Detail string `json:"detail"`
}

type ProductPricing struct {
	SellingPrice float64 `json:"selling_price"`
	Cost         float64 `json:"cost"`
}

type DashboardChartDataResponse struct {
	CostRevenue     []CostRevenueChartItem     `json:"cost_revenue"`
	ProfitThreshold []ProfitThresholdChartItem `json:"profit_threshold"`
}

type CostRevenueChartItem struct {
	Month     string  `json:"month"`
	Revenue   float64 `json:"revenue"`
	Cost      float64 `json:"cost"`
	NetProfit float64 `json:"net_profit"`
}

type NetProfitChartItem struct {
	Month  string  `json:"month"`
	Profit float64 `json:"profit"`
}

type ProfitThresholdChartItem struct {
	Name      string  `json:"name"`
	Threshold float64 `json:"threshold"`
}

type StockBatchDetail struct {
	Time          time.Time `json:"time"`
	Quantity      int       `json:"quantity"`
	RecipeName    string    `json:"recipe_name"`
	StockDetailId string    `json:"stock_detail_id"`
	// OrderProductQuantity int       `json:"order_product_quantity"`
	CreatedAt time.Time `json:"created_at"`
}

type StockBatchList struct {
	StockID     string   `json:"stock_id"`
	StockBatchs []string `json:"stock_batchs"`
}

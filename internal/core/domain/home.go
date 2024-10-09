package domain

type UnreadNotification struct {
	UnreadNotiAmount int `json:"unread_noti_amount"`
}

type FilterSellingRequest struct {
	UserID       string   `json:"user_id"`
	FilterType   string   `json:"filter_type"`
	SalesChannel []string `json:"sales_channel,omitempty"`
	OrderTypes   []string `json:"order_types,omitempty"`
	UnitType     string   `json:"unit_type,omitempty"`
	SortType     string   `json:"sort_type,omitempty"`
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
	NetProfit       []NetProfitChartItem       `json:"net_profit"`
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

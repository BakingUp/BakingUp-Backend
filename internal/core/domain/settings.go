package domain

type UserLanguage struct {
	Language string `json:"language"`
}

type ChangeUserLanguage struct {
	UserID   string `json:"user_id"`
	Language string `json:"language"`
}

type FixCostSetting struct {
	Rent          float64 `json:"rent"`
	Salaries      float64 `json:"salaries"`
	Insurance     float64 `json:"insurance"`
	Subscriptions float64 `json:"subscriptions"`
	Advertising   float64 `json:"advertising"`
	Electricity   float64 `json:"electricity"`
	Water         float64 `json:"water"`
	Gas           float64 `json:"gas"`
	Other         float64 `json:"other"`
	Note          string  `json:"note"`
}

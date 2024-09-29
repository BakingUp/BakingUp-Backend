package domain

type UserLanguage struct {
	Language string `json:"language"`
}

type ChangeUserLanguage struct {
	UserID   string `json:"user_id"`
	Language string `json:"language"`
}

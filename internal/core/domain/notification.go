package domain

type NotificationItem struct {
	NotiID       string `json:"noti_id"`
	Title        string `json:"title"`
	Message      string `json:"message"`
	CreatedAt    string `json:"created_at"`
	IsRead       bool   `json:"is_read"`
	NotiType     string `json:"noti_type"`
	ItemID       string `json:"item_id"`
	ItemName     string `json:"item_name"`
	NotiItemType string `json:"noti_item_type"`
}

type NotificationList struct {
	Notis []NotificationItem `json:"notis"`
}

type CreateNotificationItem struct {
	UserID       string `json:"user_id"`
	ThaiTitle    string `json:"thai_title"`
	EngTitle     string `json:"eng_title"`
	ThaiMessage  string `json:"thai_message"`
	EngMessage   string `json:"eng_message"`
	IsRead       bool   `json:"is_read"`
	NotiType     string `json:"noti_type"`
	ItemID       string `json:"item_id"`
	ItemName     string `json:"item_name"`
	NotiItemType string `json:"noti_item_type"`
}

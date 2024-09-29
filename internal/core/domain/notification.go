package domain

type NotificationItem struct {
	Title     string `json:"title"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
	IsRead    bool   `json:"is_read"`
	NotiType  string `json:"noti_type"`
}

type NotificationList struct {
	Notis []NotificationItem `json:"notis"`
}

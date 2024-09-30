package util

import "github.com/BakingUp/BakingUp-Backend/prisma/db"

func GetNotificationTitle(notification *db.NotificationsModel, language *db.Language) string {
	if *language == db.LanguageTh {
		return notification.ThaiTitle
	}

	return notification.EngTitle
}

func GetNotificationMessage(notification *db.NotificationsModel, language *db.Language) string {
	if *language == db.LanguageTh {
		return notification.ThaiMessage
	}

	return notification.EngMessage
}

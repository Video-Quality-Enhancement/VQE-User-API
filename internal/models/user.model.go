package models

import "time"

type User struct {
	UserId                 string    `json:"userId" bson:"userId"`
	NotificationInterfaces []string  `json:"notificationInterfaces" bson:"notificationInterfaces"`
	FCMtokens              []string  `json:"FCMtokens" bson:"FCMtokens"`
	WhatsAppNumber         string    `json:"whatsAppNumber" bson:"whatsAppNumber"`
	DiscordId              string    `json:"discordId" bson:"discordId"`
	TelegramNumber         string    `json:"telegramNumber" bson:"telegramNumber"`
	Webhooks               []string  `json:"webhooks" bson:"webhooks"`
	CreatedAt              time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt" bson:"updatedAt"`
}

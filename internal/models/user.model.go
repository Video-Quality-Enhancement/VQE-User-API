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

type WhatsAppRequest struct {
	WhatsAppNumber string `json:"whatsAppNumber" bson:"whatsAppNumber" binding:"required"`
}

type DiscordRequest struct {
	DiscordId string `json:"discordId" bson:"discordId" binding:"required"`
}

type TelegramRequest struct {
	TelegramNumber string `json:"telegramNumber" bson:"telegramNumber" binding:"required"`
}

type NotificationInterfacesRequest struct {
	NotificationInterfaces []string `json:"notificationInterfaces" bson:"notificationInterfaces" binding:"required"`
}

type FCMtokensRequest struct {
	FCMtokens []string `json:"FCMtokens" bson:"FCMtokens" binding:"required"`
}

type WebhooksRequest struct {
	Webhooks []string `json:"webhooks" bson:"webhooks" binding:"required"`
}

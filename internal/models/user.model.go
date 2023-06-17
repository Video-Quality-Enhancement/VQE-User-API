package models

import "time"

type User struct {
	UserId                 string    `json:"userId" bson:"userId"`
	NotificationInterfaces []string  `json:"notificationInterfaces,omitempty" bson:"notificationInterfaces,omitempty"`
	FCMtokens              []string  `json:"fcmTokens,omitempty" bson:"fcmTokens,omitempty"`
	WhatsAppNumber         string    `json:"whatsAppNumber,omitempty" bson:"whatsAppNumber,omitempty"`
	DiscordId              string    `json:"discordId,omitempty" bson:"discordId,omitempty"`
	TelegramNumber         string    `json:"telegramNumber,omitempty" bson:"telegramNumber,omitempty"`
	Webhooks               []string  `json:"webhooks,omitempty" bson:"webhooks,omitempty"`
	CreatedAt              time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt" bson:"updatedAt"`
}

type WhatsAppRequest struct {
	WhatsAppNumber string `json:"whatsAppNumber" bson:"whatsAppNumber" binding:"required,e164"`
}

type DiscordRequest struct {
	DiscordId string `json:"discordId" bson:"discordId" binding:"required"`
}

type TelegramRequest struct {
	TelegramNumber string `json:"telegramNumber" bson:"telegramNumber" binding:"required,e164"`
}

type NotificationInterfacesRequest struct {
	NotificationInterfaces []string `json:"notificationInterfaces" bson:"notificationInterfaces" binding:"required,are-notification-interfaces-valid"`
}

type FCMtokenRequest struct {
	FCMtoken string `json:"fcmTokens" bson:"fcmTokens" binding:"required"`
}

type FCMtokensRequest struct {
	FCMtokens []string `json:"fcmTokens" bson:"fcmTokens"`
}

type WebhooksRequest struct {
	Webhooks []string `json:"webhooks" bson:"webhooks" binding:"required,are-webhooks-valid"`
}

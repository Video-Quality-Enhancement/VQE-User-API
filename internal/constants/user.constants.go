package constants

type NotificationInterface string

const (
	UI       NotificationInterface = "ui"
	Email    NotificationInterface = "email"
	Whatsapp NotificationInterface = "whatsapp"
	Discord  NotificationInterface = "discord"
	Telegram NotificationInterface = "telegram"
	Webhooks NotificationInterface = "webhooks"
)

func (r NotificationInterface) String() string {
	return string(r)
}

func GetNotificationInterfaces() [4]NotificationInterface {
	return [...]NotificationInterface{Email, UI, Whatsapp, Discord}
}

func GetNotificationInterfaceSet() map[NotificationInterface]struct{} {
	notificationInterfaces := GetNotificationInterfaces()
	notificationInterfacesSet := make(map[NotificationInterface]struct{})
	for _, notificationInterface := range notificationInterfaces {
		notificationInterfacesSet[notificationInterface] = struct{}{}
	}
	return notificationInterfacesSet
}

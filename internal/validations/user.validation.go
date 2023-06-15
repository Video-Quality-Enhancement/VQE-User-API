package validations

import (
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/constants"
	validator "github.com/go-playground/validator/v10"
	"golang.org/x/exp/slog"
)

func ValidateNotificationInterfaces(fl validator.FieldLevel) bool {
	notificationInterfaceSet := constants.GetNotificationInterfaceSet()
	notificationInterfaces := fl.Field().Interface().([]string)
	for _, notificationInterface := range notificationInterfaces {
		if _, ok := notificationInterfaceSet[constants.NotificationInterface(notificationInterface)]; !ok {
			slog.Error("Invalid notification interface", "notificationInterface", notificationInterface)
			return false
		}
	}
	return true
}

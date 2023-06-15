package validations

import (
	"net/http"

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

func ValidateWebhooks(fl validator.FieldLevel) bool {
	webhooks := fl.Field().Interface().([]string)
	for _, webhook := range webhooks {

		// Create a new request
		req, err := http.NewRequest("POST", webhook, nil)
		if err != nil {
			slog.Error("Invalid webhook", "webhook", webhook)
			return false
		}

		// Send the request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			slog.Error("Invalid webhook", "webhook", webhook)
			return false
		}

		if resp.StatusCode != http.StatusOK {
			slog.Error("Invalid webhook", "webhook", webhook)
			return false
		}
	}
	return true
}

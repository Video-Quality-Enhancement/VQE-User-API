package validations

import (
	"github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

func RegisterUserValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("are-notification-interfaces-valid", ValidateNotificationInterfaces)
	}
}

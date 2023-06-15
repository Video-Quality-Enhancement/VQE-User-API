package routes

import (
	"net/http"

	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/controllers"
	"github.com/gin-gonic/gin"
)

func testController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API call test successful",
	})
}

func RegisterUserRoutes(router *gin.RouterGroup, authorization gin.HandlerFunc, controller controllers.UserController) {

	router.Use(authorization)
	router.GET("/test", testController)

	router.PUT("/", controller.UpsertUser)
	router.GET("/", controller.GetUser)

	router.PUT("/whatsapp", controller.EditWhatsAppNumber)
	router.GET("/whatsapp", controller.GetWhatsAppNumber)

	router.PUT("/discord", controller.EditDiscordId)
	router.GET("/discord", controller.GetDiscordId)

	router.PUT("/telegram", controller.EditTelegramNumber)
	router.GET("/telegram", controller.GetTelegramNumber)

	router.PUT("/notificationInterfaces", controller.EditNotificationInterfaces)
	router.GET("/notificationInterfaces", controller.GetNotificationInterfaces)

	router.PUT("/fcmTokens", controller.EditFCMtokens)
	router.GET("/fcmTokens", controller.GetFCMtokens)

	router.PUT("/webhooks", controller.EditWebhooks)
	router.GET("/webhooks", controller.GetWebhooks)

	router.DELETE("/", controller.DeleteUser)

}

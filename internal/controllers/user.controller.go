package controllers

import (
	"net/http"

	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/services"
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	UpsertUser(c *gin.Context)
	GetUser(c *gin.Context)

	EditWhatsAppNumber(c *gin.Context)
	GetWhatsAppNumber(c *gin.Context)

	EditDiscordId(c *gin.Context)
	GetDiscordId(c *gin.Context)

	EditTelegramNumber(c *gin.Context)
	GetTelegramNumber(c *gin.Context)

	EditNotificationInterfaces(c *gin.Context)
	GetNotificationInterfaces(c *gin.Context)

	AddFCMtoken(c *gin.Context)
	DeleteFCMtoken(c *gin.Context)
	GetFCMtokens(c *gin.Context)

	EditWebhooks(c *gin.Context)
	GetWebhooks(c *gin.Context)

	DeleteUser(c *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (controller *userController) UpsertUser(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isUpserted, err := controller.userService.UpsertUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if isUpserted {
		c.JSON(http.StatusCreated, gin.H{"isUpserted": isUpserted})
		return
	}

	c.JSON(http.StatusOK, gin.H{"isUpserted": isUpserted})
}

func (controller *userController) GetUser(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := controller.userService.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller *userController) EditWhatsAppNumber(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var whatsAppRequest models.WhatsAppRequest
	err = c.ShouldBindJSON(&whatsAppRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.userService.EditWhatsAppNumber(userId, whatsAppRequest.WhatsAppNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, whatsAppRequest)
}

func (controller *userController) GetWhatsAppNumber(c *gin.Context) {

	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	whatsAppNumber, err := controller.userService.GetWhatsAppNumber(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.WhatsAppRequest{WhatsAppNumber: whatsAppNumber})

}

func (controller *userController) EditDiscordId(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var discordRequest models.DiscordRequest
	err = c.ShouldBindJSON(&discordRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.userService.EditDiscordId(userId, discordRequest.DiscordId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, discordRequest)
}

func (controller *userController) GetDiscordId(c *gin.Context) {

	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	discordId, err := controller.userService.GetDiscordId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.DiscordRequest{DiscordId: discordId})

}

func (controller *userController) EditTelegramNumber(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var telegramRequest models.TelegramRequest
	err = c.ShouldBindJSON(&telegramRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.userService.EditTelegramNumber(userId, telegramRequest.TelegramNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, telegramRequest)
}

func (controller *userController) GetTelegramNumber(c *gin.Context) {

	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	telegramNumber, err := controller.userService.GetTelegramNumber(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.TelegramRequest{TelegramNumber: telegramNumber})

}

func (controller *userController) EditNotificationInterfaces(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var notificationInterfacesRequest models.NotificationInterfacesRequest
	err = c.ShouldBindJSON(&notificationInterfacesRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.userService.EditNotificationInterfaces(userId, notificationInterfacesRequest.NotificationInterfaces)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notificationInterfacesRequest)
}

func (controller *userController) GetNotificationInterfaces(c *gin.Context) {

	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notificationInterfaces, err := controller.userService.GetNotificationInterfaces(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.NotificationInterfacesRequest{NotificationInterfaces: notificationInterfaces})

}

func (controller *userController) AddFCMtoken(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fcmTokensRequest models.FCMtokenRequest
	err = c.ShouldBindJSON(&fcmTokensRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.userService.AddFCMtoken(userId, fcmTokensRequest.FCMtoken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fcmTokensRequest)
}

func (controller *userController) DeleteFCMtoken(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fcmTokensRequest models.FCMtokenRequest
	err = c.ShouldBindJSON(&fcmTokensRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.userService.DeleteFCMtoken(userId, fcmTokensRequest.FCMtoken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fcmTokensRequest)
}

func (controller *userController) GetFCMtokens(c *gin.Context) {

	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fcmTokens, err := controller.userService.GetFCMtokens(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.FCMtokensRequest{FCMtokens: fcmTokens})

}

func (controller *userController) EditWebhooks(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var webhooksRequest models.WebhooksRequest
	err = c.ShouldBindJSON(&webhooksRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.userService.EditWebhooks(userId, webhooksRequest.Webhooks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, webhooksRequest)
}

func (controller *userController) GetWebhooks(c *gin.Context) {

	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webhooks, err := controller.userService.GetWebhooks(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.WebhooksRequest{Webhooks: webhooks})

}

func (controller *userController) DeleteUser(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.userService.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

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
	EditDiscordId(c *gin.Context)
	EditTelegramNumber(c *gin.Context)
	EditNotificationInterfaces(c *gin.Context)
	EditFCMtokens(c *gin.Context)
	EditWebhooks(c *gin.Context)
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

func (controller *userController) EditFCMtokens(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fcmTokensRequest models.FCMtokensRequest
	err = c.ShouldBindJSON(&fcmTokensRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.userService.EditFCMtokens(userId, fcmTokensRequest.FCMtokens)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fcmTokensRequest)
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

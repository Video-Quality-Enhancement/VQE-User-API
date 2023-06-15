package services

import (
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/producers"
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/repositories"
	"golang.org/x/exp/slog"
)

type UserService interface {
	UpsertUser(userId string) (bool, error)
	GetUser(userId string) (*models.User, error)
	EditWhatsAppNumber(userId string, whatsAppNumber string) error
	EditDiscordId(userId string, discordId string) error
	EditTelegramNumber(userId string, telegramNumber string) error
	EditNotificationInterfaces(userId string, notificationInterfaces []string) error
	EditFCMtokens(userId string, FCMtokens []string) error
	EditWebhooks(userId string, webhooks []string) error
	DeleteUser(userId string) error
}

type userService struct {
	userRepository repositories.UserRepository
	producer       producers.WelcomeProducer
}

func NewUserService(userRepository repositories.UserRepository, producer producers.WelcomeProducer) UserService {
	return &userService{
		userRepository: userRepository,
		producer:       producer,
	}
}

func (s *userService) UpsertUser(userId string) (bool, error) {

	user := &models.User{
		UserId: userId,
	}
	isUpserted, err := s.userRepository.Upsert(user)
	if err != nil {
		slog.Error("Failed to upsert user", "error", err, "userId", userId)
		return false, err
	}

	if isUpserted {
		err := s.producer.Publish(userId)
		if err != nil {
			slog.Error("Failed to publish welcome message", "error", err, "userId", userId)
			return false, err
		}
	}

	slog.Debug("Upserted user", "userId", userId, "isUpserted", isUpserted)
	return isUpserted, nil

}

func (s *userService) GetUser(userId string) (*models.User, error) {

	user, err := s.userRepository.FindByUserId(userId)
	if err != nil {
		slog.Error("Failed to get user by userId", "error", err, "userId", userId)
		return nil, err
	}

	slog.Debug("Got user by userId", "userId", userId)
	return user, nil

}

func (s *userService) EditWhatsAppNumber(userId string, whatsAppNumber string) error {

	err := s.userRepository.UpdateWhatsAppNumber(userId, whatsAppNumber)
	if err != nil {
		slog.Error("Failed to edit WhatsApp number", "error", err, "userId", userId, "whatsAppNumber", whatsAppNumber)
		return err
	}

	slog.Debug("Edited WhatsApp number", "userId", userId, "whatsAppNumber", whatsAppNumber)
	return nil

}

func (s *userService) EditDiscordId(userId string, discordId string) error {

	err := s.userRepository.UpdateDiscordId(userId, discordId)
	if err != nil {
		slog.Error("Failed to edit Discord ID", "error", err, "userId", userId, "discordId", discordId)
		return err
	}

	slog.Debug("Edited Discord ID", "userId", userId, "discordId", discordId)
	return nil

}

func (s *userService) EditTelegramNumber(userId string, telegramNumber string) error {

	err := s.userRepository.UpdateTelegramNumber(userId, telegramNumber)
	if err != nil {
		slog.Error("Failed to edit Telegram number", "error", err, "userId", userId, "telegramNumber", telegramNumber)
		return err
	}

	slog.Debug("Edited Telegram number", "userId", userId, "telegramNumber", telegramNumber)
	return nil

}

func (s *userService) EditNotificationInterfaces(userId string, notificationInterfaces []string) error {

	err := s.userRepository.UpdateNotificationInterfaces(userId, notificationInterfaces)
	if err != nil {
		slog.Error("Failed to edit Telegram number", "error", err, "userId", userId, "notificationInterfaces", notificationInterfaces)
		return err
	}

	slog.Debug("Edited Telegram number", "userId", userId, "notificationInterfaces", notificationInterfaces)
	return nil

}

func (s *userService) EditFCMtokens(userId string, FCMtokens []string) error {

	err := s.userRepository.UpdateFCMtokens(userId, FCMtokens)
	if err != nil {
		slog.Error("Failed to edit FCM tokens", "error", err, "userId", userId, "FCMtokens", FCMtokens)
		return err
	}

	slog.Debug("Edited FCM tokens", "userId", userId, "FCMtokens", FCMtokens)
	return nil

}

func (s *userService) EditWebhooks(userId string, webhooks []string) error {

	err := s.userRepository.UpdateWebhooks(userId, webhooks)
	if err != nil {
		slog.Error("Failed to edit Telegram number", "error", err, "userId", userId, "webhooks", webhooks)
		return err
	}

	slog.Debug("Edited Telegram number", "userId", userId, "webhooks", webhooks)
	return nil

}

func (s *userService) DeleteUser(userId string) error {

	err := s.userRepository.Delete(userId)
	if err != nil {
		slog.Error("Failed to delete user", "error", err, "userId", userId)
		return err
	}

	slog.Debug("Deleted user", "userId", userId)
	return nil

}

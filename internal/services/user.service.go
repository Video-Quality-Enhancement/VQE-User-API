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
	GetWhatsAppNumber(userId string) (string, error)

	EditDiscordId(userId string, discordId string) error
	GetDiscordId(userId string) (string, error)

	EditTelegramNumber(userId string, telegramNumber string) error
	GetTelegramNumber(userId string) (string, error)

	EditNotificationInterfaces(userId string, notificationInterfaces []string) error
	GetNotificationInterfaces(userId string) ([]string, error)

	EditFCMtokens(userId string, FCMtokens []string) error
	GetFCMtokens(userId string) ([]string, error)

	EditWebhooks(userId string, webhooks []string) error
	GetWebhooks(userId string) ([]string, error)

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
		slog.Error("Failed to get user", "error", err, "userId", userId)
		return nil, err
	}

	slog.Debug("Got user", "userId", userId)
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

func (s *userService) GetWhatsAppNumber(userId string) (string, error) {

	whatsAppNumber, err := s.userRepository.FindWhatsAppNumber(userId)
	if err != nil {
		slog.Error("Failed to get WhatsApp number", "error", err, "userId", userId)
		return "", err
	}

	slog.Debug("Got WhatsApp number", "userId", userId)
	return whatsAppNumber, nil

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

func (s *userService) GetDiscordId(userId string) (string, error) {

	discordId, err := s.userRepository.FindDiscordId(userId)
	if err != nil {
		slog.Error("Failed to get Discord ID", "error", err, "userId", userId)
		return "", err
	}

	slog.Debug("Got Discord ID", "userId", userId)
	return discordId, nil

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

func (s *userService) GetTelegramNumber(userId string) (string, error) {

	telegramNumber, err := s.userRepository.FindTelegramNumber(userId)
	if err != nil {
		slog.Error("Failed to get Telegram number", "error", err, "userId", userId)
		return "", err
	}

	slog.Debug("Got Telegram number", "userId", userId)
	return telegramNumber, nil

}

func (s *userService) EditNotificationInterfaces(userId string, notificationInterfaces []string) error {

	err := s.userRepository.UpdateNotificationInterfaces(userId, notificationInterfaces)
	if err != nil {
		slog.Error("Failed to edit notification Interfaces", "error", err, "userId", userId, "notificationInterfaces", notificationInterfaces)
		return err
	}

	slog.Debug("Edited notification Interfaces", "userId", userId, "notificationInterfaces", notificationInterfaces)
	return nil

}

func (s *userService) GetNotificationInterfaces(userId string) ([]string, error) {

	notificationInterfaces, err := s.userRepository.FindNotificationInterfaces(userId)
	if err != nil {
		slog.Error("Failed to get notification Interfaces", "error", err, "userId", userId)
		return nil, err
	}

	slog.Debug("Got notification Interfaces", "userId", userId)
	return notificationInterfaces, nil

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

func (s *userService) GetFCMtokens(userId string) ([]string, error) {

	FCMtokens, err := s.userRepository.FindFCMtokens(userId)
	if err != nil {
		slog.Error("Failed to get FCM tokens", "error", err, "userId", userId)
		return nil, err
	}

	slog.Debug("Got FCM tokens", "userId", userId)
	return FCMtokens, nil

}

func (s *userService) EditWebhooks(userId string, webhooks []string) error {

	err := s.userRepository.UpdateWebhooks(userId, webhooks)
	if err != nil {
		slog.Error("Failed to edit Webhooks", "error", err, "userId", userId, "webhooks", webhooks)
		return err
	}

	slog.Debug("Edited Webhooks", "userId", userId, "webhooks", webhooks)
	return nil

}

func (s *userService) GetWebhooks(userId string) ([]string, error) {

	webhooks, err := s.userRepository.FindWebhooks(userId)
	if err != nil {
		slog.Error("Failed to get Webhooks", "error", err, "userId", userId)
		return nil, err
	}

	slog.Debug("Got Webhooks", "userId", userId)
	return webhooks, nil

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

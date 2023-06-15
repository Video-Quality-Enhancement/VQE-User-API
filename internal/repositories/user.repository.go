package repositories

import (
	"context"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slog"
)

type UserRepository interface {
	Upsert(user *models.User) (bool, error)

	FindByUserId(userId string) (*models.User, error)

	UpdateWhatsAppNumber(userId string, whatsAppNumber string) error
	FindWhatsAppNumber(userId string) (string, error)

	UpdateDiscordId(userId string, discordId string) error
	FindDiscordId(userId string) (string, error)

	UpdateTelegramNumber(userId string, telegramNumber string) error
	FindTelegramNumber(userId string) (string, error)

	UpdateNotificationInterfaces(userId string, notificationInterfaces []string) error
	FindNotificationInterfaces(userId string) ([]string, error)

	UpdateFCMtokens(userId string, FCMtokens []string) error
	FindFCMtokens(userId string) ([]string, error)

	UpdateWebhooks(userId string, webhooks []string) error
	FindWebhooks(userId string) ([]string, error)

	Delete(userId string) error
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &userRepository{
		collection: collection,
	}
}

func (r *userRepository) Upsert(user *models.User) (bool, error) {
	now := time.Now().UTC()
	user.CreatedAt = now
	user.UpdatedAt = now

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": user.UserId}
	update := bson.D{{
		Key: "$set",
		Value: bson.M{
			"userId":    user.UserId,
			"createdAt": user.CreatedAt,
			"updatedAt": user.UpdatedAt,
		},
	}}
	opts := options.Update().SetUpsert(true)

	updatedResult, err := r.collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		slog.Error("Failed to upsert", "error", err, "userId", user.UserId)
		return false, err
	}

	slog.Debug("Upserted", "userId", user.UserId, "updatedResult", updatedResult)

	if updatedResult.UpsertedCount == 1 {
		return true, nil
	}

	return false, nil
}

func (r *userRepository) FindByUserId(userId string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}

	var user models.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		slog.Error("Failed to find user", "error", err, "userId", userId)
		return nil, err
	}

	slog.Debug("Found user", "userId", userId, "user", user)
	return &user, nil
}

func (r *userRepository) UpdateWhatsAppNumber(userId string, whatsAppNumber string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	update := bson.D{{
		Key: "$set",
		Value: bson.M{
			"whatsAppNumber": whatsAppNumber,
			"updatedAt":      time.Now().UTC(),
		},
	}}

	updatedResult, err := r.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		slog.Error("Failed to update WhatsApp number", "error", err, "userId", userId, "whatsAppNumber", whatsAppNumber)
		return err
	}

	slog.Debug("Updated WhatsApp number", "userId", userId, "whatsAppNumber", whatsAppNumber, "updatedResult", updatedResult)
	return nil
}

func (r *userRepository) FindWhatsAppNumber(userId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	opts := options.FindOne().SetProjection(bson.M{"whatsAppNumber": 1})

	var user models.User
	err := r.collection.FindOne(ctx, filter, opts).Decode(&user)

	if err != nil {
		slog.Error("Failed to find WhatsApp number", "error", err, "userId", userId)
		return "", err
	}

	slog.Debug("Found WhatsApp number", "userId", userId, "whatsAppNumber", user.WhatsAppNumber)
	return user.WhatsAppNumber, nil
}

func (r *userRepository) UpdateDiscordId(userId string, discordId string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	update := bson.D{{
		Key: "$set",
		Value: bson.M{
			"discordId": discordId,
			"updatedAt": time.Now().UTC(),
		},
	}}

	updatedResult, err := r.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		slog.Error("Failed to update Discord ID", "error", err, "userId", userId, "discordId", discordId)
		return err
	}

	slog.Debug("Updated Discord ID", "userId", userId, "discordId", discordId, "updatedResult", updatedResult)
	return nil
}

func (r *userRepository) FindDiscordId(userId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	opts := options.FindOne().SetProjection(bson.M{"discordId": 1})

	var user models.User
	err := r.collection.FindOne(ctx, filter, opts).Decode(&user)

	if err != nil {
		slog.Error("Failed to find Discord ID", "error", err, "userId", userId)
		return "", err
	}

	slog.Debug("Found Discord ID", "userId", userId, "discordId", user.DiscordId)
	return user.DiscordId, nil
}

func (r *userRepository) UpdateTelegramNumber(userId string, telegramNumber string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	update := bson.D{{
		Key: "$set",
		Value: bson.M{
			"telegramNumber": telegramNumber,
			"updatedAt":      time.Now().UTC(),
		},
	}}

	updatedResult, err := r.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		slog.Error("Failed to update WhatsApp number", "error", err, "userId", userId, "telegramNumber", telegramNumber)
		return err
	}

	slog.Debug("Updated WhatsApp number", "userId", userId, "telegramNumber", telegramNumber, "updatedResult", updatedResult)
	return nil

}

func (r *userRepository) FindTelegramNumber(userId string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	opts := options.FindOne().SetProjection(bson.M{"telegramNumber": 1})

	var user models.User
	err := r.collection.FindOne(ctx, filter, opts).Decode(&user)

	if err != nil {
		slog.Error("Failed to find Telegram number", "error", err, "userId", userId)
		return "", err
	}

	slog.Debug("Found Telegram number", "userId", userId, "telegramNumber", user.TelegramNumber)
	return user.TelegramNumber, nil

}

func (r *userRepository) UpdateNotificationInterfaces(userId string, notificationInterfaces []string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	update := bson.D{{
		Key: "$set",
		Value: bson.M{
			"notificationInterfaces": notificationInterfaces,
			"updatedAt":              time.Now().UTC(),
		},
	}}

	updatedResult, err := r.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		slog.Error("Failed to update WhatsApp number", "error", err, "userId", userId, "notificationInterfaces", notificationInterfaces)
		return err
	}

	slog.Debug("Updated WhatsApp number", "userId", userId, "notificationInterfaces", notificationInterfaces, "updatedResult", updatedResult)
	return nil

}

func (r *userRepository) FindNotificationInterfaces(userId string) ([]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	opts := options.FindOne().SetProjection(bson.M{"notificationInterfaces": 1})

	var user models.User
	err := r.collection.FindOne(ctx, filter, opts).Decode(&user)

	if err != nil {
		slog.Error("Failed to find Notification Interfaces", "error", err, "userId", userId)
		return nil, err
	}

	slog.Debug("Found Notification Interfaces", "userId", userId, "notificationInterfaces", user.NotificationInterfaces)
	return user.NotificationInterfaces, nil

}

func (r *userRepository) UpdateFCMtokens(userId string, FCMtokens []string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	update := bson.D{{
		Key: "$set",
		Value: bson.M{
			"FCMtokens": FCMtokens,
			"updatedAt": time.Now().UTC(),
		},
	}}

	updatedResult, err := r.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		slog.Error("Failed to update WhatsApp number", "error", err, "userId", userId, "FCMtokens", FCMtokens)
		return err
	}

	slog.Debug("Updated WhatsApp number", "userId", userId, "FCMtokens", FCMtokens, "updatedResult", updatedResult)
	return nil

}

func (r *userRepository) FindFCMtokens(userId string) ([]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	opts := options.FindOne().SetProjection(bson.M{"FCMtokens": 1})

	var user models.User
	err := r.collection.FindOne(ctx, filter, opts).Decode(&user)

	if err != nil {
		slog.Error("Failed to find FCM tokens", "error", err, "userId", userId)
		return nil, err
	}

	slog.Debug("Found FCM tokens", "userId", userId, "FCMtokens", user.FCMtokens)
	return user.FCMtokens, nil

}

func (r *userRepository) UpdateWebhooks(userId string, webhooks []string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	update := bson.D{{
		Key: "$set",
		Value: bson.M{
			"webhooks":  webhooks,
			"updatedAt": time.Now().UTC(),
		},
	}}

	updatedResult, err := r.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		slog.Error("Failed to update WhatsApp number", "error", err, "userId", userId, "webhooks", webhooks)
		return err
	}

	slog.Debug("Updated WhatsApp number", "userId", userId, "webhooks", webhooks, "updatedResult", updatedResult)
	return nil

}

func (r *userRepository) FindWebhooks(userId string) ([]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}
	opts := options.FindOne().SetProjection(bson.M{"webhooks": 1})

	var user models.User
	err := r.collection.FindOne(ctx, filter, opts).Decode(&user)

	if err != nil {
		slog.Error("Failed to find Webhooks", "error", err, "userId", userId)
		return nil, err
	}

	slog.Debug("Found Webhooks", "userId", userId, "webhooks", user.Webhooks)
	return user.Webhooks, nil

}

func (r *userRepository) Delete(userId string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}

	deletedResult, err := r.collection.DeleteOne(ctx, filter)

	if err != nil {
		slog.Error("Failed to delete", "error", err, "userId", userId)
		return err
	}

	slog.Debug("Deleted", "userId", userId, "deletedResult", deletedResult)
	return nil

}

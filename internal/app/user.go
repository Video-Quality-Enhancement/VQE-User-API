package app

import (
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/controllers"
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/middlewares"
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/producers"
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/repositories"
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/routes"
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/services"
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/validations"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpUser(router *gin.RouterGroup, collection *mongo.Collection, conn config.AMQPconnection, firebaseClient config.FirebaseClient) {

	repository := repositories.NewUserRepository(collection)
	producer := producers.NewWelcomeProducer(conn)
	service := services.NewUserService(repository, producer)
	controller := controllers.NewUserController(service)
	validations.RegisterUserValidations()
	authorization := middlewares.Authorization(firebaseClient)
	routes.RegisterUserRoutes(router, authorization, controller)

}

func SetUpUserRepositoryIndexes(collection *mongo.Collection) {

	repository := repositories.NewUserRepositorySetup(collection)
	repository.MakeUserIdUniqueIndex()

}

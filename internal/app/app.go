package app

import (
	"os"

	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpApp(router *gin.Engine, database *mongo.Database, conn config.AMQPconnection, firebaseClient config.FirebaseClient) {

	collection := database.Collection(os.Getenv("USER_COLLECTION"))
	userRouter := router.Group("/api/user")

	SetUpUser(userRouter, collection, conn, firebaseClient)
	// SetUpUserVideo(userVideoRouter, collection, conn, firebaseClient)

}

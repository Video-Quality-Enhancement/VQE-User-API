package app

import (
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpUser(router *gin.RouterGroup, collection *mongo.Collection, conn config.AMQPconnection, firebaseClient config.FirebaseClient) {

}

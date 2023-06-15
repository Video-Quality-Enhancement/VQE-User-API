package main

import (
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/app"
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-User-API/internal/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	gin.DefaultWriter = config.NewSlogInfoWriter()
	gin.DefaultErrorWriter = config.NewSlogErrorWriter()
}

func main() {

	logFile := config.SetupSlogOutputFile()
	defer logFile.Close()

	router := gin.New()
	router.Use(middlewares.JSONlogger())
	router.Use(gin.Recovery())

	configurations := cors.DefaultConfig()
	configurations.AllowAllOrigins = true
	configurations.AllowCredentials = true
	configurations.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS"}
	configurations.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Origin", "Cache-Control", "X-Requested-With"}
	configurations.ExposeHeaders = []string{"Content-Length"}
	router.Use(cors.New(configurations))

	client := config.NewMongoClient()
	database := client.ConnectToDB()
	defer client.Disconnect()

	conn := config.NewAMQPconnection()
	defer conn.DisconnectAll()

	firebaseClient := config.NewFirebaseClient()

	app.SetUpApp(router, database, conn, firebaseClient)

	router.Run()

}

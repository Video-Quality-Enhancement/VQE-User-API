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

}

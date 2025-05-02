package routes

import "github.com/gin-gonic/gin"

func RegisterServer(server *gin.Engine) {
	authRouter := server.Group("/auth")
	authRouter.GET("/getUser")
}

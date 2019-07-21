package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/convenient-utils/response"
	"github.com/sillyhatxu/table-backend/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

func InitialAPI() {
	logrus.Info("---------- initial internal api start ----------")
	router := SetupRouter()
	router.Run(config.Conf.Http.Listen)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": response.Success, "message": "OK"})
	})
	userGroup := router.Group("/users")
	{
		userGroup.GET("/login", login)
		userGroup.POST("/login", login)
		userGroup.GET("/main/:id", getById)
	}
	managerGroup := router.Group("/managers")
	{
		managerGroup.GET("/login", managerLogin)
		managerGroup.POST("/login", managerLogin)
		userGroup.GET("", list)
		userGroup.PUT("/enable/:id", enable)
		userGroup.PUT("/disable/:id", disable)
		userGroup.POST("/export", export)
	}
	return router
}

func login(context *gin.Context) {
	context.HTML(http.StatusOK, "users/login.html", gin.H{
		"title": "Users",
	})
}

func managerLogin(context *gin.Context) {
	context.HTML(http.StatusOK, "managers/login.html", gin.H{
		"title": "Managers",
	})
}

func userLogin(context *gin.Context) {

}

func list(context *gin.Context) {

}

func getById(context *gin.Context) {

}

func enable(context *gin.Context) {

}

func disable(context *gin.Context) {

}

func export(context *gin.Context) {

}

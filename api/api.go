package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/convenient-utils/response"
	"github.com/sillyhatxu/table-backend/config"
	"github.com/sillyhatxu/table-backend/dto"
	"github.com/sillyhatxu/table-backend/service"
	"github.com/sirupsen/logrus"
	"net/http"
)

var Context *response.PageContext

func InitialAPI(c *response.PageContext) {
	Context = c
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
	router.GET("/login", loginPage)
	router.POST("/login", login)
	managerGroup := router.Group("/managers")
	{
		managerGroup.GET("/table-one", tableOne)
		managerGroup.POST("/table-one/add", addTableOne)
		managerGroup.POST("/table-one/update", updateTableOne)
		managerGroup.POST("/table-one/export", export)
		managerGroup.PUT("/table-one/get/:id", getById)
		managerGroup.PUT("/table-one/enable/:id", enable)
		managerGroup.PUT("/table-one/disable/:id", disable)
	}
	return router
}

func loginPage(context *gin.Context) {
	context.HTML(http.StatusOK, "managers/login.html", Context.Page(*response.ServerSuccess(dto.LoginDTO{LoginName: "ghy", Password: "123"}, nil)))
}

func tableOne(context *gin.Context) {
	array, err := service.TableOneList()
	if err != nil {
		context.JSON(http.StatusOK, response.ServerError(nil, err.Error(), nil))
		return
	}
	context.HTML(http.StatusOK, "managers/tableOne.html", Context.Page(*response.ServerSuccess(array, nil)))
}

func login(context *gin.Context) {
	var requestBody dto.LoginDTO
	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusOK, response.ServerParamsValidateError(nil, err.Error()))
		return
	}
	if !(requestBody.LoginName == "ghy" && requestBody.Password == "123") {
		context.JSON(http.StatusOK, response.ServerUnauthorizedError("", nil))
		return
	}
	context.JSON(http.StatusOK, response.ServerSuccess(nil, nil))
}

func addTableOne(context *gin.Context) {
	var requestBody dto.AddDTO
	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusOK, response.ServerParamsValidateError(nil, err.Error()))
		return
	}
	logrus.Infof("addTableOne;requestBody : %v", requestBody)
	err = service.TableOneAdd(requestBody)
	if err != nil {
		context.JSON(http.StatusOK, response.ServerError(nil, err.Error(), nil))
		return
	}
	context.JSON(http.StatusOK, response.ServerSuccess(nil, nil))
}

func updateTableOne(context *gin.Context) {
	var requestBody dto.AddDTO
	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusOK, response.ServerParamsValidateError(nil, err.Error()))
		return
	}
	err = service.TableOneAdd(requestBody)
	if err != nil {
		context.JSON(http.StatusOK, response.ServerError(nil, err.Error(), nil))
		return
	}
	context.JSON(http.StatusOK, response.ServerSuccess(nil, nil))
}

func getById(context *gin.Context) {

}

func enable(context *gin.Context) {

}

func disable(context *gin.Context) {

}

func export(context *gin.Context) {

}

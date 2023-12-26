package router1

import (
	"ginchat/docs"
	"ginchat/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	swaggerURL := ginSwagger.URL("http://localhost:8081/swagger/doc.json") // Swagger 文档的 URL
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, swaggerURL))
	

	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/user/createUser", service.CreateUser)

	return r
}

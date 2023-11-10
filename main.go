package main

import (

	//"go-tools/durable"
	"go-tools/router"
	"go-tools/utils"

	"context"
	"log"

	_ "github.com/swaggo/swag"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title go-tools
// @version 1.0
// @description Go-Tools
// @contact.name Mohamed Sameem
// @contact.email mmmohamedsameem@gmail.com

// @BasePath /

func main() {

	//durable.GormDB = durable.InitMysqlDb()
	//defer durable.CloseDbConn(durable.GormDB)

	r := router.SetupRouter(GinContextToContextMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8036")
	log.Println("Starting streamerx-backend service at port : 8036")
	utils.HandleError("Error while starting server ", nil)

}
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

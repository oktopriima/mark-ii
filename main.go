package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-ii/controller/ping"
	"github.com/oktopriima/mark-ii/controller/users"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("ping", ping.PingController)

	{
		userRoute := r.Group("user")
		userRoute.POST("", users.CreateController)
		userRoute.PUT("", )
		userRoute.GET("", )
		userRoute.GET(":id", )
		userRoute.DELETE(":id", )
	}

	{
		orderRoute := r.Group("order")
		orderRoute.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
			return
		})
	}

	r.Run(":9000")
}

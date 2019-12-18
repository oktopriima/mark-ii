package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-ii/controller/ping"
	"github.com/oktopriima/mark-ii/controller/users"
)

func main() {
	r := gin.Default()

	r.POST("ping", ping.PingController)

	{
		userRoute := r.Group("user")
		userRoute.POST("", users.CreateController)
		userRoute.GET(":id", users.FindController)
		userRoute.GET("", users.FindByController)
		//userRoute.PUT("", )
		//userRoute.GET("", )
		//userRoute.GET(":id", )
		//userRoute.DELETE(":id", )
	}

	r.Run(":9000")
}

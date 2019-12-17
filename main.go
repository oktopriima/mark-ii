package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-ii/controller/ping"
	"github.com/oktopriima/mark-ii/controller/users"
)

func main() {
	r := gin.Default()

	r.POST("ping", ping.PingController)
	r.POST("user", users.CreateController)

	r.Run(":9000")
}

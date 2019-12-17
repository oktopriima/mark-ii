package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-ii/conf"
	"net/http"
)

func PingController(ctx *gin.Context) {
	cfg := conf.NewConfig()

	db, err := conf.MysqlConnection(cfg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	_ = db

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})

	return
}

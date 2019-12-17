package users

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-ii/conf"
	"github.com/oktopriima/mark-ii/model"
	"github.com/oktopriima/mark-ii/services"
	"net/http"
)

func CreateController(ctx *gin.Context) {
	var err error
	cfg := conf.NewConfig()

	err, db := conf.MysqlConnection(cfg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userContract := services.NewUserServiceContract(db)

	tx := db.Begin()
	defer tx.Rollback()

	user := new(model.User)
	user.Name = ""
	user.Email = ""
	user.Password = ""

	err = userContract.Create(user, tx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})

	return
}

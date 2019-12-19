package users

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/copier"
	"github.com/oktopriima/mark-ii/conf"
	"github.com/oktopriima/mark-ii/model"
	"github.com/oktopriima/mark-ii/request/users"
	"github.com/oktopriima/mark-ii/services"
	"github.com/oktopriima/mark-v/httpresponse"
	"net/http"
	"runtime"
)

func CreateController(ctx *gin.Context) {
	runtime.GOMAXPROCS(1)

	var err error
	cfg := conf.NewConfig()

	db, err := conf.MysqlConnection(cfg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var req users.CrateRequest

	if err = ctx.ShouldBindWith(&req, binding.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userContract := services.NewUserServiceContract(db)

	tx := db.Begin()
	defer tx.Rollback()

	user := new(model.User)
	err = copier.Copy(&user, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	passTemp, err := conf.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user.Password = string(passTemp)

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

func FindController(ctx *gin.Context) {
	runtime.GOMAXPROCS(2)

	var err error
	cfg := conf.NewConfig()

	db, err := conf.MysqlConnection(cfg)
	if err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	var req users.FindRequest

	if err = ctx.ShouldBindUri(&req); err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	userContract := services.NewUserServiceContract(db)
	data, err := userContract.Find(req.ID)

	httpresponse.NewSuccessResponse(ctx, data)
	return
}

func FindByController(ctx *gin.Context) {
	var err error
	cfg := conf.NewConfig()

	db, err := conf.MysqlConnection(cfg)
	if err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	var req users.FindRequest

	if err = ctx.ShouldBind(&req); err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	userContract := services.NewUserServiceContract(db)

	data, err := userContract.FindBy(nil)

	if err = ctx.ShouldBind(&req); err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, data)
	return
}

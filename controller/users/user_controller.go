package users

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/copier"
	"github.com/oktopriima/mark-ii/conf"
	"github.com/oktopriima/mark-ii/model"
	"github.com/oktopriima/mark-ii/request/users"
	"github.com/oktopriima/mark-ii/services"
	"github.com/oktopriima/mark-ii/validation/user"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"runtime"
)

func CreateController(ctx *gin.Context) {
	runtime.GOMAXPROCS(2)

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
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("createuservalidator", user.CreateUserValidator)
	}

	if err = ctx.ShouldBind(&req); err != nil {
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var req users.FindRequest

	if err = ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userContract := services.NewUserServiceContract(db)
	data, err := userContract.Find(1)

	if err = ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"data": data,
	})
	return
}

func FindByController(ctx *gin.Context) {
	runtime.GOMAXPROCS(2)

	var err error
	cfg := conf.NewConfig()

	db, err := conf.MysqlConnection(cfg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var req users.FindRequest

	if err = ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userContract := services.NewUserServiceContract(db)
	criteria := make(map[string]interface{})
	criteria["email"] = req.Email

	data, err := userContract.FindBy(criteria)

	if err = ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"data": data,
	})
	return
}

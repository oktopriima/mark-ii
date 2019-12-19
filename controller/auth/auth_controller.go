package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-ii/conf"
	"github.com/oktopriima/mark-ii/request/auth"
	"github.com/oktopriima/mark-ii/services"
	"github.com/oktopriima/mark-v/jwtmiddleware"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginController(ctx *gin.Context) {
	var err error
	cfg := conf.NewConfig()

	db, err := conf.MysqlConnection(cfg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var req auth.LoginRequest
	if err = ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userContract := services.NewUserServiceContract(db)
	user, err := userContract.FindOneBy(map[string]interface{}{
		"email": req.Email,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//	 compare password db with request
	byteHash := []byte(user.Password)
	bytePlain := []byte(req.Password)

	if err := bcrypt.CompareHashAndPassword(byteHash, bytePlain); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})
		return
	}

	tokenStruct := new(jwtmiddleware.TokenRequestStructure)
	tokenStruct.Email = user.Email
	tokenStruct.UserID = user.ID

	signInKey := "secret"
	g := jwtmiddleware.NewCustomAuth([]byte(signInKey))
	token, err := g.GenerateToken(*tokenStruct)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": token,
	})
	return

}

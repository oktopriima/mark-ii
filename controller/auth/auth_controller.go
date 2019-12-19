package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-ii/conf"
	"github.com/oktopriima/mark-ii/request/auth"
	"github.com/oktopriima/mark-ii/services"
	"github.com/oktopriima/mark-v/httpresponse"
	"github.com/oktopriima/mark-v/jwtmiddleware"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginController(ctx *gin.Context) {
	var err error
	cfg := conf.NewConfig()

	db, err := conf.MysqlConnection(cfg)
	if err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	var req auth.LoginRequest
	if err = ctx.ShouldBind(&req); err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	userContract := services.NewUserServiceContract(db)
	user, err := userContract.FindOneBy(map[string]interface{}{
		"email": req.Email,
	})
	if err != nil {
		httpresponse.NewErrorException(ctx, http.StatusBadRequest, err)
		return
	}

	//	 compare password from db with request
	byteHash := []byte(user.Password) // password from db
	bytePlain := []byte(req.Password) // password from request

	if err := bcrypt.CompareHashAndPassword(byteHash, bytePlain); err != nil {
		httpresponse.NewErrorException(ctx, http.StatusForbidden, err)
		return
	}

	tokenStruct := new(jwtmiddleware.TokenRequestStructure)
	tokenStruct.Email = user.Email
	tokenStruct.UserID = user.ID

	signInKey := "secret"
	g := jwtmiddleware.NewCustomAuth([]byte(signInKey))
	token, err := g.GenerateToken(*tokenStruct)
	if err != nil {
		httpresponse.NewErrorException(ctx, http.StatusForbidden, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, token)
	return

}

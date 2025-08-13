package middleware

import (
	"chathenon/util"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorization(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	tokenString := strings.Split(auth, " ")

	if len(tokenString) != 2 || tokenString[0] != "Bearer" {
		ctx.Error(errors.New("error unauthorized"))
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, err := util.ParseToken(tokenString[1])
	if err != nil {
		ctx.Error(errors.New("error unauthorized"))
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	id, err := claims.GetSubject()
	if err != nil {
		ctx.Error(errors.New("error claims"))
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	subId, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(errors.New("internal server error"))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Set("sub", subId)

	ctx.Next()
}

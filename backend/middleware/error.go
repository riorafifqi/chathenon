package middleware

import (
	"chathenon/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		util.ResponseMsg(ctx, false, ctx.Errors[0].Error(), nil, http.StatusInternalServerError)
	}
}

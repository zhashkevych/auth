package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/auth/pkg/auth"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase auth.UseCase) {
	h := newHandler(usecase)

	router.POST("/sign-up", h.signUp)
	router.POST("/sign-in", h.signIn)
}

package v1

import (
	"backend/internal/entity"
	"backend/internal/interfaces"
	"context"

	"github.com/gin-gonic/gin"
)

type httpUsecase interface {
	Start(context.Context) (entity.Challenge, error)
	AddPassword(context.Context, entity.Password) error
	CheckHash(context.Context, []entity.Hash) (bool, entity.Challenge, []entity.Password, error)
}

type HTTPController struct {
	router *gin.Engine
	uc     httpUsecase
	log    interfaces.Logger
}

func New(log interfaces.Logger, router *gin.Engine, uc httpUsecase) *HTTPController {
	return &HTTPController{
		router: router,
		uc:     uc,
		log:    log,
	}
}

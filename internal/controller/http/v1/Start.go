package v1

import (
	"backend/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr *HTTPController) Start(gCtx *gin.Context) {
	var (
		err       error
		challenge entity.Challenge
	)

	ctx := gCtx.Request.Context()

	if challenge, err = ctr.uc.Start(ctx); err != nil {
		status, errorMessage := ctr.processError(err)
		gCtx.JSON(status, gin.H{"error": errorMessage})
		return
	}

	response := CheckPasswordRes{
		NeedMore:          true,
		NewChallenge:      challenge.Get(),
		ProbablePasswords: nil,
	}

	gCtx.JSON(http.StatusOK, response)
}

package v1

import (
	"backend/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr *HTTPController) AddPassword(gCtx *gin.Context) {
	var (
		err  error
		json AddPasswordReq
	)

	if err := gCtx.ShouldBindJSON(&json); err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := entity.Password{}

	if err = password.Set(json.Password); err != nil {
		status, errorMessage := ctr.processError(err)
		gCtx.JSON(status, gin.H{"error": errorMessage})
		return
	}

	ctx := gCtx.Request.Context()

	if err = ctr.uc.AddPassword(ctx, password); err != nil {
		status, errorMessage := ctr.processError(err)
		gCtx.JSON(status, gin.H{"error": errorMessage})
		return
	}

	gCtx.JSON(http.StatusOK, gin.H{"message": "Done"})
}

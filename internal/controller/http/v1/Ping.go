package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr *HTTPController) Ping(gCtx *gin.Context) {
	gCtx.JSON(http.StatusOK, gin.H{"message": "Pong"})
}

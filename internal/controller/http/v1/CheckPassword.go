package v1

import (
	"backend/internal/entity"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func (ctr *HTTPController) CheckPassword(gCtx *gin.Context) {
	var (
		err  error
		json CheckPasswordReq
	)

	if err := gCtx.ShouldBindJSON(&json); err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sort.SliceStable(json.Hashes, func(i, j int) bool {
		return json.Hashes[i].Number < json.Hashes[j].Number
	})

	hashes := []entity.Hash{}

	for _, rawHash := range json.Hashes {
		hash := entity.Hash{}
		if err = hash.Set(rawHash.Hash); err != nil {
			status, errorMessage := ctr.processError(err)
			gCtx.JSON(status, gin.H{"error": errorMessage})
			return
		}

		hashes = append(hashes, hash)
	}

	ctx := gCtx.Request.Context()

	needMore, challenge, passwords, err := ctr.uc.CheckHash(ctx, hashes)
	if err != nil {
		status, errorMessage := ctr.processError(err)
		gCtx.JSON(status, gin.H{"error": errorMessage})
		return
	}

	probPasswords := []string{}
	for _, pass := range passwords {
		probPasswords = append(probPasswords, pass.Get())
	}

	response := CheckPasswordRes{
		NeedMore:          needMore,
		NewChallenge:      challenge.Get(),
		ProbablePasswords: probPasswords,
	}

	gCtx.JSON(http.StatusOK, response)
}

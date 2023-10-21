package v1

type AddPasswordReq struct {
	Password string `json:"password" binding:"required"`
}

type CheckPasswordReq struct {
	Hashes []struct {
		Number int    `json:"number" binding:"required"`
		Hash   string `json:"hash" binding:"required"`
	} `json:"hashes" binding:"required"`
}

type CheckPasswordRes struct {
	NeedMore          bool     `json:"need_more"`
	NewChallenge      int      `json:"new_challenge"`
	ProbablePasswords []string `json:"probable_passwords"`
}

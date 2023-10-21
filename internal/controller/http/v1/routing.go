package v1

func (ctr *HTTPController) Routing() {
	ctr.router.GET("/ping", ctr.Ping)

	ctr.router.GET("/start", ctr.Start)
	ctr.router.POST("/add-password", ctr.AddPassword)
	ctr.router.POST("/check-password", ctr.CheckPassword)
}

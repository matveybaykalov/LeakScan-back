package v1

import (
	"backend/internal/config"
	"fmt"
	"net/http"
)

func (ctr *HTTPController) StartServer(cfg config.HTTPConfig) *http.Server {
	ctr.Routing()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: ctr.router,
	}

	ctr.log.Info("starting server")
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			ctr.log.Fatalf("listen: %s\n", err)
		}
	}()

	return srv
}

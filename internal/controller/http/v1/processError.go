package v1

import "net/http"

func (ctr *HTTPController) processError(err error) (int, interface{}) {
	switch {
	default:
		ctr.log.Errorf("!!!!ALARM!!!! unknown error: %s", err)
		return http.StatusInternalServerError, "internal error"
	}
}

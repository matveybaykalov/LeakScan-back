package v1

import (
	"net/http"

	"backend/internal/errors"

	errs "github.com/pkg/errors"
)

func unwrapError(err error) error {
	unwrapped := errs.Unwrap(err)
	if unwrapped == nil {
		return err
	}
	return unwrapError(unwrapped)
}

func (ctr *HTTPController) processError(err error) (int, interface{}) {
	var (
		validationError errors.ValidationError
	)

	switch {
	case errs.As(err, &validationError):
		ctr.log.Warningf("User produce error: %s", err)
		return http.StatusBadRequest, unwrapError(err).Error()

	default:
		ctr.log.Errorf("!!!!ALARM!!!! unknown error: %s", err)
		return http.StatusInternalServerError, "internal error"
	}
}

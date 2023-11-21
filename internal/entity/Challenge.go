package entity

import (
	errs "backend/internal/errors"
	"fmt"
)

type Challenge struct {
	value int
}

func (cg *Challenge) Set(value int) error {
	var (
		minValue int = 1000
		maxValue int = 1000000
	)

	if value < minValue || value > maxValue {
		err := fmt.Errorf("chellenge must be between 1000 and 1000000")
		return errs.ValidationError{Err: err}
	}

	cg.value = value

	return nil
}

func (cg *Challenge) Get() int {
	return cg.value
}

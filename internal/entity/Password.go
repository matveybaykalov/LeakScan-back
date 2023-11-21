package entity

import (
	errs "backend/internal/errors"
	"fmt"
)

type Password struct {
	value string
}

func (pwd *Password) Set(value string) error {
	var (
		minLen int = 1
		maxLen int = 100
	)

	if len(value) < minLen || len(value) > maxLen {
		err := fmt.Errorf("password length must be between 1 and 100")
		return errs.ValidationError{Err: err}
	}

	pwd.value = value

	return nil
}

func (pwd *Password) Get() string {
	return pwd.value
}

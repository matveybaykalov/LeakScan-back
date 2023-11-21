package entity

import (
	errs "backend/internal/erorrs"
	"fmt"
	"regexp"
)

type Hash struct {
	value string
}

func (hs *Hash) Set(value string) error {
	reg := regexp.MustCompile(`^[0-9a-f]{16}$`)
	if !reg.MatchString(value) {
		err := fmt.Errorf("bad hash input")
		return errs.ValidationError{Err: err}
	}

	hs.value = value

	return nil
}

func (hs *Hash) Get() string {
	return hs.value
}

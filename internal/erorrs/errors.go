package erorrs

type ValidationError struct {
	Err error
}

func (e ValidationError) Error() string {
	return e.Err.Error()
}

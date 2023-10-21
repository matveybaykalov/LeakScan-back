package entity

type Password struct {
	value string
}

func (pwd *Password) Set(value string) error {
	// TODO: добавиь валидацию

	pwd.value = value

	return nil
}

func (pwd *Password) Get() string {
	return pwd.value
}

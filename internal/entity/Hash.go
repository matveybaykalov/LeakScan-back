package entity

type Hash struct {
	value string
}

func (hs *Hash) Set(value string) error {
	// TODO: добавить валидацию
	hs.value = value

	return nil
}

func (hs *Hash) Get() string {
	return hs.value
}

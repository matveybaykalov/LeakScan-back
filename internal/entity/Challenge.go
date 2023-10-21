package entity

type Challenge struct {
	value int
}

func (cg *Challenge) Set(value int) error {
	//TODO: можно сделать валидацию
	cg.value = value

	return nil
}

func (cg *Challenge) Get() int {
	return cg.value
}

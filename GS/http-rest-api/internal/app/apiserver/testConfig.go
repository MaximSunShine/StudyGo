package apiserver

type Son struct {
	goToMorning string
}

func NewSon() (*Son, error) {
	return &Son{}, nil
}

func (s *Son) NewSonConfig() error {
	s.goToMorning = "work"
	return nil
}

package apiserver

import (
	"fmt"
)

type Father struct { //объявление
	goToMorning string
	Son         *Son
}

func f() {
	s := NewFather()    //иннициализация
	s.NewFatherConfig() //конфигурация
	fmt.Println(s)

}
func NewFather() *Father {
	return &Father{}
}

func (s *Father) NewFatherConfig() {
	s.goToMorning = "work"
}

func (s *Father) ConfigureSon() error {
	son, err := NewSon()
	if err != nil {
		return err
	}
	son.NewSonConfig()
	if err != nil {
		return err
	}

	s.Son = son
	return nil
}

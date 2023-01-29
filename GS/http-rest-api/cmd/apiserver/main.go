package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/GS/http-rest-api/internal/app/apiserver"
	"log"
)

var configPath string

func init() {
	//передача параметра в программу (флага) путь к файлу содержащего конфиг данные для полей структуры Config
	flag.StringVar(&configPath, "config_path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	cnf := apiserver.NewConfig() //возвращаем структуру Config, предварительно иннициадизируем структуру данными
	// заполняем структуру Config данными из файла адрес которого заполнен значением по умолчанию
	_, err := toml.DecodeFile(configPath, cnf) // читаем из файла значения полей (аналогия с джейсоном)

	if err != nil {
		log.Fatal(err) // не удалось прочитать данные с файла
	}
	// заполняем поле структуры APIServer новыми значением: структурой Config
	if err := apiserver.Start(cnf); err != nil {
		log.Fatal(err)
	}
}

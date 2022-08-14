package database

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
)

var (
	DB     *gorm.DB
	err    error
	Host   string
	Port   string
	DBname string
	User   string
	Pass   string
	Config *configStruct
)

type configStruct struct {
	Host   string `json:"Host"`
	Port   string `json:"Port"`
	DBname string `json:"DBname"`
	User   string `json:"User"`
	Pass   string `json:"Pass"`
}

func ConectaComBancoDeDados() {
	file, err := ioutil.ReadFile("database/config.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal(file, &Config)

	Host = Config.Host
	Port = Config.Port
	DBname = Config.DBname
	User = Config.User
	Pass = Config.Pass

	stringConexao := User + ":" + Pass + "@tcp(" + Host + ":" + Port + ")/" + DBname
	DB, err = gorm.Open(mysql.Open(stringConexao))

	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados.")
	}
}

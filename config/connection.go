package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	models "bookstore/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var Data map[string]interface{}

func ReadDbConfigFile() map[string]string {
	conf, _ := ioutil.ReadFile("config/db.json")

	err := json.Unmarshal(conf, &Data)

	HandleError(err)

	return interfaceToMapStringString(Data["db"])
}

func interfaceToMapStringString(inter interface{}) map[string]string {
	mapStringString := make(map[string]string)
	mapStringInterface := inter.(map[string]interface{})

	for key, value := range mapStringInterface {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		mapStringString[strKey] = strValue
	}

	return mapStringString
}

func dbConfig() map[string]string {
	conf := ReadDbConfigFile()

	return conf
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func InitDb() {
	config := dbConfig()

	driver := config["driver"]
	host := config["host"]
	port := config["port"]
	user := config["user"]
	pass := config["pass"]
	dbname := config["dbname"]

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)

	database, err := gorm.Open(driver, mysqlInfo)

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Book{})

	DB = database
}

package conf

import (
	"log"

	"github.com/asdine/storm"
	"gopkg.in/ini.v1"
	"gopkg.in/macaron.v1"
)

// Cfg represents the pointer to configuration file
var (
	Cfg    *ini.File
	UserDB *storm.DB
)

// find configuration file
func init() {
	var err error
	Cfg, err = macaron.SetConfig("conf/app.ini")
	if err != nil {
		if isDbConnParamsInEnvVariables() {
			log.Printf("[conf/Init] Error during app.ini reading. Error: %s\n", err.Error())
		} else {
			log.Fatalf("[conf/Init] Error during app.ini reading. Error: %s\n", err.Error())
		}
	}
	UserDB, err = storm.Open("./user.db")
	if err != nil {
		log.Fatalln("Error opening UserDB: ", err.Error())
		return
	}
}

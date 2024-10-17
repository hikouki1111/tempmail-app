package utility

import (
	"encoding/json"
	tempmail "github.com/hikouki1111/tempmail-wrapper"
	"log"
	"os"
)

type Userdata struct {
	Accounts []tempmail.Account `json:"accounts"`
}

var filename = "userdata.json"

func NewUserdata() *Userdata {
	if !fileExists(filename) {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		userdata := Userdata{Accounts: make([]tempmail.Account, 0)}
		jsonData, err := json.Marshal(userdata)
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.Write(jsonData)
		if err != nil {
			log.Fatal(err)
		}

		return &userdata
	}

	jsonData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var userdata Userdata
	err = json.Unmarshal(jsonData, &userdata)
	if err != nil {
		log.Fatal(err)
	}

	return &userdata
}

func (ud *Userdata) Store() {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jsonData, err := json.Marshal(*ud)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return err == nil
}

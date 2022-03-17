package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type singleton map[string]map[string]string

var (
	instance singleton
)

func NewClass() singleton {

	if instance == nil {

		instance = make(singleton)
		jsonFile, err := os.Open("account.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			panic(err)
		}
		fmt.Println("Successfully Opened users.json")
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(byteValue, &instance)
		if err != nil {
			panic(err)
		}
	}

	return instance
}

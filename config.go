package main

import (
	"encoding/json"
	"io/ioutil"
)

type tab struct {
	Title  string
	Handle string
}
type config struct {
	Title string
	Tabs  []tab
	Port  int
}

var Config = config{}

func init() {
	txt, err := ioutil.ReadFile("config.json")
	checkErr(err)
	json.Unmarshal(txt, &Config)
}

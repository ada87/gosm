package main

import (
	"encoding/json"
	//	"fmt"
	"io/ioutil"
)

type Tab struct {
	Title  string
	Handle string
}
type config struct {
	Title string
	Tabs  []Tab
}

var Config = config{}

func init() {
	txt, err := ioutil.ReadFile("config.json")
	checkErr(err)
	json.Unmarshal(txt, &Config)
}

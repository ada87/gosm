package main

import (
	"fmt"
	"net/http"
	"os"
)

func (c *Control) HanderOs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("os")
}

func cmd(cmd string) {
	os.Environ()
}

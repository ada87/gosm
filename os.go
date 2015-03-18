package main

import (
	"fmt"
	"net/http"
)

func (c *Control) HanderOs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("os")
}

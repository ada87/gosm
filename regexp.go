package main

import (
	"fmt"
	"net/http"
)

func (c *Control) HanderRegexp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HanderRegexp")
}

package main

import (
	"fmt"
	"net/http"
)

func (c *Controllers) PageRegexp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HanderRegexp")
}

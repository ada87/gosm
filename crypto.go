package main

import (
	"fmt"
	"net/http"
)

func (c *Controllers) PageCrypto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("crypto")
}

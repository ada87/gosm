package main

import (
	"fmt"
	"net/http"
)

func (c *Control) HanderCrypto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("crypto")
}

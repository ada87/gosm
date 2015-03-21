package main

import (
	"net/http"
)

func (c *Controllers) PageRegexp(w http.ResponseWriter, r *http.Request) {
	Portal(w, "regexp.tpl", nil)
}

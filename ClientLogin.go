//Similar to get

package main

import (
	"encoding/base64"
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	//"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	vat := r.Header.Get("authorization")
	if vat[:5] != "Basic" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	data, err := base64.StdEncoding.DecodeString(vat[6:])
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	auth := authenticate(string(data))
	if auth != true {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	APIrequest(w, r)
}

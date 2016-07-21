// APIrequest
package main

import (
	//"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
)

func APIrequest(w http.ResponseWriter, r *http.Request) {

	jsonbody, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var jsonAfterUnmarshal map[string]interface {
	}
	if err := json.Unmarshal(jsonbody, &jsonAfterUnmarshal); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var requestDetails []string
	requestDetails = append(requestDetails, jsonAfterUnmarshal["MobileNo"].(string))
	requestDetails = append(requestDetails, string(jsonbody))
	requestDetails = append(requestDetails, jsonAfterUnmarshal["Service"].(string))
	requestDetails = append(requestDetails, jsonAfterUnmarshal["Date"].(string))
	reqID := InsertAPIrequest(requestDetails, "INSERT INTO APIrequest SET MDN=?,REQUEST=?,SERVICE=?,Date=?")
	fmt.Println(reqID)
	APIresponse(w,r,reqID)

}

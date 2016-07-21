// APIresponse
package main

import (

	"net/http"
)

func APIresponse(w http.ResponseWriter, r *http.Request, reqID int64) {

		var str []string
		str = getAPIrequestdata(reqID)

		// mdn := str[0] + "94"
		//req := str[1] + "ABC"
		// res := str[2] + "DEF"
    resp:=[]byte(`{"Status":"200 OK","Date":"2015-07-20"}`)
		 var newstr []string
		 newstr=append(newstr,string(resp))
		 newstr=append(newstr,str[3])

	InsertAPIresponse(newstr,reqID,"INSERT INTO APIresponse SET RESID=?, RESPONSE = ?,Date = ?")

	APIerrorlog(w,r,reqID,"Immmmerr")

}

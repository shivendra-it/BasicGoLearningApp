// APIerrorLog
package main

import(
	"net/http"
)

func APIerrorlog(w http.ResponseWriter, r *http.Request, reqID int64,err string) {

	var str []string
	str = getAPIrequestdata(reqID)

	var newstr []string
	newstr=append(newstr,err)
	newstr=append(newstr,str[3])
InsertAPIerror(newstr,reqID,"INSERT INTO APIerror SET ERRID=?, ERROR = ?,Date = ?")

}

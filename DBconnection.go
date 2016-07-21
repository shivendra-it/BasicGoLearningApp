package main


import (
  "log"
  "database/sql"
_ "github.com/go-sql-driver/mysql"
  "fmt"
  "strings"
)

var db *sql.DB
var err1 error

const (
    DB_HOST = "tcp(127.0.0.1:3306)"
    DB_NAME = "FinalAPI"
    DB_USER = "root"
    DB_PASS = "12345"
)




func checkErr(err error){
if err != nil {
    panic(err.Error()) // proper error handling instead of panic in your app
}
}



func InsertUserPass (s string, query string){
	st := strings.Split(s, ":")
	st1, st2 := st[0], st[1]
  stmtIns, err := db.Prepare(query)
  checkErr(err)
  res, err := stmtIns.Exec(st1,st2)
            checkErr(err)
  id, err := res.LastInsertId()
            checkErr(err)
            fmt.Println(id)
}


func InsertAPIrequest (str []string, query string) int64{

  stmtIns, err := db.Prepare(query)
  checkErr(err)
  res, err := stmtIns.Exec(str[0],str[1],str[2],str[3])
            checkErr(err)
  id, err := res.LastInsertId()
            checkErr(err)

            return id
}


func InsertAPIresponse (str []string, n int64, query string){

  //n1= int(n)
  stmtIns, err := db.Prepare(query)
  checkErr(err)
  _, err = stmtIns.Exec(n,str[0],str[1])
            checkErr(err)
}



func InsertAPIerror (str []string, n int64, query string){

  stmtIns, err := db.Prepare(query)
  checkErr(err)
  _, err = stmtIns.Exec(n,str[0],str[1])
            checkErr(err)
}



func authenticate(codstr string) bool {
rows, err := db.Query("SELECT username,password FROM ClientLogin")
checkErr(err)

s := string(codstr[:])
st := strings.Split(s, ":")
st1, st2 := st[0], st[1]


for rows.Next() {
    var us string
    var ps string
    err = rows.Scan(&us,&ps)
    checkErr(err)

  if st1==us{
      if st2==ps {
  	fmt.Println("\nPassword Matched\n")
  //	http.Error(w, `Successfully login`, http.StatusOK)
  	return true
      }else{
  	fmt.Println("\nPassword didn't Matched\n")
  //	http.Error(w, `Invalid input parameters!`, http.StatusUnauthorized)
  	return false
  }
  }
}
return false
}


func getAPIrequestdata(reqID int64) []string{

  rows, err := db.Query("SELECT ID,MDN,REQUEST,SERVICE,Date FROM APIrequest")
  checkErr(err)

  for rows.Next() {
      var id int64
      var mdn string
      var req string
      var res string
      var dat string
      err = rows.Scan(&id,&mdn,&req,&res,&dat)
      checkErr(err)

    if reqID == id {
      var str []string
	    str=append(str,mdn)
	    str=append(str,req)
	    str=append(str,res)
	    str=append(str,dat)
      return str
    }
  }
  return nil
}


func EndConnection(){
  defer db.Close()
}



func DBconnection(){
     dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
     db, err1 = sql.Open("mysql", dsn)
     if err1 != nil {
         log.Fatal(err1)
     }
     //defer db.Close()

}

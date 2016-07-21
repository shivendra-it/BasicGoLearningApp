package main
import(
"log"
"net/http"
)
func main(){
DBconnection()
 //Insert("HAsegPY:bsdfgdduddy","INSERT INTO T VALUES(?)")
//Query1("select userpass from T")

	http.HandleFunc("/login", handler)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

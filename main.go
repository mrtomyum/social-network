package main


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

type API struct {
	Message string "json:message"
}

type User struct {
	//ID int "json:id"
	Name string "json:username"
	Email string "json:email"
	First string "json:first"
	Last string "json:last"
}

const (
	DB_HOST = "tcp(nava.work:3306)"
	DB_NAME = "nava"
	DB_USER = /*"root"*/ "root"
	DB_PASS = /*""*/ "mypass"
)

var db *sql.DB

func init()  {
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	db, err := sql.Open("mysql", dsn )
	if err != nil {
		log.Fatal(">>>Open Connection Error: ",err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(">>>DB Ping Error: ",err)
	}
}

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/user/create", CreateUser).Methods("GET")
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	NewUser := User{}
	NewUser.Name = r.FormValue("user")
	NewUser.Email = r.FormValue("email")
	NewUser.First = r.FormValue("first")
	NewUser.Last = r.FormValue("last")
	output, err := json.Marshal(NewUser)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	sql := "INSERT INTO users set user_nickname='" + NewUser.Name + "'," +
		"user_first='" + NewUser.First + "'," +
		"user_last= '" + NewUser.Last + "'," +
		"user_email='" + NewUser.Email + "'"
	fmt.Println(sql)

	q, err := db.Exec(sql)

	if err != nil {
		fmt.Println(">>>>Err after db.Exec(sql):", err)
	}
	fmt.Println(">>>>Success!", q)
}

//func Hello(w http.ResponseWriter, r *http.Request) {
//	urlParams := r.URL.Query()
//	name := urlParams.Get(":name")
//	HelloMessage := "Hello, " + name
//	message := API{HelloMessage}
//	output, err := json.Marshal(message)
//	if err != nil {
//		fmt.Println("Something went wrong!")
//	}
//	fmt.Fprintf(w, string(output))
//}

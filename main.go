package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello "+r.Form.Get("name"))
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		http.Redirect(w, r, "/hello?name=winjo", 302)
	}
}

func main() {
	http.HandleFunc("/hello", sayHelloName)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(nil)
	}
	// db, err := sql.Open("mysql", "root:abc123@/sample?charset=utf8mb4,utf8")
	// checkErr(err)

	// //插入数据
	// stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	// checkErr(err)

	// res, err := stmt.Exec("astaxie", string([]byte{255}), "2012-12-09")
	// checkErr(err)

	// id, err := res.LastInsertId()
	// checkErr(err)

	// fmt.Println(id)

	// rows, err := db.Query("SELECT * FROM userinfo")
	// checkErr(err)

	// for rows.Next() {
	// 	var uid int
	// 	var username string
	// 	var department string
	// 	var created string
	// 	err = rows.Scan(&uid, &username, &department, &created)
	// 	checkErr(err)
	// 	fmt.Println(uid)
	// 	fmt.Println(username)
	// 	fmt.Println(department)
	// 	fmt.Println(created)
	// }
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

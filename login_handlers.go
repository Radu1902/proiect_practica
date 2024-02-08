package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func serveLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func (catalog *Catalog) loginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body any
	decoder.Decode(&body)
	email := body.(string)

	var adminName string
	var studentName string
	catalog.db.QueryRow("select nume from admini where email = $1", email).Scan(&adminName)
	catalog.db.QueryRow("select nume from studenti where email = $1", email).Scan(&studentName)

	if adminName != "" {
		http.HandleFunc("/admin", serveAdmin)
		http.HandleFunc("/admin/studenti", catalog.getStudentList)
	} else if studentName != "" {
		http.HandleFunc("/student", serveStudent)
		http.HandleFunc("/student/info", catalog.getStudentInfo(w, r, email))
	} else {
		fmt.Println("email does not exist in database")
	}
	fmt.Println(email)
}

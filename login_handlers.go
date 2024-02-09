package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	studentPasswordQuery = "select parola from studenti where email = $1"
	adminPasswordQuery   = "select parola from admini where email = $1"
)

func serveLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func (catalog *Catalog) loginHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var body any
	decoder.Decode(&body)
	credentials := body.(string)
	splitCredentials := strings.Split(credentials, ",")
	email := splitCredentials[0]
	parola := splitCredentials[1]

	var parolaAdmin string
	var parolaStudent string
	catalog.db.QueryRow(studentPasswordQuery, email).Scan(&parolaStudent)
	catalog.db.QueryRow(adminPasswordQuery, email).Scan(&parolaAdmin)

	if parolaAdmin != "" {
		if parola == parolaAdmin {
			// serveMux.HandleFunc("/admin", serveAdmin)
			// serveMux.HandleFunc("/admin/studenti", catalog.getStudentList)
			http.HandleFunc("/"+email, serveAdmin)
			http.HandleFunc("/"+email+"/studenti", catalog.getStudentList)
		} else {
			fmt.Println("Parola gresita")
		}
	} else if parolaStudent != "" {
		if parola == parolaStudent {
			// serveMux.HandleFunc("/student", serveStudent)
			// serveMux.HandleFunc("/student/info", catalog.getStudentInfo(email))
			http.HandleFunc("/"+email, serveStudent)
			http.HandleFunc("/"+email+"/info", catalog.getStudentInfo)
			http.HandleFunc("/"+email+"/update", catalog.updateStudent)
			http.HandleFunc("/"+email+"/remove", catalog.removeStudent)
		} else {
			fmt.Println("Parola gresita")
		}
	} else {
		fmt.Println("Email-ul nu exista in baza de date")
	}
	fmt.Println(email)
}

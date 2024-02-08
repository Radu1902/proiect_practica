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
			http.HandleFunc("/admin", serveAdmin)
			http.HandleFunc("/admin/studenti", catalog.getStudentList)
		} else {
			fmt.Println("Parola gresita")
		}
	} else if parolaStudent != "" {
		if parola == parolaStudent {
			http.HandleFunc("/student", serveStudent)
			http.HandleFunc("/student/info", catalog.getStudentInfo(email))
		} else {
			fmt.Println("Parola gresita")
		}
	} else {
		fmt.Println("Email-ul nu exista in baza de date")
	}
	fmt.Println(email)
}

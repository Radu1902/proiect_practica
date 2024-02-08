package main

import "net/http"

func serveStudent(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func (catalog *Catalog) getStudentInfo(w http.ResponseWriter, r *http.Request, email string) {

}

package main

import "net/http"

func serveAdmin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin.html")
}

func (catalog Catalog) getStudentList(w http.ResponseWriter, r *http.Request) {

}

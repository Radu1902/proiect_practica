package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	studentQuery = "SELECT * FROM studenti WHERE email = $1"
)

func serveStudent(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "student.html")
}

func (catalog *Catalog) getStudentInfo(email string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student Student
		err := catalog.db.QueryRow(studentQuery, email).Scan(&student.Email, &student.Nume, &student.Parola, &student.Grupa, &student.Specializare, &student.Firma, &student.Start_date, &student.End_date)
		if err == sql.ErrNoRows {
			fmt.Println(err)
		}

		json.NewEncoder(w).Encode(student)

	}
}

func (catalog *Catalog) updateStudent(email string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//

	}
}

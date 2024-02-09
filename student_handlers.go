package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	studentQuery        = `SELECT * FROM studenti WHERE email = $1`
	updatePracticaQuery = `
	UPDATE studenti
	SET 
	firma = $1,
	start_date = $2,
	end_date = $3
	WHERE email = $4`
	removePracticaQuery = `
	UPDATE studenti
	SET 
	firma = null,
	start_date = null,
	end_date = null
	WHERE email = $1`
)

func serveStudent(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "student.html")
}

func (catalog *Catalog) getStudentInfo(w http.ResponseWriter, r *http.Request) {
	hostName := r.URL
	email := strings.Split(hostName.String(), "/")[1]
	fmt.Println(email)

	var student Student
	err := catalog.db.QueryRow(studentQuery, email).Scan(&student.Email, &student.Nume, &student.Parola, &student.Grupa, &student.Specializare, &student.Firma, &student.Start_date, &student.End_date)
	if err == sql.ErrNoRows {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(student)
}

func (catalog *Catalog) updateStudent(w http.ResponseWriter, r *http.Request) {
	hostName := r.URL
	email := strings.Split(hostName.String(), "/")[1]
	fmt.Println(email)

	decoder := json.NewDecoder(r.Body)
	var body any
	decoder.Decode(&body)
	practicaData := body.(string)
	splitPracticaData := strings.Split(practicaData, ",")
	firma := splitPracticaData[0]
	startDate := splitPracticaData[1]
	endDate := splitPracticaData[2]

	fmt.Println(practicaData)

	catalog.db.Query(updatePracticaQuery, firma, startDate, endDate, email)
}

func (catalog *Catalog) removeStudent(w http.ResponseWriter, r *http.Request) {
	hostName := r.URL
	email := strings.Split(hostName.String(), "/")[1]
	fmt.Println(email, " was removed practica")

	catalog.db.Query(removePracticaQuery, email)
}

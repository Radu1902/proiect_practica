package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	studentListQuery   = `select email, nume, grupa, specializare, firma, start_date, end_date from studenti`
	modifyStudentQuery = `
	UPDATE studenti
	SET 
	email = $1,
	nume = $2,
	grupa = $3,
	specializare = $4,
	firma = $5,
	start_date = $6,
	end_date = $7
	WHERE email = $8`
	removeStudentPracticaQuery = `
	UPDATE studenti
	SET 
	firma = null,
	start_date = null,
	end_date = null
	WHERE email = $1`
	removeStudentQuery = `
	DELETE FROM studenti
	WHERE email = $1`
)

func serveAdmin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin.html")
}

func (catalog Catalog) getStudentList(w http.ResponseWriter, r *http.Request) {
	hostName := r.URL
	adminEmail := strings.Split(hostName.String(), "/")[1]
	fmt.Println(adminEmail)

	var studentList []Student
	rows, err := catalog.db.Query(studentListQuery)
	if err != nil {
		fmt.Println(err)
		// return
	}
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.Email, &student.Nume, &student.Grupa, &student.Specializare, &student.Firma, &student.Start_date, &student.End_date)
		if err != nil {
			fmt.Println(err)
			// return
		}
		studentList = append(studentList, student)
	}
	fmt.Println(studentList)

	json.NewEncoder(w).Encode(studentList)
}

func (catalog *Catalog) removeStudentPractica(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body any
	decoder.Decode(&body)
	email := body.(string)

	catalog.db.Query(removeStudentPracticaQuery, email)
}

func (catalog *Catalog) removeStudent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body any
	decoder.Decode(&body)
	email := body.(string)

	catalog.db.Query(removeStudentQuery, email)
}

func (catalog *Catalog) modifyStudentData(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body any
	decoder.Decode(&body)
	studentData := body.(string)
	newStudentData := strings.Split(studentData, ",")
	oldEmail := newStudentData[0]

	newEmail := newStudentData[1]
	newName := newStudentData[2]
	newGroup := newStudentData[3]
	newSpec := newStudentData[4]
	newFirm := newStudentData[5]
	newStart := newStudentData[6]
	newEnd := newStudentData[7]

	catalog.db.Query(modifyStudentQuery, newEmail, newName, newGroup, newSpec, newFirm, newStart, newEnd, oldEmail)
}

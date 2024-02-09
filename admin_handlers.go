package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	studentListQuery = `select email, nume, grupa, specializare, firma, start_date, end_date from studenti`
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

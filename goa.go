// goa.go
package main

import (
	"encoding/json"
	//	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var g_db *sql.DB = nil

func handleHello(w http.ResponseWriter, r *http.Request) {
	person := struct {
		Name string `json:"name", db:"pk"`
		Age  int    `json:"age"`
	}{"Tom", 12}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

// func handleCourseList(w http.ResponseWriter, r *http.Request) {
// 	courses := []struct {
// 		CourseId      int    `json:"courseId"`
// 		CourseName    string `json:"courseName"`
// 		CourseSummary string `json:"courseSummary"`
// 	}{
// 		{1, "Math", "Math for grade 6"},
// 		{2, "Art", "Art for grade 6"},
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(courses)
// }

type GoaBaseRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (res *GoaBaseRes) setError(code string, message string) {
	res.Code = code
	res.Message = message
}

func (res *GoaBaseRes) setOK() {
	res.Code = "0"
	res.Message = "OK"
}

type Course struct {
	Id      int    `json:"courseId"`
	Name    string `json:"courseName"`
	Summary string `json:"courseSummary"`
}

type CourseListRes struct {
	GoaBaseRes
	Courses []Course `json:"course_list"`
}

func handleCourseList(w http.ResponseWriter, r *http.Request) {
	resCourseList := CourseListRes{}
	resCourseList.setOK()

	rows, err := g_db.Query("SELECT id, course_name, course_summary from course")
	if err == nil {
		columns, err := rows.Columns()
		if err == nil {
			log.Println(columns)
		} else {
			log.Fatal(err)
		}

		resCourseList.Courses = make([]Course, 0, 10)

		for rows.Next() {
			course := Course{}
			err = rows.Scan(&course.Id, &course.Name, &course.Summary)
			if err != nil {
				log.Fatal(err)
				break
			}

			resCourseList.Courses = append(resCourseList.Courses, course)
		}
	} else {
		resCourseList.setError("1", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resCourseList)
}

func main() {
	// open the database

	db, err := sql.Open("mysql", "goa:Goa$1234@/goa_college")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}

	g_db = db

	// setup the api server

	http.HandleFunc("/", handleHello)
	http.HandleFunc("/course/list", handleCourseList)

	port := ":5400"
	log.Println("Goa start at port", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

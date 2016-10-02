// goa.go
package main

import (
	"encoding/json"
	//	"fmt"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	person := struct {
		Name string `json:"name", db:"pk"`
		Age  int    `json:"age"`
	}{"Tom", 12}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func handleCourseList(w http.ResponseWriter, r *http.Request) {
	courses := []struct {
		CourseId      int    `json:"courseId"`
		CourseName    string `json:"courseName"`
		CourseSummary string `json:"courseSummary"`
	}{
		{1, "Math", "Math for grade 6"},
		{2, "Art", "Art for grade 6"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func main() {
	http.HandleFunc("/", handleHello)
	http.HandleFunc("/course/list", handleCourseList)

	port := ":5400"
	log.Println("Goa start at port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

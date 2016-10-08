// goa.go
package goa

import (
	"encoding/json"
	//	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
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

type DatabaseConfig struct {
	Type string `yaml:"type"`
	Url  string `yaml:"url"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

func main() {
	// config yaml file
	filename, err := filepath.Abs("./goa.yaml")
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("filename = ", filename)

	yamlContent, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("yaml content = ", string(yamlContent))

	var config Config
	err = yaml.Unmarshal(yamlContent, &config)
	if err != nil {
		log.Fatal(err)
		return
	}

	if len(config.Database.Type) == 0 {
		log.Fatal("config file error: no database type defined")
		return
	}

	if len(config.Database.Url) == 0 {
		log.Fatal("config file error: no database url defined")
		return
	}

	log.Println("database type = ", config.Database.Type)
	log.Println("database url = ", config.Database.Url)

	// open the database

	//db, err_db := sql.Open("mysql", "goa:Goa$1234@/goa_college")
	db, err_db := sql.Open(config.Database.Type, config.Database.Url)
	if err_db != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	log.Println("ping database ...")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("database ok")

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

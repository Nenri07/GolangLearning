package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(" from here i am gonna do the webapis")
	r := mux.NewRouter()
	courses = append(courses, Course{CourseId: "88", CourseName: "TBW", CoursePrice: 400000, Author: &Author{AuthorName: "Sir Ihtisham", AuthorBday: "18/2/1989"}})

	courses = append(courses, Course{CourseId: "200009888", CourseName: "Concurency with Aws", CoursePrice: 2222, Author: &Author{AuthorName: "Umer", AuthorBday: "11/06/2002"}})

	r.HandleFunc("/", homeRoute).Methods("GET")
	r.HandleFunc("/course", addCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteById).Methods("DELETE")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCoursebyId).Methods("GET")

	log.Fatal(http.ListenAndServe(":3999", r))
}

// /models for controllers just practice
type Course struct {
	CourseName     string  `json:"coursename"`
	CourseId       string  `json:"courseid"`
	CoursePrice    int     `json:"courseprice"`
	CourseDuration float64 `json:"courseduration"`
	Author         *Author `json:"author"`
}

type Author struct {
	AuthorName string `json:"authorname"`
	AuthorBday string `json:"authorbday"`
}

// fake db for courses to add
var courses []Course

// assosiated obect properties
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//controller part

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> This is HOme Welcome </h1>"))

}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fetching courses.....")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

////seeding is when we are testing we inject data by ourself

func getCoursebyId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fetching data by id...........")
	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	for _, values := range courses {
		if values.CourseId == params["id"] {
			json.NewEncoder(w).Encode(values)
			return
		}
	}
	json.NewEncoder(w).Encode("No Result found of course by this id")

}

func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Addig course in db.......")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("please provide data")
		return
	}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("please enter the name")
		return
	}
	//enrate random id and then convert to int fro string

	course.CourseId = strconv.Itoa(rand.IntN(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updaing by id......")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop, id, expel, replace id by param id(rewrite), rewrite new course at that id from body

	for key, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:key], courses[key+1:]...)
			var updatedcourse Course
			json.NewDecoder(r.Body).Decode(&updatedcourse)
			updatedcourse.CourseId = params["id"]
			courses = append(courses, updatedcourse)
			fmt.Println("updatedcourse value", updatedcourse)

			json.NewEncoder(w).Encode(updatedcourse)
			return
		}

	}

}

func deleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting by id......")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for key, course := range courses {
		if course.CourseId == id {
			courses = append(courses[:key], courses[key+1:]...)
			json.NewEncoder(w).Encode("the value of the following id has been deleted")
			break
		}
		json.NewEncoder(w).Encode("no course find by this id")
		return

	}

}

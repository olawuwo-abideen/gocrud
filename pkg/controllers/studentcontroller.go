package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/olawuwo-abideen/studentmanagementsystem/pkg/models"
	"github.com/olawuwo-abideen/studentmanagementsystem/pkg/utils"
)

var NewStudent models.Student

func GetStudent(w http.ResponseWriter, r *http.Request) {
	newStudents := models.GetAllStudents()
	res, _ := json.Marshal(newStudents)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	studentDetails, _ := models.GetStudentById(ID)
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	CreateStudent := &models.Student{}
	utils.ParseBody(r, CreateStudent)
	b := CreateStudent.CreateStudent()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	student := models.DeleteStudent(ID)
	res, _ := json.Marshal(student)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var updateStudent = &models.Student{}
	utils.ParseBody(r, updateStudent)
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	studentDetails, db := models.GetStudentById(ID)
	if updateStudent.FirstName != "" {
		studentDetails.FirstName = updateStudent.FirstName
	}
	if updateStudent.LastName != "" {
		studentDetails.LastName = updateStudent.LastName
	}
	if updateStudent.Email != "" {
		studentDetails.Email = updateStudent.Email
	}
	if updateStudent.Department != "" {
		studentDetails.Department = updateStudent.Department
	}
	db.Save(&studentDetails)
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

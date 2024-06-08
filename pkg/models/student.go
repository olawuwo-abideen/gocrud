package models

import (
	"github.com/jinzhu/gorm"
	"github.com/olawuwo-abideen/studentmanagementsystem/pkg/config"
)

var db *gorm.DB

type Student struct {
	ID         uint `gorm:"primaryKey"`
	FirstName  string
	LastName   string
	Email      string
	Department string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
}

func (b *Student) CreateStudent() *Student {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}

func GetStudentById(Id int64) (*Student, *gorm.DB) {
	var getStudent Student
	db := db.Where("ID=?", Id).Find(&getStudent)
	return &getStudent, db
}

func DeleteStudent(ID int64) Student {
	var student Student
	db.Where("ID=?", ID).Delete(student)
	return student
}

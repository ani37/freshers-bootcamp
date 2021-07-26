package Models

import (
	"fmt"
	"Exercise2/Config"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllStudents Fetch all student data
func GetAllStudents(student *[]Student) (err error) {
	if err = Config.DB.Find(student).Error; err != nil {
		return err
	}
	return nil
}

//AddStudentDetails ... Insert New data
func AddStudentDetails(student *Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

//GetStudentByID ... Fetch only one student by ID
func GetStudentByID(student *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}

//UpdateStudentDetails ... Update student
func UpdateStudentDetails(student *Student, id string) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}

//DeleteStudentDetails ... Delete student
func DeleteStudentDetails(student *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil
}
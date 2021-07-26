package Models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	ID        uint
	FirstName string
	LastName  string
	Dob       string
}

type StudentAddress struct {
	gorm.Model
	StudentID uint		`gorm:"primaryKey"`
	Address   string	`gorm:"primaryKey"`
}

type StudentMarks struct {
	gorm.Model
	StudentID 	uint   `gorm:"primaryKey"`
	Subject   	string `gorm:"primaryKey"`
	Marks     	uint
}

func (b *Student) TableName() string {
	return "student"
}

func (b *StudentMarks) TableName() string {
	return "studentmarks"
}

func (b *StudentAddress) TableName() string {
	return "studentaddress"
}
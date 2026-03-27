package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&Student{})
	return db, nil
}

func AddStudent(student Student) error {
	db, err := Init()
	if err != nil {
		return err
	}
	result := db.Create(&student)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Sucessuful! Student created.")
	return nil
}

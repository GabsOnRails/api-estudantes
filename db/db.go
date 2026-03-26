package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	CPF    int
	Email  string
	Age    int
	Active bool
}

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&Student{})
	return db, nil
}

func AddStudent() {
	db, err := Init()
	if err != nil {
		log.Fatalln(err)
	}
	student := Student{Name: "John Doe", CPF: 123456789, Email: "john.doe@example.com", Age: 20, Active: true}
	result := db.Create(&student)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	fmt.Println("Sucessuful! Student created.")
}

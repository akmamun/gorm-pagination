package main

import (
	"encoding/json"
	"fmt"

	"github.com/akmamun/gorm-pagination/pagination"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Example struct {
	Id   int    `json:"id"`
	Data string `json:"data" binding:"required"`
}

var example Example

func autoMigration() {
	db.AutoMigrate(&example)

}
func insertedData() {
	insertedData := Example{Data: "data"}
	db.Create(&insertedData)
	fmt.Println("Inserted Data!")

}

func main() {
	insertedData := Example{Data: "data"}

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("db connection close")
	} else {
		autoMigration()
	}

	var example Example
	limit := 10
	offset := 0

	query := db.Model(&example).Where("id = ?", 1)
	data, err := pagination.Paginate[Example](query, limit, offset)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)

}

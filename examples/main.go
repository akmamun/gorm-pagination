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

func main() {
	var example Example
	insertedData := Example{Data: "data"}

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err == nil {
		db.AutoMigrate(&example)
		db.Create(&insertedData)
		fmt.Println("Inserted!")
	} else {
		fmt.Println(err)
		return
	}
	ex := pagination.Paginate(&pagination.Param{
		DB:      db,
		Limit:   10,
		Offset:  1,
		OrderBy: "id ASC",
	}, &example)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ex)

}

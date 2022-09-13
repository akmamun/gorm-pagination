package main

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Example struct {
	Id        int        `json:"id"`
	Data      string     `json:"data" binding:"required"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,string,omitempty"`
}

func main() {
	var example Example
	insertedData := Example{Data: "data"}

	db, err := gorm.Open("sqlite3", "example.db")

	if err != nil {
		db.AutoMigrate(&example)
		db.Create(insertedData)
		fmt.Println("Insert OK!")
	} else {
		fmt.Println(err)
		return
	}
	pagination.Paginate(&pagination.Param{
		DB:      db,
		Limit:   10,
		Offset:  1,
		OrderBy: "id ASC",
	}, &example)

}

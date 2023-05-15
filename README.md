# gorm-pagination
 Gorm Pagination library

## Installation 
```shell
go get github.com/akmamun/gorm-pagination
```

## Usage Example 
- Example 1
```go
    type Example struct {
    Id        int        `json:"id"`
    Data      string     `json:"data" binding:"required"`
    CreatedAt *time.Time `json:"created_at,string,omitempty"`
    UpdatedAt *time.Time `json:"updated_at,string,omitempty"`
    }

    var example []Example
     
    pagination.Paginate[Example](*gorm.DB, limit, Offset)
```
- Example 2
```go
    type Example struct {
    Id        int        `json:"id"`
    Data      string     `json:"data" binding:"required"`
    CreatedAt *time.Time `json:"created_at,string,omitempty"`
    UpdatedAt *time.Time `json:"updated_at,string,omitempty"`
    }

    data, err := pagination.Paginate[Example](*gorm.DB, limit, Offset)

```
### Pagination View
- Input Params `limit` and `offset`
- Example Url `localhost:8000/test?limit=1&offset=1`
- Output Format
```json
{
  "total_record": 17,
  "total_page": 17,
  "offset": 1,
  "limit": 1,
  "prev_page": 1,
  "next_page": 2,
  "results": [
    {
      "id": 2,
      "data": "this is test data",
      "created_at": "2022-09-13T17:42:36.358116Z",
      "updated_at": "2022-09-13T17:42:36.358116Z"
    }
  ]
}
```
- Full Example
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/akmamun/gorm-pagination"  

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
  db.Model()
  data, err := pagination.Paginate[Example](*gorm.DB, limit, Offset)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)
}
```
### Credit https://github.com/hellokaton/gorm-paginator

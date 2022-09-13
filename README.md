gorm-pagination

## Usage
```shell
go get github.com/akmamun/gormpagination
```

```go
type Example struct {
Id        int        `json:"id"`
Data      string     `json:"data" binding:"required"`
CreatedAt *time.Time `json:"created_at,string,omitempty"`
UpdatedAt *time.Time `json:"updated_at,string,omitempty"`
}

var example []Example
db = db.Where("id > ?", 0)
pagination.Paginate(&pagination.Param{
DB:      db,
Limit:   limit,
Offset:  offset,
OrderBy: "id ASC",
}, &example)
```
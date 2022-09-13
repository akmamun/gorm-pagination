gorm-pagination

## Usage
```shell
go get github.com/akmamun/gorm-pagination
```

```go
type User struct {
	ID       int
	UserName string `gorm:"not null;size:100;unique"`
}

var example []User
db = db.Where("id > ?", 0)
pagination.Paginate(&pagination.Param{
DB:      db,
Limit:   limit,
Offset:  offset,
OrderBy: "id ASC",
}, &example)
```
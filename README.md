# gorm-pagination

## Installation 
```shell
go get github.com/akmamun/gorm-pagination/pagination
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
     
    pagination.Paginate(&pagination.Param{
    DB:      *gorm.DB,
    Limit:   limit,
    Offset:  offset,
    OrderBy: "id ASC",
    }, &example)
```
- Example 2
```go
    type Example struct {
    Id        int        `json:"id"`
    Data      string     `json:"data" binding:"required"`
    CreatedAt *time.Time `json:"created_at,string,omitempty"`
    UpdatedAt *time.Time `json:"updated_at,string,omitempty"`
    }

    var example []Example

    data := pagination.Paginate(&pagination.Param{
            DB:      *gorm.DB,
            Limit:   limit,
            Offset:  offset,
            OrderBy: "id ASC",
            Query : Example{Id:1}
    }, &example)
```
### Credit https://github.com/hellokaton/gorm-paginator

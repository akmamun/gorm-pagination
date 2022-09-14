# gorm-pagination
 Gorm Pagination library

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
### Credit https://github.com/hellokaton/gorm-paginator

package pagination

import (
	"gorm.io/gorm"
	"math"
)

type Param struct {
	DB      *gorm.DB
	Offset  int64
	Limit   int64
	OrderBy string
	Query   interface{}
}

type Result struct {
	TotalRecord int64       `json:"total_record"`
	TotalPage   int64       `json:"total_page"`
	Offset      int64       `json:"offset"`
	Limit       int64       `json:"limit"`
	PrevPage    int64       `json:"prev_page"`
	NextPage    int64       `json:"next_page"`
	Results     interface{} `json:"results"`
}

func Paginate(param *Param, resultData interface{}) *Result {
	db := param.DB

	var result Result
	var count, offset int64
	done := make(chan bool, 1)

	go countResults(db, resultData, done, &count)

	if param.Limit == 0 {
		param.Limit = 10
	}

	if param.Offset == 0 {
		offset = 0
	} else {
		offset = param.Offset
	}

	if err := param.Query; err == nil {param.Query = "" }
	
	db.Offset(int(offset)).
		Limit(int(param.Limit)).
		Order(param.OrderBy).
		Where(param.Query).
		Find(resultData)

	<-done

	result.TotalRecord = count
	result.Results = resultData

	result.Offset = offset
	result.Limit = param.Limit
	result.TotalPage = int64(math.Ceil(float64(count) / float64(param.Limit)))

	if param.Offset > 1 {
		result.PrevPage = param.Offset - 1
	} else {
		result.PrevPage = param.Offset
	}

	if param.Offset == result.TotalPage {
		result.NextPage = param.Offset
	} else {
		result.NextPage = param.Offset + 1
	}
	return &result
}

// count through separate channel
func countResults(db *gorm.DB, anyType interface{}, done chan bool, count *int64) {
	db.Model(anyType).Count(count)
	done <- true
}

package pagination

import (
	"gorm.io/gorm"
	"math"
)

type Result[T any] struct {
	TotalRecord int64 `json:"total_record"`
	TotalPage   int64 `json:"total_page"`
	Offset      int64 `json:"offset"`
	Limit       int64 `json:"limit"`
	PrevPage    int64 `json:"prev_page"`
	NextPage    int64 `json:"next_page"`
	Results     []T   `json:"results"`
}

func Paginate[T any](db *gorm.DB, limit, offset int64) (*Result[T], error) {
	var result Result[T]
	var data []T
	var count int64

	if limit == 0 {
		limit = 10
	}

	db.Count(&count)
	err := db.
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&data).Error

	if err != nil {
		return nil, err
	}

	result.TotalRecord = count
	result.Results = data

	result.Offset = offset
	result.Limit = limit
	result.TotalPage = int64(math.Ceil(float64(count) / float64(limit)))

	if offset > 1 {
		result.PrevPage = offset - 1
	} else {
		result.PrevPage = offset
	}

	if offset == result.TotalPage {
		result.NextPage = offset
	} else {
		result.NextPage = offset + 1
	}
	return &result, nil
}

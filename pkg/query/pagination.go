package query

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Page         uint   `json:"page,omitempty"`
	Limit        uint   `json:"limit,omitempty"`
	TotalPage    uint   `json:"total_page,omitempty"`
	TotalData    uint   `json:"total_data,omitempty"`
	NextPage     string `json:"next_page,omitempty"`
	PreviousPage string `json:"previous_page,omitempty"`
}

func (p *Pagination) GetPage() uint {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetOffset() uint {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() uint {
	if p.Limit == 0 {
		p.Limit = 5
	}
	return p.Limit
}

func Paginate(data interface{}, pagination Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var TotalData int64

	db.Model(data).Count(&TotalData)
	pagination.TotalData = uint(TotalData)
	pagination.TotalPage = uint(math.Ceil(float64(TotalData) / float64(pagination.Limit)))

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(pagination.GetOffset())).Limit(int(pagination.GetLimit()))
	}
}

package response

import (
	"gorm.io/gorm"
)

const (
	EmptyPage    = 0
	EmptyPerPage = 0
)

type Pagination struct {
	List    interface{}            `json:"list"`
	Page    int                    `json:"page"`
	PerPage int                    `json:"perPage"`
	Total   int64                  `json:"total"`
	Extra   map[string]interface{} `json:"extra"`
}

func NewPagination(page, perPage int, list interface{}, tx *gorm.DB) (Pagination, error) {
	p := Pagination{
		Page:    page,
		PerPage: perPage,
		List:    list,
		Extra:   make(map[string]interface{}),
	}

	if EmptyPage == p.Page {
		p.Page = 1
	}
	if EmptyPerPage == p.PerPage {
		p.PerPage = 15
	}

	var err error

	err = tx.Count(&p.Total).Error
	if err != nil {
		return p, err
	}

	err = tx.Limit(perPage).Offset((page - 1) * perPage).Find(p.List).Error
	if err != nil {
		return p, err
	}

	return p, err
}

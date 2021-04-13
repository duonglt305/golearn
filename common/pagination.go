package common

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type PaginationResponse struct {
	From        int   `json:"from"`
	To          int   `json:"to"`
	CurrentPage int   `json:"current_page"`
	LastPage    int   `json:"last_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
}
type Pagination struct {
	CurrentPage int
	PerPage     int
	Total       int64
}

func (p *Pagination) Next() int {
	next := p.CurrentPage + 1
	lastPage := p.LastPage()
	if next > lastPage {
		next = lastPage
	}
	return next
}

func (p *Pagination) Previous() int {
	prev := p.CurrentPage - 1
	if prev < 1 {
		prev = 1
	}
	return prev
}

func (p *Pagination) LastPage() int {
	totalPage := p.Total / int64(p.PerPage)
	remainder := p.Total % int64(p.PerPage)
	if remainder != 0 {
		totalPage += 1
	}
	return int(totalPage)
}

func (p *Pagination) From() int {
	return p.To() - p.PerPage
}

func (p *Pagination) To() int {
	return p.CurrentPage * p.PerPage
}

func NewPagination() *Pagination {
	return &Pagination{}
}

func Paginate(c *gin.Context, p *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}
		limit, _ := strconv.Atoi(c.Query("limit"))
		if limit > 100 {
			limit = 100
		} else if limit <= 0 {
			limit = 15
		}
		offset := (page - 1) * limit
		p.PerPage = limit
		p.CurrentPage = page
		db.Session(&gorm.Session{}).Select("id").Count(&p.Total)
		return db.Offset(offset).Limit(limit)
	}
}

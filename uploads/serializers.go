package uploads

import (
	"github.com/gin-gonic/gin"
	"time"
)

type ListMediaSerializer struct {
	Context *gin.Context `json:"-"`
	Media   []Media
}

type ListMediaResponse struct {
	Data        []MediaResponse `json:"data"`
	From        int             `json:"from"`
	To          int             `json:"to"`
	CurrentPage int             `json:"current_page"`
	LastPage    int             `json:"last_page"`
	PerPage     int             `json:"per_page"`
	Total       int             `json:"total"`
}

type MediaResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	ParentID  uint      `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s ListMediaSerializer) Response() ListMediaResponse {
	resp := ListMediaResponse{
		Data:        nil,
		CurrentPage: 0,
		LastPage:    0,
		Total:       0,
	}
	for _, media := range s.Media {
		resp.Data = append(
			resp.Data,
			MediaResponse{
				ID:        media.ID,
				Name:      media.Name,
				Slug:      media.Slug,
				ParentID:  media.ParentID,
				CreatedAt: media.CreatedAt,
				UpdatedAt: media.CreatedAt,
			},
		)
	}
	return resp
}

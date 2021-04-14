package uploads

import (
	"github.com/gin-gonic/gin"
	"golearn/common"
	"time"
)

type ListMediaSerializer struct {
	Context    *gin.Context `json:"-"`
	Media      []MediaItem
	Pagination *common.Pagination
}

type ListMediaResponse struct {
	common.PaginationResponse
	Data []interface{} `json:"data"`
}

type MediaResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	ParentID  uint      `json:"parent_id"`
	OwnerID   uint      `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FileResponse struct {
	MediaResponse
	Mimes string `json:"mimes"`
	Size  uint   `json:"size"`
	Path  string `json:"path"`
}

func (s *ListMediaSerializer) Response() ListMediaResponse {
	resp := ListMediaResponse{
		Data: nil,
		PaginationResponse: common.PaginationResponse{
			From:        s.Pagination.From(),
			To:          s.Pagination.To(),
			CurrentPage: s.Pagination.CurrentPage,
			LastPage:    s.Pagination.LastPage(),
			PerPage:     s.Pagination.PerPage,
			Total:       s.Pagination.Total,
		},
	}
	for _, media := range s.Media {
		var m interface{}
		m = MediaResponse{
			ID:        media.ID,
			Name:      media.Name,
			ParentID:  media.ParentID,
			CreatedAt: media.CreatedAt,
			UpdatedAt: media.CreatedAt,
		}
		if media.Type == File {
			m = FileResponse{
				MediaResponse: m.(MediaResponse),
				Mimes:         "",
				Size:          0,
				Path:          "",
			}
		}
		resp.Data = append(resp.Data, m)
	}
	return resp
}

type MediaItemSerializer struct {
	Context   *gin.Context `json:"-"`
	MediaItem MediaItem
}

func (s *MediaItemSerializer) Response() *MediaResponse {
	return &MediaResponse{
		ID:        s.MediaItem.ID,
		Name:      s.MediaItem.Name,
		ParentID:  s.MediaItem.ParentID,
		OwnerID:   s.MediaItem.OwnerID,
		CreatedAt: s.MediaItem.CreatedAt,
		UpdatedAt: s.MediaItem.UpdatedAt,
	}
}

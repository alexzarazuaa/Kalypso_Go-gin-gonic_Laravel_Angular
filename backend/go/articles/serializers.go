package articles

import (
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/users"
	"gopkg.in/gin-gonic/gin.v1"
)
type ArticleUserSerializer struct {
	C *gin.Context
	ArticleUserModel
}

func (s *ArticleUserSerializer) Response() users.ProfileResponse {
	response := users.ProfileSerializer{s.C, s.ArticleUserModel.UserModel}
	return response.Response()
}

type ArticleSerializer struct {
	C *gin.Context
	ArticleModel
}

type ArticleResponse struct {
	ID             uint                  `json:"-"`
	Slug           string                `json:"slug"`
	Name           string  				 `json:"name"`
	Brand          string  				 `json:"brand"`
	Img            string  				 `json:"img"`
	Description    string                `json:"description"`
	Rating         int                `json:"rating"`
	Category       string                `json:"category"`
	CreatedAt      string                `json:"createdAt"`
	UpdatedAt      string                `json:"updatedAt"`
	Author         users.ProfileResponse `json:"author"`
}

type ArticlesSerializer struct {
	C        *gin.Context
	Articles []ArticleModel
}

func (s *ArticleSerializer) Response() ArticleResponse {
	// myUserModel := s.C.MustGet("my_user_model").(users.UserModel)
	authorSerializer := ArticleUserSerializer{s.C, s.Author}
	response := ArticleResponse{
		ID:          s.ID,
		Slug:        s.Slug,
		Name:        s.Name,
		Brand:       s.Brand,
		Img:         s.Img,
		Description: s.Description,
		Rating:      s.Rating,
		Category:    s.Category,
		CreatedAt:   s.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		UpdatedAt:   s.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		Author:      authorSerializer.Response(),
	}
	return response
}

func (s *ArticlesSerializer) Response() []ArticleResponse {
	response := []ArticleResponse{}
	for _, article := range s.Articles {
		serializer := ArticleSerializer{s.C, article}
		response = append(response, serializer.Response())
	}
	return response
}
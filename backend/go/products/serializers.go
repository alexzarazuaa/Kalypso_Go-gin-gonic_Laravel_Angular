package products

import (
	"goKa/users"
	"github.com/gin-gonic/gin"
)
type ProductUserSerializer struct {
	C *gin.Context
	ProductUsers
}

func (s *ProductUserSerializer) Response() users.ProfileResponse {
	response := users.ProfileSerializer{s.C, s.ProductUsers.Users}
	return response.Response()
}

type ProductSerializer struct {
	C *gin.Context
	ProductModel
}

type ProductResponse struct {
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

type ProductsSerializer struct {
	C        *gin.Context
	Products []ProductModel
}

func (s *ProductSerializer) Response() ProductResponse {
	// myUsers := s.C.MustGet("my_user_model").(users.Users)
	authorSerializer := ProductUserSerializer{s.C, s.Author}
	response := ProductResponse{
		ID:          s.ID,
		Slug:        s.Slug,
		Name:        s.Name,
		Brand:       s.Brand,
		Img:         s.Img,
		Description: s.Description,
		Rating:      s.Rating,
		Category:    s.Category,
		Author:      authorSerializer.Response(),
	}
	return response
}

func (s *ProductsSerializer) Response() []ProductResponse {
	response := []ProductResponse{}
	for _, product := range s.Products {
		serializer := ProductSerializer{s.C, product}
		response = append(response, serializer.Response())
	}
	return response
}
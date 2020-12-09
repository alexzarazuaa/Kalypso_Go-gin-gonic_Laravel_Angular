package buy_products

import (
	"goKa/users"
	"github.com/gin-gonic/gin"
)
type Buy_ProductUserSerializer struct {
	C *gin.Context
	Buy_ProductUserModel
}

func (s *Buy_ProductUserSerializer) Response() users.ProfileResponse {
	response := users.ProfileSerializer{s.C, s.Buy_ProductUserModel.UserModel}
	return response.Response()
}

type Buy_ProductSerializer struct {
	C *gin.Context
	Buy_ProductModel
}

type Buy_ProductResponse struct {
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

type Buy_ProductsSerializer struct {
	C        *gin.Context
	Buy_Products []Buy_ProductModel
}

func (s *Buy_ProductSerializer) Response() Buy_ProductResponse {
	// myUserModel := s.C.MustGet("my_user_model").(users.UserModel)
	authorSerializer := Buy_ProductUserSerializer{s.C, s.Author}
	response := Buy_ProductResponse{
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

func (s *Buy_ProductsSerializer) Response() []Buy_ProductResponse {
	response := []Buy_ProductResponse{}
	for _, buy_product := range s.Buy_Products {
		serializer := Buy_ProductSerializer{s.C, buy_product}
		response = append(response, serializer.Response())
	}
	return response
}
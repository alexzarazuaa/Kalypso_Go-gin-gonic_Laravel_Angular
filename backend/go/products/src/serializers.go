package products

import (
	"github.com/gin-gonic/gin"
)
type ProductUserSerializer struct {
	C *gin.Context
	ProductUsers
}

func (s *ProductUserSerializer) Response() ProfileResponse {
	response := ProfileSerializer{s.C, s.ProductUsers.Users}
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
	Author         ProfileResponse `json:"author"`
	Favorite       bool                  `json:"favorited"`
	FavoritesCount uint                  `json:"favoritesCount"`
}

type ProductsSerializer struct {
	C        *gin.Context
	Products []ProductModel
}

func (s *ProductSerializer) Response() ProductResponse {
	myUserModel := s.C.MustGet("my_user_model").(Users)
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
		Favorite:       s.isFavoriteBy(GetProductUsers(myUserModel)),
		FavoritesCount: s.favoritesCount(),
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



//-----------------PROFILE-----------------------------//

type ProfileSerializer struct {
	C *gin.Context
	Users
}

// Declare your response schema here
type ProfileResponse struct {
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Image     *string `json:"image"`
	Karma 	   int    `json:"karma"`
}

// Put your response logic including wrap the userModel here.
func (self *ProfileSerializer) Response() ProfileResponse {
	profile := ProfileResponse{
		Username:  self.Username,
		Image:     self.Image,
		Karma:	   self.Karma,
		Email:	   self.Email,
	}
	return profile
}

//----------------END PROFILE-------------------------//
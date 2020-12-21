package buy_products

import (
	"github.com/gin-gonic/gin"
)
type Buy_ProductUserSerializer struct {
	C *gin.Context
	Buy_ProductUsers
}

func (s *Buy_ProductUserSerializer) Response() ProfileResponse {
	response := ProfileSerializer{s.C, s.Buy_ProductUsers.Users}
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
	Author         ProfileResponse `json:"author"`
}

type Buy_ProductsSerializer struct {
	C        *gin.Context
	Buy_Products []Buy_ProductModel
}

func (s *Buy_ProductSerializer) Response() Buy_ProductResponse {
	// myUsers := s.C.MustGet("my_user_model").(users.Users)
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
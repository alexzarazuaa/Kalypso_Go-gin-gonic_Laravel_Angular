package buy_products

import (
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/common"
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/users"
	"github.com/gosimple/slug"
	"gopkg.in/gin-gonic/gin.v1"
)

type Buy_ProductModelValidator struct {
	Buy_Product struct {
		Name       string   `form:"name" json:"name" binding:"exists,min=4"`
		Description string   `form:"description" json:"description" binding:"max=2048"`
		Brand        string   `form:"brand" json:"brand" binding:"max=2048"`
		Category        string   `form:"category" json:"category" binding:"max=2048"`

	} `json:"buy_product"`
	buy_productModel Buy_ProductModel `json:"-"`
}

func NewBuy_ProductModelValidator() Buy_ProductModelValidator {
	return Buy_ProductModelValidator{}
}

func (s *Buy_ProductModelValidator) Bind(c *gin.Context) error {
	myUserModel := c.MustGet("my_user_model").(users.UserModel)

	err := common.Bind(c, s)
	if err != nil {
		return err
	}
		
	s.buy_productModel.Slug = slug.Make(s.Buy_Product.Name)
	s.buy_productModel.Name=(s.Buy_Product.Name)
	s.buy_productModel.Brand=(s.Buy_Product.Brand)
	s.buy_productModel.Description = s.Buy_Product.Description
	s.buy_productModel.Rating=0
	s.buy_productModel.Description = s.Buy_Product.Category
	s.buy_productModel.Author = GetBuy_ProductUserModel(myUserModel)

	return nil
}

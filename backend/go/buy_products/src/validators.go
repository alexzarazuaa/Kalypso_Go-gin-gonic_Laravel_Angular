package buy_products

import (
	"gobuys_products/common"
	"github.com/gosimple/slug"
	"github.com/gin-gonic/gin"
)

type Buy_ProductModelValidator struct {
	Buy_Product struct {
		Name       string   `form:"name" json:"name" binding:"required,min=4"`
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
	myUsers := c.MustGet("my_user_model").(users.Users)

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
	s.buy_productModel.Author = GetBuy_ProductUsers(myUsers)

	return nil
}

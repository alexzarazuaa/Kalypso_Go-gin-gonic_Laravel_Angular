package products

import (
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/common"
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/users"
	"github.com/gosimple/slug"
	"gopkg.in/gin-gonic/gin.v1"
)

type ProductModelValidator struct {
	Product struct {
		Name       string   `form:"name" json:"name" binding:"exists,min=4"`
		Description string   `form:"description" json:"description" binding:"max=2048"`
		Brand        string   `form:"brand" json:"brand" binding:"max=2048"`
		Category        string   `form:"category" json:"category" binding:"max=2048"`

	} `json:"product"`
	productModel ProductModel `json:"-"`
}

func NewProductModelValidator() ProductModelValidator {
	return ProductModelValidator{}
}

func NewProductModelValidatorFillWith(productModel ProductModel) ProductModelValidator {
	productModelValidator := NewProductModelValidator()
	return productModelValidator
}

func (s *ProductModelValidator) Bind(c *gin.Context) error {
	myUserModel := c.MustGet("my_user_model").(users.UserModel)

	err := common.Bind(c, s)
	if err != nil {
		return err
	}
		
	s.productModel.Slug = slug.Make(s.Product.Name)
	s.productModel.Name=(s.Product.Name)
	s.productModel.Brand=(s.Product.Brand)
	s.productModel.Description = s.Product.Description
	s.productModel.Rating=0
	s.productModel.Description = s.Product.Category
	s.productModel.Author = GetProductUserModel(myUserModel)

	return nil
}

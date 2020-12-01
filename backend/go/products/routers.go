package products

import (
	"errors"
	"net/http"
	"fmt"


	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/common"
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/users"
	"gopkg.in/gin-gonic/gin.v1"
)

func ProductsRegister(router *gin.RouterGroup) {
	router.POST("/", ProductCreate)
	router.PUT("/:slug", ProductUpdate)
	router.DELETE("/:slug", ProductDelete)
}

func ProductsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", ProductList)
	router.GET("/:slug", ProductRetrieve)

}


func ProductCreate(c *gin.Context) {
	productModelValidator := NewProductModelValidator()
	if err := productModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	//fmt.Println(productModelValidator.productModel.Author.UserModel)

	if err := SaveOne(&productModelValidator.productModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := ProductSerializer{c, productModelValidator.productModel}
	c.JSON(http.StatusCreated, gin.H{"product": serializer.Response()})
}

func ProductList(c *gin.Context) {
	productModels, modelCount, err := FindManyProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid param")))
		return
	}
	serializer := ProductsSerializer{c, productModels}
	c.JSON(http.StatusOK, gin.H{"products": serializer.Response(), "productsCount": modelCount})
}
func ProductFeed(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	myUserModel := c.MustGet("my_user_model").(users.UserModel)
	if myUserModel.ID == 0 {
		c.AbortWithError(http.StatusUnauthorized, errors.New("{error : \"Require auth!\"}"))
		return
	}
	productUserModel := GetProductUserModel(myUserModel)
	productModels, modelCount, err := productUserModel.GetProductFeed(limit, offset)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid param")))
		return
	}
	serializer := ProductsSerializer{c, productModels}
	c.JSON(http.StatusOK, gin.H{"products": serializer.Response(), "productsCount": modelCount})
}

func ProductRetrieve(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "feed" {
		ProductFeed(c)
		return
	}
	productModel, err := FindOneProduct(&ProductModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid slug")))
		return
	}
	serializer := ProductSerializer{c, productModel}
	c.JSON(http.StatusOK, gin.H{"product": serializer.Response()})
}

func ProductUpdate(c *gin.Context) {
	slug := c.Param("slug")
	productModel, err := FindOneProduct(&ProductModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid slug")))
		return
	}
	productModelValidator := NewProductModelValidatorFillWith(productModel)
	if err := productModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	productModelValidator.productModel.ID = productModel.ID
	if err := productModel.Update(productModelValidator.productModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := ProductSerializer{c, productModel}
	c.JSON(http.StatusOK, gin.H{"product": serializer.Response()})
}

func ProductDelete(c *gin.Context) {
	slug := c.Param("slug")
	fmt.Println(slug);
	fmt.Println("------------------------------")
	err := DeleteProductModel(&ProductModel{Slug: slug})
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid slug")))
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": "Delete success"})
}
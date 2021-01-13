package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"goProducts/common"
	"goProducts/src"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&products.ProductModel{})
	db.AutoMigrate(&products.ProductUsers{})
	db.AutoMigrate(&products.FavoriteModel{})
}

func main() {


	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	MakeRoutes(r)

	v1 := r.Group("/api")


	v1.Use(products.AuthMiddleware(false))
	products.ProductsAnonymousRegister(v1.Group("/products"))


	v1.Use(products.AuthMiddleware(true))
	products.ProductsRegister(v1.Group("/products"))

	r.Run(":3000")
}

func MakeRoutes(r *gin.Engine) {
	cors := func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()

	}
	r.Use(cors)
}
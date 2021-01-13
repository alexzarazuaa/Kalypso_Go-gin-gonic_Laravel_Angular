package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"gobuys_products/common"
	"gobuys_products/src"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&buy_products.Buy_ProductModel{})
	db.AutoMigrate(&buy_products.Buy_ProductUsers{})
}

func main() {


	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	MakeRoutes(r)

	v1 := r.Group("/api")

	v1.Use(buy_products.AuthMiddleware(true))
	buy_products.Buy_ProductsRegister(v1.Group("/buy_products"))

	r.Run(":3001")
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
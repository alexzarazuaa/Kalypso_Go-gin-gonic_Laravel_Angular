package main

import (
	"fmt"

	"github.com/gin-gonic/gin"


	"goKa/routers"
	"goKa/common"
	"goKa/users"
	// "github.com/jinzhu/gorm"
)


func main() {


	db := common.Init()

	defer db.Close()

	r := gin.Default()

	MakeRoutes(r)

	v1 := r.Group("/api")


	routers.ProductList(v1.Group("/products"))
	routers.ProductsAnonymousRegister(v1.Group("/products"))
	
	fmt.Printf("0.0.0.0:3002")
	r.Run(":3002")
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
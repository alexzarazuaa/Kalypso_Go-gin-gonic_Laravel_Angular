package main

import (
	"fmt"

	"gopkg.in/gin-gonic/gin.v1"

<<<<<<< HEAD
	"github.com/jinzhu/gorm"
	"github.com/yomogan/6_gin_gonic_thinkster/articles"
	"github.com/yomogan/6_gin_gonic_thinkster/common"
	"github.com/yomogan/6_gin_gonic_thinkster/users"
=======
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/buy_products"
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/products"
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/common"
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/users"
	"github.com/jinzhu/gorm"
>>>>>>> 54f4ab9460419a42b520998c60f9fa0be7b23b8d
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
<<<<<<< HEAD
	db.AutoMigrate(&articles.ArticleModel{})
	db.AutoMigrate(&articles.TagModel{})
	db.AutoMigrate(&articles.FavoriteModel{})
	db.AutoMigrate(&articles.ArticleUserModel{})
	db.AutoMigrate(&articles.CommentModel{})
=======
	db.AutoMigrate(&buy_products.Buy_ProductModel{})
	db.AutoMigrate(&products.ProductModel{})
	db.AutoMigrate(&buy_products.Buy_ProductUserModel{})
>>>>>>> 54f4ab9460419a42b520998c60f9fa0be7b23b8d
}

func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	MakeRoutes(r)

	v1 := r.Group("/api")

	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))
<<<<<<< HEAD
	articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	articles.TagsAnonymousRegister(v1.Group("/tags"))

	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))
	users.ProfileRegister(v1.Group("/profiles"))
	articles.ArticlesRegister(v1.Group("/articles"))
=======
	products.ProductsAnonymousRegister(v1.Group("/products"))


	v1.Use(users.AuthMiddleware(true))
	buy_products.Buy_ProductsRegister(v1.Group("/buy_products"))
	users.UserRegister(v1.Group("/user"))
	users.ProfileRegister(v1.Group("/profiles"))
>>>>>>> 54f4ab9460419a42b520998c60f9fa0be7b23b8d

	// testAuth := r.Group("/api/ping")
	// testAuth.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// test 1 to 1
	// tx1 := db.Begin()
	// userA := users.UserModel{
	// 	Username: "AAAAAAAAAAAAAAAA",
	// 	Email:    "aaaa@g.cn",
	// 	Bio:      "hehddeda",
	// 	Image:    nil,
	// }
	// tx1.Save(&userA)
	// tx1.Commit()
	// fmt.Println(userA)

<<<<<<< HEAD
	//db.Save(&ArticleUserModel{
	//    UserModelID:userA.ID,
	//})
	//var userAA ArticleUserModel
	//db.Where(&ArticleUserModel{
=======
	//db.Save(&Buy_ProductUserModel{
	//    UserModelID:userA.ID,
	//})
	//var userAA Buy_ProductUserModel
	//db.Where(&Buy_ProductUserModel{
>>>>>>> 54f4ab9460419a42b520998c60f9fa0be7b23b8d
	//    UserModelID:userA.ID,
	//}).First(&userAA)
	//fmt.Println(userAA)

	//r.Run() // listen and serve on 0.0.0.0:8080
	fmt.Printf("0.0.0.0:3000")
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

		/*
			fmt.Printf("c.Request.Method \n")
			fmt.Printf(c.Request.Method)
			fmt.Printf("c.Request.RequestURI \n")
			fmt.Printf(c.Request.RequestURI)
		*/
	}
	r.Use(cors)
<<<<<<< HEAD
}
=======
}
>>>>>>> 54f4ab9460419a42b520998c60f9fa0be7b23b8d

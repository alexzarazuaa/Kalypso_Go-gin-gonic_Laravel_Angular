package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"goUsers/common"
	"goUsers/src"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
}

func main() {


	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	MakeRoutes(r)

	v1 := r.Group("/api")


	v1.Use(users.AuthMiddleware(false))
	users.UsersRegister(v1.Group("/users"))


	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/users"))
	users.ProfileRegister(v1.Group("/profiles"))

	// testAuth := r.Group("/api/ping")
	// testAuth.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// test 1 to 1
	// tx1 := db.Begin()
	// userA := users.Users{
	// 	Username: "AAAAAAAAAAAAAAAA",
	// 	Email:    "aaaa@g.cn",
	// 	Bio:      "hehddeda",
	// 	Image:    nil,
	// }
	// tx1.Save(&userA)
	// tx1.Commit()
	// fmt.Println(userA)

	//db.Save(&Buy_ProductUsers{
	//    UsersID:userA.ID,
	//})
	//var userAA Buy_ProductUsers
	//db.Where(&Buy_ProductUsers{
	//    UsersID:userA.ID,
	//}).First(&userAA)
	//fmt.Println(userAA)

	//r.Run() // listen and serve on 0.0.0.0:3002
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

		/*
			fmt.Printf("c.Request.Method \n")
			fmt.Printf(c.Request.Method)
			fmt.Printf("c.Request.RequestURI \n")
			fmt.Printf(c.Request.RequestURI)
		*/
	}
	r.Use(cors)
}
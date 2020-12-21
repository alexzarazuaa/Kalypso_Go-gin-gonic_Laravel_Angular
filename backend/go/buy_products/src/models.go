package buy_products

import (
	"time"
	"strconv"
	"github.com/jinzhu/gorm"
	"gobuys_products/common"
)

type Buy_ProductModel struct {
	gorm.Model
	Slug         string  `gorm:"unique_index"`
	Name     	 string  `gorm:"column:name"`
	Brand        string  `gorm:"column:brand"`
	Img        	 string  `gorm:"column:img"`
	Description  string  `gorm:"column:description"`
	Rating 		 int     `gorm:"column:rating"`
	Category 	 string  `gorm:"column:category"`
	Author      Buy_ProductUsers
	AuthorID    uint
}

type Users struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
	Bearer     	 string  `gorm:"column:bearer"`
	Karma 		 int  `gorm:"column:karma"`
	Type 		 string  `gorm:"column:type"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}



type Buy_ProductUsers struct {
	gorm.Model
	Users      Users
	UsersID    uint
	Buy_ProductModels  []Buy_ProductModel  `gorm:"ForeignKey:AuthorID"`
}


func GetBuy_ProductUsers(userModel Users) Buy_ProductUsers {
	var buy_productUsers Buy_ProductUsers
	if userModel.ID == 0 {
		return buy_productUsers
	}
	db := common.GetDB()
	db.Where(&Buy_ProductUsers{
		UsersID: userModel.ID,
	}).FirstOrCreate(&buy_productUsers)
	buy_productUsers.Users = userModel
	return buy_productUsers
}

func FindManyBuy_Products() ([]Buy_ProductModel, int, error) {
	db := common.GetDB()
	var count int
	var models []Buy_ProductModel
	db.Model(&models).Count(&count)
	err :=db.Find(&models).Error
	return models,count, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneBuy_Product(condition interface{}) (Buy_ProductModel, error) {
	db := common.GetDB()
	var model Buy_ProductModel
	err :=db.Where(condition).First(&model).Error
	return model, err
}

func (self *Buy_ProductUsers) GetBuy_ProductFeed(limit, offset string) ([]Buy_ProductModel, int, error) {
	db := common.GetDB()
	var models []Buy_ProductModel
	var count int

	offset_int, err := strconv.Atoi(offset)
	if err != nil {
		offset_int = 0
	}
	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		limit_int = 20
	}

	tx := db.Begin()
	//followings := self.Users.GetFollowings()
	 var buy_productUserss []uint
	// for _, following := range followings {
	// 	buy_productUsers := GetBuy_ProductUsers(following)
	// 	buy_productUserss = append(buy_productUserss, buy_productUsers.ID)
	// }

	tx.Where("author_id in (?)", buy_productUserss).Offset(offset_int).Limit(limit_int).Find(&models)

	for i, _ := range models {
		tx.Model(&models[i]).Related(&models[i].Author, "Author")
		tx.Model(&models[i].Author).Related(&models[i].Author.Users)
	}
	err = tx.Commit().Error
	return models, count, err
}
package buy_products

import (
	"strconv"
	"github.com/jinzhu/gorm"

	"goKa/common"
	"goKa/users"
)

type Buy_ProductModel struct {
	gorm.Model
	Slug         string `gorm:"unique_index"`
	Name     	 string  `gorm:"column:name"`
	Brand        string  `gorm:"column:brand"`
	Img        	 string  `gorm:"column:img"`
	Description  string  `gorm:"column:description"`
	Rating 		 int     `gorm:"column:rating"`
	Category 	 string  `gorm:"column:category"`
	Author      Buy_ProductUserModel
	AuthorID    uint
}

type Buy_ProductUserModel struct {
	gorm.Model
	UserModel      users.UserModel
	UserModelID    uint
	Buy_ProductModels  []Buy_ProductModel  `gorm:"ForeignKey:AuthorID"`
}


func GetBuy_ProductUserModel(userModel users.UserModel) Buy_ProductUserModel {
	var buy_productUserModel Buy_ProductUserModel
	if userModel.ID == 0 {
		return buy_productUserModel
	}
	db := common.GetDB()
	db.Where(&Buy_ProductUserModel{
		UserModelID: userModel.ID,
	}).FirstOrCreate(&buy_productUserModel)
	buy_productUserModel.UserModel = userModel
	return buy_productUserModel
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

func (self *Buy_ProductUserModel) GetBuy_ProductFeed(limit, offset string) ([]Buy_ProductModel, int, error) {
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
	followings := self.UserModel.GetFollowings()
	var buy_productUserModels []uint
	for _, following := range followings {
		buy_productUserModel := GetBuy_ProductUserModel(following)
		buy_productUserModels = append(buy_productUserModels, buy_productUserModel.ID)
	}

	tx.Where("author_id in (?)", buy_productUserModels).Order("updated_at desc").Offset(offset_int).Limit(limit_int).Find(&models)

	for i, _ := range models {
		tx.Model(&models[i]).Related(&models[i].Author, "Author")
		tx.Model(&models[i].Author).Related(&models[i].Author.UserModel)
	}
	err = tx.Commit().Error
	return models, count, err
}
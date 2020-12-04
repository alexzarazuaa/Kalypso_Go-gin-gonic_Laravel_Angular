package products

import (
	"strconv"
	"github.com/jinzhu/gorm"

	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/common"
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/users"
)

type ProductModel struct {
	gorm.Model
	Slug         string `gorm:"unique_index"`
	Name     	 string  `gorm:"column:name"`
	Brand        string  `gorm:"column:brand"`
	Img        	 string  `gorm:"column:img"`
	Description  string  `gorm:"column:description"`
	Rating 		 int     `gorm:"column:rating"`
	Category 	 string  `gorm:"column:category"`
	Author      ProductUserModel
	AuthorID    uint
}

type ProductUserModel struct {
	gorm.Model
	UserModel      users.UserModel
	UserModelID    uint
	ProductModels  []ProductModel  `gorm:"ForeignKey:AuthorID"`
}


func GetProductUserModel(userModel users.UserModel) ProductUserModel {
	var productUserModel ProductUserModel
	if userModel.ID == 0 {
		return productUserModel
	}
	db := common.GetDB()
	db.Where(&ProductUserModel{
		UserModelID: userModel.ID,
	}).FirstOrCreate(&productUserModel)
	productUserModel.UserModel = userModel
	return productUserModel
}

func FindManyProducts() ([]ProductModel, int, error) {
	db := common.GetDB()
	var count int
	var models []ProductModel
	db.Model(&models).Count(&count)
	err :=db.Find(&models).Error
	return models,count, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneProduct(condition interface{}) (ProductModel, error) {
	db := common.GetDB()
	var model ProductModel
	err :=db.Where(condition).First(&model).Error
	return model, err
}

func (self *ProductUserModel) GetProductFeed(limit, offset string) ([]ProductModel, int, error) {
	db := common.GetDB()
	var models []ProductModel
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
	var productUserModels []uint
	for _, following := range followings {
		productUserModel := GetProductUserModel(following)
		productUserModels = append(productUserModels, productUserModel.ID)
	}

	tx.Where("author_id in (?)", productUserModels).Order("updated_at desc").Offset(offset_int).Limit(limit_int).Find(&models)

	for i, _ := range models {
		tx.Model(&models[i]).Related(&models[i].Author, "Author")
		tx.Model(&models[i].Author).Related(&models[i].Author.UserModel)
	}
	err = tx.Commit().Error
	return models, count, err
}
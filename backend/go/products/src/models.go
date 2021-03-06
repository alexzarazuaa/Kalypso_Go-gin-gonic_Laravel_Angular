package products

import (
	// "fmt"
	"time"
	"strconv"
	"github.com/jinzhu/gorm"
	"goProducts/common"

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
	Author      ProductUsers
	AuthorID    uint 
}

type BrandsKarma struct {
	gorm.Model
	Name         string `gorm:"unique_index"`
	Rating 		 int     `gorm:"column:rating"`
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


type ProductUsers struct {
	gorm.Model
	Users      Users
	UsersID    uint
	FavoriteModels []FavoriteModel `gorm:"ForeignKey:FavoriteByID"`
}


type FavoriteModel struct {
	gorm.Model
	Favorite     ProductModel
	FavoriteID   uint
	FavoriteBy   ProductUsers
	FavoriteByID uint
}


func GetProductUsers(user Users) ProductUsers {
	var productUsers ProductUsers
	if user.ID == 0 {
		return productUsers
	}
	db := common.GetDB()

	db.Where(&ProductUsers{
		UsersID: user.ID,
	}).FirstOrCreate(&productUsers)
	productUsers.Users = user
	return productUsers
}

func FindManyProducts(limit, offset, favorited string) ([]ProductModel, int, error) {
	db := common.GetDB()
	var count int
	var models []ProductModel
	tx := db.Begin()

	offset_int, err := strconv.Atoi(offset)
	if err != nil {
		offset_int = 0
	}

	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		limit_int = 20
	}
	if favorited != "" {
		var userModel Users
		tx.Where(Users{Username: favorited}).First(&userModel)
		productUserModel := GetProductUsers(userModel)
		if productUserModel.ID != 0 {
			var favoriteModels []FavoriteModel
			tx.Where(FavoriteModel{
				FavoriteByID: productUserModel.ID,
			}).Offset(offset_int).Limit(limit_int).Find(&favoriteModels)

			count = tx.Model(&productUserModel).Association("FavoriteModels").Count()
			for _, favorite := range favoriteModels {
				var model ProductModel
				tx.Model(&favorite).Related(&model, "Favorite")
				models = append(models, model)
			}
		}
	}
	db.Model(&models).Count(&count)
	err =db.Find(&models).Error
	return models,count, err
}

func (product ProductModel) favoritesCount() uint {
	db := common.GetDB()
	var count uint
	db.Model(&FavoriteModel{}).Where(FavoriteModel{
		FavoriteID: product.ID,
	}).Count(&count)
	return count
}

func (product ProductModel) isFavoriteBy(user ProductUsers) bool {
	db := common.GetDB()
	var favorite FavoriteModel

	if ( user.UsersID != 0){
		db.Where(FavoriteModel{
			FavoriteID:   product.ID,
			FavoriteByID: user.UsersID,
		}).First(&favorite)
	}
	return favorite.ID != 0
}

func (product ProductModel) favoriteBy(user ProductUsers) error {
	db := common.GetDB()
	var favorite FavoriteModel
	err := db.FirstOrCreate(&favorite, &FavoriteModel{
		FavoriteID:   product.ID,
		FavoriteByID: user.UsersID,
	}).Error
	return err
}

func (product ProductModel) unFavoriteBy(user ProductUsers) error {
	db := common.GetDB()
	err := db.Where(FavoriteModel{
		FavoriteID:   product.ID,
		FavoriteByID: user.UsersID,
	}).Delete(FavoriteModel{}).Error
	return err
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

	// db := common.GetDB()
	// var model ProductModel
	// tx := db.Begin()
	// tx.Where(condition).First(&model)
	// tx.Model(&model).Related(&model.Author, "Author")
	// tx.Model(&model.Author).Related(&model.Author.UserModel)
	// // tx.Model(&model).Related(&model.Tags, "Tags")
	// err := tx.Commit().Error
	// return model, err
}

func (self *ProductUsers) GetProductFeed(limit, offset string) ([]ProductModel, int, error) {
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
	var productUserss []uint

	tx.Where("author_id in (?)", productUserss).Offset(offset_int).Limit(limit_int).Find(&models)

	for i, _ := range models {
		tx.Model(&models[i]).Related(&models[i].Author, "Author")
		tx.Model(&models[i].Author).Related(&models[i].Author.Users)
	}
	err = tx.Commit().Error
	return models, count, err
}

func ProductsbyBrands(condition interface{}) ([]ProductModel, error) {
	db := common.GetDB()
	var model []ProductModel
	err :=db.Where(condition).Find(&model).Error
	return model, err
}

func GetBrands() ([]BrandsKarma, error) {
	db := common.GetDB()
	var model []BrandsKarma
	err :=db.Order("rating desc").Find(&model).Error
	return model, err
}

func FindOneUser(condition interface{}) (Users, error) {
	db := common.GetDB()
	var model Users
	err :=db.Where(condition).First(&model).Error
	return model, err
}
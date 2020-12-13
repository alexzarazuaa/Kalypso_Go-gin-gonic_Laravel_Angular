package products

// import (
// 	"strconv"
// 	"github.com/jinzhu/gorm"

// 	"goKa/common"
// 	"goKa/users"
// )

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

type ProductUsers struct {
	gorm.Model
	Users      users.Users
	UsersID    uint
	ProductModels  []ProductModel  `gorm:"ForeignKey:AuthorID"`
}


// func GetProductUsers(userModel users.Users) ProductUsers {
// 	var productUsers ProductUsers
// 	if userModel.ID == 0 {
// 		return productUsers
// 	}
// 	db := common.GetDB()
// 	db.Where(&ProductUsers{
// 		UsersID: userModel.ID,
// 	}).FirstOrCreate(&productUsers)
// 	productUsers.Users = userModel
// 	return productUsers
// }

// func FindManyProducts() ([]ProductModel, int, error) {
// 	db := common.GetDB()
// 	var count int
// 	var models []ProductModel
// 	db.Model(&models).Count(&count)
// 	err :=db.Find(&models).Error
// 	return models,count, err
// }

// func SaveOne(data interface{}) error {
// 	db := common.GetDB()
// 	err := db.Save(data).Error
// 	return err
// }

// func FindOneProduct(condition interface{}) (ProductModel, error) {
// 	db := common.GetDB()
// 	var model ProductModel
// 	err :=db.Where(condition).First(&model).Error
// 	return model, err
// }

// func (self *ProductUsers) GetProductFeed(limit, offset string) ([]ProductModel, int, error) {
// 	db := common.GetDB()
// 	var models []ProductModel
// 	var count int

// 	offset_int, err := strconv.Atoi(offset)
// 	if err != nil {
// 		offset_int = 0
// 	}
// 	limit_int, err := strconv.Atoi(limit)
// 	if err != nil {
// 		limit_int = 20
// 	}

// 	tx := db.Begin()
// 	followings := self.Users.GetFollowings()
// 	var productUserss []uint
// 	for _, following := range followings {
// 		productUsers := GetProductUsers(following)
// 		productUserss = append(productUserss, productUsers.ID)
// 	}

// 	tx.Where("author_id in (?)", productUserss).Offset(offset_int).Limit(limit_int).Find(&models)

// 	for i, _ := range models {
// 		tx.Model(&models[i]).Related(&models[i].Author, "Author")
// 		tx.Model(&models[i].Author).Related(&models[i].Author.Users)
// 	}
// 	err = tx.Commit().Error
// 	return models, count, err
// }
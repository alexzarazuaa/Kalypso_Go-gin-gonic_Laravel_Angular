package articles

import (
	"strconv"
	"github.com/jinzhu/gorm"

	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/common"
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/users"
)

type ArticleModel struct {
	gorm.Model
	Slug         string `gorm:"unique_index"`
	Name     	 string  `gorm:"column:name"`
	Brand        string  `gorm:"column:brand"`
	Img        	 string  `gorm:"column:img"`
	Description  string  `gorm:"column:description"`
	Rating 		 int     `gorm:"column:rating"`
	Category 	 string  `gorm:"column:category"`
	Author      ArticleUserModel
	AuthorID    uint
}

type ArticleUserModel struct {
	gorm.Model
	UserModel      users.UserModel
	UserModelID    uint
	ArticleModels  []ArticleModel  `gorm:"ForeignKey:AuthorID"`
}


func GetArticleUserModel(userModel users.UserModel) ArticleUserModel {
	var articleUserModel ArticleUserModel
	if userModel.ID == 0 {
		return articleUserModel
	}
	db := common.GetDB()
	db.Where(&ArticleUserModel{
		UserModelID: userModel.ID,
	}).FirstOrCreate(&articleUserModel)
	articleUserModel.UserModel = userModel
	return articleUserModel
}

func FindManyArticles() ([]ArticleModel, int, error) {
	db := common.GetDB()
	var count int
	var models []ArticleModel
	db.Model(&models).Count(&count)
	err :=db.Find(&models).Error
	return models,count, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneArticle(condition interface{}) (ArticleModel, error) {
	db := common.GetDB()
	var model ArticleModel
	err :=db.Where(condition).First(&model).Error
	return model, err
}

func (self *ArticleUserModel) GetArticleFeed(limit, offset string) ([]ArticleModel, int, error) {
	db := common.GetDB()
	var models []ArticleModel
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
	var articleUserModels []uint
	for _, following := range followings {
		articleUserModel := GetArticleUserModel(following)
		articleUserModels = append(articleUserModels, articleUserModel.ID)
	}

	tx.Where("author_id in (?)", articleUserModels).Order("updated_at desc").Offset(offset_int).Limit(limit_int).Find(&models)

	for i, _ := range models {
		tx.Model(&models[i]).Related(&models[i].Author, "Author")
		tx.Model(&models[i].Author).Related(&models[i].Author.UserModel)
	}
	err = tx.Commit().Error
	return models, count, err
}

func (model *ArticleModel) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Update(data).Error
	return err
}

func DeleteArticleModel(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(ArticleModel{}).Error
	return err
}

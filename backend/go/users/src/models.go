package users

import (
	// "fmt"
	"errors"
	"github.com/jinzhu/gorm"
	"goUsers/common"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Users struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username;unique_index"`
	Email        string  `gorm:"column:email;unique_index"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
	Bearer     	 string  `gorm:"column:bearer"`
	Karma 		 int  `gorm:"column:karma"`
	Type 		 string  `gorm:"column:type"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type FollowModel struct {
	gorm.Model
	Following    Users
	FollowingID  uint
	FollowedBy   Users
	FollowedByID uint
}


type BrandsKarma struct {
	gorm.Model
	Name         string `gorm:"unique_index"`
	Karma     	 int  `gorm:"column:Karma"`
}

// Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&Users{})
	db.AutoMigrate(&FollowModel{})
}

func (u *Users) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

// Database will only save the hashed string, you should check it by util function.
// 	if err := serModel.checkPassword("password0"); err != nil { password error }
func (u *Users) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// You could input the conditions and it will return an Users in database with error info.
// 	userModel, err := FindOneUser(&Users{Username: "username0"})
func FindOneUser(condition interface{}) (Users, error) {
	db := common.GetDB()
	var model Users
	err := db.Where(condition).First(&model).Error
	return model, err
}

// You could input an Users which will be saved in database returning with error info
// 	if err := SaveOne(&userModel); err != nil { ... }
func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

// You could update properties of an Users to database returning with error info.
//  err := db.Model(userModel).Update(Users{Username: "wangzitian0"}).Error
func (model *Users) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Update(data).Error
	return err
}


func (model *Users) InsertToken(data interface{}) error {
	db := common.GetDB()
	err :=db.Model(&model).Update(data).Error
	return err
}

// You could add a following relationship as userModel1 following userModel2
// 	err = userModel1.following(userModel2)
func (u Users) following(v Users) error {
	db := common.GetDB()
	var follow FollowModel
	err := db.FirstOrCreate(&follow, &FollowModel{
		FollowingID:  v.ID,
		FollowedByID: u.ID,
	}).Error
	return err
}

// You could check whether  userModel1 following userModel2
// 	followingBool = myUsers.isFollowing(self.Users)
func (u Users) isFollowing(v Users) bool {
	db := common.GetDB()
	var follow FollowModel
	db.Where(FollowModel{
		FollowingID:  v.ID,
		FollowedByID: u.ID,
	}).First(&follow)
	return follow.ID != 0
}

// You could delete a following relationship as userModel1 following userModel2
// 	err = userModel1.unFollowing(userModel2)
func (u Users) unFollowing(v Users) error {
	db := common.GetDB()
	err := db.Where(FollowModel{
		FollowingID:  v.ID,
		FollowedByID: u.ID,
	}).Delete(FollowModel{}).Error
	return err
}

// You could get a following list of userModel
// 	followings := userModel.GetFollowings()
func (u Users) GetFollowings() []Users {
	db := common.GetDB()
	tx := db.Begin()
	var follows []FollowModel
	var followings []Users
	tx.Where(FollowModel{
		FollowedByID: u.ID,
	}).Find(&follows)
	for _, follow := range follows {
		var userModel Users
		tx.Model(&follow).Related(&userModel, "Following")
		followings = append(followings, userModel)
	}
	tx.Commit()
	return followings
}

func UpdateBrands(name string, karma int) (error) {
	db := common.GetDB()
	err:= db.Model(BrandsKarma{}).Where("name = ?", name).Updates(BrandsKarma{Karma: karma}).Error
	return err
}
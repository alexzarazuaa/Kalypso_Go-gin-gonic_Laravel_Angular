package users

import (
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/common"
)

//-----------------PROFILE-----------------------------//

type ProfileSerializer struct {
	C *gin.Context
	UserModel
}

// Declare your response schema here
type ProfileResponse struct {
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Image     *string `json:"image"`
	Karma 	   int    `json:"karma"`
}

// Put your response logic including wrap the userModel here.
func (self *ProfileSerializer) Response() ProfileResponse {
	profile := ProfileResponse{
		Username:  self.Username,
		Image:     self.Image,
		Karma:	   self.Karma,
		Email:	   self.Email,
	}
	return profile
}

//----------------END PROFILE-------------------------//



//-------------------LOGIN---------------------------//

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Image     *string `json:"image"`
	Karma 	   int    `json:"karma"`
	Type	   string  `json:"type"`
	Token    string  `json:"token"`
}

func (self *UserSerializer) Response() UserResponse {
	myUserModel := self.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Image:    myUserModel.Image,
		Karma: 	   myUserModel.Karma,
		Type:	  myUserModel.Type,
		Token:    common.GenToken(myUserModel.ID),
	}
	return user
}

//--------------------END LOGIN------------------------//



//--------------------ADMIN---------------------------//
type AdminSerializer struct {
	C *gin.Context
	UserModel
}

type AdminResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Image     *string `json:"image"`
	Karma 	   int    `json:"karma"`
	Type	   string  `json:"type"`
}

func (self *AdminSerializer) Response() AdminResponse {
	admin := AdminResponse{
		Username:  self.Username,
		Image:     self.Image,
		Karma:	   self.Karma,
		Email:	   self.Email,
		Type:	   self.Type,
	}
	return admin
}

//------------------------END ADMIN---------------------//

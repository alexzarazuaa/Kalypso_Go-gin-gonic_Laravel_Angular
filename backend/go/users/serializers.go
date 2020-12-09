package users

import (
	"github.com/gin-gonic/gin"

	"goKa/common"
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



//--------------------NO NORMAL TYPE---------------------------//

type NoTypeSerializer struct {
	C *gin.Context
	UserModel
}

type NoTypeResponse struct {
	Type	   string  `json:"type"`
}

func (self *NoTypeSerializer) Response() NoTypeResponse {
	user := NoTypeResponse{
		Type:	   self.Type,
	}
	return user
}

//------------------------END NO NORMAL TYPE---------------------//
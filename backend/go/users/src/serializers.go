package users

import (
	"github.com/gin-gonic/gin"
	"goApp/common"
)

//-----------------PROFILE-----------------------------//

type ProfileSerializer struct {
	C *gin.Context
	Users
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
	myUsers := self.c.MustGet("my_user_model").(Users)
	user := UserResponse{
		Username: myUsers.Username,
		Email:    myUsers.Email,
		Image:    myUsers.Image,
		Karma: 	   myUsers.Karma,
		Type:	  myUsers.Type,
		Token:    common.GenToken(myUsers.ID),
	}
	return user
}

//--------------------END LOGIN------------------------//



//--------------------ADMIN---------------------------//

type AdminSerializer struct {
	C *gin.Context
	Users
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


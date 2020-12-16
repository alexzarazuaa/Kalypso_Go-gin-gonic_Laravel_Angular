package users

import (
	"fmt"
	"strings"
	"errors"
	"goKa/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", UsersRegistration)
	router.POST("/login", UsersLogin)
}

func UserRegister(router *gin.RouterGroup) {
	router.GET("/", UserRetrieve)
	router.PUT("/", UserUpdate)
}

func ProfileRegister(router *gin.RouterGroup) {
	router.GET("/:username", ProfileRetrieve)
	router.POST("/:username/follow", ProfileFollow)
	router.DELETE("/:username/follow", ProfileUnfollow)
}

func ProfileRetrieve(c *gin.Context) {
	username := c.Param("username")
	userModel, err := FindOneUser(&Users{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("Invalid username")))
		return
	}
	profileSerializer := ProfileSerializer{c, userModel}
	c.JSON(http.StatusOK, gin.H{"profile": profileSerializer.Response()})
}

func ProfileFollow(c *gin.Context) {
	username := c.Param("username")
	userModel, err := FindOneUser(&Users{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("Invalid username")))
		return
	}
	myUsers := c.MustGet("my_user_model").(Users)
	err = myUsers.following(userModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := ProfileSerializer{c, userModel}
	c.JSON(http.StatusOK, gin.H{"profile": serializer.Response()})
}

func ProfileUnfollow(c *gin.Context) {
	username := c.Param("username")
	userModel, err := FindOneUser(&Users{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("Invalid username")))
		return
	}
	myUsers := c.MustGet("my_user_model").(Users)

	err = myUsers.unFollowing(userModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := ProfileSerializer{c, userModel}
	c.JSON(http.StatusOK, gin.H{"profile": serializer.Response()})
}

func UsersRegistration(c *gin.Context) {
	userModelValidator := NewUsersValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("my_user_model", userModelValidator.userModel)
	serializer := UserSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func UsersLogin(c *gin.Context) {

	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	userModel, err := FindOneUser(&Users{Email: loginValidator.userModel.Email})

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	if userModel.checkPassword(loginValidator.User.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}


	if ((userModel.Type)=="client"){	//Type client -> Login

		UpdateContextUsers(c, userModel.ID)
		serializer := UserSerializer{c}
		c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})

	}else if ((userModel.Type)=="admin"){	//Type admin -> show user information
		// create bearer
		serializer1 := UserSerializer{c}
		bearer:=serializer1.Response().Token

		// insert bearer in DB
		err := userModel.InsertToken(&Users{Bearer:bearer})
		if err != nil {
			c.JSON(http.StatusNotFound, common.NewError("DB", errors.New("Error Update Login Admin")))
			return
		}
		
		//create code to save in json redis
		s := strings.Split(userModel.Email, "@") 
		code:= fmt.Sprint(userModel.ID) + userModel.Username[0:3] + s[0][len(s[0])-3:] + userModel.Username[len(userModel.Username)-2:] + bearer[0:5]
		
		//Create json to save in redis
		client := common.NewClient()
		user:=`{"username":"`+userModel.Username+`", "email":"`+userModel.Email+`", "type": "`+userModel.Type+`", "bearer":"`+bearer+`", "code":"`+code+`"}`

		// fmt.Println(code)
			// user:=`{"hola":"adios"}`
		//save json in redis
		err_redis := common.SetUser(userModel.Email, user, client)
		if err_redis != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		 //all right
		serializer := AdminSerializer{c, userModel}
		c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})

	} else{		//No normal type -> show type
		
		c.JSON(http.StatusOK, gin.H{"Does not have a normal type": userModel.Type })

	}

}

func UserRetrieve(c *gin.Context) {
	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

func UserUpdate(c *gin.Context) {
	myUsers := c.MustGet("my_user_model").(Users)
	userModelValidator := NewUsersValidatorFillWith(myUsers)
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	userModelValidator.userModel.ID = myUsers.ID
	if err := myUsers.Update(userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	UpdateContextUsers(c, myUsers.ID)
	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}
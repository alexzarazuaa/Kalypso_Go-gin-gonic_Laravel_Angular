package products

import (
	// "fmt"
	// "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"goProducts/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Strips 'TOKEN ' prefix from token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}

// Extract  token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

// Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

// // A helper to write user_id and user_model to the context
func UpdateContextUsers(c *gin.Context, my_user_id uint) {
	var myUsers Users
	if my_user_id != 0 {
		db := common.GetDB()
		db.First(&myUsers, my_user_id)
	}
	c.Set("my_user_id", my_user_id)
	c.Set("my_user_model", myUsers)
}


func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextUsers(c, 0)
		client := common.NewClient()

		err, val := common.Get("user", client)

		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
				}
			return
		}
		
		uncrypt:= strings.Split(val, "*")

		bearerR:=uncrypt[0]+ `.`+uncrypt[2]+`.`+uncrypt[4]
		mail:=uncrypt[1]+`@`+uncrypt[3]

		bearerH, err:=stripBearerPrefixFromTokenString(strings.Join(c.Request.Header["Authorization"]," "))

			if ( bearerH != bearerR) {

				if auto401 {
					c.AbortWithError(http.StatusUnauthorized, err)
				}
				return
			}

			user, err_user := FindOneUser(&Users{Email: mail})

			if (err_user != nil){
				c.AbortWithError(http.StatusUnauthorized, err_user)
				return
			}

			UpdateContextUsers(c, user.ID)
	}
}


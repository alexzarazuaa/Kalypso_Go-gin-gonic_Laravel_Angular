package articles

import (
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/common"
	"github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular/backend/go/users"
	"github.com/gosimple/slug"
	"gopkg.in/gin-gonic/gin.v1"
)

type ArticleModelValidator struct {
	Article struct {
		Name       string   `form:"name" json:"name" binding:"exists,min=4"`
		Description string   `form:"description" json:"description" binding:"max=2048"`
		Brand        string   `form:"brand" json:"brand" binding:"max=2048"`
		Category        string   `form:"category" json:"category" binding:"max=2048"`

	} `json:"article"`
	articleModel ArticleModel `json:"-"`
}

func NewArticleModelValidator() ArticleModelValidator {
	return ArticleModelValidator{}
}

func NewArticleModelValidatorFillWith(articleModel ArticleModel) ArticleModelValidator {
	articleModelValidator := NewArticleModelValidator()
	return articleModelValidator
}

func (s *ArticleModelValidator) Bind(c *gin.Context) error {
	myUserModel := c.MustGet("my_user_model").(users.UserModel)

	err := common.Bind(c, s)
	if err != nil {
		return err
	}
		
	s.articleModel.Slug = slug.Make(s.Article.Name)
	s.articleModel.Name=(s.Article.Name)
	s.articleModel.Brand=(s.Article.Brand)
	s.articleModel.Description = s.Article.Description
	s.articleModel.Rating=0
	s.articleModel.Description = s.Article.Category
	s.articleModel.Author = GetArticleUserModel(myUserModel)

	return nil
}

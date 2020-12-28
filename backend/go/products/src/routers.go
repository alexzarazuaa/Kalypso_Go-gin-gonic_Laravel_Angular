package products

import (
	"fmt"
	"time"
	"strconv"
	"encoding/json"
	"errors"
	"net/http"
	"goProducts/common"
	"github.com/gin-gonic/gin"
)

func ProductsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", ProductList)
	router.GET("/:slug", ProductRetrieve)
	router.POST("/object/:key/:object/:point", Karma_redis)
	router.PUT("/", Proof)
}


//General function to karma in products and brands
func Karma_redis(c *gin.Context){
	client := common.NewClient()
	key := c.Param("key")
	object := c.Param("object")
	point_str:=c.Param("point")
	point, err_int:=strconv.Atoi(point_str)//string to int
	 
	if(err_int!=nil){
		c.JSON(http.StatusBadRequest, gin.H{"error": err_int.Error()})
		return
	}

	//obtain data(brands or products) from redis 
	err_get, val := common.Get(key, client)
	if err_get != nil {//If not exist data in redis we storege first brand or product

		objects := map[string]int{ object: point }
		err_SetMarshal:= SetMarshal(objects,key)//Object -> Byte and storage in redis
		if(err_SetMarshal!= nil){//Any mistakes
			c.JSON(http.StatusBadRequest, gin.H{"error": err_SetMarshal.Error()})
			return
		}

	c.JSON(http.StatusOK, gin.H{"result": "okey_first_store_object_redis"})
	return
	}

	//If exists data in redis 
	objects := map[string]int{}
	json.Unmarshal([]byte(val), &objects)

	if((objects[object])==0){//If This brand or product not stored in redis
		objects[object]=point
	
	}else{//If this brand or product is stored in redis
		objects[object]+=point
	}

	//Then stored data in redis
	err_SetMarshal:= SetMarshal(objects,key)
	if(err_SetMarshal!= nil){
		c.JSON(http.StatusBadRequest, gin.H{"error": err_SetMarshal.Error()})
		return
	}

	//At 10 minutes
	timeDelay := 600000 * time.Millisecond

	var endTime <-chan time.Time
	
	endTime = time.After(timeDelay)

    for {
        select {
        case <-endTime:
		
			//We pass redis(products & brands) in database
			for key, element := range objects {
				err_update := UpdateBrands(key, element)

				if (err_update!=nil){
					c.JSON(http.StatusBadRequest, gin.H{"error": err_update.Error()})
					return
				}
			}
			c.JSON(http.StatusOK, gin.H{"result": objects})
			return
        }
	}

}

func Proof(c *gin.Context){
 

}
func SetMarshal(objects map[string]int, key string )  error {
	client := common.NewClient()


	objects_marshal, err_marshal := json.Marshal(objects)
	if err_marshal != nil {
		return err_marshal
	}

	err_set := common.Set(key, objects_marshal, client)
	if err_set != nil {
		return err_set
	}

	return nil
}

func ProductList(c *gin.Context) {
	productModels, modelCount, err := FindManyProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid param")))
		return
	}
	serializer := ProductsSerializer{c, productModels}
	c.JSON(http.StatusOK, gin.H{"products": serializer.Response(), "productsCount": modelCount})
}

func ProductFeed(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	myUsers := c.MustGet("my_user_model").(Users)
	if myUsers.ID == 0 {
		c.AbortWithError(http.StatusUnauthorized, errors.New("{error : \"Require auth!\"}"))
		return
	}
	productUsers := GetProductUsers(myUsers)
	productModels, modelCount, err := productUsers.GetProductFeed(limit, offset)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid param")))
		return
	}
	serializer := ProductsSerializer{c, productModels}
	c.JSON(http.StatusOK, gin.H{"products": serializer.Response(), "productsCount": modelCount})
}

func ProductRetrieve(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "feed" {
		ProductFeed(c)
		return
	}
	productModel, err := FindOneProduct(&ProductModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid slug")))
		return
	}
	serializer := ProductSerializer{c, productModel}
	c.JSON(http.StatusOK, gin.H{"product": serializer.Response()})
}


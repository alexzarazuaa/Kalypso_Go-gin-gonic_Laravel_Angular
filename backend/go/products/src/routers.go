package products

import (
	"fmt"
	"sort"
	"strings"
	// "reflect"
	// "time"
	// "strconv"
	"encoding/json"
	"errors"
	"net/http"
	"goProducts/common"
	"github.com/gin-gonic/gin"
)

func ProductsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/:slug", ProductList)
	// router.PUT("/:types", Proof)
}


func ProductsRegister(router *gin.RouterGroup) {
	router.POST("/:slug/favorite", ProductFavorite)
	router.DELETE("/:slug/favorite", ProductUnfavorite)
}

// func Proof(c *gin.Context){
// 	client := common.NewClient()
// 	types := c.Param("types")
// 	err_get, val := common.Get(types, client)
// 	fmt.Println(err_get)
// 	fmt.Println(val)
// }

//General function to karma in products and brands
func Karma_redis( types string, id string, karma int) error{
	client := common.NewClient()
	fmt.Println("-------------------------------------")

	//obtain data(brands or products) from redis 
	err_get, val := common.Get(types, client)
	if err_get != nil {//If not exist data in redis we storege first brand or product

		fmt.Println("NO HAY NADAAA")

		objects := map[string]int{ id: karma }
		err_SetMarshal:= SetMarshal(objects,types)//Object -> Byte and storage in redis

		if (err_SetMarshal!= nil){//Any mistakes
			return err_SetMarshal
		}


		fmt.Println(" HAY NADAAA")

	return nil
	}

	//If exists data in redis 
	objects := map[string]int{}
	json.Unmarshal([]byte(val), &objects)

	if ((objects[id])==0){//If This brand or product not stored in redis
		objects[id]=karma
	
	}else{//If this brand or product is stored in redis
		objects[id]+=karma
	}

	//Then stored data in redis
	err_SetMarshal:= SetMarshal(objects,types)
	if(err_SetMarshal!= nil){
		return err_SetMarshal
	}

	// //At 10 minutes
	// // timeDelay := 60000typytypess0 * time.Millisecond

	// // var endTime <-chan time.Time
	
	// // endTime = time.After(timeDelay)

    // // for {
    // //     select {
    // //     case <-endTime:
		
	// // 		//We pass redis(products & brands) in database
	// // 		for key, element := range objects {
	// // 			err_update := UpdateBrands(key, element)

	// // 			if (err_update!=nil){
	// // 				c.JSON(http.StatusBadRequest, gin.H{"error": err_update.Error()})
	// // 				return
	// // 			}
	// // 		}
		
    // //     }
	// // }
		return nil
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

func ProductFavorite(c *gin.Context) {
	slug := c.Param("slug")
	productModel, err := FindOneProduct(&ProductModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("protucts", errors.New("Invalid slug")))
		return
	}
	myUserModel := c.MustGet("my_user_model").(Users)
	err = productModel.favoriteBy(GetProductUsers(myUserModel))


	err_karmaProd:= Karma_redis("product", productModel.Slug, 10)

	if err_karmaProd != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err_karmaProd.Error()})
	return
	}

	err_karmaBrd:= Karma_redis("brands", productModel.Brand, 5)
	if err_karmaBrd != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err_karmaBrd.Error()})
	return
	}
	serializer := ProductSerializer{c, productModel}
	c.JSON(http.StatusOK, gin.H{"product": serializer.Response()})
}

func ProductUnfavorite(c *gin.Context) {
	slug := c.Param("slug")
	productModel, err := FindOneProduct(&ProductModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid slug")))
		return
	}
	myUserModel := c.MustGet("my_user_model").(Users)
	err = productModel.unFavoriteBy(GetProductUsers(myUserModel))

	err_karma:= Karma_redis("product", productModel.Slug, -10)
	if err_karma != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err_karma.Error()})
	return
	}	
	serializer := ProductSerializer{c, productModel}
	c.JSON(http.StatusOK, gin.H{"product": serializer.Response()})
}

func ProductList(c *gin.Context) {
	slug := c.Param("slug")

	if (slug=="list"){
		favorited := c.Query("favorited")
		limit := c.Query("limit")
		offset := c.Query("offset")
	
		productModels, modelCount, err := FindManyProducts(limit, offset, favorited)
		if err != nil {
			c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid param")))
			return
		}
		serializer := ProductsSerializer{c, productModels}
		c.JSON(http.StatusOK, gin.H{"products": serializer.Response(), "productsCount": modelCount})
	} else if (slug=="home"){

		client := common.NewClient()
		err_get, val := common.Get("brands", client)
		if err_get != nil {
	
			//AQUI LO BUSCAREMOS DE BASE DE DATOS/////////////////////////
			c.JSON(http.StatusBadRequest, gin.H{"error": err_get.Error()})	
			}
	
			objects := map[string]int{}

			json.Unmarshal([]byte(val), &objects)

			type object struct {
				Key   string
				Value int
			}
		
			var objectsort []object
			for k, v := range objects {
				objectsort = append(objectsort, object{k, v})
			}

		
			sort.Slice(objectsort, func(i, j int) bool {
				return objectsort[i].Value > objectsort[j].Value
			})

			keys := make([]string, 0)

			for k := range objectsort {
				keys= append(keys,objectsort[k].Key )
			}

			c.JSON(http.StatusOK, gin.H{"brands": keys})


	}else if (strings.Contains(slug, "brands")){
		brands:= strings.Split(slug, ",")
		brand:=brands[1]

		products, err := ProductsbyBrands(&ProductModel{Brand: brand})


		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
		}

		product := make([]ProductModel, 0)

		for i := range products {
			product= append(product,products[i])
		}


		err_karmaBrd:= Karma_redis("brands", brand, 5)
		if err_karmaBrd != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err_karmaBrd.Error()})
		return
		}
	

		c.JSON(http.StatusOK, gin.H{"product": product})
				
	}else{
		if slug == "feed" {
			ProductFeed(c)
			return
		}
		productModel, err := FindOneProduct(&ProductModel{Slug: slug})
		if err != nil {
			c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid slug")))
			return
		}
	
		err_karma:= Karma_redis("product", productModel.Slug, 5)
		if err_karma != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err_karma.Error()})
		return
		}
	
		serializer := ProductSerializer{c, productModel}
		c.JSON(http.StatusOK, gin.H{"product": serializer.Response()})
	}
	
}

func BrandsHome(c *gin.Context){
	client := common.NewClient()
	err_get, val := common.Get("brands", client)
	if err_get != nil {

		//AQUI LO BUSCAREMOS DE BASE DE DATOS/////////////////////////
		c.JSON(http.StatusBadRequest, gin.H{"error": err_get.Error()})



		}

		fmt.Println(val)
	}


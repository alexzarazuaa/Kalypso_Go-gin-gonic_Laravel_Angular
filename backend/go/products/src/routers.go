package products

import (
	"fmt"
	"strconv"
	"sort"
	"strings"
	"encoding/json"
	"errors"
	"net/http"
	"goProducts/common"
	"github.com/gin-gonic/gin"
)

func ProductsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/:slug", ProductList)
	router.POST("/:slug", UpKarmaProduct)

}


func ProductsRegister(router *gin.RouterGroup) {
	router.POST("/:slug/favorite", ProductFavorite)
	router.DELETE("/:slug/favorite", ProductUnfavorite)
}


func UpKarmaProduct(c *gin.Context) {
	data := c.Param("slug")
	s := strings.Split(data, ",")

	err_karma:= Karma_redis("products", s[0], 10)
	if err_karma != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err_karma.Error()})
	return
	}	

	err_karma = Karma_redis("brands", s[1], 10)
	if err_karma != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err_karma.Error()})
	return
	}	
	c.JSON(http.StatusOK, gin.H{"karma": "okey"})
}


//General function to karma in products and brands
func Karma_redis( types string, id string, karma int) error{
	client := common.NewClient()


	//obtain data(brands or products) from redis 
	err_get, val := common.Get(types, client)
	if err_get != nil {//If not exist data in redis we storege first brand or product

		objects := map[string]int{ id: karma }
		err_SetMarshal:= SetMarshal(objects,types)//Object -> Byte and storage in redis

		if (err_SetMarshal!= nil){//Any mistakes
			return err_SetMarshal

		}

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


	err_karma:= Karma_redis("products", productModel.Slug, 10)
	if err_karma != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err_karma.Error()})
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

	err_karma:= Karma_redis("products", productModel.Slug, -10)
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

	} else if (strings.Contains(slug, "home")){

		mode:= strings.Split(slug, ",")


		var vars [2]string
		vars[0] = "brands"
		vars[1] = "products"

		data := make(map[string]interface{})
		products := []map[string]interface{}{}
		value:= false
		
		if (mode[1]=="admin"){ value=true}
		client := common.NewClient()

		//bucleamos para que busque brands y product
		for v := range vars {

			err_get, val := common.Get(vars[v], client)

			if err_get != nil { ///SI no existen datos en redis los buscara en la base de datos
				if (vars[v] == "products"){//Si los datos que no existen son los de productos

					favorited := c.Query("favorited")
					limit := c.Query("limit")
					offset := c.Query("offset")
					productModels, _, err := FindManyProducts(limit, offset, favorited)
					if err != nil {
						c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid param")))
						return
					}

					str:=`{`
					for k := range productModels {
						if (	k ==(len(productModels)-1)) {
							str=str+ `"`+ productModels[k].Slug+`":`+ strconv.Itoa(productModels[k].Rating)
						}else{
							str=str+ `"`+ productModels[k].Slug+`":`+ strconv.Itoa(productModels[k].Rating)+`,`
						}
					}
			   		str=str+ `}`

					keys:= order_redis(str, value)

					products:= ProductsKarma(keys, products)
					data["products"]=products

				}else{  // Si los datos que no existen son los de marcas
					brands, err := GetBrands()
					if err != nil {
						c.JSON(http.StatusNotFound, common.NewError("products", errors.New("Invalid param")))
						return
					}

					str:=`{`
					for k := range brands {
						if (k ==(len(brands)-1)) {
							str=str+ `"`+ brands[k].Name+`":`+ strconv.Itoa(brands[k].Rating)
						}else{
							str=str+ `"`+ brands[k].Name+`":`+ strconv.Itoa(brands[k].Rating)+`,`
						}
			 		}
			  		str=str+ `}`

					keys:= order_redis(str, value)
					data["brands"]=keys
				}
			}else{  // Si existen datos en redis los buscara en redis

				keys:= order_redis(val, value)

					if (vars[v] == "products"){ // si los datos existentes en redis son de productos
						products:= ProductsKarma(keys, products)
						data["products"]=products
					}else{// si los datos existentes en redis son de marcas
						data["brands"]=keys
					}
			}
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
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

		 err, productModel:=detail(slug, true)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		serializer := ProductSerializer{c, productModel}
		c.JSON(http.StatusOK, gin.H{"product": serializer.Response()})
		
	}
}

func order_redis( val string, value bool )  []map[string]interface{} {
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


	data := []map[string]interface{}{}


	for k := range objectsort {
		object := map[string]interface{}{
			"key" : objectsort[k].Key, 
			"value" : objectsort[k].Value,
		}

		data = append(data, object)

		if (value==false && k==4 ) {break}
	}

	return data
}

func detail (slug string, key bool) (error,ProductModel) {
	
	productModel, err := FindOneProduct(&ProductModel{Slug: slug})
	if err != nil {
		return err , productModel
	}

	if (key==true){
		err_karma:= Karma_redis("products", productModel.Slug, 5)
		if err_karma != nil {
		return err_karma, productModel
		}
	}
	

	return nil, productModel
}

func ProductsKarma (keys, products []map[string]interface{}) []map[string]interface{} {
	for k := range keys {

					
					 
		err, productModel:=detail(fmt.Sprintf("%v", keys[k]["key"]), false)

		if err == nil {
			product := map[string]interface{}{
			   "key" : productModel,
			   "value" : keys[k]["value"],
		   }

		   products= append(products,product) 
	   }
   }
   return products
}

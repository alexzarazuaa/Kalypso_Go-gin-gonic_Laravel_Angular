<?php

namespace App\Http\Controllers\Api;


use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\RealWorld\Transformers\BuyProductTransformer;
use App\Model_buysProducts;
use App\Http\Requests\Api\CreateBuyProduct;


class buysProducts extends ApiController
{

    /**
     * buysProducts constructor.
     *
     * @param BuyProductTransformer $transformer
     */
    public function __construct(BuyProductTransformer $transformer)
    {
        $this->transformer = $transformer;

        // $this->middleware('auth.api')->except(['index', 'show']);
        // $this->middleware('auth.api:optional')->only(['index', 'show']);
    }


    /**
     * INDEX FOR GET ALL BUYSPRODUCTS
     */

    public function index()
    {
        $buyProduct = Model_buysProducts::all();
    
        return response() -> json($buyProduct);
    }

    public function store(CreateBuyProduct $request){
        
        $buyProduct = new Model_buysProducts();
        $buyProduct -> id_user = $request -> id_user;
        $buyProduct -> slug = $request -> slug;
        $buyProduct -> name = $request -> name;
        $buyProduct -> brand = $request -> brand;
        $buyProduct -> image = $request -> image;
        $buyProduct -> desc = $request -> desc;
        $buyProduct -> rating = $request -> rating;
        $buyProduct -> category = $request -> category;

     


  
        return response() -> json($buyProduct);
         //return $this->respondWithTransformer($buyProduct);
    }



    /**
     * RETURN ONE BUY PRODUCT
     */
    public function show(Model_buysProducts $buyProduct)
    {
        return $this->respondWithTransformer($buyProduct);
    }// end_showSong
    



    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $id)
    {
        $model_buysProducts = Model_buysProducts::find($id);

        if (!$model_buysProducts) return response() -> json('Not Found');

     
        $model_buysProducts -> id_user = $request -> id_user;
        $model_buysProducts -> slug = $request -> slug;
        $model_buysProducts -> name = $request -> name;
        $model_buysProducts -> brand = $request -> brand;
        $model_buysProducts -> image = $request -> image;
        $model_buysProducts -> desc = $request -> desc;
        $model_buysProducts -> rating = $request -> rating;
        $model_buysProducts -> category = $request -> category;
    

        $model_buysProducts -> save();

        return response() -> json($model_buysProducts);
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        $model_buysProducts = Model_buysProducts::find($id);
        
        if(!$model_buysProducts) return response() -> json('Not Found');

        $model_buysProducts -> delete();

        return response() -> json($model_buysProducts);
    }
}
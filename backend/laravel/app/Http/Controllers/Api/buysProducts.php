<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\Model_buysProducts;


class buysProducts extends ApiController
{

    public function store(Request $request){
        
        $model_buysProducts = new Model_buysProducts();
        $model_buysProducts -> id_user = $request -> id_user;
        $model_buysProducts -> slug = $request -> slug;
        $model_buysProducts -> name = $request -> name;
        $model_buysProducts -> brand = $request -> brand;
        $model_buysProducts -> image = $request -> image;
        $model_buysProducts -> desc = $request -> desc;
        $model_buysProducts -> rating = $request -> rating;
        $model_buysProducts -> category = $request -> category;


        //print_r($model_buysProducts);

        $model_buysProducts -> save();

        return response() -> json($model_buysProducts);
        // return $this->respondWithPagination($model_buysProducts);
    }


    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show()
    {
        $model_buysProducts = Model_buysProducts::all();

        return response() -> json($model_buysProducts);
    }

    public function index()
    {
        $model_buysProducts = Model_buysProducts::all();

        return response() -> json($model_buysProducts);
    }


    /**
     * RETURN ONE BUY PRODUCT
     */
    public function showBuyProduct($id) {

        $model_buysProducts = Model_buysProducts::find($id);
        print_r($id);
        return response() -> json($model_buysProducts);

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
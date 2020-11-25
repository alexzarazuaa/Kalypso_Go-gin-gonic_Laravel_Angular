<?php

namespace App\Http\Controllers\Api;

use App\Http\Requests\Api\createBuyProduct;
use App\Http\Requests\Api\UpdateProduct;
use App\Http\Requests\Api\DeleteBuyProduct;
use App\RealWorld\Transformers\BuyProductTransformer;
use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\Model_buysProducts;


class buysProducts extends ApiController
{


    /**
     * ArticleController constructor.
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
     * GET ALL BUY PRODUCTS
     */
    public function index()
    {
        $model_buysProducts = Model_buysProducts::all();

        return response() -> json($model_buysProducts);
    }

    
    public function store(CreateBuyProduct $request){
        
        $user = auth()->user();

        $buyProduct = $user->model_buysProducts()->create([
            'slug' =>$request->input('buyProduct.slug'),
            'name' => $request->input('buyProduct.name'),
            'brand' => $request->input('buyProduct.brand'),
            'image' =>$request->input('buyProduct.image'),
            'desc' =>$request->input('buyProduct.desc'),
            'rating' => $request->input('buyProduct.rating'),
            'category' => $request->input('buyProduct.category')

        ]);


        return $this->respondWithTransformer($buyProduct);
    }


    /**
     * RETURN ONE BUY PRODUCT
     */


    public function show($id)
    {        $model_buysProducts = Model_buysProducts::find($id);
        print_r($id);
        return response() -> json($model_buysProducts);
    }


    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(UpdateProduct $request,Model_buysProducts $buyProduct)
    {
        if ($request->has('buyProduct')) {
            $buyProduct->update($request->get('buyProduct'));
        }

        return $this->respondWithTransformer($buyProduct);
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy(DeleteBuyProduct $request, Model_buysProducts $buyProduct)
    {
        $buyProduct->delete();

        return $this->respondSuccess();
    }
}

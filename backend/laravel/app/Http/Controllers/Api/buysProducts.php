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

    /**
     * RETURN ONE BUY PRODUCT
     */
    public function show($id) {

        $model_buysProducts = Model_buysProducts::find($id);
        return response() -> json($model_buysProducts);

    }// end_product

 
}
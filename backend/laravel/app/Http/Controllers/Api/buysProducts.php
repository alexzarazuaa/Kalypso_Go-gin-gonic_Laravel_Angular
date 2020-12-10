<?php

namespace App\Http\Controllers\Api;


use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\RealWorld\Transformers\BuyProductTransformer;
use App\Model_buys_products;
use App\Http\Requests\Api\CreateBuyProduct;


class buys_products extends ApiController
{

    /**
     * buys_products constructor.
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
     * INDEX FOR GET ALL BUYS_pRODUCTS
     */

    public function index()
    {
        $buyProduct = Model_buys_products::all();
    
        return response() -> json($buyProduct);
    }

    /**
     * RETURN ONE BUY PRODUCT
     */
    public function show($id) {

        $model_buys_products = Model_buys_products::find($id);
        return response() -> json($model_buys_products);

    }// end_product

 
}
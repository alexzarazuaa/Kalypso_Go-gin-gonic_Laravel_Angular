<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\RealWorld\Transformers\BuyProductTransformer;
use App\Http\Requests\Api\CreateProduct;
use App\Http\Requests\Api\DeleteProduct;
use App\Product;



class Products extends ApiController
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        $Products = Product::all();
    
        return response() -> json($Products);
    }


    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    // public function create()
    // {
    //     //
    // }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
    
        $product = new Product();
        // $name = $request -> product['name'];
        $product -> slug = str_slug($request -> product['name']);  
        $product -> name = $request -> product['name'];;
        $product -> brand = $request -> product['brand'];
        $product -> description = $request -> product['description'];
        $product -> rating = $request -> product['rating'];
        $product -> category = $request -> product['category'];

    
        $product -> save();

        return response() -> json($product);
        
    }


    /**
     * Get the product given by its slug.
     *
     * @param Product $product
     * @return \Illuminate\Http\JsonResponse
     */
    public function show(Product $product)
    {
        return response() -> json($product);
    }

    // public function show(Article $article)
    // {
    //     return $this->respondWithTransformer($article);
    // }


    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $slug
     * @return \Illuminate\Http\Response
     */
    public function edit($slug)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $slug
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $slug)
    {
        $product = Product::find($slug);
        if(!$product) return response() -> json('Product Not Found');

        $product -> slug = str_slug($request -> product['name']);  
        $product -> name = $request -> product['name'];;
        $product -> brand = $request -> product['brand'];
        $product -> description = $request -> product['description'];
        $product -> rating = $request -> product['rating'];
        $product -> category = $request -> product['category'];

    
        $product -> save();

        return response() -> json($product);
    }
    /**
     * Delete the product given by its slug.
     *
     * @param DeleteProduct $request
     * @param Product $product
     * @return \Illuminate\Http\JsonResponse
     */

    public function destroy($slug) {
        $product = Product::where('slug', $slug)->delete();
        return $this -> respondSuccess();
    }// end_delete
}

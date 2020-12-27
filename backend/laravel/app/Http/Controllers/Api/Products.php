<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\RealWorld\Transformers\BuyProductTransformer;
use App\Product;



class Products extends Controller
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
    
        // echo($request);
        // echo("-----------------");
        // print_r($request->product['name']);

        $product = new Product();
        $product -> name = $request -> product['name'];
        $product -> brand = $request -> product['brand'];
        // $product -> img = $request -> product['img'];
        $product -> description = $request -> product['description'];
        $product -> rating = $request -> product['rating'];
        $product -> category = $request -> product['category'];

    
        $product -> save();

        return response() -> json($product);
         //return $this->respondWithTransformer($buyProduct);
        
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        // echo("holagola");
        $product = Product::find($id);
        return response() -> json($product);
    }

    // public function show(Article $article)
    // {
    //     return $this->respondWithTransformer($article);
    // }


    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $id)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        $Products = Product::find($id);
        
        if(!$Products) return response() -> json('Product not Found');

        $Products -> delete();

        return response() -> json($Products);
    }
}

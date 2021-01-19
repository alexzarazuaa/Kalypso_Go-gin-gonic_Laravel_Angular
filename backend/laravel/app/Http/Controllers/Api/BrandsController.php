<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\Brand;
use App\Product;

use Illuminate\Support\Facades\Redis;


class Brands extends ApiController
{

    public function index()
    {

        $types=['brands', 'products'];

        foreach($types as $key_t) {
            
            $redis = json_decode(Redis::get($key_t), true);

            foreach($redis as $key => $value) {

                if ( $key_t == "brands"){
                    $object = Brand::where('name',$key)->first();
                }else{
                    $object = Product::where('slug',$key)->first();

                }
                    if($object){
                        $object -> rating = $value;
                        $object -> save();
                    }
            }

        }

    return $this -> respondSuccess("okey");

    }

    public function store(Request $request)
    {
            //
    }



    public function edit($slug)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     */
    public function update(Request $request)
    {
   
       
      error_log('HEY THERE ENTRA update');
        return $this -> respondSuccess("OKEY");

    }
    /**
     * Delete Redis
     */

    public function destroy() {

        error_log('HEY THERE  ENTRA DESTROY');
        Redis::del('brands');
        Redis::del('products');

        return $this -> respondSuccess("okey");
    }// end_delete
    
}

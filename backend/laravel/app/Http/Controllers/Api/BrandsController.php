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
    
}

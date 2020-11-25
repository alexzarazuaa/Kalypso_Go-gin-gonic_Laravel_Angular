<?php

namespace App\RealWorld\Transformers;

class BuyProductTransformer extends Transformer
{
    protected $resourceName = 'buyProduct';

    public function transform($data)
    {
        return [
            'id_user'              => $data['id_user'],
            'slug'                 => $data ['slug'],
            'name'                 => $data['name'],
            'brand'                =>$data['brand'],
            'image'                =>$data['image'],
            'desc'                 =>$data['desc'],
            'rating'               => $data['rating'],
            'category'             => $data['category']

         

        ];
    }
}

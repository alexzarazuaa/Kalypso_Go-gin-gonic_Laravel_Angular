<?php

namespace App\RealWorld\Transformers;

class ProfileTransformer extends Transformer
{
    protected $resourceName = 'profile';

    public function transform($data)
    {
        return [
            'username'  => $data['username'],
            'image'     => $data['image'],
            'following' => $data['following'],
        ];
    }
}
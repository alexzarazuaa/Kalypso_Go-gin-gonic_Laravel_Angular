<?php

namespace App\RealWorld\Transformers;

class UserTransformer extends Transformer
{
    protected $resourceName = 'user';

    public function transform($data)
    {
        return [
            'email'     => $data['email'],
            'token'     => $data['token'],
            'username'  => $data['username'],
            'image'     => $data['image'],
            'type'      => $data['type']
        ];
    }
}
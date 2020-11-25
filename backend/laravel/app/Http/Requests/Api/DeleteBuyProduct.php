<?php

namespace App\Http\Requests\Api;

class DeleteBuyProduct extends ApiRequest
{
    /**
     * Determine if the user is authorized to make this request.
     *
     * @return bool
     */
    public function authorize()
    {
        $buyProduct = $this->route('buyProduct');

        // return $buyProduct->user_id == auth()->id();
    }
}

<?php

namespace App\Http\Requests\Api;

class CreateBuyProduct extends ApiRequest
{
    /**
     * Get data to be validated from the request.
     *
     * @return array
     */
    protected function validationData()
    {
        return $this->get('buyProduct') ?: [];
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'id_user' => 'required|integer|max:255',
            'slug' => 'required|integer|max:255',
            'name' => 'required|string|max:255',
            'brand' => 'required|string|max:255',
            'image' => 'required|string|max:255',
            'desc' => 'required|string|max:255',
            'rating' => 'required|integer|max:255',
            'category' => 'required|integer|max:255'
        ];
          
    }
}

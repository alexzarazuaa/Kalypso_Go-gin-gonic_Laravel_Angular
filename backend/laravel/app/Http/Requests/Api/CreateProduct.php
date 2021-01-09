<?php

namespace App\Http\Requests\Api;

class CreateProduct extends ApiRequest
{
    /**
     * Get data to be validated from the request.
     *
     * @return array
     */
    protected function validationData()
    {
        return $this->get('product') ?: [];
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'slug' => 'required|integer|max:255',
            'name' => 'required|string|max:255',
            'brand' => 'required|string|max:255',
            'description' => 'required|string|max:255',
            'rating' => 'required|integer|max:255',
            'category' => 'required|integer|max:255'
        ];
    }
}


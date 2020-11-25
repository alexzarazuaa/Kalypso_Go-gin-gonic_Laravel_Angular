<?php

namespace App\Http\Requests\Api;

class UpdateProduct extends ApiRequest
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
     * Determine if the user is authorized to make this request.
     *
     * @return bool
     */
    public function authorize()
    {
        $buyProduct = $this->route('buyProduct');

        // return $buyProduct->user_id == auth()->id();
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [

            'name' => 'required|string|max:255',
            'brand' => 'required|string',
            'rating' => 'required|integer',
            'category' => 'required|string'
        ];
    }
}

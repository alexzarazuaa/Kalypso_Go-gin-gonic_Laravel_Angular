<?php

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::group(['namespace' => 'Api'], function () {


        
    // Route::post('products', 'buys_products@create');
    // Route::get('products', 'buys_products@show');
    // Route::get('products/{id}', 'buys_products@showBuyProduct');
    // Route::put('products/{id}', 'buys_products@update');
    // Route::delete('products/{id}', 'buys_products@delete');

     Route::resource('buys_products', 'buys_products');
     Route::resource('products', 'Products');

    Route::post('users/login', 'AuthController@login');
    Route::post('users', 'AuthController@register');

    Route::get('user', 'UserController@index');
    Route::match(['put', 'patch'], 'user', 'UserController@update');

    Route::get('profiles/{user}', 'ProfileController@show');
    Route::post('profiles/{user}/follow', 'ProfileController@follow');
    Route::delete('profiles/{user}/follow', 'ProfileController@unFollow');

    Route::get('articles/feed', 'FeedController@index');
    Route::post('articles/{article}/favorite', 'FavoriteController@add');
    Route::delete('articles/{article}/favorite', 'FavoriteController@remove');

    Route::resource('articles', 'ArticleController', [
        'except' => [
            'create', 'edit'
        ]
    ]);

    Route::resource('articles/{article}/comments', 'CommentController', [
        'only' => [
            'index', 'store', 'destroy'
        ]
    ]);



});
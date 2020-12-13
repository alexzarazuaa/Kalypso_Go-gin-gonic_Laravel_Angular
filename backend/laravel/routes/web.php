<?php

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/


use Illuminate\Support\Facades\Redis;



Route::get('/', function () {
    return view('welcome');
});

// Route::get('/order', 'OrderController@index');
// Auth::routes();
// Route::get('/home', 'HomeController@index')->name('home');

///////////////////////////num_visits
Route::get('/redis', function () {
    $visits = Redis::Incr('visits');
    return "Num visits: " . $visits;
});

///////////////////////////downloads video_id
Route::get('/videos/{id}', function($id){
    $downloads = Redis::get('videos.{$id}.downloads');
    return view('welcome')->withDownloads($downloads);
});
Route::get('/videos/{id}/downloads', function($id){
    Redis::Incr('videos.{$id}.downloads');
    return back();
});

///////////////////////////posts_cache
Route::get('/posts', function (){
 
    $redis = Redis::connection();
    try {
        $posts = \App\Post::all();
        Redis::set('posts.all', $posts);

        if(Redis::exists('posts.all')){
            return json_decode(Redis::get('posts.all'));
        }
    } catch (\Predis\Connection\ConnectionException $e) {
        $posts = \App\Post::all();
        return json_decode($posts);
    }

    // $redis = Redis::connection();
    // try {
    //     $redis->ping();
    //     if(Redis::exists('posts.all')){
    //         return json_decode(Redis::get('posts.all'));
    //     }

    //     $posts = \App\Post::all();
    //     Redis::set('posts.all', $posts);

    //     return json_decode($posts);
    // } catch (\Predis\Connection\ConnectionException $e) {
    //     $posts = \App\Post::all();
    //     return json_decode($posts);
    // }
});

Route::get('/posts_cache', function (){
    return \Cache::remember('posts.all', 60 * 60, function () {
        return json_decode(Redis::get('posts.all'));
    });
});

///////////////////////////posts_favorite
Route::get('favorite-post', function () {
    Redis::hincrby('posts.1.stats', 'favorites', 1);
    Redis::hincrby('posts.1.stats', 'watchLaters', 1);
    Redis::hincrby('posts.1.stats', 'completions', 1);
    return redirect('/posts/1/stats');
});
Route::get('posts/{id}/stats', function ($id) {
    return Redis::hgetall("posts.{$id}.stats");
});

///////////////////////////posts_trending
Route::get('posts/trending', function () {
    $trending = Redis::zrevrange('trending_posts', 0, 2);
    $trending = \App\Post::hydrate(
        array_map('json_decode', $trending)
    );
    return $trending;
});
Route::get('posts/{post}', function (\App\Post $post) {
    Redis::zincrby('trending_posts', 1, $post);
    return $post;
});
<?php

/*
|--------------------------------------------------------------------------
| Model Factories
|--------------------------------------------------------------------------
|
| Here you may define all of your model factories. Model factories give
| you a convenient way to create models for testing and seeding your
| database. Just tell the factory how a default model should look.
|
*/

$factory->define(App\User::class, function (\Faker\Generator $faker) {

    return [
        'username' => str_replace('.', '', $faker->unique()->userName),
        'email' => $faker->unique()->safeEmail,
        'password' => 'secret',
        'bio' => $faker->sentence,
        'image' => 'https://cdn.worldvectorlogo.com/logos/laravel.svg',
        'type' => 'type',
    ];
});

$factory->define(App\Article::class, function (\Faker\Generator $faker) {

    static $reduce = 9;

    return [
        'title' => $faker->sentence,
        'description' => $faker->sentence(10),
        'body' => $faker->paragraphs($faker->numberBetween(1, 3), true),
        'created_at' => \Carbon\Carbon::now()->subSeconds($reduce--),
    ];
});
$factory->define(App\Model_buys_products::class , function(\Faker\Generator $faker){

    static $reduce = 999;
    return[
        'id_user' => $faker->randomDigit($faker->numberBetween(1, 99999), true),
        'slug' =>$faker->randomDigit($faker->numberBetween(1, 99999), true),
        'name' => $faker->firstName,
        'brand'  => $faker->company,
        'image' => $faker->imageUrl($width = 640, $height = 480),
        'desc' => $faker->sentence($nbWords = 6, $variableNbWords = true),
        'rating' => $faker->randomDigit($faker->numberBetween(1, 99999), true),
        'category' => $faker->word
        
    ];
});

//
$factory->define(App\Product::class , function(\Faker\Generator $faker){

    static $reduce = 999;
    return[
        'id_user' => $faker->randomDigit($faker->numberBetween(1, 99999), true),
        'slug' =>$faker->randomDigit($faker->numberBetween(1, 99999), true),
        'name' => $faker->firstName,
        'brand'  => $faker->company,
        'image' => $faker->imageUrl($width = 640, $height = 480),
        'desc' => $faker->sentence($nbWords = 6, $variableNbWords = true),
        'rating' => $faker->randomDigit($faker->numberBetween(1, 99999), true),
        'category' => $faker->word
        
    ];
});

$factory->define(App\Comment::class, function (\Faker\Generator $faker) {

    static $users;
    static $reduce = 999;

    $users = $users ?: \App\User::all();

    return [
        'body' => $faker->paragraph($faker->numberBetween(1, 5)),
        'user_id' => $users->random()->id,
        'created_at' => \Carbon\Carbon::now()->subSeconds($reduce--),
    ];
});

$factory->define(App\Tag::class, function (\Faker\Generator $faker) {

    return [
        'name' => $faker->unique()->word,
    ];
});

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
        'image' => 'https://cdn.worldvectorlogo.com/logos/laravel.svg',
        'password' => 'secret',
        'karma' => $faker->randomDigit($faker->numberBetween(1, 99999), true),
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
$factory->define(App\Model_buysProducts::class , function(\Faker\Generator $faker){

    static $reduce = 999;
    return[
        'slug' =>$faker->randomDigit($faker->numberBetween(10, 99999), true),
        'name' => $faker->firstName,
        'brand'  => $faker->company,
        'img' => $faker->imageUrl($width = 640, $height = 480),
        'description' => $faker->sentence($nbWords = 6, $variableNbWords = true),
        'rating' => $faker->randomDigit($faker->numberBetween(1, 99999), true),
        'category' => $faker->word
        
    ];
});

//
$factory->define(App\Product::class , function(\Faker\Generator $faker){

    static $reduce = 999;
    return[
        'slug' =>$faker->randomDigit($faker->numberBetween(1, 99999), true),
        'name' => $faker->firstName,
        'brand'  => $faker->company,
        'img' => $faker->imageUrl($width = 640, $height = 480),
        'description' => $faker->sentence($nbWords = 6, $variableNbWords = true),
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

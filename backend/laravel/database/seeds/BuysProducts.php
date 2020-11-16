<?php

use Illuminate\Database\Seeder;
use Faker\Factory as Faker;

class BuysProducts extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $faker = Faker::create();
        for ($i=0; $i < 5; $i++) {
            \DB::table('BuysProducts')->insert(array(
                   'id_user' => $faker->randomDigit,
                   'name' => $faker->firstName,
                   'brand'  => $faker->company,
                   'rating' => $faker->randomDigit,
                   'category' => $faker->word
            ));
        }

    }
}

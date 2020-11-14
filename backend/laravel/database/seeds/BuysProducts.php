<?php

use Illuminate\Database\Seeder;

class BuysProducts extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        DB::table('BuysProducts')->insert([
            'name' => str_random(3),
            'brand' => str_random(3),
            'rating' => str_random(3),
            'category' => str_random(3)
        ]);

          DB::table('BuysProducts')->insert([
            'name' => str_random(3),
            'brand' => str_random(3),
            'rating' => str_random(3),
            'category' => str_random(3)
        ]);

    }
}

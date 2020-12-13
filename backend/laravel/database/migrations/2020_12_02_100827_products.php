<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class Products extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        if (!Schema::hasTable('product_models'))
        {
        Schema::create('product_models', function (Blueprint $table) {
            $table->increments('id');
            $table->integer('slug');
            $table->string('name');
            $table->string('brand');
            $table->string('img');
            $table->string('description');
            $table->integer('rating');
            $table->string('category');
            $table->timestamps();
        });
    }
}

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    // public function down()
    // {
    //     Schema::dropIfExists('Products');
    // }
}

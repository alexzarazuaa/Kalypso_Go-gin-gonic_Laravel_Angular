<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CrearTablaBuysProducts extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('BuysProducts', function (Blueprint $table) {
            $table->increments('id');
            $table->integer('id_user');
            $table->integer('slug');
            $table->string('name');
            $table->string('brand');
            $table->string('image');
            $table->string('desc');
            $table->integer('rating');
            $table->string('category');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('BuysProducts');
    }
}

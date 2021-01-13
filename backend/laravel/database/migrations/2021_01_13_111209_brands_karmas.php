<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class BrandsKarmas extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        if (!Schema::hasTable('brands_karmas'))
        {
        Schema::create('brands_karmas', function (Blueprint $table) {
            $table->increments('id');
            $table->string('name');
            $table->string('rating');
            $table->timestamps();
        });
    }
    }

    // /**
    //  * Reverse the migrations.
    //  *
    //  * @return void
    //  */
    // public function down()
    // {
    //     Schema::dropIfExists('brands_karmas');
    // }
}

<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class UserModel extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
            if (!Schema::hasTable('users'))
        {
                Schema::create('users', function (Blueprint $table) {
                    $table->increments('id');
                    $table->string('username');
                    $table->string('email');
                    $table->string('image', 2048)->nullable();
                    $table->string('karma');
                    $table->string('password');
                    $table->string('type');
                    $table->rememberToken();

                    $table->unique(['username', 'email']);
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
    //     Schema::dropIfExists('users');
    // }

}
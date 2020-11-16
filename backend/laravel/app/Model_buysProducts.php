<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Model_buysProducts extends Model
{
    protected $table = 'BuysProducts';
    protected $fillable = ['id_user','name','brand','rating','category'];
}

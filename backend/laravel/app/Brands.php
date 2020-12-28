<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Article extends Model
{
    protected $table ='brands';
    protected $fillable=['name', 'karma'];
}

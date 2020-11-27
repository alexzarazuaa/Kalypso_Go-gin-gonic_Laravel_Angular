<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Model_buysProducts extends Model
{
    protected $table = 'BuysProducts';
    protected $fillable = ['id_user','slug','name','brand','image','desc','rating','category'];


    /**
     * Get the key name for route model binding.
     *
     * @return string
     */
    public function getRouteKeyName()
    {
        return 'slug';
    }

        /**
     * Get the attribute name to slugify.
     *
     * @return string
     */
    public function getSlugSourceColumn()
    {
        return 'slug';
    }

}



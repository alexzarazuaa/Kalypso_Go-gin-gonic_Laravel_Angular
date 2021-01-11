import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpParams } from '@angular/common/http';
import { ApiService } from './api.service';
import { Products } from '../models';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ProductsService {

  constructor(private apiService: ApiService) { }

  getAll(): Observable<Products[]> {
    const params = {};
    return this.apiService.get('/products/');
  }// end_query


  //GET PRODUCTS BY BACKEND GO
  
  getAll_goProd(): Observable<Products[]> {
    const params = {};
    return this.apiService.get_goProd('/list')
    .pipe(map(data => {
      console.log('data in service',data.products);
      return data.products;
    }));
  }// end_query

  getOneGO(slug): Observable<Products> {
    return this.apiService.get_goProd('/' + slug)
      .pipe(map(data => {
        console.log('slug detail product GO',data);
        return data;
      }));
  }// end_get

  

  getOne(id): Observable<Products> {
    return this.apiService.get('/products/' + id)
      .pipe(map(data => {
        console.log('data in service',data);
        return data;
      }));
  }// end_get




  favorite(slug): Observable<Products> {
    return this.apiService.post_goProd('/' + slug + '/favorite');

    //http://localhost:3000/api/products/levis-crop-jeans/favorite
  }

  unfavorite(slug): Observable<Products> {
    return this.apiService.delete_goProd('/' + slug + '/favorite');
  }

  // SAVE PRODUCT FOR EDITOR

  save(product): Observable<Products> {
    // If we're updating an existing product
    if (product.slug) {
      return this.apiService.put('/products/' + product.slug, {product: product})
        .pipe(map(data => data.product));

    // Otherwise, create a new product
    } else {
      return this.apiService.post('/products/', {product: product})
        .pipe(map(data => data.product));
    }

  }

  destroy(slug) {
    return this.apiService.delete('/products/' + slug);
  }



 


}


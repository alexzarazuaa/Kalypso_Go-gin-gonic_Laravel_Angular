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
  

  getOne(id): Observable<Products> {
    return this.apiService.get('/products/' + id)
      .pipe(map(data => {
        console.log('data in service',data);
        return data;
      }));
  }// end_get


  destroy(slug) {
    return this.apiService.delete('/product/' + slug);
  }

}




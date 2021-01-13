import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpParams } from '@angular/common/http';
import { ApiService } from './api.service';
import { BuyProduct } from '../models';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class BuysProductsService {

  constructor(private apiService: ApiService) { }

  getAll(): Observable<BuyProduct[]> {
    const params = {};
    return this.apiService.get('/buysproducts/');
  }// end_query
  

  getOne(id): Observable<BuyProduct> {
    return this.apiService.get('/buysproducts/' + id)
      .pipe(map(data => {
 
        return data;
      }));
  }// end_get


  destroy(slug) {
    return this.apiService.delete('/product/' + slug);
  }
  
  insert(slug){
    return this.apiService.post_buys('/' + slug)
    .pipe(map(data => {

      return data;
    }));
  }

  

}




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
    return this.apiService.get('/products');
  }// end_query


  get(id): Observable<BuyProduct> {
    return this.apiService.get('/products/' + id)
    
      .pipe(map(data => {
        console.log(data);
        return data.product;
      }));
  }// end_get


  destroy(slug) {
    return this.apiService.delete('/product/' + slug);
  }

}




import { Injectable } from '@angular/core';
import { HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

import { ApiService } from './api.service';
import { BuyProduct } from '../models';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class BuysProducts {

  constructor(private apiService: ApiService) {
  }

  query(): Observable<BuyProduct[]> {
    const params = {};
    return this.apiService.get('/products');
  }// end_query
  

  get(slug): Observable<BuyProduct> {
    return this.apiService.get('/products/' + slug)
      .pipe(map(data => {
        console.log(data);
        return data.buyproduct;
      }));
  }// end_get



}
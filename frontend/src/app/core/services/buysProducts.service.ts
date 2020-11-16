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

  query(): Observable<{BuyProduct: BuyProduct[]}> {
    const params = {};

    return this.apiService.get('/products', 'laravel');
  }// end_query

  get(slug): Observable<BuyProduct> {
    return this.apiService.get('/product/' + slug)
      .pipe(map(data => data.BuyProduct));
  }// end_get
}
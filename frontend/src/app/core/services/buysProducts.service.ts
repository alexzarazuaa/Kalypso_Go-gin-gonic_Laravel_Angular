import { Injectable } from '@angular/core';

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

  query(): Observable<{ buyProducts: BuyProduct[] }> {

    return this.apiService.get('/products');
  }// end_query

  get(slug): Observable<BuyProduct> {
    return this.apiService.get('/product/' + slug)
      .pipe(map(data => {
        console.log(data);
        return data;
      }));
  }// end_get


}
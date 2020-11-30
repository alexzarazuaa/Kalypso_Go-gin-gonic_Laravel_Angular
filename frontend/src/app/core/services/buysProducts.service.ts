import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { ApiService } from './api.service';
import { BuyProduct  } from '../models';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class BuysProductsService {

  constructor(private apiService: ApiService) {}

  getAll(): Observable<BuyProduct[]> {

    return this.apiService.get('/products');
  }

  
  get(slug): Observable<BuyProduct> {
    return this.apiService.get('/products/' + slug)
      .pipe(map(data => data.product));
  }// end_get


  destroy(slug) {
    return this.apiService.delete('/product/' + slug);
  }

}




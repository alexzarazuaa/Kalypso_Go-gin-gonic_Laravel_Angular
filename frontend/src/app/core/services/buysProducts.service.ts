import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpParams } from '@angular/common/http';
import { ApiService } from './api.service';
import { BuyProduct, BuyProductListConfig } from '../models';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class BuysProductsService {

  constructor(private apiService: ApiService) { }

  getAll(config: BuyProductListConfig): Observable<{ products: BuyProduct[] }> {

    const params = {};
    return this.apiService
      .get(
        '/products/' + ((config.type === 'feed') ? 'feed' : ''),
        new HttpParams({ fromObject: params })
      );
  }



  get(slug): Observable<BuyProduct> {
    return this.apiService.get('/products/' + slug)
      .pipe(map(data => data.product));
  }// end_get


  destroy(slug) {
    return this.apiService.delete('/product/' + slug);
  }

}




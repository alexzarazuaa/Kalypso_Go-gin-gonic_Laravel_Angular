import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { ApiService } from './api.service';
import { PanelGo  } from '../models';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class PanelService {

  constructor(private apiService: ApiService) {}


  getAll(): Observable<PanelGo[]> {
    const params = {};
    return this.apiService.get_Go('/products/')
    .pipe(map(data => {
            console.log('dataGo in service',data);
            return data;
          }));
  }// end_query


  // getOne(id): Observable<BuyProduct> {
  //   return this.apiService.get('/products/' + id)
  //     .pipe(map(data => {
  //       console.log('data in service',data);
  //       return data.product;
  //     }));
  // }// end_get


  // destroy(slug) {
  //   return this.apiService.delete('/product/' + slug);
  // }


}




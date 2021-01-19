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
    return this.apiService.get_Go('/products/')
    .pipe(map(data => {
            return data;
          }));
  }// end_query




}




import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
// import { HttpClientModule } from '@angular/common/http';
import { Products } from '../models';
import { ApiService } from './api.service';
import { map } from 'rxjs/operators';

@Injectable({
    providedIn: 'root'
  })
export class BrandsService {
    constructor(
        private apiService: ApiService
    ) { }

    getBrands(mode): Observable<[string]> {
        return this.apiService.get_goProd('/home'+ mode)
            .pipe(map(data => {
                console.log(data)
                return data;
            }));
    }

    filterBrands(brand): Observable<[string]> {
        return this.apiService.get_goProd('/' + brand)
            .pipe(map(data => {
                console.log(data)
                return data;
            }));
    }

}

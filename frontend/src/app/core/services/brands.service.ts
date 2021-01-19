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
              
                return data;
            }));
    }

    filterBrands(brand): Observable<[string]> {
        return this.apiService.get_goProd('/' + brand)
            .pipe(map(data => {
              
                return data;
            }));
    }

    insertRedisDb(): Observable<[string]> {
        return this.apiService.get('/brands')
            .pipe(map(data => {
              
                return data;
            }));
    }

    DelRatingRedis(): Observable<[string]> {
        return this.apiService.delete('/brands')
            .pipe(map(data => {
              
                return data;
            }));
        }

    DelRatingDB(): Observable<[string]> {
            return this.apiService.put('/brands')
                .pipe(map(data => {
                  
                    return data;
                }));
            }


    DelRedisService() : Observable<[string]>  {
        console.log('HEY ENTRE SERVICE ++++ ------')
        return this.apiService.delete('/brands/deleteRedis')
        .pipe(map(data => {
            console.log(data)
            return data;
        }));
      }
    
      UpRedisService() : Observable<[string]> {
        console.log('HEY ENTRE SERVICE +++ ------')
        return this.apiService.put('/brands/updateRedis')
        .pipe(map(data => {
            console.log(data)
            return data;
        }));
      }


}

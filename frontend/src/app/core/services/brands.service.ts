import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
// import { HttpClientModule } from '@angular/common/http';
import { ApiService } from './api.service';
import { map } from 'rxjs/operators';

@Injectable()
export class BrandsService {
    constructor(
        private apiService: ApiService
    ) { }

    getBrands(): Observable<[string]> {
        return this.apiService.get_goProd('/home')
            .pipe(map(data => {
                console.log(data)
                return data;
            }));
    }

}

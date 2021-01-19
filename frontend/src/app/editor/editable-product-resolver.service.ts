import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, Router, RouterStateSnapshot } from '@angular/router';
import { Observable } from 'rxjs';

import { Products, ProductsService } from '../core';
import { catchError, map } from 'rxjs/operators';

@Injectable()
export class EditableProduct implements Resolve<Products> {
  constructor(
    private productsService: ProductsService,
    private router: Router,
  ) { }

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> {

    return this.productsService.getOne(route.params['slug'])
      .pipe(
        map(
          product => {
            return product
          }
        ),
        catchError((err) => this.router.navigateByUrl('/'))
      );
  }
}


import { Injectable, } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, Router, RouterStateSnapshot } from '@angular/router';
import { Observable } from 'rxjs';

import { Products, ProductsService, UserService } from '../core';
import { catchError } from 'rxjs/operators';

@Injectable()
export class ProductsResolver implements Resolve<Products> {
  constructor(
    private productsService: ProductsService,
    private router: Router,
    private userService: UserService
  ) {}

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> {
    return this.productsService.getOne(route.params['id'])
      .pipe(catchError((err) => {
        console.log('sadjfg');
        return this.router.navigateByUrl('/');
      }));
  }
}
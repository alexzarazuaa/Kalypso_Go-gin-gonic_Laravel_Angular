
import { Injectable, } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, Router, RouterStateSnapshot } from '@angular/router';
import { Observable } from 'rxjs';

import { BuyProduct, BuysProductsService, UserService } from '../core';
import { catchError } from 'rxjs/operators';

@Injectable()
export class ProductsResolver implements Resolve<BuyProduct> {
  constructor(
    private buysProductsService: BuysProductsService,
    private router: Router,
    private userService: UserService
  ) {}

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> {
    return this.buysProductsService.getOne(route.params['id'])
      .pipe(catchError((err) => {
        console.log('sadjfg');
        return this.router.navigateByUrl('/');
      }));
  }
}
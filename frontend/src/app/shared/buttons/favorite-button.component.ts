import { Component, EventEmitter, Input, Output } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';

import { Products, ProductsService, UserService } from '../../core';
import { of } from 'rxjs';
import { concatMap ,  tap } from 'rxjs/operators';

@Component({
  selector: 'app-favorite-button',
  templateUrl: './favorite-button.component.html'
})
export class FavoriteButtonComponent {
  constructor(
    private productsService: ProductsService,
    private router: Router,
    private userService: UserService,
    private toastr: ToastrService
  ) {}

  @Input() product: Products;
  @Output() toggle = new EventEmitter<boolean>();
  isSubmitting = false;

  toggleFavorite() {
    this.isSubmitting = true;
    console.log(this.product);

    this.userService.isAuthenticated.pipe(concatMap(
      (authenticated) => {
        // Not authenticated? Push to login screen
        if (!authenticated) {
          this.router.navigateByUrl('/login');
          return of(null);
        }

        // Favorite the products if it isn't favorited yet
        if (!this.product.favorited) {
          return this.productsService.favorite(this.product.slug)
          .pipe(tap(
            data => {
              console.log('jajaj')
              this.isSubmitting = false;
              this.toastr.success('PRODUCT FAVORITE')
              this.toggle.emit(true);
            },
            err => this.isSubmitting = false
          ));

        // Otherwise, unfavorite the products
        } else {
          return this.productsService.unfavorite(this.product.slug)
          .pipe(tap(
            data => {
              this.isSubmitting = false;
              this.toastr.success('FAVORITE DELETE')
              this.toggle.emit(false);
            },
            err => this.isSubmitting = false
          ));
        }

      }
    )).subscribe();
  }
}

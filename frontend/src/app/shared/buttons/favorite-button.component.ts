import { Component, EventEmitter, Input, Output, OnDestroy } from '@angular/core';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { Products, ProductsService, UserService } from '../../core';
import { of, Subscription } from 'rxjs';

@Component({
  selector: 'app-favorite-button',
  templateUrl: './favorite-button.component.html'
})
export class FavoriteButtonComponent{
  onDestroyEvent: EventEmitter<string> = new EventEmitter();



  constructor(
    private productsService: ProductsService,
    private router: Router,
    private userService: UserService,
    private toastr: ToastrService


  ) {}
  private subscription: Subscription;

  @Input() product: Products;
  @Output() toggle = new EventEmitter<boolean>();
  isSubmitting = false;


  ngOnDestroy() {
    this.subscription.unsubscribe()
  }

  toggleFavorite(event) {
    event.stopPropagation();
    this.isSubmitting = true;


    this.subscription = this.userService.isAuthenticated.subscribe(

      (authenticated) => {
        console.log(authenticated)
        if (!authenticated) {
          this.router.navigateByUrl('/login');
          return of(null);
        }


        if (!this.product.favorited) {
          return this.productsService.favorite(this.product.slug)
          .subscribe(
            _ => {
  
              this.product.favorited = true;
              this.isSubmitting = false;
              this.toggle.emit(true);
              this.product.favoritesCount++;
              this.toastr.success('PRODUCT FAVORITE');
   
            },  
            
            _ => this.isSubmitting = false
          );

        } else {
          return this.productsService.unfavorite(this.product.slug)
          .subscribe(
            _ => {
              this.product.favorited = false;
              this.isSubmitting = false;
              this.toggle.emit(false);
              this.product.favoritesCount--;
              this.toastr.success('FAVORITE DELETE');
            },
            _ => this.isSubmitting = false
          );
          
        }



      }
    );
  }
}
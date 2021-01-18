import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Products, ProductsService, UserService, User, BuysProductsService } from '../core';
import { ToastrService } from 'ngx-toastr';
import { concatMap, flatMap, tap } from 'rxjs/operators';
import { of } from 'rxjs';


@Component({
  selector: 'app-product-page',
  styleUrls: ['./product.component.css'],
  templateUrl: './product.component.html'
})
export class ProductComponent implements OnInit {

  product: Products;



  currentUser: User;
  canModify: boolean;
  isSubmitting = false;
  isDeleting = false;
  constructor(
    private userService: UserService,
    private productsService: ProductsService,
    private buysProducts: BuysProductsService,
    private route: ActivatedRoute,
    private router: Router,
    private toastr: ToastrService
  ) { }

  ngOnInit() {
    // Retreive the prefetched product
    this.route.data.subscribe(
      (data: { product: Products; }) => {
        this.product = data.product;
      }
    );

    // Load the current user's data
    this.userService.currentUser.subscribe(
      (userData: User) => {
        this.currentUser = userData;

      }
    );
  }


  deleteProduct() {
    this.productsService.destroy(this.product['product'].slug)
      .subscribe(
        _ => {
          this.isDeleting = true;
          this.toastr.success('Producto Eliminado', 'Eliminado');
          this.router.navigateByUrl('/');
        }
      );
  }


  BuyProduct() {

    this.userService.isAuthenticated.pipe(concatMap(
      (authenticated) => {
        // Not authenticated? Push to login screen
        if (!authenticated) {
          this.router.navigateByUrl('/login');
          return of(null);
        }
        this.buysProducts.insert(this.product["product"].slug)
          .subscribe(data => {
            if (data['data'] == "okey") {
              this.toastr.success('Producto Comprado', 'Comprado');

              let data = this.product["product"].slug + ',' + this.product["product"].brand;
              this.productsService.UpKarmaProduct(data).subscribe()


            }
          })


      }
    )).subscribe();
  }


}
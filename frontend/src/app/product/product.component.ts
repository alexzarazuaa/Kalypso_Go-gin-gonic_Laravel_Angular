import { Component, OnInit } from '@angular/core';
import { ActivatedRoute ,Router} from '@angular/router';
import { Products,ProductsService, UserService, User } from '../core';
import { ToastrService } from 'ngx-toastr';


@Component({
  selector: 'app-product-page',
  styleUrls: ['./product.component.css'],
  templateUrl: './product.component.html'
})
export class ProductComponent implements OnInit {

  product: Products;
  productsService : ProductsService;
  

  currentUser: User;
  canModify: boolean;
  isSubmitting = false;
  isDeleting = false;
  constructor(
    private userService: UserService,
    private route: ActivatedRoute,
    private router: Router,
    private toastr: ToastrService
  ) { }

  ngOnInit() {
    // Retreive the prefetched product
    this.route.data.subscribe(
      (data: { product: Products; }) => {
        console.log(data.product,'detail')
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
    this.isDeleting = true;
    console.log('click ---->   ',this.product['product'].slug);
    this.productsService.destroy(this.product['product'].slug)
      .subscribe(
        success => {
          this.toastr.success('Producto Eliminado', 'Eliminado');
          this.router.navigateByUrl('/');
        }
      );
  }
}
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Products, UserService, User } from '../core';


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
    private route: ActivatedRoute,
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
}
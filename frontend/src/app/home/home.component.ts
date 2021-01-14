import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { BrandsService, UserService, Products } from '../core';

@Component({
  selector: 'app-home-page',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  constructor(
    private router: Router,
    private userService: UserService,
    private BrandsService: BrandsService
  ) { }



  isAuthenticated: boolean;

  brands: Array<string> = [];
  brandsLoaded = false;
  products: Products[];



  ngOnInit() {


    this.userService.isAuthenticated.subscribe(
      (authenticated) => {
        this.isAuthenticated = authenticated;

      }
    );


    this.BrandsService.getBrands(',client')
      .subscribe(data => {
        this.brands = data['data']['brands']

        this.brandsLoaded = true;
      });

  }

  // setListTo(type: string = '', ) {
  //   // If feed is requested but user is not authenticated, redirect to login
  //   if (type === 'feed' && !this.isAuthenticated) {
  //     this.router.navigateByUrl('/login');
  //     return;
  //   }

  // }

  FilterBrand(brand) {
    brand = "brands," + brand
    this.router.navigateByUrl('/shop', { state: { data: brand.key } });
  }
}

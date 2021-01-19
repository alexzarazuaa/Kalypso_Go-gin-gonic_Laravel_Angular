import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { OwlOptions } from 'ngx-owl-carousel-o';

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

  customOptions: OwlOptions = {
    loop: true,
    mouseDrag: false,
    touchDrag: false,
    pullDrag: false,
    dots: true,
    navSpeed: 700,
    navText: ['<', '>'],
    responsive: {
      0: {
        items: 1
      },
      400: {
        items: 1
      },
      740: {
        items: 1
      },
      940: {
        items: 1
      }
    },
    nav: true
  }



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
      });

  }


  FilterBrand(brand) {
    brand = "brands," + brand
    this.router.navigateByUrl('/shop', { state: { data: brand } });
  }
}

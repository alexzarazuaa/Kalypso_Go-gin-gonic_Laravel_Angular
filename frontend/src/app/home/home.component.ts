import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { ArticleListConfig, TagsService, BrandsService, UserService, Products } from '../core';

@Component({
  selector: 'app-home-page',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  constructor(
    private router: Router,
    //private tagsService: TagsService,
    private userService: UserService,
    private BrandsService: BrandsService
  ) { }



  isAuthenticated: boolean;
  listConfig: ArticleListConfig = {
    type: 'all',
    filters: {}
  };
  // tags: Array<string> = [];
  brands: Array<string> = [];
  brandsLoaded = false;
  products : Products[];

  // tagsLoaded = false;

  ngOnInit() {

    // this.userService.currentUser.subscribe(
    //   (userData) => {
    //     console.log("-+-----", userData)

    //     if(userData.type == 'admin'){
    //       this.router.navigateByUrl('/admin')
    //     }
    //   })

    this.userService.isAuthenticated.subscribe(
      (authenticated) => {
        this.isAuthenticated = authenticated;


        // set the article list accordingly
        if (authenticated) {
          this.setListTo('feed');
        } else {
          this.setListTo('all');
        }
      }
    );  


    this.BrandsService.getBrands()
      .subscribe(data => {
        this.brands = data['data']['brands']
        this.brandsLoaded = true;
      });

  }

  setListTo(type: string = '', filters: Object = {}) {
    // If feed is requested but user is not authenticated, redirect to login
    if (type === 'feed' && !this.isAuthenticated) {
      this.router.navigateByUrl('/login');
      return;
    }

    // Otherwise, set the list object
    this.listConfig = { type: type, filters: filters };
  }

  FilterBrand(brand) {
    console.log(brand)
    brand= "brands," + brand
    this.router.navigateByUrl('/shop', { state: { data :brand } });

  
  }
}

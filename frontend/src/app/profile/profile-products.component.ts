import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import { ProductListConfig, Profile } from '../core';

@Component({
  selector: 'app-profile-products',
  templateUrl: './profile-products.component.html'
})
export class ProfileProductsComponent implements OnInit {
  constructor(
    private route: ActivatedRoute,
    private router: Router
  ) {}

  profile: Profile;
  productsConfig: ProductListConfig = {
    type: 'all',
    filters: {}
  };

  ngOnInit() {
    this.route.parent.data.subscribe(
      (data: {profile: Profile}) => {
        this.profile = data.profile;
        this.productsConfig = {
          type: 'all',
          filters: {}
        }; // Only method I found to refresh article load on swap
        this.productsConfig.filters.author = this.profile.username;
      }
    );
  }

}

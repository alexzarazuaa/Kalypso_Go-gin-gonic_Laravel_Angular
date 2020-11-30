import { Component, Input } from '@angular/core';

import { BuyProduct, BuyProductListConfig, BuysProductsService } from '../../core';
@Component({
  selector: 'app-product-list',
  styleUrls: ['product-list.component.css'],
  templateUrl: './product-list.component.html'
})
export class ProductListComponent {
  constructor (
    private buysProductsService: BuysProductsService
  ) {}

  @Input() limit: number;
  @Input()
  set config(config: BuyProductListConfig) {
    if (config) {
      this.query = config;
      this.runQuery();
    }
  }

  query: BuyProductListConfig;
  results: BuyProduct[];
  loading = false;



  runQuery() {
    this.loading = true;
    this.results = [];


    this.buysProductsService.getAll(this.query)
    .subscribe(data => {
      this.loading = false;
      this.results = data.products;


    });
  }
}

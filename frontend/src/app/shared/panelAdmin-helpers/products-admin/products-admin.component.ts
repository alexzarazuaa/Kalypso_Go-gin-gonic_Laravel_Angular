import { Component, OnInit, Input } from '@angular/core';
import { Products, ProductsService } from '../../../core';

@Component({
  selector: 'app-products-admin',
  templateUrl: './products-admin.component.html',
  styleUrls: ['./products-admin.component.css']
})
export class ProductsAdminComponent implements OnInit {
  constructor(
    private productsService: ProductsService) { }

  results: Products[];

  ngOnInit() {
    this.results = [];
    this.productsService.getAll().subscribe(data => {
      this.results = data;
    })
  }

}

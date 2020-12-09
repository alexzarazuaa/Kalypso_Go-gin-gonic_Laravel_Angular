import { Component, Input, OnInit } from '@angular/core';

import { Products } from '../../../core';

@Component({
  selector: 'app-products-admin-preview',
  templateUrl: './products-admin-preview.html'
})
export class ProductsAdminPreviewComponent implements OnInit{
  @Input() products: Products;
  ngOnInit() {
        
  }
}
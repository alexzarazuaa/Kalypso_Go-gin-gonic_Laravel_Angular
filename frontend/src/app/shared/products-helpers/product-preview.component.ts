import { Component, Input } from '@angular/core';

import { BuyProduct } from '../../core';

@Component({
  selector: 'app-products-preview',
  templateUrl: './product-preview.component.html',
  styleUrls: ['./product-preview.component.css']
})
export class ProductsPreviewComponent {
  @Input() product: BuyProduct;


  ngOnInit() {
    this.product;
    console.log('Entra en el preview component', this.product);


  }
}
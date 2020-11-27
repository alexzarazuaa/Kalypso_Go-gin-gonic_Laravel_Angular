import { Component, Input } from '@angular/core';

import { BuyProduct } from '../../core';

@Component({
  selector: 'app-products-preview',
  templateUrl: './products-preview.component.html',
  styleUrls: ['./products-preview.component.css']
})
export class ProductsPreviewComponent {
  @Input() product: BuyProduct;


  ngOnInit() {
    this.product;
    console.log('Entra en el preview component', this.product);


  }
}
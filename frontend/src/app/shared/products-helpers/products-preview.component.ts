import { Component, Input } from '@angular/core';

import { BuyProduct } from '../../core';

@Component({
  selector: 'app-products-preview',
  templateUrl: './products-preview.component.html'
})
export class ProductsPreviewComponent {
  @Input() buyProducts: BuyProduct;


}
import { Component, Input } from '@angular/core';

import { Products } from '../../../core';

@Component({
  selector: 'app-products-admin-preview',
  templateUrl: './products-admin-preview.html'
})
export class ProductsAdminPreviewComponent{
  @Input() products: Products;
}
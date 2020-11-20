import { ModuleWithProviders, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { ProductComponent } from './product.component';
import { SharedModule } from '../shared';
import { ProductRoutingModule } from './products-routing.module';


@NgModule({
  declarations: [ProductComponent],
  imports: [
    SharedModule, 
    ProductRoutingModule
  ],

})
export class ProductModule {}
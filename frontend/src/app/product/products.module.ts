import { ModuleWithProviders, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { ProductComponent } from './product.component';
import { SharedModule } from '../shared';
import { ProductRoutingModule } from './products-routing.module';
import { ProductsResolver } from './product-resolver.service';

@NgModule({
  declarations: [
    ProductComponent
  ],
  imports: [
    SharedModule, 
    ProductRoutingModule,
    
  ],
  providers:[
    ProductsResolver
  ]

})
export class ProductModule {}
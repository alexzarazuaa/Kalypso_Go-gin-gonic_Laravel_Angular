import { ModuleWithProviders, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { ShopComponent } from './shop.component';
import { SharedModule } from '../shared';
import { ShopRoutingModule } from './shop-routing.module';

@NgModule({

  imports: [
    SharedModule, 
    ShopRoutingModule,
    
  ],
  declarations: [
    ShopComponent,
  ],

})
export class ShopModule {}
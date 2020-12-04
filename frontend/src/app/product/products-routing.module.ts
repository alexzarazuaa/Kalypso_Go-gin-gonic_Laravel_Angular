import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ProductComponent } from './product.component';
import { ProductsResolver } from './product-resolver.service';




const routes: Routes = [
  {
    path: ':id',
    component: ProductComponent,
    resolve:{
        product : ProductsResolver,
    }
  }

];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProductRoutingModule {}
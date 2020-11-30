import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { from } from 'rxjs';
import { ProductComponent } from './product.component';
import { ProductsResolver } from './product-resolver.service';




const routes: Routes = [
  {
    path: ':slug', component: ProductComponent,
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
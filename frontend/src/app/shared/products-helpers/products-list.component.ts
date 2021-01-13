import { Component, OnInit, Input } from '@angular/core';
import { Products, ProductsService,BrandsService } from '../../core';
import { Router } from '@angular/router';
import { flatMap } from 'rxjs/operators';


@Component({
  selector: 'app-products',
  templateUrl: './products-list.component.html',
  styleUrls: ['./products-list.component.css']
})
export class ProductslistComponent implements OnInit {
  constructor(
    private productsService: ProductsService,
    private BrandsService: BrandsService,
    private router: Router) { 
  }

  
    products : Products[];

  ngOnInit() {
if(history.state.data){
    if((history.state.data).includes('brands')){
          this.BrandsService.filterBrands(history.state.data)
      .subscribe(data => {
        this.products=data['product'];
      })
    }
  }else if(this.router.url==="/"){
    this.BrandsService.getBrands(',client')
      .subscribe(data => {

        this.products = data['data']['products']
      });
  }else{

  
     this.products = [];  
      this.productsService.getAll_goProd().subscribe(data => {
      this.products = data;

    })
  }
}

}



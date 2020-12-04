import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { BuyProduct, BuysProductsService } from '../core';


@Component({
  templateUrl: './shop.component.html'
})
export class ShopComponent implements OnInit {

    constructor(
        private buysProducts: BuysProductsService) { }
    
    
        buyProducts : BuyProduct[];
    
      ngOnInit() {
 

          console.log('products laravel');
        
      }
    
}
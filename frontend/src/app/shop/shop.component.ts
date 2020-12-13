import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Products, ProductsService } from '../core';


@Component({
  templateUrl: './shop.component.html',
  styleUrls: ['./shop.component.css']
})
export class ShopComponent implements OnInit {

    constructor(
        private productsService: ProductsService) { }
    
    
        products : Products[];
    
      ngOnInit() {


          console.log('SHOOOOOP');
        
      }
    
}
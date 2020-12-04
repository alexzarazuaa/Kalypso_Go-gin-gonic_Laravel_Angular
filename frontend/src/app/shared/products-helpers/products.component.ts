import { Component, OnInit, Input } from '@angular/core';
import { BuyProduct, BuysProductsService } from '../../core';

@Component({
  selector: 'app-products',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css']
})
export class ProductsComponent implements OnInit {
  constructor(
    private buysProducts: BuysProductsService) { }


    buyProducts : BuyProduct[];

  ngOnInit() {
     this.buyProducts = [];
    this.buysProducts.getAll().subscribe(data => {
      this.buyProducts = data;
      console.log(this.buyProducts,'products laravel');
    })
  }

}


/**
 * NG ON INIT TESTING DATA GO
 */

// ngOnInit() {
//   this.buyProducts = [];
//  this.buysProducts.getAll().subscribe(data => {
//    this.buyProducts = data;
//    console.log(this.buyProducts,'products laravel');
//  })
// }

import { Component, OnInit, Input } from '@angular/core';
import { BuyProduct, BuysProducts } from '../../core';

@Component({
  selector: 'app-products',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css']
})
export class ProductsComponent implements OnInit {
  constructor(
    private buysProducts: BuysProducts) { }
    buyProducts = [];

  ngOnInit() {
     this.buyProducts = [];
    console.log('Entra en el oninit');
    this.buysProducts.query().subscribe(data => {
      this.buyProducts = data;
      console.log(this.buyProducts,'products laravel');
    })
  }

}
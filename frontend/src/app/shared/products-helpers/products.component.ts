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

  ngOnInit() {
    console.log('holi');
    this.buysProducts.query().subscribe(data => {
      console.log(data);
    })
  }

}
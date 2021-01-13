import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Products, BrandsService  } from '../core';


@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})
export class AdminComponent implements OnInit {
  brands: Array<string> = [];
  products : Products[];

  constructor(
    private router: Router,
    private BrandsService : BrandsService
  ) { }

  ngOnInit() {
    console.log('CONSOLE PANEL ADMIN');

    this.BrandsService.getBrands(',admin')
    .subscribe(data => {
      this.brands = data['data']['brands']
      this.products = data['data']['products'];
    });




  }

  insertBD(){

    console.log('-++-+--+-- entra ------')
    this.BrandsService.insertRedisDb()
    .subscribe(data =>{
          console.log(data)
    })
  }

}
import { Component, OnInit } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { Products, BrandsService } from '../core';


@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})
export class AdminComponent implements OnInit {
  brands: Array<string> = [];
  products: Products[];

  constructor(
    private BrandsService: BrandsService,
    private toastr: ToastrService
  ) { }

  ngOnInit() {

    this.BrandsService.getBrands(',admin')
      .subscribe(data => {
        this.brands = data['data']['brands']
        this.products = data['data']['products'];
      });

  }


  insertBD() {
    this.BrandsService.insertRedisDb()
      .subscribe(data => {
        this.toastr.success('INSERT IN BD');
      })
  }

}
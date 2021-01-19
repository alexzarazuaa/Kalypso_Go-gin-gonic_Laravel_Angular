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

<<<<<<< HEAD
  DelRedis() {
    console.log('HEY ENTRE DEL ------')
    this.BrandsService.DelRedisService()
      .subscribe(data => {
        console.log(data)
        this.toastr.success('DELETE REDIS');
      })
  }

  UpRedis() {
    console.log('HEY ENTRE UP ------')
    this.BrandsService.UpRedisService()
      .subscribe(data => {
        console.log(data)
        this.toastr.success('UPDATE REDIS');
=======

  DelDB() {
    this.BrandsService.DelRatingDB()
      .subscribe(data => {
        this.toastr.success('Del Rating');
      })
  }

  DelRedis() {
    this.BrandsService.DelRatingRedis()
      .subscribe(data => {
        this.toastr.success('Del Rating');
>>>>>>> c5ac40264167a551fe2480c516dcb58c2813269a
      })
  }

}
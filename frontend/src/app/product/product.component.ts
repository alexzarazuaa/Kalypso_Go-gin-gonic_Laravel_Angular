import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

//import { Award, AwardsService } from '../core';

@Component({
  selector: 'app-product-page',
  styleUrls: ['./product.component.css'],
  templateUrl: './product.component.html'
})
export class ProductComponent implements OnInit {
  constructor() {
  }

  ngOnInit() {
    // console.log("EIII ENTRA EN Product Component");
  }
}
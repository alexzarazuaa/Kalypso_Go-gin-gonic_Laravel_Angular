import { Component } from '@angular/core';
import { OwlOptions } from 'ngx-owl-carousel-o';

@Component({
  selector: 'app-slider',
  templateUrl: './slider.component.html',
  styleUrls: ['./slider.component.css']
})

export class SliderComponent {

  slide = [
     {
       id: 1,
       src:'https://via.placeholder.com/600/92c952',
       alt:'Side 1',
       title:'Side 1'
     },
     {
       id: 2,
       src:'https://via.placeholder.com/600/771796',
       alt:'Side 2',
       title:'Side 2'
     },
     {
       id: 3,
       src:'https://via.placeholder.com/600/24f355',
       alt:'Side 3',
       title:'Side 3'
     },
     {
       id: 4,
       src:'https://via.placeholder.com/600/d32776',
       alt:'Side 4',
       title:'Side 4'
     },
     {
       id: 5,
       src:'https://via.placeholder.com/600/f66b97',
       alt:'Side 5',
       title:'Side 5'
     }
   ]
 
   constructor() { }

  customOptions: OwlOptions = {
    loop: true,
    mouseDrag: false,
    touchDrag: false,
    pullDrag: false,
    dots: false,
    navSpeed: 600,
    navText: ['·', '·'],
    responsive: {
      0: {
        items: 1 
      },
      400: {
        items: 2
      },
      760: {
        items: 3
      },
      1000: {
        items: 4
      }
    },
    nav: true
  }

}
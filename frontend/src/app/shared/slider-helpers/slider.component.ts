import { Component, OnInit } from '@angular/core';


@Component({
  selector: 'app-slider',
  templateUrl: './slider.component.html',
  styleUrls: ['./slider.component.css']
})

export class SliderComponent {

  constructor() {}

  
  customOptions: any = {
    loop: true,
    items: 1,
    autoplay:true,
    dots: true
  }

  slides = [
    {
      image: "frontend/src/img/Polo1.webp",
      text : "JAJA"
    },
    {
      image: "../../img/preview2.png"
    },
  ]


}

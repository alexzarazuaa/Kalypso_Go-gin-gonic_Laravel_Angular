
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PanelGo, PanelService } from '../core';


@Component({
  selector: 'app-admin-panel',
  templateUrl: './panel.component.html',
  styleUrls: ['./panel.component.css']
})
export class PanelComponent implements OnInit {
  constructor(
    private panelService: PanelService) { }


    algo : PanelGo[];

  ngOnInit() {
     this.algo = [];
    this.panelService.getAll().subscribe(data => {
      this.algo = data;
      console.log(this.algo,'products laravel');
    })
  }

}
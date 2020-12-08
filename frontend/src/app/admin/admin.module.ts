
import { NgModule } from '@angular/core';
import { SharedModule } from '../shared';
import { CommonModule } from '@angular/common';
import {AdminComponent } from './admin.component';
import {AdminRoutingModule } from './admin-routing.module';


@NgModule({
  imports: [
    CommonModule,
    SharedModule,
   AdminRoutingModule
  ],
  declarations: [
   AdminComponent
  ]
})
export class AdminModule { }

import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { RouterModule } from '@angular/router';
import { ProductslistComponent , ProductsPreviewComponent } from './products-helpers';
import { FavoriteButtonComponent, FollowButtonComponent } from './buttons';
import { ListErrorsComponent } from './list-errors.component';
import { ShowAuthedDirective } from './show-authed.directive';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    RouterModule
  ],
  declarations: [
    FavoriteButtonComponent,
    FollowButtonComponent,
    ListErrorsComponent,
    ShowAuthedDirective,
    ProductslistComponent,
    ProductsPreviewComponent

  ],
  exports: [
    CommonModule,
    FavoriteButtonComponent,
    FollowButtonComponent,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    ListErrorsComponent,
    RouterModule,
    ShowAuthedDirective,
    ProductslistComponent,
    ProductsPreviewComponent

  ]
})
export class SharedModule {}

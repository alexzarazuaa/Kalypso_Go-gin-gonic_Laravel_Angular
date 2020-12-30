import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { CarouselModule } from 'ngx-owl-carousel-o';

import { AppComponent } from './app.component';
import { AuthModule } from './auth/auth.module';
import { HomeModule } from './home/home.module';
import { ProductModule } from './product/products.module';
import {
  FooterComponent,
  HeaderComponent,
  SharedModule
} from './shared';
import { SliderComponent } from './shared/slider-helpers/slider.component';
import { AppRoutingModule } from './app-routing.module';
import { CoreModule } from './core/core.module';
import { from } from 'rxjs';

@NgModule({
  declarations: [AppComponent, FooterComponent, HeaderComponent,SliderComponent],
  imports: [
    BrowserModule,
    CoreModule,
    SharedModule,
    HomeModule,
    AuthModule,
    AppRoutingModule,
    ProductModule,
    CarouselModule    
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {}

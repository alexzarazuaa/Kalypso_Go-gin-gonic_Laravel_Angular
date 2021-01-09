import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpHeaders, HttpClient, HttpParams } from '@angular/common/http';
import { Observable ,  throwError } from 'rxjs';

import { JwtService } from './jwt.service';
import { catchError } from 'rxjs/operators';

@Injectable()
export class ApiService {
  constructor(
    private http: HttpClient,
    private jwtService: JwtService
  ) {}

  private formatErrors(error: any) {
    return  throwError(error.error);
  }

  get(path: string, params: HttpParams = new HttpParams()): Observable<any> {
    return this.http.get(`${environment.laravel}${path}`, { params })
      .pipe(catchError(this.formatErrors));
  }

  put(path: string, body: Object = {}): Observable<any> {
    return this.http.put(
      `${environment.laravel}${path}`,
      JSON.stringify(body)
    ).pipe(catchError(this.formatErrors));
  }

  post(path: string, body: Object = {}): Observable<any> {
    return this.http.post(
      `${environment.laravel}${path}`,
      JSON.stringify(body)
    ).pipe(catchError(this.formatErrors));
  }

  delete(path): Observable<any> {
    return this.http.delete(
      `${environment.laravel}${path}`
    ).pipe(catchError(this.formatErrors));
  }


  /* ===================  GO METHODS ===================*/

    /**
   * Method go enviroment GET data
   * @param path 
   * @param params 
   */

  get_Go(path: string, params: HttpParams = new HttpParams()): Observable<any> {
    return this.http.get(`${environment.go}${path}`, { params })
      .pipe(catchError(this.formatErrors));
  }

  /**
   * 
   * @param path  METHOD POST GO  ENVIROMENT DATA
   * @param body 
   */

  post_Go(path: string, body: Object = {}): Observable<any> {
    return this.http.post(
      `${environment.go}${path}`,
      JSON.stringify(body),
    ).pipe(catchError(this.formatErrors));
  }
  put_Go(path: string, body: Object = {}): Observable<any> {
    return this.http.put(
      `${environment.go}${path}`,
      JSON.stringify(body)
    ).pipe(catchError(this.formatErrors));
  }


  delete_go(path): Observable<any> {
    return this.http.delete(
      `${environment.go}${path}`
    ).pipe(catchError(this.formatErrors));
  }



    /* ===================  GO PRODUCTS METHODS ===================*/

    /**
   * Method PRODUCTS Go enviroment GET data
   * @param path 
   * @param params 
   */

  // get_goProd(path: string, params: HttpParams = new HttpParams()): Observable<any> {
  //   return this.http.get(`${environment.go_prods}${path}`, { params })
  //     .pipe(catchError(this.formatErrors));
  // }


  get_goProd(path: string, params: HttpParams = new HttpParams()): Observable<any> {
    return this.http.get(`${environment.go_prods}${path}`, { params })
      .pipe(catchError(this.formatErrors));
  }

  post_goProd(path: string, body: Object = {}): Observable<any> {
    return this.http.post(
      `${environment.go_prods}${path}`,
      JSON.stringify(body),
    ).pipe(catchError(this.formatErrors));
  }

  delete_goProd(path): Observable<any> {
    return this.http.delete(
      `${environment.go_prods}${path}`
    ).pipe(catchError(this.formatErrors));
  }


}

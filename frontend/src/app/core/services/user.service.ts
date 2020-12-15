import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, BehaviorSubject, ReplaySubject } from 'rxjs';

import { ApiService } from './api.service';
import { JwtService } from './jwt.service';
import { User } from '../models';
import { map, distinctUntilChanged } from 'rxjs/operators';
import { Certificate } from 'crypto';


@Injectable()
export class UserService {
  private currentUserSubject = new BehaviorSubject<User>({} as User);
  public currentUser = this.currentUserSubject.asObservable().pipe(distinctUntilChanged());

  private userCredentials;

  private isAuthenticatedSubject = new ReplaySubject<boolean>(1);
  public isAuthenticated = this.isAuthenticatedSubject.asObservable();

  constructor(
    private apiService: ApiService,
    private http: HttpClient,
    private jwtService: JwtService
  ) { }

  // Verify JWT in localstorage with server & load user's info.
  // This runs once on application startup.
  populate() {
    // If JWT detected, attempt to get & store user's info
    if (this.jwtService.getToken()) {
      this.apiService.get_Go('/user/')
        .subscribe(
          data => this.setAuth(data.user),
          err => this.purgeAuth()
        );
    } else {
      // Remove any potential remnants of previous auth states
      this.purgeAuth();
    }
  }

  setAuth(user: User) {
    // Save JWT sent from server in localstorage
    this.jwtService.saveToken(user.token);
    console.log("TOKEN==========> : ");
    console.log(user.token);



    this.jwtService.saveToken(user.token);
    // Set current user data into observable
    this.currentUserSubject.next(user);
    // Set isAuthenticated to true
    this.isAuthenticatedSubject.next(true);
  }

  purgeAuth() {
    // Remove JWT from localstorage
    this.jwtService.destroyToken();
    // Set current user to an empty object
    this.currentUserSubject.next({} as User);
    // Set auth status to false
    this.isAuthenticatedSubject.next(false);
  }

  attemptAuth(type, credentials): Observable<User> {
    const route = (type === 'login') ? 'login' : '';
    console.log('=>>>>>>>>>>>>>>>>>>>', credentials);



    let algo = this.apiService.post_Go('/users/' + route, { user: credentials })

    algo.subscribe(
      data => {
        console.log(data)

        if (data.user.type == "client") {
          this.setAuth(data.user);
          return data
        } else {
          console.log("--------------------")
          let algo1 = this.apiService.post('/users/' + route, { user: credentials })

          algo1.subscribe(
            data => {
              if (Object.keys(data.user).length !== 0) {
                this.setAuth(data.user);
              }
            }
          )

        }
      }
    );
    return algo
    // return credentials;


    // return this.apiService.post('/users/' + route, { user: credentials })
    //           .pipe(map(
    //             data => {
    //               console.log("---------------")

    //               this.setAuth(data.user);
    //               console.log(data.user.type);
    //               return data;
    //             }
    //           ));



  }

  login_admin() {

    if ((Object.keys(this.getCurrentUser()).length === 0) && (this.userCredentials !== undefined)) {

      console.log("********************************************")
      console.log(this.userCredentials)

    } else {
      console.log(this.getCurrentUser())
    }

  }
  getCurrentUser(): User {
    return this.currentUserSubject.value
  }

  // Update the user on the server (email, pass, etc)
  update(user): Observable<User> {
    return this.apiService
      .put('/user/', { user })
      .pipe(map(data => {
        // Update the currentUser observable
        this.currentUserSubject.next(data.user);
        return data.user;
      }));
  }

}

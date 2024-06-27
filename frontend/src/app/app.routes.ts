import { Routes } from '@angular/router';
import {HeaderLayout} from "./layout/header/header.component";
import {MainComponent} from "./pages/admin/main/main.component";
import {SignInComponent} from "./pages/sign-in/sign-in.component";
import {SignUpComponent} from "./pages/sign-up/sign-up.component";
import {HomeComponent} from "./pages/home/home.component";

export const routes: Routes = [
  {
    path:'', component: HomeComponent
  },
  {
    path:'admin',component:HeaderLayout , children:[
      {
        path:'main',component:MainComponent
      }
    ]
  },
  {
    path:'signIn',
    component:SignInComponent
  },
  {
    path:'signUp',component:SignUpComponent
  }
];

import { Component } from '@angular/core';

import {MatSidenavModule} from "@angular/material/sidenav";
import {MatButtonModule} from "@angular/material/button";
import {MainComponent} from "../../pages/admin/main/main.component";

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [
    MatSidenavModule, MatButtonModule, MainComponent,
  ],
  templateUrl: './header.component.html',
  styleUrl: './header.component.scss'
})
export class HeaderLayout {
  showFiller = false;
}

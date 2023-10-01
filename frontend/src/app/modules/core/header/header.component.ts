import { Component } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { IconsModule } from '../../icons/icons.module';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss'],
  standalone: true,
  imports: [RouterLink, RouterLinkActive, IconsModule],
})
export class HeaderComponent {}

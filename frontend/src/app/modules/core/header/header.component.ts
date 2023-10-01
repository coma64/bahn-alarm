import { Component } from '@angular/core';
import { FeatherModule } from 'angular-feather';
import { RouterLink, RouterLinkActive } from '@angular/router';

@Component({
    selector: 'app-header',
    templateUrl: './header.component.html',
    styleUrls: ['./header.component.scss'],
    standalone: true,
    imports: [RouterLink, RouterLinkActive, FeatherModule]
})
export class HeaderComponent {

}

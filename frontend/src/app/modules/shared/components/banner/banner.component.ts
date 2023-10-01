import { Component, HostBinding, Input } from '@angular/core';
import { FeatherModule } from 'angular-feather';

@Component({
    selector: 'app-banner',
    templateUrl: './banner.component.html',
    styleUrls: ['./banner.component.scss'],
    standalone: true,
    imports: [FeatherModule],
})
export class BannerComponent {
  @Input() @HostBinding('class') type = 'error' as const;
}

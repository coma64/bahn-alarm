import { Component, HostBinding, Input } from '@angular/core';
import { IconsModule } from '../../../icons/icons.module';

@Component({
  selector: 'app-banner',
  templateUrl: './banner.component.html',
  styleUrls: ['./banner.component.scss'],
  standalone: true,
  imports: [IconsModule],
})
export class BannerComponent {
  @Input() @HostBinding('class') type = 'error' as const;
}

import { Component } from '@angular/core';
import { Select } from '@ngxs/store';
import { Observable } from 'rxjs';
import { State } from '../../../state/state';
import { trackById } from '../../shared/track-by-id';
import { SpinnerComponent } from '../../shared/components/spinner/spinner.component';
import { NgIf, NgFor, AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-alarmed-devices-list',
    templateUrl: './alarmed-devices-list.component.html',
    styleUrls: ['./alarmed-devices-list.component.scss'],
    standalone: true,
    imports: [
        NgIf,
        NgFor,
        SpinnerComponent,
        AsyncPipe,
    ],
})
export class AlarmedDevicesListComponent {
  @Select() alarmedDevices$!: Observable<State['alarmedDevices']>;
  protected readonly trackById = trackById;
}

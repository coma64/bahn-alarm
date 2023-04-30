import { Component } from '@angular/core';
import { Select } from '@ngxs/store';
import { Observable } from 'rxjs';
import { State } from '../../../state/state';
import { trackById } from '../../shared/track-by-id';

@Component({
  selector: 'app-alarmed-devices-list',
  templateUrl: './alarmed-devices-list.component.html',
  styleUrls: ['./alarmed-devices-list.component.scss'],
})
export class AlarmedDevicesListComponent {
  @Select() alarmedDevices$!: Observable<State['alarmedDevices']>;
  protected readonly trackById = trackById;
}

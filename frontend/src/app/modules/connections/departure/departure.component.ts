import { Component, Input } from '@angular/core';
import { TrackedDeparture } from '../../../api';
import { ToHumanStatusPipe } from '../to-human-status.pipe';
import { NextRelativeTimePipe } from '../../shared/pipes/next-relative-time.pipe';
import { ToRelativeTimePipe } from '../../shared/pipes/to-relative-time.pipe';
import { NgIf, AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-departure',
    templateUrl: './departure.component.html',
    styleUrls: ['./departure.component.scss'],
    standalone: true,
    imports: [
        NgIf,
        AsyncPipe,
        ToRelativeTimePipe,
        NextRelativeTimePipe,
        ToHumanStatusPipe,
    ],
})
export class DepartureComponent {
  @Input() departure?: TrackedDeparture;
}

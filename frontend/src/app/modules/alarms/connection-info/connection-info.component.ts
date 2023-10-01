import { Component, Input } from '@angular/core';
import { SimpleConnection } from '../../../api';
import { NextRelativeTimePipe } from '../../shared/pipes/next-relative-time.pipe';
import { ToRelativeTimePipe } from '../../shared/pipes/to-relative-time.pipe';
import { NgIf, AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-connection-info',
    templateUrl: './connection-info.component.html',
    styleUrls: ['./connection-info.component.scss'],
    standalone: true,
    imports: [
        NgIf,
        AsyncPipe,
        ToRelativeTimePipe,
        NextRelativeTimePipe,
    ],
})
export class ConnectionInfoComponent {
  @Input() connection?: SimpleConnection;
}

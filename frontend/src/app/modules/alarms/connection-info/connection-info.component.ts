import { Component, Input } from '@angular/core';
import { SimpleConnection } from '../../../api';

@Component({
  selector: 'app-connection-info',
  templateUrl: './connection-info.component.html',
  styleUrls: ['./connection-info.component.scss'],
})
export class ConnectionInfoComponent {
  @Input() connection?: SimpleConnection;
}

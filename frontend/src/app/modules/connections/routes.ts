import { Routes } from '@angular/router';
import { ConnectionsComponent } from './connections/connections.component';
import { EditConnectionComponent } from './edit-connection/edit-connection.component';

const routes: Routes = [
  {
    path: '',
    component: ConnectionsComponent,
  },
  {
    path: 'add',
    component: EditConnectionComponent,
  },
  {
    path: 'edit/:connectionId',
    component: EditConnectionComponent,
  },
];

export default routes;

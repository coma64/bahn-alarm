import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
import './init-dayjs';
import { AppModule } from './app/app.module';
import LogRocket from 'logrocket';

LogRocket.init('qmpbyd/bahn-alarm');

platformBrowserDynamic()
  .bootstrapModule(AppModule)
  .catch((err) => console.error(err));

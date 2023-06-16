export * from './alarms.service';
import { AlarmsService } from './alarms.service';
export * from './auth.service';
import { AuthService } from './auth.service';
export * from './bahn.service';
import { BahnService } from './bahn.service';
export * from './notifications.service';
import { NotificationsService } from './notifications.service';
export * from './tracking.service';
import { TrackingService } from './tracking.service';
export const APIS = [
  AlarmsService,
  AuthService,
  BahnService,
  NotificationsService,
  TrackingService,
];

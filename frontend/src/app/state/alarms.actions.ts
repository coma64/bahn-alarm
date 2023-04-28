import { Alarm } from '../api';

export namespace AlarmsActions {
  export class Fetched {
    static readonly type = '[Alarms] Fetched';
    constructor(readonly items: Array<Alarm>) {}
  }

  export class Deleted {
    static readonly type = '[Alarms] Deleted';
    constructor(readonly deletedId: number) {}
  }
}

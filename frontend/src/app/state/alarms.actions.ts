import { Urgency } from '../api';

export namespace AlarmsActions {
  export class Fetch {
    static readonly type = '[Alarms] Fetch';
  }

  export class Delete {
    static readonly type = '[Alarms] Delete';
    constructor(readonly targetId: number) {}
  }

  export class FilterByUrgency {
    static readonly type = '[Alarms] Filter by urgency';
    constructor(readonly urgency?: Urgency) {}
  }

  export class SetPage {
    static readonly type = '[Alarms] Set page';
    constructor(readonly newPage: number) {}
  }
}

import { TrackingStats } from '../api';

export namespace ConnectionStats {
  export class Fetched {
    static readonly type = '[ConnectionStats] Fetched';
    constructor(public readonly stats: TrackingStats) {}
  }
}

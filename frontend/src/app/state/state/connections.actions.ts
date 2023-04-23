import { TrackedConnection } from '../../api';

export namespace Connections {
  export class Fetched {
    static readonly type = '[Connections] Fetched';
    constructor(public readonly connections: Array<TrackedConnection>) {}
  }
}
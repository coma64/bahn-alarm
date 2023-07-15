import { TrackedConnection } from '../api';

export namespace Connections {
  export class Fetched {
    static readonly type = '[Connections] Fetched';
    constructor(public readonly connections: TrackedConnection[]) {}
  }

  export class Created {
    static readonly type = '[Connections] Created';

    constructor(public readonly newConnection: TrackedConnection) {}
  }

  export class Updated {
    static readonly type = '[Connections] Updated';

    constructor(public readonly updatedConnection: TrackedConnection) {}
  }
}

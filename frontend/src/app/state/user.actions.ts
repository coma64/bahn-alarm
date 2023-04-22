export namespace User {
  export class LoginSuccess {
    static readonly type = '[User] login success';
    constructor(public readonly name: string) {}
  }

  export class Logout {
    static readonly type = '[User] logging out';
  }
}

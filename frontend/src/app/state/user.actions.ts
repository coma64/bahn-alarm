import { User } from '../api';

export namespace UserActions {
  export class LoginSuccess {
    static readonly type = '[User] login success';
    constructor(public readonly user: User) {}
  }

  export class Logout {
    static readonly type = '[User] logging out';
  }

  export class RegisteredPushNotifications {
    static readonly type = '[User] Registered push notifications';
    constructor(readonly pushSubId: number) {}
  }
}

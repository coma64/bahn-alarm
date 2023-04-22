export class LoginSuccess {
  static readonly type = '[User] login success';
  constructor(public readonly name: string) {}
}

import * as dayjs from 'dayjs';
import * as utc from 'dayjs/plugin/utc';
import * as relativeTime from 'dayjs/plugin/relativeTime';
import * as duration from 'dayjs/plugin/duration';

dayjs.extend(utc);
dayjs.extend(relativeTime);
dayjs.extend(duration);

export class RelativeTime {
  static readonly ZERO_DATE = dayjs.utc('0001-01-01T00:01:01Z');

  private readonly utc: dayjs.Dayjs;
  private readonly local: dayjs.Dayjs;
  readonly str: string;

  constructor(dateTime: dayjs.Dayjs) {
    this.utc = RelativeTime.copyTime(dateTime, dayjs.utc());
    this.local = this.utc.local();
    this.str = this.local.format('HH:mm');
  }

  static fromIso(dateTime: string): RelativeTime {
    return new RelativeTime(dayjs.utc(dateTime));
  }

  static fromTimeInput(time: string): RelativeTime {
    const [hours, minutes] = time.split(':').map((v) => parseInt(v));
    return new RelativeTime(
      dayjs().set('hours', hours).set('minutes', minutes).utc(),
    );
  }

  static now(): RelativeTime {
    return new RelativeTime(dayjs.utc());
  }

  static copyTime(from: dayjs.Dayjs, to: dayjs.Dayjs): dayjs.Dayjs {
    return to
      .set('hour', from.hour())
      .set('minute', from.minute())
      .set('second', from.second())
      .set('millisecond', from.millisecond());
  }

  toIso(): string {
    return this.utc.toISOString();
  }

  toIsoZeroBased(): string {
    const zeroBasedDateTime = RelativeTime.copyTime(
      this.utc,
      RelativeTime.ZERO_DATE,
    );
    return zeroBasedDateTime.toISOString();
  }

  timeUntil(alwaysShowMinutes: boolean): string {
    // TODO: add "in" prefix option and "now"
    let nextTime = this.utc;
    const now = dayjs.utc();
    if (nextTime.isBefore(now)) {
      nextTime = nextTime.add(1, 'day');
    }

    const diffMinutes = nextTime.diff(now, 'minutes');
    const diffHours = nextTime.diff(now, 'hours');
    if (diffHours > 0) {
      if (alwaysShowMinutes) {
        return `${diffHours}h ${diffMinutes - diffHours * 60}m`;
      }
      return `${diffHours}h`;
    } else {
      return `${diffMinutes}m`;
    }
  }
}

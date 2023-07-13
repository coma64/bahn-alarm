import dayjs from "dayjs/esm";
import utc from 'dayjs/esm/plugin/utc';
import relativeTime from 'dayjs/esm/plugin/relativeTime';
import duration from 'dayjs/esm/plugin/duration';
import localizedFormat from 'dayjs/esm/plugin/localizedFormat';

dayjs.extend(localizedFormat);
dayjs.extend(utc);
dayjs.extend(relativeTime);
dayjs.extend(duration);

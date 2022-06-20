export enum DateUnit {
  second = 1000,
  minute = DateUnit.second * 60,
  hour = DateUnit.minute * 60
}

export enum SecondsIn {
  minutes = 60,
  hour = minutes * 60,
  days = hour * 24
}

export function getCurrentDatePlus(hours: number, minutes = 0): Date {
  const date = new Date();
  date.setHours(date.getHours() + hours);
  date.setMinutes(date.getMinutes() + minutes);

  return date;
}

export function getDateDiffIn(dateUnit: DateUnit, isoDate?: string): string {
  if (!isoDate) {
    return '';
  }

  const targetDate = new Date(isoDate);
  const now = new Date();
  const diffTime = Math.abs(targetDate.getTime() - now.getTime());
  const diffInIndicatedUnit = Math.ceil(diffTime / dateUnit);

  return diffInIndicatedUnit.toString();
}

export function prettyDiffDate(isoDate?: string): string {
  if (!isoDate) {
    return '';
  }

  const targetDate = new Date(isoDate);
  const now = new Date();
  const diffTime = Math.abs(targetDate.getTime() - now.getTime());
  const diffInIndicatedUnit = Math.ceil(diffTime / DateUnit.second);

  if (diffInIndicatedUnit < SecondsIn.minutes) {
    return `${diffInIndicatedUnit} seconds ago`;
  }

  if (diffInIndicatedUnit < SecondsIn.hour) {
    return `${Math.ceil(diffInIndicatedUnit / SecondsIn.minutes)} minutes ago`;
  }

  if (diffInIndicatedUnit < SecondsIn.days) {
    return `${Math.ceil(diffInIndicatedUnit / SecondsIn.hour)} hours ago`;
  }

  return `${Math.ceil(diffInIndicatedUnit / SecondsIn.days)} days ago`;
}

export enum DateUnit {
  second = 1000,
  minute = DateUnit.second * 60,
  hour = DateUnit.minute * 60
}

// TODO: it is actually a diff between now and date, rename
export function getCurrentDatePlus(hours: number, minutes = 0): Date {
  const date = new Date();
  date.setHours(date.getHours() + hours);
  date.setMinutes(date.getMinutes() + minutes);

  return date;
}

// TODO: Not very good name
export function getDateUnitRemains(dateUnit: DateUnit, isoDate?: string): string {
  if (!isoDate) {
    return '';
  }

  const targetDate = new Date(isoDate);
  const now = new Date();
  const diffTime = Math.abs(targetDate.getTime() - now.getTime());
  const diffInIndicatedUnit = Math.ceil(diffTime / dateUnit);

  return diffInIndicatedUnit.toString();
}

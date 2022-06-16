export function getCurrentDatePlus(hours: number, minutes = 0): Date {
  const date = new Date();
  date.setHours(date.getHours() + hours);
  date.setMinutes(date.getMinutes() + minutes);

  return date;
}

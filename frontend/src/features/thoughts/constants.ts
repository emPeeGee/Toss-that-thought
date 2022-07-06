import { getCurrentDatePlus } from 'utils/date';

interface LifetimeOptionModel {
  value: string;
  label: string;
}

export const lifetimeOptions: LifetimeOptionModel[] = [
  { value: getCurrentDatePlus(168).toISOString(), label: '7 days' },
  { value: getCurrentDatePlus(72).toISOString(), label: '3 days' },
  { value: getCurrentDatePlus(24).toISOString(), label: '1 days' },
  { value: getCurrentDatePlus(12).toISOString(), label: '12 hours' },
  { value: getCurrentDatePlus(4).toISOString(), label: '4 hours' },
  { value: getCurrentDatePlus(1).toISOString(), label: '1 hour' },
  { value: getCurrentDatePlus(0, 30).toISOString(), label: '30 minutes' },
  { value: getCurrentDatePlus(0, 5).toISOString(), label: '5 minutes' },
  { value: getCurrentDatePlus(0, 1).toISOString(), label: '1 minutes' }
];

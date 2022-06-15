import React, { ReactNode } from 'react';
import { createStyles } from '@mantine/core';

type Props = {
  children: ReactNode;
};

const useStyles = createStyles(() => ({
  shell: {
    width: '100%',
    height: '100vh',
    padding: 0,
    margin: 0,
    display: 'flex',
    flexDirection: 'column',

    '& > *': {
      width: '100%',
      flexGrow: 1
    }
  }
}));

export function AppShell({ children }: Props) {
  const { classes } = useStyles();
  return <main className={classes.shell}>{children}</main>;
}

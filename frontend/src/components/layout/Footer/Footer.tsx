import React from 'react';
import { Text, createStyles, Anchor } from '@mantine/core';
import logo from 'assets/logo.svg';

const useStyles = createStyles((theme, _params, getRef) => ({
  footer: {
    marginTop: 'auto'
  },

  container: {
    display: 'flex',
    flexDirection: 'column',
    marginTop: '64px',
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[1],
    padding: '16px 32px'
  },

  info: {
    ref: getRef('info'),
    display: 'flex',
    flexDirection: 'column',

    p: {
      margin: 0
    }
  },

  links: {
    padding: '2rem 0',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',

    h4: {
      margin: 0
    },

    ul: {
      margin: 0,
      padding: 0,
      listStyle: 'none',
      display: 'flex',
      justifyContent: 'flex-start',
      alignItems: 'center',

      li: {
        color: theme.colors.gray[7],
        padding: '0 1rem'
      }
    }
  },

  logo: {
    display: 'flex',
    justifyContent: 'flex-start',
    alignItems: 'center',

    h3: {
      margin: 0
    },

    img: {
      width: '46px',
      height: '46px'
    }
  }
}));

export function Footer() {
  const { classes } = useStyles();

  return (
    <footer className={classes.footer}>
      <div className={classes.container}>
        <div className={classes.info}>
          <div className={classes.logo}>
            <img src={logo} alt="" />
            <h3>Toss that thought.</h3>
          </div>
          <div>
            <Text color="dimmed">
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ad, alias aliquid animi{' '}
            </Text>
          </div>
        </div>

        <nav className={classes.links}>
          <h4>Links</h4>
          <ul>
            <li>
              <Anchor href="/">About</Anchor>
            </li>
            <li>
              <Anchor href="/">Code</Anchor>
            </li>
            <li>
              <Anchor href="/">Other</Anchor>
            </li>
          </ul>
        </nav>

        <Text color="gray">Build with 💙 by emPeeGee.</Text>
      </div>
    </footer>
  );
}

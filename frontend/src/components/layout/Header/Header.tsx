import React, { useContext, useEffect, useState } from 'react';
import { MoonStars, Sun } from 'tabler-icons-react';
import { ActionIcon, Text, useMantineColorScheme } from '@mantine/core';

import { UserContext, tokenIdentifier } from 'features/authentication/';
import { Anchor } from 'components';
import logo from 'assets/logo.svg';
import {
  IconGroup,
  Icon,
  Items,
  ItemsRight,
  ListItem,
  UnorderedList,
  Wrapper
} from './Header.styles';

export function Header() {
  const { colorScheme, toggleColorScheme } = useMantineColorScheme();
  const [isDark, setIsDark] = useState(false);
  const userContext = useContext(UserContext);

  useEffect(() => {
    setIsDark(colorScheme === 'dark');
  }, [colorScheme]);

  const logout = () => {
    console.log('Logout');
    userContext?.setUser(null);
    userContext?.setToken(null);
    localStorage.removeItem(tokenIdentifier);
  };

  return (
    <Wrapper>
      <Items>
        <IconGroup to="/" tabIndex={0}>
          <Icon src={logo} alt="Application logo" />
          <Text weight={700} color={isDark ? 'white' : 'black'}>
            Toss That Thought
          </Text>
        </IconGroup>
        <ItemsRight>
          <UnorderedList>
            <ListItem>
              <ActionIcon
                variant="outline"
                color={isDark ? 'yellow' : 'blue'}
                onClick={() => toggleColorScheme()}
                title="Toggle color scheme">
                {isDark ? <Sun size={18} /> : <MoonStars size={18} />}
              </ActionIcon>
            </ListItem>

            {userContext?.user ? (
              <>
                <ListItem>
                  <Anchor to="profile" title={userContext?.user?.username ?? 'Profile'} />
                </ListItem>
                <ListItem>
                  <Anchor to="profile/recent" title="Recent thoughts" />
                </ListItem>
                <ListItem onClick={() => logout()}>
                  <Anchor to="/" title="Log out" />
                </ListItem>
              </>
            ) : (
              <>
                <ListItem>
                  <Anchor to="sign-in" title="Sign In" />
                </ListItem>

                <ListItem>
                  <Anchor to="sign-up" title="Sign Up" />
                </ListItem>
              </>
            )}
          </UnorderedList>
        </ItemsRight>
      </Items>
    </Wrapper>
  );
}

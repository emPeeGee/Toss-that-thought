import React, { useEffect, useState } from 'react';
import { MoonStars, Sun } from 'tabler-icons-react';
import { ActionIcon, Text, useMantineColorScheme } from '@mantine/core';

import { Anchor } from 'components/navigation/Anchor/Anchor';
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

  useEffect(() => {
    setIsDark(colorScheme === 'dark');
  }, [colorScheme]);

  return (
    <Wrapper>
      <Items>
        <IconGroup to="/">
          <Icon src={logo} alt="Application logo" />
          <Text weight={700} color={isDark ? 'white' : 'black'}>
            Toss That Thought
          </Text>
        </IconGroup>
        <ItemsRight>
          <UnorderedList>
            <ActionIcon
              variant="outline"
              color={isDark ? 'yellow' : 'blue'}
              onClick={() => toggleColorScheme()}
              title="Toggle color scheme">
              {isDark ? <Sun size={18} /> : <MoonStars size={18} />}
            </ActionIcon>
            <ListItem>
              <Anchor to="sign-up" title="Create Account" />
            </ListItem>
            <ListItem>
              <Anchor to="about" title="About" />
            </ListItem>
            <ListItem>
              <Anchor to="sign-in" title="Sign In" />
            </ListItem>
          </UnorderedList>
        </ItemsRight>
      </Items>
    </Wrapper>
  );
}

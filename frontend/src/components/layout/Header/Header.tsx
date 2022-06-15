import React from 'react';
import { Anchor } from 'components/navigation/Anchor/Anchor';
import logo from 'assets/logo.svg';

import {
  IconGroup,
  Icon,
  IconTitle,
  Items,
  ItemsRight,
  ListItem,
  UnorderedList,
  Wrapper
} from './Header.styles';

export function Header() {
  return (
    <Wrapper>
      <Items>
        <IconGroup to="/">
          <Icon src={logo} />
          <IconTitle>Toss That Thought</IconTitle>
        </IconGroup>
        <ItemsRight>
          <UnorderedList>
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

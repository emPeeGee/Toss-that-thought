import React from 'react';
import logo from '../../../assets/logo.svg';
import { Anchor } from '../../navigation/Anchor/Anchor';

import {
  Icon,
  IconTitle,
  Items,
  ItemsLeft,
  ItemsRight,
  ListItem,
  UnorderedList,
  Wrapper
} from './Header.styles';

export function Header() {
  return (
    <Wrapper>
      <Items>
        <ItemsLeft>
          <Icon src={logo} />
          <IconTitle>Toss That Thought</IconTitle>
        </ItemsLeft>
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

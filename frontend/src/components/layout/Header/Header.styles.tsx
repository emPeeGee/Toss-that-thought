import styled from 'styled-components';
import { Link } from 'react-router-dom';

export const Wrapper = styled.header`
  padding: 16px 8px;
  margin-bottom: 32px;
  border-bottom: 1px solid #e0e3e8;
`;

export const IconGroup = styled(Link)`
  display: flex;
  justify-content: center;
  align-items: center;
  text-decoration: none;

  &:focus {
    outline: 2px solid orange;
  }
`;

export const Icon = styled.img`
  width: 36px;
  height: 36px;
`;

export const Items = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
`;

export const ItemsRight = styled.div``;

export const UnorderedList = styled.ul`
  list-style: none;
  margin: 0;
  display: flex;
  justify-content: center;
  align-items: center;
`;
export const ListItem = styled.li`
  padding: 4px 12px;
`;

import React from 'react';
import { StyledAnchor } from './Anchor.styles';

export function Anchor({ to, title }: { to: string; title: string }) {
  return <StyledAnchor to={to}>{title}</StyledAnchor>;
}

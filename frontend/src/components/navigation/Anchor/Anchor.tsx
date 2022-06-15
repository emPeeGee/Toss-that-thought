import React from 'react';
import { Anchor as MantineAnchor } from '@mantine/core';
import { Link } from 'react-router-dom';

export function Anchor({ to, title }: { to: string; title: string }) {
  return (
    <MantineAnchor component={Link} to={to}>
      {title}
    </MantineAnchor>
  );
}

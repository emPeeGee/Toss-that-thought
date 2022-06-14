import React from 'react';
import { Container, Textarea, Input, Select } from '@mantine/core';

export function Thought() {
  return (
    <Container size="md">
      <h2>Paste a password, secret message or private link below.</h2>
      <h3>Keep sensitive info out of your email and chat logs.</h3>

      <Textarea
        aria-label="Paste a password, secret message or private link below"
        placeholder="Autosize with no rows limit"
        label="Autosize with no rows limit"
        autosize
        minRows={10}
      />

      <Input placeholder="Your email" radius="md" size="md" />
      <Select
        label="Your favorite framework/library"
        placeholder="Pick one"
        data={[
          { value: '7', label: '7 days' },
          { value: '3', label: '3 days' },
          { value: '1', label: '1 days' }
        ]}
      />
    </Container>
  );
}

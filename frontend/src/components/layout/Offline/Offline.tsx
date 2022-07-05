import React from 'react';
import { Button, Container, Group, Text, Title } from '@mantine/core';
import { PlugConnectedX } from 'tabler-icons-react';

export function Offline() {
  return (
    <Container>
      <Group direction="column" align="center">
        <PlugConnectedX size={80} strokeWidth={0.7} />
        <Title order={1}>Oops. No internet connection.</Title>
        <Text>Make sure wifi or cellular data is turned on and then try again.</Text>
        <Button variant="outline">Try again</Button>
      </Group>
    </Container>
  );
}

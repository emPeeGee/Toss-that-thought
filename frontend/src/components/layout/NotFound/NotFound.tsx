import React from 'react';
import { Button, Container, Paper, Text, Title } from '@mantine/core';
import { ArrowBarLeft } from 'tabler-icons-react';

export function NotFound() {
  return (
    <Container size="md">
      <Paper shadow="xs" p="md" my="md" withBorder>
        <Title order={1} my="md">
          Page Not Found
        </Title>
        <Text my="md">
          Looks like you&apos;ve followed a broken link or entered a URL that doesn&apos;t exist on
          this app.
        </Text>
        <Button variant="outline" my="md" leftIcon={<ArrowBarLeft />}>
          Back to main page
        </Button>
      </Paper>
    </Container>
  );
}

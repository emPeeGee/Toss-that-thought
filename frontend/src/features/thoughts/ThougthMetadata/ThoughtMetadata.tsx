import React from 'react';
import { Button, Code, Container, Divider, Grid, Paper, Text } from '@mantine/core';
import { Link, useParams } from 'react-router-dom';
import { Bolt, MessageCircle2 } from 'tabler-icons-react';

export function ThoughtMetadata() {
  const { metadataKey } = useParams();

  return (
    <Container my="xl">
      <Paper shadow="xs" p="md" my="md" withBorder>
        <Text size="xl" weight="500">
          Share the link:
        </Text>
        <Code color="yellow" my="xs" style={{ fontSize: '20px' }}>
          http://localhost:9000/thought/{metadataKey}
        </Code>
        <Text color="gray" size="sm">
          Requires a passphrase.
        </Text>
      </Paper>
      <Paper shadow="xs" p="md" my="md" withBorder>
        <Text size="xl" weight="500">
          The passphrase:
        </Text>
        <Code color="orange" my="xs" style={{ fontSize: '20px' }}>
          http://localhost:9000/thought/{metadataKey}
        </Code>
      </Paper>
      <Grid align="center" mx={0} my="lg">
        <Text size="xl" weight="500">
          Expires in 7 days.
        </Text>
        <Text color="dimmed" pl="sm">
          (2022-06-15@06:11:33 UTC)
        </Text>
      </Grid>
      <Button<typeof Link>
        to="/metadata/1234"
        leftIcon={<Bolt size={24} />}
        variant="light"
        my="lg"
        fullWidth
        component={Link}>
        Burn this thought
      </Button>
      <Divider my="md" />
      <Text color="dimmed" my="lg" style={{ fontStyle: 'italic' }}>
        * Burning a secret will delete it before it has been read (click to confirm).
      </Text>
      <Divider my="md" />
      <Button<typeof Link>
        to="/"
        leftIcon={<MessageCircle2 size={24} />}
        variant="light"
        my="lg"
        fullWidth
        component={Link}>
        Create another thought
      </Button>
    </Container>
  );
}

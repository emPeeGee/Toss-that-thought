import React from 'react';
import { Alert, Button, Container, Divider, Text, TextInput } from '@mantine/core';
import { ArrowBackUp, Bolt } from 'tabler-icons-react';
// import { useParams } from 'react-router-dom';

export function ThoughtBurn() {
  // const { metadataKey } = useParams();

  return (
    <Container size="md" my="xl">
      <Text>Thought {`1543gf ${4}`}</Text>
      <TextInput my="md" placeholder="Enter passphrase here" />

      <Button fullWidth my="lg" leftIcon={<Bolt size={24} />}>
        Burn this thought
      </Button>
      <Button fullWidth my="lg" leftIcon={<ArrowBackUp size={24} />}>
        Cancel
      </Button>

      <Divider my="md" />

      <Alert
        withCloseButton
        color="red"
        title="Advice"
        closeButtonLabel="Close advice"
        onClose={() => {}}>
        Burning a secret is permanent and cannot be undone
      </Alert>
    </Container>
  );
}

import React, { useState } from 'react';
import { Alert, Button, Container, Divider, Text, TextInput, Title } from '@mantine/core';
import { ArrowBackUp, Bolt } from 'tabler-icons-react';
import { Link, useParams } from 'react-router-dom';
// import { useParams } from 'react-router-dom';

export function ThoughtBurn() {
  const { metadataKey } = useParams();
  const [isAdviceAlertVisible, setIsAdviceAlertVisible] = useState(true);

  return (
    <Container size="md" my="xl">
      <Title order={1} my="lg">
        Burn thought 🔥
      </Title>
      <Text>Thought {`1543gf ${4}`}</Text>
      <TextInput my="md" placeholder="Enter passphrase here" />

      <Button<typeof Link>
        component={Link}
        to={`/metadata/${metadataKey}`}
        fullWidth
        my="lg"
        color="orange"
        leftIcon={<Bolt size={24} />}>
        Burn this thought
      </Button>
      <Button<typeof Link>
        component={Link}
        to={`/metadata/${metadataKey}`}
        fullWidth
        variant="outline"
        color="gray"
        my="lg"
        leftIcon={<ArrowBackUp size={24} />}>
        Cancel
      </Button>

      {isAdviceAlertVisible && (
        <>
          <Divider my="md" />
          <Alert
            withCloseButton
            color="red"
            title="Advice"
            closeButtonLabel="Close advice"
            onClose={() => {
              setIsAdviceAlertVisible(false);
            }}>
            Burning a secret is permanent and cannot be undone
          </Alert>
        </>
      )}
    </Container>
  );
}

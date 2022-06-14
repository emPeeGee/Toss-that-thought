import React from 'react';
import {
  Container,
  Textarea,
  Select,
  TextInput,
  Button,
  Text,
  Alert,
  Highlight
} from '@mantine/core';
import { AlertCircle } from 'tabler-icons-react';

export function Thought() {
  return (
    <Container size="md">
      <Text my="lg" align="center" color="dark" weight={700} style={{ fontSize: 30 }}>
        Paste a password, secret message or private link below.
      </Text>
      <Text my="lg" align="center" size="md" color="gray" weight={500}>
        Keep sensitive info out of your email and chat logs.
      </Text>

      <Textarea
        aria-label="Paste a password, secret message or private link below"
        my="lg"
        placeholder="Your thought goes here"
        autosize
        minRows={10}
      />

      <TextInput
        required
        variant="filled"
        my="md"
        label="Passphrase"
        description="A word or a passphrase that is difficult to guess"
        radius="md"
        size="md"
      />

      <Select
        required
        label="Lifetime"
        description="After passing lifetime thought will be no more available"
        size="md"
        my="md"
        variant="filled"
        data={[
          { value: '7', label: '7 days' },
          { value: '3', label: '3 days' },
          { value: '1', label: '1 days' }
        ]}
        transition="pop-top-left"
        transitionDuration={80}
        transitionTimingFunction="ease"
      />

      <Button variant="light" my="lg" fullWidth style={{ marginBottom: '64px' }}>
        Create that thought
      </Button>

      <Text color="gray" align="center">
        * A thought link only works once and then disappears forever.
      </Text>
      <Text color="gray" align="center">
        Sign up for a free account to set passphrases for extra security along with additional
        privacy options. We will even email the link for you if you want.
      </Text>

      <Alert
        my="md"
        icon={<AlertCircle size={16} />}
        withCloseButton
        closeButtonLabel="Close create a account alert">
        <Highlight
          highlight="Stay anonymous!"
          highlightStyles={(theme) => ({
            background: 'transparent',
            color: theme.colors.blue[9],
            fontWeight: 500
          })}>
          Stay anonymous! Create a account
        </Highlight>
      </Alert>
    </Container>
  );
}

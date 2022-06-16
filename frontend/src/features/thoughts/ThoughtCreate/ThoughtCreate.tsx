import React, { useState } from 'react';
import {
  Container,
  Textarea,
  Select,
  TextInput,
  Button,
  Text,
  Alert,
  Title,
  Anchor
} from '@mantine/core';
import { AlertCircle, MessageCircle2 } from 'tabler-icons-react';
import { Link } from 'react-router-dom';

export function ThoughtCreate() {
  const [isCreateAccountAlertVissible, setIsCreateAccountAlertVissible] = useState(true);
  return (
    <Container>
      <Title order={1} my="lg" align="center" color="dark" style={{ fontSize: 30 }}>
        Paste a password, secret message or private link below. 💭
      </Title>
      <Text my="lg" align="center" size="md" color="gray" weight={500}>
        Keep sensitive info out of your email and chat logs. 👀
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

      <Button<typeof Link>
        component={Link}
        fullWidth
        to="/metadata/1234"
        variant="light"
        my="lg"
        leftIcon={<MessageCircle2 size={24} />}
        style={{ marginBottom: '64px' }}>
        Create that thought
      </Button>

      <Text color="gray" align="center" style={{ fontStyle: 'italic' }}>
        * A thought link only works once and then disappears forever.
      </Text>
      <Text color="gray" align="center">
        Sign up for a{' '}
        <Anchor component={Link} to="/sign-up" aria-label="Create free account">
          free account
        </Anchor>{' '}
        to set passphrases for extra security along with additional privacy options. We will even
        email the link for you if you want.
      </Text>

      {isCreateAccountAlertVissible && (
        <Alert
          my="md"
          color="yellow"
          icon={<AlertCircle size={28} />}
          title="Stay anonymous"
          withCloseButton
          closeButtonLabel="Close alert 'Create a account using a temporary email address'"
          onClose={() => setIsCreateAccountAlertVissible(false)}>
          Create an account using a temporary email address.
        </Alert>
      )}
    </Container>
  );
}

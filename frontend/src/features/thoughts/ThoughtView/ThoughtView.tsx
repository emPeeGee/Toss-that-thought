import React, { useState } from 'react';
import { Alert, Button, Container, Divider, Text, Textarea, TextInput } from '@mantine/core';
import { ArrowForwardUp, Eye } from 'tabler-icons-react';
import { Link } from 'react-router-dom';
// import { useParams } from 'react-router-dom';

export function ThoughtView() {
  // const { thoughtKey} = useParams();
  const [isCarefulAlertVisible, setIsCarefulAlertVisible] = useState(true);
  const [isThoughtValid] = useState(true);
  const [isPassphraseCorrect, setIsPassphraseCorrect] = useState(true);

  if (!isThoughtValid) {
    return (
      <Container size="md" my="xl">
        <Alert color="red" title="Hmmm">
          It either never existed or has already been viewed.
        </Alert>
      </Container>
    );
  }

  return (
    <Container size="md" my="xl">
      {isPassphraseCorrect ? (
        <>
          <Textarea
            disabled
            label="This thought is for you:"
            variant="filled"
            size="xl"
            my="xl"
            value="Lorem"
          />
          <Button<typeof Link>
            to="/"
            component={Link}
            fullWidth
            my="lg"
            leftIcon={<ArrowForwardUp size={24} />}>
            Reply with another thought
          </Button>
        </>
      ) : (
        <>
          <Text size="xl">This thought requires a passphrase:</Text>
          <TextInput variant="filled" my="md" placeholder="Enter your passphrase here" />
          <Button
            fullWidth
            my="lg"
            leftIcon={<Eye size={24} />}
            onClick={() => setIsPassphraseCorrect(true)}>
            View thought
          </Button>
        </>
      )}

      {isCarefulAlertVisible && (
        <>
          <Divider my="md" />
          <Alert
            withCloseButton
            color="orange"
            title="Careful"
            closeButtonLabel="Close advice"
            onClose={() => {
              setIsCarefulAlertVisible(false);
            }}>
            We will only show it once
          </Alert>
        </>
      )}
    </Container>
  );
}

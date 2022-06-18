import React, { useEffect, useState } from 'react';
import {
  Alert,
  Button,
  Container,
  Divider,
  LoadingOverlay,
  PasswordInput,
  Text,
  Textarea
} from '@mantine/core';
import { ArrowForwardUp, Eye, Lock } from 'tabler-icons-react';
import { useForm } from 'react-hook-form';
import { Link, useParams } from 'react-router-dom';
import { api } from 'services/http';
import { ThoughtPassphraseRequest, ThoughtResponse } from '../thought.model';

export function ThoughtView() {
  const { register, handleSubmit } = useForm<ThoughtPassphraseRequest>({
    mode: 'onChange'
  });
  const { thoughtKey } = useParams();
  const [isCarefulAlertVisible, setIsCarefulAlertVisible] = useState(true);
  const [isError, setIsError] = useState(false);
  const [isPassphrasePhasePassed, setIsPassphrasePhasePassed] = useState(false);
  const [isPassphraseCorrect, setIsPassphraseCorrect] = useState(true);
  const [isLoading, setIsLoading] = useState(false);
  const [thought, setThought] = useState('');

  useEffect(() => {
    setIsLoading(true);
    api
      .get({ url: `thought/${thoughtKey}` })
      .catch((err) => {
        console.log(err);
        setIsError(true);
      })
      .finally(() => {
        setIsLoading(false);
      });
  }, []);

  const viewThought = (data: ThoughtPassphraseRequest) => {
    setIsLoading(true);

    api
      .post<ThoughtPassphraseRequest, ThoughtResponse>({ url: `thought/${thoughtKey}`, body: data })
      .then((response) => {
        setThought(response.thought);
        setIsPassphrasePhasePassed(true);
      })
      .catch((err) => {
        console.log(err);
        setIsPassphraseCorrect(false);
      })
      .finally(() => setIsLoading(false));
  };

  if (isError) {
    return (
      <Container size="md" my="xl">
        <Alert color="red" title="Hmm">
          It either never existed or has already been viewed.
        </Alert>
      </Container>
    );
  }

  return (
    <Container size="md" my="xl">
      <LoadingOverlay visible={isLoading} />
      {isPassphrasePhasePassed ? (
        <>
          <Textarea
            disabled
            label="This thought is for you:"
            variant="filled"
            size="xl"
            my="xl"
            value={thought}
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
          {!isPassphraseCorrect && (
            <Alert title="Oops..." color="red" my="lg">
              Double check that passphrase
            </Alert>
          )}

          <form onSubmit={handleSubmit(viewThought)}>
            <Text size="xl">This thought requires a passphrase:</Text>
            <PasswordInput
              {...register('passphrase', { required: false, value: '' })}
              my="md"
              placeholder="Enter passphrase here"
              toggleTabIndex={0}
              icon={<Lock size={16} />}
            />
            <Button fullWidth my="lg" variant="light" leftIcon={<Eye size={24} />} type="submit">
              View thought
            </Button>
          </form>
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

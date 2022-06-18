import React, { useEffect, useState } from 'react';
import {
  Alert,
  Button,
  Container,
  Divider,
  LoadingOverlay,
  PasswordInput,
  Text,
  Title
} from '@mantine/core';
import { useForm } from 'react-hook-form';
import { ArrowBackUp, Bolt, Lock } from 'tabler-icons-react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import { api } from 'services/http';
import { isObjectEmpty } from 'utils/is-empty';
import { ThoughtBurnRequest, ThoughtMetadataModel } from '../thought.model';

// TODO: Parse date
export function ThoughtBurn() {
  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm<ThoughtBurnRequest>({
    mode: 'onChange'
  });
  const { metadataKey } = useParams();
  const [thoughtMetadata, setThoughtMetadata] = useState<ThoughtMetadataModel>();
  const [isAdviceAlertVisible, setIsAdviceAlertVisible] = useState(true);
  const [isLoading, setIsLoading] = useState(false);
  const [isError, setIsError] = useState(false);
  const [isBurnError, setIsBurnError] = useState(false);

  useEffect(() => {
    setIsLoading(true);
    api
      .get<ThoughtMetadataModel>({ url: `metadata/${metadataKey}` })
      .then((response) => {
        setThoughtMetadata(response);
      })
      .catch(() => setIsError(true))
      .finally(() => setIsLoading(false));
  }, [metadataKey]);

  const burnThought = (data: ThoughtBurnRequest) => {
    setIsLoading(true);
    setIsBurnError(false);

    api
      .post({
        url: `thought/${metadataKey}/burn`,
        body: data
      })
      .then(() => {
        navigate(`/metadata/${metadataKey}`, {
          replace: true
        });
      })
      .catch((err) => {
        setIsBurnError(true);
        console.log(err);
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  if (isError) {
    return (
      <Container>
        <Alert color="red" title="Oops...">
          Such thought does not exist
        </Alert>
      </Container>
    );
  }

  // TODO: Not ready yet, but should be checked when though will be viewed
  // TODO: Kinda redundant with below isBurned
  if (thoughtMetadata?.isViewed) {
    return (
      <Container size="md" my="xl">
        <Alert my="md" color="orange" title="Cannot burn!">
          The thought was already viewed 3 seconds ago
        </Alert>
        <Button<typeof Link>
          fullWidth
          to={`/metadata/${metadataKey}`}
          component={Link}
          variant="outline"
          leftIcon={<ArrowBackUp size={28} />}>
          Back
        </Button>
      </Container>
    );
  }

  if (thoughtMetadata?.isBurned) {
    // TODO: More friendly time
    return (
      <Container size="md" my="xl">
        <Alert my="md" color="orange" title="Cannot burn!">
          The thought was already burned{' '}
          {new Date(thoughtMetadata?.burnedDate?.Time ?? '').toDateString()}
        </Alert>
        <Button<typeof Link>
          fullWidth
          to={`/metadata/${metadataKey}`}
          component={Link}
          variant="outline"
          leftIcon={<ArrowBackUp size={28} />}>
          Back
        </Button>
      </Container>
    );
  }

  // TODO: Check empty password flow
  return (
    <Container size="md" my="xl">
      <LoadingOverlay visible={isLoading} />

      {isBurnError && (
        <Alert color="red" title="Oops...">
          Double check that passphrase
        </Alert>
      )}
      <Title order={1} my="lg">
        Burn thought 🔥
      </Title>
      <PasswordInput
        {...register('passphrase', { required: false, value: '' })}
        my="md"
        placeholder="Enter passphrase here"
        toggleTabIndex={0}
        icon={<Lock size={16} />}
      />

      <Container px={0} my="lg">
        <Button
          fullWidth
          variant="light"
          color="orange"
          leftIcon={<Bolt size={24} />}
          onClick={handleSubmit(burnThought)}
          disabled={!isObjectEmpty(errors)}>
          Burn this thought
        </Button>
        {errors.passphrase && <Text color="red">Passphrase is required</Text>}
      </Container>
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
            color="violet"
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

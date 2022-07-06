import React, { useEffect, useState } from 'react';
import {
  Alert,
  Button,
  Container,
  Divider,
  LoadingOverlay,
  PasswordInput,
  Title
} from '@mantine/core';
import { useForm } from 'react-hook-form';
import { ArrowBackUp, Bolt, Lock } from 'tabler-icons-react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import { api } from 'services/http';
import { prettyDiffDate } from 'utils';
import { ThoughtPassphraseRequest, ThoughtMetadataModel } from '../thought.model';

export function ThoughtBurn() {
  const navigate = useNavigate();
  const { register, handleSubmit } = useForm<ThoughtPassphraseRequest>({
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

  const burnThought = (data: ThoughtPassphraseRequest) => {
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

  // TODO: Kinda redundant with below isBurned
  if (thoughtMetadata?.isViewed) {
    return (
      <Container size="md" my="xl">
        <Alert my="md" color="orange" title="Cannot burn!">
          The thought was already viewed {prettyDiffDate(thoughtMetadata.viewedDate)}.
          {new Date(thoughtMetadata?.viewedDate ?? '').toLocaleString()}
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
    return (
      <Container size="md" my="xl">
        <Alert my="md" color="orange" title="Cannot burn!">
          The thought was already burned {prettyDiffDate(thoughtMetadata.burnedDate)}.{' '}
          {new Date(thoughtMetadata?.burnedDate ?? '').toLocaleString()}
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

  return (
    <Container size="md" my="xl">
      <LoadingOverlay visible={isLoading} />

      {isBurnError && (
        <Alert color="red" title="Oops...">
          Double check that passphrase
        </Alert>
      )}
      <form onSubmit={handleSubmit(burnThought)}>
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
            type="submit"
            leftIcon={<Bolt size={24} />}>
            Burn this thought
          </Button>
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
      </form>

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

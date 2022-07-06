import React, { useContext, useState } from 'react';
import {
  Container,
  Textarea,
  Select,
  TextInput,
  Button,
  Text,
  Alert,
  Title,
  Anchor,
  LoadingOverlay
} from '@mantine/core';
import { AlertCircle, MessageCircle2 } from 'tabler-icons-react';
import { Link, useNavigate } from 'react-router-dom';
import { SubmitHandler, useForm, Controller } from 'react-hook-form';
import { api } from 'services/http';
import { UserContext } from 'features/authentication';
import { ThoughtCreateRequest, lifetimeOptions, ThoughtMetadataModel } from 'features/thoughts';

export function ThoughtCreate() {
  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    control,
    formState: { errors }
  } = useForm<ThoughtCreateRequest>({
    defaultValues: {
      lifetime: lifetimeOptions[1].value
    }
  });
  const [isCreateAccountAlertVisible, setIsCreateAccountAlertVisible] = useState(true);
  const [isLoading, setIsLoading] = useState(false);
  const [isError, setIsError] = useState(false);
  const userContext = useContext(UserContext);

  const thoughtSubmit: SubmitHandler<ThoughtCreateRequest> = async (data) => {
    setIsLoading(true);
    api
      .post<ThoughtCreateRequest, ThoughtMetadataModel>({
        url: 'create',
        body: data,
        token: userContext?.token
      })
      .then((response) => {
        navigate(`metadata/${response.metadataKey}`, {
          replace: true,
          state: response
        });
      })
      .catch((reason) => {
        setIsError(true);
        console.log(reason);
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  return (
    <Container>
      <LoadingOverlay visible={isLoading} />
      <Title order={1} my="lg" align="center" color="dark" style={{ fontSize: 30 }}>
        Paste a password, secret message or private link below. 💭
      </Title>
      <Text my="lg" align="center" size="md" color="gray" weight={500}>
        Keep sensitive info out of your email and chat logs. 👀
      </Text>

      <Textarea
        {...register('thought', { required: true, maxLength: 1000 })}
        aria-label="Paste a password, secret message or private link below"
        my="lg"
        placeholder="Your thought goes here"
        autosize
        minRows={10}
        error={errors.thought ? 'Field is invalid' : null}
      />

      <TextInput
        variant="filled"
        my="md"
        label="Passphrase"
        description="A word or a passphrase that is difficult to guess"
        radius="md"
        size="md"
        {...register('passphrase')}
        error={errors.passphrase ? 'Field is invalid' : null}
      />

      <Controller
        name="lifetime"
        control={control}
        rules={{ required: true }}
        render={({ field }) => (
          <Select
            {...field}
            required
            label="Lifetime"
            description="After passing lifetime thought will be no more available"
            size="md"
            my="md"
            variant="filled"
            data={lifetimeOptions}
            transition="pop-top-left"
            transitionDuration={80}
            transitionTimingFunction="ease"
            error={errors.lifetime ? 'Field is invalid' : null}
          />
        )}
      />

      <Container my="lg" px={0} style={{ marginBottom: '64px' }}>
        <Button<typeof Link>
          component={Link}
          fullWidth
          to="/metadata/1234"
          variant="light"
          leftIcon={<MessageCircle2 size={24} />}
          onClick={handleSubmit(thoughtSubmit)}>
          Create that thought
        </Button>
        {isError && <Text color="red">An unknown error occurred</Text>}
      </Container>

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

      {isCreateAccountAlertVisible && (
        <Alert
          my="md"
          color="yellow"
          icon={<AlertCircle size={28} />}
          title="Stay anonymous"
          withCloseButton
          closeButtonLabel="Close alert 'Create a account using a temporary email address'"
          onClose={() => setIsCreateAccountAlertVisible(false)}>
          Create an account using a temporary email address.
        </Alert>
      )}
    </Container>
  );
}

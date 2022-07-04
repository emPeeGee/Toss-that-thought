import React, { useState } from 'react';
import { Button, Container, LoadingOverlay, PasswordInput, TextInput, Title } from '@mantine/core';
import { ArrowBackUp, Bolt, Lock, UserCircle } from 'tabler-icons-react';
import { Link, useNavigate } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import {
  AuthenticationResponse,
  CredentialsModel
} from 'features/authentication/authentication.model';
import { api } from 'services/http';
import { useToken } from 'hooks/useToken';

export function SignIn() {
  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm<CredentialsModel>({
    mode: 'onChange'
  });

  const navigate = useNavigate();
  const { token, setToken } = useToken();
  const [isLoading, setIsLoading] = useState(false);

  if (token) {
    console.log(token, ' you are authenticated');
  }

  const signIn = (data: CredentialsModel) => {
    setIsLoading(true);
    api
      .post<CredentialsModel, AuthenticationResponse>({
        url: 'signIn',
        body: data,
        auth: true
      })
      .then((response) => {
        console.log(response);
        setToken(response.token);

        navigate(`/profile/`, {
          replace: true
        });
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  return (
    <Container>
      <LoadingOverlay visible={isLoading} />

      <Title order={1} my="lg">
        Sign In
      </Title>

      <form onSubmit={handleSubmit(signIn)}>
        <TextInput
          {...register('username', { required: true, value: '' })}
          required
          label="Username"
          placeholder="Enter your username"
          error={errors.username ? 'Username is required' : null}
          icon={<UserCircle size={14} />}
        />

        <PasswordInput
          {...register('password', { required: true, value: '' })}
          required
          my="md"
          label="Password"
          placeholder="Enter your password"
          error={errors.password ? 'Password is required' : null}
          toggleTabIndex={0}
          icon={<Lock size={16} />}
        />

        <Container px={0} my="lg">
          <Button
            fullWidth
            variant="light"
            color="primary"
            type="submit"
            leftIcon={<Bolt size={24} />}>
            Sign in
          </Button>
          <Button<typeof Link>
            component={Link}
            to="/"
            fullWidth
            variant="outline"
            color="gray"
            my="lg"
            leftIcon={<ArrowBackUp size={24} />}>
            Go home
          </Button>
        </Container>
      </form>
    </Container>
  );
}

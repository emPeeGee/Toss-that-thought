import React, { useContext, useEffect, useState } from 'react';
import { Button, Container, LoadingOverlay, PasswordInput, TextInput, Title } from '@mantine/core';
import { ArrowBackUp, Bolt, Lock, UserCircle } from 'tabler-icons-react';
import { Link, useNavigate } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import { AuthenticationResponse, CredentialsModel, UserContext } from 'features/authentication';
import { api } from 'services/http';

export function SignIn() {
  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm<CredentialsModel>({
    mode: 'onChange'
  });

  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(false);
  const userContext = useContext(UserContext);

  useEffect(() => {
    if (userContext?.user) {
      navigate(`/profile/`, {
        replace: true
      });
    }
  }, [userContext?.user]);

  const signIn = (data: CredentialsModel) => {
    setIsLoading(true);
    api
      .post<CredentialsModel, AuthenticationResponse>({
        url: 'signIn',
        body: data,
        auth: true
      })
      .then((response) => {
        userContext?.setToken(response.token);
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

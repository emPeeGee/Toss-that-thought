import React, { useContext } from 'react';
import { Code, Container, Text, Group, Title } from '@mantine/core';
import { UserContext } from 'features/authentication/user.context';

export function Profile() {
  const userContext = useContext(UserContext);

  return (
    <Container>
      <Title order={1}>Profile</Title>
      <Group direction="column">
        <Group>
          <Text>Your name: </Text>
          <Code color="violet" my="xs" style={{ fontSize: '20px' }}>
            {userContext?.user?.username}
          </Code>
        </Group>

        <Group>
          <Text>Your username: </Text>
          <Code color="primary" my="xs" style={{ fontSize: '20px' }}>
            {userContext?.user?.name}
          </Code>
        </Group>

        <Group>
          <Text>User created at: </Text>
          <Code color="yellow" my="xs" style={{ fontSize: '20px' }}>
            {userContext?.user?.createdAt}
          </Code>
        </Group>

        <Group>
          <Text>Last user update: </Text>
          <Code color="yellow" my="xs" style={{ fontSize: '20px' }}>
            {userContext?.user?.updatedAt}
          </Code>
        </Group>
      </Group>
    </Container>
  );
}

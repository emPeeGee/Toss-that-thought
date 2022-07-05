import React, { useContext, useEffect, useState } from 'react';
import { Container, Title, Text, LoadingOverlay } from '@mantine/core';
import { ThoughtMetadataModel } from 'features/thoughts/thought.model';
import { UserContext } from 'features/authentication';
import { api } from 'services/http';
import { Received, ReceivedComponent } from 'features/thoughts/RecentThoughts/Received/Received';
import {
  NotReceived,
  NotReceivedComponent
} from 'features/thoughts/RecentThoughts/NotReceived/NotReceived';

type RecentThoughtsComponent = React.FunctionComponent & {
  NotReceived: NotReceivedComponent;
  Received: ReceivedComponent;
};

// eslint-disable-next-line react/function-component-definition
export const RecentThoughts: RecentThoughtsComponent = () => {
  const userContext = useContext(UserContext);
  const [thoughts, setThoughts] = useState<ThoughtMetadataModel[]>([]);
  const [isError, setIsError] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    setIsLoading(true);
    api
      .get<ThoughtMetadataModel[]>({
        url: 'recent',
        token: userContext?.token
      })
      .then((response) => {
        setThoughts(response);
      })
      .catch(() => {
        setIsError(true);
      })
      .finally(() => {
        setIsLoading(false);
      });
  }, [userContext?.token]);

  if (isError) {
    return (
      <Container>
        <Text color="red">Something went wrong</Text>
      </Container>
    );
  }

  return (
    <Container>
      <LoadingOverlay visible={isLoading} />

      <Title order={1} my="lg">
        Recent thoughts 🍜
      </Title>

      <RecentThoughts.NotReceived
        thoughts={thoughts.filter((thought) => !thought.isBurned && !thought.isViewed)}
      />

      <RecentThoughts.Received
        thoughts={thoughts.filter((thought) => thought.isBurned || thought.isViewed)}
      />
    </Container>
  );
};

RecentThoughts.NotReceived = NotReceived;
RecentThoughts.Received = Received;

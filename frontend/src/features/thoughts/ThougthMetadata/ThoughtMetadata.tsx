import React, { useEffect, useState } from 'react';
import { Alert, Button, Code, Container, Divider, Grid, Paper, Text, Title } from '@mantine/core';
import { Link, useLocation, useParams } from 'react-router-dom';
import { Bolt, MessageCircle2 } from 'tabler-icons-react';
import { isObjectEmpty } from 'utils/is-empty';
import { ThoughtMetadataModel } from 'features/thoughts/thought.model';

// TODO: Check if such metadata exists
export function ThoughtMetadata() {
  const { state: routerState } = useLocation();
  const { metadataKey } = useParams();
  const [isAdviceAlertVisible, setIsAdviceAlertVisible] = useState(true);
  const [thoughtMetadata, SetThoughtMetadata] = useState<ThoughtMetadataModel>();

  useEffect(() => {
    console.log(routerState);
    if (isObjectEmpty(routerState)) {
      console.log('state is empty');
    } else {
      SetThoughtMetadata({ ...(routerState as ThoughtMetadataModel) });
      window.history.replaceState({}, document.title);

      console.log('not empty');
      console.log(thoughtMetadata);
    }
  }, [routerState]);

  return (
    <Container my="xl">
      <Title order={1} my="lg">
        Thought metadata 💭
      </Title>
      <Paper shadow="xs" p="md" my="md" withBorder>
        <Text size="xl" weight="500">
          Share the link:
        </Text>
        <Code color="yellow" my="xs" style={{ fontSize: '20px' }}>
          http://localhost:9000/thought/{metadataKey}
        </Code>
        <Text color="gray" size="sm">
          Requires a passphrase.
        </Text>
      </Paper>
      <Paper shadow="xs" p="md" my="md" withBorder>
        <Text size="xl" weight="500">
          The passphrase:
        </Text>
        <Code color="orange" my="xs" style={{ fontSize: '20px' }}>
          http://localhost:9000/thought/{metadataKey}
        </Code>
      </Paper>
      <Grid align="center" mx={0} my="lg">
        <Text size="xl" weight="500">
          Expires in 7 days.
        </Text>
        <Text color="dimmed" pl="sm">
          (2022-06-15@06:11:33 UTC)
        </Text>
      </Grid>
      <Button<typeof Link>
        to={`/thought/${metadataKey}/burn`}
        leftIcon={<Bolt size={24} />}
        variant="light"
        color="orange"
        my="lg"
        fullWidth
        component={Link}>
        Burn this thought
      </Button>
      {isAdviceAlertVisible && (
        <>
          <Divider my="md" />
          <Alert
            color="red"
            withCloseButton
            title="Advice"
            onClose={() => {
              setIsAdviceAlertVisible(false);
            }}
            closeButtonLabel="Close advice">
            Burning a thought will delete it before it has been read (click to confirm).
          </Alert>
          <Divider my="md" />
        </>
      )}
      <Button<typeof Link>
        to="/"
        leftIcon={<MessageCircle2 size={24} />}
        variant="light"
        my="lg"
        fullWidth
        component={Link}>
        Create another thought
      </Button>
    </Container>
  );
}

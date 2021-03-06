/* eslint-disable no-nested-ternary */
import React, { useEffect, useState } from 'react';
import {
  Alert,
  Button,
  Code,
  Container,
  Divider,
  Grid,
  LoadingOverlay,
  Paper,
  Text,
  Title
} from '@mantine/core';
import { Link, useLocation, useParams } from 'react-router-dom';
import { Bolt, MessageCircle2 } from 'tabler-icons-react';
import { ThoughtMetadataModel } from 'features/thoughts/thought.model';
import * as Thought from 'features/thoughts';
import { DateUnit, getDateDiffIn, prettyDiffDate, isObjectEmpty } from 'utils';
import { api } from 'services/http';

export function ThoughtMetadata() {
  const { state: routerState } = useLocation();
  const { metadataKey } = useParams();
  const [isAdviceAlertVisible, setIsAdviceAlertVisible] = useState(true);
  const [thoughtMetadata, setThoughtMetadata] = useState<ThoughtMetadataModel>();
  const [thoughtLink, setThoughtLink] = useState<string | undefined>(undefined);
  const [isError, setIsError] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    if (isObjectEmpty(routerState)) {
      setIsLoading(true);
      api
        .get<ThoughtMetadataModel>({ url: `metadata/${metadataKey}` })
        .then((response) => {
          setThoughtMetadata(response);
        })
        .catch(() => {
          setIsError(true);
        })
        .finally(() => {
          setIsLoading(false);
        });
    } else {
      setThoughtMetadata({ ...(routerState as ThoughtMetadataModel) });
      setThoughtLink(
        `http://localhost:3000/thought/${(routerState as ThoughtMetadataModel)?.thoughtKey}`
      );
      window.history.replaceState({}, document.title);
    }
  }, [metadataKey, routerState]);

  if (isError) {
    return (
      <Container>
        <Alert color="red" title="Oops...">
          Looks like such thought does not exist.
        </Alert>
      </Container>
    );
  }

  return (
    <Container my="xl">
      <LoadingOverlay visible={isLoading} />

      <Title order={1} my="lg">
        Thought metadata ????
      </Title>

      {thoughtMetadata?.thoughtKey && <Thought.ThoughtLink link={thoughtLink} />}

      <Paper shadow="xs" p="md" my="md" withBorder>
        <Text size="xl" weight="500">
          Thought ({thoughtMetadata?.abbreviatedThoughtKey})
        </Text>
        <Code color="orange" my="xs" style={{ fontSize: '20px' }}>
          This message is encrypted with your passphrase.
        </Code>
      </Paper>

      {thoughtMetadata?.isViewed ? (
        <Grid align="center" mx={0} my="lg">
          <Text size="xl" weight="500">
            Viewed {getDateDiffIn(DateUnit.minute, thoughtMetadata?.viewedDate)} minutes ago.
          </Text>
          <Text color="dimmed" pl="sm">
            {thoughtMetadata?.viewedDate}
          </Text>
        </Grid>
      ) : thoughtMetadata?.isBurned ? (
        <Code color="blue" style={{ fontSize: '20px' }}>
          Burned {prettyDiffDate(thoughtMetadata.burnedDate)}.{' '}
          {new Date(thoughtMetadata?.burnedDate ?? '').toLocaleString()}
        </Code>
      ) : (
        <>
          <Grid align="center" mx={0} my="lg">
            <Text size="xl" weight="500">
              Expires in {getDateDiffIn(DateUnit.hour, thoughtMetadata?.lifetime)} hours.
            </Text>
            <Text color="dimmed" pl="sm">
              {thoughtMetadata?.lifetime}
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
                color="violet"
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

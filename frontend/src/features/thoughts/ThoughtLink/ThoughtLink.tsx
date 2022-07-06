import React from 'react';
import { ActionIcon, Alert, Code, Grid, Paper, Text } from '@mantine/core';
import { showNotification } from '@mantine/notifications';
import { selectElement, copyTextToClipboard } from 'utils';

interface ThoughtLinkProps {
  link?: string;
}

export function ThoughtLink({ link }: ThoughtLinkProps) {
  const copyToClipboard = () => {
    copyTextToClipboard(link)
      .then(() => {
        showNotification({
          title: 'You did great ✅',
          message: 'The thought link was copied to clipboard',
          color: 'green'
        });
      })
      .catch(() => {
        showNotification({
          title: 'Something went wrong. ❌',
          message: 'The thought link was not copied to clipboard.',
          color: 'red'
        });
      });
  };

  return (
    <Paper shadow="xs" p="md" my="md" withBorder>
      <Text size="xl" weight="500">
        Share the link:
      </Text>
      <Grid justify="space-between" align="center">
        <Grid.Col span={11}>
          <Code
            color="yellow"
            my="xs"
            style={{ fontSize: '20px', cursor: 'copy' }}
            onClick={(e) => selectElement(e.target as HTMLElement)}>
            {link}
          </Code>
        </Grid.Col>
        <Grid.Col span={1}>
          <ActionIcon
            aria-label="Copy thought link to clipboard"
            variant="outline"
            onClick={() => copyToClipboard()}
            size="lg"
            style={{ marginLeft: 'auto' }}>
            💾
          </ActionIcon>
        </Grid.Col>
      </Grid>
      <Text color="gray" size="sm">
        Requires a passphrase.
      </Text>
      <Alert color="orange" title="Warning" my="sm">
        It will disappear forever after you leave the page
      </Alert>
    </Paper>
  );
}

ThoughtLink.defaultProps = {
  link: undefined
};

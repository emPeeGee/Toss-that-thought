import React from 'react';
import { Anchor, Button, Group, Mark, Text, Timeline, Title } from '@mantine/core';
import { Link } from 'react-router-dom';
import { ThoughtMetadataModel } from 'features/thoughts/thought.model';

export type NotReceivedComponent = React.FC<{ thoughts: ThoughtMetadataModel[] }>;

interface ReceivedProps {
  thoughts: ThoughtMetadataModel[];
}

// eslint-disable-next-line react/function-component-definition
export const NotReceived: NotReceivedComponent = ({ thoughts }: ReceivedProps) => {
  return (
    <>
      <Title order={2} my="lg">
        Not Received thoughts 💨
      </Title>

      <Timeline active={1} bulletSize={24} lineWidth={2} my="md">
        {thoughts.map((thought) => (
          <Timeline.Item active lineActive key={thought.metadataKey}>
            <Anchor component={Link} to={`/metadata/${thought.metadataKey}`} weight={500}>
              {thought.abbreviatedThoughtKey}
            </Anchor>

            <Group position="apart">
              <Group direction="column" align="left" spacing={0}>
                <Text size="sm">
                  Created:{' '}
                  <Mark color="primary" p={1}>
                    {new Date(thought?.createdDate ?? '').toLocaleString()}
                  </Mark>
                </Text>

                <Text size="sm">
                  Active till:
                  <Mark color="green" p={1}>
                    {new Date(thought.lifetime ?? '').toLocaleString()}
                  </Mark>
                </Text>
              </Group>

              <Button<typeof Link>
                to={`/thought/${thought.metadataKey}/burn`}
                component={Link}
                aria-label={`Go to page burn thought ${thought.abbreviatedThoughtKey}`}
                variant="outline"
                color="orange">
                Burn
              </Button>
            </Group>
          </Timeline.Item>
        ))}
      </Timeline>
    </>
  );
};

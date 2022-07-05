import { ThoughtMetadataModel } from 'features/thoughts/thought.model';
import { Anchor, Group, Mark, Text, Timeline, Title } from '@mantine/core';
import { Link } from 'react-router-dom';
import React from 'react';

export type ReceivedComponent = React.FunctionComponent<{ thoughts: ThoughtMetadataModel[] }>;

interface ReceivedProps {
  thoughts: ThoughtMetadataModel[];
}

// eslint-disable-next-line react/function-component-definition
export const Received: ReceivedComponent = ({ thoughts }: ReceivedProps) => {
  return (
    <>
      <Title order={2} my="lg">
        Received thoughts ✅
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

                {thought.isViewed && (
                  <Text size="sm">
                    Viewed:
                    <Mark color="orange" p={1}>
                      {new Date(thought.viewedDate ?? '').toLocaleString()}
                    </Mark>
                  </Text>
                )}

                {thought.isBurned && (
                  <Text size="sm">
                    Burned:
                    <Mark color="orange" p={1}>
                      {new Date(thought.burnedDate ?? '').toLocaleString()}
                    </Mark>
                  </Text>
                )}

                <Text size="sm">
                  Active till:
                  <Mark color="green" p={1}>
                    {new Date(thought.lifetime ?? '').toLocaleString()}
                  </Mark>
                </Text>
              </Group>
            </Group>
          </Timeline.Item>
        ))}
      </Timeline>
    </>
  );
};

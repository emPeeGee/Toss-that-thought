import React, { useEffect, useMemo, useState } from 'react';
import { Alert, Button, Container, Divider, LoadingOverlay, Textarea } from '@mantine/core';
import { ArrowForwardUp } from 'tabler-icons-react';
import { Link, useParams } from 'react-router-dom';
import { api } from 'services/http';
import { ThoughtPassphraseInfo, ThoughtViewContext } from 'features/thoughts';
import { GetThoughtForm, GetThoughtFormComponent } from './GetThoughForm/GetThoughtForm';

type ThoughtViewComponent = React.FunctionComponent & {
  GetThoughtForm: GetThoughtFormComponent;
};

// TODO: A bug or a feature? When I burn a thought, but have a page with view thought,
// if I enter password, it says that password is not correct, but on refresh,
// it says it does not exists, the same for burn view

// eslint-disable-next-line react/function-component-definition
export const ThoughtView: ThoughtViewComponent = () => {
  const { thoughtKey } = useParams();
  const [isCarefulAlertVisible, setIsCarefulAlertVisible] = useState(true);
  const [isError, setIsError] = useState(false);
  const [isPassphrasePhasePassed, setIsPassphrasePhasePassed] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [thought, setThought] = useState('');
  const [canSkipPassphrase, setCanSkipPassphrase] = useState(false);
  const value = useMemo(
    () => ({
      isLoading,
      setIsLoading,
      isPassphrasePhasePassed,
      setIsPassphrasePhasePassed,
      setThought
    }),
    [isLoading, canSkipPassphrase]
  );

  useEffect(() => {
    setIsLoading(true);
    api
      .get<ThoughtPassphraseInfo>({ url: `thought/${thoughtKey}` })
      .then((response) => {
        setCanSkipPassphrase(response.canPassphraseBeSkipped);
      })
      .catch((err) => {
        console.log(err);
        setIsError(true);
      })
      .finally(() => {
        setIsLoading(false);
      });
  }, []);

  if (isError) {
    return (
      <Container size="md" my="xl">
        <Alert color="red" title="Hmm">
          It either never existed or has already been viewed.
        </Alert>
      </Container>
    );
  }

  return (
    <ThoughtViewContext.Provider value={value}>
      <Container size="md" my="xl">
        <LoadingOverlay visible={isLoading} />
        {isPassphrasePhasePassed ? (
          <>
            <Textarea
              disabled
              label="This thought is for you:"
              variant="filled"
              size="xl"
              my="xl"
              value={thought}
            />
            <Button<typeof Link>
              to="/"
              component={Link}
              fullWidth
              my="lg"
              leftIcon={<ArrowForwardUp size={24} />}>
              Reply with another thought
            </Button>
          </>
        ) : (
          <ThoughtView.GetThoughtForm
            thoughtKey={thoughtKey}
            canSkipPassphrase={canSkipPassphrase}
          />
        )}

        {isCarefulAlertVisible && (
          <>
            <Divider my="md" />
            <Alert
              withCloseButton
              color="orange"
              title="Careful"
              closeButtonLabel="Close advice"
              onClose={() => {
                setIsCarefulAlertVisible(false);
              }}>
              We will only show it once
            </Alert>
          </>
        )}
      </Container>
    </ThoughtViewContext.Provider>
  );
};

ThoughtView.GetThoughtForm = GetThoughtForm;

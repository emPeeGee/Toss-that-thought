import React, { useContext, useState } from 'react';
import { useForm } from 'react-hook-form';
import { ThoughtPassphraseRequest, ThoughtResponse } from 'features/thoughts/thought.model';
import { api } from 'services/http';
import { Alert, Button, PasswordInput, Text } from '@mantine/core';
import { Eye, Lock } from 'tabler-icons-react';
import { ThoughtViewContext, ThoughtViewContextModel } from 'features/thoughts/';

interface GetThoughtFormProps {
  thoughtKey?: string;
  canSkipPassphrase: boolean;
}

export type GetThoughtFormComponent = React.FunctionComponent<GetThoughtFormProps>;

// eslint-disable-next-line react/function-component-definition
export const GetThoughtForm: GetThoughtFormComponent = ({
  thoughtKey,
  canSkipPassphrase
}: GetThoughtFormProps) => {
  const thoughtViewContext = useContext<ThoughtViewContextModel | null>(ThoughtViewContext);
  const [isPassphraseCorrect, setIsPassphraseCorrect] = useState(true);
  const { register, handleSubmit } = useForm<ThoughtPassphraseRequest>({
    mode: 'onChange'
  });

  const viewThought = (data: ThoughtPassphraseRequest) => {
    thoughtViewContext?.setIsLoading(true);

    api
      .post<ThoughtPassphraseRequest, ThoughtResponse>({ url: `thought/${thoughtKey}`, body: data })
      .then((response) => {
        thoughtViewContext?.setThought(response.thought);
        thoughtViewContext?.setIsPassphrasePhasePassed(true);
      })
      .catch(() => {
        setIsPassphraseCorrect(false);
      })
      .finally(() => thoughtViewContext?.setIsLoading(false));
  };

  return (
    <>
      {!isPassphraseCorrect && (
        <Alert title="Oops..." color="red" my="lg">
          Double check that passphrase
        </Alert>
      )}

      <form onSubmit={handleSubmit(viewThought)}>
        {!canSkipPassphrase ? (
          <>
            <Text size="xl">This thought requires a passphrase:</Text>
            <PasswordInput
              {...register('passphrase', { required: false, value: '' })}
              my="md"
              placeholder="Enter passphrase here"
              toggleTabIndex={0}
              icon={<Lock size={16} />}
            />
          </>
        ) : (
          <Text size="xl">Click the button to continue 👇</Text>
        )}
        <Button fullWidth my="lg" variant="light" leftIcon={<Eye size={24} />} type="submit">
          View thought
        </Button>
      </form>
    </>
  );
};

GetThoughtForm.defaultProps = {
  thoughtKey: undefined
};

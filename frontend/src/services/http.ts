// To make a config file
const testUrl = 'http://localhost:9000/api/';

interface GetRequest {
  url?: string;
}

interface PostRequest<T> {
  url?: string;
  body?: T;
}

async function handleErrors<T>(response: Response): Promise<T> {
  if (!response.ok) {
    // eslint-disable-next-line prefer-promise-reject-errors
    return Promise.reject((await response.json()) as T);
  }

  return (await response.json()) as T;
}

// TODO: To check if this approach is good
export const api = {
  get: <K>({ url }: GetRequest) =>
    fetch(`${testUrl}${url}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then((response) => handleErrors<K>(response)),

  post: <T, K>({ url, body }: PostRequest<T>) =>
    fetch(`${testUrl}${url}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(body)
    }).then((response) => handleErrors<K>(response))
};

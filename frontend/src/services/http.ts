// To make a config file
const testUrl = 'http://localhost:9000/api/';
const authUrl = 'http://localhost:9000/auth/';

interface GetRequest {
  url: string;
  token?: string | null;
}

interface PostRequest<T> {
  url: string;
  body?: T;
  auth?: boolean;
  token?: string | null;
}

async function handleErrors<T>(response: Response): Promise<T> {
  if (!response.ok) {
    // eslint-disable-next-line prefer-promise-reject-errors
    return Promise.reject((await response.json()) as T);
  }

  return (await response.json()) as T;
}

export const api = {
  get: <K>({ url, token }: GetRequest) =>
    fetch(`${testUrl}${url}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `${token ? `Bearer ${token}` : ''}`
      }
    }).then((response) => handleErrors<K>(response)),

  post: <T, K>({ url, body, token, auth = false }: PostRequest<T>) =>
    fetch(`${auth ? authUrl : testUrl}${url}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `${token ? `Bearer ${token}` : ''}`
      },
      body: JSON.stringify(body)
    }).then((response) => handleErrors<K>(response))
};

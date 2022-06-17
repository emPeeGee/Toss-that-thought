// To make a config file
const testUrl = 'http://localhost:9000/api/';

interface PostRequest<T> {
  url?: string;
  body?: T;
}

// TODO: To check if this approach is good
export const api = {
  // get: () =>
  //   fetch('http://localhost:9000/api/create', {
  //     method: 'POST',
  //     headers: {
  //       'Content-Type': 'application/json'
  //     },
  //     body: JSON.stringify(data)
  //   }).then((value) => value.json()),

  post: <T, K>({ url, body }: PostRequest<T>) =>
    fetch(`${testUrl}${url}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(body)
    }).then(async (response) => (await response.json()) as K)
};

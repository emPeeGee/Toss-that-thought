/* eslint-disable @typescript-eslint/return-await */

export async function copyTextToClipboard(text?: string): Promise<string | void> {
  if (!text) {
    return Promise.reject();
  }

  if ('clipboard' in navigator) {
    return await navigator.clipboard.writeText(text);
  }

  document.execCommand('copy', true, text);
  return Promise.resolve();
}

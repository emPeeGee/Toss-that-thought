export function selectElement(target: HTMLElement): void {
  const range = document.createRange();
  const selection = window.getSelection();
  range.selectNodeContents(target);

  selection?.removeAllRanges();
  selection?.addRange(range);
}

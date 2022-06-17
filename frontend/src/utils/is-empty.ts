export const isObjectEmpty = (target: any) => {
  if (!target) {
    return true; // If the target is null or undefined, return true, meaning object is empty
  }

  return Object.keys(target).length === 0;
};

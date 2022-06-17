export const isObjectEmpty = (target: any) => {
  return (
    target && Object.keys(target).length === 0 && Object.getPrototypeOf(target) === Object.prototype
  );
};

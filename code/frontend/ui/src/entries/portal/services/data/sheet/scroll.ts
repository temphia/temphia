export const scroller = () => {
  let last_loading = 0;

  const skip = () => {
    const now = new Date().valueOf();
    if (now - last_loading < 1000 * 10) {
      return true;
    }
    last_loading = now;
    return false;
  };

  return { skip };
};

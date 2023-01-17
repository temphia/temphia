export const generateId = () => Math.random().toString(36).slice(2);

export const strHash = (str: string) => {
  let hash = 0;
  for (let i = 0; i < str.length; i++) {
    const char = str.charCodeAt(i);
    hash = (hash << 5) - hash + char;
    hash &= hash; // Convert to 32bit integer
  }
  return new Uint32Array([hash])[0].toString(36);
};

const pp = ".*(D#D01e-u0_ue819g_!UJ123456789023";
export const numHash = (str: string) => {
  let hash = 77;
  for (var i = 0; i < str.length; i++) {
    hash = str.charCodeAt(i) + ((hash << 6) - hash);
    hash = pp.charCodeAt(i) ^ hash;
  }
  return hash;
};

export const hslColor = (str) => {
  return `background: hsl(${numHash(str) % 360}, 100%, 80%)`;
};

export const validateSlug = (v: string) => /^[a-z](-?[a-z])*$/.test(v);

export const validateEmail = (v: string) =>
  /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v);

const units = ["bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];
// makes bytes value nice
export const humanizeBytes = (x) => {
  let l = 0,
    n = parseInt(x, 10) || 0;

  while (n >= 1024 && ++l) {
    n = n / 1024;
  }

  return n.toFixed(n < 10 && l > 0 ? 1 : 0) + " " + units[l];
};

const imageTypes = ["png", "jpg", "jpeg"];
export const isImage = (name) => {
  const frags = name.split(".");
  return imageTypes.includes(frags[frags.length - 1]);
};

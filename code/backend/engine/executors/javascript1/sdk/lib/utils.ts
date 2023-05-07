const no_found = "upper: no more rows in this result set";
const already_exists = "duplicate key value violates";

interface Response {
  payload: {
    ok: boolean;
    data?: any;
    message?: string;
  };
}

export const utils = {
  is_db_not_found: (err: string) => {
    return err.indexOf(no_found) !== -1;
  },

  is_db_already_exists: (err: string) => {
    return err.indexOf(already_exists) !== -1;
  },

  ab2str: (buf: ArrayBuffer) => {
    return String.fromCharCode.apply(null, new Uint16Array(buf));
  },
  str2ab: (str: string) => {
    var buf = new ArrayBuffer(str.length * 2);
    var bufView = new Uint16Array(buf);
    for (var i = 0, strLen = str.length; i < strLen; i++) {
      bufView[i] = str.charCodeAt(i);
    }
    return buf;
  },
  is_arraybuffer(value: any): boolean {
    return (
      typeof ArrayBuffer === "function" &&
      (value instanceof ArrayBuffer ||
        toString.call(value) === "[object ArrayBuffer]")
    );
  },
  generate_str_id: () => Math.random().toString(36).slice(2),
};

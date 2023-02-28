import * as ct from "./column";

const istclass = [
  ct.CtypeShortText,
  ct.CtypePhone,
  ct.CtypeNumber,
  ct.CtypeSelect,
];

const text_cols = ["name", "title"];

export const generate_column_order = (columns: {
  [_: string]: object;
}): string[] => {
  const doneCols = {};
  const orderedColumns = [];

  // then first class colums
  istclass.forEach((cType) => {
    Object.values(columns).forEach((val) => {
      if (doneCols[val["slug"]]) {
        return;
      }

      if (val["ref_type"]) {
        return;
      }

      if (val["ctype"] !== cType) {
        return;
      }

      orderedColumns.push(val["slug"]);
      doneCols[val["slug"]] = true;
    });
  });

  // then remaining columns expect ref types
  Object.values(columns).forEach((val) => {
    if (istclass.includes(val["ctype"])) {
      return;
    }
    if (doneCols[val["slug"]]) {
      return;
    }

    if (val["ref_type"]) {
      return;
    }

    orderedColumns.push(val["slug"]);
    doneCols[val["slug"]] = true;
  });

  // atlast ref types
  Object.values(columns).forEach((val) => {
    if (doneCols[val["slug"]]) {
      return;
    }

    if (!val["ref_type"]) {
      return;
    }

    orderedColumns.push(val["slug"]);
    doneCols[val["slug"]] = true;
  });

  return orderedColumns;
};

export interface Order {
  name?: string;
  image?: string;
  description?: string;
  user?: string;
  tag?: string;
  tags?: string[];
  other: string[];
}

export const calculate_card_order = (
  columns: { [_: string]: object },
  view: object
) => {
  const sk: Order = {
    other: [],
  };

  const done = {};
  for (let index = 0; index < text_cols.length; index++) {
    const element = text_cols[index];
    const col = columns[element];

    if (col && col["ctype"] === ct.CtypeShortText) {
      sk.name = element;
      done[element] = true;
    }
    break;
  }

  if (columns["image"]) {
    sk.image = "image";
    done["image"] = true;
  }

  if (columns["description"]) {
    sk.description = "description";
    done["description"] = true;
  }

  if (columns["user"]) {
    sk.user = "user";
    done["user"] = true;
  }

  if (columns["tag"]) {
    sk.tag = "tag";
    done["tag"] = true;
  }

  const colkeys = Object.keys(columns);
  for (let index = 0; index < colkeys.length; index++) {
    const element = colkeys[index];
    if (done[element]) {
      continue;
    }

    sk.other.push(element);
  }

  return sk;
};

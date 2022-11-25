import {
  CtypeShortText,
  CtypePhone,
  CtypeSelect,
  CtypeRFormula,
  CtypeFile,
  CtypeMultiFile,
  CtypeCheckBox,
  CtypeCurrency,
  CtypeNumber,
  CtypeLocation,
  CtypeDateTime,
  CtypeMultSelect,
  CtypeLongText,
  CtypeSingleUser,
  CtypeMultiUser,
  CtypeEmail,
  CtypeJSON,
  CtypeRangeNumber,
  CtypeColor,
} from "../core/fields/field";

export interface Order {
  name?: string;
  image?: string;
  description?: string;
  user?: string;
  tag?: string;
  tags?: string[];
  other: string[];
}

const text_cols = ["name", "title"];

export const calculate_order = (
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

    if (col && col["ctype"] === CtypeShortText) {
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

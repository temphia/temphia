import { CtypeIcons } from "../../../../data/tableui/core/fields/field";
import type { Column, ColumnRef, State, Table } from "./builder";

export const Validate = (group: State): string => {
  if (group.slug.length > 20) {
    return "error group slug is bigger than 20";
  }

  for (var i = 0; i < group.tables.length; i++) {
    const tbl = group.tables[i];

    if (tbl.slug.length > 20) {
      return `error table [${tbl.name}] slug is bigger than 20 : ${tbl.slug}`;
    }

    for (var j = 0; j < tbl.columns.length; j++) {
      const col = tbl.columns[j];

      if (col.slug.length > 20) {
        return `error column [${tbl.name} / ${col.name}] slug is bigger than 20 : ${tbl.slug}`;
      }

      if (!CtypeIcons[col.ctype]) {
        `error unknown column type [${tbl.name} / ${col.name}] : ${col.ctype}`;
      }

      col.ctype;
    }
  }

  return "";
};

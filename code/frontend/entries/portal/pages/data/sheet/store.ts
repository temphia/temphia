import { Writable, writable } from "svelte/store";
import {
  Sheet,
  SheetCell,
  SheetColTypeBoolean,
  SheetColTypeDate,
  SheetColTypeFile,
  SheetColTypeLongText,
  SheetColTypeText,
  SheetColumn,
  SheetRow,
} from "./sheets";

interface State {
  columns: SheetColumn[];
  rows: SheetRow[];
  cells: { [_: string]: { [_: string]: SheetCell } } | any;
  sheets: Sheet[];
}

export const CreateStore = () => {
  const state: Writable<State> = writable({
    cells: {
      "1": { "1": { value: "Ram" }, "2": { value: "This is some notes" } },
      "2": { "1": { value: "Hari" }, "2": { value: "Something something" } },
      "3": { "1": { value: "bin" }, "2": { value: "the mr bean ..." } },
    },

    columns: [
      {
        ctype: SheetColTypeText,
        id: 1,
        name: "name",
        sheet_id: 1,
      },

      {
        ctype: SheetColTypeLongText,
        id: 2,
        name: "notes",
        sheet_id: 1,
      },
      {
        ctype: SheetColTypeBoolean,
        id: 3,
        name: "Done",
        sheet_id: 1,
      },
      {
        ctype: SheetColTypeDate,
        id: 4,
        name: "Created At",
        sheet_id: 1,
      },
    ],
    rows: [
      {
        id: 1,
        sheet_id: 1,
      },

      {
        id: 2,
        sheet_id: 1,
      },
      {
        id: 3,
        sheet_id: 1,
      },
    ],
    sheets: [
      { id: 1, name: "example 1" },
      { id: 2, name: "example 2" },
    ],
  });

  const add_column = (name: string, ctype: string, opts: object) => {
    state.update((old) => {
      const colIndex = old.columns.reduce(
        (prev, curr) => (curr.id > prev ? curr.id : prev),
        0
      );

      return {
        ...old,
        columns: [
          ...old.columns,
          { name, ctype, exta_options: opts, id: colIndex + 1, sheet_id: 1 },
        ],
      };
    });
  };

  return { state, add_column };
};

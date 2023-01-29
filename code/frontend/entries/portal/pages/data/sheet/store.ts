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
        __id: 1,
        name: "name",
        sheetid: 1,
      },

      {
        ctype: SheetColTypeLongText,
        __id: 2,
        name: "notes",
        sheetid: 1,
      },
      {
        ctype: SheetColTypeBoolean,
        __id: 3,
        name: "Done",
        sheetid: 1,
      },
      {
        ctype: SheetColTypeDate,
        __id: 4,
        name: "Created At",
        sheetid: 1,
      },
    ],
    rows: [
      {
        __id: 1,
        sheetid: 1,
      },

      {
        __id: 2,
        sheetid: 1,
      },
      {
        __id: 3,
        sheetid: 1,
      },
    ],
    sheets: [
      { __id: 1, name: "example 1" },
      { __id: 2, name: "example 2" },
    ],
  });

  const add_column = (name: string, ctype: string, opts: object) => {
    state.update((old) => {
      const colIndex = old.columns.reduce(
        (prev, curr) => (curr.__id > prev ? curr.__id : prev),
        0
      );

      return {
        ...old,
        columns: [
          ...old.columns,
          { name, ctype, exta_options: opts, __id: colIndex + 1, sheetid: 1 },
        ],
      };
    });
  };

  return { state, add_column };
};

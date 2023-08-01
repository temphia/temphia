export const DataModTypeSheetInsert = "sheet_insert";
export const DataModTypeSheetUpdate = "sheet_update";
export const DataModTypeSheetDelete = "sheet_delete";

export type ModType =
  | typeof DataModTypeSheetInsert
  | typeof DataModTypeSheetUpdate
  | typeof DataModTypeSheetDelete;

export interface DataSheetMod {
  group?: string;
  sheet_id: number;
  rows: number[];
  mod_type: ModType;
  data: any;
}

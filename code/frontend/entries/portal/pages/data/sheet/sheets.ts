export interface Sheet {
  id: number;
  name: string;
}

export interface SheetColumn {
  id: number;
  name: string;
  ctype: "TEXT" | "LONGTEXT" | "INTEGER" | "FLOAT" | "DATE" | "BOOL" | "RATINGS" | "LOCATION" | "FILE";
  sheet_id: number;
  exta_options: {};
  color: string;
}

export interface SheetRow {
  id: number;
  sheet_id: string;
  color: string;
}

export interface SheetCell {
  id: number;
  value: string;
  num_value: number;
  sheet_id: number;
  row_id: number;
}

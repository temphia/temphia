export interface FilterItem {
  column: string;
  cond: string;
  value: any;
}

export interface Column {
  slug: string;
  name: string;
  ctype: string;
  options: string[];
  description: string;
  pattern: string;
  strict_pattern: boolean;
  ref_id: string;
  ref_type: string;
  ref_copy: string;
  ref_target: string;
  ref_object: string;
}

export interface DataWidget {
  id: number;
  name: string;
  type: string;
  sub_type: string;
  plug_id: string;
  agent_id: string;
  icon: string;
  payload: string;
}

export type DataModType =
  | typeof DataModTypeInsert
  | typeof DataModTypeUpdate
  | typeof DataModTypeDelete
  | typeof DataModTypeComment;

export const DataModTypeInsert = "insert";
export const DataModTypeUpdate = "update";
export const DataModTypeDelete = "delete";
export const DataModTypeComment = "comment";

export interface DataModification {
  group?: string;
  table: string;
  rows: number[];
  mod_type: DataModType;
  data: any;
}

export interface TableExecData {
  invoker_type: "data_table";
  table_id: string;
  data_group: string;
  table_slug: string;
  source: string

  cells: { [_: string]: { [_: string]: any } };
  rows: number[];
  columns?: Column[];
}

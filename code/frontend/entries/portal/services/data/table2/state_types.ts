import type { Column } from "./table_types";
import type { FilterItem } from "./table_types";

export interface ViewData {
  name?: string;
  filter_conds: FilterItem[];
  count: number;
  selects: string[];
  main_column: string;
  search_term: string;
}

export interface NavData {
  loading: boolean;
  lastTry: Date;

  loading_error: string;
  last_page: boolean;
  active_page: number;

  active_view: ViewData;
}

export interface DirtyData {
  rowid: number;
  data: object;
}

export const defaultViewData = () => ({
  count: 20,
  filter_conds: [],
  main_column: "",
  search_term: "",
  selects: [],
  page: 0,
});

export interface DataState {
  indexed_column: { [_: string]: Column };
  column_order: string[];

  rows: number[];
  indexed_rows: { [_: string]: object };

  ref_rows_cache: { [_: string]: { [_: string]: object } }; // <column <row_id|column_id, object>>

  sparse_rows: number[];
  remote_dirty: { [_: string]: true };
  views: object[];
}

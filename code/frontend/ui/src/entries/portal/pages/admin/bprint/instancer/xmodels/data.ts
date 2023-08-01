// syncme => xbprint/data.go

export interface NewTableGroup {
  name: string;
  slug: string;
  description: string;
  tables: NewTable[];
  exec_order: string[];
}

export interface NewTable {
  name: string;
  slug: string;
  description: string;
  icon: string;
  main_column?: string;
  activity_type: string;
  sync_type: string;
  columns: object[];
  indexes: object[];
  unique_indexes: object[];
  fts_index?: object;
  column_refs: object[];
  deleted_at: boolean;
  views: object[];
  seed_data?: object;
}

export interface DataGroupRequest {
  dyndb_source: string;
  group_name: string;
  group_slug: string;
  cabinet_source: string;
  cabinet_folder: string;
  seed_source: string; // autogen | data
}

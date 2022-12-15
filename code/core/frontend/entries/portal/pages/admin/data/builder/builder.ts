import { get, writable, Writable } from "svelte/store";

export interface Table {
  name: string;
  slug: string;
  description: string;
  icon: string;
  main_column: string;
  columns: Column[];
  column_refs: ColumnRef[];
}

export interface Column {
  name: string;
  slug: string;
  ctype: string;
  description: string;
  icon: string;
  options: string[];
  not_nullable: boolean;
  pattern: string;
  strict_pattern: boolean;
}

export interface ColumnRef {
  slug: string;
  type: string;
  target: string;
  from_cols: string[];
  to_cols: string[];
}

export interface TableBasic {
  name: string;
  slug: string;
  description: string;
}

export interface State {
  name: string;
  slug: string;
  description: string;
  tables: Table[];
}

export class Builder {
  state: Writable<State>;

  static new(name: string, slug: string, description?: string): Builder {
    const b = new Builder();

    b.state = writable({
      name: name,
      slug: slug,
      description: description || "",
      tables: [],
    });

    return b;
  }

  static from_batch(data: State): Builder {
    const b = new Builder();
    b.state = writable(data);
    return b;
  }

  // table
  add_table = (data: TableBasic) => {
    this.state.update((state) => {
      return {
        ...state,
        tables: [
          ...state.tables,
          {
            ...data,
            column_refs: [],
            columns: [],
            icon: "",
            main_column: "",
          },
        ],
      };
    });
  };

  edit_table = (table: string, data: TableBasic) => {
    const tidx = this._get_table_index(table);

    this.state.update((state) => {
      const oldtable = state.tables[tidx];

      state.tables[tidx] = { ...oldtable, ...data };

      return {
        ...state,
        tables: [...state.tables],
      };
    });
  };

  delete_table = (table: string) => {
    const tidx = this._get_table_index(table);

    this.state.update((state) => {
      const removed = state.tables.splice(tidx, 1)[0];

      state.tables.forEach((tbl, tidx) => {
        let modified = false;

        tbl.column_refs.forEach((ref, refidx) => {
          if (ref.target !== table) {
            return;
          }
          modified = true;
          tbl.column_refs[refidx] = null;
        });

        if (!modified) {
          return;
        }

        state.tables[tidx].column_refs = state.tables[tidx].column_refs.filter(
          (val) => val !== null
        );
      });

      return {
        ...state,
      };
    });
  };

  // column
  add_column = (table: string, data: Column, ref_data?: ColumnRef) => {
    const tidx = this._get_table_index(table);
    this.state.update((state) => {
      const tbl = state.tables[tidx];
      tbl.columns.push(data);

      return { ...state };
    });
  };

  edit_column = (table: string, column: string, data: Column) => {
    const [tidx, cidx] = this._get_column_index(table, column);

    this.state.update((state) => {
      const tbl = state.tables[tidx];
      const col = tbl.columns[cidx];

      tbl.columns[cidx] = { ...col, ...data };

      return { ...state };
    });
  };

  delete_column = (table: string, column: string) => {
    const [tidx, cidx] = this._get_column_index(table, column);
    this.state.update((state) => {
      const tbl = state.tables[tidx];
      tbl.columns.splice(cidx, 1)[0];

      state.tables.forEach((ctbl) => {
        ctbl.column_refs.forEach((cref, crefIdx) => {
          if (cref.target !== table) {
            return;
          }

          if (!(cref.to_cols || []).includes(column)) {
            return;
          }

          ctbl.column_refs[crefIdx] = null;
        });

        ctbl.column_refs = ctbl.column_refs.filter((val) => val !== null);
      });

      return { ...state };
    });
  };

  // refs
  add_column_ref = (table: string) => {};
  delete_column_ref = (table: string) => {};

  add_index = (table: string) => {};
  delete_index = (table: string) => {};

  add_view = (table: string) => {};
  delete_view = (table: string) => {};

  // private

  _get_table_index = (table: string) => {
    const state = get(this.state);
    let tidx = -1;
    state.tables.forEach((val, idx) => {
      if (val.slug === table) {
        tidx = idx;
      }
    });
    return tidx;
  };

  _get_column_index = (table: string, column: string) => {
    const state = get(this.state);
    let tidx = -1;
    let cidx = -1;
    state.tables.forEach((tval, idx) => {
      if (tval.slug !== table) {
        return;
      }
      tidx = idx;

      tval.columns.forEach((cval, idx) => {
        if (cval.slug !== column) {
          return;
        }
        cidx = idx;
      });
    });

    return [tidx, cidx];
  };
}

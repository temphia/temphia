import { get, writable } from "svelte/store";
import type { Writable} from "svelte/store";
import { DataAPI, FolderTktAPI } from "../../apiv2";
import type { DataSheetAPI } from "../../apiv2/data_sheet";
import type {
  SheetCell,
  SheetColumn,
  SheetRow,
  Sheet,
  SheetWidget,
  SheetExecData,
} from "./sheet_types";
import { formatCells } from "./format";

import type { SockdBuilder } from "../../portal/sockd/builder";

import {
  MESSAGE_SERVER_PUBLISH,
  type Sockd,
  type SockdMessage,
} from "../../portal/sockd";

import { scroller } from "./scroll";
import { formatRefCells } from "./format";
import type { DataSheetMod } from "./sync";

export class SheetGroupService {
  source: string;
  group_slug: string;
  active_sheets: Map<string, SheetService>;

  sheets: Writable<Sheet[]>;

  folder_api: FolderTktAPI;
  data_api: DataAPI;
  data_sheet_api: DataSheetAPI;
  sockd_builder: SockdBuilder;
  sockd_conn: Sockd;

  profile_genrator: (string) => string;

  constructor(
    source: string,
    group: string,
    api: DataAPI,
    sockd_builder: SockdBuilder,
    profile_genrator: (string) => string
  ) {
    this.source = source;
    this.group_slug = group;
    this.active_sheets = new Map();
    this.sheets = writable([]);
    this.data_api = api;
    this.data_sheet_api = api.sheet_api();
    this.profile_genrator = profile_genrator;
    this.sockd_builder = sockd_builder;
  }

  init = async () => {
    const resp = await this.data_sheet_api.list_sheet_group();
    if (!resp.ok) {
      console.log("@err", resp);
      return;
    }

    this.sheets.set(resp.data["sheets"] || []);

    const folder_ticket = resp.data["folder_ticket"] || "";
    this.folder_api = new FolderTktAPI(
      this.data_api.api_base_url,
      folder_ticket
    );

    this.sockd_conn = await this.sockd_builder.build(
      this.data_api.sockd_url(),
      this.sockd_handle
    );
  };

  private sockd_handle = (msg: SockdMessage) => {
    if (msg.type !== MESSAGE_SERVER_PUBLISH || !msg.payload["sheet_id"]) {
      return;
    }

    const payload = msg.payload as DataSheetMod;
    const service = this.active_sheets.get(String(payload.sheet_id));
    service.on_sockd(payload);
  };

  get_sheet_service = async (sheetid: string, gotorow?: number) => {
    let ssvc = this.active_sheets.get(sheetid);
    if (ssvc) {
      return ssvc;
    }

    const sheet = get(this.sheets).filter((v) => v.__id === Number(sheetid));

    ssvc = new SheetService(
      this,
      sheetid,
      sheet[0].name,
      this.profile_genrator
    );

    await ssvc.init(gotorow);

    this.active_sheets.set(sheetid, ssvc);
    return ssvc;
  };

  refetch_sheets = async () => {
    const resp = await this.data_sheet_api.list_sheets();
    if (!resp.ok) {
      return;
    }

    this.sheets.set(resp.data);
  };
}

export interface SheetState {
  columns: SheetColumn[];
  cells: { [_: string]: { [_: string]: SheetCell } };
  rows: SheetRow[];

  ref_columns: SheetColumn[];
  widgets: SheetWidget[];
  loading: boolean;
}

export class SheetService {
  group: SheetGroupService;
  sheetid: string;
  state: Writable<SheetState>;
  api: DataSheetAPI;
  force_render_index: Writable<number>;
  sheet_name: string;

  scroll_skip: () => boolean;
  max_row: number;
  min_row: number;
  loading: boolean;

  profile_genrator: (string) => string;
  scroller?: (rowid: string) => void;
  close_big_modal?: () => void;
  close_small_modal?: () => void;

  constructor(
    group: SheetGroupService,
    sheetid: string,
    sheet_name: string,
    profile_genrator: (string) => string
  ) {
    this.group = group;
    this.sheetid = sheetid;
    this.sheet_name = sheet_name;
    this.profile_genrator = profile_genrator;
    this.api = group.data_sheet_api;

    this.state = writable({
      cells: {},
      columns: [],
      rows: [],
      widgets: [],
      ref_columns: [],
      loading: true,
    });

    this.force_render_index = writable(0);
    this.state.subscribe((state) => console.log("@state", state));

    this.max_row = 0;
    this.min_row = 0;
    this.loading = false;

    this.scroll_skip = scroller().skip;
  }

  init = async (gotorow?: number) => {
    const opts = {};
    if (gotorow) {
      opts["row_cursor_id"] = gotorow;
    }

    const resp = await this.api.load_sheet(this.sheetid, opts);
    if (!resp.ok) {
      return false;
    }

    this.apply(resp.data, {});

    return true;
  };

  apply = (data: any, prevcells: object) => {
    const cell: SheetCell[] = data["cells"] || [];
    const pcells = { ...prevcells, ...formatCells(cell) };

    const rows = Object.keys(pcells)
      .map((v) => ({ __id: Number(v), sheetid: Number(this.sheetid) }))
      .sort((a, b) => {
        return Number(a["__id"]) - Number(b["__id"]);
      });

    this.min_row = (rows[0] || {}).__id || 0;
    this.max_row = (rows[rows.length - 1] || {}).__id || 0;

    const nextstate = {
      columns: data["columns"] || [],
      cells: pcells,
      rows,
      loading: false,
    };

    if (data["widget_apps"]) {
      nextstate["widgets"] = data["widget_apps"];
    }

    if (data["reverse_ref_cols"]) {
      nextstate["ref_columns"] = data["reverse_ref_cols"];
    }

    this.state.update((old) => ({ ...old, ...nextstate }));
  };

  add_sheet = async (name: string, opts: any) => {
    const resp = await this.api.new_sheet({
      name,
      opts,
    });
    if (!resp.ok) {
      return;
    }

    await this.group.refetch_sheets();

    this.force_render_index.update((old) => old + 1);
  };

  remove_sheet = async () => {
    await this.api.delete_sheet(this.sheetid);
    this.group.active_sheets.delete(this.sheetid);
    await this.group.refetch_sheets();
  };

  add_column = async (opts: any) => {
    const resp = await this.api.new_column(this.sheetid, opts);
    if (!resp.ok) {
      return;
    }
    await this.refetch_columns();
  };

  update_column = async (cid: string, data: any) => {
    const resp = await this.api.update_column(this.sheetid, cid, data);
    if (!resp.ok) {
      return;
    }

    await this.refetch_columns();
  };

  remove_column = async (cid: string) => {
    const resp2 = await this.api.delete_column(this.sheetid, cid);
    if (resp2.ok) {
      return;
    }

    await this.refetch_columns();
  };

  add_row_cell = async (data: { [_: number]: { [_: string]: any } }) => {
    const resp = await this.api.new_row_cell(this.sheetid, data);
    if (!resp.ok) {
      return resp;
    }

    const rowid = resp.data[0]["rowid"];
    this.insert_one_row(rowid, resp.data);

    return resp;
  };

  update_row_cell = async (
    rid: string,
    data: { [_: number]: { [_: string]: any } }
  ) => {
    const resp = await this.api.update_row_cell(this.sheetid, rid, data);
    if (!resp.ok) {
      return resp;
    }

    this.update_one_row(Number(rid), resp.data);
    return resp;
  };

  remove_row_cell = async (rid: string) => {
    const resp = await this.api.delete_row_cell(this.sheetid, rid);
    if (resp.ok) {
      const id = Number(rid);
      this.state.update((old) => {
        const newrows = old.rows.filter((v) => v.__id !== id);
        return { ...old, rows: newrows };
      });
    }
  };

  ref_sheet_query = async (opts) => {
    const resp = await this.api.ref_query_sheet(this.sheetid, opts);
    if (!resp.ok) {
      return {};
    }
    return formatRefCells(resp.data);
  };

  private refetch_columns = async () => {
    const resp = await this.api.list_columns(this.sheetid);
    if (!resp.ok) {
      return;
    }

    console.log("@refetch_column", resp.data);

    this.state.update((old) => ({ ...old, columns: resp.data }));
  };

  get_relations = (rid: string, refsheet: string, refcol: string) => {
    return this.api.get_row_relation(this.sheetid, rid, refsheet, refcol);
  };

  search = (opts: object) => {
    return this.api.search(this.sheetid, opts);
  };

  goto_row = (rowid: string) => {
    if (get(this.state).cells[rowid]) {
      if (this.scroller) {
        this.scroller(rowid);
      }
    } else {
      this.init(Number(rowid) - 1);
    }
  };

  scroll_top = async () => {
    this.loading = true;

    if (this.scroll_skip()) {
      return;
    }

    const resp = await this.api.query_sheet(this.sheetid, {
      row_cursor_id: this.min_row,
    });

    if (!resp.ok) {
      this.loading = false;
      return;
    }

    this.apply(resp.data, get(this.state).cells);
    this.loading = false;
  };

  scroll_bottom = async () => {
    this.loading = true;

    if (this.scroll_skip()) {
      return;
    }

    const resp = await this.api.query_sheet(this.sheetid, {
      row_cursor_id: this.max_row,
    });

    if (!resp.ok) {
      this.loading = false;

      return;
    }

    this.apply(resp.data, get(this.state).cells);
    this.loading = false;
  };

  export = () => {
    const sheets = get(this.group.sheets).map((v) => v.__id);
    return this.api.export(sheets);
  };

  get_invoker(widget: SheetWidget, launcher: any) {
    return new SheetInvoker(this, widget, launcher);
  }

  get_exec_data(rows: number[]): SheetExecData {
    const state = get(this.state);
    const cells = {};

    rows.forEach((row) => {
      cells[row] = state.cells[row];
    });

    return {
      data_group: this.group.group_slug,
      invoker_type: "data_sheet",
      sheet_id: Number(this.sheetid),
      sheet_name: this.sheet_name,
      cells,
      columns: state.columns,
      rows: rows.map((r) => ({ __id: r, sheetid: Number(this.sheetid) })),
      source: this.group.source,
    };
  }

  on_sockd = (payload: DataSheetMod) => {
    switch (payload.mod_type) {
      case "sheet_insert":
        this.insert_one_row(payload.rows[0], payload.data);
        break;
      case "sheet_update":
        this.update_one_row(payload.rows[0], payload.data);
        break;
      case "sheet_delete":
        const id = Number(payload.rows[0]);
        this.state.update((old) => {
          const newrows = old.rows.filter((v) => v.__id !== id);
          return { ...old, rows: newrows };
        });
        break;

      default:
        break;
    }
  };

  insert_one_row = (rowid: number, rcells: SheetCell[]) => {
    const data = get(this.state);

    if (data.rows.findIndex((val) => val.__id === rowid) !== -1) {
      console.log("@skipping insert row", rowid);
      return;
    }

    this.state.update((old) => {
      old.rows.push({
        __id: rowid,
        sheetid: Number(this.sheetid),
      });

      rcells.forEach((cell) => {
        const cells = old.cells[cell.rowid] || {};
        cells[cell.colid] = cell;
        old.cells[cell.rowid] = cells;
      });

      return { ...old };
    });
  };

  update_one_row = (rowid: number, rcells: SheetCell[]) => {
    const data = get(this.state);

    if (data.rows.findIndex((val) => val.__id === rowid) === -1) {
      console.log("@skipping update row, cause outof range ?", rowid);
      return;
    }

    this.state.update((old) => {
      rcells.forEach((cell) => {
        const cells = old.cells[cell.rowid] || {};
        cells[cell.colid] = cell;
        old.cells[cell.rowid] = cells;
      });

      return { ...old };
    });
  };
}

export class SheetInvoker {
  widget: SheetWidget;
  service: SheetService;
  instance_id?: string;
  launcher: any;

  constructor(service: SheetService, widget: SheetWidget, launcher: any) {
    this.widget = widget;
    this.service = service;
    this.launcher = launcher;
  }

  init = (instance_id: string) => {
    this.instance_id = instance_id;
  };

  handle = (msg_id: string, data: any) => {
    console.log("@instance_handle", msg_id, data);
  };

  close = () => {};
}

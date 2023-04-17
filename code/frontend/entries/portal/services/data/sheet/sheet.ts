import { writable, Writable } from "svelte/store";
import { DataAPI, FolderTktAPI } from "../../../../../lib/apiv2";
import type { DataSheetAPI } from "../../../../../lib/apiv2/data_sheet";
import type {
  SheetCell,
  SheetColumn,
  SheetRow,
  Sheet,
  SheetWidget,
} from "../../../pages/data/sheet/sheets";
import { formatCells } from "./format";

export class SheetGroupService {
  source: string;
  group_slug: string;
  active_sheets: Map<string, SheetService>;

  sheets: Writable<Sheet[]>;

  folder_api: FolderTktAPI;
  data_api: DataAPI;
  data_sheet_api: DataSheetAPI;
  profile_genrator: (string) => string;

  constructor(
    source: string,
    group: string,
    api: DataAPI,
    profile_genrator: (string) => string
  ) {
    this.source = source;
    this.group_slug = group;
    this.active_sheets = new Map();
    this.sheets = writable([]);
    this.data_api = api;
    this.data_sheet_api = api.sheet_api();
    this.profile_genrator = profile_genrator;
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
  };

  get_sheet_service = async (sheetid: string) => {
    let ssvc = this.active_sheets.get(sheetid);
    if (ssvc) {
      return ssvc;
    }

    ssvc = new SheetService(this, sheetid, this.profile_genrator);

    await ssvc.init();

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
  profile_genrator: (string) => string;

  constructor(
    group: SheetGroupService,
    sheetid: string,
    profile_genrator: (string) => string
  ) {
    this.group = group;
    this.sheetid = sheetid;
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
  }

  init = async () => {
    const resp = await this.group.data_sheet_api.load_sheet(this.sheetid, {});
    if (!resp.ok) {
      return false;
    }

    const cell: SheetCell[] = resp.data["cells"] || [];
    const pcells = formatCells(cell);

    const rows = Object.keys(pcells)
      .map((v) => ({ __id: Number(v), sheetid: Number(this.sheetid) }))
      .sort((a, b) => {
        return Number(a["__id"]) - Number(b["__id"]);
      });

    this.state.set({
      columns: resp.data["columns"] || [],
      cells: pcells,
      rows,
      loading: false,
      widgets: resp.data["widget_apps"] || [],
      ref_columns: resp.data["reverse_ref_cols"] || [],
    });

    return true;
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
    await this.api.new_row_cell(this.sheetid, data);
  };

  update_row_cell = async (
    rid: string,
    data: { [_: number]: { [_: string]: any } }
  ) => {
    await this.api.update_row_cell(this.sheetid, rid, data);
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

  get_sibling_sheet = async (sheetid) => {
    return await this.group.get_sheet_service(sheetid);
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

  search = (search: string) => {
    return this.api.search(this.sheetid, search);
  };

  get_invoker(widget: SheetWidget) {
    return new SheetInvoker(this, widget)
  }
}

export class SheetInvoker {
  widget: SheetWidget;
  service: SheetService;
  constructor(service: SheetService, widget: SheetWidget) {
    this.widget = widget;
    this.service = service;
  }

  handle = (instance_id: string, msg_id: string, data: any) => {
    console.log("@instance_handle", instance_id, msg_id, data);
  };

  close = (instance_id: string) => {
    console.log("@close", instance_id);
  };
}

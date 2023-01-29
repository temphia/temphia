import { writable, Writable } from "svelte/store";
import type { DataAPI, FolderTktAPI } from "../../../../lib/apiv2";
import type {
  SheetCell,
  SheetColumn,
  SheetRow,
} from "../../pages/data/sheet/sheets";

export class SheetGroupService {
  data_api: DataAPI;

  source: string;
  group_slug: string;
  active_sheets: Map<string, SheetService>;

  sheets: object[];
  folder_api: FolderTktAPI;

  constructor(source: string, group: string, api: DataAPI) {
    this.data_api = api;
    this.source = source;
    this.group_slug = group;
    this.active_sheets = new Map();
    this.sheets = [];
    this.folder_api = null;
  }

  init = async () => {
    const resp = await this.data_api.list_sheets();
    if (!resp.ok) {
      console.log("@err", resp);
      return;
    }

    this.sheets = resp.data["sheets"] || [];
  };

  get_sheet_service = async (sheetid: string) => {
    let ssvc = this.active_sheets.get(sheetid);
    if (ssvc) {
      return ssvc;
    }

    ssvc = new SheetService(this, sheetid);

    await ssvc.init();

    this.active_sheets.set(sheetid, ssvc);
    return ssvc;
  };
}

export interface SheetState {
  columns: SheetColumn[];
  cells: { [_: string]: { [_: string]: SheetCell } };
  rows: SheetRow[];
}

export class SheetService {
  group: SheetGroupService;
  sheetid: string;
  state: Writable<SheetState>;

  constructor(group: SheetGroupService, sheetid: string) {
    this.group = group;
    this.sheetid = sheetid;

    this.state = writable({
      cells: {},
      columns: [],
      rows: [],
    });

    this.state.subscribe((state) => console.log("@state", state))
  }

  init = async () => {
    const resp = await this.group.data_api.load_sheet(this.sheetid, {});
    if (!resp.ok) {
      return false;
    }

    const cell: SheetCell[] = resp.data["cells"] || [];

    const pcells = cell.reduce((prev, cell) => {
      let row = prev[cell.rowid];
      if (!row) {
        row = {};
        prev[cell.rowid] = row;
      }

      row[cell.colid] = cell;

      return prev;
    }, {});

    const rows = Object.keys(pcells)
      .map((v) => ({ __id: Number(v), sheetid: Number(this.sheetid) }))
      .sort((a, b) => {
        return Number(a["__id"]) - Number(b["__id"]);
      });

    this.state.set({
      columns: resp.data["columns"],
      cells: pcells,
      rows,
    });

    return true;
  };
}

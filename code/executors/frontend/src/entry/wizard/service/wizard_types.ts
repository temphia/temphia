import type { Writable } from "svelte/store";

export interface State {
  stageTitle?: string;
  final: boolean;
  flowState:
    | "NOT_LOADED"
    | "SPLASH_LOADED"
    | "STAGE_LOADED"
    | "STAGE_PROCESSING"
    | "FINISHED";
  fields: object[];
  data_sources: { [_: string]: any };
  message?: string;
  epoch: number;
  errors?: { [_: string]: string };
  prev_data?: {[_: string]: any}
}

export interface Manager {
  wizard_title?: string

  get_state(): Writable<State>;
  get_field_store(field: string): FieldStore;

  init(): Promise<void>;
  splash_next(): Promise<void>;
  stage_next(): Promise<void>;
  stage_back(): Promise<void>;
}

export interface FieldStore {
  set_value(val: any): void;
  register_before_submit(fn: () => void): void;
  set_validity(valid: boolean): void;
  field_query(action: string, data: any): Promise<any>;
  verify_remote(data: any): Promise<any>;
}

export const BASIC = "basic";
export const BASIC_SHORTTEXT = "basic.shorttext";
export const BASIC_LONGTEXT = "basic.longtext";
export const BASIC_RANGE = "basic.range";
export const BASIC_SELECT = "basic.select";
export const BASIC_MULTI_SELECT = "basic.multiselect";
export const BASIC_PHONE = "basic.phone";
export const BASIC_CHECKBOX = "basic.checkbox";
export const BASIC_COLOR = "basic.color";
export const BASIC_DATE = "basic.date";
export const BASIC_DATETIME = "basic.datetime";
export const BASIC_EMAIL = "basic.email";
export const BASIC_NUMBER = "basic.number";
export const BASIC_PARAGRAPH = "basic.paragraph";

export const SELECT_MONTH = "select.month";
export const SELECT_WEEK = "select.week";
export const SELECT_NESTED = "select.nested";

export const IMAGE = "image";
export const IMAGE_INLINE = "image.inline";
export const FILE = "file";
export const FILE_INLINE = "file.inline";

export const MARKDOWN = "markdown";
export const MARKDOWN_PREVIEW = "markdown.preview";
export const SECRET = "secret";
export const QUESTION = "question";
export const FULLNAME = "fullname";

export const JSON_MULTI_SELECT = "json.select";
export const JSON_MULTI_INLINE = "json.inline";
export const JSON_MULTI_NESTED = "json.nested";
export const JSON_SINGLE_SELECT = "json.select";
export const JSON_SINGLE_INLINE = "json.inline";
export const JSON_SINGLE_NESTED = "json.nested";

export const LOCAT = "locat";
export const LOCAT_CIRCLE = "locat.circle";
export const LOCAT_AREA = "locat.area";
export const LOCAT_ADDRESS = "locat.addr"; // https://github.com/kelvins/geocoder
export const HTML = "html";

export const VIEW_IMAGE = "view.image";
export const VIEW_FILE = "view.file";
export const VIEW_CARDS = "view.cards";
export const VIEW_ALBUM = "view.album";
export const VIEW_CHARTJS = "view.chartjs";
export const VIEW_AUTOTABLE = "view.autotable";
export const VIEW_METRICS_CARD = "view.metrics_card";
export const VIEW_METRICS_TABLE = "view.metrics_table";

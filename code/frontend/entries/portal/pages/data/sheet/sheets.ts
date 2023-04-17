// syncme => xbprint/sheet.go

export const SheetColTypeText = "shorttext";
export const SheetColTypeLongText = "longtext";
export const SheetColTypeSelect = "select";
export const SheetColTypeMultiSelect = "multi_select";
export const SheetColTypeNumber = "number";
export const SheetColTypeDate = "datetime";
export const SheetColTypeBoolean = "bool";
export const SheetColTypeRatings = "ratings";
export const SheetColTypeLocation = "location";
export const SheetColTypeFile = "file";
export const SheetColTypeUser = "user";
export const SheetColTypeReference = "reference";
export const SheetColTypeRemote = "remote";

export const SheetColTypes = [
  SheetColTypeText,
  SheetColTypeSelect,
  SheetColTypeMultiSelect,
  SheetColTypeLongText,
  SheetColTypeNumber,
  SheetColTypeDate,
  SheetColTypeBoolean,
  SheetColTypeRatings,
  SheetColTypeLocation,
  SheetColTypeFile,
  SheetColTypeUser,
  SheetColTypeReference,
  SheetColTypeRemote,
];

export const SheetCtypeShapes = {
  text: [
    SheetColTypeText,
    SheetColTypeLongText,
    SheetColTypeRemote,
    SheetColTypeSelect,
    SheetColTypeMultiSelect,
    SheetColTypeUser,
  ],
  number: [SheetColTypeNumber, SheetColTypeReference],
};

export const SheetCtypeIcons = {
  [SheetColTypeText]: "annotation",
  [SheetColTypeSelect]: "menu-alt-4",
  [SheetColTypeMultiSelect]: "menu-alt-4",
  [SheetColTypeLongText]: "document-text",
  [SheetColTypeNumber]: "hashtag",
  [SheetColTypeDate]: "calendar",
  [SheetColTypeBoolean]: "check",
  [SheetColTypeRatings]: "star",
  [SheetColTypeLocation]: "map",
  [SheetColTypeFile]: "document",
  [SheetColTypeReference]: "paper-clip",
  [SheetColTypeUser]: "users",
  [SheetColTypeRemote]: "external-link",
};

export interface Sheet {
  __id: number;
  name: string;
}

export interface SheetColumn {
  __id: number;
  name: string;
  ctype: string;
  sheetid: number;
  opts?: string;
  extraopts?: object;
  color?: string;
  refsheet?: number;
  refcolumn?: number;
  remotehook?: number;
}

export interface SheetRow {
  __id: number;
  sheetid: number;
}

export interface SheetCell {
  __id: number;
  value?: string;
  num_value?: number;
  sheetid: number;
  rowid: number;
  colid: number;
}

export interface SheetWidget {
  id: number;
  name: string;
  icon?: string;
  context: string;
}

export interface SheetExecData {
  invoker_type: "data_sheet";
  sheet_id: string;
  data_group: string;

  cells: { [_: string]: { [_: string]: SheetCell } };
  rows: SheetRow[];
  columns?: SheetColumn[];
}

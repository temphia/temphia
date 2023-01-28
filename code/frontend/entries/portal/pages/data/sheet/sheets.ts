
// syncme => xbprint/sheet.go

export const SheetColTypeText = "shorttext";
export const SheetColTypeLongText = "longtext";
export const SheetColTypeNumber = "number";
export const SheetColTypeDate = "datetime";
export const SheetColTypeBoolean = "bool";
export const SheetColTypeRatings = "ratings";
export const SheetColTypeLocation = "location";
export const SheetColTypeFile = "file";
export const SheetColTypeReference = "reference"
export const SheetColTypeRemote    = "remote"

export const SheetColTypes = [
  SheetColTypeText,
  SheetColTypeLongText,
  SheetColTypeNumber,
  SheetColTypeDate,
  SheetColTypeBoolean,
  SheetColTypeRatings,
  SheetColTypeLocation,
  SheetColTypeFile,
  SheetColTypeReference,
  SheetColTypeRemote
];

export interface Sheet {
  id: number;
  name: string;
}

export interface SheetColumn {
  id: number;
  name: string;
  ctype: string;
  sheetid: number;
  exta_options?: {};
  color?: string;
}

export interface SheetRow {
  id: number;
  sheetid: number;
  color?: string;
}

export interface SheetCell {
  id: number;
  value?: string;
  num_value?: number;
  sheetid: number;
  rowid: number;
  colid: number;
}

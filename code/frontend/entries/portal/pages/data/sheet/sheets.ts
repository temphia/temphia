// syncme => xbprint/sheet.go

export const SheetColTypeText = "shorttext";
export const SheetColTypeLongText = "longtext";
export const SheetColTypeNumber = "number";
export const SheetColTypeDate = "datetime";
export const SheetColTypeBoolean = "bool";
export const SheetColTypeRatings = "ratings";
export const SheetColTypeLocation = "location";
export const SheetColTypeFile = "file";
export const SheetColTypeReferenceNum = "ref_text";
export const SheetColTypeReferenceText = "ref_number";
export const SheetColTypeRemoteText = "remote_text";
export const SheetColTypeRemoteNum = "remote_number";

export const SheetColTypes = [
  SheetColTypeText,
  SheetColTypeLongText,
  SheetColTypeNumber,
  SheetColTypeDate,
  SheetColTypeBoolean,
  SheetColTypeRatings,
  SheetColTypeLocation,
  SheetColTypeFile,
  SheetColTypeReferenceNum,
  SheetColTypeReferenceText,
  SheetColTypeRemoteNum,
  SheetColTypeRemoteText,
];

export const SheetCtypeIcons = {
  [SheetColTypeText]: "annotation",
  [SheetColTypeLongText]: "document-text",
  [SheetColTypeNumber]: "hashtag",
  [SheetColTypeDate]: "calendar",
  [SheetColTypeBoolean]: "check",
  [SheetColTypeRatings]: "star",
  [SheetColTypeLocation]: "map",
  [SheetColTypeFile]: "document",
  [SheetColTypeReferenceNum]: "paper-clip",
  [SheetColTypeReferenceText]: "paper-clip",
  [SheetColTypeRemoteNum]: "external-link",
  [SheetColTypeRemoteText]: "external-link",
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
  exta_options?: {};
  color?: string;
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

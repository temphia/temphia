// syncme => xbprint/sheet.go

export interface NewSheetGroup {
  name: string;
  sheets: NewSheet[];
}

export interface NewSheet {
  name: string;
  columns: NewSheetColumn[];
  seed_data: { [_: string]: any }[];
}

export interface NewSheetColumn {
  name: string;
  ctype: string;
  color: string;
  extra_options: { [_: string]: string }[];
}

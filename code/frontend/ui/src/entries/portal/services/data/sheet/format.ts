import type { SheetCell, SheetColumn } from "../../../pages/data/sheet/sheets";
export const formatRefCells = (rdata: {
  columns: SheetColumn[];
  cells: SheetCell[];
  sheet_id: number;
}) => {
  const cells = formatCells(rdata.cells || []);

  return {
    rows: Object.keys(cells).map((v) => ({ __id: v, sheetid: rdata.sheet_id })),
    cells,
    columns: rdata.columns,
  };
};

export const formatCells = (cells: SheetCell[]) => {
  return cells.reduce((prev, cell) => {
    let row = prev[cell.rowid];
    if (!row) {
      row = {};
      prev[cell.rowid] = row;
    }

    row[cell.colid] = cell;

    return prev;
  }, {});
};

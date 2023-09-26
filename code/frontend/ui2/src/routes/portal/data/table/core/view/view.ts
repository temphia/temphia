import type { TableService } from "$lib/services/data";

export interface ViewModal {
  open: (compo, options) => void;
  close: () => void;
}

export interface DataContext {
  get_modal: () => ViewModal;
  table_service: TableService;
}

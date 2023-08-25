import type { Writable } from "svelte/store";
import type { DirtyData } from "./state_types";

export class DirtyRowService {
  dirtyStore: Writable<DirtyData>;
  callbacks: Map<string, () => void>;
  constructor(store: Writable<DirtyData>) {
    this.dirtyStore = store;
    this.callbacks = new Map();
  }

  register_before_save(field: string, callback: () => void): void {
    this.callbacks.set(field, callback);
  }

  on_ohange(_field: string, _value: any): void {
    this.set_value(_field, _value);
  }

  // row stuff
  start_modify_row = (row: number) => {
    this.callbacks.clear();
    this.dirtyStore.set({ rowid: row, data: {} });
  };

  start_new_row = () => {
    this.callbacks.clear();
    this.dirtyStore.set({ rowid: 0, data: {} });
  };

  set_value = (_filed: string, value: any) => {
    this.dirtyStore.update((old) => ({
      ...old,
      data: { ...old.data, [_filed]: value },
    }));
  };

  clear_dirty_row = () => {
    this.dirtyStore.set({ rowid: 0, data: {} });
  };

  set_ref_copy(column: string, value: any) {
    this.dirtyStore.update((old) => ({
      ...old,
      data: { ...old.data, [column]: value },
    }));
  }

  before_save() {
    this.callbacks.forEach((val) => val());
  }
}

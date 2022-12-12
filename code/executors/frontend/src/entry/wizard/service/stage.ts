import type { FieldStore, Manager } from "./wizard_types";

export class StageStore {
  _manager: Manager;
  _values: { [_: string]: any };
  _eventHandlers: Array<() => void>;

  constructor(m: Manager) {
    this._manager = m;
    this._values = new Map();
    this._eventHandlers = new Array(0);
  }

  get_field_store = (field: string) => {
    return new fieldStore(this, field);
  };

  set_value = (field: string, val: any) => {
    this._values[field] = val;
  };

  field_query(field: string, action: string, data: any): Promise<any> {
    return null;
  }

  verify_remote(field: string, data: any): Promise<any> {
    return null;
  }

  register_before_submit(fn: () => void) {
    this._eventHandlers.push(fn);
  }

  // private

  _apply_event() {
    this._eventHandlers.forEach((eh) => {
      eh();
    });
  }

  _get_values() {
    this._apply_event();
    return this._values;
  }
}

export class fieldStore implements FieldStore {
  _stage: StageStore;
  _field_name: string;

  constructor(stage: StageStore, field_name: string) {
    this._stage = stage;
    this._field_name = field_name;
  }

  set_value(val: any): void {
    this._stage.set_value(this._field_name, val);
  }

  register_before_submit(fn: () => void): void {
    this._stage.register_before_submit(fn);
  }

  set_validity(valid: boolean): void {}

  field_query(action: string, data: any): Promise<any> {
    return null;
  }

  verify_remote(data: any): Promise<any> {
    return null;
  }
}

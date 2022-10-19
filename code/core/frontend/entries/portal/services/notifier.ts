import { writable, Writable } from "svelte/store";
import type { SelfAPI } from "../../../lib/apiv2";

export interface Options {
  self_api: SelfAPI;
}

export interface State {
  messages: object[];
  loading: boolean;
  cursor: number;
}

export class Notifier {
  self_api: SelfAPI;
  state: Writable<State>;

  constructor(opts: Options) {
    this.self_api = opts.self_api;
    this.state = writable({ messages: [], cursor: 0, loading: false });
  }

  init() {}
  read_message() {}
  delete_message() {}

  toggle() {}
}

import { writable, Writable } from "svelte/store";
import type { SelfAPI } from "../../../lib/apiv2";

export interface State {
  messages: object[];
  loading: boolean;
  cursor: number;
}

export class Notifier {
  self_api: SelfAPI;
  state: Writable<State>;

  constructor(self_api: SelfAPI) {
    this.self_api = self_api;
    this.state = writable({ messages: [], cursor: 0, loading: false });
  }

  async init() {}
  read_message(id: number) {}
  delete_message(id: number) {}
}

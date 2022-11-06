import { writable, Writable } from "svelte/store";

export class Launcher {
  display_state: Writable<"HIDDEN" | "FLOATING" | "NOT_FLOATING">;
  constructor() {
    this.display_state = writable("HIDDEN");
  }

  hide() {
    this.display_state.set("HIDDEN");
  }

  float() {
    this.display_state.set("FLOATING");
  }

  not_float() {
    this.display_state.set("NOT_FLOATING");
  }
}

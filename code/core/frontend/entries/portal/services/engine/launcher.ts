import { writable, Writable } from "svelte/store";
import { generateId } from "../../../../lib/utils";

interface LauncherState {
  display: "HIDDEN" | "FLOATING" | "SHOW";
  active_instance?: string;
  instances: Instance[];
}

interface Instance {
  id: string;
  target_id: number;
  plug_id: string;
  agent_id: string;
  name: string;
  invoker?: {
    close_instance: (id: string) => void;
    handle_message: (id: string, data: any) => void;
  };
  invoker_name: string;
}

export class Launcher {
  state: Writable<LauncherState>;
  constructor() {
    this.state = writable({
      display: "HIDDEN",
      instances: [],
    });
  }

  plane_hide() {
    this.state.update((old) => ({ ...old, display: "HIDDEN" }));
  }

  plane_float() {
    this.state.update((old) => ({ ...old, display: "FLOATING" }));
  }

  plane_show() {
    this.state.update((old) => ({ ...old, display: "SHOW" }));
  }

  instance_by_target(target_app: object): string {
    const instance_id = generateId();

    this.state.update((old) => ({
      ...old,
      active_instance: instance_id,
      instances: [
        ...old.instances,
        {
          id: instance_id,
          invoker_name: "fixme",
          name: "Test1",
          plug_id: "fixme",
          agent_id: "fixme",
          target_id: 0,
        },
      ],
    }));

    return instance_id;
  }
}

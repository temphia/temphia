import { writable, Writable } from "svelte/store";
import { generateId } from "../../../../lib/utils";

interface LauncherState {
  display: "HIDDEN" | "FLOATING" | "SHOW";
  active_instance?: string;
  instances: Instance[];
}

export interface Instance {
  id: string;
  target_id: string;
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

  target_index: { [_: string]: string };

  constructor() {
    this.state = writable({
      display: "HIDDEN",
      instances: [],
    });

    this.state.subscribe((_state) => {
      this.target_index = {};
      _state.instances.forEach((instance) => {
        this.target_index[instance.target_id] = instance.id;
      });
    });

    this.state.subscribe((lstate) => console.log("@launcher_state", lstate));
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
          target_id: target_app["target_id"],
        },
      ],
    }));

    return instance_id;
  }

  instance_change(instance_id: string) {
    // fixme => check if instance is still in instances array
    this.state.update((old) => ({ ...old, active_instance: instance_id }));
  }

  instance_close(instance_id: string) {
    this.state.update((old) => {
      const instances = old.instances.filter((v) => v.id !== instance_id);

      return {
        ...old,
        active_instance: instances.length > 0 ? instances[0].id : undefined,
        instances,
      };
    });
  }
}

import { get, writable, Writable } from "svelte/store";
import { generateId } from "../../../../lib/utils";

interface LauncherState {
  display: "HIDDEN" | "FLOATING" | "SHOW";
  active_instance?: string;
  instances: Instance[];
}

export interface Instance {
  id: string;
  target_id: string;
  target_type: string;
  name: string;
  invoker?: {
    close_instance: (id: string) => void;
    handle_message: (id: string, data: any) => void;
  };
  invoker_name: string;
}

export interface InvokerOptions {
  invoker_name: string;
  invoker?: any;
  target_name?: string;
  target_type?: string;
  target_id: string;
}

export class Launcher {
  state: Writable<LauncherState>;
  bootloader?: string;

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

  get_bootloader = async () => {
    if (this.bootloader) {
      return this.bootloader;
    }

    const resp = await fetch("/z/assets/build/executor_bootloader_iframe.js");
    this.bootloader = await resp.text();
    return this.bootloader;
  };

  plane_hide() {
    this.state.update((old) => ({ ...old, display: "HIDDEN" }));
  }

  plane_not_float() {
    if (get(this.state).display === "FLOATING") {
      this.state.update((old) => ({ ...old, display: "HIDDEN" }));
    }
    
  }

  plane_float() {
    this.state.update((old) => ({ ...old, display: "FLOATING" }));
  }

  plane_show() {
    this.state.update((old) => ({ ...old, display: "SHOW" }));
  }

  instance_by_target(topts: InvokerOptions): string {
    const instances = get(this.state).instances;
    const old = instances.filter((v) => v.target_id === topts.target_id);

    if (old.length > 0) {
      this.instance_change(old[0].id);
      return;
    }

    const instance_id = generateId();

    this.state.update((old) => ({
      ...old,
      active_instance: instance_id,
      instances: [
        ...old.instances,
        {
          id: instance_id,
          invoker_name: topts.invoker_name || "",
          name: topts.target_name || "",
          target_type: topts.target_type || "",
          target_id: topts.target_id,
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

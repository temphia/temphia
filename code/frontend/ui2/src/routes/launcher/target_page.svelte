<script lang="ts">
  import { getContext, onDestroy, afterUpdate, onMount } from "svelte";
  import type { PortalService } from "$lib/services/portal/portal";

  let target = ""

  const get_params = () => {
    let currtarget = location.hash; // fixme => parse it
    let name = "";
    if (!target) {
      target = currtarget 
    }
    return [name, target];
  };

  const app = getContext("__app__") as PortalService;
  let lopts = app.nav.options || {};

  const launcher = app.launcher;

  onMount(async () => {
    const [name, currtarget] = get_params()

    console.log("@on_mount target |> ", currtarget);
    let instance = launcher.target_index[currtarget];
    if (instance) {
      launcher.instance_change(instance);
    } else {
      instance = launcher.instance_by_target({
        invoker_name: "user_app",
        target_id: currtarget,
        target_name: name,
        target_type: lopts["target_type"] || "",
        startup_payload: {},
        invoker: null,
      });
    }
    launcher.plane_show();
  });

  afterUpdate(() => {
    const [name, currtarget] = get_params()

    const final_target = currtarget;
    if (final_target && target !== final_target) {
      target = final_target;
      let instance = launcher.target_index[final_target];
      launcher.instance_change(instance);
    }
    console.log(
      "This is after update (target|final_target)",
      target,
      "|",
      final_target
    );
  });

  onDestroy(() => {
    launcher.plane_hide();
  });
</script>

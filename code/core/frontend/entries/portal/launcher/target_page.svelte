<script lang="ts">
  import { getContext, onDestroy, afterUpdate, onMount } from "svelte";
  import type { PortalService } from "../services";
  import { params } from "svelte-hash-router";

  export let target = $params.target;

  const app = getContext("__app__") as PortalService;
  const launcher = app.launcher;

  onMount(async () => {
    console.log("@on_mount target |> ", target);
    let instance = launcher.target_index[target];
    if (instance) {
      launcher.instance_change(instance);
    } else {
      const name = $params._ ? window.atob($params._) : "";
      instance = launcher.instance_by_target({ target_id: target, name });
    }
    launcher.plane_show();
  });

  afterUpdate(() => {
    const final_target = $params.target;
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

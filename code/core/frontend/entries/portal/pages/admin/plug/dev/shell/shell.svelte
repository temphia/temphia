<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { LoadingSpinner } from "../../../../../../xcompo";
  import { DevShellService } from "../../../../../services/engine/dev_shell";
  import type { PortalService } from "../../../core";

  import ShellInner from "./inner.svelte";

  export let pid = $params.pid;
  export let aid = $params.aid;

  const app: PortalService = getContext("__app__");

  let message = "";
  let loading = true;
  let interface_loaded = false;
  let iface_data = {};

  let dev_shell_service = new DevShellService(app.api_manager, pid, aid);

  dev_shell_service.init().then((val) => {
    if (typeof val === "string") {
      message = val;
    } else {
      loading = false;
    }
  });
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <ShellInner {aid} {pid} service={dev_shell_service} {app} />
{/if}

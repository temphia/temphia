<script lang="ts">
  import type { PortalService } from "../../../core";
  import Draggable from "./_draggable.svelte";
  import Plug from "./plug.svelte";
  import { getContext } from "svelte";

  export let pid: string;

  const app = getContext("__app__") as PortalService;

  let rootElem;

  const data = {
    plug: {
      id: "chc50lom4q7efu3enuq0",
      name: "Simpledoc App",
      live: true,
      dev: true,
      bprint_id: "chc50kom4q7efu3enul0",
      tenant_id: "default0",
    },
    agents: [
      {
        id: "default",
        name: "default",
        type: "web",
        executor: "javascript1",
        plug_id: "chc50lom4q7efu3enuq0",
      },
    ],
    agent_links: {
      default: {
        id: 1,
        name: "link2adaper",
        from_plug_id: "chc50lom4q7efu3enuq0",
        from_agent_id: "default",
        to_plug_id: "adapter-1",
        to_agent_id: "default",
        to_handler: "test",
        tenant_id: "default0",
      },
    },
    target_apps: [
      {
        id: 1,
        name: "Documents",
        target_type: "user_group_app",
        target: "super_admin",
        context_type: "app.1",
        plug_id: "chc50lom4q7efu3enuq0",
        agent_id: "default",
        tenant_id: "default0",
      },
    ],
  };

  let loading = true;

  const papi = app.api_manager.get_admin_plug_api();

  const load = async () => {
    const resp = await papi.flowmap(pid);
    if (!resp.ok) {
      console.log("Err", resp);
      return;
    }
    console.log("lll", resp.data);
  };

  load();
</script>

<div class="h-full w-full max-h-screen p-2" bind:this={rootElem}>
  <div
    class="h-full w-full rounded overflow-scroll border border-slate-900 bg-white"
  >
    <div
      class="relative"
      style="min-width:5000px; min-height:5000px; transform: scale(1); 
        transform-origin: 0% 0% 0px;
        background-image: radial-gradient(rgba(15, 15, 16, 0.33) 1px, transparent 1px); 
              background-size: 13px 13px; background-color: rgba(71, 211, 255, 0.06);
        "
    >
      <Draggable left={2500} top={2500}>
        <Plug />
      </Draggable>
    </div>
  </div>
</div>

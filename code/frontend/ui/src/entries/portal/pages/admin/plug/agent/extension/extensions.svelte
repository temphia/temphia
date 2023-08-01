<script lang="ts">
    import { getContext } from "svelte";
    import {
      AutoTable,
      LoadingSpinner,
      FloatingAdd,
      PortalService,
    } from "../../../core";
    import { params } from "svelte-hash-router";
  
    export let pid = $params.pid;
    export let aid = $params.aid;
  
    let datas = [];
    let loading = true;
    const app = getContext("__app__") as PortalService;
  
    const api = app.api_manager.get_admin_plug_api();
  
    const load = async () => {
      const resp = await api.list_agent_ext(pid, aid);
      if (!resp.ok) {
        return;
      }
  
      datas = resp.data;
      loading = false;
    };
  
    load();
  
    // actions
  
    const action_edit = (id: string) => {};
    const action_delete = async (id: string) => {
      await api.delete_agent_ext(pid, aid, id);
      load();
    };
  
    const action_new = () => app.nav.admin_agent_ext_new(pid, aid);
  </script>
  
  {#if loading}
    <LoadingSpinner />
  {:else}
    <AutoTable
      action_key="id"
      actions={[
        {
          Name: "Edit",
          Action: action_edit,
          drop: true,
          icon: "trash",
        },
        {
          Name: "Delete",
          Action: action_delete,
          drop: true,
          icon: "trash",
        },
      ]}
      key_names={[
        ["id", "ID"],
        ["name", "Name"],
        ["plug_id", "Plug"],
        ["agent_id", "Agent"],
        ["bprint_id", "Bprint"],
        ["ref_file", "File"],
      ]}
      {datas}
    />
  {/if}
  
  <FloatingAdd onClick={action_new} />
  
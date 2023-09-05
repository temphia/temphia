<script lang="ts">
    import { getContext } from "svelte";
    import {
      AutoTable,
      LoadingSpinner,
      FloatingAdd,
      PortalService,
    } from "$lib/core";
    import { params } from "$lib/params";
    import NewPicker from "./_new_picker.svelte";
  
    export let pid = $params["pid"];
    export let aid = $params["aid"];
  
    let datas = [];
    let loading = true;
    const app = getContext("__app__") as PortalService;
  
    const api = app.api_manager.get_admin_plug_api();
  
    const load = async () => {
      const resp = await api.list_agent_link(pid, aid);
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
      await api.delete_agent_link(pid, aid, id);
      load();
    };
  
    const action_new = () => {
      app.utils.small_modal_open(NewPicker, { service: app, pid, aid });
    };
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
        ["from_plug_id", "From Plug"],
        ["from_agent_id", "From Agent"],
        ["to_plug_id", "To Plug"],
        ["to_agent_id", "To Agent"],
      ]}
      {datas}
    />
  {/if}
  
  <FloatingAdd onClick={action_new} />
  
<script lang="ts">
    import { getContext } from "svelte";
    import {
      AutoTable,
      LoadingSpinner,
      FloatingAdd,
      PortalService,
    } from "../../core";
  
    const app = getContext("__app__") as PortalService;
  
    let datas = [];
    let loading = true;
    const api = app.api_manager.get_admin_target_api();
  
    const load = async () => {
      const resp = await api.listHook();
      if (!resp.ok) {
        return;
      }
  
      datas = resp.data;
      loading = false;
    };
  
    load();
  
    // actions
  
    const action_edit = (id: string) => app.nav.admin_target_hook_edit(Number(id));
    const action_delete = async (id: string, data: object) => {
        const resp = await api.deleteHook(data["target_type"], Number(id));
        if (resp.ok) {
          return;
        }
        load();
    };
  
    const action_new = () => app.nav.admin_target_hook_new();
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
          icon: "pencil-alt",
        },
        {
          Name: "Delete",
          Class: "bg-red-400",
          Action: action_delete,
          icon: "trash",
        },
      ]}
      key_names={[
        ["name", "Name"],
        ["id", "Id"],
        ["target_type", "Type"],
        ["target", "Target"],
      ]}
      color={["target_type"]}
      {datas}
    />
  {/if}
  
  <FloatingAdd onClick={action_new} />
  
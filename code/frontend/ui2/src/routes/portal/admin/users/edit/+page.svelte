<script lang="ts">
    import { getContext } from "svelte";
    import { AutoForm, LoadingSpinner, PortalService } from "$lib/core";
    import { params } from "$lib/params";
    
    export let uid = $params["userid"];
  
    const app = getContext("__app__") as PortalService;
    const api = app.api_manager.get_admin_user_api();
  
    let message = "";
    let data = {};
    let loading = true;
  
    const load = async () => {
      const resp = await api.get(uid);
      if (!resp.ok) {
        return;
      }
  
      data = resp.data;
      loading = false;
    };
  
    load();
  
    const save = async (_data) => {
      const resp = await api.update(uid, _data);
      if (!resp.ok) {
        message = resp.data;
        return;
      }
      app.nav.admin_users();
    };
  </script>
  
  {#if loading}
    <LoadingSpinner />
  {:else}
    <AutoForm
      {message}
      schema={{
        fields: [
          {
            name: "Full Name",
            ftype: "TEXT",
            key_name: "full_name",
          },
          {
            name: "Bio",
            ftype: "LONG_TEXT",
            key_name: "bio",
          },
        ],
        name: "Edit User",
        required_fields: [],
      }}
      onSave={save}
      {data}
    />
  {/if}
  
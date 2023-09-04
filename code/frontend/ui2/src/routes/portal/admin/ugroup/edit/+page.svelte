<script lang="ts">
    import { getContext } from "svelte";
    import { AutoForm, LoadingSpinner, PortalService } from "$lib/core";
    import { params  } from "$lib/params";
    
    let data = $params;

    let ugroup = data["ugroup"];
  
    const app = getContext("__app__") as PortalService;
    const api = app.api_manager.get_admin_ugroup_api();
  
    let message = "";
    let rdata = {};
    let loading = true;
  
    const load = async () => {
      const resp = await api.get(ugroup);
      if (!resp.ok) {
        return;
      }
  
      rdata = resp.data;
      loading = false;
    };
  
    load();
  
    const save = async (_data) => {
      const resp = await api.update(ugroup, _data);
      if (!resp.ok) {
        message = resp.data;
        return;
      }
      app.nav.admin_ugroups();
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
            name: "Name",
            ftype: "TEXT",
            key_name: "name",
          },
          {
            name: "Slug",
            ftype: "TEXT_SLUG",
            key_name: "slug",
            slug_gen: null, // fixme
          },
  
          {
            name: "Scopes",
            ftype: "MULTI_TEXT",
            key_name: "scopes",
          },
  
          {
            name: "Features",
            ftype: "MULTI_TEXT",
            key_name: "features",
          },
  
          {
            name: "Feature Options",
            ftype: "KEY_VALUE_TEXT",
            key_name: "feature_opts",
          },
  
          {
            name: "Extra Meta",
            ftype: "KEY_VALUE_TEXT",
            key_name: "extra_meta",
          },
        ],
        name: "Edit Data Group",
        required_fields: [],
      }}
      onSave={save}
      data={rdata}
    />
  {/if}
  
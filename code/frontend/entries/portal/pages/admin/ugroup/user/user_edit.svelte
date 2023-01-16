<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../../core";
  import { params } from "svelte-hash-router";

  export let ugroup = $params.ugroup;
  export let userid = $params.userid;


  const app = getContext("__app__") as PortalService;

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const api = await app.api_manager.get_ugroup_tkt_api(ugroup);
    if (!api) {
      return;
    }

    const resp = await api.get(userid);
    if (!resp.ok) {
      return;
    }

    data = resp.data;
    loading = false;
  };

  const save = async (_data) => {
    const api = await app.api_manager.get_ugroup_tkt_api(ugroup);
    if (!api) {
      console.log("ugroup tkt api not found");
      return;
    }

    const resp = await api.update(userid, _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_ugroup_users(ugroup)
  };

  load();
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

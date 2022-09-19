<script lang="ts">
  import { getContext } from "svelte";
  import type { AuthApp } from "../../../lib/app/auth/auth";

  import { SelfAPI } from "../../../lib/core/api";
  import { apiURL } from "../../../lib/core/site";
  import UserCard from "../common/user_card.svelte";

  const app: AuthApp = getContext("_auth_app_");

  let userdata;
  let loading = true;

  app.clear_preauthed_data();

  (async () => {
    const data = app.site_manager.getAuthedData();
    const selfapi = new SelfAPI(apiURL(data.tenant_id), data.user_token);

    await selfapi.init();

    const resp = await selfapi.get_self_info();
    if (resp.status !== 200) {
      return;
    }
    userdata = resp.data;
    loading = false;
  })();
</script>

{#if loading}
  <div>loading...</div>
{:else}
  <UserCard
    bio={userdata["bio"] || ""}
    full_name={userdata["full_name"] || ""}
    group_name={userdata["group_name"] || ""}
    tenant_id={userdata["tenant_id"] || ""}
    tenant_name={userdata["tenant_name"] || ""}
    user_id={userdata["user_id"] || ""}
  />
{/if}

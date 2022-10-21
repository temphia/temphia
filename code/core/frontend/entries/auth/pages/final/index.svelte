<script lang="ts">
  import { getContext } from "svelte";
  import { apiURL } from "../../../../lib/utils/site";
  import type { AuthApp } from "../../app";
  import UserCard from "../common/user_card.svelte";

  const app: AuthApp = getContext("_auth_app_");

  let userdata;
  let loading = true;

  app.clear_preauthed_data();

  (async () => {
    console.log("FIXME")
    // const data = app.site_manager.getAuthedData();
    // userdata = resp.data;
    // loading = false;
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

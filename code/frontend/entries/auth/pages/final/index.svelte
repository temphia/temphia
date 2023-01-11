<script lang="ts">
  import { getContext } from "svelte";
  import { apiURL } from "../../../../lib/utils/site";
  import type { AuthService } from "../../services";
  import UserCard from "../common/user_card.svelte";

  const app: AuthService = getContext("_auth_app_");

  let userdata;
  let orgdata;
  let loading = true;

  app.clear_preauthed_data();

  (async () => {
    const resp = await app.about();
    console.log(resp);
    userdata = resp["user_info"];
    orgdata = resp["org_info"];
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
    tenant_id={orgdata["slug"] || ""}
    
    tenant_name={orgdata["name"] || ""}
    user_id={userdata["user_id"] || ""}
  />
{/if}


<script lang="ts">
  import { getContext } from "svelte";
  import { FloatingAdd } from "../../../../../shared";
  import type { PortalApp } from "../../../app";
  import Layout from "../layout.svelte";
  import Users from "./users.svelte";

  export let id;

  const app: PortalApp = getContext("__app__");
</script>

<Layout current_item={"user_groups"}>
  <Users user_group={id} />
</Layout>

<FloatingAdd
  onClick={() => {
    app.navigator.goto_admin_new_user_page(id);
  }}
/>

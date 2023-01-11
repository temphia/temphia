<script lang="ts">
  import AdminLayout from "../../layout/admin.svelte";
  import { Router } from "svelte-hash-router";
  import { registerEditors } from "./bprint/editor/register";
  import { getContext } from "svelte";
  import type { PortalService } from "./core";

  const app = getContext("__app__") as PortalService;

  registerEditors(app.xtmgr);
</script>

<AdminLayout>
  <Router />
</AdminLayout>

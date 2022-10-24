<script lang="ts">
  import { params } from "svelte-hash-router";
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";
  import Skeleton from "./listings/_skeleton.svelte";

  const app = getContext("__app__") as PortalService;
  let sources;

  (async () => {
    sources = await app.api_manager.get_repo_sources();
    app.nav.repo_source(Object.entries(sources)[0][0]);
  })();
</script>

<Skeleton />

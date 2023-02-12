<script lang="ts">
  import Builder from "./_builder/builder.svelte";
  import { params } from "svelte-hash-router";
  import type { EasypageService } from "../service/easypage";
  import { getContext } from "svelte";

  export let pid = $params.pid;

  const service = getContext("__easypage_service__") as EasypageService;
</script>

<Builder page_slug={pid} {service} />

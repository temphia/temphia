<script lang="ts">
  import { LoadingSpinner } from "../../xcompo";
  import { PageDashService } from "./service";
  import Pagedash from "./pagedash.svelte";
  import Tailwind from "../../xcompo/common/_tailwind.svelte";
  import type { Environment } from "../../../lib/engine/environment";

  export let env: Environment;

  let loading = false;
  let data = {};
  let service: PageDashService;

  const load = async () => {
    if (!service) {
      service = env.Extend(new PageDashService(env));
    }

    const resp = await service.load({});
    if (!resp.ok) {
      return;
    }
    data = resp.data;
    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <Pagedash {service} {data} />
{/if}

<Tailwind />

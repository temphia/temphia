<script lang="ts">
  import { LoadingSpinner } from "../../xcompo";
  import { PageDashService } from "./service";
  import Pagedash from "./pagedash.svelte";
  import Tailwind from "../../xcompo/common/_tailwind.svelte";

  export let env: any;

  let loading = false;
  let data = {};
  let service;

  const load = async () => {
    if (!service) {
      service = new PageDashService(env);
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

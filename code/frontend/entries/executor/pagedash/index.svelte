<script lang="ts">
  import { LoadingSpinner } from "../../xcompo";
  import { PageDashService } from "./service";
  export let env: any;

  let loading = false;

  const load = async () => {
    const service = new PageDashService(env);
    const data = await service.load({});
    if (!data) {
      return;
    }
    console.log("@resp", data);
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <div>Pagedash</div>
{/if}

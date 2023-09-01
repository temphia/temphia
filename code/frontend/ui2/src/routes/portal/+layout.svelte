<script lang="ts">
  import { onMount } from "svelte";
  import { PortalService } from "$lib/services/portal/portal";
  import { SiteUtils, baseURL } from "$lib/utils/site";
  import { LoadingSpinner } from "$lib/compo";

  import Noauth from "./layout/noauth.svelte";
  import Root from "./layout/root.svelte";

  import type { Registry } from "$lib/services/portal/registry/registry";

  let loading = true;
  let ok = true;
  let app;

  const load = async () => {
    const site = new SiteUtils();

    if (!site.isLogged()) {
      ok = false;
      loading = false;
    }

    const sdata = site.getAuthedData();
    app = new PortalService({
      base_url: baseURL(),
      registry: window["__registry__"] as Registry<any>,
      site_utils: site,
      tenant_id: sdata.tenant_id,
      user_token: sdata.user_token,
    });

    await app.init();
    loading = false;
  };

  load();

  onMount(() => {
    window.onunhandledrejection = (e) => {
      console.log("we got exception, but the app has crashed", e);
      e.preventDefault();
    };
  });
</script>

{#if loading}
  <LoadingSpinner />
{:else if !ok}
  <Noauth />
{:else}
  <Root launcher={app.launcher} {app}>
    <svelte:fragment>
      <slot />
    </svelte:fragment>
  </Root>
{/if}

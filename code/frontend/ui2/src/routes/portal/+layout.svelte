<script lang="ts">
  import { onMount } from "svelte";
  import { PortalService } from "$lib/services/portal/portal";
  import { baseURL } from "$lib/utils/site";

  import LayoutInner from "./_layout_inner.svelte";
  import { LoadingSpinner } from "$lib/compo";

  let loading = true;

  const app = new PortalService({
    base_url: baseURL(),
    registry: null,
    site_utils: null,
    tenant_id: "",
    user_token: "",
  });

  onMount(() => {
    window.onunhandledrejection = (e) => {
      console.log("we got exception, but the app has crashed", e);
      e.preventDefault();
    };
  });
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <LayoutInner pending_notification={false} launcher={null}>
    <svelte:fragment>
      <slot />
    </svelte:fragment>
  </LayoutInner>
{/if}

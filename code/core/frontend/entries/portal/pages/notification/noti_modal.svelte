<script lang="ts">
  import Inner from "./notification.svelte";
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../lib/app/portal";

  const app: PortalApp = getContext("__app__");
  const show = app.notifier.is_open;
</script>

{#if $show}
  <div
    class="w-full h-full bg-gray-800 bg-opacity-90 top-0 overflow-y-auto overflow-x-hidden fixed sticky-0 z-50 left-14"
  >
    <div
      class="w-full absolute h-full overflow-x-hidden transform translate-x-0 transition ease-in-out duration-700 "
    >
      <Inner noti={app.notifier} />
    </div>
  </div>
{/if}

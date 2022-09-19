<script lang="ts">
  import { getContext, setContext } from "svelte";
  import { buildApp } from "../../../src/lib/app/portal";
  import { Toaster } from "../../../src/entries/common";
  import NotiModal from "../../../src/entries/portal/pages/notification/noti_modal.svelte";
  import NavDesktop from "./nav_desktop.svelte";
  import NavMobile from "./nav_mobile.svelte";

  const { close, open } = getContext("simple-modal");

  let loaded = false;
  let toast_error;
  let toast_success;

  const app = buildApp(open, close, {
    error: (msg) => toast_error(msg),
    success: (msg) => toast_success(msg),
  });

  app.init().then(() => {
    loaded = true;
  });

  setContext("__app__", app);
</script>

<Toaster bind:error={toast_error} bind:success={toast_success} />

{#if loaded}
  <NotiModal />
  {#if app.is_mobile()}
    <div class="h-screen w-screen flex flex-col">
      <NavMobile {app} />
      <div class="h-full w-full overflow-auto">
        <slot />
      </div>
    </div>
  {:else}
    <div class="h-screen w-screen flex">
      <NavDesktop {app} />

      <div
        class="h-screen overflow-auto"
        style="width:95%;width: calc(100vw - 3rem);"
      >
        <slot />
      </div>
    </div>
  {/if}
{:else}
  <div>loading..</div>
{/if}

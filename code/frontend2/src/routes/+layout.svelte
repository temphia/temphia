<script lang="ts">
  import "@skeletonlabs/skeleton/themes/theme-hamlindigo.css";
  import "@skeletonlabs/skeleton/styles/skeleton.css";
  import {
    AppShell,
    AppBar,
    AppRail,
    AppRailAnchor,
    AppRailTile,
  } from "@skeletonlabs/skeleton";

  import { WSService } from "$lib";

  let currentTile = 0;

  const ws = new WSService();
</script>

<AppShell>
  <svelte:fragment slot="sidebarLeft">
    <AppRail>
      <AppRailTile
        bind:group="{currentTile}"
        name="tile-1"
        value="home"
        title="Home"
      >
        <svelte:fragment slot="lead">Home</svelte:fragment>
        <span>Home</span>
      </AppRailTile>

      <AppRailTile
        bind:group="{currentTile}"
        name="tile-2"
        value="{1}"
        title="tile-2"
      >
        <svelte:fragment slot="lead">(icon)</svelte:fragment>
        <span>Tile 2</span>
      </AppRailTile>

      <!-- --- -->
      <svelte:fragment slot="trail">
        <AppRailAnchor href="/" target="_blank" title="Account"
          >(icon)</AppRailAnchor
        >
      </svelte:fragment>
    </AppRail>
  </svelte:fragment>

  <slot />
</AppShell>
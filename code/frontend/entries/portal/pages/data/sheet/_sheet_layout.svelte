<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { createEventDispatcher } from "svelte";
  import { ActionNormal } from "../shared";
  import type { Sheet, SheetWidget } from "./sheets";

  export let sheets: Sheet[];
  export let active_sheet: number;
  export let selected_rows = [];
  export let widgets: SheetWidget[];

  const dispatch = createEventDispatcher();
</script>

<div class="flex flex-col p-2 rounded">
  <nav class="flex flex-row border flex-nowrap overflow-auto">
    {#each sheets as sheet}
      {#if sheet.__id === active_sheet}
        <button
          class="p-2 hover:text-blue-500 focus:outline-none text-blue-500 border-b-2 font-medium border-blue-500 inline-flex"
          >{sheet.name}
          &nbsp;&nbsp;
          <span on:click={() => dispatch("remove_sheet")}>
            <Icon
              name="x-circle"
              class="w-5 pt-1 text-gray-500 hover:text-red-500"
            />
          </span>
        </button>
      {:else}
        <button
          on:click={() => dispatch("change_sheet", sheet.__id)}
          class="text-gray-600 p-2 block hover:text-blue-500 focus:outline-none"
          >{sheet.name}</button
        >
      {/if}
    {/each}

    <button
      on:click={() => dispatch("add_sheet")}
      class="m-2 p-1 rounded hover:bg-blue-200 border border-blue-200"
    >
      <Icon name="plus" class="w-4 h-4" />
    </button>
  </nav>

  <div class="flex p-1 gap-1">
    <ActionNormal
      onClick={() => dispatch("action_refresh")}
      icon="refresh"
      name="Refresh"
    />
    <ActionNormal
      onClick={() => dispatch("action_goto_rawtable")}
      icon="hashtag"
      name="Raw"
    />

    <ActionNormal
      onClick={() => dispatch("action_goto_history")}
      icon="calendar"
      name="History"
    />

    <ActionNormal
      onClick={() => dispatch("action_extra")}
      icon="dots-horizontal"
      name="Extra"
    />

    {#key selected_rows}
      {#if selected_rows.length === 1}
        <ActionNormal
          onClick={() => dispatch("action_delete_trash", selected_rows[0])}
          icon="trash"
          name="Delete"
        />
      {/if}
    {/key}

    <ActionNormal
      onClick={() => dispatch("action_search")}
      icon="document-search"
      name="Search"
    />

    <div class="h-full border mx-1" />

    {#each widgets as widget}
      <ActionNormal
        onClick={() => dispatch("action_run_widget", widget)}
        icon="puzzle"
        name={widget.name}
      />
    {/each}
  </div>

  <div id="sheet-main" class="bg-white border rounded overflow-auto relative">
    <slot />
  </div>
</div>

<div class="fixed bottom-4 z-5 right-14">
  <button
    on:click={() => dispatch("add_row")}
    class="p-0 w-8 h-8 bg-blue-500 rounded-full hover:bg-blue-800 active:shadow-lg mouse shadow transition ease-in duration-200 focus:outline-none"
  >
    <Icon name="plus" class="w-6 h-6 inline-block text-white" />
  </button>
</div>

<style>
  #sheet-main {
    height: calc(-7rem + 100vh);
  }
</style>

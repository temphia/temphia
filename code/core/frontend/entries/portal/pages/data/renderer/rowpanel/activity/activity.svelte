<script lang="ts">
  import Card from "./_card.svelte";
  import EmojiSelector from "svelte-emoji-selector";
  import UserSelector from "./_user_selector.svelte";
  import type { TableService } from "../../../../../services/data/table";

  export let table_service: TableService;
  export let rowid: number;

  const row_editor = table_service.get_row_service();

  let timelines = [];
  let loaded = false;
  $: _expanded_id = 0;

  const load = () => {
    if (rowid !== 0) {
      row_editor.list_activity(rowid).then((resp) => {
        timelines = resp.data;
        loaded = true;
      });
    }
  };

  load();

  let message = "";
</script>

{#key _expanded_id}
  {#if loaded}
    <div class="flex-grow flex flex-col h-32 p-2 space-y-1 overflow-y-auto">
      {#each timelines as item}
        <Card
          data={item}
          expanded={_expanded_id === item.id}
          onClick={() => {
            if (_expanded_id === item.id) {
              _expanded_id = 0;
            } else {
              _expanded_id = item.id;
            }
          }}
        />
      {/each}
    </div>

    <div class="flex flex-col gap-2 relative">
      <div class="flex justify-end gap-4">
        <div class="px-2 py-1 rounded-full border hover:bg-yellow-200">
          <EmojiSelector on:emoji={(ev) => {}} />
        </div>

        <div class="p-1 rounded-full border hover:bg-yellow-200">
          <UserSelector />
        </div>
      </div>

      <textarea
        bind:value={message}
        class="shadow w-full bg-gray-50 border focus:bg-gray-100 p-1"
        placeholder="comment something..."
      />
      <button
        class="rounded bg-blue-500 hover:bg-blue-300 text-white p-1"
        on:click={async () => {
          if (!message) {
            return;
          }
          await row_editor.comment_row(rowid, message);
          load();
        }}>comment</button
      >
    </div>
  {:else}
    <div>Loading..</div>
  {/if}
{/key}

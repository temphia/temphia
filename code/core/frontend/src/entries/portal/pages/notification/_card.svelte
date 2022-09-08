<script lang="ts">
  import SvgMessage from "./_svg_message.svelte";
  import SvgClose from "./_svg_close.svelte";
  import { time_ago } from "../../../../lib/vendor/timeago";

  export let delete_notif;
  export let read_notif;

  export let nmsg;
</script>

<div
  class="w-full p-3 mt-4 bg-white rounded shadow flex flex-shrink-0 cursor-pointer relative"
  on:click|stopPropagation={() => {
    if (nmsg.read) {
      return;
    }
    read_notif(nmsg.id);
  }}
>
  {#if !nmsg.read}
    <div class="absolute">
      <div class="h-2 w-2 rounded-full bg-green-400" />
    </div>
  {/if}

  <div
    tabindex="0"
    aria-label="group icon"
    role="img"
    class="focus:outline-none w-8 h-8 border rounded-full border-gray-200 flex flex-shrink-0 items-center justify-center"
  >
    <SvgMessage />
  </div>

  {#if nmsg["type"] === "user_message"}
    <div class="pl-3 w-full">
      <div class="flex items-center justify-between w-full">
        <p tabindex="0" class="focus:outline-none text-sm leading-none">
          <span class="text-indigo-700">{nmsg["from_user"] || ""}</span>
          messaged you:
          <span class="italic">{nmsg["contents"] || ""}</span>
        </p>
        <SvgClose
          onClick={() => {
            delete_notif(nmsg.id);
          }}
        />
      </div>
      {#if nmsg["created_at"]}
        <p class="focus:outline-none text-xs leading-3 pt-1 text-gray-500">
          {time_ago(nmsg["created_at"])}
        </p>
      {/if}
    </div>
  {:else}
    <div>Not implemented</div>
  {/if}
</div>

<script lang="ts">
  import Card from "./panels/_card.svelte";
  import Processing from "./panels/processing.svelte";
  import { createEventDispatcher } from "svelte";
  import { Quotes } from "$lib/services/portal/notifier/quotes";

  export let loading = true;
  export let messages: object[] = [];

  const dispatch = createEventDispatcher();

  const refresh = () => dispatch("refresh");
  const toggle_npanel = () => dispatch("toggle_npanel");
  const nread = (msg) => dispatch("nread", msg);
  const ndelete = (msg) => dispatch("ndelete", msg);

  $: _messages = messages.sort((x, y) => x["id"] - y["id"]);
</script>

<div class="p-8">
  <div class="flex items-center justify-between">
    <p
      class="focus:outline-none text-2xl font-semibold leading-6 text-gray-800"
    >
      Notifications
    </p>

    <div>
      <button
        aria-label="mark all message as read"
        class="focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 rounded-md cursor-pointer"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-6 w-6 text-gray-400"
          viewBox="0 0 20 20"
          fill="currentColor"
        >
          <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
          <path
            fill-rule="evenodd"
            d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z"
            clip-rule="evenodd"
          />
        </svg>
      </button>

      <button
        aria-label="refresh"
        on:click={refresh}
        class="focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 rounded-md cursor-pointer mx-5"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-6 w-6 text-gray-400"
          viewBox="0 0 20 20"
          fill="currentColor"
        >
          <path
            fill-rule="evenodd"
            d="M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z"
            clip-rule="evenodd"
          />
        </svg>
      </button>

      <button
        aria-label="close notification modal"
        class="focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 rounded-md cursor-pointer"
        on:click={toggle_npanel}
      >
        <svg
          class="h-6 w-6"
          viewBox="0 0 24 24"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M18 6L6 18"
            stroke="#4B5563"
            stroke-width="1.25"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
          <path
            d="M6 6L18 18"
            stroke="#4B5563"
            stroke-width="1.25"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </button>
    </div>
  </div>

  {#if loading}
    <Processing text={"loading"} />
  {:else}
    {#each _messages as nmsg}
      <Card {nmsg} delete_notif={ndelete} read_notif={nread} />
    {/each}
  {/if}

  <div class="flex items-center justiyf-between">
    <hr class="w-full" />
    <p
      class="focus:outline-none text-sm flex flex-shrink-0 leading-normal px-3 py-16 text-gray-500"
    >
      {Quotes[Math.floor(Math.random() * Quotes.length)]}
    </p>
    <hr class="w-full" />
  </div>
</div>

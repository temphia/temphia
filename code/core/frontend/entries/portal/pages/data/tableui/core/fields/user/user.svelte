<script lang="ts">
  import type { Column, RowService } from "../../../../../services/data";
  import UserPanel from "./_user_panel.svelte";

  export let value;
  export let column: Column;
  export let onChange: (value: any) => void;
  export let row_service: RowService;

  $: __value = [];

  const link_gen = (user: string) => ""; // fixme

  const unSelectUser = (userId: string) => () => {
    __value = [...__value.filter((val) => val !== userId)];
  };
</script>

<div class="flex flex-col w-full h-full p-1">
  <div
    class="flex flex-wrap w-full p-1 space-x-2 border border-dashed rounded-lg bg-gray-50"
    style="min-height: 2rem;"
  >
    {#each __value as item}
      <div class="flex flex-col relative border p-1 rounded-lg">
        <svg
          on:click={() => unSelectUser(item)}
          xmlns="http://www.w3.org/2000/svg"
          class="h-4 w-4 absolute -right-1 text-blue-500 border rounded-full bg-white cursor-pointer"
          viewBox="0 0 20 20"
          fill="currentColor"
        >
          <path
            fill-rule="evenodd"
            d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
            clip-rule="evenodd"
          />
        </svg>
        <img
          class="w-10 h-auto rounded-full"
          src={link_gen(item)}
          alt=""
        />
        <span class="text-gray-800 bg-blue-50 rounded">{item}</span>
      </div>
    {/each}
  </div>

  <div
    class="p-2 flex justify-end text-gray-600 cursor-pointer"
    on:click={() => {
      row_service.open_model(UserPanel, { profile: link_gen });
    }}
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      viewBox="0 0 20 20"
      fill="currentColor"
    >
      <path
        d="M5.5 13a3.5 3.5 0 01-.369-6.98 4 4 0 117.753-1.977A4.5 4.5 0 1113.5 13H11V9.413l1.293 1.293a1 1 0 001.414-1.414l-3-3a1 1 0 00-1.414 0l-3 3a1 1 0 001.414 1.414L9 9.414V13H5.5z"
      />
      <path d="M9 13h2v5a1 1 0 11-2 0v-5z" />
    </svg>
    <span>Pick</span>
  </div>
</div>

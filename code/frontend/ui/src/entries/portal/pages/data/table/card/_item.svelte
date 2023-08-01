<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { FolderTktAPI } from "../../../../../../lib/apiv2";
  import type { Column, Order } from "../../../../services/data";
  import Cicon from "../core/cicon/cicon.svelte";

  export let order: Order;
  export let row: object;
  export let columns: { [_: string]: Column };
  export let onEdit = () => {};
  export let folder_api: FolderTktAPI;

  const image_column = order.image ? row[order.image]  : null
  const images = image_column ? (row[order.image] || "").split(",") : [];
</script>

<div class="relative mx-auto w-full">
  <div
    class="relative inline-block duration-300 ease-in-out transition-transform transform hover:border border-blue-500 w-full"
  >
    <div class="shadow p-2 rounded-lg bg-white">
      <div class="flex justify-center relative rounded-lg overflow-hidden h-52">
        <div
          class="w-full flex gap-1 flex-wrap bg-gray-50"
        >
          {#each images as img}
            <img
              src={folder_api.getFilePreviewUrl(img)}
              class="h-20 p-1 border bg-white cursor-pointer"
              alt=""
            />
          {/each}
        </div>

        <div class="absolute flex justify-center bottom-0 mb-3">
          <div
            class="flex bg-white px-4 py-1 space-x-5 rounded-full overflow-hidden shadow hover:bg-slate-200"
          >
            <button
              class="inline-flex uppercase text-slate-600 text-sm"
              on:click={onEdit}
            >
              <Icon name="pencil-alt" solid class="h-5 w-5" />
              Edit
            </button>
          </div>
        </div>

        {#if order.tag}
          <span
            class="absolute top-0 left-0 inline-flex mt-3 ml-3 px-3 py-2 rounded-lg z-10 bg-red-500 text-sm font-medium text-white select-none"
          >
            {row[order.tag] || ""}
          </span>
        {/if}

        <span
          class="absolute top-2 right-2 py-1 px-2 rounded-full z-10 bg-blue-400 text-sm font-medium text-white select-none"
        >
          {row["__id"] || ""}
        </span>
      </div>

      <div class="mt-4">
        {#if order.name}
          <h2
            class="font-medium text-base md:text-lg text-gray-800 line-clamp-1"
            title={row[order.name]}
          >
            {row[order.name] || ""}
          </h2>
        {/if}

        {#if order.description}
          <p class="mt-2 text-sm text-gray-800 line-clamp-1">
            {row[order.description] || ""}
          </p>
        {/if}
      </div>

      <div
        class="grid grid-cols-2 gap-y-4 gap-x-2 mt-8 overflow-y-auto justify-self-start content-start"
        style="max-height: 15rem"
      >
        {#each order.other as ocol}
          {@const column = columns[ocol]}

          <p class="flex justify-start items-start text-gray-800">
            <Cicon ctype={column.ctype} />
            <span> {column.name || ""} </span>
          </p>

          <p class="flex items-start text-gray-800 flex-wrap truncate">
            {row[ocol] || ""}
          </p>
        {/each}
      </div>

      {#if order.user}
        <div class="grid grid-cols-2 mt-8">
          <div class="flex items-center">
            <div class="relative">
              <div class="rounded-full w-6 h-6 md:w-8 md:h-8 bg-gray-200" />
              <span
                class="absolute top-0 right-0 inline-block w-3 h-3 bg-primary-red rounded-full"
              />
            </div>

            <p class="ml-2 text-gray-800 line-clamp-1">
              {row[order.user] || ""}
            </p>
          </div>
        </div>
      {/if}
    </div>
  </div>
</div>

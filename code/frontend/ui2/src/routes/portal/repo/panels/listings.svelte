<script lang="ts">
  import Icon from "./icon.svelte";

  export let items: any[] = [];
  export let sources: { [_: number]: string };
  export let currentSource = "";
  export let onChangeSource;
  export let onItemSelect;
</script>

<div
  class="w-full h-full py-10 mx-auto overflow-auto bg-gradient-to-b from-purple-100 to-indigo-100"
>
  <div class="flex justify-center w-full mb-10">
    <div class="relative">
      <input
        type="text"
        class="h-14 w-40 md:w-96 pr-8 pl-5 rounded z-0 focus:shadow focus:outline-none"
        placeholder="Search anything..."
      />
      <div class="absolute top-4 right-3 flex">
        <svg class="h-5 w-5 text-gray-500" viewBox="0 0 24 24" fill="none"
          ><path
            d="M21 21L15 15M17 10C17 13.866 13.866 17 10 17C6.13401 17 3 13.866 3 10C3 6.13401 6.13401 3 10 3C13.866 3 17 6.13401 17 10Z"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          /></svg
        >
      </div>
    </div>

    <div class="ml-1 relative inline-flex">
      <svg
        class="w-2 h-2 absolute top-0 right-0 m-4 pointer-events-none"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 412 232"
        ><path
          d="M206 171.144L42.678 7.822c-9.763-9.763-25.592-9.763-35.355 0-9.763 9.764-9.763 25.592 0 35.355l181 181c4.88 4.882 11.279 7.323 17.677 7.323s12.796-2.441 17.678-7.322l181-181c9.763-9.764 9.763-25.592 0-35.355-9.763-9.763-25.592-9.763-35.355 0L206 171.144z"
          fill="#648299"
          fill-rule="nonzero"
        /></svg
      >
      <select
        value={currentSource}
        on:change={(ev) => {
          if (ev.target["value"] !== "Source") {
            onChangeSource(ev.target["value"]);
          }
        }}
        class="border border-gray-300 rounded text-gray-600 h-14 pl-2 pr-8 bg-white hover:border-gray-400 focus:outline-none appearance-none"
      >
        <option>Source</option>
        {#each Object.entries(sources) as [skey, sval]}
          <option value={skey}>{`${skey} [${sval}]`}</option>
        {/each}
      </select>
    </div>
  </div>

  <div class="flex flex-wrap justify-center gap-2 p-2">
    {#each items as item, i}
      <div
        class="p-4 2xl:w-1/6 xl:w-1/5 lg:1/4 md:w-1/3 w-full bg-white rounded-md hover:border-purple-600 border cursor-pointer relative"
        on:click={onItemSelect(item)}
      >
        <div
          class="bg-purple-400 text-white text-xs absolute rounded -right-2 -top-1 p-1"
        >
          <span>{item.type || item.group}</span>
        </div>

        <div class="h-auto max-w-full p-4 bg-gray-200">
          <Icon src={item["icon"] || ""} />
        </div>
        <div class="mt-2 h-32 w-full space-y-3">
          <h1 class="text-xl font-semibold text-gray-700">
            {item.name}
          </h1>

          <div class="flex">
            {#each item.tags || [] as tag}
              <span class="text-blue-400 p-1">#{tag} </span>
            {/each}
          </div>

          <p class="font-medium text-gray-700 mb-4 truncate overflow-clip ">
            {item.description}
          </p>
        </div>
      </div>
    {/each}
  </div>
</div>

<script lang="ts">
  export let name = "";
  export let slug = "";
  export let source = "";
  export let type = "";
  export let subtype = "";
  export let description = "";
  export let files = [];

  export let last_stage_name = "Import";

  export let last_page_func;
  export let final_func;

  let _activeTab = 0;
  let _tabs = ["About", "Options", last_stage_name];
  let _final_called = false;

  const onimport = async () => {
    _activeTab = _activeTab + 1;
    await last_page_func();
    _final_called = true;
  };

  console.warn("@@=>", document.currentScript);


</script>

<div class="flex flex-col w-full h-full bg-gray-50">
  <header class="flex-shrink h-16 border-b pb-1">
    <ul class="flex justify-center items-center">
      {#each _tabs as tab, index}
        <li
          class="cursor-pointer py-2 px-4 text-gray-500 border-b-8 {_activeTab ===
          index
            ? 'text-blue-500 border-blue-500'
            : ''}"
        >
          {tab}
        </li>
      {/each}
    </ul>
  </header>
  <div class="flex-grow h-32 overflow-auto">
    <div class="w-full p-5 overflow-auto">
      {#if _activeTab === 0}
        <div class="w-full border bg-white shadow">
          <div class="p-4 border-b">
            <h2 class="text-2xl ">Blueprint Information</h2>
            <p class="text-sm text-gray-500">Description and properties.</p>
          </div>
          <div>
            <div
              class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
            >
              <p class="text-gray-600">Name</p>
              <p>{name}</p>
            </div>
            <div
              class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
            >
              <p class="text-gray-600">Slug</p>
              <p>{slug}</p>
            </div>
            <div
              class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
            >
              <p class="text-gray-600">Type</p>
              <p>{type}</p>
            </div>
            <div
              class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
            >
              <p class="text-gray-600">Sub Type</p>
              <p>{subtype}</p>
            </div>
            <div
              class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
            >
              <p class="text-gray-600">Description</p>
              <p>
                {description}
              </p>
            </div>
            <div
              class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4"
            >
              <p class="text-gray-600">Files</p>
              <div class="space-y-2">
                {#each files || [] as file}
                  <div
                    class="border-2 flex items-center p-2 rounded justify-between space-x-2"
                  >
                    <div class="space-x-2 truncate">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="fill-current inline text-gray-500"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        ><path
                          d="M17 5v12c0 2.757-2.243 5-5 5s-5-2.243-5-5v-12c0-1.654 1.346-3 3-3s3 1.346 3 3v9c0 .551-.449 1-1 1s-1-.449-1-1v-8h-2v8c0 1.657 1.343 3 3 3s3-1.343 3-3v-9c0-2.761-2.239-5-5-5s-5 2.239-5 5v12c0 3.866 3.134 7 7 7s7-3.134 7-7v-12h-2z"
                        /></svg
                      >
                      <span> {file} </span>
                    </div>
                    <a href="#" class="text-purple-700 hover:underline">
                      Download
                    </a>
                  </div>
                {/each}
              </div>
            </div>
          </div>
        </div>
      {:else if _activeTab === 1}
        <slot name="options" />
      {:else if _activeTab === 2}
        <slot name="final" />
      {/if}
    </div>
  </div>
  <footer
    class="flex-shrink h-16 flex bg-white gap-4 justify-center border-t p-2"
  >
    {#if _activeTab === 0}
      <button
        on:click={() => (_activeTab = _activeTab + 1)}
        class="py-2 px-4 border rounded-md cursor-pointer uppercase text-sm font-bold bg-blue-500 text-white hover:bg-blue-700"
        >Next</button
      >
    {:else if _activeTab === 1}
      <button
        on:click={() => (_activeTab = _activeTab - 1)}
        class="py-2 px-4 border rounded-md cursor-pointer uppercase text-sm font-bold bg-blue-500 text-white hover:bg-blue-700"
        >Back</button
      >

      <button
        on:click={onimport}
        class="py-2 px-4 border rounded-md cursor-pointer uppercase text-sm font-bold bg-blue-500 text-white hover:bg-blue-700"
        >Finish</button
      >
    {:else if _activeTab === 2}
      {#if _final_called}
        <button
          on:click={final_func}
          class="py-2 px-4 border rounded-md cursor-pointer uppercase text-sm font-bold bg-blue-500 text-white hover:bg-blue-700"
          >Close</button
        >
      {/if}
    {/if}
  </footer>
</div>

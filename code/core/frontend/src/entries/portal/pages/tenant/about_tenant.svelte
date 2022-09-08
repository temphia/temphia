<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../lib/app/portal";

  export let show_edit = false;

  const app: PortalApp = getContext("__app__");
</script>

<div class="h-full w-full overflow-auto">
  <div class="md:p-12 bg-indigo-100 flex flex-row flex-wrap">
    <div
      class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg relative"
    >
      <div class="text-2xl text-indigo-900">Organization</div>

      {#if show_edit}
        <div class="absolute right-0 p-2">
          <button
            on:click={app.navigator.goto_admin_org_edit}
            class="p-1 rounded hover:scale-110 text-white bg-blue-500"
            >Edit</button
          >
        </div>
      {/if}

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Name</label>
        <input
          type="text"
          disabled
          value="xyz"
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Slug</label>
        <input
          type="text"
          value="mno"
          disabled
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Bio</label>
        <textarea value={"this is blah, blah "} disabled class="border" />
      </div>
    </div>
  </div>
</div>

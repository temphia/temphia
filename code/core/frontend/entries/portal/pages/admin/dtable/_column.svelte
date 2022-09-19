<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../app";
  import Kveditor from "../../../../_shared/common/kveditor.svelte";

  import { CtypeConvertables } from "../../dtable/renderer/fields/field";
  import { DynAdminAPI } from "./dtable2";

  export let data;

  export let source;
  export let group;
  export let table;

  const app: PortalApp = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  let getExtraMeta;
  let extra_meta_modified;

  $: _unsafe_type_change = false;

  let _ctype = data["ctype"];

  let mod_data = {};

  const set = (name: string) => (ev) => {
    mod_data[name] = ev.target.value;
  };
</script>

<div class="h-full w-full overflow-auto">
  <div class="md:p-12 bg-indigo-100 flex flex-row flex-wrap">
    <div
      class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg"
    >
      <div class="text-2xl text-indigo-900">Column</div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Name</label>
        <input
          type="text"
          value={data["name"] || ""}
          on:change={set("name")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Description</label>
        <textarea
          type="text"
          value={data["description"] || ""}
          on:change={set("description")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          placeholder="About this data group..."
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <div class="absolute right-0">
          <label>
            <input type="checkbox" bind:checked={_unsafe_type_change} />
            Allow unsafe type change
          </label>
        </div>

        <label class="pb-2 text-gray-700 font-semibold">Type</label>
        <select
          class="p-2"
          value={_ctype}
          on:change={(ev) => {
            mod_data["type"] = ev.target["value"];
          }}
        >
          {#if _unsafe_type_change}
            {#each Object.keys(CtypeConvertables) as item}
              <option>{item}</option>
            {/each}
          {:else}
            <option>{_ctype}</option>
            {#each CtypeConvertables[_ctype] || [] as item}
              <option>{item}</option>
            {/each}
          {/if}
        </select>
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Options</label>
        <div class="flex flex-col bg">
          {#each data["options"] || [] as opt}
            <div class="p-2 border ">{opt}</div>
          {/each}
        </div>
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Pattern</label>
        <input
          type="text"
          value={data["pattern"] || ""}
          on:change={set("pattern")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Ref type</label>
        <input
          type="text"
          value={data["ref_type"] || ""}
          on:change={set("ref_type")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Ref target</label>
        <input
          type="text"
          value={data["ref_target"] || ""}
          on:change={set("ref_type")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Ref object</label>
        <input
          type="text"
          value={data["ref_object"] || ""}
          on:change={set("ref_object")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Strict Pattern</label>
        <input
          type="checkbox"
          value={data["strict_pattern"] || false}
          on:change={(ev) => {
            mod_data["strict_pattern"] = ev.target["checked"];
          }}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
        <Kveditor
          data={data["extra_meta"] || {}}
          bind:getData={getExtraMeta}
          bind:modified={extra_meta_modified}
        />
      </div>

      <div
        class="flex justify-between space-x-1"
        on:click={async () => {
          const _data = { ...mod_data };
          if (extra_meta_modified) {
            _data["extra_meta"] = getExtraMeta();
          }

          dynapi.edit_column(source, group, table, data["slug"], _data);
          app.navigator.goto_admin_dtable_page(source, group, table);
        }}
      >
        <button class="p-2 bg-blue-400 text-white rounded">Save</button>
      </div>
    </div>
  </div>
</div>

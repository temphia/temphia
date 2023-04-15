<script lang="ts">
  import Kveditor from "../../../../../../xcompo/common/kveditor.svelte";
  import type { SheetService } from "../../../../../services/data";
  import { LoadingSpinner, MultiText } from "../../../../admin/core";

  import {
    Sheet,
    SheetColTypeMultiSelect,
    SheetColTypeReference,
    SheetColTypes,
    SheetColTypeSelect,
    SheetColTypeText,
    SheetCtypeShapes,
  } from "../../sheets";
  import Layout from "./_layout.svelte";

  export let sheets: Sheet[];
  export let sheetid;
  export let onAdd = (opts: {
    name: string;
    ctype: string;
    extraopts: object;
  }) => {};

  export let service: SheetService;

  let name = "";
  let ctype = SheetColTypeText;
  let extraopts = {};
  let options_value = "";

  let refsheet = "";
  let refcolumn = "";
  let remotehook = "";

  let refcolumn_loading = false;
  let refcols = [];
  const loadRefColumns = async (_sid) => {
    if (!_sid) {
      return;
    }

    refcolumn_loading = true;
    const resp = await service.api.list_columns(_sid);
    if (!resp.ok) {
      return;
    }
    refcols = resp.data;
    refcolumn_loading = false;
  };

  const doOnAdd = async () => {
    const data = { name, ctype, extraopts };
    if (ctype === SheetColTypeReference) {
      data["refsheet"] = Number(refsheet);
      data["refcolumn"] = Number(refcolumn);
    }

    if (options_value) {
      data["opts"] = options_value;
    }

    return onAdd(data);
  };
</script>

<Layout title="New Column" onClick={doOnAdd}>
  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="name"
      >Name</label
    >
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
      id="name"
      type="text"
      bind:value={name}
      placeholder="name"
    />
  </div>

  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="type"
      >Type</label
    >

    <select
      id="type"
      bind:value={ctype}
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
    >
      {#each SheetColTypes as st}
        <option value={st}>{st}</option>
      {/each}
    </select>
  </div>

  {#if ctype === SheetColTypeReference}
    <div class="mb-4">
      <label class="block mb-2 text-sm font-bold text-gray-700" for="refsheet"
        >Ref Sheet</label
      >

      <select
        id="refsheet"
        bind:value={refsheet}
        on:change={(ev) => {
          refcolumn = "";
          refcols = [];
          loadRefColumns(ev.target["value"]);
        }}
        class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
      >
        {#each sheets as sheet}
          {#if sheetid != sheet.__id}
            <option value={sheet.__id}>{sheet.name}</option>
          {/if}
        {/each}
      </select>
    </div>

    {#if refcolumn_loading}
      <LoadingSpinner classes="" />
    {:else}
      <div class="mb-4">
        <label
          class="block mb-2 text-sm font-bold text-gray-700"
          for="refcolumn">Ref Column</label
        >

        <select
          id="refcolumn"
          bind:value={refcolumn}
          class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
        >
          {#if ctype === SheetColTypeReference}
            {#each refcols as rf}
              {#if SheetCtypeShapes["number"].includes(ctype)}
                <option value={rf.__id}>{rf.name}</option>
              {/if}
            {/each}
          {/if}
        </select>
      </div>
    {/if}
  {/if}

  {#if ctype === SheetColTypeMultiSelect || ctype === SheetColTypeSelect}
    <div class="mb-4">
      <label class="block mb-2 text-sm font-bold text-gray-700" for="opts"
        >Options</label
      >
      <MultiText bind:value={options_value} />
    </div>
  {/if}

  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="extraopts"
      >Extra Options</label
    >
    <Kveditor />
  </div>
</Layout>

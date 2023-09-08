<script lang="ts">
  import Kveditor from "$lib/compo/common/kveditor.svelte";
  import type { SheetService } from "$lib/services/data";
  import type { SheetColumn } from "../../sheets";
  import Layout from "./_layout.svelte";

  export let column: SheetColumn;
  export let service: SheetService;

  let kvmodified = false;
  let kvGetData;

  let modified = false;

  const newdata = {};

  const setValue = (field) => (ev) => {
    newdata[field] = ev.target["value"];
    modified = true;
  };

  const setNumValue = (field) => (ev) => {
    newdata[field] = Number(ev.target["value"]);
    modified = true;
  };
</script>

<Layout
  title="Edit Column {column.__id}"
  onSave={async () => {
    await service.update_column(String(column.__id), newdata);
    service.close_small_modal();
    service.init();
  }}
  onDelete={async () => {
    await service.remove_column(String(column.__id));
    service.close_small_modal();
    service.init();
  }}
>
  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="name"
      >Name</label
    >
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
      id="name"
      type="text"
      on:change={setValue("name")}
      value={column.name}
    />
  </div>

  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="options"
      >Options
    </label>
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
      id="options"
      type="text"
      on:change={setValue("opts")}
      value={column.opts || ""}
    />
  </div>

  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="color"
      >Color
    </label>
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
      id="color"
      type="color"
      on:change={setValue("color")}
      value={column.color || ""}
    />
  </div>

  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="refsheet"
      >Ref Sheet
    </label>
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
      id="refsheet"
      type="number"
      on:change={setNumValue("refsheet")}
      value={column.refsheet || 0}
    />
  </div>

  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="refcol"
      >Ref Column
    </label>
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
      id="refcol"
      type="number"
      on:change={setNumValue("refcolumn")}
      value={column.refcolumn || 0}
    />
  </div>

  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="remotehook"
      >Remote Hook
    </label>
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
      id="remotehook"
      type="number"
      on:change={setNumValue("remotehook")}
      value={column.remotehook || 0}
    />
  </div>

  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="extraopts"
      >Extra Options</label
    >
    <Kveditor
      bind:modified={kvmodified}
      data={column.extraopts || {}}
      bind:getData={kvGetData}
    />
  </div>
</Layout>

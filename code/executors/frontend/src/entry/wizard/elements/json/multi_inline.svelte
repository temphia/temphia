<script lang="ts">
  import { getContext } from "svelte";
  import type { ModalControl } from "../../core";


  import JsonTable from "./_json_table.svelte";
  import Nested from "./_nested.svelte";

  export let field: object = {};
  export let data_source: any = {};
  export let data: any;
  export let field_store: any;

  $: _data = data ? data : data_source[field["source"]] || [];

  const mhandle: ModalControl = getContext("__modal__");

  const onRemove = (data: any) => {
    // console.log("CONSOLE @=>", data);
  };

  const options = field["options"] || {};
  const nested_message = options["message"] || "";
  const nested_fields = options["fields"] || [];
  const column_names = nested_fields.map((val) => val["name"]);
</script>

<div class="flex justify-between p-1">
  <div>{field["name"] || ""}</div>

  <button
    class="p-1 bg-blue-400 rounded text-white hover:bg-blue-500 flex text-sm"
    on:click={() => {
      mhandle.show_big(Nested, {
        message: nested_message,
        fields: nested_fields,
      });
    }}
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      viewBox="0 0 20 20"
      fill="currentColor"
    >
      <path
        fill-rule="evenodd"
        d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z"
        clip-rule="evenodd"
      />
    </svg>
    Add</button
  >
</div>

<JsonTable
  key={"__id"}
  {onRemove}
  datas={_data}
  options={{
    column_names,
  }}
/>

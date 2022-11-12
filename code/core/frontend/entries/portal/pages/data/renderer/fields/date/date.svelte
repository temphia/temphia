<script lang="ts">
  import Flatpickr from "svelte-flatpickr";
  import type { Column } from "../../manager/dtypes";

  export let column: Column;
  export let onChange: (val: string) => void;
  export let value: string;

  $: _id = `row-edit-${column.slug}`;
  $: _element = `#${_id}`;
  $: _date = value;

  const defaultOpts = {
    element: _element,
    enableTime: true,
    altInput: true,
    wrap: true,
    altFormat: "F j, Y [ H:i ]",
    dateFormat: "Y-m-d",
    parseDate: (dstr) => (dstr ? new Date(dstr) : new Date()),
  };
</script>

<svelte:head>
  <link
    rel="stylesheet"
    type="text/css"
    href="https://npmcdn.com/flatpickr/dist/themes/material_green.css"
  />
</svelte:head>

<Flatpickr
  options={defaultOpts}
  value={_date}
  element={_element}
  on:change={(ev) => {
    const _new_date = ev.detail[0][0].toISOString();
;
    onChange(_new_date);
  }}
>
  <div class="mb-5 w-full ">
    <div class="flatpickr relative" id={_id}>
      <input
        type="text"
        placeholder="Select Date.."
        data-input
        class="p-2 shadow w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      />

      <div class="absolute top-0 right-0 px-3 py-2">
        <svg
          class="h-6 w-6 text-gray-400"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0
                00-2 2v12a2 2 0 002 2z"
          />
        </svg>
      </div>
    </div>
  </div>
</Flatpickr>

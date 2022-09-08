<script lang="ts">
  import Flatpickr from "svelte-flatpickr";

  export let date: string;
  export let id: string;
  export let onChange: (val: string) => void;
  export let flatpickrOpts: object;

  $: _element = `#${id}`;

  const defaultOpts = {
    element: _element,
    enableTime: true,
    altInput: true,
    wrap: true,
    altFormat: "F j, Y [ H:i ]",
    dateFormat: "Y-m-d",
    parseDate: (dstr) => (dstr ? new Date(dstr) : new Date()),
  };

  $: _final_opts = flatpickrOpts
    ? { ...defaultOpts, ...flatpickrOpts }
    : defaultOpts;

  $: console.log("@=>>>", date);
</script>

<svelte:head>
  <link
    rel="stylesheet"
    href="https://cdn.jsdelivr.net/npm/flatpickr/dist/flatpickr.min.css"
  />
  <link
    rel="stylesheet"
    type="text/css"
    href="https://npmcdn.com/flatpickr/dist/themes/material_green.css"
  />
</svelte:head>

<Flatpickr
  options={_final_opts}
  value={date}
  element={_element}
  on:change={(ev) => {
    console.log(ev);
    const _new_date = ev.detail[0][0].toString()
    if (date!== _new_date) {
        return
    }
    onChange(_new_date)
  }}
>
  <div class="mb-5 w-full ">
    <div class="flatpickr relative" {id}>
      <input
        type="text"
        placeholder="Select Date.."
        data-input
        class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1"
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

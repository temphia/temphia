<script lang="ts">
  import { getContext } from "svelte";
  import LocationSelect from "./_location_panel.svelte";
  export let value: any;
  export let column: object;
  export let onChange: (_value: any) => void;

  const { open, close } = getContext("simple-modal");

  $: __value = value;
  const callback = (_lat: number, _lng: number) => {
    __value = [_lat, _lng];
    close();

    if (value === __value) {
      return;
    }

    onChange(__value);
  };

  const selectLocation = () => {
    const props = {
      callback,
    };
    if (__value && Array.isArray(__value)) {
      props["lat"] = __value[0];
      props["lng"] = __value[1];
    }
    open(LocationSelect, props);
  };
</script>

<div class="flex">
  <input
    type="text"
    class="p-2 shadow w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    placeholder="Select a Place"
    disabled
    value={__value}
  />

  <span class="p-2 cursor-pointer text-gray-600" on:click={selectLocation}>
    <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"
      />
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"
      />
    </svg>
  </span>
</div>

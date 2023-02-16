<script lang="ts">
  import {
    FieldShortText,
    FieldLongText,
    FieldEmail,
    FieldCheckbox,
    FieldNumber,
    FieldPhone,
    FieldSelect,
    FieldRange,
    FieldColor,
  } from "./field";

  export let name: string;
  export let info: string;
  export let type: string;
  export let options: string[];
  export let html_attr: { [_: string]: string } = {};
  export let onChange: (newValye) => void;
  export let value: any;

  const change = (ev) => {
    onChange(ev.target["value"]);
  };
</script>

<label for={name} class="pb-2 text-gray-700 flex uppercase"> {name}</label>

{#if type === FieldShortText}
  <div class="flex w-full">
    <input
      type="text"
      id={name}
      on:change={change}
      value={value || ""}
      class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      {...html_attr}
    />
  </div>
{:else if type === FieldPhone}
  <div class="flex w-full">
    <input
      type="tel"
      id={name}
      on:change={change}
      value={value || ""}
      class="p-2 w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      {...html_attr}
    />
  </div>
{:else if type === FieldSelect}
  <div class="flex w-full">
    <select
      class="w-full p-2 bg-gray-50 border"
      value={value || ""}
      on:change={change}
      {...html_attr}
    >
      {#each options || [] as opt}
        <option value={opt}>{opt}</option>
      {/each}
    </select>
  </div>
{:else if type === FieldCheckbox}
  <div class="flex w-full">
    <input
      type="checkbox"
      id={name}
      on:change={null}
      checked={false}
      class="form-checkbox h-5 w-5 text-gray-600"
      {...html_attr}
    />
  </div>
{:else if type === FieldNumber}
  <div class="flex w-full">
    <input
      type="number"
      id={FieldNumber}
      on:change={() => {}}
      value={value || ""}
      class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      {...html_attr}
    />
  </div>
{:else if type === FieldLongText}
  <div class="flex w-full">
    <textarea
      id={name}
      on:change={change}
      value={value || ""}
      class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      placeholder="write something..."
      {...html_attr}
    />
  </div>
{:else if type === FieldEmail}
  <div class="flex w-full">
    <input
      type="email"
      id={name}
      on:change={change}
      value={value || ""}
      class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      placeholder="mail@example.com"
      {...html_attr}
    />
  </div>
{:else if type === FieldRange}
  <div class="flex w-full">
    <input
      type="range"
      id={name}
      on:change={change}
      value={value || ""}
      class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      {...html_attr}
    />
  </div>
{:else if type === FieldColor}
  <div class="flex w-full">
    <input
      type="color"
      id={name}
      on:change={change}
      value={value || ""}
      class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      {...html_attr}
    />
  </div>
{:else}
  <div>Not Implemented</div>
{/if}

<p class="text-sm italic mb-4">{info}</p>

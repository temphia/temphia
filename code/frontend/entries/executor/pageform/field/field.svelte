<script lang="ts">
  import * as f from "./field";

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

  const change_num = (ev) => {
    onChange(Number(ev.target["value"]));
  };
</script>

<label for={name} class="pb-2 text-gray-700 flex uppercase"> {name}</label>

<div class="flex flex-col w-full">
  {#if type === f.FieldShortText}
    <input
      type="text"
      id={name}
      on:change={change}
      value={value || ""}
      class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      {...html_attr}
    />

    <p class="text-sm italic mb-4">{info}</p>
  {:else if type === f.FieldPhone}
    <input
      type="tel"
      id={name}
      on:change={change}
      value={value || ""}
      class="p-2 w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      {...html_attr}
    />

    <p class="text-sm italic mb-4">{info}</p>
  {:else if type === f.FieldSelect}
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

    <p class="text-sm italic mb-4">{info}</p>
  {:else if type === f.FieldCheckbox}
    <input
      type="checkbox"
      id={name}
      on:change={null}
      checked={false}
      class="form-checkbox h-5 w-5 text-gray-600"
      {...html_attr}
    />
    <p class="text-sm italic mb-4">{info}</p>
  {:else if type === f.FieldNumber}
    <input
      type="number"
      id={name}
      on:change={change_num}
      value={value || ""}
      class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      {...html_attr}
    />
    <p class="text-sm italic mb-4">{info}</p>
  {:else if type === f.FieldLongText}
    <textarea
      id={name}
      on:change={change}
      value={value || ""}
      class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      placeholder="write something..."
      {...html_attr}
    />

    <p class="text-sm italic mb-4">{info}</p>
  {:else if type === f.FieldEmail}
    <input
      type="email"
      id={name}
      on:change={change}
      value={value || ""}
      class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      placeholder="mail@example.com"
      {...html_attr}
    />

    <p class="text-sm italic mb-4">{info}</p>
  {:else if type === f.FieldRange}
    <input
      type="range"
      id={name}
      on:change={change_num}
      value={value || ""}
      class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      {...html_attr}
    />
    <p class="text-sm italic mb-4">{info}</p>
  {:else if type === f.FieldColor}
    <input
      type="color"
      id={name}
      on:change={change}
      value={value || ""}
      class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      {...html_attr}
    />
    <p class="text-sm italic mb-4">{info}</p>
  {:else if type === f.FieldColor}
    <input
      type="color"
      id={name}
      on:change={change}
      value={value || ""}
      class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      {...html_attr}
    />

    <p class="text-sm italic mb-4">{info}</p>
  {:else if type === f.FieldInfo}
    <div class="rounded bg-yellow-100 w-full p-2" {...html_attr}>
      <p>{value || info || ""}</p>
    </div>
  {:else}
    <div>Not Implemented</div>
  {/if}
</div>

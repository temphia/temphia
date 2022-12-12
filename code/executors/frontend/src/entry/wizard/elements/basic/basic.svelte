<script lang="ts">
  import type { FieldStore } from "../../service/wizard_types";
  import * as Elem from "../../service/wizard_types";
  import EmojiSelector from "svelte-emoji-selector";

  export let field: object;
  export let data_source: any;
  export let data: any;
  export let field_store: FieldStore;
  export let error: string = undefined;

  let value = data === undefined ? "" : data;

  let type = field["type"];
  let name = field["name"];
  let attrs = field["attrs"] || {};

  let hide_emojipicker = false;
  if (attrs["hide_emojipicker"]) {
    hide_emojipicker = true;
    delete attrs["hide_emojipicker"];
  }

  const change = (ev) => {
    value = ev.target.value;
    field_store.set_value(value);
  };

  const changeBool = (ev) => {
    value = ev.target.checked;
    field_store.set_value(value);
  };

  const changeNum = (ev) => {
    value = ev.target.checked;
    field_store.set_value(Number(value));
  };

  const onEmoji = (ev) => {
    value = value ? value + ev.detail : ev.detail;
    field_store.set_value(value);
  };

  let mdata = data || [];
  const changeMultiSelect = (opt) => () => {
    if (mdata.includes(opt)) {
      mdata = [...mdata.filter((v) => v !== opt)];
    } else {
      mdata = [...mdata, opt];
    }
    field_store.set_value(mdata);
  };

  $: console.log("@mdata_ping_pong =>", mdata);

  const validate = (ev) => {};
</script>

<label for={name} class="text-gray-800 text-base font-bold leading-tight tracking-normal">{name}</label>
{#if type === Elem.BASIC_SHORTTEXT}
  <div class="flex w-full mb-5 mt-2">
    <input
      type="text"
      id={name}
      on:change={change}
      value={value || ""}
      class="text-gray-600 focus:outline-none focus:border focus:border-indigo-700 font-normal w-full h-10 flex items-center pl-3 text-sm border-gray-300 rounded border relative mr-2"
      placeholder=""
      {...attrs}
    />
    <!-- FIXME => copy this css style to below elements -->
    {#if !hide_emojipicker}
      <EmojiSelector on:emoji={onEmoji} />
    {/if}
  </div>
{:else if type === Elem.BASIC_LONGTEXT}
  <div class="flex w-full">
    <textarea
      on:change={change}
      value={value || ""}
      class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
      placeholder="write something..."
      {...attrs}
    />
    {#if !hide_emojipicker}
      <EmojiSelector on:emoji={onEmoji} />
    {/if}
  </div>
{:else if type === Elem.BASIC_RANGE}
  <input
    type="range"
    id={name}
    on:change={change}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_SELECT}
  <div class="flex w-full mt-5">
    <select class="w-full p-2" {value} on:change={change}>
      {#each data_source || [] as opt}
        <option value={opt}>{opt}</option>
      {/each}
    </select>
  </div>
{:else if type === Elem.BASIC_MULTI_SELECT}
  <div class="flex flex-col w-full h-full p-1 overflow-auto">
    <div
      class="flex flex-col p-1 space-y-1 border border-dashed rounded-lg bg-gray-100 text-gray-800"
      style="min-height: 2rem;"
      {...attrs}
    >
      {#each data_source || [] as opt}
        <label>
          <input
            checked={mdata.includes(opt)}
            type="checkbox"
            on:change={changeMultiSelect(opt)}
            class="form-checkbox h-5 w-5 text-gray-600"
          />
          {opt}
        </label>
      {/each}
    </div>
  </div>
{:else if type === Elem.BASIC_PHONE}
  <input
    type="tel"
    id={name}
    on:change={change}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_CHECKBOX}
  <input
    type="checkbox"
    id={name}
    on:change={changeBool}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_COLOR}
  <input
    type="color"
    id={name}
    on:change={change}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_DATE}
  <input
    type="date"
    id={name}
    on:change={change}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_DATETIME}
  <input
    type="datetime-local"
    id={name}
    on:change={change}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_EMAIL}
  <input
    type="email"
    id={name}
    on:change={change}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_NUMBER}
  <input
    type="number"
    id={name}
    on:change={changeNum}
    value={value || ""}
    class="p-2 shadow w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1 mt-5"
    {...attrs}
  />
{:else if type === Elem.BASIC_PARAGRAPH}
  <p>{data_source}</p>
{/if}

{#if error}
  <span class="font-sans text-sm italic text-red-500">{error}</span>
{:else}
  <span class="font-sans text-sm italic">{field["info"] || ""}</span>
{/if}

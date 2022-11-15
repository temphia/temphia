<script lang="ts">
  import EmojiSelector from "svelte-emoji-selector";

  import File from "./file/file.svelte";
  import Location from "./location/location.svelte";
  import DateField from "./date/date.svelte";
  import MultiSelect from "./multiselect/multiselect.svelte";
  import User from "./user/user.svelte";
  import Json from "./json/json.svelte";
  import RefPrimary from "./reference/ref_primary.svelte";
  import RefText from "./reference/ref_text.svelte";

  import {
    CtypeShortText,
    CtypePhone,
    CtypeSelect,
    CtypeRFormula,
    CtypeFile,
    CtypeMultiFile,
    CtypeCheckBox,
    CtypeCurrency,
    CtypeNumber,
    CtypeLocation,
    CtypeDateTime,
    // new
    CtypeMultSelect,
    CtypeLongText,
    CtypeSingleUser,
    CtypeMultiUser,
    CtypeEmail,
    CtypeJSON,
    CtypeRangeNumber,
    CtypeColor,
    RefHardPriId,
    RefSoftPriId,
    RefSoftText,
    RefHardText,
  } from "./field";
  import type {
    Column,
    DataService,
    RowService,
  } from "../../../../services/data";

  export let row: object;
  export let column: Column;
  export let onChange: (value: any) => void;
  export let row_service: RowService;

  let dirty_store = row_service.state.dirty_store;

  $: _value = $dirty_store.data[column.slug] || row[column.slug] || "";

  const change = (ev) => {
    _value = ev.target.value;
    onChange(_value);
  };

  const changeBool = (ev) => {
    _value = ev.target.checked;
    onChange(_value);
  };

  const changeNum = (ev) => {
    _value = Number(ev.target.value);
    onChange(_value);
  };

  function onEmoji(ev) {
    _value += ev.detail;
    onChange(_value);
  }

  console.log("EDITing ", row, column);
</script>

<label for={column.slug} class="pb-2 text-gray-700 uppercase"
  >{column.name}</label
>

{#if column.ref_type === RefHardPriId || column.ref_type === RefSoftPriId}
  <RefPrimary {column} {onChange} value={_value} {row_service} />
{:else if column.ref_type === RefSoftText || column.ref_type === RefHardText}
  <RefText />
{:else if column.ctype === CtypeShortText}
  <div class="flex w-full">
    <input
      type="text"
      id={column.slug}
      on:change={change}
      value={_value}
      class="p-2 shadow w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      placeholder=""
    />
    <EmojiSelector on:emoji={onEmoji} />
  </div>
{:else if column.ctype === CtypePhone}
  <div class="flex w-full">
    <input
      type="tel"
      id={column.slug}
      on:change={change}
      value={_value}
      class="p-2 shadow w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      placeholder=""
    />
  </div>
{:else if column.ctype === CtypeSelect}
  <div class="flex w-full">
    <select
      class="w-full p-2 bg-gray-50 border"
      value={_value}
      on:change={change}
    >
      {#each column.options || [] as opt}
        <option value={opt}>{opt}</option>
      {/each}
    </select>
  </div>
{:else if column.ctype === CtypeRFormula}
  <div class="flex w-full">
    <input
      type="text"
      id={column.slug}
      on:change={change}
      value={_value}
      class="p-2 shadow w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      placeholder=""
    />
  </div>
{:else if column.ctype === CtypeCheckBox}
  <div class="flex w-full">
    <input
      type="checkbox"
      id={column.slug}
      on:change={changeBool}
      checked={row[column.slug] === undefined ? false : row[column.slug]}
      class="form-checkbox h-5 w-5 text-gray-600"
    />
  </div>
{:else if column.ctype === CtypeCurrency || column.ctype === CtypeNumber}
  <div class="flex w-full">
    <input
      type="number"
      id={column.slug}
      on:change={changeNum}
      value={_value}
      class="p-2 shadow w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      placeholder=""
    />
  </div>
{:else if column.ctype === CtypeFile}
  <File multi={false} value={_value} {column} {onChange} {row_service} />
{:else if column.ctype === CtypeMultiFile}
  <File multi={true} value={_value} {column} {onChange} {row_service} />
{:else if column.ctype === CtypeLocation}
  <Location {column} {onChange} value={_value} {row_service} />
{:else if column.ctype === CtypeDateTime}
  <!-- <Reference {column} {onChange} {value} {manager} /> -->

  <DateField {column} {onChange} value={_value} />
{:else if column.ctype === CtypeMultSelect}
  <MultiSelect {column} {onChange} value={_value} />
{:else if column.ctype === CtypeLongText}
  <div class="flex w-full">
    <textarea
      id={column.slug}
      on:change={change}
      value={_value}
      class="p-2 shadow w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      placeholder="write something..."
    />
  </div>
{:else if column.ctype === CtypeSingleUser}
  <User {column} {onChange} value={_value} />
{:else if column.ctype === CtypeMultiUser}
  <User {column} {onChange} value={_value} />
{:else if column.ctype === CtypeEmail}
  <div class="flex w-full">
    <input
      type="email"
      id={column.slug}
      on:change={change}
      value={_value}
      class="p-2 shadow w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
      placeholder="mail@example.com"
    />
  </div>
{:else if column.ctype === CtypeJSON}
  <Json {column} value={_value} {row_service} />
{:else if column.ctype === CtypeRangeNumber}
  <div class="flex w-full">
    <input
      type="range"
      id={column.slug}
      on:change={change}
      value={_value}
      class="p-2 shadow w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    />
  </div>
{:else if column.ctype === CtypeColor}
  <div class="flex w-full">
    <input
      type="color"
      id={column.slug}
      on:change={change}
      value={_value}
      class="p-2 shadow w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    />
  </div>
{:else}
  <div>Not Implemented</div>
{/if}

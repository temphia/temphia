<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import MapPanel from "./_map_panel.svelte";
  import FilePanel from "./_file_panel.svelte";
  import UserAvatar from "./_user_avatar.svelte";

  import {
    SheetColTypeBoolean,
    SheetColTypeDate,
    SheetColTypeFile,
    SheetColTypeLocation,
    SheetColTypeLongText,
    SheetColTypeMultiSelect,
    SheetColTypeNumber,
    SheetColTypeRatings,
    SheetColTypeReference,
    SheetColTypeRemote,
    SheetColTypeSelect,
    SheetColTypeUser,
    SheetColumn,
    SheetCtypeIcons,
  } from "../sheets";
  import CellActions from "./_cell_actions.svelte";
  import Reference from "./_reference.svelte";
  import Remote from "./_remote.svelte";
  import ColorPanel from "./_color_panel.svelte";
  import Point from "./_point.svelte";
  import type { SheetService } from "../../../../services/data";
  import Selectable from "./_selectable.svelte";
  import UserPanel from "./_user_panel.svelte";

  export let column: SheetColumn;
  export let open_column;
  export let celldata = {};
  export let onCellChange = (data) => {};
  export let service: SheetService;

  const folder_api = service.group.folder_api;
  const id = `cell-${column.__id}`;
  let value = celldata["value"] || "";
  let value_num = celldata["numval"] || 0;
  let color = celldata["color"];
  let color_open = false;

  $: _is_open = open_column === column.__id;

  const toggle = () => {
    if (_is_open) {
      open_column = null;
    } else {
      open_column = column.__id;
      color_open = false;
    }
  };

  const close = () => {
    if (_is_open) {
      open_column = null;
    }
  };

  const picker_icons = {
    [SheetColTypeLocation]: "location-marker",
    [SheetColTypeFile]: "photograph",
    [SheetColTypeReference]: "paper-clip",
    [SheetColTypeRemote]: "external-link",
    [SheetColTypeUser]: "users",
  };

  const onColorChange = (ev) => {
    color = ev.target["value"];
    onCellChange({
      color,
    });
  };
</script>

<div
  class="py-2 border-b pl-2 border-l-4 border-l-white rounded   border-l-{color}-400"
  style="border-left-color: {color};"
>
  <label class="mb-2 text-sm font-bold text-gray-700 uppercase" for={id}>
    <span class="inline-flex">
      <Icon
        name={SheetCtypeIcons[column.ctype]}
        class="h-5 w-5 mr-1 text-gray-500"
        solid
      />
      {column.name || `Column ${column.__id}`}
    </span>
  </label>

  {#if column.ctype === SheetColTypeLongText}
    <textarea
      {id}
      {value}
      on:change={(ev) => onCellChange({ value: ev.target["value"] })}
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline hover:border-blue-400"
    />
  {:else if column.ctype === SheetColTypeBoolean}
    <label
      class="text-teal-500 inline-flex justify-between items-center hover:bg-gray-200 px-2 py-2 rounded-lg cursor-pointer hover:border-blue-400"
    >
      <input
        {id}
        class="shadow"
        type="checkbox"
        checked={value == "true"}
        on:change={(ev) => {
          let v = "false";
          if (ev.target["checked"]) {
            v = "true";
          }
          onCellChange({
            value: v,
          });
        }}
      />
    </label>
  {:else if column.ctype === SheetColTypeNumber}
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline hover:border-blue-400"
      {id}
      value={Number(value_num)}
      on:change={(ev) => {
        onCellChange({
          numval: Number(ev.target["value"]),
        });
      }}
      type="number"
    />
  {:else if column.ctype === SheetColTypeDate}
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline hover:border-blue-400"
      {id}
      {value}
      on:change={(ev) => {
        console.log("@value", ev.target["value"]);
      }}
      type="datetime-local"
    />
  {:else if column.ctype === SheetColTypeFile}
    <div
      class="p-1 rounded bg-gray-50 hover:border-blue-400 flex gap-1 flex-wrap"
      style="min-height: 2rem;"
    >
      {#if value}
        {#each value.split(",") as cd}
          <div class="relative">
            <button
              on:click={() => {
                const valueArray = value.split(",");
                value = valueArray.filter((v) => v !== cd).join();

                onCellChange({
                  value,
                });
              }}
              class="text-slate-700 -top-2 -right-2 absolute hover:text-red-600 border rounded-full bg-white"
            >
              <Icon solid name="x" class="w-4 h-4" />
            </button>

            <img
              class="h-8 w-auto"
              src={folder_api && folder_api.getFilePreviewUrl(cd)}
              alt=""
            />
          </div>
        {/each}
      {/if}
    </div>
  {:else if column.ctype === SheetColTypeRatings}
    <div class="flex p-1 gap-1">
      {#each [1, 2, 3, 4, 5] as rt}
        <button
          on:click={() => {
            if (value_num === rt) {
              value_num = 0;
            } else {
              value_num = rt;
            }

            onCellChange({
              numval: value_num,
            });
          }}
        >
          <Icon
            name="star"
            class="h-8 w-8 text-gray-400 {rt <= value_num
              ? 'text-yellow-400'
              : ''} "
            solid={true}
          />
        </button>
      {/each}
    </div>
  {:else if column.ctype === SheetColTypeReference}
    <div class="flex gap-1">
      {#if value_num}
        <span class="bg-blue-100 rounded p-0.5 text-gray-600">
          Ref:
          {value_num}
          <strong class="font-semibold text-gray-700">{value}</strong>
        </span>
      {/if}
    </div>
  {:else if column.ctype === SheetColTypeRemote}
    <div class="flex gap-1">
      <span class="bg-green-100 rounded p-0.5 text-gray-600">
        Remote:
        <strong class="font-semibold text-gray-700">{value}</strong>
      </span>
    </div>
  {:else if column.ctype === SheetColTypeLocation}
    {#key value}
      <Point {value} />
    {/key}
  {:else if column.ctype === SheetColTypeSelect}
    <select
      class="w-full border border-dashed rounded-lg bg-gray-50 p-1"
      {value}
      on:change={(ev) => onCellChange({ value: ev.target["value"] })}
    >
      {#each column.opts ? column.opts.split(",") : [] as opt}
        <option value={opt}>{opt}</option>
      {/each}
    </select>
  {:else if column.ctype === SheetColTypeMultiSelect}
    <Selectable
      onChange={(nval) => onCellChange({ value: nval })}
      options={column.opts ? column.opts.split(",") : []}
      {value}
    />
  {:else if column.ctype === SheetColTypeUser}
    <div class="flex gap-1">
      {#each value.split(",") as cd}
        <UserAvatar name={cd} />
      {/each}
    </div>
  {:else}
    <input
      {id}
      {value}
      on:change={(ev) => onCellChange({ value: ev.target["value"] })}
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline hover:border-blue-400"
      type="text"
    />
  {/if}

  <CellActions
    onCellColor={() => {
      open_column = null;
      color_open = !color_open;
    }}
    is_cell_open={_is_open}
    onCellClose={close}
    onPick={toggle}
    pick_icon={picker_icons[column.ctype]}
  />

  <!-- PICKUP PANEL -->

  {#if _is_open}
    <div
      class="p-1 border rounded shadow h-64 mt-2 border-green-500 overflow-auto"
    >
      {#if column.ctype === SheetColTypeLocation}
        <MapPanel
          onChange={(val) => {
            console.log("@value changed", val);
            value = val;
            onCellChange({ value });
          }}
          {value}
        />
      {:else if column.ctype === SheetColTypeRemote}
        <Remote />
      {:else if column.ctype === SheetColTypeReference}
        <Reference
          {service}
          {column}
          current={value_num}
          onSelect={(val) => {
            value = val.ref_value;
            value_num = val.__id;
            onCellChange({ value, numval: value_num });
          }}
        />
      {:else if column.ctype === SheetColTypeUser}
        <UserPanel {column} {service} />
      {:else if column.ctype === SheetColTypeFile}
        <FilePanel
          {folder_api}
          onFileAdd={(file) => {
            console.log("@file", file);

            let valueArray = value.split(",");
            valueArray = valueArray.filter((v) => v !== "");
            if (valueArray.includes(file)) {
              return;
            }

            valueArray.push(file);

            value = valueArray.join();

            onCellChange({
              value,
            });
          }}
        />
      {/if}
    </div>
  {/if}

  {#if !open_column && color_open}
    <ColorPanel {onColorChange} {color} />
  {/if}
</div>

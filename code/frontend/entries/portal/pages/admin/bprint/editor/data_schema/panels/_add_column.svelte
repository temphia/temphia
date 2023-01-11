<script lang="ts">
  import * as f from "../../../../../data/tableui/core/fields/field";
  import { ActionButton } from "../../../../core";

  const ctypes = [
    f.CtypeShortText,
    f.CtypePhone,
    f.CtypeSelect,
    f.CtypeRFormula,
    f.CtypeFile,
    f.CtypeMultiFile,
    f.CtypeCheckBox,
    f.CtypeCurrency,
    f.CtypeNumber,
    f.CtypeLocation,
    f.CtypeDateTime,
    f.CtypeMultSelect,
    f.CtypeLongText,
    f.CtypeSingleUser,
    f.CtypeMultiUser,
    f.CtypeEmail,
    f.CtypeJSON,
    f.CtypeRangeNumber,
    f.CtypeRatings,
    f.CtypeColor,
  ];

  export let callback: (data: any, ref_data?: any) => void;
  export let current_schema: any;
  export let current_table;

  let name = "";
  let slug = "";
  let description = "";
  let ctype = f.CtypeShortText;
  let icon: string = "";
  let options: string[] = [];
  let not_nullable: boolean = true;
  let pattern: string = "";
  let strict_pattern: boolean = false;

  let ref_enable = false;
  let ref_table = "";
  let ref_column = "";
  let ref_hard = false;

  $: _possible_ref_tables = current_schema.tables.filter(
    (val) => val.slug !== current_table
  );

  $: _possible_ref_columns = (
    current_schema.tables.filter((val) => val.slug === ref_table)[0] || {
      columns: [],
    }
  ).columns;

  $: console.log(
    "$REF_Tables",
    _possible_ref_tables,
    "$REF_COL",
    _possible_ref_columns
  );

  $: __refable =
    f.CtypeShapes.text.includes(ctype) || f.CtypeShapes.number.includes(ctype);

  const onAdd = () => {
    const data = {
      name,
      slug,
      description,
      ctype,
      icon,
      options,
      not_nullable,
      pattern,
      strict_pattern,
    };

    let ref_type = "";
    if (ref_column === f.KeyPrimary) {
      if (ref_hard) {
        ref_type = f.RefHardPriId;
      } else {
        ref_type = f.RefSoftPriId;
      }
    } else if (f.CtypeShapes.text.includes(ctype)) {
      if (ref_hard) {
        ref_type = f.RefHardText;
      } else {
        ref_type = f.RefSoftText;
      }
    } else {
      console.log("fixme => Handle ref");
    }

    const ref_data = {
      slug: "",
      type: ref_type,
      target: ref_table,
      from_cols: [],
      to_cols: [],
    };

    callback(data, ref_enable ? ref_data : null);
  };
</script>

<div class="p-2">
  <h2 class="font-medium leading-tight text-xl">Add Column</h2>

  <div class="flex flex-col">
    <div class="flex-col flex py-3">
      <label class="pb-2 text-gray-700 font-semibold">Name </label>
      <input
        type="text"
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        bind:value={name}
      />
    </div>

    <div class="flex-col flex py-3">
      <label class="pb-2 text-gray-700 font-semibold">Slug </label>

      <input
        type="text"
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        bind:value={slug}
      />
    </div>

    <div class="flex-col flex py-3">
      <label class="pb-2 text-gray-700 font-semibold">Ctype </label>
      <select bind:value={ctype} class="p-2 rounded">
        {#each ctypes as ct}
          <option value={ct}>{ct}</option>
        {/each}
      </select>
    </div>

    <div class="flex-col flex py-3">
      <label class="pb-2 text-gray-700 font-semibold">Info </label>
      <textarea
        type="text"
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        bind:value={description}
      />
    </div>

    <div class="flex-col flex py-3">
      <label class="pb-2 text-gray-700 font-semibold"
        >Not Nullable
        <input type="checkbox" bind:checked={not_nullable} />
      </label>
    </div>

    {#if __refable}
      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold"
          >Reference
          <input type="checkbox" bind:checked={ref_enable} />
        </label>
      </div>
    {/if}

    {#if ref_enable && __refable}
      <div class="flex-col flex p-2 border">
        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Table</label>
          <select
            bind:value={ref_table}
            class="p-2 rounded"
            on:change={() => {
              ref_column = "";
            }}
          >
            {#each _possible_ref_tables || [] as tbl}
              <option value={tbl.slug}>{tbl.name}</option>
            {/each}
          </select>
        </div>

        {#key ref_table}
          <div class="flex-col flex py-3">
            <label class="pb-2 text-gray-700 font-semibold">Column</label>
            <select bind:value={ref_column} class="p-2 rounded">
              <option value={f.KeyPrimary}>{f.KeyPrimary}</option>

              {#each _possible_ref_columns || [] as col}
                <option value={col.slug}>{col.name}</option>
              {/each}
            </select>
          </div>
        {/key}
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold"
          >Hard Reference
          <input type="checkbox" bind:checked={ref_hard} />
        </label>
      </div>
    {/if}

    <div class="flex justify-end">
      {#if name && slug && ctype}
        <ActionButton icon_name="plus" name="Add" onClick={onAdd} />
      {/if}
    </div>
  </div>
</div>

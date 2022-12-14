<script lang="ts">
  import * as f from "../../../data/tableui/core/fields/field";
  import { ActionButton } from "../../core";

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

  export let callback: (data: any) => void;
  export let data = {};

  let name = data["name"] || "";
  let slug = data["slug"] || "";
  let info = data["description"] || "";
  let ctype = data["ctype"] || f.CtypeShortText;

  const onEdit = () => {
    const data = {
      name,
      slug,
      info,
    };

    callback(data);
  };
</script>

<div class="p-2">
  <h2 class="font-medium leading-tight text-xl">Edit Column</h2>

  <div class="flex flex-col">
    <div class="flex-col flex py-3">
      <label class="pb-2 text-gray-700 font-semibold">Slug </label>
      <input
        type="text"
        disabled
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        bind:value={slug}
      />
    </div>

    <div class="flex-col flex py-3">
      <label class="pb-2 text-gray-700 font-semibold">Name </label>
      <input
        type="text"
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        bind:value={name}
      />
    </div>

    <div class="flex-col flex py-3">
      <label class="pb-2 text-gray-700 font-semibold">Ctype </label>
      <select value={ctype} class="p-2 rounded" disabled>
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
        bind:value={info}
      />
    </div>

    <div class="flex justify-end">
      {#if name && slug && ctype}
        <ActionButton icon_name="pencil-alt" name="Edit" onClick={onEdit} />
      {/if}
    </div>
  </div>
</div>

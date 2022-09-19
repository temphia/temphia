<script lang="ts">
  import { hslColor } from "../../../../../lib/utils";
  import FilesPreview from "./fields/file/_files_preview.svelte";
  import {
    ctypeFileDecode,
    //
    CtypeFile,
    CtypeMultiFile,
    CtypeSelect,
    CtypeNumber,
  } from "./fields/field";

  import { getContext } from "svelte";
  import type { FolderTktAPI } from "../../../../../lib/core/tktapi";
  import type { Column } from "../../../../../lib/service/dyn";

  const { open } = getContext("simple-modal");

  export let row: object;
  export let column: object;
  export let folder_api: FolderTktAPI;

  const _column = column as Column;

  const resolve_ref = (id) => {
    // fixme => check in store for reverse ref
    return id;
  };
</script>

<div class="text-gray-700 truncate overflow-hidden text-sm p-1">
  {#if row[_column.slug] === null}
    <div />
  {:else if _column.slug in row}
    {#if _column.ctype === CtypeNumber}
      {#if _column.ref_type}
        {resolve_ref(row[_column.slug])}
      {:else}
        {row[_column.slug]}
      {/if}
    {:else if _column.ctype === CtypeFile || _column.ctype === CtypeMultiFile}
      <div class="flex flex-row">
        {#each ctypeFileDecode(row[_column.slug]) as file}
          <img
            on:click={() =>
              open(FilesPreview, {
                folder_api,
                files: ctypeFileDecode(row[_column.slug]),
                current_file: file,
              })}
            src={folder_api.get_file_preview_link(file)}
            class="h-11 p-1 border cursor-pointer"
            alt=""
          />
        {/each}
      </div>
    {:else if _column.ctype === CtypeSelect}
      <span class="rounded p-1" style={hslColor(row[_column.slug])}
        >{row[_column.slug]}</span
      >
    {:else}
      {row[_column.slug]}
    {/if}
  {/if}
</div>

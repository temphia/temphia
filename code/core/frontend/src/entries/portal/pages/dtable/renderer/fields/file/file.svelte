<script lang="ts">
  import { getContext } from "svelte";
  import FilesPreview from "./_files_preview.svelte";
  import { ctypeFileDecode } from "../field";
  import FilepickDialog from "./_file_panel.svelte";
  import type Column from "../../../../admin/dtable/_column.svelte";
  import type {
    DataTableService,
    RowEditor,
  } from "../../../../../../../lib/service/dyn";

  export let multi: boolean;
  export let value: any;
  export let column: Column;
  export let manager: DataTableService;
  export let row_editor: RowEditor;

  export let onChange: (_value: any) => void;

  const { open, close } = getContext("simple-modal");

  $: _files = ctypeFileDecode(value);

  const showPreview = (file: string) => {
    open(FilesPreview, {
      folder_api: manager.FolderTktAPI,
      files: _files,
      current_file: file,
    });
  };

  const showDialog = () => {
    open(FilepickDialog, {
      folder_api: manager.FolderTktAPI,
      onSelect: (file: string) => {
        if (multi) {
          _files = Array.from(new Set([..._files, file]));
        } else {
          _files = [file];
        }
        const final = _files.toString();
        if (final !== value) {
          onChange(final);
        }
        close();
      },
    });
  };

  const removeFile = (file: string) => {
    if (multi) {
      _files = Array.from(new Set(_files.filter((val) => val !== file)));
    } else {
      _files = [file];
    }
    const final = _files.toString();
    if (final !== value) {
      onChange(final);
      return;
    }
  };
</script>

<div class="flex flex-col w-full h-full p-1 overflow-auto">
  <div
    class="flex flex-wrap p-1 space-x-1 space-y-1 border border-dashed rounded-lg bg-gray-50"
    style="min-height: 2rem;"
  >
    {#each _files as file}
      <div class="relative">
        <svg
          on:click={() => removeFile(file)}
          xmlns="http://www.w3.org/2000/svg"
          class="h-6 w-6 absolute right-0 text-blue-500 border rounded-full bg-white cursor-pointer"
          viewBox="0 0 20 20"
          fill="currentColor"
        >
          <path
            fill-rule="evenodd"
            d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
            clip-rule="evenodd"
          />
        </svg>

        <img
          on:click={() => showPreview(file)}
          src={manager.FolderTktAPI.get_file_preview_link(file)}
          class="h-20 p-1 border bg-white cursor-pointer"
          alt=""
        />
      </div>
    {/each}
  </div>

  <div
    class="p-2 flex justify-end text-gray-600 cursor-pointer"
    on:click={showDialog}
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      viewBox="0 0 20 20"
      fill="currentColor"
    >
      <path
        d="M5.5 13a3.5 3.5 0 01-.369-6.98 4 4 0 117.753-1.977A4.5 4.5 0 1113.5 13H11V9.413l1.293 1.293a1 1 0 001.414-1.414l-3-3a1 1 0 00-1.414 0l-3 3a1 1 0 001.414 1.414L9 9.414V13H5.5z"
      />
      <path d="M9 13h2v5a1 1 0 11-2 0v-5z" />
    </svg>
    <span>Pick</span>
  </div>
</div>

<script lang="ts">
  import { getContext } from "svelte";
  import { CEditor } from "../../../../../../components";
  import type { Editor } from "./editor";

  import PanelAddFile from "./panels/add_file.svelte";
  import PanelSave from "./panels/save.svelte";
  import PanelSaveFile from "./panels/save_file.svelte";
  import PanelCheckFile from "./panels/check_file.svelte";

  export let beditor: Editor;

  const files = beditor.files;

  const MODE_FILE = "file";
  const MODE_RAW = "raw";
  const MODE_COMMON = "common";
  const MODE_SCHEMA = "schema";

  $: _collapsed = false;
  $: _mode = MODE_COMMON;
  $: _modified = beditor.modified;
  $: _mod_files = beditor.fileMods;
  $: _current_file = "";
  let code_mode = "js";
  let codemirror;

  const { open, close } = getContext("simple-modal");

  const update_field = (key) => (ev) => {
    beditor.updateField(key, ev.target.value);
  };

  const get_field = (key) => beditor.getField(key);

  const set_modified = () => {
    beditor.setModified();
  };

  const set_file_modified = (file) => () => {
    beditor.setModifiedFile(file);
  };

  const change_view = (new_view) => () => {
    switch (_mode) {
      case MODE_FILE:
        if (!codemirror) {
          break;
        }
        if ($_mod_files[_current_file]) {
          const value = codemirror.getValue();
          console.log("SAVING DIRTY FILE", value);
          beditor.updateFile(_current_file, value);
        }

        break;
      case MODE_RAW:
        if (!codemirror) {
          break;
        }
        const data = codemirror.getValue();
        beditor.update(JSON.parse(data));

        break;
      default:
        break;
    }
    _mode = new_view;
  };

  const change_view_to_file = (file) => () => {
    if (file.endsWith(".js")) {
      code_mode = "js";
    } else if (file.endsWith(".css")) {
      code_mode = "css";
    } else if (file.endsWith(".json")) {
      code_mode = "json";
    } else if (file.endsWith(".html")) {
      code_mode = "html";
    }
    change_view(MODE_FILE)();
    _current_file = file;
  };

  const add = () => {
    open(PanelAddFile, { beditor });
  };
  const save = () => {
    console.log(_mode);

    switch (_mode) {
      case MODE_FILE:
        open(PanelSaveFile, {
          filename: _current_file,
          beditor,
          finalSave: async () => {
            const value = codemirror.getValue();
            beditor.updateFile(_current_file, value);
            const resp = await beditor.saveFile(_current_file);
            close();
          },
        });
        break;
      default:
        open(PanelSave, {});
    }
  };

  const check = () => {
    open(PanelCheckFile, {
      filename: _current_file,
      beditor,
      finalSave: async () => {},
    });
  };

  const toggle = () => {
    _collapsed = !_collapsed;
  };
</script>

<div class="w-full h-full flex">
  <div class="absolute bottom-1 z-10 p-1">
    <button
      on:click={toggle}
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-2 rounded"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M4 6h16M4 12h16M4 18h7"
        />
      </svg>
    </button>
  </div>
  {#if !_collapsed}
    <div class="flex-none w-64 flex flex-col bg-white border-r">
      <div class="flex items-center justify-center h-14 border-b">
        <div>Blueprint Editor</div>
      </div>
      <div class="overflow-y-auto flex flex-col overflow-x-hidden">
        <ul class="flex flex-col py-4 space-y-1">
          <li class="px-5">
            <div class="flex flex-row items-center h-8">
              <div class="text-sm font-light tracking-wide text-gray-500">
                System
              </div>

              {#if $_modified}
                <span
                  class="px-2 py-0.5 ml-auto text-xs font-medium tracking-wide text-red-500 bg-red-50 rounded-full"
                  >Modified</span
                >
              {/if}
            </div>
          </li>
          <li>
            <a
              on:click={change_view(MODE_COMMON)}
              href="#"
              class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-indigo-500 pr-6 {_mode ===
              MODE_COMMON
                ? 'border-indigo-600'
                : ''}"
            >
              <span class="inline-flex justify-center items-center ml-4">
                <svg
                  class="w-5 h-5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                  ><path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
                  /></svg
                >
              </span>
              <span class="ml-2 text-sm tracking-wide truncate">Common</span>
            </a>
          </li>
          <li>
            <a
              href="#"
              on:click={change_view(MODE_RAW)}
              class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-indigo-500 pr-6 {_mode ===
              MODE_RAW
                ? 'border-indigo-600'
                : ''}"
            >
              <span class="inline-flex justify-center items-center ml-4">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-6 w-6"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"
                  />
                </svg>
              </span>
              <span class="ml-2 text-sm tracking-wide truncate">Raw</span>
            </a>
          </li>
          <li>
            <a
              href="#"
              on:click={change_view(MODE_SCHEMA)}
              class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-indigo-500 pr-6 {_mode ===
              MODE_SCHEMA
                ? 'border-indigo-600'
                : ''}"
            >
              <span class="inline-flex justify-center items-center ml-4">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-6 w-6"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14"
                  />
                </svg>
              </span>
              <span class="ml-2 text-sm tracking-wide truncate"
                >Inline Schema</span
              >
            </a>
          </li>
          <li class="px-5">
            <div class="flex flex-row items-center h-8">
              <div class="text-sm font-light tracking-wide text-gray-500">
                Files
              </div>
            </div>
          </li>

          {#each $files as file}
            <li>
              <a
                href="#"
                on:click={change_view_to_file(file)}
                class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-indigo-500 pr-6 {_mode ===
                  MODE_FILE && _current_file === file
                  ? 'border-indigo-600'
                  : ''}"
              >
                <span class="inline-flex justify-center items-center ml-4">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-6 w-6"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                    />
                  </svg>
                </span>
                <span class="ml-2 text-sm tracking-wide truncate">{file}</span>
                {#if $_mod_files[file]}
                  <span
                    class="px-2 py-0.5 ml-auto text-xs font-medium tracking-wide text-red-500 bg-red-50 rounded-full"
                    >Modified</span
                  >
                {/if}
              </a>
            </li>
          {/each}
        </ul>

        <div class="border flex flex-col p-2">
          {#if _mode !== MODE_COMMON}
            <div class="w-full">
              <span class="text-base p-1 font-medium text-gray-900 p-2"
                >Lang</span
              >
              <select class="p-1 rounded" bind:value={code_mode}>
                <option value="js">JS</option>
                <option value="json">JSON</option>
                <option value="yaml">YAML</option>
                <option value="html">HTML</option>
                <option value="css">CSS</option>
              </select>
            </div>
          {/if}

          <div class="p-2 flex justify-center space-x-1">
            <button
              on:click={add}
              class="bg-indigo-500 hover:bg-indigo-700 text-white text-xs py-1 px-2 rounded inline-flex"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z"
                  clip-rule="evenodd"
                />
              </svg>
              Add
            </button>

            {#if _mode === MODE_FILE}
              <button
                on:click={check}
                class="bg-green-500 hover:bg-green-700 text-white text-xs py-1 px-2 rounded inline-flex"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd"
                  />
                </svg>
                Check
              </button>

              <button
                on:click={save}
                class="bg-blue-500 hover:bg-blue-700 text-white text-xs py-1 px-2 rounded inline-flex"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    d="M7.707 10.293a1 1 0 10-1.414 1.414l3 3a1 1 0 001.414 0l3-3a1 1 0 00-1.414-1.414L11 11.586V6h5a2 2 0 012 2v7a2 2 0 01-2 2H4a2 2 0 01-2-2V8a2 2 0 012-2h5v5.586l-1.293-1.293zM9 4a1 1 0 012 0v2H9V4z"
                  />
                </svg>
                Save
              </button>
            {/if}
          </div>
        </div>
      </div>
    </div>
  {/if}

  {#if _mode === MODE_COMMON}
    <div class="h-full w-full bg-indigo-100 p-10 overflow-auto">
      <div class="p-5 bg-white w-full ">
        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Id</label>
          <input
            type="text"
            disabled
            value={get_field("id")}
            on:change={update_field("id")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder="twgvhsysabh"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Name</label>
          <input
            type="text"
            value={get_field("name")}
            on:change={update_field("name")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder="name.."
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Group</label>
          <select
            value={get_field("group")}
            on:change={update_field("group")}
            class="form-select block w-full  p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          >
            <option value="plug">Plug</option>
            <option value="tschema">Table Schema</option>
          </select>
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Sub Group</label>
          <input
            type="text"
            value={get_field("sub_group")}
            on:change={update_field("sub_group")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder=""
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Schema</label>
          <textarea
            value={get_field("schema")}
            on:change={update_field("schema")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Description</label>
          <textarea
            on:change={update_field("description")}
            value={get_field("description")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Source</label>
          <input
            type="text"
            on:change={update_field("source")}
            value={get_field("source")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex py-3">
          <button class="p-2 bg-blue-400 m-1 w-20 text-white rounded"
            >Save</button
          >
        </div>
      </div>
    </div>
  {:else if _mode === MODE_RAW}
    <div style="width: inherit;">
      <CEditor
        bind:editor={codemirror}
        code={typeof beditor.data === "string"
          ? beditor.data
          : JSON.stringify(beditor.data, null, 4)}
        container_style="height:100vh;"
        on:change={set_modified}
      />
    </div>
  {:else if _mode === MODE_SCHEMA}
    <div style="width: inherit;">
      <CEditor
        bind:editor={codemirror}
        container_style="height:100vh;"
        on:change={set_modified}
      />
    </div>
  {:else if _mode === MODE_FILE}
    {#await beditor.getFile(_current_file)}
      <p>...waiting</p>
    {:then code}
      <div style="width: inherit;">
        {#key code_mode}
          <CEditor
            mode={code_mode}
            code={typeof code === "string"
              ? code
              : JSON.stringify(code, null, 4)}
            bind:editor={codemirror}
            on:change={set_file_modified(_current_file)}
            container_style="height:100vh;"
          />
        {/key}
      </div>
    {:catch error}
      <p style="color: red">File load err :{error.message}</p>
    {/await}
  {/if}
</div>

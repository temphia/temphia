<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../app";

  import ListAuth from "./auth/list_auth.svelte";
  import ListHook from "./hook/list_hook.svelte";
  import ListPlug from "./plug/list_plug.svelte";
  import ListData from "./data/list_data.svelte";
  import AuthNewTemplate from "./auth/auth_new_template.svelte";
  import ArrayStrEdit from "../../../../_shared/common/array_str_edit.svelte";

  export let name = "";
  export let slug = "";
  export let icon = "";
  export let enable_pass_auth = true;
  export let open_sign_up = false;
  export let scopes = "";

  export let onSave;
  export let new_mode = false;

  export let app: PortalApp;

  const { open, close } = getContext("simple-modal");

  const save = async () => {
    await onSave(slug, {
      name,
      slug,
      icon,
      enable_pass_auth,
      scopes,
      open_sign_up,
    });

    app.navigator.goto_admin_usergroups_page();
  };

  $: _show_dropdown = false;
</script>

<div class="h-full w-full bg-indigo-100 overflow-auto">
  <div class="md:p-12 flex flex-row flex-wrap">
    <div
      class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg relative"
    >
      <div class="absolute right-2">
        <button
          class="relative z-10 block rounded-md border bg-white p-2 focus:outline-none"
          on:click={() => (_show_dropdown = !_show_dropdown)}
        >
          <svg
            class="h-5 w-5 text-gray-800"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fill-rule="evenodd"
              d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
              clip-rule="evenodd"
            />
          </svg>
        </button>

        {#if _show_dropdown}
          <div
            x-show="dropdownOpen"
            class="absolute border right-0 mt-2 py-2 w-48 bg-white rounded-md shadow-xl z-20"
          >
            <a
              href="#"
              class="block px-4 py-2 text-sm capitalize text-gray-700 hover:bg-blue-500 hover:text-white"
            >
              Generate token
            </a>
          </div>
        {/if}
      </div>

      <div class="text-2xl text-indigo-900">User Group</div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Name</label>
        <input
          type="text"
          bind:value={name}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Slug</label>
        <input
          type="text"
          disabled={!new_mode}
          bind:value={slug}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Scopes</label>
        <ArrayStrEdit bind:value={scopes} />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold"
          >Enable Password Auth</label
        >
        <div>
          <input
            type="checkbox"
            bind:checked={enable_pass_auth}
            class="p-2 h-5 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Open Sign ups</label>

        <div>
          <input
            type="checkbox"
            bind:checked={open_sign_up}
            class="p-2 h-5 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>
      </div>

      <div class="flex py-3 justify-end">
        <button
          on:click={save}
          class="p-2 bg-blue-400 hover:bg-blue-600 m-1 w-20 text-white rounded"
          >Save</button
        >
      </div>
    </div>
  </div>
</div>

{#if !new_mode}
  <div class="md:p-12 flex flex-row flex-wrap">
    <div
      class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg relative"
    >
      <div class="text-2xl text-indigo-900">Extra</div>

      <div class="flex-col flex py-3 relative">
        <button
          on:click={() => {
            app.simple_modal_open(AuthNewTemplate, { app, gid: slug });
          }}
          class="absolute right-0 p-1 bg-blue-400 text-white rounded hover:scale-110"
        >
          add
        </button>
        <label class="pb-2 text-gray-700 font-semibold">Auth Providers</label>
        <ListAuth {app} gid={slug} />
      </div>

      <div class="flex-col flex py-3 relative">
        <button
          class="absolute right-0 p-1 bg-blue-400 text-white rounded hover:scale-110"
          on:click={() => app.navigator.goto_admin_user_hook_new(slug)}
        >
          add
        </button>
        <label class="pb-2 text-gray-700 font-semibold">Hooks</label>
        <ListHook {app} user_group={slug} />
      </div>

      <div class="flex-col flex py-3 relative">
        <button
          class="absolute right-0 p-1 bg-blue-400 text-white rounded hover:scale-110"
          on:click={() => app.navigator.goto_admin_user_plug_new(slug)}
        >
          add
        </button>
        <label class="pb-2 text-gray-700 font-semibold">Plugs</label>
        <ListPlug {app} user_group={slug} />
      </div>

      <div class="flex-col flex py-3 relative">
        <button
          class="absolute right-0 p-1 bg-blue-400 text-white rounded hover:scale-110"
          on:click={() => app.navigator.goto_admin_user_data_new(slug)}
        >
          add
        </button>
        <label class="pb-2 text-gray-700 font-semibold">Datas</label>
        <ListData {app} user_group={slug} />
      </div>
    </div>
  </div>
{/if}

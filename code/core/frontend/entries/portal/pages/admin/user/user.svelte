<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../app";
  import Layout from "../layout.svelte";

  export let id = "";

  const app: PortalApp = getContext("__app__");

  let user = {};

  $: _mod_data = {};
  $: _show_dropdown = false;
  $: _loading = true;

  app.get_apm().get_user_api().then(async (uapi) => {
    const resp = await uapi.get_user_by_id(id);
    _loading = false;
    user = resp.data;
  });

  const onFieldChange = (field) => (ev) => {
    _mod_data = { ..._mod_data, [field]: ev.target.value };
  };

  const onSave = async () => {
    const uapi = await app.get_apm().get_user_api();
    uapi.update_user(id, _mod_data);
    app.navigator.goto_admin_usergroups_page();
  };
</script>

<Layout loading={_loading} current_item={"user_groups"}>
  <div class="h-full w-full overflow-auto">
    <div class="md:p-12 bg-indigo-100 flex flex-row flex-wrap">
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
                Change Email
              </a>
              <a
                href="#"
                class="block px-4 py-2 text-sm capitalize text-gray-700 hover:bg-blue-500 hover:text-white"
              >
                Change Pub Key
              </a>

              <a
                href="#"
                class="block px-4 py-2 text-sm capitalize text-gray-700 hover:bg-blue-500 hover:text-white"
              >
                Set Data
              </a>
            </div>
          {/if}
        </div>
        <div class="text-2xl text-indigo-900">User</div>

        <div class="text-center p-6  border-b">
          <img
            class="h-24 w-24 rounded-full mx-auto"
            src={app.user_profile_image_link(id)}
            alt="user profile"
          />
          <p class="pt-2 text-lg font-semibold">{user["full_name"] || ""}</p>
          <p class="text-sm text-gray-600">{user["email"] || ""}</p>
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Full Name</label>
          <input
            type="text"
            value={user["full_name"] || ""}
            on:change={onFieldChange("full_name")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder="John"
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">User Id</label>
          <input
            type="text"
            disabled={true}
            value={user["user_id"]}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder="user1"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Email</label>
          <input
            type="text"
            value={user["email"] || ""}
            disabled
            placeholder="user@example.com"
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Public Key</label>
          <textarea
            type="text"
            disabled
            value={user["public_key"] || ""}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Group</label>
          <input
            type="text"
            value={user["group_id"]}
            on:change={onFieldChange("group_id")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder="group1"
          />
        </div>

        <div class="flex justify-between space-x-1">
          <button on:click={onSave} class="p-2 bg-blue-400 text-white rounded"
            >Save</button
          >
        </div>
      </div>
    </div>
  </div>
</Layout>

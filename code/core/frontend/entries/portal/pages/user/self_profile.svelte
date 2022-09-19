<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../app";

  const app: PortalApp = getContext("__app__");
  const bapi = app.get_apm().get_basic_api();

  $: _show_dropdown = false;
  $: _loading = true;
  let data = {};

  bapi.get_self_info().then((resp) => {
    data = resp.data;
    _loading = false;
  });

  let moddata = {};
  let data_modified = false;
  const onModData = (key) => (ev) => {
    moddata[key] = ev.target.value;
    data_modified = true;
  };
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
          </div>
        {/if}
      </div>

      <div class="text-2xl text-indigo-900">User Profile</div>

      <div class="text-center p-6  border-b">
        <img
          class="h-24 w-24 rounded-full border p-1 mx-auto"
          src={app.user_profile_image_link("id")}
          alt="user profile"
        />
      </div>
      {#if _loading}
        <div>Loading..</div>
      {:else}
        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Name</label>
          <input
            type="text"
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            value={data["full_name"] || ""}
            on:change={onModData("full_name")}
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Bio</label>
          <textarea
            type="text"
            value={data["bio"] || ""}
            on:change={onModData("bio")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder="I like stuff.."
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">User Id</label>
          <input
            type="text"
            disabled
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            value={data["user_id"] || ""}
            placeholder="John"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Your Email</label>
          <input
            type="text"
            disabled
            value={data["email"] || ""}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Group</label>
          <input
            type="text"
            value={data["group"] || ""}
            disabled
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder="group1"
          />
        </div>

        {#if data["pub_key"]}
          <div class="flex-col flex py-3 relative">
            <label class="pb-2 text-gray-700 font-semibold">Pub Key</label>
            <textarea
              type="text"
              value={data["pub_key"] || ""}
              on:change={onModData("pub_key")}
              class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            />
          </div>
        {/if}

        <div class="flex justify-between space-x-1">
          <div />
          {#if data_modified}
            <button
              class="px-2 py-1 border border-blue700 bg-blue-400 hover:bg-blue-600 text-white rounded"
              >Save</button
            >
          {/if}
        </div>
      {/if}
    </div>
  </div>
</div>

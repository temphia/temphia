<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";
  import { LoadingSpinner } from "../admin/core";
  import TopActions from "../admin/core/top_actions.svelte";

  const app: PortalService = getContext("__app__");
  const api = app.api_manager.get_self_api();

  let loading = true;
  let data = {};
  let moddata = {};
  let data_modified = false;

  const load = async () => {
    const resp = await api.self();
    if (!resp.ok) {
      return;
    }

    data = resp.data;
    loading = false;
  };

  load();

  const onModData = (key) => (ev) => {
    moddata[key] = ev.target.value;
    data_modified = true;
  };
</script>

<div class="h-full w-full bg-indigo-100 overflow-auto">
  <TopActions
    actions={{
      "Devices and Logins": () => app.nav.self_devices(),
    }}
  />

  {#if loading}
    <LoadingSpinner />
  {:else}
    <div class="md:p-12 flex flex-row flex-wrap">
      <div
        class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg relative"
      >
        <div class="text-2xl text-indigo-900">User Profile</div>

        <div class="text-center p-6  border-b">
          <img
            class="w-24 rounded-full border p-1 mx-auto"
            src={app.get_user_profile(data["user_id"] || "")}
            alt="user profile"
          />
        </div>

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

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Pub Key</label>
          <textarea
            value={data["pub_key"] || ""}
            on:change={onModData("pub_key")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex justify-between space-x-1">
          <div />
          {#if data_modified}
            <button
              class="px-2 py-1 border border-blue700 bg-blue-400 hover:bg-blue-600 text-white rounded"
              >Save</button
            >
          {/if}
        </div>
      </div>
    </div>
  {/if}
</div>

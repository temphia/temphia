<script lang="ts">
  import UserGroup from "../../../xcompo/svg/user_group.svelte";
  import User from "../../../xcompo/svg/user.svelte";
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";
  import { params } from "svelte-hash-router";
  import { LoadingSpinner } from "../admin/core";

  export let id = $params.id;

  const app: PortalService = getContext("__app__");

  const api = app.api_manager.get_self_api();

  let userData = {};
  let loading = true;
  let message = "";

  const load = async () => {
    const resp = await api.user_profile(id);
    if (!resp.ok) {
      return;
    }
    userData = resp.data;
    loading = false;
  };

  load();

  const sendMessage = async () => api.user_message(id, message);
</script>

<div class="h-full w-full bg-indigo-100 overflow-auto">
  <div class="md:p-12 flex flex-row flex-wrap">
    <div
      class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg relative"
    >
      {#if loading}
        <LoadingSpinner />
      {:else}
        <div class="text-center p-6  border-b">
          <img
            class="w-24 rounded-full mx-auto"
            src={app.get_user_profile(id)}
            alt="user proile {id}"
          />
          <p class="pt-2 text-lg font-semibold">
            {userData["full_name"] || ""}
          </p>
          <p class="text-sm text-gray-600">{userData["email"] || ""}</p>
        </div>
        <div class="border-b">
          <a href="#" class="px-4 py-2 hover:bg-gray-100 flex">
            <div class="text-gray-800">
              <User />
            </div>
            <div class="pl-3">
              <p class="text-sm font-medium text-gray-800 leading-none">
                User Id
              </p>
              <p class="text-xs text-gray-500">{id}</p>
            </div>
          </a>

          <a href="#" class="px-4 py-2 hover:bg-gray-100 flex">
            <div class="text-gray-800">
              <UserGroup />
            </div>
            <div class="pl-3">
              <p class="text-sm font-medium text-gray-800 leading-none">
                Group
              </p>
              <p class="text-xs text-gray-500">{userData["group"] || ""}</p>
            </div>
          </a>

          <a href="#" class="px-4 py-2 hover:bg-gray-100 flex">
            <div class="text-gray-800">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                stroke-width="2"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z"
                />
              </svg>
            </div>
            <div class="pl-3">
              <p class="text-sm font-medium text-gray-800 leading-none">Bio</p>
              <p class="text-xs text-gray-500">{userData["bio"] || ""}</p>
            </div>
          </a>
        </div>

        <div class="flex-col flex py-3 relative">
          <label for="message" class="pb-2 text-gray-700 font-semibold"
            >Message</label
          >
          <textarea
            id="message"
            bind:value={message}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder="Hey, ..."
          />

          <div class="p-2 flex justify-end">
            {#if message}
              <button
                on:click={sendMessage}
                class="px-2 py-1 border w-30 border-blue700 bg-blue-400 hover:bg-blue-600 text-white rounded"
                >Send</button
              >
            {/if}
          </div>
        </div>
      {/if}
    </div>
  </div>
</div>

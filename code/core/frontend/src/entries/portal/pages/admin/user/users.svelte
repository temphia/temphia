<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../lib/app/portal";

  export let user_group = "";

  const app: PortalApp = getContext("__app__");

  let users = [];

  const load = async () => {
    const uapi = await app.get_apm().get_user_api();
    const resp = await uapi.list_users(user_group);
    users = resp.data;
  };

  const profile = app.user_profile_image_link;

  const delUser = (id) => async () => {
    const uapi = await app.get_apm().get_user_api();
    const resp = await uapi.remove_user(id);
    if (resp.status !== 200) {
      console.log("Err", resp);
      return;
    }
    load()
  };
  const editUser = (id) => () => app.navigator.goto_admin_user_page(id);

  load();
</script>

<div class="p-8">
  <div>
    <h2 class="text-2xl font-semibold leading-tight">Users</h2>
  </div>
  <div class="my-2 flex sm:flex-row flex-col">
    <div class="flex flex-row mb-1 sm:mb-0">
      <div class="relative">
        <select
          class="appearance-none h-full rounded-l border block appearance-none w-full bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
        >
          <option>5</option>
          <option>10</option>
          <option>20</option>
        </select>
        <div
          class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
        >
          <svg
            class="fill-current h-4 w-4"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
          >
            <path
              d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
            />
          </svg>
        </div>
      </div>
      <div class="relative">
        <select
          class="appearance-none h-full rounded-r border-t sm:rounded-r-none sm:border-r-0 border-r border-b block appearance-none w-full bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:border-l focus:border-r focus:bg-white focus:border-gray-500"
        >
          <option>All</option>
          <option>Active</option>
          <option>Inactive</option>
        </select>
        <div
          class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
        >
          <svg
            class="fill-current h-4 w-4"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
          >
            <path
              d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
            />
          </svg>
        </div>
      </div>
    </div>
    <div class="block relative">
      <span class="h-full absolute inset-y-0 left-0 flex items-center pl-2">
        <svg viewBox="0 0 24 24" class="h-4 w-4 fill-current text-gray-500">
          <path
            d="M10 4a6 6 0 100 12 6 6 0 000-12zm-8 6a8 8 0 1114.32 4.906l5.387 5.387a1 1 0 01-1.414 1.414l-5.387-5.387A8 8 0 012 10z"
          />
        </svg>
      </span>
      <input
        placeholder="Search"
        class="appearance-none rounded-r rounded-l sm:rounded-l-none border border-gray-400 border-b block pl-8 pr-6 py-2 w-full bg-white text-sm placeholder-gray-400 text-gray-700 focus:bg-white focus:placeholder-gray-600 focus:text-gray-700 focus:outline-none"
      />
    </div>
  </div>

  <div class="inline-block min-w-full shadow rounded-lg overflow-hidden">
    <table class="min-w-full leading-normal">
      <thead>
        <tr>
          <th
            class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
          >
            User
          </th>
          <th
            class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
          >
            Group
          </th>
          <th
            class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
          >
            Email
          </th>
          <th
            class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
          >
            Public Key
          </th>
          <th
            class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
          >
            Created at
          </th>
          <th
            class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
          >
            Actions
          </th>
        </tr>
      </thead>
      <tbody>
        {#each users as user}
          <tr>
            <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
              <div class="flex items-center">
                <div class="flex-shrink-0 w-10 h-10">
                  <img
                    class="w-24 h-AUTO rounded-full"
                    src={profile(user.user_id)}
                    alt=""
                  />
                </div>
                <div class="ml-3">
                  <p class="text-gray-900 whitespace-no-wrap">
                    {user.full_name}
                  </p>
                </div>
              </div>
            </td>
            <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
              <p class="text-gray-900 whitespace-no-wrap">{user.group_id}</p>
            </td>
            <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
              <p class="text-gray-900 whitespace-no-wrap">{user.email}</p>
            </td>
            <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
              <p class="text-gray-900 whitespace-no-wrap">
                {user.pub_key}
              </p>
            </td>
            <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
              <p class="text-gray-900 whitespace-no-wrap">
                {user.created_at}
              </p>
            </td>
            <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
              <button
                on:click={() => app.navigator.goto_user_profile(user.user_id)}
                class="p-1 m-1 text-sm font-semibold text-white rounded transform hover:scale-110 bg-green-400"
                >Profile</button
              >

              <button
                on:click={editUser(user.user_id)}
                class="p-1 m-1 text-sm font-semibold text-white rounded transform hover:scale-110 bg-blue-400"
                >Edit</button
              ><button
                on:click={delUser(user.user_id)}
                class="p-1 m-1 text-sm font-semibold text-white rounded transform hover:scale-110 bg-red-400"
                >Delete</button
              >
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>

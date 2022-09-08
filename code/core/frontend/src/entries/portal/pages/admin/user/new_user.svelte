<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../lib/app/portal";
  import Layout from "../layout.svelte";

  const app: PortalApp = getContext("__app__");

  const options = app.navigator.nav_options || {};

  let group_id = options["group_id"] || "";
  let user_id = "";
  let full_name = "";
  let email = "";
  let password = "";
  let pub_key = "";

  const onSave = async () => {
    const uapi = await app.get_apm().get_user_api();
    await uapi.add_user({
      user_id,
      full_name,
      email,
      group_id,
      password,
      pub_key,
    });

    app.navigator.goto_admin_usergroups_page();
  };
</script>

<Layout current_item={"user_groups"}>
  <div class="h-full w-full overflow-auto">
    <div class="md:p-12 bg-indigo-100 flex flex-row flex-wrap">
      <div
        class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg"
      >
        <div class="text-2xl text-indigo-900">User</div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Full Name</label>
          <input
            type="text"
            bind:value={full_name}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder="John"
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">User Id</label>
          <input
            type="text"
            bind:value={user_id}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder="user1"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Email</label>
          <input
            type="text"
            bind:value={email}
            placeholder="john@mail.com"
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Password</label>
          <input
            type="password"
            bind:value={password}
            placeholder="*********"
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Group</label>
          <input
            type="text"
            bind:value={group_id}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder="group1"
          />
        </div>

        <div class="flex justify-between space-x-1" on:click={onSave}>
          <button class="p-2 bg-blue-400 text-white rounded">Save</button>
        </div>
      </div>
    </div>
  </div>
</Layout>

<script lang="ts">
  import { getContext } from "svelte";
  import { AutoTable, ActionAddButton, PortalService } from "../core";

  const app: PortalService = getContext("__app__");
  const tapi = app.api_manager.get_admin_tenant_api();

  let domains = [];

  let tenant = {};
  let loaded = false;

  const load = async () => {
    const resp1 = tapi.get();
    const resp2 = tapi.get_domains();

    const [r1, r2] = await Promise.all([resp1, resp2]);

    tenant = r1.data;
    domains = r2.data;
    loaded = true;
  };

  load();
</script>

<div class="h-full w-full overflow-auto">
  <div class="md:p-12 bg-indigo-100 flex flex-row flex-wrap">
    <div
      class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg relative"
    >
      <div class="text-2xl text-indigo-900">Organization</div>
      {#if loaded}
        <div class="absolute right-0 p-2">
          <button
            on:click={() => app.nav.admin_tenant_edit()}
            class="p-1 bg-blue-400 text-white text-sm font-semibold flex rounded hover:scale-110"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                d="M17.414 2.586a2 2 0 00-2.828 0L7 10.172V13h2.828l7.586-7.586a2 2 0 000-2.828z"
              />
              <path
                fill-rule="evenodd"
                d="M2 6a2 2 0 012-2h4a1 1 0 010 2H4v10h10v-4a1 1 0 112 0v4a2 2 0 01-2 2H4a2 2 0 01-2-2V6z"
                clip-rule="evenodd"
              />
            </svg> Edit</button
          >
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Name</label>
          <input
            type="text"
            disabled
            value={tenant["name"] || ""}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Slug</label>
          <input
            type="text"
            value={tenant["slug"]}
            disabled
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">About</label>
          <textarea
            value={tenant["org_bio"] || ""}
            disabled
            class="border p-1 rounded focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div class="flex-col flex py-3 relative border rounded p-2">
          <div class="absolute right-1">
            <ActionAddButton
              onClick={() => app.nav.admin_tenant_domain_new()}
            />
          </div>

          <label class="pb-2 text-gray-700 font-semibold">Domains</label>

          <AutoTable
            action_key="id"
            actions={[
              {
                Name: "Adapter Editor",
                Action: (id) => app.nav.admin_tenant_domain_adapter_editor(id),
                icon: "lightning-bolt",
              },

              {
                Name: "Edit",
                Action: (id) => app.nav.admin_tenant_domain_edit(id),
                icon: "pencil",
              },
              {
                Name: "Delete",
                Class: "bg-red-400",
                icon: "trash",
                Action: async (id) => {
                  await tapi.delete_domain(id);
                  load();
                },
              },
            ]}
            key_names={[
              ["id", "ID"],
              ["name", "Name"],
              ["adapter_type", "Http Adapter"],
              ["about", "About"],
            ]}
            datas={domains}
            color={["adapter_type"]}
          />
        </div>
      {/if}
    </div>
  </div>
</div>

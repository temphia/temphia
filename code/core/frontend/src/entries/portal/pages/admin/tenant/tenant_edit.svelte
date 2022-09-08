<script lang="ts">
  import { getContext } from "svelte";

  import type { PortalApp } from "../../../../../lib/app/portal";
  import { KvEditor, ActionAddButton, ActionButton } from "../../../../common";

  import PlugPick from "../core/plug_pick.svelte";
  import UserGroupPick from "../core/user_group_pick.svelte";
  import Layout from "../layout.svelte";

  const app: PortalApp = getContext("__app__");

  let tenant = app.navigator.nav_options;
  let modified = false;
  let message = "";

  let getMetaData;
  let meta_modified = false;

  const set = (name) => (ev) => {
    tenant = { ...tenant, [name]: ev.target.value };
    modified = true;
  };

  const save = async () => {
    const data = { ...tenant };
    if (meta_modified) {
      data["extra_meta"] = getMetaData();
    }

    const tapi = await app.get_apm().get_tenant_id();
    const resp = await tapi.update_tenant(tenant);
    if (resp.status !== 200) {
      message = resp.data;
      return;
    }

    message = "";
    app.navigator.goto_admin_org();
  };

  const rootPick = () => {
    app.simple_modal_open(PlugPick, {
      app,
      selected_plug: tenant["root_plug_id"],
      selected_agent: tenant["root_agent_id"],
      selected_method: tenant["root_handler"],
      onSelected: (obj) => {
        tenant["root_plug_id"] = obj.plug;
        tenant["root_agent_id"] = obj.agent;
        tenant["root_handler"] = obj.method;
        tenant = tenant;
        modified = true;
        app.simple_modal_close();
      },
    });
  };

  const ugPick = () => {
    app.simple_modal_open(UserGroupPick, {
      app,
      selected_ugroup: tenant["default_ugroup"],
      onSelected: (ug) => {
        tenant["default_ugroup"] = ug;
        tenant = tenant;
        app.simple_modal_close();
      },
    });
  };
</script>

<Layout current_item={"ns"}>
  <div class="h-full w-full overflow-auto">
    <div class="md:p-12 bg-indigo-100 flex flex-row flex-wrap">
      <div
        class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg relative"
      >
        <div class="text-2xl text-indigo-900">Organization</div>

        <p class="text-red-500" />

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Name</label>
          <input
            type="text"
            value={tenant["name"] || ""}
            on:change={set("name")}
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
            on:change={set("org_bio")}
            class="border p-1 rounded focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <div class="absolute right-1">
            <ActionButton icon_name="link" name="pick" onClick={rootPick} />
          </div>

          <label class="pb-2 text-gray-700 font-semibold">Root Plug</label>
          <input
            type="text"
            value={tenant["root_plug_id"] || ""}
            on:change={set("root_plug_id")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Root Agent</label>
          <input
            type="text"
            value={tenant["root_agent_id"] || ""}
            on:change={set("root_agent_id")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Root Hanlder</label>
          <input
            type="text"
            value={tenant["root_handler"] || ""}
            on:change={set("root_handler")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <div class="flex-col flex py-3 relative">
            <label class="pb-2 text-gray-700 font-semibold">SMTP User</label>
            <input
              type="text"
              value={tenant["smtp_user"] || ""}
              on:change={set("smtp_user")}
              class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            />
          </div>

          <div class="flex-col flex py-3 relative">
            <label class="pb-2 text-gray-700 font-semibold">SMTP Passport</label
            >
            <input
              type="text"
              value={tenant["smtp_pass"] || ""}
              on:change={set("smtp_pass")}
              class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            />
          </div>

          <div class="flex-col flex py-3 relative">
            <div class="absolute right-1">
              <ActionButton icon_name="link" name="pick" onClick={ugPick} />
            </div>
            <label class="pb-2 text-gray-700 font-semibold"
              >Default User Group</label
            >

            <input
              type="text"
              value={tenant["default_ugroup"] || ""}
              on:change={set("default_ugroup")}
              class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            />
          </div>

          <div class="flex-col flex py-3 relative">
            <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
            <KvEditor
              data={tenant["extra_meta"] || {}}
              bind:getData={getMetaData}
              bind:modified={meta_modified}
            />
          </div>

          {#if modified}
            <div class="flex py-3">
              <button
                on:click={save}
                class="p-2 bg-blue-400 m-1 w-20 text-white rounded">Save</button
              >
            </div>
          {/if}
        </div>
      </div>
    </div>
  </div></Layout
>

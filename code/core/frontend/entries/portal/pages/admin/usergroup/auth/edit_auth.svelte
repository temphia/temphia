<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import Kveditor from "../../../../../common/kveditor.svelte";
  import Layout from "../../layout.svelte";

  export let gid = "";
  export let id = "";

  let data;
  let mod_data = {};

  const app: PortalApp = getContext("__app__");

  let getMetaData;
  let meta_modified;
  let modified;

  const modifyField = (field: string) => (ev) => {
    modified = true;
    mod_data[field] = ev.target["value"];
  };

  const load = async () => {
    const uapi = await app.get_apm().get_user_api();
    const resp = await uapi.user_group_get_auth(gid, Number(id));
    if (resp.status !== 200) {
      console.log("Err ", resp);
      return;
    }
    data = resp.data;
  };

  const save = async () => {
    const uapi = await app.get_apm().get_user_api();
    const resp = await uapi.user_group_update_auth(gid, Number(id), mod_data);
    if (resp.status !== 200) {
      console.log("Err ", resp);
      return;
    }
    app.navigator.goto_admin_usergroups_page();
  };

  load();
</script>

<Layout current_item={"user_groups"} loading={data === null}>
  {#if data}
    <div class="w-full h-full p-10">
      <div class="bg-white p-2">
        <div class="text-2xl text-indigo-900">Auth Provider</div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Id</label>
          <input
            type="text"
            disabled
            value={data["id"]}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Name</label>
          <input
            type="text"
            value={data["name"] || ""}
            on:change={modifyField("name")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Type</label>
          <input
            type="text"
            value={data["type"] || ""}
            on:change={modifyField("type")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Provider</label>
          <input
            type="text"
            value={data["provider"] || ""}
            on:change={modifyField("provider")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Auth URL</label>
          <input
            type="text"
            value={data["auth_url"] || ""}
            on:change={modifyField("auth_url")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>
        

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Token URL</label>
          <input
            type="text"
            value={data["token_url"] || ""}
            on:change={modifyField("token_url")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>
        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Scopes</label>
          <textarea
            value={data["scopes"] || ""}
            on:change={modifyField("scopes")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Client Id</label>
          <input
            type="text"
            value={data["client_id"] || ""}
            on:change={modifyField("client_id")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Client Secret</label>
          <input
            type="text"
            value={data["client_secret"] || ""}
            on:change={modifyField("client_secret")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Payload</label>
          <textarea
            value={data["payload"] || ""}
            on:change={modifyField("payload")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Policy</label>
          <textarea
            type="text"
            value={data["policy"] || ""}
            on:change={modifyField("policy")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
          <Kveditor
            data={data["extra_meta"] || {}}
            bind:getData={getMetaData}
            bind:modified={meta_modified}
          />
        </div>

        <div class="flex justify-end">
          <button
            on:click={save}
            class="p-2 bg-blue-400 hover:bg-blue-600 m-1 w-20 text-white rounded"
            >Save</button
          >
        </div>
      </div>
    </div>
  {/if}
</Layout>

<script lang="ts">
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import { CEditor } from "../../../../../../components";
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import OpenRPC from "./openrpc.svelte";
  import { DevTktAPI } from "../../../../../../lib/core/tktapi";

  export let pid: string;
  export let aid: string;
  export let tkt: string;
  export let app: PortalApp;

  let watch_mode = false;

  const tkt_api = new DevTktAPI(app._base_app._api_url, tkt);

  let method = "";
  let editor;
  let resp_data;

  const submit = async () => {
    console.log("@editor", editor);
    if (!method) {
      return;
    }
    resp_data = await tkt_api.exec_run(pid, aid, method, editor.getValue());
  };
</script>

<div class="h-full w-full p-2">
  <div
    class="w-full h-full flex flex-col p-4 rounded bg-white"
    style="height:calc(100vh - 4.5rem);"
  >
    <div class="p-1 flex justify-between h-14">
      <div class="p-2">
        <label class="uppercase text-gray-700 font-semibold text-lg">
          Method

          <input
            type="text"
            bind:value={method}
            class="p-2 h-full rounded border w-20 sm:w-28 md:w-48 lg:w-96 "
          />

          <button class="border rounded">
            <Icon name="chevron-down" class="h-5 w-5" />
          </button>
        </label>
      </div>
      <div class="flex gap-2 items-center">
        <button
          on:click={submit}
          class="p-1 rounded bg-green-500 shadow hover:bg-green-900 flex text-white"
        >
          <Icon name="play" class="h-6 w-6" solid />
          SUBMIT
        </button>

        <button
          on:click={() => {
            app.big_modal_open(OpenRPC, {});
          }}
          class="p-1 rounded bg-blue-300 shadow hover:bg-blue-600 flex text-white"
        >
          <Icon name="information-circle" class="h-6 w-6" solid />
        </button>
      </div>
    </div>

    <div class="flex flex-col h-96 flex-grow">
      <CEditor
        code={`{"data": 11}`}
        bind:editor
        container_style={"min-height:20rem;"}
      />
      <div class="h-1/2">
        <nav class="flex flex-col sm:flex-row">
          <button
            on:click={() => (watch_mode = false)}
            class="text-gray-600 p-1 block hover:text-blue-500 focus:outline-non {watch_mode
              ? ''
              : 'text-blue-500 border-b-2 font-medium border-blue-500'}"
          >
            Results
          </button>
          <button
            on:click={() => (watch_mode = true)}
            class="text-gray-600 p-1 block hover:text-blue-500 focus:outline-none  {watch_mode
              ? 'text-blue-500 border-b-2 font-medium border-blue-500'
              : ''}"
          >
            Watch
          </button>
        </nav>

        <div class="flex p-1">
          {#if !watch_mode}
            <div class="p-1 border rounded bg-gray-50 w-full">
              {#if resp_data}
                <pre class="p-1">{JSON.stringify(resp_data, null, 4)}</pre>
              {/if}
            </div>
          {:else}
            <div class="p-1 border rounded bg-gray-50 w-full">
              <span>Not watching anything</span>
            </div>
          {/if}
        </div>
      </div>
    </div>
  </div>
</div>

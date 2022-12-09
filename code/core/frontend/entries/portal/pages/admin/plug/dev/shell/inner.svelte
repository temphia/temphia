<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import OpenRPC from "../docs/_openrpc.svelte";

  import { PortalService, CEditor } from "../../../core";
  import type { DevShellService } from "../../../../../services/engine/dev_shell";
  import Jsonview from "../../../../../../xcompo/jsonview/jsonview.svelte";

  export let pid: string;
  export let aid: string;
  export let service: DevShellService;
  export let app: PortalService;

  let watch_mode = false;

  let method = "";
  let code = "{}";
  let message = "";
  let editor;
  let resp_data;

  const submit = async () => {
    if (!method) {
      message = "Empty method";
      return;
    }

    message = "";
    const resp = await service.dev_api.exec_run_agent_action(
      pid,
      aid,
      method,
      JSON.parse(editor.getValue() || "{}")
    );
    if (!resp.ok) {
      message = resp.data;
      resp_data = "";
      return;
    }

    resp_data = resp.data;

    putLocal();
  };

  const tryLocal = () => {
    if (window.localStorage) {
      const old = JSON.parse(
        localStorage.getItem("__temphia_portal_dev_shell_") || "{}"
      );
      if (old["code"]) {
        code = old["code"];
      }
      if (old["method"]) {
        method = old["method"];
      }
    }
  };

  const putLocal = () => {
    if (window.localStorage) {
      localStorage.setItem(
        "__temphia_portal_dev_shell_",
        JSON.stringify({
          code,
          message,
        })
      );
    }
  };

  tryLocal();
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
          on:click={() => app.utils.big_modal_open(OpenRPC, {})}
          class="p-1 rounded bg-blue-300 shadow hover:bg-blue-600 flex text-white"
        >
          <Icon name="information-circle" class="h-6 w-6" solid />
        </button>
      </div>
    </div>

    <p class="text-red-500">{message}</p>

    <div class="flex flex-col h-96 flex-grow">
      <CEditor {code} bind:editor container_style={"min-height:20rem;"} />
      <div class="flex flex-col flex-grow h-52">
        <nav class="flex">
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

        <div class="flex flex-col p-1 grow overflow-auto">
          {#if !watch_mode}
            <div class="p-1 border rounded bg-gray-50 w-full grow h-content">
              {#if resp_data}
                <Jsonview json={resp_data} />
              {/if}
            </div>
          {:else}
            <div class="p-1 border rounded bg-gray-50 w-full h-full">
              <span>Not watching anything</span>
            </div>
          {/if}
        </div>
      </div>
    </div>
  </div>
</div>

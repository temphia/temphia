<script lang="ts">
  import { getContext } from "svelte";
  import Ceditor from "../../../xcompo/ceditor/ceditor.svelte";
  import { LoadingSpinner } from "../core";
  import type { EasypageService } from "../service/easypage";
  import Preview from "./_editor/preview.svelte";
  import { params } from "svelte-hash-router";

  export let pid = $params.pid;

  let preview = false;
  let editor;
  let code = "## test";
  let loading = true;

  const service = getContext("__easypage_service__") as EasypageService;

  const load = async () => {
    const resp = await service.getPageData(pid);
    if (!resp.ok) {
      loading = false;
      return;
    }
    
    if (typeof resp.data === "object") {
      code = resp.data["code"] || "## test";
    } else if (typeof resp.data === "string") {
      code = JSON.parse(resp.data)["code"] || "## test";
    }

    loading = false;
  };

  load();
</script>

<div class="h-full bg-blue-50 overflow-auto">
  {#if loading}
    <LoadingSpinner />
  {:else}
    <div class="flex p-1 justify-between bg-white">
      <nav class="flex flex-row">
        <button
          on:click={() => {
            preview = false;
          }}
          class="text-gray-600 p-2 block hover:text-blue-500 focus:outline-none {!preview
            ? 'text-blue-500 border-b-2 font-medium border-blue-500'
            : ''}"
        >
          Markdown
        </button>

        <button
          on:click={() => {
            if (!preview) {
              code = editor.getValue();
              preview = true;
            }
          }}
          class="text-gray-600 p-2 block hover:text-blue-500 focus:outline-none {preview
            ? 'text-blue-500 border-b-2 font-medium border-blue-500'
            : ''}"
        >
          Preview
        </button>
      </nav>

      <div class="flex gap-2 justify-end p-1">
        <button
          class="p-1 rounded bg-gray-50 hover:bg-gray-200 border"
          on:click={async () => {
            loading = true;
            code = editor.getValue();
            await service.setPageData(pid, JSON.stringify({ code, type: "post" }));
            loading = false;
          }}
        >
          Save
        </button>
        <button
          on:click={() => {
            location.hash = "/";
          }}
          class="p-1 rounded bg-gray-50 hover:bg-gray-200 border"
        >
          Home
        </button>
      </div>
    </div>

    {#if preview}
      <Preview {code} />
    {:else}
      <Ceditor bind:editor {code} mode="md" container_style="height:100%;" />
    {/if}
  {/if}
</div>

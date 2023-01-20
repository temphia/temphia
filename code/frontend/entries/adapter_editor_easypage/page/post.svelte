<script lang="ts">
  import Ceditor from "../../xcompo/ceditor/ceditor.svelte";
  import Preview from "./_editor/preview.svelte";

  let preview = false;
  let editor;
  let code = "## test";
</script>

<div class="p-1 h-full bg-blue-50 overflow-auto">
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
      <button class="p-1 rounded bg-gray-50 hover:bg-gray-200 border">
        Save
      </button>
    </div>
  </div>

  {#if preview}
    <Preview {code} />
  {:else}
    <Ceditor
      bind:editor
      {code}
      mode="md"
      on:change={(ev) => {}}
      container_style="height:100%;"
    />
  {/if}
</div>

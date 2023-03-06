<script lang="ts">
  import { CEditor } from "../../admin/core";
  import { ClientJS, ServerJS } from "./conts";

  export let file = ServerJS;

  let client = "//client";
  let server = "//server";
  let editor;

  export const changeFile = (tofile) => {
    console.log("@change", { tofile, code, file });

    if (file === tofile) return;
    if (tofile === ClientJS) {
      if (editor) {
        server = editor.getValue();
      }
    } else if (tofile === ServerJS) {
      if (editor) {
        client = editor.getValue();
      }
    }

    file = tofile;
  };

  $: code = file === ServerJS ? server : client;
</script>

{#key code}
  <CEditor {code} bind:editor container_style={"height:100%;"} />
{/key}

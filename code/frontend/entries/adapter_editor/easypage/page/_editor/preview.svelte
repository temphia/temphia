<script lang="ts">
  import WithScript from "../../../portal/pages/admin/plug/dev/flowmap/_with_script.svelte";
  import PreviewRender from "./_preview_render.svelte";

  export let code = "";
</script>

<div class="h-full p-1 border">
  <WithScript script="/z/assets/lib/marked.js" key="marked">
    <PreviewRender {code} />
  </WithScript>
</div>

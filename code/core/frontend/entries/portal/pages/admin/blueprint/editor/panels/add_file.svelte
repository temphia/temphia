<script lang="ts">
  import { getContext } from "svelte";

  import { Uploader } from "../../../../../../../shared";
  export let beditor;

  const { close } = getContext("simple-modal");
</script>

<Uploader
  uploadFile={(file, data) => {
    beditor.uploadNewFile(file, data);
    beditor.onlyLoadFiles()
    close();
  }}
/>

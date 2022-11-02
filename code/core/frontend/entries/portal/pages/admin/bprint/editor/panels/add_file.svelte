<script lang="ts">
  import { Uploader } from "../../../core";

  export let beditor;
  export let close_modal;
</script>

<Uploader
  uploadFile={(file, data) => {
    beditor.uploadNewFile(file, data);
    beditor.onlyLoadFiles();
    close_modal();
  }}
/>

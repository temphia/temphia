<script lang="ts">
  import { getContext } from "svelte";
  import { PortalService, Uploader } from "$lib/core";

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_bprint_api();

  let loading = false;

  const save = async (data: FormData) => {
    loading = true;
    await api.create_from_zip(data);
    app.nav.admin_bprints();
  };
</script>

<div class="flex justify-center p-4">
  <div class="bg-white rounded p-2 " style="width: 725px;">
    <Uploader placeholder="bprint.zip" uploadFile={(fname, data) => save(data)} />
  </div>
</div>

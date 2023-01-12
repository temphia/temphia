<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { PortalService } from "../../core";
  import InstanceBundlePicker from "../instancer/bundle_picker.svelte";
  import { instance_helper } from "../instancer/instance";

  export let app: PortalService;
  export let id: string;

  const manual = async () => {
    const api = app.api_manager.get_admin_bprint_api();
    const resp = await api.get(id);
    if (!resp.ok) {
      console.log("@@");
      return;
    }

    const bprint = resp.data;

    const file = bprint["files"].filter(
      (v) => v !== "schema.json" || v !== "schema.yaml"
    )[0];

    app.utils.small_modal_close();
    instance_helper(app, bprint["type"], bprint, file, InstanceBundlePicker);
  };

  const automatic = () => {
    app.utils.small_modal_close();
    app.nav.admin_bprint_auto_instancer(id);
  };
</script>

<div class="flex">
  <div class="w-1/3 pt-6 flex justify-center">
    <Icon
      name="document-download"
      class="w-16 h-16 bg-blue-600 text-white p-3 rounded-full"
    />
  </div>
  <div class="w-full pt-9 pr-4">
    <h3 class="font-bold text-blue-700">
      How do you want to instance the Blueprint?
    </h3>
    <p class="py-4 text-sm text-gray-400">
      Automatic instance creates all necessary objects and resources with
      sensible defaults or you can instance manually.
    </p>
  </div>
</div>

<div class="p-4 flex space-x-4">
  <button
    on:click={manual}
    class="w-1/2 px-4 py-3 text-center bg-gray-100 text-gray-400 hover:bg-gray-200 hover:text-black font-bold rounded-lg text-sm"
    >Manually</button
  >

  <button
    on:click={automatic}
    class="w-1/2 px-4 py-3 text-center text-blue-100 bg-blue-600 rounded-lg hover:bg-blue-700 hover:text-white font-bold text-sm"
    >Automatic</button
  >
</div>

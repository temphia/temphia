<script lang="ts">
  import { getContext } from "svelte";
  import { CEditor } from "../../../../../../components";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import InstallerLayout from "../_layout.svelte";
  import { installAsDgroup } from "./installer";

  export let bid;
  export let data;
  export let schema;

  export let sources = [];

  const app: PortalApp = getContext("__app__");

  let installer;
  const bapi = app.get_apm().get_bprint_api().then((bapi) => {
    installer = installAsDgroup(bapi);
  });

  let groupSlug = "";
  let groupName = "";
  let editor;
  let seed_from = "";

  $: _source = sources[0];
  $: _schema_changed = false;
</script>

<InstallerLayout
  last_stage_name="Install"
  description={data["description"]}
  files={data["files"]}
  type={data["type"]}
  subtype={data["sub_type"]}
  name={data["name"]}
  slug={data["slug"]}
  source=""
  final_func={app.big_modal_close}
  last_page_func={() =>
    installer({
      bprint_id: bid,
      cabinet_folder: "",
      cabinet_source: "",
      name: groupName,
      slug: groupSlug,
      schema: _schema_changed ? editor.getValue() : "",
      seed_from,
    })}
>
  <div slot="options" class="h-full w-full">
    <div class="flex flex-col relative">
      <button
        class="p-1 text-gray-500 right-1 top-1/2 border absolute bg-white rounded"
        >Check</button
      >
      <label class="leading-loose"
        >Slug
        <input
          type="text"
          bind:value={groupSlug}
          class="px-4 py-2 border w-full sm:text-sm rounded focus:outline-none focus:border-indigo-500"
          placeholder="Optional"
        />
      </label>
    </div>

    <div class="flex flex-col mt-5">
      <label class="leading-loose"
        >Group name
        <input
          bind:value={groupName}
          type="text"
          class="px-4 py-2 border w-full sm:text-sm rounded focus:outline-none focus:border-indigo-500"
          placeholder="My data collection 1"
        />
      </label>
    </div>

    <div class="flex flex-col relative mt-5">
      <label for="" class="leading-loose">Database source </label>
      <select class="p-1 w-20" bind:value={_source}>
        {#each sources as source}
          <option>{source}</option>
        {/each}
      </select>
    </div>

    <div class="flex flex-col relative mt-5">
      <button
        class="p-1 text-gray-500 right-1 top-1/2 border absolute bg-white rounded"
        >Select</button
      >
      <label class="leading-loose"
        >Cabinet source
        <input
          type="text"
          disabled
          class="px-4 py-2 border w-full sm:text-sm rounded focus:outline-none focus:border-indigo-500"
          placeholder="cabinet/folder"
        />
      </label>
    </div>

    <div class="flex flex-col relative mt-5">
      <label class="leading-loose">Seed From </label>

      <select bind:value={seed_from}>
        <option value="">None</option>
        <option value="data">Data File</option>
        <option value="autogen">Auto Generate</option>
      </select>
    </div>

    <div class="flex-col flex py-3 mt-5">
      <legend class="text-base font-medium text-gray-900 mt-5">Schema</legend>

      <div class="p-2 bg-white border">
        <CEditor
          bind:editor
          code={JSON.stringify(schema, null, 4)}
          container_style="height:20rem;"
          on:change={(ev) => {
            _schema_changed = true;
          }}
        />
      </div>
    </div>
  </div>

  <div slot="final" class="">Final</div>
</InstallerLayout>

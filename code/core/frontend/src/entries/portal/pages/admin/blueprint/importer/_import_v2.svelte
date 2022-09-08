<script>
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { ToggleButton } from "../../../../../../components";

  import { generateId } from "../../../../../../lib/utils";
  import { BetterTextInput, PrimaryButton } from "../../../../../common";

  export let source;
  export let group;
  export let data;
  export let importFunc;

  const all_files = data["files"] || [];
  let files = [];

  let new_id = (data["slug"] || "") + "_" + generateId();

  let loading = true;
</script>

<div class="space-y-4">
  <BetterTextInput
    {loading}
    label="slug"
    value={new_id}
    placeholder="my_brpint_123"
  />
  <div class="space-y-2">
    <legend class=" text-base  text-1.5xl font-medium text-gray-900">
      Files
    </legend>

    <div class="flex items-start flex-col gap-2">
      {#each all_files as file, idx}
        <div class="flex items-center h-5">
          <input
            id="file-{idx}"
            name="comments"
            checked={!files.includes(file)}
            on:click={() => {
              if (files.includes(file)) {
                files = files.filter((v) => file !== v);
              } else {
                files = [...files, file];
              }
            }}
            type="checkbox"
            class="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 rounded"
          />
          <label for="file-{idx}" class=" ml-3 font-mediuregular text-gray-700"
            >{file}</label
          >
        </div>
      {/each}

      <span class="text-xs italic">
        * Recommended importing all files, unless you know what you are doing.
      </span>
    </div>
  </div>

  <div>
    <legend class=" text-base  text-1.5xl font-medium text-gray-900">
      Instance after import?
    </legend>

    <ToggleButton checked />
  </div>

  <div class="flex justify-end">
    <PrimaryButton
      icon="arrow-circle-down"
      label="IMPORT"
      onClick={async () => {
        const resp = await importFunc({
          source: Number(source),
          group,
          slug: data.slug,
          skip_files: files,
          new_id: new_id,
        });

        console.log("@import", resp);
        //  importFunc({})
      }}
    />
  </div>
</div>

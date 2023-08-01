<script lang="ts">
  import { createEventDispatcher } from "svelte";
  export let table;

  $: _options = {};

  const dispatch = createEventDispatcher();

  const onChange = () => {
    dispatch("table_change", _options);
    console.log("@table_options", _options)
  };
</script>

<tr class="hover:bg-grey-lighter">
  <td class="p-1 border-b border-grey-light">{table.name}</td>
  <td class="p-1 border-b border-grey-light">{table.slug}</td>
  <td class="p-1 border-b border-grey-light">
    {table.description}
  </td>

  
  <td class="p-1 border-b border-grey-light">
    <select
      class="p-1 rounded bg-slate-300"
      on:change={(ev) => {
        _options["sync_type"] = ev.target["value"];
        onChange();
      }}
      value={_options["sync_type"] || ""}
    >
      <option>strict</option>
      <option>lazy</option>
      <option>none</option>
    </select>
  </td>
  <td class="p-1 border-b border-grey-light">
    <select
      class="p-1 rounded bg-slate-300"
      on:change={(ev) => {
        _options["activity_type"] = ev.target["value"];
        onChange();
      }}
      value={_options["activity_type"] || ""}
    >
      <option>none</option>
      <option>event_only</option>
      <option>event_and_data</option>
    </select>
  </td>

  <td class="p-1 border-b border-grey-light">
    <input
      type="checkbox"
      on:change={() => {
        _options["seed"] = !_options["seed"];
        onChange()
      }}
      checked={!!_options["seed"]}
    />
  </td>
</tr>

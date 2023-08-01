<script lang="ts">
  import { HSplitPane, VSplitPane } from "svelte-split-pane";
  import Ceditor from "../../../xcompo/ceditor/ceditor.svelte";
  let tableresp = {
    columns: [
      {
        name: "What msg",
        slug: "wm",
        ctype: "shorttext",
        description: "what msg here",
        group_id: "pqqq",
        options: [],
        not_nullable: true,
        table_id: "what",
        tenant_id: "demo3",
      },
      {
        name: "Notes",
        slug: "notes",
        ctype: "shorttext",
        description: "this have notes",
        group_id: "pqqq",
        options: [],
        table_id: "what",
        tenant_id: "demo3",
      },
      {
        name: "Options",
        slug: "opts",
        ctype: "select",
        description: "i like options",
        group_id: "pqqq",
        options: ["apple", "banana", "orange"],
        table_id: "what",
        tenant_id: "demo3",
      },
    ],
    rows: [
      {
        __id: 1,
        __mod_sig: null,
        __version: 0,
        created_at: "2022-01-06T23:57:04.712383Z",
        deleted_at: null,
        notes: "jean shorts",
        opts: "apple",
        updated_at: "2022-01-06T23:57:04.712383Z",
        wm: "ennui",
      },
      {
        __id: 2,
        __mod_sig: null,
        __version: 0,
        created_at: "2022-01-06T23:57:04.714058Z",
        deleted_at: null,
        notes: "yr",
        opts: "banana",
        updated_at: "2022-01-06T23:57:04.714058Z",
        wm: "fingerstache",
      },
      {
        __id: 3,
        __mod_sig: null,
        __version: 0,
        created_at: "2022-01-06T23:57:04.715327Z",
        deleted_at: null,
        notes: "pabst",
        opts: "apple",
        updated_at: "2022-01-06T23:57:04.715327Z",
        wm: "kickstarter",
      },
      {
        __id: 4,
        __mod_sig: null,
        __version: 0,
        created_at: "2022-01-06T23:57:04.716679Z",
        deleted_at: null,
        notes: "artisan",
        opts: "orange",
        updated_at: "2022-01-06T23:57:04.716679Z",
        wm: "cornhole",
      },
      {
        __id: 5,
        __mod_sig: null,
        __version: 0,
        created_at: "2022-01-06T23:57:04.718163Z",
        deleted_at: null,
        notes: "mumblecore",
        opts: "banana",
        updated_at: "2022-01-06T23:57:04.718163Z",
        wm: "salvia",
      },
      {
        __id: 6,
        __mod_sig: null,
        __version: 0,
        created_at: "2022-01-06T23:57:04.719478Z",
        deleted_at: null,
        notes: "truffaut",
        opts: "orange",
        updated_at: "2022-01-06T23:57:04.719478Z",
        wm: "readymade",
      },
      {
        __id: 7,
        __mod_sig: null,
        __version: 0,
        created_at: "2022-01-06T23:57:04.720898Z",
        deleted_at: null,
        notes: "roof",
        opts: "apple",
        updated_at: "2022-01-06T23:57:04.720898Z",
        wm: "lo-fi",
      },
      {
        __id: 8,
        __mod_sig: null,
        __version: 0,
        created_at: "2022-01-06T23:57:04.722984Z",
        deleted_at: null,
        notes: "keffiyeh",
        opts: "banana",
        updated_at: "2022-01-06T23:57:04.722984Z",
        wm: "fixie",
      },
      {
        __id: 9,
        __mod_sig: null,
        __version: 0,
        created_at: "2022-01-06T23:57:04.724297Z",
        deleted_at: null,
        notes: "kombucha",
        opts: "banana",
        updated_at: "2022-01-06T23:57:04.724297Z",
        wm: "migas",
      },
      {
        __id: 10,
        __mod_sig: null,
        __version: 0,
        created_at: "2022-01-06T23:57:04.725796Z",
        deleted_at: null,
        notes: "messenger bag",
        opts: "banana",
        updated_at: "2022-01-06T23:57:04.725796Z",
        wm: "put a bird on it",
      },
      {
        __id: 11,
        __mod_sig: null,
        __version: 0,
        created_at: "2022-01-06T23:57:04.727164Z",
        deleted_at: null,
        notes: "mixtape",
        opts: "orange",
        updated_at: "2022-01-06T23:57:04.727164Z",
        wm: "selvage",
      },
    ],
  };
</script>

<VSplitPane
  updateCallback={() => {
    console.log("VSplitPane Updated!");
  }}
>
  <svelte:fragment slot="top">
    <Ceditor container_style={"height:100%;"} />
  </svelte:fragment>

  <svelte:fragment slot="down">
    <div class="flex justify-end gap-x-2 p-1 h-10">
      <label for="" class="p-1">Examples</label>
      <select>
        <option />
        <option>ex1</option>
        <option>ex1</option>
      </select>

      <button
        class="bg-red-100 hover:bg-red-600 w-20 text-red-600 text-sm hover:text-white rounded"
        >Run</button
      >
    </div>

    <div class="p-2 w-full overflow-auto">
      <table class="w-full border">
        <thead>
          <tr
            class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal"
          >
            {#each tableresp.columns as col}
              <th class="py-3 px-6 text-left">{col.name}</th>
            {/each}
          </tr>
        </thead>
        <tbody class="text-gray-600 text-sm font-light">
          {#each tableresp.rows as row}
            <tr class="border-b border-gray-200 hover:bg-gray-100">
              {#each tableresp.columns as col}
                <td class="py-3 px-6 text-left">{row[col.slug] || ""}</td>
              {/each}
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </svelte:fragment>
</VSplitPane>

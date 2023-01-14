<script lang="ts">
  import type { AgentIface } from "./iface";
  import Value from "./_value_type.svelte";
  import SectionLayout from "./_section_layout.svelte";
  export let data: AgentIface;
</script>

<div class="flex flex-col">
  <SectionLayout name="Methods">
    <table class="text-left border w-full">
      <thead>
        <tr>
          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Name</th
          >
          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Info</th
          >

          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Arg</th
          >

          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Return Type</th
          >

          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Errors</th
          >
        </tr>
      </thead>
      <tbody>
        {#each Object.entries(data.methods || {}) as [mkey, mdata]}
          <tr class="hover:bg-grey-lighter">
            <td class="py-2 px-3 border-b border-grey-light">{mkey}</td>
            <td class="py-2 px-3 border-b border-grey-light">
              {mdata.info}
            </td>

            <td class="py-2 px-3 border-b border-grey-light">
              {#if mdata.arg}
                <Value data={mdata.arg} />
              {/if}
            </td>

            <td class="py-2 px-3 border-b border-grey-light">
              {#if mdata.return_type}
                <Value data={mdata.return_type} />
              {/if}
            </td>

            <td class="py-2 px-3 border-b border-grey-light">
              <div class="flex flex-col">
                {#each Object.entries(mdata.error_types || {}) as [ekey, edata]}
                  <div>
                    <span class="rounded bg-slate-100 ">{ekey}</span>
                    {edata}
                  </div>
                {/each}
              </div>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </SectionLayout>

  <SectionLayout name="Events">
    <table class="text-left border w-full">
      <thead>
        <tr>
          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Name</th
          >
          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Info</th
          >

          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Async</th
          >

          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Arg</th
          >

          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Return Type</th
          >
        </tr>
      </thead>
      <tbody>
        {#each Object.entries(data.events || {}) as [mkey, mdata]}
          <tr class="hover:bg-grey-lighter">
            <td class="py-2 px-3 border-b border-grey-light">{mkey}</td>
            <td class="py-2 px-3 border-b border-grey-light">
              {mdata.info || ""}
            </td>

            <td class="py-2 px-3 border-b border-grey-light">
              {mdata.async || "false"}
            </td>

            <td class="py-2 px-3 border-b border-grey-light">
              {#if mdata.arg_data}
                <Value data={mdata.arg_data} />
              {/if}
            </td>

            <td class="py-2 px-3 border-b border-grey-light">
              {#if mdata.return_data}
                <Value data={mdata.return_data} />
              {/if}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </SectionLayout>

  <SectionLayout name="Schemas">
    <table class="text-left border w-full">
      <thead>
        <tr>
          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Name</th
          >
          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Type</th
          >

          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Property</th
          >

          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Ref</th
          >
        </tr>
      </thead>
      <tbody>
        {#each Object.entries(data.schemas || {}) as [mkey, mdata]}
          <tr id="#schema-schema-{mkey}" class="hover:bg-grey-lighter">
            <td class="py-2 px-3 border-b border-grey-light">{mkey}</td>

            <td class="py-2 px-3 border-b border-grey-light">
              {mdata.type || ""}
            </td>

            <td class="py-2 px-3 border-b border-grey-light">
              {mdata.property || ""}
            </td>

            <td class="py-2 px-3 border-b border-grey-light">
              {#if mdata.ref}
                <a href="#schema-{mdata.ref}">{mdata.ref || ""}</a>
              {/if}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </SectionLayout>

  <SectionLayout name="Definations">
    <table class="text-left border w-full">
      <thead>
        <tr>
          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Name</th
          >

          <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >Value</th
          >
        </tr>
      </thead>
      <tbody>
        {#each Object.entries(data.definations || {}) as [mkey, mdata]}
          <tr id="#schema-schema-{mkey}" class="hover:bg-grey-lighter">
            <td class="py-2 px-3 border-b border-grey-light">{mkey}</td>

            <td class="py-2 px-3 border-b border-grey-light">
              <details>
                <summary>Data</summary>
                <pre class="bg-gray-100 rounded">{JSON.stringify(mdata)}</pre>
              </details>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </SectionLayout>
</div>

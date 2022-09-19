<script lang="ts">
  import { FloatingAdd } from "../../../_shared";
  import { deleteTenant, goto, goto_tenant, listTenant } from "../../operator";

  let tenants = [];
  const load = () => {
    listTenant().then((data) => {
      if (typeof data !== "object") {
        return;
      }
      tenants = data;
    });
  };

  load();
</script>

<div class="p-8">
  <div>
    <h2 class="text-2xl font-semibold leading-tight">Tenants</h2>
  </div>

  <div class="inline-block min-w-full shadow rounded-lg overflow-hidden mt-4">
    <table class="min-w-full leading-normal">
      <thead>
        <tr>
          <th
            class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
          >
            Name
          </th>
          <th
            class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
          >
            Slug
          </th>
          <th
            class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
          >
            Actions
          </th>
        </tr>
      </thead>
      <tbody>
        {#each tenants as ten}
          <tr>
            <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
              <div class="flex items-center">
                <div class="flex-shrink-0 w-10 h-10">
                  <img class="w-24 h-AUTO rounded-full" alt="" />
                </div>
                <div class="ml-3">
                  <p class="text-gray-900 whitespace-no-wrap">
                    {ten.name}
                  </p>
                </div>
              </div>
            </td>
            <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
              <p class="text-gray-900 whitespace-no-wrap">{ten.slug}</p>
            </td>

            <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
              <button
                on:click={() => goto_tenant(ten.slug)}
                class="p-1 m-1 text-sm font-semibold text-white rounded transform hover:scale-110 bg-blue-400"
                >Edit</button
              >
              <button
                class="p-1 m-1 text-sm font-semibold text-white rounded transform hover:scale-110 bg-yellow-400"
                >Ensure</button
              >

              <button
                class="p-1 m-1 text-sm font-semibold text-white rounded transform hover:scale-110 bg-green-400"
                >Enter</button
              >
              <button
                on:click={async () => {
                  await deleteTenant(ten.slug);
                  load();
                }}
                class="p-1 m-1 text-sm font-semibold text-white rounded transform hover:scale-110 bg-red-400"
                >Delete</button
              >
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>

<FloatingAdd onClick={goto("/z/operator/tenant_new")} />

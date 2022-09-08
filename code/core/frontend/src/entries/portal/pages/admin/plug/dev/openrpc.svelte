<script>
  import PetOpenrpc from "./_pet.openrpc.json";
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";

  // https://spec.open-rpc.org/
  // https://raw.githubusercontent.com/open-rpc/meta-schema/master/schema.json

  const info = PetOpenrpc.info || {};
  const components =  PetOpenrpc.components || {};

  const methods = PetOpenrpc.methods || [];
  const schemas = components.schemas || {};
  const content_schemas = components.contentDescriptors || {};
</script>

<div class="w-full h-full p-2">
  <div class="w-full rounded bg-white py-4 px-2">
    <div class="pb-4 border-b-1 ">
      <h1 class="mb-1 text-3xl text-zinc-600">{info.title}</h1>
      <div class="flex justify-start">
        <span class="rounded bg-gray-50 p-1 text-sm"
          >Version {info.version}</span
        >
        <span class="flex ml-1">
          <svg
            class="h-6 w-6 flex-none fill-sky-100 stroke-sky-500 stroke-2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"
            />
          </svg>
          <p class="ml-1">{info.license.name}</p>
        </span>
      </div>
    </div>

    <div class="grid grid-cols-6 gap-2 rounded h-full">
      <div class="col-span-1 bg-gray-50 rounded">
        <ul class="px-4 py-6 space-y-3">
          <li class="block font-medium text-gray-500 dark:text-gray-300">
            <span class="cursor-pointer hover:underline flex text-zinc-600">
              Methods
              <Icon name="code" class="h-4 w-4 ml-1 mt-1" />
            </span>

            <ul class="px-4 py-1 space-y-3">
              {#each methods as method}
                <li
                  class="block font-medium text-gray-500 dark:text-gray-300 hover:underline cursor-pointer"
                >
                  {method.name}
                </li>
              {/each}
            </ul>
          </li>
          <li class="block font-medium text-gray-500 dark:text-gray-300">
            <span class="cursor-pointer hover:underline flex text-zinc-600">
              Content Descriptors
              <Icon name="chevron-double-right" class="h-4 w-4 ml-1 mt-1" />
            </span>

            <ul class="px-4 py-1 space-y-1">
              {#each Object.keys(content_schemas) as skey}
                <li
                  class="block font-medium text-gray-500 dark:text-gray-300 hover:underline cursor-pointer"
                >
                  {skey}
                </li>
              {/each}
            </ul>
          </li>
          <li class="block font-medium text-gray-500 dark:text-gray-300">
            <span class="cursor-pointer hover:underline flex text-zinc-600">
              Schemas
              <Icon name="database" class="h-4 w-4 ml-1 mt-1" />
            </span>

            <ul class="px-4 py-1 space-y-1">
              {#each Object.keys(schemas) as skey}
                <li
                  class="block font-medium text-gray-500 dark:text-gray-300 hover:underline cursor-pointer"
                >
                  {skey}
                </li>
              {/each}
            </ul>
          </li>
        </ul>
      </div>
      <div class="col-span-5">
        <div class="flex flex-col p-1">
          <h1 class="text-2xl font-semibold text-gray-700 capitalize mb-4">
            Methods
          </h1>
          <table class="text-left w-full border ">
            <thead>
              <tr>
                <th
                  class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >Name</th
                >
                <th
                  class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >Summary</th
                >

                <th
                  class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >Params</th
                >

                <th
                  class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >Result</th
                >

                <th
                  class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >Tags</th
                >
              </tr>
            </thead>
            <tbody>
              {#each methods as method}
                <tr class="hover:bg-grey-lighter">
                  <td class="py-2 px-3 border-b border-grey-light"
                    >{method.name}</td
                  >
                  <td class="py-2 px-3 border-b border-grey-light">
                    {method.summary}
                  </td>

                  <td class="py-2 px-3 border-b border-grey-light">
                    <div class="flex-col">
                      {#each method.params as param}
                        <div class="p-1 rounded border">
                          {#if param.$ref}
                            <a href="#">{param.$ref}</a>
                          {:else}
                            <h3 class="font-semibold text-lg text-zinc-900">
                              {param.name}
                            </h3>
                            <p>{param.description}</p>

                            <pre class="bg-gray-100 rounded">{JSON.stringify(
                                param.schema || "",
                                null,
                                4
                              )}</pre>
                          {/if}
                        </div>
                      {/each}
                    </div>
                  </td>

                  <td class="py-2 px-3 border-b border-grey-light">
                    <div class="flex flex-col">
                      <div class="border rounded p-1">
                        {#if method.result !== undefined}
                          <h2 class="font-semibold text-lg text-zinc-900 ">
                            Result
                          </h2>
                          <h3 class="text-lg text-zinc-900">
                            {method.result.name || ""}
                          </h3>
                          <p>{method.result.description || ""}</p>
                          {#if method.result.schema !== undefined}
                            <pre class="bg-gray-100 rounded">
                              {JSON.stringify(method.result.schema, null, 4)}
                            </pre>
                          {/if}
                        {/if}
                      </div>

                      <div class="border rounded p-1">
                        {#if method.errors}
                          <h2 class="font-semibold text-lg text-zinc-900 ">
                            Errors
                          </h2>

                          {#if method.errors !== undefined}
                            <pre class="bg-gray-100 rounded">{JSON.stringify(
                                method.errors,
                                null,
                                4
                              )}</pre>
                          {/if}
                        {/if}
                      </div>
                    </div>
                  </td>

                  <td class="py-2 px-3 border-b border-grey-light">
                    {#each method.tags as tag}
                      <span
                        class="bg-blue-100 text-xs text-blue-900 rounded p-1"
                        >{tag.name}</span
                      >
                    {/each}
                  </td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>

        <div class="flex flex-col p-1">
          <h1 class="text-2xl font-semibold text-gray-700 capitalize mb-4">
            Content Descriptors
          </h1>
          <table class="text-left w-full border ">
            <thead>
              <tr>
                <th
                  class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >Name</th
                >
                <th
                  class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >Summary</th
                >

                <th
                  class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >Schema</th
                >
              </tr>
            </thead>
            <tbody>
              {#each Object.entries(content_schemas) as [skey, schema]}
                <tr class="hover:bg-grey-lighter">
                  <td class="py-2 px-3 border-b border-grey-light"
                    >{schema.name}</td
                  >
                  <td class="py-2 px-3 border-b border-grey-light">
                    {schema.description}
                  </td>

                  <td class="py-2 px-3 border-b border-grey-light">
                    <a href="#">{schema.schema.$ref || ""}</a>
                  </td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>

        <div class="flex flex-col p-1">
          <h1 class="text-2xl font-semibold text-gray-700 capitalize mb-4">
            Schemas
          </h1>
          <table class="text-left w-full border ">
            <thead>
              <tr>
                <th
                  class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >Name</th
                >
                <th
                  class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >Schema</th
                >
              </tr>
            </thead>
            <tbody>
              {#each Object.entries(schemas) as [skey, schema]}
                <tr class="hover:bg-grey-lighter">
                  <td class="py-2 px-3 border-b border-grey-light">{skey}</td>
                  <td class="py-2 px-3 border-b border-grey-light">
                    <pre class="bg-gray-100 rounded">{JSON.stringify(
                        schema,
                        null,
                        4
                      )}</pre>
                  </td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</div>

<style>
  a {
    text-decoration: underline;
    color: blueviolet;
  }
</style>

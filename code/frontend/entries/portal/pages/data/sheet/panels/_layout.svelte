<script>
  export let title;
  export let onClick;
</script>

<div class="flex flex-col h-full">
  <div class="mb-4 border-b">
    <h1 class="block mb-2 text-lg font-bold text-gray-700" for="slug">
      {title}
    </h1>
  </div>

  <div class="flex flex-col flex-grow gap-2">
    <slot />
  </div>

  <div class="flex justify-end border-t">
    <button
      on:click={onClick}
      class="px-2 py-1 font-bold text-white bg-blue-500 hover:bg-blue-700 focus:outline-none focus:shadow-outline rounded mt-1"
      >Save</button
    >
  </div>
</div>

<div class="hidden">
  <div class="border-l-red-400" />
  <div class="border-l-blue-400" />
  <div class="border-l-green-400" />
  <div class="border-l-gray-400" />
  <div class="border-l-yellow-400" />
  <div class="border-l-zinc-400" />
  <div class="border-l-emerald-400" />
  <div class="border-l-orange-400" />
  <div class="border-l-lime-400" />
  <div class="border-l-amber-400" />
  <div class="border-l-teal-400" />
  <div class="border-l-cyan-400" />
  <div class="border-l-sky-400" />
  <div class="border-l-indigo-400" />
  <div class="border-l-violet-400" />
  <div class="border-l-purple-400" />
  <div class="border-l-fuchsia-400" />
  <div class="border-l-pink-400" />
</div>

<script>
  export let ticket;
  export let close;
</script>

<div class="w-full relative flex flex-col justify-center items-center bg-white">
  <div>
    <h1
      role="main"
      class="text-3xl lg:text-4xl font-semibold leading-7 lg:leading-9 text-center text-gray-800"
    >
      Your developer Token
    </h1>
  </div>
  <div class="w-full p-4">
    <pre
      class="overflow-scroll rounded bg-gray-200 h-full"
      style="max-width: 95%; margin:auto; ">{JSON.stringify(ticket, null, 2)}</pre>
  </div>
  <button
    on:click={close}
    class="mt-5 px-2 py-1 flex text-white font-semibold rounded bg-green-400 hover:bg-green-600"
    >OK</button
  >
</div>

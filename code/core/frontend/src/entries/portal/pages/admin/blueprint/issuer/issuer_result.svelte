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
      class="overflow-scroll h-10 rounded bg-gray-200"
      style="max-width: 95%; margin:auto;">{ticket}</pre>
  </div>
  <button
    on:click={close}
    class="mt-5  text-base leading-4 text-center text-white p-4 rounded focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-800 bg-gray-800 hover:bg-black"
    >OK</button
  >
</div>

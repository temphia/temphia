<script>
  let breatheText = "";
  let running = false;

  const delay = (ms) => new Promise((res) => setTimeout(res, ms));

  async function listener(ev) {
    breatheText = "Breathe In";
    await delay(4000);
    breatheText = "Hold";
    await delay(4000);
    breatheText = "Breathe Out";
    await delay(4000);
    breatheText = "Hold";
    await delay(4000);
  }
</script>

{#if running}
  {#key running}
    <div class="h-full w-full flex justify-center" id="main">
      <div class="container p-40 flex flex-col">
        <div
          class="square box-border mx-auto my-0  bg-transparent border-2 border-black "
          style="width: 40vh;height: 40vh;"
        >
          <div
            on:animationstart={listener}
            on:animationiteration={listener}
            class="circle bg-white rounded-full"
          />
          <div id="breatheText">{breatheText}</div>
        </div>

        <div class="mt-10 flex justify-center">
          <button
            on:click={() => {
              running = false;
              breatheText = ""
            }}
            class="p-2 rounded bg-red-500 hover:bg-red-700 mt-10 text-white"
            >stop</button
          >
        </div>
      </div>
    </div>
  {/key}
{:else}
  <div
    class="h-full w-full flex justify-center justify-items-center content-center items-center"
  >
    <button
      on:click={() => {
        running = true;
      }}
      class="p-2 rounded bg-blue-500 hover:bg-blue-700 h-10 text-white"
      >Start</button
    >

    <div class="fixed bottom-10">
      <a href="https://github.com/JT4A/Breathing-Excercise">Adapted from</a>
    </div>
  </div>
{/if}

<style>
  #main {
    background: red;
    background: linear-gradient(360deg, #0d0e3a, #2196f3);
    background-size: 400% 400%;
    animation: AnimationName 16s ease infinite;
  }

  .circle {
    height: 5vh;
    width: 5vh;
    animation: translate 16s infinite;
    animation-direction: normal;
    animation-timing-function: linear;
  }

  #breatheText {
    padding-top: 8vh;
    font-size: 6vh;
    color: White;
    margin: 0 auto;

    text-align-last: center;
    vertical-align: middle;
  }

  @keyframes translate {
    0%,
    100% {
      transform: translate(-3vh, 36vh) scale(1);
    }
    25% {
      transform: translate(-3vh, -3vh) scale(2);
    }
    50% {
      transform: translate(36vh, -3vh) scale(2);
    }
    75% {
      transform: translate(36vh, 36vh) scale(1);
    }
  }

  @keyframes AnimationName {
    0% {
      background-position: 50% 0%;
    }
    50% {
      background-position: 50% 65%;
    }
    100% {
      background-position: 50% 0%;
    }
  }
</style>

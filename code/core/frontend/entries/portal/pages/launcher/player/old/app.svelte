<script>
  import { sleep } from "yootils";

  let counter = 0;

  export let name = "";

  (async () => {
    while (true) {
      counter = counter + 1;
      await sleep(5000);
      console.log("@APP WAKE", name, "@couner", counter);
      if (counter > 999999999) break;
    }
  })();
</script>

<h2>{"APP =>" + name}</h2>

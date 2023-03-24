import { registerExecLoaderFactory } from "../../../lib/engine/register";
import PageDash from "./index.svelte";

registerExecLoaderFactory("pagedash.loader", (opts) => {
  console.log("@pagedash.loader", opts);

  new PageDash({
    target: opts.target,
    props: {
      env: opts.env,
    },
  });
});

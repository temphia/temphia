import { registerExecLoaderFactory } from "../../../lib/engine/register";
import PageDash from "./index.svelte";

registerExecLoaderFactory("pagequery.loader", (opts) => {
  console.log("@pagequery.loader", opts);

  new PageDash({
    target: opts.target,
    props: {
      env: opts.env,
    },
  });
});

import PageDash from "./pagedash/index.svelte";
import { registerExecLoaderFactory } from "../lib";

registerExecLoaderFactory("pagedash.loader", (opts) => {
  new PageDash({
    target: opts.target,
    props: {
      env: opts.env,
    },
  });
});

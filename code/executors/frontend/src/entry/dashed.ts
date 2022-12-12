import DashedApp from "./dashed/index.svelte";
import { registerExecLoaderFactory } from "../lib";

// fixme => change to dash.loader

registerExecLoaderFactory("simpledash.main", (opts) => {
  new DashedApp({
    target: document.getElementById("plugroot"), // opts.target,
    props: {
      env: opts.env,
    },
  });
});

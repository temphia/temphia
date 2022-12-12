import DashedApp from "./dashed/index.svelte";
import type { FactoryOptions } from "../lib";
import { registerExecLoaderFactory } from "../lib";

// fixme => change to dash.loader

registerExecLoaderFactory("simpledash.main", (opts: FactoryOptions) => {
  new DashedApp({
    target: document.getElementById("plugroot"), // opts.target,
    props: {
      env: opts.env,
    },
  });
});

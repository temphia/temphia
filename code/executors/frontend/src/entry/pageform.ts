import PageForm from "./pageform/index.svelte";
import { registerExecLoaderFactory } from "../lib";

registerExecLoaderFactory("pageform.loader", (opts) => {
  new PageForm({
    target: opts.target,
    props: {
      env: opts.env,
    },
  });
});

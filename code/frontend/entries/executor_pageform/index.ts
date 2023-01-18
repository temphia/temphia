import { registerExecLoaderFactory } from "../../lib/engine/register";
import PageForm from "./index.svelte";

registerExecLoaderFactory("pageform.loader", (opts) => {
  console.log("@@pagefrom.loader", opts);

  new PageForm({
    target: opts.target,
    props: {
      env: opts.env,
    },
  });
});

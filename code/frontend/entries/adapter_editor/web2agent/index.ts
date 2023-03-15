import type { Registry } from "../../../lib/registry/registry";
import AdapterEditorWeb2Agent from "./index.svelte";

const r = window["__registry__"] as Registry<any>;
r.RegisterFactory("temphia.adapter_editor.loader", `web2agent.main`, (opts) => {
  new AdapterEditorWeb2Agent({
    target: opts.target,
    props: {},
  });
});

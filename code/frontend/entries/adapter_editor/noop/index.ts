import type { Registry } from "../../../lib/registry/registry";
import AdapterEditorNoop from "./index.svelte";

const r = window["__registry__"] as Registry<any>;
r.RegisterFactory("temphia.adapter_editor.loader", `noop.main`, (opts) => {
  new AdapterEditorNoop({
    target: opts.target,
    props: {},
  });
});

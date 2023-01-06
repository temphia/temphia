import type { Registry } from "../../lib/registry/registry";
import AdapterEditorEasypage from "./index.svelte";

const r = window["__registry__"] as Registry<any>;
r.RegisterFactory("temphia.adapter_editor.loader", `easypage.main`, (opts) => {
  new AdapterEditorEasypage({
    target: opts.target,
    props: { env: opts.env },
  });
});

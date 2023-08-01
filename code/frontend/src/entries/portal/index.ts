import { initRegistry } from "../../lib/engine/putils";
import PortalApp from "./app.svelte";

initRegistry();

const __svelte_app__ = new PortalApp({
  target: document.body,
  props: {},
});

export default __svelte_app__;

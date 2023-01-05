import { initRegistry } from "../../lib/engine/putils";
initRegistry();

window.addEventListener("load", (ev) => {
  console.log("@adapter_editor_loader");
});

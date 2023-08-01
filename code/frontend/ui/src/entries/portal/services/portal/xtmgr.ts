import type { Registry } from "../../../../lib/registry/registry";

export class XtMgr {
  app: any;
  registry: Registry<any>;
  bprint_editors: Map<string, any>;

  constructor(app: any, registry: Registry<any>) {
    this.app = app;
    this.registry = registry;
    this.bprint_editors = new Map();

    console.log("XTMGR ", this)
  }

  init = () => {
    const pexts = this.registry.GetAll("portal_extensions");
    pexts.forEach((pext) => {
      pext({
        xmgr: this,
      });
    });
  };

  register_bprint_editor = (name: string, compo: any) => {
    this.bprint_editors.set(name, compo);
  };
}

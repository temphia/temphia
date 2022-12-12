import WizardApp from "./wizard/index.svelte";
import {
  registerExecLoaderFactory,
  FactoryOptions,
} from "../lib";

// fixme => change to wizard.loader

registerExecLoaderFactory("simplewizard.main", (opts: FactoryOptions) => {
  const __simple_wizard_app__ = new WizardApp({
    target: opts.target,
    props: {
      env: opts.env,
      options: opts.payload,
    },
  });
});

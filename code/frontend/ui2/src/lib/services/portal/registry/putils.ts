import { Registry } from "../registry/registry";

export const initRegistry = () => {
  console.log("init registry");
  
  if (window["__registry__"]) {
    console.warn("Registry already loaded, skipping...");
    return;
  }
  const r = new Registry<any>();
  r.RegisterFactory(
    "loader.factory",
    "std.loader",
    async (opts: any) => {
      await opts.registry.WatchLoad("plug.factory", opts.entry, 200000);
      const factory = opts.registry.Get("plug.factory", opts.entry);
      if (!factory) {
        console.warn("could not load plug factory");
        return;
      }
      factory({
        plug: opts.plug,
        agent: opts.agent,
        entry: opts.entry,
        env: opts.env,
        target: opts.target,
        registry: opts.registry,
      });
    }
  );
  
  window["__registry__"] = r;
  window["__register_factory__"] = r.RegisterFactory;
  window["RegisterFactory"] = r.RegisterFactory;
};

// it will find appoprate loader and call loader
// then its loader responsibility to start registered factories
// plugStart => loader => actual_plug_factory_start (using entry_name)
export const plugStart = async (opts: {
  exec_loader?: string;
  plug: string;
  agent: string;
  entry: string;
  env: any; // Environment;
  target: HTMLElement;
}) => {
  console.log("let there be light", opts);

  const registry = window["__registry__"] as Registry<any>;
  if (!registry) {
    console.warn("registry not found");
    return;
  }

  if (!opts.exec_loader) {
    opts.exec_loader = "std.loader";
  }

  try {
    await registry.WatchLoad("loader.factory", opts.exec_loader, 100000);
  } catch (error) {
    console.warn("could not load, error occured:", error);
    return;
  }

  const loaderFactory = registry.Get("loader.factory", opts.exec_loader);
  if (!opts.target) {
    opts.target = document.body;
  }

  loaderFactory({
    plug: opts.plug,
    agent: opts.agent,
    entry: opts.entry,
    env: opts.env,
    registry: registry,
    target: opts.target,
  });
};

export const registerPlugFactory = (entryName: string, factory: (opts: any) => void) => registerFactory("plug.factory", entryName, factory)
export const registerExecLoaderFactory = (name: string, factory: (opts: any) => void) => registerFactory("loader.factory", name, factory)
export const registerFactory = (ftype: string, name: string, factory: (opts: any) => void) => {
    const pf = window["__register_factory__"];
    if (!pf) {
        console.warn("factory registry not found");
        return;
    }
    pf(ftype, name, factory);
}


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


// sync this
export interface Environment {
  PreformAction: (name: string, data: any) => Promise<any>;
  PreformParentAction: (name: string, data: any) => Promise<any>;
  OnParentAction: (handler: (data: any) => {}) => void;

  GetRegistry: () => any;
  GetFolderTktAPI: (ticket: string) => any;
  GetRoomTktAPI: (room: string, ticket?: string) => Promise<any>;
  GetDataTableTktAPI: (ticket: string) => any;
}

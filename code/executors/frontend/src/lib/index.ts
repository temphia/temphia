export type registerExecLoaderFactory = any;

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

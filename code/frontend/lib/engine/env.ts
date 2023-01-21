import type { AdminDataAPI, FolderTktAPI } from "../apiv2";
import type { AdminPlugStateTktAPI } from "../apiv2/admin/plug_state";
import type { Registry } from "../registry/registry";

export interface Environment {
  PreformAction: (name: string, data: any) => Promise<any>;
  PreformParentAction: (name: string, data: any) => Promise<any>;
  OnParentAction: (handler: (data: any) => {}) => void;

  GetRegistry: () => Registry<any>;
  
  GetRoomTktAPI: (ticket: string) => any;

  GetFolderTktAPI: (ticket: string) => FolderTktAPI;
  GetDataTableTktAPI: (ticket: string) => AdminDataAPI;
  GetPlugStateTktAPI: (ticket: string) => AdminPlugStateTktAPI;
}

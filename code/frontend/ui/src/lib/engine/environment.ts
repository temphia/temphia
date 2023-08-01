import type { Registry } from "../registry/registry";

export interface ExecVariables {
  api_base_url: string
  tenant_id: string
  plug_id: string
  agent_id: string
  executor?: string
  launcher?: string
  exec_data?: any
  invoker_name?: string
  target_id?: number
  user?: string
  user_group?: string
}

export interface AssetManager {
  GetAgentAssetURL(name: string): string;
  GetExecutorAssetURL(name: string): string;
  ImportDyanmic(name: string): Promise<any>;
  SheduleWorker(name: string): Worker;
}

export interface Extendable {
  OnEvent(evname: string, handler: (ev: any) => void): void
}

export type ExtendableType = Extendable | any

export interface Environment<S=ExtendableType> {
  PreformAction: (name: string, data: any) => Promise<{data:any,ok:boolean,status:number}>;
  PreformParentAction: (name: string, data: any) => Promise<any>;
  OnParentAction: (handler: (data: any) => {}) => void;

  GetRegistry: () => Registry<any>;
  GetExecVars: () => ExecVariables;
  GetAssetManager: () => AssetManager;
  GetExecApiManager: () => Promise<any>
  Extend(s: S): S
}


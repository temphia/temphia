export interface Environment {
  PreformAction: (name: string, data: any) => Promise<ActionResponse>;
  PreformParentAction: (name: string, data: any) => Promise<any>;
  OnParentAction: (handler: (data: any) => {}) => void;

  GetRegistry: () => any;
  GetFolderTktAPI: (ticket: string) => any;
  GetRoomTktAPI: (room: string, ticket?: string) => Promise<any>;
  GetDtableTktAPI: (ticket: string) => any;
}

export interface ActionResponse {
  status_ok: boolean;
  content_type?: string;
  body: any;
}

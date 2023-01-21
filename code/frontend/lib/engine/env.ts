export interface Environment {
  PreformAction: (name: string, data: any) => Promise<any>;
  PreformParentAction: (name: string, data: any) => Promise<any>;
  OnParentAction: (handler: (data: any) => {}) => void;

  GetRegistry: () => any;
  GetRoomTktAPI: (room: string, ticket?: string) => any;
  GetFolderTktAPI: (ticket: string) => any;
  GetDataTableTktAPI: (ticket: string) => any;
  GetPlugStateTktAPI: (ticket: string) => any;
}

//   PreformAction: (name: string, data: any) => Promise<Result<HttpResponse>>;

export type Result<T = any> = {
  inner: T | Error;
  ok(): boolean;
  err_message(): string;
};

type HttpResponse = {
  data: any;
  headers: { [_: string]: string };
  status: number;
};

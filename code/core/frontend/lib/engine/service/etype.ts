export interface ExecInstance {
  exec_id: string;
  handle(exec_id: string, action: string, data: any): void;
  send(action: string, data: any): void
}

export interface Logger {
  info(message?: any, ...optionalParams: any[]): void;
  log(message?: any, ...optionalParams: any[]): void;
  err(message?: any, ...optionalParams: any[]): void;
}

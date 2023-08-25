import type { Logger } from "../../../../lib/logger";

export class PortalLogger implements Logger {
  constructor() {}

  info(message?: any, ...optionalParams: any[]): void {}
  log(message?: any, ...optionalParams: any[]): void {}
  err(message?: any, ...optionalParams: any[]): void {}
}

declare function _self_list_resource(): [any, string];
declare function _self_get_resource(name: string): [any, string];
declare function _self_inlinks(): [any, string];
declare function _self_outlinks(): [any, string];
declare function _self_new_module(name: string, data: any): [number, string];

declare function _self_module_execute(
  mid: number,
  method: string,
  data: any
): [any, string];

declare function _self_link_execute(
  name: string,
  method: string,
  path: string,
  data: any,
  async: boolean,
  detached: boolean
): [any, string];
declare function _self_fork_execute(method: string, data: any): string;

export const self = {
  list_resource: () => _self_list_resource(),
  get_resource: (name: string) => _self_get_resource(name),
  inlinks: () => _self_inlinks(),
  outlinks: () => _self_outlinks(),
  new_module: (name: string, data: any) => _self_new_module(name, data),
  module_execute: (mid: number, method: string, data: any) =>
    _self_module_execute(mid, method, data),

  link_execute: (
    name: string,
    method: string,
    path: string,
    data: any,
    async: boolean,
    detached: boolean
  ) => _self_link_execute(name, method, path, data, async, detached),
  fork_execute: (method: string, data: any) => _self_fork_execute(method, data),
};

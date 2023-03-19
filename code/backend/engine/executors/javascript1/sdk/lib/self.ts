declare function _self_list_resource(): [any, string];
declare function _self_get_resource(name: string): [any, string];
declare function _self_inlinks(): [any, string];
declare function _self_outlinks(): [any, string];
declare function _self_module_execute(
  name: string,
  method: string,
  path: string,
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
  list_resource: _self_list_resource,
  get_resource: _self_get_resource,
  inlinks: _self_inlinks,
  outlinks: _self_outlinks,
  module_execute: _self_module_execute,
  link_execute: _self_link_execute,
  fork_execute: _self_fork_execute,
};
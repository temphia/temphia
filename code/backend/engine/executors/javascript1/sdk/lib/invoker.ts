declare function _invoker_name(): string;
declare function _invoker_exec_method(
  method: string,
  path: string,
  data: any
): string;
declare function _invoker_context_user(): any;
declare function _invoker_context_user_info(): [any, string];
declare function _invoker_context_user_message(): any;

export const invoker = {
  name: _invoker_name,
  exec_method: _invoker_exec_method,
  context_user: _invoker_context_user,
  context_user_info: _invoker_context_user_info,
  context_user_message: _invoker_context_user_message,
};

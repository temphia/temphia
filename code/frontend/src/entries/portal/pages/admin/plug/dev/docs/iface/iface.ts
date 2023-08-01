// syncme => agentiface.go

export interface AgentIface {
  name: string;
  definations: { [_: string]: any };
  methods: { [_: string]: Method };
  events: { [_: string]: EventType };
  schemas: { [_: string]: ValueType };
}

export interface Method {
  info: string;
  arg: ValueType;
  return_type: ValueType;
  error_types: { [_: string]: string };
}

export interface EventType {
  info: string;
  async: boolean
  arg_data: ValueType;
  return_data: ValueType;
}

export interface ValueType {
  type?: string;
  values?: ValueType[];
  property?: string;
  ref?: string;
}

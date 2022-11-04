export interface Schema {
  name: string;
  fields: Field[];
  required_fields: string[];
  save_action_name: string
}

export interface Field {
  name: string;
  key_name: string;
  disabled: boolean;
  ftype:
    | "TEXT"
    | "MULTI_TEXT"
    | "LONG_TEXT"
    | "TEXT_POLICY"
    | "BOOL"
    | "INT"
    | "KEY_VALUE_TEXT"

    | "USER_GROUP"
    | "USER"
    | "BPRINT"
    | "PLUG"
    | "AGENT"
    | "HANDLER"
    | "RESOURCE";
  linked_fields?: string[];
  options?: string[];
}

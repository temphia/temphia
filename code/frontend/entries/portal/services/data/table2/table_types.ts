export interface FilterItem {
  column: string;
  cond: string;
  value: any;
}

export interface Column {
  slug: string;
  name: string;
  ctype: string;
  options: string[];
  description: string;
  pattern: string;
  strict_pattern: boolean;
  ref_id: string;
  ref_type: string;
  ref_copy: string;
  ref_target: string;
  ref_object: string;
}

export interface DataWidget {
  id: number;
  name: string;
  type: string;
  sub_type: string;
  plug_id: string;
  agent_id: string;
  icon: string;
  payload: string;
}

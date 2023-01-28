// syncme => xbprint/resource.go

export interface NewResource {
  name: string;
  type: string;
  sub_type: string;
  payload: string;
  policy: string;
  meta: { [_: string]: string };
}

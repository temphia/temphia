export type Record = { [_: string]: any };

export type PredicateFunc = (r: Record[]) => boolean | undefined;

export const filterWith =
  (records: Record[]) =>
  (...predicates: PredicateFunc[]) => {
    records.filter((r) => {
      const resp = predicates.forEach((p) => p.apply(r));
    });
  };

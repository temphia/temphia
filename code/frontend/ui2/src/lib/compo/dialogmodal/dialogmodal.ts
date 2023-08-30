export interface DialogModal {
  show_small(comp: any, options: object): void;
  close_small(): void;

  show_big(comp: any, options: object): void;
  close_big(): void;
}

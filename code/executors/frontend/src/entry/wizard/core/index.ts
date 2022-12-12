export interface ModalControl {
  show_big: (_compo: any, _props: object) => void;
  close_big: () => void;
  show_small: (_compo: any, _props: object) => void;
  close_small: () => void;
}

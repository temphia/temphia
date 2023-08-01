import type { DialogModal } from "../../../xcompo/dialogmodal";
import type { PageQueryService } from "./service";

export const KEY = "__pagequery__";

export interface Context {
  get_root(): any;
  get_service(): PageQueryService;
  get_modal(): DialogModal;
}

import { derived, Readable, writable, Writable } from "svelte/store";

import type { SelfAPI } from "../../core/api";
import type { SockdMessage } from "../../core/sockd/stypes";
import type { SockdService } from "../sockd";

export interface Options {
  basicAPI: SelfAPI;
  sockdMuxer: SockdService;
}

interface State {
  messages: object[];
  loading: boolean;
  cursor: number;
}

export class Notification {
  basicAPI: SelfAPI;
  state: Writable<State>;
  isPendingRead: Readable<boolean>;
  is_open: Writable<boolean>;

  constructor(opts: Options) {
    const room = opts.sockdMuxer.get_notification_room();
    room.onServer(this.handler);

    this.basicAPI = opts.basicAPI;
    this.state = writable({
      cursor: 0,
      loading: false,
      messages: [],
    });

    this.isPendingRead = derived([this.state], ([state]) => {
      let pending = false;
      state.messages.forEach((msg) => {
        if (!msg["read"]) {
          pending = true;
          return;
        }
      });
      return pending;
    });

    this.is_open = writable(false);
  }

  handler = (message: SockdMessage) => {
    switch (message.payload["type"]) {
      case "new":
        this.state.update((old) => {
          return {
            ...old,
            messages: [...old.messages, message.payload["data"]],
          };
        });
        break;
      default:
        console.log("@message =>", message);
        break;
    }
  };

  init = async () => {
    this.state.update((old) => ({ ...old, loading: true }));

    const resp = await this.basicAPI.list_messages({});
    if (resp.status !== 200) {
      console.warn("Error happend", resp);
      return;
    }

    this.state.update((old) => {
      return {
        ...old,
        cursor: 0,
        loading: false,
        messages: resp.data,
      };
    });
  };

  set_read_notifications = async (id: number) => {
    await this.basicAPI.modify_messages({
      ops: "read",
      ids: [id],
    });

    return this.init();
  };

  delete_notification = async (id: number) => {
    await this.basicAPI.modify_messages({
      ops: "delete",
      ids: [id],
    });

    return this.init();
  };

  toggle_notification = () => {
    this.is_open.update((old) => !old);
  };
}

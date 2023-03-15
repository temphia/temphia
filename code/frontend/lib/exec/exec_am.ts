import { NewDataTableApi } from "./data";
import { NewFolderApi } from "./folder";
import { NewPlugStateApi } from "./plug_state";
import { NewSockdRoom } from "./sockd";

// ExecAM stands for execution api manager
export class ExecAM {
  api_base_url: string;
  constructor(api_base_url: string) {
    this.api_base_url = api_base_url;
  }

  new_data_api = (token: string) => {
    return NewDataTableApi(this.api_base_url, token);
  };

  new_folder_api = (token: string) => {
    return NewFolderApi(this.api_base_url, token);
  };

  new_sockd_room = async (token: string) => {
    return NewSockdRoom(`${this.api_base_url}/engine/ws?ticket=${token}`);
  };

  new_plug_state = (token: string) => {
    return NewPlugStateApi(this.api_base_url, token);
  };
}

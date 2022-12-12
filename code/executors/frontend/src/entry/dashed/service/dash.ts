import { writable, Writable } from "svelte/store"
import type { Environment } from "../../../lib"

export interface DashOptions {
    env: Environment
}

interface StateData {
    loaded: boolean
    inner?: State
}

interface State {
    name: string
    sections: Section[]
    data: { [_: string]: any }
}

export interface Section {
    name: string,
    layout: string
    panels: Panel[]
}
export interface Panel {
    name: string,
    width: number,
    type: string
    interval: string
    source: string
    options: {[_: string]: any}
}


export class DashClient {
    _env: Environment
    _state: Writable<StateData>

    constructor(opts: DashOptions) {
        this._state = writable({
            loaded: false,
        })

        this._env = opts.env
    }

    init = async () => {
        let resp = await this._env.PreformAction("generate", {})
        if (!resp.status_ok) {
            console.log("Error Initilizing", resp)
            return
        }

        this._state.set({
            loaded: true,
            inner: resp.body
        })
    }

    refresh = () => {

    }

    panel_query = () => {

    }

}




import { writable, Writable, get } from "svelte/store"
import type { BprintAPI } from "../../../../../../lib/core/api"


export class Editor {
    bapi: BprintAPI
    bid: string
    filecache: Map<string, any>
    fileMods: Writable<Object>
    modified: Writable<boolean>
    files: Writable<string[]>
    data: any
    loaded: boolean

    constructor(bid: string, bapi: BprintAPI) {
        this.bid = bid;
        this.bapi = bapi;
        this.filecache = new Map()
        this.fileMods = writable({})
        this.modified = writable(false)
        this.files = writable([])
        this.data = {}
        this.loaded = false
    }

    load = async () => {
        const resp = await this.bapi.bprint_get(this.bid);
        if (resp.status !== 200) {
            console.warn("could not load bprint")
            return
        }

        this.data = resp.data;
        this.files.set(resp.data["files"] || [])
        this.loaded = true;
        this.modified.set(false)
    }

    onlyLoadFiles = async () => {
        const resp = await this.bapi.bprint_get(this.bid);
        if (resp.status !== 200) {
            console.warn("could not load bprint")
            return
        }

        this.files.set(resp.data["files"] || [])
    }

    getField = (key: string) => {
        return this.data[key] || ""
    }
    setModified() {
        this.modified.set(true)
    }


    update = (data: any) => {
        this.data = { ...this.data, ...data }
    }


    save = async () => {
        const resp = await this.bapi.bprint_update(this.bid, { ...this.data });
        if (resp.status !== 200) {
            // fixme =>
            return
        }
        await this.load()
    }

    updateField = (key: string, value: any) => {
        if (this.data[key] === value) {
            return
        }
        this.data[key] = value
        this.modified.set(true)
    }

    // file related

    uploadNewFile = (file: string, data: any) => {
        return this.bapi.bprint_new_file(this.bid, file, data)
    }

    updateFile = (file: string, data: any) => {
        this.filecache.set(file, data)
        this.setModifiedFile(file)
    }

    saveFile = async (filename: string) => {
        const file = this.filecache.get(filename)
        const resp = await this.bapi.bprint_update_file(this.bid, filename, file)
        if (resp.status !== 200) {
            console.warn("could not save file", resp)
            return
        }
        this.fileMods.update((old) => ({ ...old, [filename]: false }))

    }

    getFile = async (filename: string) => {
        const cacheFile = this.filecache.get(filename)
        if (cacheFile) {
            return cacheFile
        }

        const resp = await this.bapi.bprint_get_file(this.bid, filename)
        if (resp.status !== 200) {
            return
        }
        const data = get_string(resp.data)
        this.filecache.set(filename, data)
        this.fileMods.update((old) => ({ ...old, [filename]: false }))
        return data
    }

    checkData = async (data: string, dataType: string) => {

    }

    setModifiedFile = (file: string) => {
        this.fileMods.update((old) => ({ ...old, [file]: true }))
    }


}


const get_string = (data: any) => ((typeof data === "string") ? data : JSON.stringify(data, null, 4))
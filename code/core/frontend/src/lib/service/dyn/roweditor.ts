import type { Writable } from "svelte/store"
import type { DirtyData } from "./data_types"

export type Callback = () => void
export class RowEditor {
    _dirtyStore: Writable<DirtyData>
    _callbacks: Map<string, Callback>
    constructor(store: Writable<DirtyData>) {
        this._dirtyStore = store
        this._callbacks = new Map()
    }

    RegisterBeforeSave(field: string, callback: Callback): void {
        this._callbacks.set(field, callback)
    }

    OnChange(_filed: string, _value: any): void {
        this.setValue(_filed, _value)
    }

    // row stuff
    startModifyRow = (row: number) => {
        this._callbacks.clear()
        this._dirtyStore.set({ rowid: row, data: {} })
    }

    startNewRow = () => {
        this._callbacks.clear()
        this._dirtyStore.set({ rowid: 0, data: {} })
    }


    setValue = (_filed: string, value: any) => {
        this._dirtyStore.update((old) => ({ ...old, data: { ...old.data, [_filed]: value } }))
    }

    clearDirtyRow = () => {
        this._dirtyStore.set({ rowid: 0, data: {} })
    }

    setRefCopy(column: string, value: any) {
        this._dirtyStore.update((old) => ({ ...old, data: { ...old.data, [column]: value } }))
    }

    beforeSave() {
        this._callbacks.forEach((val) => val())
    }
}

import { writable, Writable } from "svelte/store"

export interface State {
    indexed_column: { [_: string]: object }
    column_order: string[]
    reverse_ref_column: object[]

    rows: number[]
    indexed_rows: { [_: number]: object }

    sparse_rows: number[]
    remote_dirty: { [_: number]: true }
    views: object[]
    hooks: object[]
}

const istclass = ["shorttext", "phonenumber", "number", "select"]

export class CommonStore {
    states: Writable<{ [_: string]: State }>

    constructor() {
        this.states = writable({})
        this.states.subscribe((state) => console.log("ALL DATA =========================>", state))
    }

    set_rows_data = (table: string, data: any, append: boolean) => {
        if (!data["rows"]) {
            return
        }

        this.states.update((old) => {
            const oldstate = old[table] || {
                reverse_ref_column: [],
                column_order: [],
                hooks: [],
                indexed_column: {},
                indexed_rows: {},
                remote_dirty: {},
                rows: [],
                sparse_rows: [],
                views: []
            }

            const indexed_column = data["columns"]
            const column_order = this.generate_column_order(indexed_column)


            const old_rows = append ? (oldstate["rows"] || []) : []
            const old_indexed = append ? (oldstate["indexed_rows"] || {}) : {}

            let reverse_ref_column = oldstate["reverse_ref_column"] || []
            let views = oldstate["views"] || []
            let hooks = oldstate["hooks"] || []

            const extra_meta = data["extra_meta"]

            if (extra_meta) {
                reverse_ref_column = extra_meta["reverse_refs"] || []
                views = extra_meta["views"] || []
                hooks = extra_meta["hooks"] || []
            }


            const _raw_rows = data["rows"] //  [{ "__id": 1, xyz: "mno" }]
            const _rows = _raw_rows.map(row => (row["__id"]))

            // fixme => implement aesc and desc
            const rows = Array.from((new Set([..._rows, ...old_rows]))).sort((a, b) => (a - b)) // only works for order_by "__id"


            const indexed_rows = _raw_rows.reduce((result, curr) => {
                result[curr.__id] = curr

                return result
            }, { ...old_indexed })


            return {
                ...old,
                [table]: {
                    ...oldstate,
                    column_order,
                    indexed_column,
                    indexed_rows,
                    reverse_ref_column,
                    rows,
                    hooks,
                    views
                }
            }
        })
    }

    set_row_data = (table: string, data: object) => {
        this.states.update((old) => {
            const state = old[table]

            const row_id = data["__id"]

            let old_row = state.indexed_rows[row_id]
            if (old_row) {
                state.indexed_rows[row_id] = {...old_row, ...data}
                return { ...old, [table]: { ...state, indexed_rows: { ...state.indexed_rows } } }
            }

            return old
        })
    }


    set_sockd_change = (data) => {
        this.states.update((old) => (old))
    }



    generate_column_order = (columns: { [_: string]: object }): string[] => {
        // fixme => also use view column order data ?


        const doneCols = {}
        const orderedColumns = []

        // then first class colums 
        istclass.forEach((cType) => {
            Object.values(columns).forEach((val) => {
                if (doneCols[val["slug"]]) {
                    return
                }

                if (val["ref_type"]) {
                    return
                }


                if (val["ctype"] !== cType) {
                    return
                }

                orderedColumns.push(val["slug"])
                doneCols[val["slug"]] = true
            })
        })

        // then remaining columns expect ref types
        Object.values(columns).forEach((val) => {
            if (istclass.includes(val["ctype"])) {
                return
            }
            if (doneCols[val["slug"]]) {
                return
            }

            if (val["ref_type"]) {
                return
            }


            orderedColumns.push(val["slug"])
            doneCols[val["slug"]] = true
        })

        // atlast ref types
        Object.values(columns).forEach((val) => {
            if (doneCols[val["slug"]]) {
                return
            }

            if (!val["ref_type"]) {
                return
            }

            orderedColumns.push(val["slug"])
            doneCols[val["slug"]] = true
        })

        return orderedColumns
    }


    destroy = () => {

    }
}

/*

{
                    name: "Record Hook 1",
                    type: "data_hook",
                    sub_type: "row",
                },
                {
                    name: "Table Hook 1",
                    type: "data_hook",
                    sub_type: "table",
                    plug_id: "umangadashed",
                    agent_id: "default",
                },

                {
                    name: "Table Hook 1",
                    type: "data_hook",
                    sub_type: "table",
                    plug_id: "umangawizard",
                    agent_id: "default",
                }

*/
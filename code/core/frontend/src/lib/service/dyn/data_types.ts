
export interface FilterItem {
    column: string
    cond: string
    value: any
}

export interface ViewData {
    name?: string
    filter_conds: FilterItem[]
    count: number
    selects: string[]
    main_column: string
    search_term: string
}

export interface NavData {
    loading: boolean
    lastTry: Date

    loading_error: string
    last_page: boolean
    active_page: number

    active_view: ViewData
}

export interface DirtyData {
    rowid: number
    data: object
}

export const defaultViewData = () => ({
    count: 20,
    filter_conds: [],
    main_column: "",
    search_term: "",
    selects: [],
    page: 0
})


export interface Column {
    slug: string
    name: string
    ctype: string
    options: string[]
    description: string
    pattern: string
    strict_pattern: boolean
    ref_id: string
    ref_type: string
    ref_copy: string
    ref_target: string
    ref_object: string
}

export interface Hook {
    id: number
    name: string
    type: string
    sub_type: string
    plug_id: string
    agent_id: string
    icon: string
    payload: string
}
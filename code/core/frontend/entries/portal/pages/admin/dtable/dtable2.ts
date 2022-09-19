// import type { PortalApp, ApiManager } from "../../../../../lib/app/portal"

export class DynAdminAPI {
    app: any
    apm: any
    constructor(app: any) {
        this.app = app
        this.apm = app.get_apm()
    }

    load_sources = async () => {
        const bapi = this.apm.get_basic_api()
        return bapi.list_dgroup_sources()
    }

    load_groups = async (src: string) => {
        const dgapi = await this.apm.get_dyn_api()
        return dgapi.list_group(src)
    }

    edit_group = async (source: string, group: string, data: any) => {
        const dgapi = await this.apm.get_dyn_api()
        return dgapi.edit_group(source, group, data)
    }

    get_group = async (source: string, group: string) => {
        const dgapi = await this.apm.get_dyn_api()
        return dgapi.get_group(source, group)
    }

    delete_dgroup = async (source: string, group: string) => {
        const dgapi = await this.apm.get_dyn_api()
        return dgapi.delete_group(source, group)
    }

    goto_dgroup = (source: string, group: string) => {
        this.app.navigator.goto_admin_dgroup_page(source, group)
    }

    goto_dtable = (source: string, group: string, table: string) => {
        this.app.navigator.goto_admin_dtable_page(source, group, table)

    }


    goto_dtabe_data = (source: string, group: string, table: string) => {
        this.app.navigator.goto_dtable(source, group, table)
    }

    // table
    load_group_tables = async (source: string, group: string) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.list_tables()
    }

    load_tables_column = async (source: string, group: string, table: string) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.list_columns(table)
    }

    get_dtable = async (source: string, group: string, table: string) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.get_table(table)
    }

    edit_dtable = async (source: string, group: string, table: string, data: any) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.edit_table(table, data)
    }

    get_column = async (source: string, group: string, table: string, column: string) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.get_column(table, column)
    }


    edit_column = async (source: string, group: string, table: string, column: string, data: any) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.edit_column(table, column, data)
    }


    // view
    list_view = async (source: string, group: string, table: string) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.list_view(table)
    }

    new_view = async (source: string, group: string, tid: string, data: any) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.new_view(tid, data)
    }

    modify_view = async (source: string, group: string, tid: string, id: number, data: any) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.modify_view(tid, id, data)
    }

    get_view = async (source: string, group: string, tid: string, id: number) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.get_view(tid, id)
    }

    del_view = async (source: string, group: string, tid: string, id: number) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.del_view(tid, id)
    }


    list_hook = async (source: string, group: string, tid: string) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.list_hook(tid)
    }


    new_hook = async (source: string, group: string, tid: string, data: any) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.new_hook(tid, data)
    }


    modify_hook = async (source: string, group: string, tid: string, id: number, data: any) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.modify_hook(tid, id, data)
    }

    get_hook = async (source: string, group: string, tid: string, id: number) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.get_hook(tid, id)
    }

    del_hook = async (source: string, group: string, tid: string, id: number) => {
        const api = await this.apm.get_dtable_api(source, group)
        return api.del_hook(tid, id)
    }

}
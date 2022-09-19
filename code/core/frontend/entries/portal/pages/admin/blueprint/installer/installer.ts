import type { BprintAPI } from "../../../../../../lib/core/api"

// sync_with(vmodels.PlugInstallOptions)
export interface PlugInstallOpts {
    bprint_id: string
    new_plug_id: string
    agents: string[]
    resources: string[]
    schema: string
}


export const installAsPlug = (bapi: BprintAPI) => async (opts: PlugInstallOpts) => {
    return bapi.bprint_install(opts.bprint_id, {
        bprint_id: opts.bprint_id,
        data: {
            agents: opts.agents,
            resources: opts.resources,
            schema: opts.schema,
            new_plug_id: opts.new_plug_id,
        }
    })
}


// sync_with(vmodels.DGroupInstallOptions)
export interface DgroupInstallOpts {
    bprint_id: string

    name: string
    slug: string
    schema: string
    cabinet_source: string
    cabinet_folder: string
    seed_from: string
    dyndb_source: string
}


export const installAsDgroup = (bapi: BprintAPI) => async (opts: DgroupInstallOpts) => {
    return bapi.bprint_install(opts.bprint_id, {
        bprint_id: opts.bprint_id,
        data: {
            group_name: opts.name,
            group_slug: opts.slug,
            schema: opts.schema,
            cabinet_source: opts.cabinet_source,
            cabinet_folder: opts.cabinet_folder,
            seed_from: opts.seed_from,
            dyndb_source: opts.dyndb_source,
        }
    })
}



// export const installAsDtable = async (bprintid: string, schema: string, target: string, name: string, slug: string) => {
//     const bapi = await getBprintAPI()
//     return bapi.bprint_install(bprintid, {
//         bprint_id: bprintid,
//         data: {
//             schema: schema,
//             target_group_id: target,
//             table_slug: slug,
//         }
//     })

// }
import type { Environment, Factory, FactoryOptions } from "../ecore/ecore"

export class Registry {
    _factories: Map<string, (opts: FactoryOptions) => void>
    _watchers: Map<string, (() => void)[]>
    _type_watchers: Map<string, ((factory: Factory) => void)[]>

    constructor() {
        this._factories = new Map()
        this._watchers = new Map()
        this._type_watchers = new Map()
    }

    RegisterFactory = (type: string, name: string, factory: Factory) => {
        console.log(`START REGISTER FACTORY => type(${type}) name(${name})`)

        const key = [type, name].toString()

        this._factories.set(key, factory)

        const watchers = this._watchers.get(key)
        if (watchers) {
            console.log("Found watchers ", watchers)
            watchers.forEach((watcher) => watcher())
        }

        const typeWatchers = this._type_watchers.get(type)
        if (typeWatchers) {
            typeWatchers.forEach((f) => f(factory))
        }

        console.log(`END REGISTER FACTORY => type(${type}) name(${name})`)
    }

    WatchLoad = async (type: string, name: string, timeout: number): Promise<void> => {
        console.log("before Watching")

        const key = [type, name].toString()
        if (this._factories.has(key)) {
            console.log("found factories already")
            return Promise.resolve()
        }

        const p = new Promise<void>((resolve, reject) => {
            console.log("making promise")

            let oldwatcher = this._watchers.get(key)
            if (!oldwatcher) {
                oldwatcher = new Array<() => void>(0)
            }
            oldwatcher.push(() => {
                resolve()
            })
            this._watchers.set(key, oldwatcher)
            setTimeout(() => {
                reject(`TimeOut loading type ${type} & name ${name}`)
            }, timeout)
        })
        return p
    }

    OnTypeLoad = (typename: string, callback: (factory: Factory) => void) => {
        let oldwatcher = this._type_watchers.get(typename)
        if (!oldwatcher) {
            oldwatcher = new Array<() => void>(0)
        }
        oldwatcher.push(callback)
    }

    Get = (type: string, name: string): Factory => {
        const key = [type, name].toString()

        return this._factories.get(key.toString())
    }

    GetAll = (type: string) => {
        const facts: Factory[] = Array(0)

        this._factories.forEach((fact, [_type, _]) => {
            if (type !== _type) {
                return
            }
            facts.push(fact)
        })

        return facts
    }

    InstanceAll = (type: string, opts: FactoryOptions) => {
        this._factories.forEach((fact, key) => {
            const [_type, _] = key.split(',')
            if (type !== _type) {
                return
            }
            fact(opts)
        })
    }

    Instance = (type: string, name: string, opts: FactoryOptions) => {
        const key = [type, name].toString()
        this._factories.get(key)(opts)
    }

}


export const initRegistry = () => {
    if (window["__registry__"]) {
        console.warn("Registry already loaded, skipping...")
        return
    }
    const r = new Registry()
    r.RegisterFactory("loader.factory", "std.loader", async (opts: FactoryOptions) => {
        await opts.registry.WatchLoad("plug.factory", opts.entry, 200000)
        const factory = opts.registry.Get("plug.factory", opts.entry)
        if (!factory) {
            console.warn("could not load plug factory")
            return
        }
        factory({
            plug: opts.plug,
            agent: opts.agent,
            entry: opts.entry,
            env: opts.env,
            target: opts.target,
            payload: opts.payload,
            registry: opts.registry
        })
    })

    console.log("GLOBAL_REGISTRY =>", r)

    window["__registry__"] = r
    window["__register_factory__"] = r.RegisterFactory
}

export interface ExecFactoryOptions {
    exec_loader?: string
    plug: string
    agent: string
    entry: string
    env: Environment
    target: HTMLElement
    payload?: any
}

// it will find appoprate loader and call loader
// then its loader responsibility to start registered factories
// plugStart => loader => actual_plug_factory_start (using entry_name)
export const plugStart = async (opts: ExecFactoryOptions) => {
    console.log("let there be light", opts)

    const registry = window["__registry__"] as Registry
    if (!registry) {
        console.warn("registry not found")
        return
    }

    if (!opts.exec_loader) {
        opts.exec_loader = "std.loader"
    }

    try {
        await registry.WatchLoad("loader.factory", opts.exec_loader, 100000)
    } catch (error) {
        console.warn("could not load, error occured:", error)
        return
    }

    const loaderFactory = registry.Get("loader.factory", opts.exec_loader)
    if (!opts.target) {
        opts.target = document.body
    }

    loaderFactory({
        plug: opts.plug,
        agent: opts.agent,
        entry: opts.entry,
        env: opts.env,
        registry: registry,
        target: opts.target,
        payload: opts.payload
    })
}

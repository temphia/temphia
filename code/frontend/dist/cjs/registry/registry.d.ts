export declare type FactoryOptions<F = any> = F;
export declare type Factory<F = any> = (opts: FactoryOptions<F>) => void;
export declare class Registry<F> {
    _factories: Map<string, (opts: FactoryOptions<F>) => void>;
    _watchers: Map<string, (() => void)[]>;
    _type_watchers: Map<string, ((factory: Factory<F>) => void)[]>;
    constructor();
    RegisterFactory: (type: string, name: string, factory: Factory<F>) => void;
    WatchLoad: (type: string, name: string, timeout: number) => Promise<void>;
    OnTypeLoad: (typename: string, callback: (factory: Factory<F>) => void) => void;
    Get: (type: string, name: string) => Factory<F>;
    GetAll: (type: string) => Factory<F>[];
    InstanceAll: (type: string, opts: FactoryOptions<F>) => void;
    Instance: (type: string, name: string, opts: FactoryOptions<F>) => void;
}

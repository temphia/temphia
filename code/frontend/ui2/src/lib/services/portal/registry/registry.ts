export type FactoryOptions<F = any> = F;
export type Factory<F = any> = (opts: FactoryOptions<F>) => any;

export class Registry<F> {
  _factories: Map<string, (opts: FactoryOptions<F>) => void>;
  _watchers: Map<string, (() => void)[]>;
  _type_watchers: Map<string, ((factory: Factory<F>) => void)[]>;

  constructor() {
    this._factories = new Map();
    this._watchers = new Map();
    this._type_watchers = new Map();
  }

  RegisterFactory = (type: string, name: string, factory: Factory<F>) => {
    console.log(`START REGISTER FACTORY => type(${type}) name(${name})`);

    const key = [type, name].toString();

    this._factories.set(key, factory);

    const watchers = this._watchers.get(key);
    if (watchers) {
      console.log("Found watchers ", watchers);
      watchers.forEach((watcher) => watcher());
    }

    const typeWatchers = this._type_watchers.get(type);
    if (typeWatchers) {
      typeWatchers.forEach((f) => f(factory));
    }

    console.log(`END REGISTER FACTORY => type(${type}) name(${name})`);
  };

  WatchLoad = async (
    type: string,
    name: string,
    timeout: number
  ): Promise<void> => {
    console.log("before Watching");

    const key = [type, name].toString();
    if (this._factories.has(key)) {
      console.log("found factories already");
      return Promise.resolve();
    }

    const p = new Promise<void>((resolve, reject) => {
      console.log("making promise");

      let oldwatcher = this._watchers.get(key);
      if (!oldwatcher) {
        oldwatcher = new Array<() => void>(0);
      }
      oldwatcher.push(() => {
        resolve();
      });
      this._watchers.set(key, oldwatcher);
      setTimeout(() => {
        reject(`TimeOut loading type ${type} & name ${name}`);
      }, timeout);
    });
    return p;
  };

  OnTypeLoad = (typename: string, callback: (factory: Factory<F>) => void) => {
    let oldwatcher = this._type_watchers.get(typename);
    if (!oldwatcher) {
      oldwatcher = new Array<() => void>(0);
    }
    oldwatcher.push(callback);
  };

  Get = (type: string, name: string): Factory<F> => {
    const key = [type, name].toString();

    return this._factories.get(key.toString());
  };

  GetAll = (type: string) => {
    const facts: Factory<F>[] = Array(0);

    this._factories.forEach((fact, [_type, _]) => {
      if (type !== _type) {
        return;
      }
      facts.push(fact);
    });

    return facts;
  };

  InstanceAll = (type: string, opts: FactoryOptions<F>) => {
    this._factories.forEach((fact, key) => {
      const [_type, _] = key.split(",");
      if (type !== _type) {
        return;
      }
      fact(opts);
    });
  };

  Instance = (type: string, name: string, opts: FactoryOptions<F>) => {
    const key = [type, name].toString();
    this._factories.get(key)(opts);
  };
}

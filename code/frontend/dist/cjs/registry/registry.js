"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Registry = void 0;
class Registry {
    constructor() {
        this.RegisterFactory = (type, name, factory) => {
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
        this.WatchLoad = async (type, name, timeout) => {
            console.log("before Watching");
            const key = [type, name].toString();
            if (this._factories.has(key)) {
                console.log("found factories already");
                return Promise.resolve();
            }
            const p = new Promise((resolve, reject) => {
                console.log("making promise");
                let oldwatcher = this._watchers.get(key);
                if (!oldwatcher) {
                    oldwatcher = new Array(0);
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
        this.OnTypeLoad = (typename, callback) => {
            let oldwatcher = this._type_watchers.get(typename);
            if (!oldwatcher) {
                oldwatcher = new Array(0);
            }
            oldwatcher.push(callback);
        };
        this.Get = (type, name) => {
            const key = [type, name].toString();
            return this._factories.get(key.toString());
        };
        this.GetAll = (type) => {
            const facts = Array(0);
            this._factories.forEach((fact, [_type, _]) => {
                if (type !== _type) {
                    return;
                }
                facts.push(fact);
            });
            return facts;
        };
        this.InstanceAll = (type, opts) => {
            this._factories.forEach((fact, key) => {
                const [_type, _] = key.split(",");
                if (type !== _type) {
                    return;
                }
                fact(opts);
            });
        };
        this.Instance = (type, name, opts) => {
            const key = [type, name].toString();
            this._factories.get(key)(opts);
        };
        this._factories = new Map();
        this._watchers = new Map();
        this._type_watchers = new Map();
    }
}
exports.Registry = Registry;

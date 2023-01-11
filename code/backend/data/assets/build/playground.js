var app = (function () {
    'use strict';

    function noop() { }
    function assign(tar, src) {
        // @ts-ignore
        for (const k in src)
            tar[k] = src[k];
        return tar;
    }
    function add_location(element, file, line, column, char) {
        element.__svelte_meta = {
            loc: { file, line, column, char }
        };
    }
    function run(fn) {
        return fn();
    }
    function blank_object() {
        return Object.create(null);
    }
    function run_all(fns) {
        fns.forEach(run);
    }
    function is_function(thing) {
        return typeof thing === 'function';
    }
    function safe_not_equal(a, b) {
        return a != a ? b == b : a !== b || ((a && typeof a === 'object') || typeof a === 'function');
    }
    function is_empty(obj) {
        return Object.keys(obj).length === 0;
    }
    function validate_store(store, name) {
        if (store != null && typeof store.subscribe !== 'function') {
            throw new Error(`'${name}' is not a store with a 'subscribe' method`);
        }
    }
    function subscribe(store, ...callbacks) {
        if (store == null) {
            return noop;
        }
        const unsub = store.subscribe(...callbacks);
        return unsub.unsubscribe ? () => unsub.unsubscribe() : unsub;
    }
    function component_subscribe(component, store, callback) {
        component.$$.on_destroy.push(subscribe(store, callback));
    }
    function create_slot(definition, ctx, $$scope, fn) {
        if (definition) {
            const slot_ctx = get_slot_context(definition, ctx, $$scope, fn);
            return definition[0](slot_ctx);
        }
    }
    function get_slot_context(definition, ctx, $$scope, fn) {
        return definition[1] && fn
            ? assign($$scope.ctx.slice(), definition[1](fn(ctx)))
            : $$scope.ctx;
    }
    function get_slot_changes(definition, $$scope, dirty, fn) {
        if (definition[2] && fn) {
            const lets = definition[2](fn(dirty));
            if ($$scope.dirty === undefined) {
                return lets;
            }
            if (typeof lets === 'object') {
                const merged = [];
                const len = Math.max($$scope.dirty.length, lets.length);
                for (let i = 0; i < len; i += 1) {
                    merged[i] = $$scope.dirty[i] | lets[i];
                }
                return merged;
            }
            return $$scope.dirty | lets;
        }
        return $$scope.dirty;
    }
    function update_slot_base(slot, slot_definition, ctx, $$scope, slot_changes, get_slot_context_fn) {
        if (slot_changes) {
            const slot_context = get_slot_context(slot_definition, ctx, $$scope, get_slot_context_fn);
            slot.p(slot_context, slot_changes);
        }
    }
    function get_all_dirty_from_scope($$scope) {
        if ($$scope.ctx.length > 32) {
            const dirty = [];
            const length = $$scope.ctx.length / 32;
            for (let i = 0; i < length; i++) {
                dirty[i] = -1;
            }
            return dirty;
        }
        return -1;
    }
    function exclude_internal_props(props) {
        const result = {};
        for (const k in props)
            if (k[0] !== '$')
                result[k] = props[k];
        return result;
    }
    function append(target, node) {
        target.appendChild(node);
    }
    function insert(target, node, anchor) {
        target.insertBefore(node, anchor || null);
    }
    function detach(node) {
        node.parentNode.removeChild(node);
    }
    function destroy_each(iterations, detaching) {
        for (let i = 0; i < iterations.length; i += 1) {
            if (iterations[i])
                iterations[i].d(detaching);
        }
    }
    function element(name) {
        return document.createElement(name);
    }
    function svg_element(name) {
        return document.createElementNS('http://www.w3.org/2000/svg', name);
    }
    function text(data) {
        return document.createTextNode(data);
    }
    function space() {
        return text(' ');
    }
    function empty() {
        return text('');
    }
    function listen(node, event, handler, options) {
        node.addEventListener(event, handler, options);
        return () => node.removeEventListener(event, handler, options);
    }
    function stop_propagation(fn) {
        return function (event) {
            event.stopPropagation();
            // @ts-ignore
            return fn.call(this, event);
        };
    }
    function attr(node, attribute, value) {
        if (value == null)
            node.removeAttribute(attribute);
        else if (node.getAttribute(attribute) !== value)
            node.setAttribute(attribute, value);
    }
    function set_custom_element_data(node, prop, value) {
        if (prop in node) {
            node[prop] = typeof node[prop] === 'boolean' && value === '' ? true : value;
        }
        else {
            attr(node, prop, value);
        }
    }
    function children(element) {
        return Array.from(element.childNodes);
    }
    function set_style(node, key, value, important) {
        if (value === null) {
            node.style.removeProperty(key);
        }
        else {
            node.style.setProperty(key, value, important ? 'important' : '');
        }
    }
    function toggle_class(element, name, toggle) {
        element.classList[toggle ? 'add' : 'remove'](name);
    }
    function custom_event(type, detail, { bubbles = false, cancelable = false } = {}) {
        const e = document.createEvent('CustomEvent');
        e.initCustomEvent(type, bubbles, cancelable, detail);
        return e;
    }

    let current_component;
    function set_current_component(component) {
        current_component = component;
    }
    function get_current_component() {
        if (!current_component)
            throw new Error('Function called outside component initialization');
        return current_component;
    }
    function onDestroy(fn) {
        get_current_component().$$.on_destroy.push(fn);
    }
    function createEventDispatcher() {
        const component = get_current_component();
        return (type, detail, { cancelable = false } = {}) => {
            const callbacks = component.$$.callbacks[type];
            if (callbacks) {
                // TODO are there situations where events could be dispatched
                // in a server (non-DOM) environment?
                const event = custom_event(type, detail, { cancelable });
                callbacks.slice().forEach(fn => {
                    fn.call(component, event);
                });
                return !event.defaultPrevented;
            }
            return true;
        };
    }
    // TODO figure out if we still want to support
    // shorthand events, or if we want to implement
    // a real bubbling mechanism
    function bubble(component, event) {
        const callbacks = component.$$.callbacks[event.type];
        if (callbacks) {
            // @ts-ignore
            callbacks.slice().forEach(fn => fn.call(this, event));
        }
    }

    const dirty_components = [];
    const binding_callbacks = [];
    const render_callbacks = [];
    const flush_callbacks = [];
    const resolved_promise = Promise.resolve();
    let update_scheduled = false;
    function schedule_update() {
        if (!update_scheduled) {
            update_scheduled = true;
            resolved_promise.then(flush);
        }
    }
    function tick() {
        schedule_update();
        return resolved_promise;
    }
    function add_render_callback(fn) {
        render_callbacks.push(fn);
    }
    // flush() calls callbacks in this order:
    // 1. All beforeUpdate callbacks, in order: parents before children
    // 2. All bind:this callbacks, in reverse order: children before parents.
    // 3. All afterUpdate callbacks, in order: parents before children. EXCEPT
    //    for afterUpdates called during the initial onMount, which are called in
    //    reverse order: children before parents.
    // Since callbacks might update component values, which could trigger another
    // call to flush(), the following steps guard against this:
    // 1. During beforeUpdate, any updated components will be added to the
    //    dirty_components array and will cause a reentrant call to flush(). Because
    //    the flush index is kept outside the function, the reentrant call will pick
    //    up where the earlier call left off and go through all dirty components. The
    //    current_component value is saved and restored so that the reentrant call will
    //    not interfere with the "parent" flush() call.
    // 2. bind:this callbacks cannot trigger new flush() calls.
    // 3. During afterUpdate, any updated components will NOT have their afterUpdate
    //    callback called a second time; the seen_callbacks set, outside the flush()
    //    function, guarantees this behavior.
    const seen_callbacks = new Set();
    let flushidx = 0; // Do *not* move this inside the flush() function
    function flush() {
        const saved_component = current_component;
        do {
            // first, call beforeUpdate functions
            // and update components
            while (flushidx < dirty_components.length) {
                const component = dirty_components[flushidx];
                flushidx++;
                set_current_component(component);
                update(component.$$);
            }
            set_current_component(null);
            dirty_components.length = 0;
            flushidx = 0;
            while (binding_callbacks.length)
                binding_callbacks.pop()();
            // then, once components are updated, call
            // afterUpdate functions. This may cause
            // subsequent updates...
            for (let i = 0; i < render_callbacks.length; i += 1) {
                const callback = render_callbacks[i];
                if (!seen_callbacks.has(callback)) {
                    // ...so guard against infinite loops
                    seen_callbacks.add(callback);
                    callback();
                }
            }
            render_callbacks.length = 0;
        } while (dirty_components.length);
        while (flush_callbacks.length) {
            flush_callbacks.pop()();
        }
        update_scheduled = false;
        seen_callbacks.clear();
        set_current_component(saved_component);
    }
    function update($$) {
        if ($$.fragment !== null) {
            $$.update();
            run_all($$.before_update);
            const dirty = $$.dirty;
            $$.dirty = [-1];
            $$.fragment && $$.fragment.p($$.ctx, dirty);
            $$.after_update.forEach(add_render_callback);
        }
    }
    const outroing = new Set();
    let outros;
    function group_outros() {
        outros = {
            r: 0,
            c: [],
            p: outros // parent group
        };
    }
    function check_outros() {
        if (!outros.r) {
            run_all(outros.c);
        }
        outros = outros.p;
    }
    function transition_in(block, local) {
        if (block && block.i) {
            outroing.delete(block);
            block.i(local);
        }
    }
    function transition_out(block, local, detach, callback) {
        if (block && block.o) {
            if (outroing.has(block))
                return;
            outroing.add(block);
            outros.c.push(() => {
                outroing.delete(block);
                if (callback) {
                    if (detach)
                        block.d(1);
                    callback();
                }
            });
            block.o(local);
        }
    }

    const globals = (typeof window !== 'undefined'
        ? window
        : typeof globalThis !== 'undefined'
            ? globalThis
            : global);

    function get_spread_update(levels, updates) {
        const update = {};
        const to_null_out = {};
        const accounted_for = { $$scope: 1 };
        let i = levels.length;
        while (i--) {
            const o = levels[i];
            const n = updates[i];
            if (n) {
                for (const key in o) {
                    if (!(key in n))
                        to_null_out[key] = 1;
                }
                for (const key in n) {
                    if (!accounted_for[key]) {
                        update[key] = n[key];
                        accounted_for[key] = 1;
                    }
                }
                levels[i] = n;
            }
            else {
                for (const key in o) {
                    accounted_for[key] = 1;
                }
            }
        }
        for (const key in to_null_out) {
            if (!(key in update))
                update[key] = undefined;
        }
        return update;
    }
    function get_spread_object(spread_props) {
        return typeof spread_props === 'object' && spread_props !== null ? spread_props : {};
    }
    function create_component(block) {
        block && block.c();
    }
    function mount_component(component, target, anchor, customElement) {
        const { fragment, on_mount, on_destroy, after_update } = component.$$;
        fragment && fragment.m(target, anchor);
        if (!customElement) {
            // onMount happens before the initial afterUpdate
            add_render_callback(() => {
                const new_on_destroy = on_mount.map(run).filter(is_function);
                if (on_destroy) {
                    on_destroy.push(...new_on_destroy);
                }
                else {
                    // Edge case - component was destroyed immediately,
                    // most likely as a result of a binding initialising
                    run_all(new_on_destroy);
                }
                component.$$.on_mount = [];
            });
        }
        after_update.forEach(add_render_callback);
    }
    function destroy_component(component, detaching) {
        const $$ = component.$$;
        if ($$.fragment !== null) {
            run_all($$.on_destroy);
            $$.fragment && $$.fragment.d(detaching);
            // TODO null out other refs, including component.$$ (but need to
            // preserve final state?)
            $$.on_destroy = $$.fragment = null;
            $$.ctx = [];
        }
    }
    function make_dirty(component, i) {
        if (component.$$.dirty[0] === -1) {
            dirty_components.push(component);
            schedule_update();
            component.$$.dirty.fill(0);
        }
        component.$$.dirty[(i / 31) | 0] |= (1 << (i % 31));
    }
    function init(component, options, instance, create_fragment, not_equal, props, append_styles, dirty = [-1]) {
        const parent_component = current_component;
        set_current_component(component);
        const $$ = component.$$ = {
            fragment: null,
            ctx: null,
            // state
            props,
            update: noop,
            not_equal,
            bound: blank_object(),
            // lifecycle
            on_mount: [],
            on_destroy: [],
            on_disconnect: [],
            before_update: [],
            after_update: [],
            context: new Map(options.context || (parent_component ? parent_component.$$.context : [])),
            // everything else
            callbacks: blank_object(),
            dirty,
            skip_bound: false,
            root: options.target || parent_component.$$.root
        };
        append_styles && append_styles($$.root);
        let ready = false;
        $$.ctx = instance
            ? instance(component, options.props || {}, (i, ret, ...rest) => {
                const value = rest.length ? rest[0] : ret;
                if ($$.ctx && not_equal($$.ctx[i], $$.ctx[i] = value)) {
                    if (!$$.skip_bound && $$.bound[i])
                        $$.bound[i](value);
                    if (ready)
                        make_dirty(component, i);
                }
                return ret;
            })
            : [];
        $$.update();
        ready = true;
        run_all($$.before_update);
        // `false` as a special case of no DOM component
        $$.fragment = create_fragment ? create_fragment($$.ctx) : false;
        if (options.target) {
            if (options.hydrate) {
                const nodes = children(options.target);
                // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
                $$.fragment && $$.fragment.l(nodes);
                nodes.forEach(detach);
            }
            else {
                // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
                $$.fragment && $$.fragment.c();
            }
            if (options.intro)
                transition_in(component.$$.fragment);
            mount_component(component, options.target, options.anchor, options.customElement);
            flush();
        }
        set_current_component(parent_component);
    }
    /**
     * Base class for Svelte components. Used when dev=false.
     */
    class SvelteComponent {
        $destroy() {
            destroy_component(this, 1);
            this.$destroy = noop;
        }
        $on(type, callback) {
            const callbacks = (this.$$.callbacks[type] || (this.$$.callbacks[type] = []));
            callbacks.push(callback);
            return () => {
                const index = callbacks.indexOf(callback);
                if (index !== -1)
                    callbacks.splice(index, 1);
            };
        }
        $set($$props) {
            if (this.$$set && !is_empty($$props)) {
                this.$$.skip_bound = true;
                this.$$set($$props);
                this.$$.skip_bound = false;
            }
        }
    }

    function dispatch_dev(type, detail) {
        document.dispatchEvent(custom_event(type, Object.assign({ version: '3.48.0' }, detail), { bubbles: true }));
    }
    function append_dev(target, node) {
        dispatch_dev('SvelteDOMInsert', { target, node });
        append(target, node);
    }
    function insert_dev(target, node, anchor) {
        dispatch_dev('SvelteDOMInsert', { target, node, anchor });
        insert(target, node, anchor);
    }
    function detach_dev(node) {
        dispatch_dev('SvelteDOMRemove', { node });
        detach(node);
    }
    function listen_dev(node, event, handler, options, has_prevent_default, has_stop_propagation) {
        const modifiers = options === true ? ['capture'] : options ? Array.from(Object.keys(options)) : [];
        if (has_prevent_default)
            modifiers.push('preventDefault');
        if (has_stop_propagation)
            modifiers.push('stopPropagation');
        dispatch_dev('SvelteDOMAddEventListener', { node, event, handler, modifiers });
        const dispose = listen(node, event, handler, options);
        return () => {
            dispatch_dev('SvelteDOMRemoveEventListener', { node, event, handler, modifiers });
            dispose();
        };
    }
    function attr_dev(node, attribute, value) {
        attr(node, attribute, value);
        if (value == null)
            dispatch_dev('SvelteDOMRemoveAttribute', { node, attribute });
        else
            dispatch_dev('SvelteDOMSetAttribute', { node, attribute, value });
    }
    function set_data_dev(text, data) {
        data = '' + data;
        if (text.wholeText === data)
            return;
        dispatch_dev('SvelteDOMSetData', { node: text, data });
        text.data = data;
    }
    function validate_each_argument(arg) {
        if (typeof arg !== 'string' && !(arg && typeof arg === 'object' && 'length' in arg)) {
            let msg = '{#each} only iterates over array-like objects.';
            if (typeof Symbol === 'function' && arg && Symbol.iterator in arg) {
                msg += ' You can use a spread to convert this iterable into an array.';
            }
            throw new Error(msg);
        }
    }
    function validate_slots(name, slot, keys) {
        for (const slot_key of Object.keys(slot)) {
            if (!~keys.indexOf(slot_key)) {
                console.warn(`<${name}> received an unexpected slot "${slot_key}".`);
            }
        }
    }
    /**
     * Base class for Svelte components with some minor dev-enhancements. Used when dev=true.
     */
    class SvelteComponentDev extends SvelteComponent {
        constructor(options) {
            if (!options || (!options.target && !options.$$inline)) {
                throw new Error("'target' is a required option");
            }
            super();
        }
        $destroy() {
            super.$destroy();
            this.$destroy = () => {
                console.warn('Component was already destroyed'); // eslint-disable-line no-console
            };
        }
        $capture_state() { }
        $inject_state() { }
    }

    /* entries/xcompo/common/_tailwind.svelte generated by Svelte v3.48.0 */

    function create_fragment$N(ctx) {
    	const block = {
    		c: noop,
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: noop,
    		p: noop,
    		i: noop,
    		o: noop,
    		d: noop
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$N.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$N($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Tailwind', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Tailwind> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Tailwind extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$N, create_fragment$N, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Tailwind",
    			options,
    			id: create_fragment$N.name
    		});
    	}
    }

    const subscriber_queue = [];
    /**
     * Creates a `Readable` store that allows reading by subscription.
     * @param value initial value
     * @param {StartStopNotifier}start start and stop notifications for subscriptions
     */
    function readable(value, start) {
        return {
            subscribe: writable(value, start).subscribe
        };
    }
    /**
     * Create a `Writable` store that allows both updating and reading by subscription.
     * @param {*=}value initial value
     * @param {StartStopNotifier=}start start and stop notifications for subscriptions
     */
    function writable(value, start = noop) {
        let stop;
        const subscribers = new Set();
        function set(new_value) {
            if (safe_not_equal(value, new_value)) {
                value = new_value;
                if (stop) { // store is ready
                    const run_queue = !subscriber_queue.length;
                    for (const subscriber of subscribers) {
                        subscriber[1]();
                        subscriber_queue.push(subscriber, value);
                    }
                    if (run_queue) {
                        for (let i = 0; i < subscriber_queue.length; i += 2) {
                            subscriber_queue[i][0](subscriber_queue[i + 1]);
                        }
                        subscriber_queue.length = 0;
                    }
                }
            }
        }
        function update(fn) {
            set(fn(value));
        }
        function subscribe(run, invalidate = noop) {
            const subscriber = [run, invalidate];
            subscribers.add(subscriber);
            if (subscribers.size === 1) {
                stop = start(set) || noop;
            }
            run(value);
            return () => {
                subscribers.delete(subscriber);
                if (subscribers.size === 0) {
                    stop();
                    stop = null;
                }
            };
        }
        return { set, update, subscribe };
    }
    function derived(stores, fn, initial_value) {
        const single = !Array.isArray(stores);
        const stores_array = single
            ? [stores]
            : stores;
        const auto = fn.length < 2;
        return readable(initial_value, (set) => {
            let inited = false;
            const values = [];
            let pending = 0;
            let cleanup = noop;
            const sync = () => {
                if (pending) {
                    return;
                }
                cleanup();
                const result = fn(single ? values[0] : values, set);
                if (auto) {
                    set(result);
                }
                else {
                    cleanup = is_function(result) ? result : noop;
                }
            };
            const unsubscribers = stores_array.map((store, i) => subscribe(store, (value) => {
                values[i] = value;
                pending &= ~(1 << i);
                if (inited) {
                    sync();
                }
            }, () => {
                pending |= (1 << i);
            }));
            inited = true;
            sync();
            return function stop() {
                run_all(unsubscribers);
                cleanup();
            };
        });
    }

    var commonjsGlobal = typeof globalThis !== 'undefined' ? globalThis : typeof window !== 'undefined' ? window : typeof global !== 'undefined' ? global : typeof self !== 'undefined' ? self : {};

    function createCommonjsModule(fn, basedir, module) {
    	return module = {
    		path: basedir,
    		exports: {},
    		require: function (path, base) {
    			return commonjsRequire(path, (base === undefined || base === null) ? module.path : base);
    		}
    	}, fn(module, module.exports), module.exports;
    }

    function commonjsRequire () {
    	throw new Error('Dynamic requires are not currently supported by @rollup/plugin-commonjs');
    }

    var urlPattern = createCommonjsModule(function (module, exports) {
    // Generated by CoffeeScript 1.10.0
    var slice = [].slice;

    (function(root, factory) {
      if (exports !== null) {
        return module.exports = factory();
      } else {
        return root.UrlPattern = factory();
      }
    })(commonjsGlobal, function() {
      var P, UrlPattern, astNodeContainsSegmentsForProvidedParams, astNodeToNames, astNodeToRegexString, baseAstNodeToRegexString, concatMap, defaultOptions, escapeForRegex, getParam, keysAndValuesToObject, newParser, regexGroupCount, stringConcatMap, stringify;
      escapeForRegex = function(string) {
        return string.replace(/[-\/\\^$*+?.()|[\]{}]/g, '\\$&');
      };
      concatMap = function(array, f) {
        var i, length, results;
        results = [];
        i = -1;
        length = array.length;
        while (++i < length) {
          results = results.concat(f(array[i]));
        }
        return results;
      };
      stringConcatMap = function(array, f) {
        var i, length, result;
        result = '';
        i = -1;
        length = array.length;
        while (++i < length) {
          result += f(array[i]);
        }
        return result;
      };
      regexGroupCount = function(regex) {
        return (new RegExp(regex.toString() + '|')).exec('').length - 1;
      };
      keysAndValuesToObject = function(keys, values) {
        var i, key, length, object, value;
        object = {};
        i = -1;
        length = keys.length;
        while (++i < length) {
          key = keys[i];
          value = values[i];
          if (value == null) {
            continue;
          }
          if (object[key] != null) {
            if (!Array.isArray(object[key])) {
              object[key] = [object[key]];
            }
            object[key].push(value);
          } else {
            object[key] = value;
          }
        }
        return object;
      };
      P = {};
      P.Result = function(value, rest) {
        this.value = value;
        this.rest = rest;
      };
      P.Tagged = function(tag, value) {
        this.tag = tag;
        this.value = value;
      };
      P.tag = function(tag, parser) {
        return function(input) {
          var result, tagged;
          result = parser(input);
          if (result == null) {
            return;
          }
          tagged = new P.Tagged(tag, result.value);
          return new P.Result(tagged, result.rest);
        };
      };
      P.regex = function(regex) {
        return function(input) {
          var matches, result;
          matches = regex.exec(input);
          if (matches == null) {
            return;
          }
          result = matches[0];
          return new P.Result(result, input.slice(result.length));
        };
      };
      P.sequence = function() {
        var parsers;
        parsers = 1 <= arguments.length ? slice.call(arguments, 0) : [];
        return function(input) {
          var i, length, parser, rest, result, values;
          i = -1;
          length = parsers.length;
          values = [];
          rest = input;
          while (++i < length) {
            parser = parsers[i];
            result = parser(rest);
            if (result == null) {
              return;
            }
            values.push(result.value);
            rest = result.rest;
          }
          return new P.Result(values, rest);
        };
      };
      P.pick = function() {
        var indexes, parsers;
        indexes = arguments[0], parsers = 2 <= arguments.length ? slice.call(arguments, 1) : [];
        return function(input) {
          var array, result;
          result = P.sequence.apply(P, parsers)(input);
          if (result == null) {
            return;
          }
          array = result.value;
          result.value = array[indexes];
          return result;
        };
      };
      P.string = function(string) {
        var length;
        length = string.length;
        return function(input) {
          if (input.slice(0, length) === string) {
            return new P.Result(string, input.slice(length));
          }
        };
      };
      P.lazy = function(fn) {
        var cached;
        cached = null;
        return function(input) {
          if (cached == null) {
            cached = fn();
          }
          return cached(input);
        };
      };
      P.baseMany = function(parser, end, stringResult, atLeastOneResultRequired, input) {
        var endResult, parserResult, rest, results;
        rest = input;
        results = stringResult ? '' : [];
        while (true) {
          if (end != null) {
            endResult = end(rest);
            if (endResult != null) {
              break;
            }
          }
          parserResult = parser(rest);
          if (parserResult == null) {
            break;
          }
          if (stringResult) {
            results += parserResult.value;
          } else {
            results.push(parserResult.value);
          }
          rest = parserResult.rest;
        }
        if (atLeastOneResultRequired && results.length === 0) {
          return;
        }
        return new P.Result(results, rest);
      };
      P.many1 = function(parser) {
        return function(input) {
          return P.baseMany(parser, null, false, true, input);
        };
      };
      P.concatMany1Till = function(parser, end) {
        return function(input) {
          return P.baseMany(parser, end, true, true, input);
        };
      };
      P.firstChoice = function() {
        var parsers;
        parsers = 1 <= arguments.length ? slice.call(arguments, 0) : [];
        return function(input) {
          var i, length, parser, result;
          i = -1;
          length = parsers.length;
          while (++i < length) {
            parser = parsers[i];
            result = parser(input);
            if (result != null) {
              return result;
            }
          }
        };
      };
      newParser = function(options) {
        var U;
        U = {};
        U.wildcard = P.tag('wildcard', P.string(options.wildcardChar));
        U.optional = P.tag('optional', P.pick(1, P.string(options.optionalSegmentStartChar), P.lazy(function() {
          return U.pattern;
        }), P.string(options.optionalSegmentEndChar)));
        U.name = P.regex(new RegExp("^[" + options.segmentNameCharset + "]+"));
        U.named = P.tag('named', P.pick(1, P.string(options.segmentNameStartChar), P.lazy(function() {
          return U.name;
        })));
        U.escapedChar = P.pick(1, P.string(options.escapeChar), P.regex(/^./));
        U["static"] = P.tag('static', P.concatMany1Till(P.firstChoice(P.lazy(function() {
          return U.escapedChar;
        }), P.regex(/^./)), P.firstChoice(P.string(options.segmentNameStartChar), P.string(options.optionalSegmentStartChar), P.string(options.optionalSegmentEndChar), U.wildcard)));
        U.token = P.lazy(function() {
          return P.firstChoice(U.wildcard, U.optional, U.named, U["static"]);
        });
        U.pattern = P.many1(P.lazy(function() {
          return U.token;
        }));
        return U;
      };
      defaultOptions = {
        escapeChar: '\\',
        segmentNameStartChar: ':',
        segmentValueCharset: 'a-zA-Z0-9-_~ %',
        segmentNameCharset: 'a-zA-Z0-9',
        optionalSegmentStartChar: '(',
        optionalSegmentEndChar: ')',
        wildcardChar: '*'
      };
      baseAstNodeToRegexString = function(astNode, segmentValueCharset) {
        if (Array.isArray(astNode)) {
          return stringConcatMap(astNode, function(node) {
            return baseAstNodeToRegexString(node, segmentValueCharset);
          });
        }
        switch (astNode.tag) {
          case 'wildcard':
            return '(.*?)';
          case 'named':
            return "([" + segmentValueCharset + "]+)";
          case 'static':
            return escapeForRegex(astNode.value);
          case 'optional':
            return '(?:' + baseAstNodeToRegexString(astNode.value, segmentValueCharset) + ')?';
        }
      };
      astNodeToRegexString = function(astNode, segmentValueCharset) {
        if (segmentValueCharset == null) {
          segmentValueCharset = defaultOptions.segmentValueCharset;
        }
        return '^' + baseAstNodeToRegexString(astNode, segmentValueCharset) + '$';
      };
      astNodeToNames = function(astNode) {
        if (Array.isArray(astNode)) {
          return concatMap(astNode, astNodeToNames);
        }
        switch (astNode.tag) {
          case 'wildcard':
            return ['_'];
          case 'named':
            return [astNode.value];
          case 'static':
            return [];
          case 'optional':
            return astNodeToNames(astNode.value);
        }
      };
      getParam = function(params, key, nextIndexes, sideEffects) {
        var index, maxIndex, result, value;
        if (sideEffects == null) {
          sideEffects = false;
        }
        value = params[key];
        if (value == null) {
          if (sideEffects) {
            throw new Error("no values provided for key `" + key + "`");
          } else {
            return;
          }
        }
        index = nextIndexes[key] || 0;
        maxIndex = Array.isArray(value) ? value.length - 1 : 0;
        if (index > maxIndex) {
          if (sideEffects) {
            throw new Error("too few values provided for key `" + key + "`");
          } else {
            return;
          }
        }
        result = Array.isArray(value) ? value[index] : value;
        if (sideEffects) {
          nextIndexes[key] = index + 1;
        }
        return result;
      };
      astNodeContainsSegmentsForProvidedParams = function(astNode, params, nextIndexes) {
        var i, length;
        if (Array.isArray(astNode)) {
          i = -1;
          length = astNode.length;
          while (++i < length) {
            if (astNodeContainsSegmentsForProvidedParams(astNode[i], params, nextIndexes)) {
              return true;
            }
          }
          return false;
        }
        switch (astNode.tag) {
          case 'wildcard':
            return getParam(params, '_', nextIndexes, false) != null;
          case 'named':
            return getParam(params, astNode.value, nextIndexes, false) != null;
          case 'static':
            return false;
          case 'optional':
            return astNodeContainsSegmentsForProvidedParams(astNode.value, params, nextIndexes);
        }
      };
      stringify = function(astNode, params, nextIndexes) {
        if (Array.isArray(astNode)) {
          return stringConcatMap(astNode, function(node) {
            return stringify(node, params, nextIndexes);
          });
        }
        switch (astNode.tag) {
          case 'wildcard':
            return getParam(params, '_', nextIndexes, true);
          case 'named':
            return getParam(params, astNode.value, nextIndexes, true);
          case 'static':
            return astNode.value;
          case 'optional':
            if (astNodeContainsSegmentsForProvidedParams(astNode.value, params, nextIndexes)) {
              return stringify(astNode.value, params, nextIndexes);
            } else {
              return '';
            }
        }
      };
      UrlPattern = function(arg1, arg2) {
        var groupCount, options, parsed, parser, withoutWhitespace;
        if (arg1 instanceof UrlPattern) {
          this.isRegex = arg1.isRegex;
          this.regex = arg1.regex;
          this.ast = arg1.ast;
          this.names = arg1.names;
          return;
        }
        this.isRegex = arg1 instanceof RegExp;
        if (!(('string' === typeof arg1) || this.isRegex)) {
          throw new TypeError('argument must be a regex or a string');
        }
        if (this.isRegex) {
          this.regex = arg1;
          if (arg2 != null) {
            if (!Array.isArray(arg2)) {
              throw new Error('if first argument is a regex the second argument may be an array of group names but you provided something else');
            }
            groupCount = regexGroupCount(this.regex);
            if (arg2.length !== groupCount) {
              throw new Error("regex contains " + groupCount + " groups but array of group names contains " + arg2.length);
            }
            this.names = arg2;
          }
          return;
        }
        if (arg1 === '') {
          throw new Error('argument must not be the empty string');
        }
        withoutWhitespace = arg1.replace(/\s+/g, '');
        if (withoutWhitespace !== arg1) {
          throw new Error('argument must not contain whitespace');
        }
        options = {
          escapeChar: (arg2 != null ? arg2.escapeChar : void 0) || defaultOptions.escapeChar,
          segmentNameStartChar: (arg2 != null ? arg2.segmentNameStartChar : void 0) || defaultOptions.segmentNameStartChar,
          segmentNameCharset: (arg2 != null ? arg2.segmentNameCharset : void 0) || defaultOptions.segmentNameCharset,
          segmentValueCharset: (arg2 != null ? arg2.segmentValueCharset : void 0) || defaultOptions.segmentValueCharset,
          optionalSegmentStartChar: (arg2 != null ? arg2.optionalSegmentStartChar : void 0) || defaultOptions.optionalSegmentStartChar,
          optionalSegmentEndChar: (arg2 != null ? arg2.optionalSegmentEndChar : void 0) || defaultOptions.optionalSegmentEndChar,
          wildcardChar: (arg2 != null ? arg2.wildcardChar : void 0) || defaultOptions.wildcardChar
        };
        parser = newParser(options);
        parsed = parser.pattern(arg1);
        if (parsed == null) {
          throw new Error("couldn't parse pattern");
        }
        if (parsed.rest !== '') {
          throw new Error("could only partially parse pattern");
        }
        this.ast = parsed.value;
        this.regex = new RegExp(astNodeToRegexString(this.ast, options.segmentValueCharset));
        this.names = astNodeToNames(this.ast);
      };
      UrlPattern.prototype.match = function(url) {
        var groups, match;
        match = this.regex.exec(url);
        if (match == null) {
          return null;
        }
        groups = match.slice(1);
        if (this.names) {
          return keysAndValuesToObject(this.names, groups);
        } else {
          return groups;
        }
      };
      UrlPattern.prototype.stringify = function(params) {
        if (params == null) {
          params = {};
        }
        if (this.isRegex) {
          throw new Error("can't stringify patterns generated from a regex");
        }
        if (params !== Object(params)) {
          throw new Error("argument must be an object or undefined");
        }
        return stringify(this.ast, params, {});
      };
      UrlPattern.escapeForRegex = escapeForRegex;
      UrlPattern.concatMap = concatMap;
      UrlPattern.stringConcatMap = stringConcatMap;
      UrlPattern.regexGroupCount = regexGroupCount;
      UrlPattern.keysAndValuesToObject = keysAndValuesToObject;
      UrlPattern.P = P;
      UrlPattern.newParser = newParser;
      UrlPattern.defaultOptions = defaultOptions;
      UrlPattern.astNodeToRegexString = astNodeToRegexString;
      UrlPattern.astNodeToNames = astNodeToNames;
      UrlPattern.getParam = getParam;
      UrlPattern.astNodeContainsSegmentsForProvidedParams = astNodeContainsSegmentsForProvidedParams;
      UrlPattern.stringify = stringify;
      return UrlPattern;
    });
    });

    function defineProp (obj, prop, value) {
      Object.defineProperty(obj, prop, { value });
    }

    // Parse schema into routes
    function parse$2 (schema = {}, notRoot, pathname, href = '#') {
      // Convert schema to options object. Schema can be:
      // + function: Svelte component
      // + string: redirect path
      // + object: options
      if (notRoot) {
        let type = typeof schema;
        schema = type === 'function' ? { $$component: schema }
          : type === 'string' ? { $$redirect: schema }
          : (type !== 'object' || schema === null) ? {} : schema;

        let c = schema.$$component;
        if (typeof c !== 'function' && c !== undefined && c !== null)
          throw new Error('Invalid Svelte component')
      }

      // Any properties not starting with $$ will be treated as routes,
      // the rest will be kept as route data. Custom data is also kept,
      // but will be replaced with internal data if duplicating names.
      let route = {};
      for (let i in schema) {
        if (/^\$\$/.test(i))
          defineProp(route, i, schema[i]);
        else
          route[i] = parse$2(schema[i], true, i, href + i);
      }

      // Define internal data
      if (notRoot) {
        defineProp(route, '$$href', href); // full path including #
        defineProp(route, '$$pathname', pathname); // scoped path
        defineProp(route, '$$pattern', new urlPattern(href));
        defineProp(route, '$$stringify', v => route.$$pattern.stringify(v));
      }

      return Object.freeze(route)
    }

    // Routes store must be set before creating any Svelte components.
    // It can only be set once. A parsed version is created after with
    // helpful internal data
    let schema = writable();
    let routes = derived(schema, $ => parse$2($));
    routes.set = v => {
      schema.set(v);
      delete routes.set;
    };

    let regex = /(#?[^?]*)?(\?.*)?/;

    function parse$1 () {
      let match = regex.exec(window.location.hash);
      let pathname = match[1] || '#/';
      let querystring = match[2];
      return { pathname, querystring }
    }

    let path = readable(parse$1(), set => {
      let update = () => set(parse$1());
      window.addEventListener('hashchange', update);
      return () => window.removeEventListener('hashchange', update)
    });

    let pathname = derived(path, $ => $.pathname); // current pathname without query
    let querystring = derived(path, $ => $.querystring);
    derived(querystring, $ => {
      return Array.from(new URLSearchParams($))
        .reduce((a, [i, e]) => { a[i] = e; return a }, {})
    });

    // Search for matching route
    function parse (active, pathname, notRoot, matches = []) {
      if (notRoot) {
        let params = active.$$pattern.match(pathname);
        if (params) {
          return !active.$$redirect
            ? { active, params, matches }
            // redirect
            : tick().then(() => {
              history.replaceState(null, null, '#' + active.$$redirect);
              window.dispatchEvent(new Event('hashchange'));
            })
        }
      }

      for (let e of Object.values(active)) {
        let result = parse(e, pathname, true, [...matches, e]);
        if (result) return result
      }
    }

    let match = derived([routes, pathname], ([$r, $p]) => parse($r, $p) || {});
    derived(match, $ => $.active || {}); // current active route
    derived(match, $ => $.params || {});
    let matches = derived(match, $ => $.matches || []); // parents of active route and itself
    let components = derived(matches, $ => $.map(e => e.$$component).filter(e => e));// components to use in <Router/>

    /* node_modules/svelte-hash-router/src/components/Router.svelte generated by Svelte v3.48.0 */

    function create_fragment$M(ctx) {
    	let switch_instance;
    	let switch_instance_anchor;
    	let current;
    	const switch_instance_spread_levels = [/*$$props*/ ctx[2]];
    	var switch_value = /*$components*/ ctx[0][/*i*/ ctx[1]];

    	function switch_props(ctx) {
    		let switch_instance_props = {};

    		for (let i = 0; i < switch_instance_spread_levels.length; i += 1) {
    			switch_instance_props = assign(switch_instance_props, switch_instance_spread_levels[i]);
    		}

    		return {
    			props: switch_instance_props,
    			$$inline: true
    		};
    	}

    	if (switch_value) {
    		switch_instance = new switch_value(switch_props());
    	}

    	const block = {
    		c: function create() {
    			if (switch_instance) create_component(switch_instance.$$.fragment);
    			switch_instance_anchor = empty();
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			if (switch_instance) {
    				mount_component(switch_instance, target, anchor);
    			}

    			insert_dev(target, switch_instance_anchor, anchor);
    			current = true;
    		},
    		p: function update(ctx, [dirty]) {
    			const switch_instance_changes = (dirty & /*$$props*/ 4)
    			? get_spread_update(switch_instance_spread_levels, [get_spread_object(/*$$props*/ ctx[2])])
    			: {};

    			if (switch_value !== (switch_value = /*$components*/ ctx[0][/*i*/ ctx[1]])) {
    				if (switch_instance) {
    					group_outros();
    					const old_component = switch_instance;

    					transition_out(old_component.$$.fragment, 1, 0, () => {
    						destroy_component(old_component, 1);
    					});

    					check_outros();
    				}

    				if (switch_value) {
    					switch_instance = new switch_value(switch_props());
    					create_component(switch_instance.$$.fragment);
    					transition_in(switch_instance.$$.fragment, 1);
    					mount_component(switch_instance, switch_instance_anchor.parentNode, switch_instance_anchor);
    				} else {
    					switch_instance = null;
    				}
    			} else if (switch_value) {
    				switch_instance.$set(switch_instance_changes);
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			if (switch_instance) transition_in(switch_instance.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			if (switch_instance) transition_out(switch_instance.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(switch_instance_anchor);
    			if (switch_instance) destroy_component(switch_instance, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$M.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    let level = 0;

    function instance$M($$self, $$props, $$invalidate) {
    	let $components;
    	validate_store(components, 'components');
    	component_subscribe($$self, components, $$value => $$invalidate(0, $components = $$value));
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Router', slots, []);
    	let i = level++;
    	onDestroy(() => level--);

    	$$self.$$set = $$new_props => {
    		$$invalidate(2, $$props = assign(assign({}, $$props), exclude_internal_props($$new_props)));
    	};

    	$$self.$capture_state = () => ({
    		level,
    		onDestroy,
    		components,
    		i,
    		$components
    	});

    	$$self.$inject_state = $$new_props => {
    		$$invalidate(2, $$props = assign(assign({}, $$props), $$new_props));
    		if ('i' in $$props) $$invalidate(1, i = $$new_props.i);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	$$props = exclude_internal_props($$props);
    	return [$components, i, $$props];
    }

    class Router extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$M, create_fragment$M, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Router",
    			options,
    			id: create_fragment$M.name
    		});
    	}
    }

    /* entries/portal/pages/launcher/start.svelte generated by Svelte v3.48.0 */

    const file$I = "entries/portal/pages/launcher/start.svelte";

    function create_fragment$L(ctx) {
    	let div;
    	let t1;
    	let a0;
    	let t3;
    	let a1;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Start";
    			t1 = space();
    			a0 = element("a");
    			a0.textContent = "Hello";
    			t3 = space();
    			a1 = element("a");
    			a1.textContent = "Admin Bprint";
    			add_location(div, file$I, 0, 0, 0);
    			attr_dev(a0, "href", "#/launch/hello1");
    			add_location(a0, file$I, 2, 0, 18);
    			attr_dev(a1, "href", "#/admin/bprint");
    			add_location(a1, file$I, 3, 0, 54);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			insert_dev(target, t1, anchor);
    			insert_dev(target, a0, anchor);
    			insert_dev(target, t3, anchor);
    			insert_dev(target, a1, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			if (detaching) detach_dev(t1);
    			if (detaching) detach_dev(a0);
    			if (detaching) detach_dev(t3);
    			if (detaching) detach_dev(a1);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$L.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$L($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Start', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Start> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Start extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$L, create_fragment$L, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Start",
    			options,
    			id: create_fragment$L.name
    		});
    	}
    }

    /* entries/portal/pages/launcher/player.svelte generated by Svelte v3.48.0 */

    const file$H = "entries/portal/pages/launcher/player.svelte";

    function create_fragment$K(ctx) {
    	let div;
    	let t0;
    	let t1;

    	const block = {
    		c: function create() {
    			div = element("div");
    			t0 = text("Player ");
    			t1 = text(/*app*/ ctx[0]);
    			add_location(div, file$H, 4, 0, 38);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, t0);
    			append_dev(div, t1);
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*app*/ 1) set_data_dev(t1, /*app*/ ctx[0]);
    		},
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$K.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$K($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Player', slots, []);
    	let { app } = $$props;
    	const writable_props = ['app'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Player> was created with unknown prop '${key}'`);
    	});

    	$$self.$$set = $$props => {
    		if ('app' in $$props) $$invalidate(0, app = $$props.app);
    	};

    	$$self.$capture_state = () => ({ app });

    	$$self.$inject_state = $$props => {
    		if ('app' in $$props) $$invalidate(0, app = $$props.app);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [app];
    }

    class Player extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$K, create_fragment$K, safe_not_equal, { app: 0 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Player",
    			options,
    			id: create_fragment$K.name
    		});

    		const { ctx } = this.$$;
    		const props = options.props || {};

    		if (/*app*/ ctx[0] === undefined && !('app' in props)) {
    			console.warn("<Player> was created without expected prop 'app'");
    		}
    	}

    	get app() {
    		throw new Error("<Player>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set app(value) {
    		throw new Error("<Player>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/portal/pages/data/loader.svelte generated by Svelte v3.48.0 */

    const file$G = "entries/portal/pages/data/loader.svelte";

    function create_fragment$J(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Loader";
    			add_location(div, file$G, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$J.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$J($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Loader', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Loader> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Loader$3 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$J, create_fragment$J, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Loader",
    			options,
    			id: create_fragment$J.name
    		});
    	}
    }

    /* entries/portal/pages/data/groups.svelte generated by Svelte v3.48.0 */

    const file$F = "entries/portal/pages/data/groups.svelte";

    function create_fragment$I(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Groups";
    			add_location(div, file$F, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$I.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$I($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Groups', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Groups> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Groups$1 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$I, create_fragment$I, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Groups",
    			options,
    			id: create_fragment$I.name
    		});
    	}
    }

    /* entries/portal/pages/data/group.svelte generated by Svelte v3.48.0 */

    const file$E = "entries/portal/pages/data/group.svelte";

    function create_fragment$H(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Group";
    			add_location(div, file$E, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$H.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$H($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Group', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Group> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Group$2 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$H, create_fragment$H, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Group",
    			options,
    			id: create_fragment$H.name
    		});
    	}
    }

    /* entries/portal/pages/data/table.svelte generated by Svelte v3.48.0 */

    const file$D = "entries/portal/pages/data/table.svelte";

    function create_fragment$G(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Table";
    			add_location(div, file$D, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$G.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$G($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Table', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Table> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Table$1 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$G, create_fragment$G, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Table",
    			options,
    			id: create_fragment$G.name
    		});
    	}
    }

    /* entries/portal/pages/cabinet/loader.svelte generated by Svelte v3.48.0 */

    const file$C = "entries/portal/pages/cabinet/loader.svelte";

    function create_fragment$F(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Loader";
    			add_location(div, file$C, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$F.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$F($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Loader', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Loader> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Loader$2 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$F, create_fragment$F, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Loader",
    			options,
    			id: create_fragment$F.name
    		});
    	}
    }

    /* entries/portal/pages/cabinet/folders.svelte generated by Svelte v3.48.0 */

    const file$B = "entries/portal/pages/cabinet/folders.svelte";

    function create_fragment$E(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Folders";
    			add_location(div, file$B, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$E.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$E($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Folders', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Folders> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Folders extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$E, create_fragment$E, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Folders",
    			options,
    			id: create_fragment$E.name
    		});
    	}
    }

    /* entries/portal/pages/cabinet/folder.svelte generated by Svelte v3.48.0 */

    const file$A = "entries/portal/pages/cabinet/folder.svelte";

    function create_fragment$D(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Folder";
    			add_location(div, file$A, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$D.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$D($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Folder', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Folder> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Folder$1 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$D, create_fragment$D, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Folder",
    			options,
    			id: create_fragment$D.name
    		});
    	}
    }

    /* entries/portal/pages/repo/loader.svelte generated by Svelte v3.48.0 */

    const file$z = "entries/portal/pages/repo/loader.svelte";

    function create_fragment$C(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Loader";
    			add_location(div, file$z, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$C.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$C($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Loader', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Loader> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Loader$1 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$C, create_fragment$C, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Loader",
    			options,
    			id: create_fragment$C.name
    		});
    	}
    }

    /* entries/portal/pages/repo/repo.svelte generated by Svelte v3.48.0 */

    const file$y = "entries/portal/pages/repo/repo.svelte";

    function create_fragment$B(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Repo";
    			add_location(div, file$y, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$B.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$B($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Repo', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Repo> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Repo$1 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$B, create_fragment$B, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Repo",
    			options,
    			id: create_fragment$B.name
    		});
    	}
    }

    /* entries/portal/pages/admin/bprint/bprint.svelte generated by Svelte v3.48.0 */

    const file$x = "entries/portal/pages/admin/bprint/bprint.svelte";

    function create_fragment$A(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Bprint";
    			add_location(div, file$x, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$A.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$A($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Bprint', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Bprint> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Bprint extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$A, create_fragment$A, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Bprint",
    			options,
    			id: create_fragment$A.name
    		});
    	}
    }

    /* entries/portal/pages/admin/bprint/bprints.svelte generated by Svelte v3.48.0 */

    const file$w = "entries/portal/pages/admin/bprint/bprints.svelte";

    function create_fragment$z(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Bprints";
    			add_location(div, file$w, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$z.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$z($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Bprints', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Bprints> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Bprints extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$z, create_fragment$z, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Bprints",
    			options,
    			id: create_fragment$z.name
    		});
    	}
    }

    /* entries/portal/pages/admin/bprint/editor/editor.svelte generated by Svelte v3.48.0 */

    const file$v = "entries/portal/pages/admin/bprint/editor/editor.svelte";

    function create_fragment$y(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Editor";
    			add_location(div, file$v, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$y.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$y($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Editor', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Editor> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Editor extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$y, create_fragment$y, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Editor",
    			options,
    			id: create_fragment$y.name
    		});
    	}
    }

    /* entries/portal/pages/admin/bprint/instancer/instancer.svelte generated by Svelte v3.48.0 */

    const file$u = "entries/portal/pages/admin/bprint/instancer/instancer.svelte";

    function create_fragment$x(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Instancer";
    			add_location(div, file$u, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$x.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$x($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Instancer', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Instancer> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Instancer extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$x, create_fragment$x, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Instancer",
    			options,
    			id: create_fragment$x.name
    		});
    	}
    }

    /* entries/portal/pages/admin/repo/repo.svelte generated by Svelte v3.48.0 */

    const file$t = "entries/portal/pages/admin/repo/repo.svelte";

    function create_fragment$w(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Repo";
    			add_location(div, file$t, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$w.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$w($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Repo', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Repo> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Repo extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$w, create_fragment$w, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Repo",
    			options,
    			id: create_fragment$w.name
    		});
    	}
    }

    /* entries/portal/pages/admin/repo/repos.svelte generated by Svelte v3.48.0 */

    const file$s = "entries/portal/pages/admin/repo/repos.svelte";

    function create_fragment$v(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Repos";
    			add_location(div, file$s, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$v.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$v($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Repos', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Repos> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Repos extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$v, create_fragment$v, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Repos",
    			options,
    			id: create_fragment$v.name
    		});
    	}
    }

    /* entries/portal/pages/admin/data/loader.svelte generated by Svelte v3.48.0 */

    const file$r = "entries/portal/pages/admin/data/loader.svelte";

    function create_fragment$u(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Loader";
    			add_location(div, file$r, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$u.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$u($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Loader', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Loader> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Loader extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$u, create_fragment$u, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Loader",
    			options,
    			id: create_fragment$u.name
    		});
    	}
    }

    /* entries/portal/pages/admin/data/group/group.svelte generated by Svelte v3.48.0 */

    const file$q = "entries/portal/pages/admin/data/group/group.svelte";

    function create_fragment$t(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Group";
    			add_location(div, file$q, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$t.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$t($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Group', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Group> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Group$1 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$t, create_fragment$t, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Group",
    			options,
    			id: create_fragment$t.name
    		});
    	}
    }

    /* entries/portal/pages/admin/data/group/groups.svelte generated by Svelte v3.48.0 */

    const file$p = "entries/portal/pages/admin/data/group/groups.svelte";

    function create_fragment$s(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Groups";
    			add_location(div, file$p, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$s.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$s($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Groups', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Groups> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Groups extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$s, create_fragment$s, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Groups",
    			options,
    			id: create_fragment$s.name
    		});
    	}
    }

    /* entries/portal/pages/admin/data/table/table.svelte generated by Svelte v3.48.0 */

    const file$o = "entries/portal/pages/admin/data/table/table.svelte";

    function create_fragment$r(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Table";
    			add_location(div, file$o, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$r.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$r($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Table', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Table> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Table extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$r, create_fragment$r, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Table",
    			options,
    			id: create_fragment$r.name
    		});
    	}
    }

    /* entries/portal/pages/admin/data/table/tables.svelte generated by Svelte v3.48.0 */

    const file$n = "entries/portal/pages/admin/data/table/tables.svelte";

    function create_fragment$q(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Tables";
    			add_location(div, file$n, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$q.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$q($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Tables', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Tables> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Tables extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$q, create_fragment$q, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Tables",
    			options,
    			id: create_fragment$q.name
    		});
    	}
    }

    /* entries/portal/pages/admin/data/column/columns.svelte generated by Svelte v3.48.0 */

    const file$m = "entries/portal/pages/admin/data/column/columns.svelte";

    function create_fragment$p(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Columns";
    			add_location(div, file$m, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$p.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$p($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Columns', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Columns> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Columns extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$p, create_fragment$p, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Columns",
    			options,
    			id: create_fragment$p.name
    		});
    	}
    }

    /* entries/portal/pages/admin/data/column/column.svelte generated by Svelte v3.48.0 */

    const file$l = "entries/portal/pages/admin/data/column/column.svelte";

    function create_fragment$o(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Column";
    			add_location(div, file$l, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$o.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$o($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Column', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Column> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Column extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$o, create_fragment$o, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Column",
    			options,
    			id: create_fragment$o.name
    		});
    	}
    }

    /* entries/portal/layout/admin.svelte generated by Svelte v3.48.0 */

    const file$k = "entries/portal/layout/admin.svelte";

    function create_fragment$n(ctx) {
    	let div;
    	let t1;
    	let current;
    	const default_slot_template = /*#slots*/ ctx[1].default;
    	const default_slot = create_slot(default_slot_template, ctx, /*$$scope*/ ctx[0], null);

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Admin Layout";
    			t1 = space();
    			if (default_slot) default_slot.c();
    			add_location(div, file$k, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			insert_dev(target, t1, anchor);

    			if (default_slot) {
    				default_slot.m(target, anchor);
    			}

    			current = true;
    		},
    		p: function update(ctx, [dirty]) {
    			if (default_slot) {
    				if (default_slot.p && (!current || dirty & /*$$scope*/ 1)) {
    					update_slot_base(
    						default_slot,
    						default_slot_template,
    						ctx,
    						/*$$scope*/ ctx[0],
    						!current
    						? get_all_dirty_from_scope(/*$$scope*/ ctx[0])
    						: get_slot_changes(default_slot_template, /*$$scope*/ ctx[0], dirty, null),
    						null
    					);
    				}
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(default_slot, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(default_slot, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			if (detaching) detach_dev(t1);
    			if (default_slot) default_slot.d(detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$n.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$n($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Admin', slots, ['default']);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Admin> was created with unknown prop '${key}'`);
    	});

    	$$self.$$set = $$props => {
    		if ('$$scope' in $$props) $$invalidate(0, $$scope = $$props.$$scope);
    	};

    	return [$$scope, slots];
    }

    class Admin$2 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$n, create_fragment$n, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Admin",
    			options,
    			id: create_fragment$n.name
    		});
    	}
    }

    /* entries/portal/pages/admin/admin.svelte generated by Svelte v3.48.0 */
    const file$j = "entries/portal/pages/admin/admin.svelte";

    // (9:0) <AdminLayout>
    function create_default_slot$2(ctx) {
    	let router;
    	let current;
    	router = new Router({ $$inline: true });

    	const block = {
    		c: function create() {
    			create_component(router.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(router, target, anchor);
    			current = true;
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(router.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(router.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(router, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_default_slot$2.name,
    		type: "slot",
    		source: "(9:0) <AdminLayout>",
    		ctx
    	});

    	return block;
    }

    function create_fragment$m(ctx) {
    	let div;
    	let t1;
    	let adminlayout;
    	let current;

    	adminlayout = new Admin$2({
    			props: {
    				$$slots: { default: [create_default_slot$2] },
    				$$scope: { ctx }
    			},
    			$$inline: true
    		});

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Admin!!!";
    			t1 = space();
    			create_component(adminlayout.$$.fragment);
    			add_location(div, file$j, 6, 0, 123);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			insert_dev(target, t1, anchor);
    			mount_component(adminlayout, target, anchor);
    			current = true;
    		},
    		p: function update(ctx, [dirty]) {
    			const adminlayout_changes = {};

    			if (dirty & /*$$scope*/ 1) {
    				adminlayout_changes.$$scope = { dirty, ctx };
    			}

    			adminlayout.$set(adminlayout_changes);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(adminlayout.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(adminlayout.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			if (detaching) detach_dev(t1);
    			destroy_component(adminlayout, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$m.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$m($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Admin', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Admin> was created with unknown prop '${key}'`);
    	});

    	$$self.$capture_state = () => ({ AdminLayout: Admin$2, Router });
    	return [];
    }

    class Admin$1 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$m, create_fragment$m, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Admin",
    			options,
    			id: create_fragment$m.name
    		});
    	}
    }

    var admin_pages = {
        $$component: Admin$1,
        "/bprint": {
            "/": Bprints,
            "/:bid": Bprint,
            "/:bid/editor": Editor,
            "/:bid/instancer": Instancer,
        },
        repo: {
            "/": Repos,
            "/:rid": Repo,
        },
        data: {
            "/": Loader,
            "/group": Groups,
            "/group/:group": Group$1,
            "/table/:group": Tables,
            "/table/:group/:table": Table,
            "/column/:group/:table": Columns,
            "/column/:group/:table/:column": Column,
        },
        resource: {},
        user: {},
        ugroup: {},
        lens: {},
        tenant: {},
    };

    /* entries/portal/pages/play/play.svelte generated by Svelte v3.48.0 */

    const file$i = "entries/portal/pages/play/play.svelte";

    function create_fragment$l(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Playground";
    			add_location(div, file$i, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$l.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$l($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Play', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Play> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Play extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$l, create_fragment$l, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Play",
    			options,
    			id: create_fragment$l.name
    		});
    	}
    }

    // lauch
    var page_routes = {
        "/": Start,
        "/launch/:app": Player,
        "/data": {
            "/": Loader$3,
            "/:source": Groups$1,
            "/:source/:dgroup": Group$2,
            "/:source/:dgroup/:dtable": Table$1,
        },
        "/cabinet": {
            "/": Loader$2,
            "/:source": Folders,
            "/:source/:folder": Folder$1,
            "/:source/:folder/:file": null,
        },
        "/repo": {
            "/": Loader$1,
            "/:source": Repo$1,
        },
        "/admin": admin_pages,
        "/play": Play,
    };

    /* node_modules/svelte-tooltip/src/SvelteTooltip.svelte generated by Svelte v3.48.0 */

    const file$h = "node_modules/svelte-tooltip/src/SvelteTooltip.svelte";
    const get_custom_tip_slot_changes = dirty => ({});
    const get_custom_tip_slot_context = ctx => ({});

    // (84:4) {:else}
    function create_else_block$2(ctx) {
    	let current;
    	const custom_tip_slot_template = /*#slots*/ ctx[9]["custom-tip"];
    	const custom_tip_slot = create_slot(custom_tip_slot_template, ctx, /*$$scope*/ ctx[8], get_custom_tip_slot_context);

    	const block = {
    		c: function create() {
    			if (custom_tip_slot) custom_tip_slot.c();
    		},
    		m: function mount(target, anchor) {
    			if (custom_tip_slot) {
    				custom_tip_slot.m(target, anchor);
    			}

    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			if (custom_tip_slot) {
    				if (custom_tip_slot.p && (!current || dirty & /*$$scope*/ 256)) {
    					update_slot_base(
    						custom_tip_slot,
    						custom_tip_slot_template,
    						ctx,
    						/*$$scope*/ ctx[8],
    						!current
    						? get_all_dirty_from_scope(/*$$scope*/ ctx[8])
    						: get_slot_changes(custom_tip_slot_template, /*$$scope*/ ctx[8], dirty, get_custom_tip_slot_changes),
    						get_custom_tip_slot_context
    					);
    				}
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(custom_tip_slot, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(custom_tip_slot, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (custom_tip_slot) custom_tip_slot.d(detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_else_block$2.name,
    		type: "else",
    		source: "(84:4) {:else}",
    		ctx
    	});

    	return block;
    }

    // (82:4) {#if tip}
    function create_if_block$6(ctx) {
    	let div;
    	let t;

    	const block = {
    		c: function create() {
    			div = element("div");
    			t = text(/*tip*/ ctx[0]);
    			attr_dev(div, "class", "default-tip svelte-3ftgoq");
    			attr_dev(div, "style", /*style*/ ctx[6]);
    			add_location(div, file$h, 82, 6, 1458);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, t);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*tip*/ 1) set_data_dev(t, /*tip*/ ctx[0]);
    		},
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$6.name,
    		type: "if",
    		source: "(82:4) {#if tip}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$k(ctx) {
    	let div1;
    	let span;
    	let t;
    	let div0;
    	let current_block_type_index;
    	let if_block;
    	let current;
    	const default_slot_template = /*#slots*/ ctx[9].default;
    	const default_slot = create_slot(default_slot_template, ctx, /*$$scope*/ ctx[8], null);
    	const if_block_creators = [create_if_block$6, create_else_block$2];
    	const if_blocks = [];

    	function select_block_type(ctx, dirty) {
    		if (/*tip*/ ctx[0]) return 0;
    		return 1;
    	}

    	current_block_type_index = select_block_type(ctx);
    	if_block = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx);

    	const block = {
    		c: function create() {
    			div1 = element("div");
    			span = element("span");
    			if (default_slot) default_slot.c();
    			t = space();
    			div0 = element("div");
    			if_block.c();
    			attr_dev(span, "class", "tooltip-slot svelte-3ftgoq");
    			add_location(span, file$h, 71, 2, 1280);
    			attr_dev(div0, "class", "tooltip svelte-3ftgoq");
    			toggle_class(div0, "active", /*active*/ ctx[5]);
    			toggle_class(div0, "left", /*left*/ ctx[4]);
    			toggle_class(div0, "right", /*right*/ ctx[2]);
    			toggle_class(div0, "bottom", /*bottom*/ ctx[3]);
    			toggle_class(div0, "top", /*top*/ ctx[1]);
    			add_location(div0, file$h, 74, 2, 1333);
    			attr_dev(div1, "class", "tooltip-wrapper svelte-3ftgoq");
    			add_location(div1, file$h, 70, 0, 1248);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div1, anchor);
    			append_dev(div1, span);

    			if (default_slot) {
    				default_slot.m(span, null);
    			}

    			append_dev(div1, t);
    			append_dev(div1, div0);
    			if_blocks[current_block_type_index].m(div0, null);
    			current = true;
    		},
    		p: function update(ctx, [dirty]) {
    			if (default_slot) {
    				if (default_slot.p && (!current || dirty & /*$$scope*/ 256)) {
    					update_slot_base(
    						default_slot,
    						default_slot_template,
    						ctx,
    						/*$$scope*/ ctx[8],
    						!current
    						? get_all_dirty_from_scope(/*$$scope*/ ctx[8])
    						: get_slot_changes(default_slot_template, /*$$scope*/ ctx[8], dirty, null),
    						null
    					);
    				}
    			}

    			let previous_block_index = current_block_type_index;
    			current_block_type_index = select_block_type(ctx);

    			if (current_block_type_index === previous_block_index) {
    				if_blocks[current_block_type_index].p(ctx, dirty);
    			} else {
    				group_outros();

    				transition_out(if_blocks[previous_block_index], 1, 1, () => {
    					if_blocks[previous_block_index] = null;
    				});

    				check_outros();
    				if_block = if_blocks[current_block_type_index];

    				if (!if_block) {
    					if_block = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx);
    					if_block.c();
    				} else {
    					if_block.p(ctx, dirty);
    				}

    				transition_in(if_block, 1);
    				if_block.m(div0, null);
    			}

    			if (dirty & /*active*/ 32) {
    				toggle_class(div0, "active", /*active*/ ctx[5]);
    			}

    			if (dirty & /*left*/ 16) {
    				toggle_class(div0, "left", /*left*/ ctx[4]);
    			}

    			if (dirty & /*right*/ 4) {
    				toggle_class(div0, "right", /*right*/ ctx[2]);
    			}

    			if (dirty & /*bottom*/ 8) {
    				toggle_class(div0, "bottom", /*bottom*/ ctx[3]);
    			}

    			if (dirty & /*top*/ 2) {
    				toggle_class(div0, "top", /*top*/ ctx[1]);
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(default_slot, local);
    			transition_in(if_block);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(default_slot, local);
    			transition_out(if_block);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div1);
    			if (default_slot) default_slot.d(detaching);
    			if_blocks[current_block_type_index].d();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$k.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$k($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('SvelteTooltip', slots, ['default','custom-tip']);
    	let { tip = "" } = $$props;
    	let { top = false } = $$props;
    	let { right = false } = $$props;
    	let { bottom = false } = $$props;
    	let { left = false } = $$props;
    	let { active = false } = $$props;
    	let { color = "#757575" } = $$props;
    	let style = `background-color: ${color};`;
    	const writable_props = ['tip', 'top', 'right', 'bottom', 'left', 'active', 'color'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<SvelteTooltip> was created with unknown prop '${key}'`);
    	});

    	$$self.$$set = $$props => {
    		if ('tip' in $$props) $$invalidate(0, tip = $$props.tip);
    		if ('top' in $$props) $$invalidate(1, top = $$props.top);
    		if ('right' in $$props) $$invalidate(2, right = $$props.right);
    		if ('bottom' in $$props) $$invalidate(3, bottom = $$props.bottom);
    		if ('left' in $$props) $$invalidate(4, left = $$props.left);
    		if ('active' in $$props) $$invalidate(5, active = $$props.active);
    		if ('color' in $$props) $$invalidate(7, color = $$props.color);
    		if ('$$scope' in $$props) $$invalidate(8, $$scope = $$props.$$scope);
    	};

    	$$self.$capture_state = () => ({
    		tip,
    		top,
    		right,
    		bottom,
    		left,
    		active,
    		color,
    		style
    	});

    	$$self.$inject_state = $$props => {
    		if ('tip' in $$props) $$invalidate(0, tip = $$props.tip);
    		if ('top' in $$props) $$invalidate(1, top = $$props.top);
    		if ('right' in $$props) $$invalidate(2, right = $$props.right);
    		if ('bottom' in $$props) $$invalidate(3, bottom = $$props.bottom);
    		if ('left' in $$props) $$invalidate(4, left = $$props.left);
    		if ('active' in $$props) $$invalidate(5, active = $$props.active);
    		if ('color' in $$props) $$invalidate(7, color = $$props.color);
    		if ('style' in $$props) $$invalidate(6, style = $$props.style);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [tip, top, right, bottom, left, active, style, color, $$scope, slots];
    }

    class SvelteTooltip extends SvelteComponentDev {
    	constructor(options) {
    		super(options);

    		init(this, options, instance$k, create_fragment$k, safe_not_equal, {
    			tip: 0,
    			top: 1,
    			right: 2,
    			bottom: 3,
    			left: 4,
    			active: 5,
    			color: 7
    		});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "SvelteTooltip",
    			options,
    			id: create_fragment$k.name
    		});
    	}

    	get tip() {
    		throw new Error("<SvelteTooltip>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set tip(value) {
    		throw new Error("<SvelteTooltip>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get top() {
    		throw new Error("<SvelteTooltip>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set top(value) {
    		throw new Error("<SvelteTooltip>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get right() {
    		throw new Error("<SvelteTooltip>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set right(value) {
    		throw new Error("<SvelteTooltip>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get bottom() {
    		throw new Error("<SvelteTooltip>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set bottom(value) {
    		throw new Error("<SvelteTooltip>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get left() {
    		throw new Error("<SvelteTooltip>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set left(value) {
    		throw new Error("<SvelteTooltip>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get active() {
    		throw new Error("<SvelteTooltip>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set active(value) {
    		throw new Error("<SvelteTooltip>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get color() {
    		throw new Error("<SvelteTooltip>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set color(value) {
    		throw new Error("<SvelteTooltip>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/xcompo/svg/user.svelte generated by Svelte v3.48.0 */

    const file$g = "entries/xcompo/svg/user.svelte";

    function create_fragment$j(ctx) {
    	let svg;
    	let path;

    	const block = {
    		c: function create() {
    			svg = svg_element("svg");
    			path = svg_element("path");
    			attr_dev(path, "stroke-linecap", "round");
    			attr_dev(path, "stroke-linejoin", "round");
    			attr_dev(path, "stroke-width", "2");
    			attr_dev(path, "d", "M5.121 17.804A13.937 13.937 0 0112 16c2.5 0 4.847.655 6.879 1.804M15 10a3 3 0 11-6 0 3 3 0 016 0zm6 2a9 9 0 11-18 0 9 9 0 0118 0z");
    			add_location(path, file$g, 7, 2, 124);
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg, "class", "h-6 w-6");
    			attr_dev(svg, "fill", "none");
    			attr_dev(svg, "viewBox", "0 0 24 24");
    			attr_dev(svg, "stroke", "currentColor");
    			add_location(svg, file$g, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, svg, anchor);
    			append_dev(svg, path);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(svg);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$j.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$j($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('User', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<User> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class User extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$j, create_fragment$j, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "User",
    			options,
    			id: create_fragment$j.name
    		});
    	}
    }

    /* entries/xcompo/svg/plugin.svelte generated by Svelte v3.48.0 */

    const file$f = "entries/xcompo/svg/plugin.svelte";

    function create_fragment$i(ctx) {
    	let svg;
    	let path;

    	const block = {
    		c: function create() {
    			svg = svg_element("svg");
    			path = svg_element("path");
    			attr_dev(path, "stroke-linecap", "round");
    			attr_dev(path, "stroke-linejoin", "round");
    			attr_dev(path, "stroke-width", "2");
    			attr_dev(path, "d", "M17 14v6m-3-3h6M6 10h2a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v2a2 2 0 002 2zm10 0h2a2 2 0 002-2V6a2 2 0 00-2-2h-2a2 2 0 00-2 2v2a2 2 0 002 2zM6 20h2a2 2 0 002-2v-2a2 2 0 00-2-2H6a2 2 0 00-2 2v2a2 2 0 002 2z");
    			add_location(path, file$f, 7, 2, 124);
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg, "class", "h-6 w-6");
    			attr_dev(svg, "fill", "none");
    			attr_dev(svg, "viewBox", "0 0 24 24");
    			attr_dev(svg, "stroke", "currentColor");
    			add_location(svg, file$f, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, svg, anchor);
    			append_dev(svg, path);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(svg);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$i.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$i($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Plugin', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Plugin> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Plugin extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$i, create_fragment$i, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Plugin",
    			options,
    			id: create_fragment$i.name
    		});
    	}
    }

    /* entries/xcompo/svg/home.svelte generated by Svelte v3.48.0 */

    const file$e = "entries/xcompo/svg/home.svelte";

    function create_fragment$h(ctx) {
    	let svg;
    	let path;

    	const block = {
    		c: function create() {
    			svg = svg_element("svg");
    			path = svg_element("path");
    			attr_dev(path, "stroke-linecap", "round");
    			attr_dev(path, "stroke-linejoin", "round");
    			attr_dev(path, "stroke-width", "2");
    			attr_dev(path, "d", "M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6");
    			add_location(path, file$e, 7, 2, 124);
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg, "class", "h-6 w-6");
    			attr_dev(svg, "fill", "none");
    			attr_dev(svg, "viewBox", "0 0 24 24");
    			attr_dev(svg, "stroke", "currentColor");
    			add_location(svg, file$e, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, svg, anchor);
    			append_dev(svg, path);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(svg);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$h.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$h($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Home', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Home> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Home extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$h, create_fragment$h, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Home",
    			options,
    			id: create_fragment$h.name
    		});
    	}
    }

    /* entries/xcompo/svg/admin.svelte generated by Svelte v3.48.0 */

    const file$d = "entries/xcompo/svg/admin.svelte";

    function create_fragment$g(ctx) {
    	let svg;
    	let path;

    	const block = {
    		c: function create() {
    			svg = svg_element("svg");
    			path = svg_element("path");
    			attr_dev(path, "stroke-linecap", "round");
    			attr_dev(path, "stroke-linejoin", "round");
    			attr_dev(path, "stroke-width", "2");
    			attr_dev(path, "d", "M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4");
    			add_location(path, file$d, 8, 2, 122);
    			attr_dev(svg, "width", "20");
    			attr_dev(svg, "height", "20");
    			attr_dev(svg, "fill", "currentColor");
    			attr_dev(svg, "class", "h-6 w-6");
    			attr_dev(svg, "viewBox", "0 0 24 24");
    			attr_dev(svg, "stroke", "currentColor");
    			add_location(svg, file$d, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, svg, anchor);
    			append_dev(svg, path);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(svg);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$g.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$g($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Admin', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Admin> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Admin extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$g, create_fragment$g, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Admin",
    			options,
    			id: create_fragment$g.name
    		});
    	}
    }

    /* entries/xcompo/svg/folder.svelte generated by Svelte v3.48.0 */

    const file$c = "entries/xcompo/svg/folder.svelte";

    function create_fragment$f(ctx) {
    	let svg;
    	let path;

    	const block = {
    		c: function create() {
    			svg = svg_element("svg");
    			path = svg_element("path");
    			attr_dev(path, "stroke-linecap", "round");
    			attr_dev(path, "stroke-linejoin", "round");
    			attr_dev(path, "stroke-width", "2");
    			attr_dev(path, "d", "M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z");
    			add_location(path, file$c, 7, 2, 124);
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg, "class", "h-6 w-6");
    			attr_dev(svg, "fill", "none");
    			attr_dev(svg, "viewBox", "0 0 24 24");
    			attr_dev(svg, "stroke", "currentColor");
    			add_location(svg, file$c, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, svg, anchor);
    			append_dev(svg, path);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(svg);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$f.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$f($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Folder', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Folder> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Folder extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$f, create_fragment$f, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Folder",
    			options,
    			id: create_fragment$f.name
    		});
    	}
    }

    /* entries/xcompo/svg/group.svelte generated by Svelte v3.48.0 */

    const file$b = "entries/xcompo/svg/group.svelte";

    function create_fragment$e(ctx) {
    	let svg;
    	let path;

    	const block = {
    		c: function create() {
    			svg = svg_element("svg");
    			path = svg_element("path");
    			attr_dev(path, "stroke-linecap", "round");
    			attr_dev(path, "stroke-linejoin", "round");
    			attr_dev(path, "stroke-width", "2");
    			attr_dev(path, "d", "M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10");
    			add_location(path, file$b, 7, 2, 124);
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg, "class", "h-6 w-6");
    			attr_dev(svg, "fill", "none");
    			attr_dev(svg, "viewBox", "0 0 24 24");
    			attr_dev(svg, "stroke", "currentColor");
    			add_location(svg, file$b, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, svg, anchor);
    			append_dev(svg, path);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(svg);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$e.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$e($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Group', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Group> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Group extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$e, create_fragment$e, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Group",
    			options,
    			id: create_fragment$e.name
    		});
    	}
    }

    /* entries/xcompo/svg/store2.svelte generated by Svelte v3.48.0 */

    const file$a = "entries/xcompo/svg/store2.svelte";

    function create_fragment$d(ctx) {
    	let svg;
    	let path;

    	const block = {
    		c: function create() {
    			svg = svg_element("svg");
    			path = svg_element("path");
    			attr_dev(path, "stroke-linecap", "round");
    			attr_dev(path, "stroke-linejoin", "round");
    			attr_dev(path, "stroke-width", "2");
    			attr_dev(path, "d", "M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z");
    			add_location(path, file$a, 7, 2, 124);
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg, "class", "h-6 w-6");
    			attr_dev(svg, "fill", "none");
    			attr_dev(svg, "viewBox", "0 0 24 24");
    			attr_dev(svg, "stroke", "currentColor");
    			add_location(svg, file$a, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, svg, anchor);
    			append_dev(svg, path);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(svg);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$d.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$d($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Store2', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Store2> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Store2 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$d, create_fragment$d, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Store2",
    			options,
    			id: create_fragment$d.name
    		});
    	}
    }

    /* entries/xcompo/svg/notification.svelte generated by Svelte v3.48.0 */

    const file$9 = "entries/xcompo/svg/notification.svelte";

    function create_fragment$c(ctx) {
    	let svg;
    	let path;

    	const block = {
    		c: function create() {
    			svg = svg_element("svg");
    			path = svg_element("path");
    			attr_dev(path, "stroke-linecap", "round");
    			attr_dev(path, "stroke-linejoin", "round");
    			attr_dev(path, "stroke-width", "2");
    			attr_dev(path, "d", "M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9");
    			add_location(path, file$9, 7, 2, 124);
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg, "class", "h-6 w-6");
    			attr_dev(svg, "fill", "none");
    			attr_dev(svg, "viewBox", "0 0 24 24");
    			attr_dev(svg, "stroke", "currentColor");
    			add_location(svg, file$9, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, svg, anchor);
    			append_dev(svg, path);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(svg);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$c.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$c($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Notification', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Notification> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Notification$1 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$c, create_fragment$c, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Notification",
    			options,
    			id: create_fragment$c.name
    		});
    	}
    }

    /* entries/xcompo/svg/logo.svelte generated by Svelte v3.48.0 */

    const file$8 = "entries/xcompo/svg/logo.svelte";

    function create_fragment$b(ctx) {
    	let svg;
    	let path0;
    	let g;
    	let path1;
    	let ellipse0;
    	let path2;
    	let ellipse1;
    	let path3;

    	const block = {
    		c: function create() {
    			svg = svg_element("svg");
    			path0 = svg_element("path");
    			g = svg_element("g");
    			path1 = svg_element("path");
    			ellipse0 = svg_element("ellipse");
    			path2 = svg_element("path");
    			ellipse1 = svg_element("ellipse");
    			path3 = svg_element("path");
    			attr_dev(path0, "d", "M.378.378h69.7v67.81H.378z");
    			attr_dev(path0, "fill", "#925bd5");
    			attr_dev(path0, "stroke", "#fff");
    			attr_dev(path0, "stroke-width", ".757");
    			add_location(path0, file$8, 5, 3, 102);
    			attr_dev(path1, "d", "M26.613 15.364a.26.26 0 0 0-.264.259l-.007.905a3.79 3.79 0 0 0-.465.06l-.228-.879a.26.26 0 1 0-.506.131l.228.88a3.79 3.79 0 0 0-.435.174l-.447-.789a.26.26 0 0 0-.455.258l.447.79a3.79 3.79 0 0 0-.373.283l-.638-.648a.26.26 0 1 0-.373.367l.636.647a3.79 3.79 0 0 0-.289.369l-.781-.459a.26.26 0 1 0-.265.451l.784.461a3.79 3.79 0 0 0-.183.431l-.876-.242a.26.26 0 0 0-.139.504l.873.242-.045.285.597.005c.193-1.592 1.544-2.79 3.148-2.792 1.624.002 2.984 1.229 3.153 2.843l.627.005-.041-.285.879-.228a.26.26 0 0 0-.131-.506l-.88.228a3.79 3.79 0 0 0-.174-.435l.789-.447a.26.26 0 1 0-.258-.455l-.79.447a3.79 3.79 0 0 0-.283-.373l.648-.638a.26.26 0 1 0-.367-.373l-.647.636a3.79 3.79 0 0 0-.369-.289l.459-.781a.26.26 0 1 0-.451-.265l-.461.784a3.79 3.79 0 0 0-.43-.183l.242-.876a.26.26 0 0 0-.504-.14l-.242.873a3.79 3.79 0 0 0-.464-.062l.007-.909a.26.26 0 0 0-.259-.264z");
    			add_location(path1, file$8, 11, 5, 228);
    			attr_dev(ellipse0, "ry", ".938");
    			attr_dev(ellipse0, "rx", ".916");
    			attr_dev(ellipse0, "cy", "19.213");
    			attr_dev(ellipse0, "cx", "27.2");
    			add_location(ellipse0, file$8, 13, 6, 1107);
    			attr_dev(path2, "d", "M43.905 15.364a.26.26 0 0 0-.264.259l-.007.905a3.79 3.79 0 0 0-.465.06l-.228-.879a.26.26 0 0 0-.506.131l.228.88a3.79 3.79 0 0 0-.435.174l-.447-.789a.26.26 0 0 0-.455.258l.447.79a3.79 3.79 0 0 0-.373.283l-.638-.648a.26.26 0 1 0-.373.367l.636.647a3.79 3.79 0 0 0-.289.369l-.781-.459a.26.26 0 1 0-.265.451l.784.461a3.79 3.79 0 0 0-.183.431l-.876-.242a.26.26 0 0 0-.139.504l.873.242-.045.285.597.005c.193-1.592 1.544-2.79 3.148-2.792 1.624.002 2.984 1.229 3.153 2.843l.627.005-.041-.285.879-.228a.26.26 0 1 0-.131-.506l-.88.228a3.79 3.79 0 0 0-.174-.435l.789-.447a.26.26 0 0 0-.258-.455l-.79.447a3.79 3.79 0 0 0-.283-.373l.648-.638a.26.26 0 1 0-.367-.373l-.647.636a3.79 3.79 0 0 0-.369-.289l.459-.781a.26.26 0 1 0-.451-.265l-.461.784a3.79 3.79 0 0 0-.43-.183l.242-.876a.26.26 0 0 0-.504-.14l-.242.873a3.79 3.79 0 0 0-.464-.062l.007-.909a.26.26 0 0 0-.259-.264z");
    			add_location(path2, file$8, 13, 59, 1160);
    			attr_dev(ellipse1, "cx", "44.493");
    			attr_dev(ellipse1, "cy", "19.307");
    			attr_dev(ellipse1, "rx", ".916");
    			attr_dev(ellipse1, "ry", ".938");
    			add_location(ellipse1, file$8, 15, 6, 2039);
    			attr_dev(path3, "d", "M17.178 25.113l30.918-.44-1.145 6.518-29.597-3.171zm11.715 5.373l8.632 2.026-.088 28.011-3.435.617z");
    			add_location(path3, file$8, 15, 61, 2094);
    			attr_dev(g, "fill", "#fff");
    			add_location(g, file$8, 10, 4, 208);
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg, "width", "35");
    			attr_dev(svg, "height", "35");
    			attr_dev(svg, "viewBox", "0 0 70.457 68.567");
    			add_location(svg, file$8, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, svg, anchor);
    			append_dev(svg, path0);
    			append_dev(svg, g);
    			append_dev(g, path1);
    			append_dev(g, ellipse0);
    			append_dev(g, path2);
    			append_dev(g, ellipse1);
    			append_dev(g, path3);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(svg);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$b.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$b($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Logo', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Logo> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Logo extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$b, create_fragment$b, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Logo",
    			options,
    			id: create_fragment$b.name
    		});
    	}
    }

    /* entries/portal/layout/main.svelte generated by Svelte v3.48.0 */
    const file$7 = "entries/portal/layout/main.svelte";

    function get_each_context$1(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[11] = list[i];
    	return child_ctx;
    }

    function get_each_context_1(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[11] = list[i];
    	return child_ctx;
    }

    // (35:10) <SvelteTooltip tip={link[1]} right color="#7c3aed">
    function create_default_slot$1(ctx) {
    	let a;
    	let switch_instance;
    	let current;
    	var switch_value = /*link*/ ctx[11][0];

    	function switch_props(ctx) {
    		return { $$inline: true };
    	}

    	if (switch_value) {
    		switch_instance = new switch_value(switch_props());
    	}

    	const block = {
    		c: function create() {
    			a = element("a");
    			if (switch_instance) create_component(switch_instance.$$.fragment);
    			attr_dev(a, "href", "");
    			attr_dev(a, "class", "h-12 px-4 flex justify-center items-center w-full text-gray-700 focus:text-orange-500");
    			add_location(a, file$7, 35, 12, 1274);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, a, anchor);

    			if (switch_instance) {
    				mount_component(switch_instance, a, null);
    			}

    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			if (switch_value !== (switch_value = /*link*/ ctx[11][0])) {
    				if (switch_instance) {
    					group_outros();
    					const old_component = switch_instance;

    					transition_out(old_component.$$.fragment, 1, 0, () => {
    						destroy_component(old_component, 1);
    					});

    					check_outros();
    				}

    				if (switch_value) {
    					switch_instance = new switch_value(switch_props());
    					create_component(switch_instance.$$.fragment);
    					transition_in(switch_instance.$$.fragment, 1);
    					mount_component(switch_instance, a, null);
    				} else {
    					switch_instance = null;
    				}
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			if (switch_instance) transition_in(switch_instance.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			if (switch_instance) transition_out(switch_instance.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(a);
    			if (switch_instance) destroy_component(switch_instance);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_default_slot$1.name,
    		type: "slot",
    		source: "(35:10) <SvelteTooltip tip={link[1]} right color=\\\"#7c3aed\\\">",
    		ctx
    	});

    	return block;
    }

    // (33:6) {#each route_links as link}
    function create_each_block_1(ctx) {
    	let li;
    	let sveltetooltip;
    	let t;
    	let current;

    	sveltetooltip = new SvelteTooltip({
    			props: {
    				tip: /*link*/ ctx[11][1],
    				right: true,
    				color: "#7c3aed",
    				$$slots: { default: [create_default_slot$1] },
    				$$scope: { ctx }
    			},
    			$$inline: true
    		});

    	const block = {
    		c: function create() {
    			li = element("li");
    			create_component(sveltetooltip.$$.fragment);
    			t = space();
    			attr_dev(li, "class", "text-white hover:bg-gray-100");
    			add_location(li, file$7, 33, 8, 1158);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, li, anchor);
    			mount_component(sveltetooltip, li, null);
    			append_dev(li, t);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			const sveltetooltip_changes = {};

    			if (dirty & /*$$scope*/ 1024) {
    				sveltetooltip_changes.$$scope = { dirty, ctx };
    			}

    			sveltetooltip.$set(sveltetooltip_changes);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(sveltetooltip.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(sveltetooltip.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(li);
    			destroy_component(sveltetooltip);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block_1.name,
    		type: "each",
    		source: "(33:6) {#each route_links as link}",
    		ctx
    	});

    	return block;
    }

    // (56:10) {#if pending_notification}
    function create_if_block_2$1(ctx) {
    	let span2;
    	let span0;
    	let t;
    	let span1;

    	const block = {
    		c: function create() {
    			span2 = element("span");
    			span0 = element("span");
    			t = space();
    			span1 = element("span");
    			attr_dev(span0, "class", "animate-ping absolute inline-flex h-3 w-3 rounded-full bg-pink-400 opacity-75");
    			add_location(span0, file$7, 57, 14, 2035);
    			attr_dev(span1, "class", "relative inline-flex rounded-full h-3 w-3 bg-pink-500");
    			add_location(span1, file$7, 60, 14, 2174);
    			attr_dev(span2, "class", "flex absolute -mt-5 ml-4");
    			add_location(span2, file$7, 56, 12, 1981);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, span2, anchor);
    			append_dev(span2, span0);
    			append_dev(span2, t);
    			append_dev(span2, span1);
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(span2);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_2$1.name,
    		type: "if",
    		source: "(56:10) {#if pending_notification}",
    		ctx
    	});

    	return block;
    }

    // (136:6) {#if pending_notification}
    function create_if_block_1$1(ctx) {
    	let span2;
    	let span0;
    	let t;
    	let span1;

    	const block = {
    		c: function create() {
    			span2 = element("span");
    			span0 = element("span");
    			t = space();
    			span1 = element("span");
    			attr_dev(span0, "class", "animate-ping absolute inline-flex h-3 w-3 rounded-full bg-pink-400 opacity-75");
    			add_location(span0, file$7, 137, 10, 4330);
    			attr_dev(span1, "class", "relative inline-flex rounded-full h-3 w-3 bg-pink-500");
    			add_location(span1, file$7, 140, 10, 4457);
    			attr_dev(span2, "class", "flex absolute -mt-5 ml-4");
    			add_location(span2, file$7, 136, 8, 4280);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, span2, anchor);
    			append_dev(span2, span0);
    			append_dev(span2, t);
    			append_dev(span2, span1);
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(span2);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_1$1.name,
    		type: "if",
    		source: "(136:6) {#if pending_notification}",
    		ctx
    	});

    	return block;
    }

    // (151:2) {#if __open_menu}
    function create_if_block$5(ctx) {
    	let div3;
    	let div2;
    	let div1;
    	let div0;
    	let button0;
    	let svg0;
    	let path0;
    	let path1;
    	let t0;
    	let ul;
    	let t1;
    	let li;
    	let button1;
    	let svg1;
    	let path2;
    	let polyline;
    	let line;
    	let current;
    	let mounted;
    	let dispose;
    	let each_value = /*route_links*/ ctx[2];
    	validate_each_argument(each_value);
    	let each_blocks = [];

    	for (let i = 0; i < each_value.length; i += 1) {
    		each_blocks[i] = create_each_block$1(get_each_context$1(ctx, each_value, i));
    	}

    	const out = i => transition_out(each_blocks[i], 1, 1, () => {
    		each_blocks[i] = null;
    	});

    	const block = {
    		c: function create() {
    			div3 = element("div");
    			div2 = element("div");
    			div1 = element("div");
    			div0 = element("div");
    			button0 = element("button");
    			svg0 = svg_element("svg");
    			path0 = svg_element("path");
    			path1 = svg_element("path");
    			t0 = space();
    			ul = element("ul");

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].c();
    			}

    			t1 = space();
    			li = element("li");
    			button1 = element("button");
    			svg1 = svg_element("svg");
    			path2 = svg_element("path");
    			polyline = svg_element("polyline");
    			line = svg_element("line");
    			attr_dev(path0, "d", "M18 6L6 18");
    			attr_dev(path0, "stroke", "#4B5563");
    			attr_dev(path0, "stroke-width", "1.25");
    			attr_dev(path0, "stroke-linecap", "round");
    			attr_dev(path0, "stroke-linejoin", "round");
    			add_location(path0, file$7, 170, 16, 5468);
    			attr_dev(path1, "d", "M6 6L18 18");
    			attr_dev(path1, "stroke", "#4B5563");
    			attr_dev(path1, "stroke-width", "1.25");
    			attr_dev(path1, "stroke-linecap", "round");
    			attr_dev(path1, "stroke-linejoin", "round");
    			add_location(path1, file$7, 177, 16, 5698);
    			attr_dev(svg0, "class", "h-6 w-6");
    			attr_dev(svg0, "viewBox", "0 0 24 24");
    			attr_dev(svg0, "fill", "none");
    			attr_dev(svg0, "xmlns", "http://www.w3.org/2000/svg");
    			add_location(svg0, file$7, 164, 14, 5284);
    			attr_dev(button0, "aria-label", "close menu modal");
    			attr_dev(button0, "class", "focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 cursor-pointer rounded border mr-2");
    			add_location(button0, file$7, 157, 12, 4982);
    			attr_dev(div0, "class", "absolute right-4");
    			add_location(div0, file$7, 156, 10, 4939);
    			attr_dev(path2, "d", "M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4");
    			add_location(path2, file$7, 221, 18, 7226);
    			attr_dev(polyline, "points", "16 17 21 12 16 7");
    			add_location(polyline, file$7, 222, 18, 7297);
    			attr_dev(line, "x1", "21");
    			attr_dev(line, "y1", "12");
    			attr_dev(line, "x2", "9");
    			attr_dev(line, "y2", "12");
    			add_location(line, file$7, 223, 18, 7354);
    			attr_dev(svg1, "class", "h-5 w-5 text-red-700");
    			attr_dev(svg1, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg1, "width", "24");
    			attr_dev(svg1, "height", "24");
    			attr_dev(svg1, "viewBox", "0 0 24 24");
    			attr_dev(svg1, "fill", "none");
    			attr_dev(svg1, "stroke", "currentColor");
    			attr_dev(svg1, "stroke-width", "2");
    			attr_dev(svg1, "stroke-linecap", "round");
    			attr_dev(svg1, "stroke-linejoin", "round");
    			add_location(svg1, file$7, 209, 16, 6800);
    			attr_dev(button1, "class", "h-12 cursor-pointer mx-auto flex justify-center items-center w-full focus:text-orange-500 hover:bg-red-200 focus:outline-none");
    			add_location(button1, file$7, 205, 14, 6576);
    			attr_dev(li, "class", "border rounded");
    			add_location(li, file$7, 204, 12, 6534);
    			attr_dev(ul, "class", "space-y-4 pt-10");
    			add_location(ul, file$7, 188, 10, 5983);
    			attr_dev(div1, "class", "h-full bg-white rounded p-5");
    			add_location(div1, file$7, 155, 8, 4887);
    			attr_dev(div2, "class", "h-full w-full absolute transform translate-x-0 transition ease-in-out duration-700 p-5");
    			add_location(div2, file$7, 152, 6, 4763);
    			attr_dev(div3, "class", "h-screen w-screen bg-gray-600 bg-opacity-90 fixed z-50");
    			add_location(div3, file$7, 151, 4, 4688);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div3, anchor);
    			append_dev(div3, div2);
    			append_dev(div2, div1);
    			append_dev(div1, div0);
    			append_dev(div0, button0);
    			append_dev(button0, svg0);
    			append_dev(svg0, path0);
    			append_dev(svg0, path1);
    			append_dev(div1, t0);
    			append_dev(div1, ul);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].m(ul, null);
    			}

    			append_dev(ul, t1);
    			append_dev(ul, li);
    			append_dev(li, button1);
    			append_dev(button1, svg1);
    			append_dev(svg1, path2);
    			append_dev(svg1, polyline);
    			append_dev(svg1, line);
    			current = true;

    			if (!mounted) {
    				dispose = listen_dev(button0, "click", /*click_handler_3*/ ctx[8], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*__open_menu, route_links*/ 6) {
    				each_value = /*route_links*/ ctx[2];
    				validate_each_argument(each_value);
    				let i;

    				for (i = 0; i < each_value.length; i += 1) {
    					const child_ctx = get_each_context$1(ctx, each_value, i);

    					if (each_blocks[i]) {
    						each_blocks[i].p(child_ctx, dirty);
    						transition_in(each_blocks[i], 1);
    					} else {
    						each_blocks[i] = create_each_block$1(child_ctx);
    						each_blocks[i].c();
    						transition_in(each_blocks[i], 1);
    						each_blocks[i].m(ul, t1);
    					}
    				}

    				group_outros();

    				for (i = each_value.length; i < each_blocks.length; i += 1) {
    					out(i);
    				}

    				check_outros();
    			}
    		},
    		i: function intro(local) {
    			if (current) return;

    			for (let i = 0; i < each_value.length; i += 1) {
    				transition_in(each_blocks[i]);
    			}

    			current = true;
    		},
    		o: function outro(local) {
    			each_blocks = each_blocks.filter(Boolean);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				transition_out(each_blocks[i]);
    			}

    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div3);
    			destroy_each(each_blocks, detaching);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$5.name,
    		type: "if",
    		source: "(151:2) {#if __open_menu}",
    		ctx
    	});

    	return block;
    }

    // (190:12) {#each route_links as link}
    function create_each_block$1(ctx) {
    	let li;
    	let a;
    	let switch_instance;
    	let t0;
    	let t1_value = /*link*/ ctx[11][1] + "";
    	let t1;
    	let current;
    	let mounted;
    	let dispose;
    	var switch_value = /*link*/ ctx[11][0];

    	function switch_props(ctx) {
    		return { $$inline: true };
    	}

    	if (switch_value) {
    		switch_instance = new switch_value(switch_props());
    	}

    	const block = {
    		c: function create() {
    			li = element("li");
    			a = element("a");
    			if (switch_instance) create_component(switch_instance.$$.fragment);
    			t0 = space();
    			t1 = text(t1_value);
    			attr_dev(a, "href", "#");
    			attr_dev(a, "class", "h-12 px-4 flex justify-center items-center w-full text-gray-700 focus:text-orange-500 uppercase");
    			add_location(a, file$7, 191, 16, 6110);
    			attr_dev(li, "class", "border rounded");
    			add_location(li, file$7, 190, 14, 6066);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, li, anchor);
    			append_dev(li, a);

    			if (switch_instance) {
    				mount_component(switch_instance, a, null);
    			}

    			append_dev(a, t0);
    			append_dev(a, t1);
    			current = true;

    			if (!mounted) {
    				dispose = listen_dev(a, "click", /*click_handler_4*/ ctx[9], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			if (switch_value !== (switch_value = /*link*/ ctx[11][0])) {
    				if (switch_instance) {
    					group_outros();
    					const old_component = switch_instance;

    					transition_out(old_component.$$.fragment, 1, 0, () => {
    						destroy_component(old_component, 1);
    					});

    					check_outros();
    				}

    				if (switch_value) {
    					switch_instance = new switch_value(switch_props());
    					create_component(switch_instance.$$.fragment);
    					transition_in(switch_instance.$$.fragment, 1);
    					mount_component(switch_instance, a, t0);
    				} else {
    					switch_instance = null;
    				}
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			if (switch_instance) transition_in(switch_instance.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			if (switch_instance) transition_out(switch_instance.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(li);
    			if (switch_instance) destroy_component(switch_instance);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block$1.name,
    		type: "each",
    		source: "(190:12) {#each route_links as link}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$a(ctx) {
    	let div9;
    	let nav0;
    	let div0;
    	let a0;
    	let logo;
    	let t0;
    	let ul;
    	let t1;
    	let div4;
    	let div3;
    	let div1;
    	let t2;
    	let notificationicon0;
    	let t3;
    	let div2;
    	let usericon0;
    	let t4;
    	let button;
    	let svg0;
    	let path0;
    	let polyline;
    	let line;
    	let t5;
    	let nav1;
    	let div5;
    	let a1;
    	let svg1;
    	let path1;
    	let t6;
    	let div6;
    	let notificationicon1;
    	let t7;
    	let t8;
    	let div7;
    	let usericon1;
    	let t9;
    	let t10;
    	let div8;
    	let current;
    	let mounted;
    	let dispose;
    	logo = new Logo({ $$inline: true });
    	let each_value_1 = /*route_links*/ ctx[2];
    	validate_each_argument(each_value_1);
    	let each_blocks = [];

    	for (let i = 0; i < each_value_1.length; i += 1) {
    		each_blocks[i] = create_each_block_1(get_each_context_1(ctx, each_value_1, i));
    	}

    	const out = i => transition_out(each_blocks[i], 1, 1, () => {
    		each_blocks[i] = null;
    	});

    	let if_block0 = /*pending_notification*/ ctx[0] && create_if_block_2$1(ctx);
    	notificationicon0 = new Notification$1({ $$inline: true });
    	usericon0 = new User({ $$inline: true });
    	notificationicon1 = new Notification$1({ $$inline: true });
    	let if_block1 = /*pending_notification*/ ctx[0] && create_if_block_1$1(ctx);
    	usericon1 = new User({ $$inline: true });
    	let if_block2 = /*__open_menu*/ ctx[1] && create_if_block$5(ctx);
    	const default_slot_template = /*#slots*/ ctx[4].default;
    	const default_slot = create_slot(default_slot_template, ctx, /*$$scope*/ ctx[10], null);

    	const block = {
    		c: function create() {
    			div9 = element("div");
    			nav0 = element("nav");
    			div0 = element("div");
    			a0 = element("a");
    			create_component(logo.$$.fragment);
    			t0 = space();
    			ul = element("ul");

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].c();
    			}

    			t1 = space();
    			div4 = element("div");
    			div3 = element("div");
    			div1 = element("div");
    			if (if_block0) if_block0.c();
    			t2 = space();
    			create_component(notificationicon0.$$.fragment);
    			t3 = space();
    			div2 = element("div");
    			create_component(usericon0.$$.fragment);
    			t4 = space();
    			button = element("button");
    			svg0 = svg_element("svg");
    			path0 = svg_element("path");
    			polyline = svg_element("polyline");
    			line = svg_element("line");
    			t5 = space();
    			nav1 = element("nav");
    			div5 = element("div");
    			a1 = element("a");
    			svg1 = svg_element("svg");
    			path1 = svg_element("path");
    			t6 = space();
    			div6 = element("div");
    			create_component(notificationicon1.$$.fragment);
    			t7 = space();
    			if (if_block1) if_block1.c();
    			t8 = space();
    			div7 = element("div");
    			create_component(usericon1.$$.fragment);
    			t9 = space();
    			if (if_block2) if_block2.c();
    			t10 = space();
    			div8 = element("div");
    			if (default_slot) default_slot.c();
    			attr_dev(a0, "class", "h-6 mx-auto");
    			attr_dev(a0, "href", "http://svelte.dev/");
    			add_location(a0, file$7, 26, 6, 1017);
    			attr_dev(div0, "class", "h-16 flex items-center w-full");
    			add_location(div0, file$7, 24, 4, 939);
    			add_location(ul, file$7, 31, 4, 1111);
    			attr_dev(div1, "class", "h-16 cursor-pointer mx-auto w-full flex justify-center items-center focus:text-orange-500 hover:bg-red-200 focus:outline-none relative");
    			add_location(div1, file$7, 48, 8, 1667);
    			attr_dev(div2, "class", "h-12 cursor-pointer mx-auto flex justify-center items-center w-full focus:text-orange-500 hover:bg-red-200 focus:outline-none");
    			add_location(div2, file$7, 69, 8, 2367);
    			attr_dev(path0, "d", "M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4");
    			add_location(path0, file$7, 92, 12, 3137);
    			attr_dev(polyline, "points", "16 17 21 12 16 7");
    			add_location(polyline, file$7, 93, 12, 3202);
    			attr_dev(line, "x1", "21");
    			attr_dev(line, "y1", "12");
    			attr_dev(line, "x2", "9");
    			attr_dev(line, "y2", "12");
    			add_location(line, file$7, 94, 12, 3253);
    			attr_dev(svg0, "class", "h-5 w-5 text-red-700");
    			attr_dev(svg0, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg0, "width", "24");
    			attr_dev(svg0, "height", "24");
    			attr_dev(svg0, "viewBox", "0 0 24 24");
    			attr_dev(svg0, "fill", "none");
    			attr_dev(svg0, "stroke", "currentColor");
    			attr_dev(svg0, "stroke-width", "2");
    			attr_dev(svg0, "stroke-linecap", "round");
    			attr_dev(svg0, "stroke-linejoin", "round");
    			add_location(svg0, file$7, 80, 10, 2783);
    			attr_dev(button, "class", "h-12 cursor-pointer mx-auto flex justify-center items-center w-full focus:text-orange-500 hover:bg-red-200 focus:outline-none");
    			add_location(button, file$7, 76, 8, 2587);
    			attr_dev(div3, "class", "flex flex-col items-center w-full");
    			add_location(div3, file$7, 47, 6, 1611);
    			attr_dev(div4, "class", "mt-auto flex items-center w-full");
    			add_location(div4, file$7, 46, 4, 1558);
    			attr_dev(nav0, "class", "flex-col items-center bg-blue-200 text-gray-700 h-full w-12 shadow-lg tx-nav-desktop svelte-1u19ika");
    			add_location(nav0, file$7, 19, 2, 803);
    			attr_dev(path1, "stroke-linecap", "round");
    			attr_dev(path1, "stroke-linejoin", "round");
    			attr_dev(path1, "stroke-width", "2");
    			attr_dev(path1, "d", "M4 6h16M4 12h16M4 18h16");
    			add_location(path1, file$7, 119, 10, 3874);
    			attr_dev(svg1, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg1, "class", "h-6 w-6 hover:text-gray-200");
    			attr_dev(svg1, "fill", "none");
    			attr_dev(svg1, "viewBox", "0 0 24 24");
    			attr_dev(svg1, "stroke", "currentColor");
    			add_location(svg1, file$7, 112, 8, 3674);
    			attr_dev(a1, "class", "navbar-burger self-center mr-12");
    			attr_dev(a1, "href", "#");
    			add_location(a1, file$7, 105, 6, 3515);
    			attr_dev(div5, "class", "py-2 flex w-full items-center");
    			add_location(div5, file$7, 104, 4, 3464);
    			attr_dev(div6, "class", "flex mr-6 items-center cursor-pointer");
    			add_location(div6, file$7, 129, 4, 4105);
    			attr_dev(div7, "class", "flex mr-6 items-center cursor-pointer");
    			add_location(div7, file$7, 145, 4, 4572);
    			attr_dev(nav1, "class", "justify-between w-screen bg-blue-200 text-gray-700 shadow-lg tx-nav-mobile svelte-1u19ika");
    			add_location(nav1, file$7, 101, 2, 3364);
    			attr_dev(div8, "class", "h-screen overflow-auto tx-main svelte-1u19ika");
    			set_style(div8, "width", "95%");
    			set_style(div8, "width", "calc(100vw - 3rem)");
    			add_location(div8, file$7, 233, 2, 7525);
    			attr_dev(div9, "class", "h-screen w-screen flex tx-root svelte-1u19ika");
    			add_location(div9, file$7, 17, 0, 733);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div9, anchor);
    			append_dev(div9, nav0);
    			append_dev(nav0, div0);
    			append_dev(div0, a0);
    			mount_component(logo, a0, null);
    			append_dev(nav0, t0);
    			append_dev(nav0, ul);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].m(ul, null);
    			}

    			append_dev(nav0, t1);
    			append_dev(nav0, div4);
    			append_dev(div4, div3);
    			append_dev(div3, div1);
    			if (if_block0) if_block0.m(div1, null);
    			append_dev(div1, t2);
    			mount_component(notificationicon0, div1, null);
    			append_dev(div3, t3);
    			append_dev(div3, div2);
    			mount_component(usericon0, div2, null);
    			append_dev(div3, t4);
    			append_dev(div3, button);
    			append_dev(button, svg0);
    			append_dev(svg0, path0);
    			append_dev(svg0, polyline);
    			append_dev(svg0, line);
    			append_dev(div9, t5);
    			append_dev(div9, nav1);
    			append_dev(nav1, div5);
    			append_dev(div5, a1);
    			append_dev(a1, svg1);
    			append_dev(svg1, path1);
    			append_dev(nav1, t6);
    			append_dev(nav1, div6);
    			mount_component(notificationicon1, div6, null);
    			append_dev(div6, t7);
    			if (if_block1) if_block1.m(div6, null);
    			append_dev(nav1, t8);
    			append_dev(nav1, div7);
    			mount_component(usericon1, div7, null);
    			append_dev(div9, t9);
    			if (if_block2) if_block2.m(div9, null);
    			append_dev(div9, t10);
    			append_dev(div9, div8);

    			if (default_slot) {
    				default_slot.m(div8, null);
    			}

    			current = true;

    			if (!mounted) {
    				dispose = [
    					listen_dev(div1, "click", /*click_handler*/ ctx[5], false, false, false),
    					listen_dev(a1, "click", /*click_handler_1*/ ctx[6], false, false, false),
    					listen_dev(div6, "click", /*click_handler_2*/ ctx[7], false, false, false)
    				];

    				mounted = true;
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*route_links*/ 4) {
    				each_value_1 = /*route_links*/ ctx[2];
    				validate_each_argument(each_value_1);
    				let i;

    				for (i = 0; i < each_value_1.length; i += 1) {
    					const child_ctx = get_each_context_1(ctx, each_value_1, i);

    					if (each_blocks[i]) {
    						each_blocks[i].p(child_ctx, dirty);
    						transition_in(each_blocks[i], 1);
    					} else {
    						each_blocks[i] = create_each_block_1(child_ctx);
    						each_blocks[i].c();
    						transition_in(each_blocks[i], 1);
    						each_blocks[i].m(ul, null);
    					}
    				}

    				group_outros();

    				for (i = each_value_1.length; i < each_blocks.length; i += 1) {
    					out(i);
    				}

    				check_outros();
    			}

    			if (/*pending_notification*/ ctx[0]) {
    				if (if_block0) ; else {
    					if_block0 = create_if_block_2$1(ctx);
    					if_block0.c();
    					if_block0.m(div1, t2);
    				}
    			} else if (if_block0) {
    				if_block0.d(1);
    				if_block0 = null;
    			}

    			if (/*pending_notification*/ ctx[0]) {
    				if (if_block1) ; else {
    					if_block1 = create_if_block_1$1(ctx);
    					if_block1.c();
    					if_block1.m(div6, null);
    				}
    			} else if (if_block1) {
    				if_block1.d(1);
    				if_block1 = null;
    			}

    			if (/*__open_menu*/ ctx[1]) {
    				if (if_block2) {
    					if_block2.p(ctx, dirty);

    					if (dirty & /*__open_menu*/ 2) {
    						transition_in(if_block2, 1);
    					}
    				} else {
    					if_block2 = create_if_block$5(ctx);
    					if_block2.c();
    					transition_in(if_block2, 1);
    					if_block2.m(div9, t10);
    				}
    			} else if (if_block2) {
    				group_outros();

    				transition_out(if_block2, 1, 1, () => {
    					if_block2 = null;
    				});

    				check_outros();
    			}

    			if (default_slot) {
    				if (default_slot.p && (!current || dirty & /*$$scope*/ 1024)) {
    					update_slot_base(
    						default_slot,
    						default_slot_template,
    						ctx,
    						/*$$scope*/ ctx[10],
    						!current
    						? get_all_dirty_from_scope(/*$$scope*/ ctx[10])
    						: get_slot_changes(default_slot_template, /*$$scope*/ ctx[10], dirty, null),
    						null
    					);
    				}
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(logo.$$.fragment, local);

    			for (let i = 0; i < each_value_1.length; i += 1) {
    				transition_in(each_blocks[i]);
    			}

    			transition_in(notificationicon0.$$.fragment, local);
    			transition_in(usericon0.$$.fragment, local);
    			transition_in(notificationicon1.$$.fragment, local);
    			transition_in(usericon1.$$.fragment, local);
    			transition_in(if_block2);
    			transition_in(default_slot, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(logo.$$.fragment, local);
    			each_blocks = each_blocks.filter(Boolean);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				transition_out(each_blocks[i]);
    			}

    			transition_out(notificationicon0.$$.fragment, local);
    			transition_out(usericon0.$$.fragment, local);
    			transition_out(notificationicon1.$$.fragment, local);
    			transition_out(usericon1.$$.fragment, local);
    			transition_out(if_block2);
    			transition_out(default_slot, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div9);
    			destroy_component(logo);
    			destroy_each(each_blocks, detaching);
    			if (if_block0) if_block0.d();
    			destroy_component(notificationicon0);
    			destroy_component(usericon0);
    			destroy_component(notificationicon1);
    			if (if_block1) if_block1.d();
    			destroy_component(usericon1);
    			if (if_block2) if_block2.d();
    			if (default_slot) default_slot.d(detaching);
    			mounted = false;
    			run_all(dispose);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$a.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$a($$self, $$props, $$invalidate) {
    	let __open_menu;
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Main', slots, ['default']);
    	let { pending_notification = false } = $$props;

    	const route_links = [
    		[Home, "start", "/z/portal"],
    		[Group, "data tables", "/z/portal/dtable_load"],
    		[Plugin, "plug apps", "/z/portal/apps_launcher"],
    		[Admin, "admin", "/z/portal/admin/plugs"],
    		[Folder, "cabinet", "/z/portal/cabinet_load"],
    		[Store2, "store", "/z/portal/store"]
    	];

    	const dispatch = createEventDispatcher();
    	const writable_props = ['pending_notification'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Main> was created with unknown prop '${key}'`);
    	});

    	const click_handler = () => dispatch("ntoggle");

    	const click_handler_1 = () => {
    		$$invalidate(1, __open_menu = !__open_menu);
    	};

    	const click_handler_2 = () => dispatch("ntoggle");

    	const click_handler_3 = () => {
    		$$invalidate(1, __open_menu = !__open_menu);
    	};

    	const click_handler_4 = () => {
    		$$invalidate(1, __open_menu = false);
    	};

    	$$self.$$set = $$props => {
    		if ('pending_notification' in $$props) $$invalidate(0, pending_notification = $$props.pending_notification);
    		if ('$$scope' in $$props) $$invalidate(10, $$scope = $$props.$$scope);
    	};

    	$$self.$capture_state = () => ({
    		createEventDispatcher,
    		SvelteTooltip,
    		AdminIcon: Admin,
    		FolderIcon: Folder,
    		GroupIcon: Group,
    		PlugIcon: Plugin,
    		HomeIcon: Home,
    		StoreIcon: Store2,
    		NotificationIcon: Notification$1,
    		UserIcon: User,
    		Logo,
    		pending_notification,
    		route_links,
    		dispatch,
    		__open_menu
    	});

    	$$self.$inject_state = $$props => {
    		if ('pending_notification' in $$props) $$invalidate(0, pending_notification = $$props.pending_notification);
    		if ('__open_menu' in $$props) $$invalidate(1, __open_menu = $$props.__open_menu);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	$$invalidate(1, __open_menu = false);

    	return [
    		pending_notification,
    		__open_menu,
    		route_links,
    		dispatch,
    		slots,
    		click_handler,
    		click_handler_1,
    		click_handler_2,
    		click_handler_3,
    		click_handler_4,
    		$$scope
    	];
    }

    class Main extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$a, create_fragment$a, safe_not_equal, { pending_notification: 0 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Main",
    			options,
    			id: create_fragment$a.name
    		});
    	}

    	get pending_notification() {
    		throw new Error("<Main>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set pending_notification(value) {
    		throw new Error("<Main>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/portal/layout/big_modal.svelte generated by Svelte v3.48.0 */

    const file$6 = "entries/portal/layout/big_modal.svelte";

    // (19:2) {#if _show && current_component}
    function create_if_block$4(ctx) {
    	let modal_wrapper;
    	let div;
    	let svg;
    	let path;
    	let t;
    	let modal_section;
    	let switch_instance;
    	let current;
    	let mounted;
    	let dispose;
    	const switch_instance_spread_levels = [/*props*/ ctx[1]];
    	var switch_value = /*current_component*/ ctx[0];

    	function switch_props(ctx) {
    		let switch_instance_props = {};

    		for (let i = 0; i < switch_instance_spread_levels.length; i += 1) {
    			switch_instance_props = assign(switch_instance_props, switch_instance_spread_levels[i]);
    		}

    		return {
    			props: switch_instance_props,
    			$$inline: true
    		};
    	}

    	if (switch_value) {
    		switch_instance = new switch_value(switch_props());
    	}

    	const block = {
    		c: function create() {
    			modal_wrapper = element("modal-wrapper");
    			div = element("div");
    			svg = svg_element("svg");
    			path = svg_element("path");
    			t = space();
    			modal_section = element("modal-section");
    			if (switch_instance) create_component(switch_instance.$$.fragment);
    			attr_dev(path, "d", "M14.53 4.53l-1.06-1.06L9 7.94 4.53 3.47 3.47 4.53 7.94 9l-4.47 4.47 1.06 1.06L9 10.06l4.47 4.47 1.06-1.06L10.06 9z");
    			add_location(path, file$6, 30, 10, 720);
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg, "width", "18");
    			attr_dev(svg, "height", "18");
    			attr_dev(svg, "viewBox", "0 0 18 18");
    			add_location(svg, file$6, 24, 8, 577);
    			attr_dev(div, "class", "modal-close absolute top-0 right-0 cursor-pointer flex flex-col items-center mt-4 mr-4 bg-white rounded-lg border-4 z-50");
    			add_location(div, file$6, 20, 6, 382);
    			set_custom_element_data(modal_section, "class", "svelte-n0txo5");
    			add_location(modal_section, file$6, 36, 6, 905);
    			set_custom_element_data(modal_wrapper, "class", "svelte-n0txo5");
    			add_location(modal_wrapper, file$6, 19, 4, 360);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, modal_wrapper, anchor);
    			append_dev(modal_wrapper, div);
    			append_dev(div, svg);
    			append_dev(svg, path);
    			append_dev(modal_wrapper, t);
    			append_dev(modal_wrapper, modal_section);

    			if (switch_instance) {
    				mount_component(switch_instance, modal_section, null);
    			}

    			current = true;

    			if (!mounted) {
    				dispose = listen_dev(div, "click", window.closeModal, false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			const switch_instance_changes = (dirty & /*props*/ 2)
    			? get_spread_update(switch_instance_spread_levels, [get_spread_object(/*props*/ ctx[1])])
    			: {};

    			if (switch_value !== (switch_value = /*current_component*/ ctx[0])) {
    				if (switch_instance) {
    					group_outros();
    					const old_component = switch_instance;

    					transition_out(old_component.$$.fragment, 1, 0, () => {
    						destroy_component(old_component, 1);
    					});

    					check_outros();
    				}

    				if (switch_value) {
    					switch_instance = new switch_value(switch_props());
    					create_component(switch_instance.$$.fragment);
    					transition_in(switch_instance.$$.fragment, 1);
    					mount_component(switch_instance, modal_section, null);
    				} else {
    					switch_instance = null;
    				}
    			} else if (switch_value) {
    				switch_instance.$set(switch_instance_changes);
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			if (switch_instance) transition_in(switch_instance.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			if (switch_instance) transition_out(switch_instance.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(modal_wrapper);
    			if (switch_instance) destroy_component(switch_instance);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$4.name,
    		type: "if",
    		source: "(19:2) {#if _show && current_component}",
    		ctx
    	});

    	return block;
    }

    // (18:0) {#key _show}
    function create_key_block(ctx) {
    	let if_block_anchor;
    	let current;
    	let if_block = /*_show*/ ctx[2] && /*current_component*/ ctx[0] && create_if_block$4(ctx);

    	const block = {
    		c: function create() {
    			if (if_block) if_block.c();
    			if_block_anchor = empty();
    		},
    		m: function mount(target, anchor) {
    			if (if_block) if_block.m(target, anchor);
    			insert_dev(target, if_block_anchor, anchor);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			if (/*_show*/ ctx[2] && /*current_component*/ ctx[0]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);

    					if (dirty & /*_show, current_component*/ 5) {
    						transition_in(if_block, 1);
    					}
    				} else {
    					if_block = create_if_block$4(ctx);
    					if_block.c();
    					transition_in(if_block, 1);
    					if_block.m(if_block_anchor.parentNode, if_block_anchor);
    				}
    			} else if (if_block) {
    				group_outros();

    				transition_out(if_block, 1, 1, () => {
    					if_block = null;
    				});

    				check_outros();
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(if_block);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(if_block);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (if_block) if_block.d(detaching);
    			if (detaching) detach_dev(if_block_anchor);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_key_block.name,
    		type: "key",
    		source: "(18:0) {#key _show}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$9(ctx) {
    	let previous_key = /*_show*/ ctx[2];
    	let key_block_anchor;
    	let current;
    	let key_block = create_key_block(ctx);

    	const block = {
    		c: function create() {
    			key_block.c();
    			key_block_anchor = empty();
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			key_block.m(target, anchor);
    			insert_dev(target, key_block_anchor, anchor);
    			current = true;
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*_show*/ 4 && safe_not_equal(previous_key, previous_key = /*_show*/ ctx[2])) {
    				group_outros();
    				transition_out(key_block, 1, 1, noop);
    				check_outros();
    				key_block = create_key_block(ctx);
    				key_block.c();
    				transition_in(key_block, 1);
    				key_block.m(key_block_anchor.parentNode, key_block_anchor);
    			} else {
    				key_block.p(ctx, dirty);
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(key_block);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(key_block);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(key_block_anchor);
    			key_block.d(detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$9.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$9($$self, $$props, $$invalidate) {
    	let _show;
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Big_modal', slots, []);
    	let current_component;
    	let props = {};

    	// fixme => export const showmodal and closemodal
    	window.showModal = (_compo, _props) => {
    		$$invalidate(0, current_component = _compo);
    		$$invalidate(1, props = _props);
    		$$invalidate(2, _show = true);
    	};

    	window.closeModal = () => {
    		$$invalidate(2, _show = false);
    	};

    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Big_modal> was created with unknown prop '${key}'`);
    	});

    	$$self.$capture_state = () => ({ current_component, props, _show });

    	$$self.$inject_state = $$props => {
    		if ('current_component' in $$props) $$invalidate(0, current_component = $$props.current_component);
    		if ('props' in $$props) $$invalidate(1, props = $$props.props);
    		if ('_show' in $$props) $$invalidate(2, _show = $$props._show);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	$$invalidate(2, _show = true);
    	return [current_component, props, _show];
    }

    class Big_modal extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$9, create_fragment$9, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Big_modal",
    			options,
    			id: create_fragment$9.name
    		});
    	}
    }

    /* entries/portal/notification/_svg_message.svelte generated by Svelte v3.48.0 */

    const file$5 = "entries/portal/notification/_svg_message.svelte";

    function create_fragment$8(ctx) {
    	let svg;
    	let path;

    	const block = {
    		c: function create() {
    			svg = svg_element("svg");
    			path = svg_element("path");
    			attr_dev(path, "d", "M4.30325 12.6667L1.33325 15V2.66667C1.33325 2.48986 1.40349 2.32029 1.52851 2.19526C1.65354 2.07024 1.82311 2 1.99992 2H13.9999C14.1767 2 14.3463 2.07024 14.4713 2.19526C14.5963 2.32029 14.6666 2.48986 14.6666 2.66667V12C14.6666 12.1768 14.5963 12.3464 14.4713 12.4714C14.3463 12.5964 14.1767 12.6667 13.9999 12.6667H4.30325ZM5.33325 6.66667V8H10.6666V6.66667H5.33325Z");
    			attr_dev(path, "fill", "#4338CA");
    			add_location(path, file$5, 7, 2, 109);
    			attr_dev(svg, "width", "16");
    			attr_dev(svg, "height", "16");
    			attr_dev(svg, "viewBox", "0 0 16 16");
    			attr_dev(svg, "fill", "none");
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			add_location(svg, file$5, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, svg, anchor);
    			append_dev(svg, path);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(svg);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$8.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$8($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Svg_message', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Svg_message> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Svg_message extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$8, create_fragment$8, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Svg_message",
    			options,
    			id: create_fragment$8.name
    		});
    	}
    }

    /* entries/portal/notification/_svg_close.svelte generated by Svelte v3.48.0 */

    const file$4 = "entries/portal/notification/_svg_close.svelte";

    function create_fragment$7(ctx) {
    	let div;
    	let svg;
    	let path0;
    	let path1;
    	let mounted;
    	let dispose;

    	const block = {
    		c: function create() {
    			div = element("div");
    			svg = svg_element("svg");
    			path0 = svg_element("path");
    			path1 = svg_element("path");
    			attr_dev(path0, "d", "M10.5 3.5L3.5 10.5");
    			attr_dev(path0, "stroke", "#4B5563");
    			attr_dev(path0, "stroke-width", "1.25");
    			attr_dev(path0, "stroke-linecap", "round");
    			attr_dev(path0, "stroke-linejoin", "round");
    			add_location(path0, file$4, 17, 4, 304);
    			attr_dev(path1, "d", "M3.5 3.5L10.5 10.5");
    			attr_dev(path1, "stroke", "#4B5563");
    			attr_dev(path1, "stroke-width", "1.25");
    			attr_dev(path1, "stroke-linecap", "round");
    			attr_dev(path1, "stroke-linejoin", "round");
    			add_location(path1, file$4, 24, 4, 458);
    			attr_dev(svg, "width", "14");
    			attr_dev(svg, "height", "14");
    			attr_dev(svg, "viewBox", "0 0 14 14");
    			attr_dev(svg, "fill", "none");
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			add_location(svg, file$4, 10, 2, 181);
    			attr_dev(div, "tabindex", "0");
    			attr_dev(div, "aria-label", "delete icon");
    			attr_dev(div, "role", "button");
    			attr_dev(div, "class", "focus:outline-none cursor-pointer");
    			add_location(div, file$4, 3, 0, 49);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, svg);
    			append_dev(svg, path0);
    			append_dev(svg, path1);

    			if (!mounted) {
    				dispose = listen_dev(
    					div,
    					"click",
    					function () {
    						if (is_function(/*onClick*/ ctx[0])) /*onClick*/ ctx[0].apply(this, arguments);
    					},
    					false,
    					false,
    					false
    				);

    				mounted = true;
    			}
    		},
    		p: function update(new_ctx, [dirty]) {
    			ctx = new_ctx;
    		},
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$7.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$7($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Svg_close', slots, []);
    	let { onClick } = $$props;
    	const writable_props = ['onClick'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Svg_close> was created with unknown prop '${key}'`);
    	});

    	$$self.$$set = $$props => {
    		if ('onClick' in $$props) $$invalidate(0, onClick = $$props.onClick);
    	};

    	$$self.$capture_state = () => ({ onClick });

    	$$self.$inject_state = $$props => {
    		if ('onClick' in $$props) $$invalidate(0, onClick = $$props.onClick);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [onClick];
    }

    class Svg_close extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$7, create_fragment$7, safe_not_equal, { onClick: 0 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Svg_close",
    			options,
    			id: create_fragment$7.name
    		});

    		const { ctx } = this.$$;
    		const props = options.props || {};

    		if (/*onClick*/ ctx[0] === undefined && !('onClick' in props)) {
    			console.warn("<Svg_close> was created without expected prop 'onClick'");
    		}
    	}

    	get onClick() {
    		throw new Error("<Svg_close>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set onClick(value) {
    		throw new Error("<Svg_close>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    // vendored => https://github.com/mikepenzin/timeago-simple/blob/master/index.js
    const time_ago = (date_time) => {
        // today date and time in milliseconds 
        var today = Date.now();
        var dateParse = Date.parse(date_time);
        //We will perform some test - if there is error, we will throw error to console and exit, no change will be on the data.
        try {
            // We need to check if we able to parse the Date (if the result is NaN, this is an issue)
            if (dateParse !== dateParse)
                throw "timeago-simple: Please check date and time format! Unable to parse the date & time: " + date_time;
        }
        catch (err) {
            console.error(err);
            return (date_time);
        }
        if ((dateParse - today) < 0) {
            return pastCalc(date_time);
        }
        else {
            return futureCalc(date_time);
        }
    };
    // General help functions for time calculations
    function pastCalc(timeData) {
        // today date and time in milliseconds 
        var today = Date.now();
        // parsing post date and time into milliseconds format
        timeData = Date.parse(timeData);
        var seconds = (today - timeData) / 1000;
        var minutes = (seconds / 60);
        var hours = (seconds / 3600);
        if (seconds < 60 && minutes < 1) {
            return (seconds === 1 ? Math.round(seconds) + " second ago" : Math.round(seconds) + " seconds ago");
        }
        if (minutes < 60 && hours < 1) {
            return (minutes === 1 ? Math.round(minutes) + " minute ago" : Math.round(minutes) + " minutes ago");
        }
        if (hours > 24) {
            var days = hours / 24;
            if (days > 30) {
                var month = days / 30;
                if (month > 12) {
                    var years = month / 12;
                    if (years > 0) {
                        return (years === 1 ? Math.ceil(years) + " year ago" : Math.ceil(years) + " years ago");
                    }
                }
                return (Math.round(month) + " month ago");
            }
            return (days === 1 ? Math.round(days) + " day ago" : Math.round(days) + " days ago");
        }
        else {
            return (hours === 1 ? Math.round(hours) + " hour ago" : Math.round(hours) + " hours ago");
        }
    }
    function futureCalc(timeData) {
        // today date and time in milliseconds 
        var today = Date.now();
        // parsing post date and time into milliseconds format
        timeData = Date.parse(timeData);
        var seconds = (timeData - today) / 1000;
        var minutes = (seconds / 60);
        var hours = (seconds / 3600);
        if (seconds < 60 && minutes < 1) {
            return (seconds === 1 ? "in " + Math.round(seconds) + " second" : "in " + Math.round(seconds) + " seconds");
        }
        if (minutes < 60 && hours < 1) {
            return (minutes === 1 ? "in " + Math.round(minutes) + " minute" : "in " + Math.round(minutes) + " minutes");
        }
        if (hours > 24) {
            var days = hours / 24;
            if (days > 30) {
                var month = days / 30;
                if (month > 12) {
                    var years = month / 12;
                    if (years > 0) {
                        return (years === 1 ? "in " + Math.ceil(years) + " year" : "in " + Math.ceil(years) + " years");
                    }
                }
                return ("in " + Math.round(month) + " month");
            }
            return (days === 1 ? "in " + Math.round(days) + " day" : "in " + Math.round(days) + " days");
        }
        else {
            return (hours === 1 ? "in " + Math.round(hours) + " hour" : "in " + Math.round(hours) + " hours");
        }
    }

    /* entries/portal/notification/_card.svelte generated by Svelte v3.48.0 */
    const file$3 = "entries/portal/notification/_card.svelte";

    // (18:2) {#if !nmsg.read}
    function create_if_block_2(ctx) {
    	let div1;
    	let div0;

    	const block = {
    		c: function create() {
    			div1 = element("div");
    			div0 = element("div");
    			attr_dev(div0, "class", "h-2 w-2 rounded-full bg-green-400");
    			add_location(div0, file$3, 19, 6, 501);
    			attr_dev(div1, "class", "absolute");
    			add_location(div1, file$3, 18, 4, 472);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div1, anchor);
    			append_dev(div1, div0);
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div1);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_2.name,
    		type: "if",
    		source: "(18:2) {#if !nmsg.read}",
    		ctx
    	});

    	return block;
    }

    // (53:2) {:else}
    function create_else_block$1(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Not implemented";
    			add_location(div, file$3, 53, 4, 1485);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_else_block$1.name,
    		type: "else",
    		source: "(53:2) {:else}",
    		ctx
    	});

    	return block;
    }

    // (33:2) {#if nmsg["type"] === "user_message"}
    function create_if_block$3(ctx) {
    	let div1;
    	let div0;
    	let p;
    	let span0;
    	let t0_value = (/*nmsg*/ ctx[2]["from_user"] || "") + "";
    	let t0;
    	let t1;
    	let span1;
    	let t2_value = (/*nmsg*/ ctx[2]["contents"] || "") + "";
    	let t2;
    	let t3;
    	let svgclose;
    	let t4;
    	let current;

    	svgclose = new Svg_close({
    			props: { onClick: /*func*/ ctx[3] },
    			$$inline: true
    		});

    	let if_block = /*nmsg*/ ctx[2]["created_at"] && create_if_block_1(ctx);

    	const block = {
    		c: function create() {
    			div1 = element("div");
    			div0 = element("div");
    			p = element("p");
    			span0 = element("span");
    			t0 = text(t0_value);
    			t1 = text("\n          messaged you:\n          ");
    			span1 = element("span");
    			t2 = text(t2_value);
    			t3 = space();
    			create_component(svgclose.$$.fragment);
    			t4 = space();
    			if (if_block) if_block.c();
    			attr_dev(span0, "class", "text-indigo-700");
    			add_location(span0, file$3, 36, 10, 1007);
    			attr_dev(span1, "class", "italic");
    			add_location(span1, file$3, 38, 10, 1104);
    			attr_dev(p, "tabindex", "0");
    			attr_dev(p, "class", "focus:outline-none text-sm leading-none");
    			add_location(p, file$3, 35, 8, 932);
    			attr_dev(div0, "class", "flex items-center justify-between w-full");
    			add_location(div0, file$3, 34, 6, 869);
    			attr_dev(div1, "class", "pl-3 w-full");
    			add_location(div1, file$3, 33, 4, 837);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div1, anchor);
    			append_dev(div1, div0);
    			append_dev(div0, p);
    			append_dev(p, span0);
    			append_dev(span0, t0);
    			append_dev(p, t1);
    			append_dev(p, span1);
    			append_dev(span1, t2);
    			append_dev(div0, t3);
    			mount_component(svgclose, div0, null);
    			append_dev(div1, t4);
    			if (if_block) if_block.m(div1, null);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			if ((!current || dirty & /*nmsg*/ 4) && t0_value !== (t0_value = (/*nmsg*/ ctx[2]["from_user"] || "") + "")) set_data_dev(t0, t0_value);
    			if ((!current || dirty & /*nmsg*/ 4) && t2_value !== (t2_value = (/*nmsg*/ ctx[2]["contents"] || "") + "")) set_data_dev(t2, t2_value);
    			const svgclose_changes = {};
    			if (dirty & /*delete_notif, nmsg*/ 5) svgclose_changes.onClick = /*func*/ ctx[3];
    			svgclose.$set(svgclose_changes);

    			if (/*nmsg*/ ctx[2]["created_at"]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);
    				} else {
    					if_block = create_if_block_1(ctx);
    					if_block.c();
    					if_block.m(div1, null);
    				}
    			} else if (if_block) {
    				if_block.d(1);
    				if_block = null;
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(svgclose.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(svgclose.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div1);
    			destroy_component(svgclose);
    			if (if_block) if_block.d();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$3.name,
    		type: "if",
    		source: "(33:2) {#if nmsg[\\\"type\\\"] === \\\"user_message\\\"}",
    		ctx
    	});

    	return block;
    }

    // (47:6) {#if nmsg["created_at"]}
    function create_if_block_1(ctx) {
    	let p;
    	let t_value = time_ago(/*nmsg*/ ctx[2]["created_at"]) + "";
    	let t;

    	const block = {
    		c: function create() {
    			p = element("p");
    			t = text(t_value);
    			attr_dev(p, "class", "focus:outline-none text-xs leading-3 pt-1 text-gray-500");
    			add_location(p, file$3, 47, 8, 1326);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, p, anchor);
    			append_dev(p, t);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*nmsg*/ 4 && t_value !== (t_value = time_ago(/*nmsg*/ ctx[2]["created_at"]) + "")) set_data_dev(t, t_value);
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(p);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_1.name,
    		type: "if",
    		source: "(47:6) {#if nmsg[\\\"created_at\\\"]}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$6(ctx) {
    	let div1;
    	let t0;
    	let div0;
    	let svgmessage;
    	let t1;
    	let current_block_type_index;
    	let if_block1;
    	let current;
    	let mounted;
    	let dispose;
    	let if_block0 = !/*nmsg*/ ctx[2].read && create_if_block_2(ctx);
    	svgmessage = new Svg_message({ $$inline: true });
    	const if_block_creators = [create_if_block$3, create_else_block$1];
    	const if_blocks = [];

    	function select_block_type(ctx, dirty) {
    		if (/*nmsg*/ ctx[2]["type"] === "user_message") return 0;
    		return 1;
    	}

    	current_block_type_index = select_block_type(ctx);
    	if_block1 = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx);

    	const block = {
    		c: function create() {
    			div1 = element("div");
    			if (if_block0) if_block0.c();
    			t0 = space();
    			div0 = element("div");
    			create_component(svgmessage.$$.fragment);
    			t1 = space();
    			if_block1.c();
    			attr_dev(div0, "tabindex", "0");
    			attr_dev(div0, "aria-label", "group icon");
    			attr_dev(div0, "role", "img");
    			attr_dev(div0, "class", "focus:outline-none w-8 h-8 border rounded-full border-gray-200 flex flex-shrink-0 items-center justify-center");
    			add_location(div0, file$3, 23, 2, 573);
    			attr_dev(div1, "class", "w-full p-3 mt-4 bg-white rounded shadow flex flex-shrink-0 cursor-pointer relative");
    			add_location(div1, file$3, 8, 0, 242);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div1, anchor);
    			if (if_block0) if_block0.m(div1, null);
    			append_dev(div1, t0);
    			append_dev(div1, div0);
    			mount_component(svgmessage, div0, null);
    			append_dev(div1, t1);
    			if_blocks[current_block_type_index].m(div1, null);
    			current = true;

    			if (!mounted) {
    				dispose = listen_dev(div1, "click", stop_propagation(/*click_handler*/ ctx[4]), false, false, true);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			if (!/*nmsg*/ ctx[2].read) {
    				if (if_block0) ; else {
    					if_block0 = create_if_block_2(ctx);
    					if_block0.c();
    					if_block0.m(div1, t0);
    				}
    			} else if (if_block0) {
    				if_block0.d(1);
    				if_block0 = null;
    			}

    			let previous_block_index = current_block_type_index;
    			current_block_type_index = select_block_type(ctx);

    			if (current_block_type_index === previous_block_index) {
    				if_blocks[current_block_type_index].p(ctx, dirty);
    			} else {
    				group_outros();

    				transition_out(if_blocks[previous_block_index], 1, 1, () => {
    					if_blocks[previous_block_index] = null;
    				});

    				check_outros();
    				if_block1 = if_blocks[current_block_type_index];

    				if (!if_block1) {
    					if_block1 = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx);
    					if_block1.c();
    				} else {
    					if_block1.p(ctx, dirty);
    				}

    				transition_in(if_block1, 1);
    				if_block1.m(div1, null);
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(svgmessage.$$.fragment, local);
    			transition_in(if_block1);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(svgmessage.$$.fragment, local);
    			transition_out(if_block1);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div1);
    			if (if_block0) if_block0.d();
    			destroy_component(svgmessage);
    			if_blocks[current_block_type_index].d();
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$6.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$6($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Card', slots, []);
    	let { delete_notif } = $$props;
    	let { read_notif } = $$props;
    	let { nmsg } = $$props;
    	const writable_props = ['delete_notif', 'read_notif', 'nmsg'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Card> was created with unknown prop '${key}'`);
    	});

    	const func = () => {
    		delete_notif(nmsg.id);
    	};

    	const click_handler = () => {
    		if (nmsg.read) {
    			return;
    		}

    		read_notif(nmsg.id);
    	};

    	$$self.$$set = $$props => {
    		if ('delete_notif' in $$props) $$invalidate(0, delete_notif = $$props.delete_notif);
    		if ('read_notif' in $$props) $$invalidate(1, read_notif = $$props.read_notif);
    		if ('nmsg' in $$props) $$invalidate(2, nmsg = $$props.nmsg);
    	};

    	$$self.$capture_state = () => ({
    		SvgMessage: Svg_message,
    		SvgClose: Svg_close,
    		time_ago,
    		delete_notif,
    		read_notif,
    		nmsg
    	});

    	$$self.$inject_state = $$props => {
    		if ('delete_notif' in $$props) $$invalidate(0, delete_notif = $$props.delete_notif);
    		if ('read_notif' in $$props) $$invalidate(1, read_notif = $$props.read_notif);
    		if ('nmsg' in $$props) $$invalidate(2, nmsg = $$props.nmsg);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [delete_notif, read_notif, nmsg, func, click_handler];
    }

    class Card extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$6, create_fragment$6, safe_not_equal, { delete_notif: 0, read_notif: 1, nmsg: 2 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Card",
    			options,
    			id: create_fragment$6.name
    		});

    		const { ctx } = this.$$;
    		const props = options.props || {};

    		if (/*delete_notif*/ ctx[0] === undefined && !('delete_notif' in props)) {
    			console.warn("<Card> was created without expected prop 'delete_notif'");
    		}

    		if (/*read_notif*/ ctx[1] === undefined && !('read_notif' in props)) {
    			console.warn("<Card> was created without expected prop 'read_notif'");
    		}

    		if (/*nmsg*/ ctx[2] === undefined && !('nmsg' in props)) {
    			console.warn("<Card> was created without expected prop 'nmsg'");
    		}
    	}

    	get delete_notif() {
    		throw new Error("<Card>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set delete_notif(value) {
    		throw new Error("<Card>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get read_notif() {
    		throw new Error("<Card>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set read_notif(value) {
    		throw new Error("<Card>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get nmsg() {
    		throw new Error("<Card>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set nmsg(value) {
    		throw new Error("<Card>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/portal/notification/processing.svelte generated by Svelte v3.48.0 */

    const file$2 = "entries/portal/notification/processing.svelte";

    function create_fragment$5(ctx) {
    	let div1;
    	let svg;
    	let path;
    	let t0;
    	let div0;
    	let t1;

    	const block = {
    		c: function create() {
    			div1 = element("div");
    			svg = svg_element("svg");
    			path = svg_element("path");
    			t0 = space();
    			div0 = element("div");
    			t1 = text(/*text*/ ctx[0]);
    			attr_dev(path, "clip-rule", "evenodd");
    			attr_dev(path, "d", "M15.165 8.53a.5.5 0 01-.404.58A7 7 0 1023 16a.5.5 0 011 0 8 8 0 11-9.416-7.874.5.5 0 01.58.404z");
    			attr_dev(path, "fill", "currentColor");
    			attr_dev(path, "fill-rule", "evenodd");
    			add_location(path, file$2, 12, 4, 303);
    			attr_dev(svg, "fill", "none");
    			attr_dev(svg, "class", "w-20 h-20 animate-spin");
    			attr_dev(svg, "viewBox", "0 0 32 32");
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			add_location(svg, file$2, 6, 2, 176);
    			add_location(div0, file$2, 19, 2, 511);
    			attr_dev(div1, "class", "flex justify-center items-center space-x-1 text-lg text-gray-700");
    			add_location(div1, file$2, 5, 0, 95);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div1, anchor);
    			append_dev(div1, svg);
    			append_dev(svg, path);
    			append_dev(div1, t0);
    			append_dev(div1, div0);
    			append_dev(div0, t1);
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*text*/ 1) set_data_dev(t1, /*text*/ ctx[0]);
    		},
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div1);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$5.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$5($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Processing', slots, []);
    	let { text = "Processing" } = $$props;
    	const writable_props = ['text'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Processing> was created with unknown prop '${key}'`);
    	});

    	$$self.$$set = $$props => {
    		if ('text' in $$props) $$invalidate(0, text = $$props.text);
    	};

    	$$self.$capture_state = () => ({ text });

    	$$self.$inject_state = $$props => {
    		if ('text' in $$props) $$invalidate(0, text = $$props.text);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [text];
    }

    class Processing extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$5, create_fragment$5, safe_not_equal, { text: 0 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Processing",
    			options,
    			id: create_fragment$5.name
    		});
    	}

    	get text() {
    		throw new Error("<Processing>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set text(value) {
    		throw new Error("<Processing>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/portal/notification/notification.svelte generated by Svelte v3.48.0 */
    const file$1 = "entries/portal/notification/notification.svelte";

    function get_each_context(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[8] = list[i];
    	return child_ctx;
    }

    // (94:2) {:else}
    function create_else_block(ctx) {
    	let each_1_anchor;
    	let current;
    	let each_value = /*_messages*/ ctx[1];
    	validate_each_argument(each_value);
    	let each_blocks = [];

    	for (let i = 0; i < each_value.length; i += 1) {
    		each_blocks[i] = create_each_block(get_each_context(ctx, each_value, i));
    	}

    	const out = i => transition_out(each_blocks[i], 1, 1, () => {
    		each_blocks[i] = null;
    	});

    	const block = {
    		c: function create() {
    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].c();
    			}

    			each_1_anchor = empty();
    		},
    		m: function mount(target, anchor) {
    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].m(target, anchor);
    			}

    			insert_dev(target, each_1_anchor, anchor);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*_messages, ndelete, nread*/ 50) {
    				each_value = /*_messages*/ ctx[1];
    				validate_each_argument(each_value);
    				let i;

    				for (i = 0; i < each_value.length; i += 1) {
    					const child_ctx = get_each_context(ctx, each_value, i);

    					if (each_blocks[i]) {
    						each_blocks[i].p(child_ctx, dirty);
    						transition_in(each_blocks[i], 1);
    					} else {
    						each_blocks[i] = create_each_block(child_ctx);
    						each_blocks[i].c();
    						transition_in(each_blocks[i], 1);
    						each_blocks[i].m(each_1_anchor.parentNode, each_1_anchor);
    					}
    				}

    				group_outros();

    				for (i = each_value.length; i < each_blocks.length; i += 1) {
    					out(i);
    				}

    				check_outros();
    			}
    		},
    		i: function intro(local) {
    			if (current) return;

    			for (let i = 0; i < each_value.length; i += 1) {
    				transition_in(each_blocks[i]);
    			}

    			current = true;
    		},
    		o: function outro(local) {
    			each_blocks = each_blocks.filter(Boolean);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				transition_out(each_blocks[i]);
    			}

    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_each(each_blocks, detaching);
    			if (detaching) detach_dev(each_1_anchor);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_else_block.name,
    		type: "else",
    		source: "(94:2) {:else}",
    		ctx
    	});

    	return block;
    }

    // (92:2) {#if loading}
    function create_if_block$2(ctx) {
    	let processing;
    	let current;

    	processing = new Processing({
    			props: { text: "loading" },
    			$$inline: true
    		});

    	const block = {
    		c: function create() {
    			create_component(processing.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(processing, target, anchor);
    			current = true;
    		},
    		p: noop,
    		i: function intro(local) {
    			if (current) return;
    			transition_in(processing.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(processing.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(processing, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$2.name,
    		type: "if",
    		source: "(92:2) {#if loading}",
    		ctx
    	});

    	return block;
    }

    // (95:4) {#each _messages as nmsg}
    function create_each_block(ctx) {
    	let card;
    	let current;

    	card = new Card({
    			props: {
    				nmsg: /*nmsg*/ ctx[8],
    				delete_notif: /*ndelete*/ ctx[5],
    				read_notif: /*nread*/ ctx[4]
    			},
    			$$inline: true
    		});

    	const block = {
    		c: function create() {
    			create_component(card.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(card, target, anchor);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			const card_changes = {};
    			if (dirty & /*_messages*/ 2) card_changes.nmsg = /*nmsg*/ ctx[8];
    			card.$set(card_changes);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(card.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(card.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(card, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block.name,
    		type: "each",
    		source: "(95:4) {#each _messages as nmsg}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$4(ctx) {
    	let div3;
    	let div1;
    	let p0;
    	let t1;
    	let div0;
    	let button0;
    	let svg0;
    	let path0;
    	let path1;
    	let t2;
    	let button1;
    	let svg1;
    	let path2;
    	let t3;
    	let button2;
    	let svg2;
    	let path3;
    	let path4;
    	let t4;
    	let current_block_type_index;
    	let if_block;
    	let t5;
    	let div2;
    	let hr0;
    	let t6;
    	let p1;
    	let t8;
    	let hr1;
    	let current;
    	let mounted;
    	let dispose;
    	const if_block_creators = [create_if_block$2, create_else_block];
    	const if_blocks = [];

    	function select_block_type(ctx, dirty) {
    		if (/*loading*/ ctx[0]) return 0;
    		return 1;
    	}

    	current_block_type_index = select_block_type(ctx);
    	if_block = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx);

    	const block = {
    		c: function create() {
    			div3 = element("div");
    			div1 = element("div");
    			p0 = element("p");
    			p0.textContent = "Notifications";
    			t1 = space();
    			div0 = element("div");
    			button0 = element("button");
    			svg0 = svg_element("svg");
    			path0 = svg_element("path");
    			path1 = svg_element("path");
    			t2 = space();
    			button1 = element("button");
    			svg1 = svg_element("svg");
    			path2 = svg_element("path");
    			t3 = space();
    			button2 = element("button");
    			svg2 = svg_element("svg");
    			path3 = svg_element("path");
    			path4 = svg_element("path");
    			t4 = space();
    			if_block.c();
    			t5 = space();
    			div2 = element("div");
    			hr0 = element("hr");
    			t6 = space();
    			p1 = element("p");
    			p1.textContent = "Existance is lie :|";
    			t8 = space();
    			hr1 = element("hr");
    			attr_dev(p0, "tabindex", "0");
    			attr_dev(p0, "class", "focus:outline-none text-2xl font-semibold leading-6 text-gray-800");
    			add_location(p0, file$1, 15, 4, 636);
    			attr_dev(path0, "d", "M10 12a2 2 0 100-4 2 2 0 000 4z");
    			add_location(path0, file$1, 33, 10, 1144);
    			attr_dev(path1, "fill-rule", "evenodd");
    			attr_dev(path1, "d", "M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z");
    			attr_dev(path1, "clip-rule", "evenodd");
    			add_location(path1, file$1, 34, 10, 1199);
    			attr_dev(svg0, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg0, "class", "h-6 w-6 text-gray-400");
    			attr_dev(svg0, "viewBox", "0 0 20 20");
    			attr_dev(svg0, "fill", "currentColor");
    			add_location(svg0, file$1, 27, 8, 974);
    			attr_dev(button0, "aria-label", "mark all message as read");
    			attr_dev(button0, "class", "focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 rounded-md cursor-pointer");
    			add_location(button0, file$1, 23, 6, 790);
    			attr_dev(path2, "fill-rule", "evenodd");
    			attr_dev(path2, "d", "M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z");
    			attr_dev(path2, "clip-rule", "evenodd");
    			add_location(path2, file$1, 53, 10, 1841);
    			attr_dev(svg1, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg1, "class", "h-6 w-6 text-gray-400");
    			attr_dev(svg1, "viewBox", "0 0 20 20");
    			attr_dev(svg1, "fill", "currentColor");
    			add_location(svg1, file$1, 47, 8, 1671);
    			attr_dev(button1, "aria-label", "refresh");
    			attr_dev(button1, "class", "focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 rounded-md cursor-pointer mx-5");
    			add_location(button1, file$1, 42, 6, 1472);
    			attr_dev(path3, "d", "M18 6L6 18");
    			attr_dev(path3, "stroke", "#4B5563");
    			attr_dev(path3, "stroke-width", "1.25");
    			attr_dev(path3, "stroke-linecap", "round");
    			attr_dev(path3, "stroke-linejoin", "round");
    			add_location(path3, file$1, 72, 10, 2636);
    			attr_dev(path4, "d", "M6 6L18 18");
    			attr_dev(path4, "stroke", "#4B5563");
    			attr_dev(path4, "stroke-width", "1.25");
    			attr_dev(path4, "stroke-linecap", "round");
    			attr_dev(path4, "stroke-linejoin", "round");
    			add_location(path4, file$1, 79, 10, 2824);
    			attr_dev(svg2, "class", "h-6 w-6");
    			attr_dev(svg2, "viewBox", "0 0 24 24");
    			attr_dev(svg2, "fill", "none");
    			attr_dev(svg2, "xmlns", "http://www.w3.org/2000/svg");
    			add_location(svg2, file$1, 66, 8, 2488);
    			attr_dev(button2, "aria-label", "close notification modal");
    			attr_dev(button2, "class", "focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 rounded-md cursor-pointer");
    			add_location(button2, file$1, 61, 6, 2271);
    			add_location(div0, file$1, 22, 4, 778);
    			attr_dev(div1, "class", "flex items-center justify-between");
    			add_location(div1, file$1, 14, 2, 584);
    			attr_dev(hr0, "class", "w-full");
    			add_location(hr0, file$1, 100, 4, 3285);
    			attr_dev(p1, "tabindex", "0");
    			attr_dev(p1, "class", "focus:outline-none text-sm flex flex-shrink-0 leading-normal px-3 py-16 text-gray-500");
    			add_location(p1, file$1, 101, 4, 3311);
    			attr_dev(hr1, "class", "w-full");
    			add_location(hr1, file$1, 107, 4, 3478);
    			attr_dev(div2, "class", "flex items-center justiyf-between");
    			add_location(div2, file$1, 99, 2, 3233);
    			attr_dev(div3, "class", "2xl:w-4/12 bg-gray-50 h-screen overflow-y-auto p-8 absolute");
    			add_location(div3, file$1, 13, 0, 508);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div3, anchor);
    			append_dev(div3, div1);
    			append_dev(div1, p0);
    			append_dev(div1, t1);
    			append_dev(div1, div0);
    			append_dev(div0, button0);
    			append_dev(button0, svg0);
    			append_dev(svg0, path0);
    			append_dev(svg0, path1);
    			append_dev(div0, t2);
    			append_dev(div0, button1);
    			append_dev(button1, svg1);
    			append_dev(svg1, path2);
    			append_dev(div0, t3);
    			append_dev(div0, button2);
    			append_dev(button2, svg2);
    			append_dev(svg2, path3);
    			append_dev(svg2, path4);
    			append_dev(div3, t4);
    			if_blocks[current_block_type_index].m(div3, null);
    			append_dev(div3, t5);
    			append_dev(div3, div2);
    			append_dev(div2, hr0);
    			append_dev(div2, t6);
    			append_dev(div2, p1);
    			append_dev(div2, t8);
    			append_dev(div2, hr1);
    			current = true;

    			if (!mounted) {
    				dispose = [
    					listen_dev(button1, "click", /*refresh*/ ctx[2], false, false, false),
    					listen_dev(button2, "click", /*toggle_npanel*/ ctx[3], false, false, false)
    				];

    				mounted = true;
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			let previous_block_index = current_block_type_index;
    			current_block_type_index = select_block_type(ctx);

    			if (current_block_type_index === previous_block_index) {
    				if_blocks[current_block_type_index].p(ctx, dirty);
    			} else {
    				group_outros();

    				transition_out(if_blocks[previous_block_index], 1, 1, () => {
    					if_blocks[previous_block_index] = null;
    				});

    				check_outros();
    				if_block = if_blocks[current_block_type_index];

    				if (!if_block) {
    					if_block = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx);
    					if_block.c();
    				} else {
    					if_block.p(ctx, dirty);
    				}

    				transition_in(if_block, 1);
    				if_block.m(div3, t5);
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(if_block);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(if_block);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div3);
    			if_blocks[current_block_type_index].d();
    			mounted = false;
    			run_all(dispose);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$4.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$4($$self, $$props, $$invalidate) {
    	let _messages;
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Notification', slots, []);
    	let { loading = true } = $$props;
    	let { messages = [] } = $$props;
    	const dispatch = createEventDispatcher();
    	const refresh = () => dispatch("refresh");
    	const toggle_npanel = () => dispatch("toggle_npanel");
    	const nread = msg => dispatch("nread", msg);
    	const ndelete = msg => dispatch("ndelete", msg);
    	const writable_props = ['loading', 'messages'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Notification> was created with unknown prop '${key}'`);
    	});

    	$$self.$$set = $$props => {
    		if ('loading' in $$props) $$invalidate(0, loading = $$props.loading);
    		if ('messages' in $$props) $$invalidate(6, messages = $$props.messages);
    	};

    	$$self.$capture_state = () => ({
    		Card,
    		Processing,
    		createEventDispatcher,
    		loading,
    		messages,
    		dispatch,
    		refresh,
    		toggle_npanel,
    		nread,
    		ndelete,
    		_messages
    	});

    	$$self.$inject_state = $$props => {
    		if ('loading' in $$props) $$invalidate(0, loading = $$props.loading);
    		if ('messages' in $$props) $$invalidate(6, messages = $$props.messages);
    		if ('_messages' in $$props) $$invalidate(1, _messages = $$props._messages);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	$$self.$$.update = () => {
    		if ($$self.$$.dirty & /*messages*/ 64) {
    			$$invalidate(1, _messages = messages.sort((x, y) => x["id"] - y["id"]));
    		}
    	};

    	return [loading, _messages, refresh, toggle_npanel, nread, ndelete, messages];
    }

    class Notification extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$4, create_fragment$4, safe_not_equal, { loading: 0, messages: 6 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Notification",
    			options,
    			id: create_fragment$4.name
    		});
    	}

    	get loading() {
    		throw new Error("<Notification>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set loading(value) {
    		throw new Error("<Notification>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get messages() {
    		throw new Error("<Notification>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set messages(value) {
    		throw new Error("<Notification>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/portal/notification/modal.svelte generated by Svelte v3.48.0 */
    const file = "entries/portal/notification/modal.svelte";

    // (7:0) {#if show}
    function create_if_block$1(ctx) {
    	let div1;
    	let div0;
    	let inner;
    	let current;

    	inner = new Notification({
    			props: {
    				loading: /*loading*/ ctx[1],
    				messages: /*messages*/ ctx[2]
    			},
    			$$inline: true
    		});

    	inner.$on("ndelete", /*ndelete_handler*/ ctx[3]);
    	inner.$on("nread", /*nread_handler*/ ctx[4]);
    	inner.$on("refresh", /*refresh_handler*/ ctx[5]);
    	inner.$on("toggle_npanel", /*toggle_npanel_handler*/ ctx[6]);

    	const block = {
    		c: function create() {
    			div1 = element("div");
    			div0 = element("div");
    			create_component(inner.$$.fragment);
    			attr_dev(div0, "class", "w-full absolute h-full overflow-x-hidden transform translate-x-0 transition ease-in-out duration-700 ");
    			add_location(div0, file, 10, 4, 296);
    			attr_dev(div1, "class", "w-full h-full bg-gray-800 bg-opacity-90 top-0 overflow-y-auto overflow-x-hidden fixed sticky-0 z-50 left-14");
    			add_location(div1, file, 7, 2, 163);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div1, anchor);
    			append_dev(div1, div0);
    			mount_component(inner, div0, null);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			const inner_changes = {};
    			if (dirty & /*loading*/ 2) inner_changes.loading = /*loading*/ ctx[1];
    			if (dirty & /*messages*/ 4) inner_changes.messages = /*messages*/ ctx[2];
    			inner.$set(inner_changes);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(inner.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(inner.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div1);
    			destroy_component(inner);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$1.name,
    		type: "if",
    		source: "(7:0) {#if show}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$3(ctx) {
    	let if_block_anchor;
    	let current;
    	let if_block = /*show*/ ctx[0] && create_if_block$1(ctx);

    	const block = {
    		c: function create() {
    			if (if_block) if_block.c();
    			if_block_anchor = empty();
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			if (if_block) if_block.m(target, anchor);
    			insert_dev(target, if_block_anchor, anchor);
    			current = true;
    		},
    		p: function update(ctx, [dirty]) {
    			if (/*show*/ ctx[0]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);

    					if (dirty & /*show*/ 1) {
    						transition_in(if_block, 1);
    					}
    				} else {
    					if_block = create_if_block$1(ctx);
    					if_block.c();
    					transition_in(if_block, 1);
    					if_block.m(if_block_anchor.parentNode, if_block_anchor);
    				}
    			} else if (if_block) {
    				group_outros();

    				transition_out(if_block, 1, 1, () => {
    					if_block = null;
    				});

    				check_outros();
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(if_block);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(if_block);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (if_block) if_block.d(detaching);
    			if (detaching) detach_dev(if_block_anchor);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$3.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$3($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Modal', slots, []);
    	let { show = false } = $$props;
    	let { loading = true } = $$props;
    	let { messages = [] } = $$props;
    	const writable_props = ['show', 'loading', 'messages'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Modal> was created with unknown prop '${key}'`);
    	});

    	function ndelete_handler(event) {
    		bubble.call(this, $$self, event);
    	}

    	function nread_handler(event) {
    		bubble.call(this, $$self, event);
    	}

    	function refresh_handler(event) {
    		bubble.call(this, $$self, event);
    	}

    	function toggle_npanel_handler(event) {
    		bubble.call(this, $$self, event);
    	}

    	$$self.$$set = $$props => {
    		if ('show' in $$props) $$invalidate(0, show = $$props.show);
    		if ('loading' in $$props) $$invalidate(1, loading = $$props.loading);
    		if ('messages' in $$props) $$invalidate(2, messages = $$props.messages);
    	};

    	$$self.$capture_state = () => ({ Inner: Notification, show, loading, messages });

    	$$self.$inject_state = $$props => {
    		if ('show' in $$props) $$invalidate(0, show = $$props.show);
    		if ('loading' in $$props) $$invalidate(1, loading = $$props.loading);
    		if ('messages' in $$props) $$invalidate(2, messages = $$props.messages);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [
    		show,
    		loading,
    		messages,
    		ndelete_handler,
    		nread_handler,
    		refresh_handler,
    		toggle_npanel_handler
    	];
    }

    class Modal extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$3, create_fragment$3, safe_not_equal, { show: 0, loading: 1, messages: 2 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Modal",
    			options,
    			id: create_fragment$3.name
    		});
    	}

    	get show() {
    		throw new Error("<Modal>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set show(value) {
    		throw new Error("<Modal>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get loading() {
    		throw new Error("<Modal>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set loading(value) {
    		throw new Error("<Modal>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get messages() {
    		throw new Error("<Modal>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set messages(value) {
    		throw new Error("<Modal>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/portal/portal.svelte generated by Svelte v3.48.0 */

    const { console: console_1 } = globals;

    // (13:0) <MainLayout pending_notification={true} on:ntoggle={(ev) => console.log(ev)}>
    function create_default_slot(ctx) {
    	let router;
    	let current;
    	router = new Router({ $$inline: true });

    	const block = {
    		c: function create() {
    			create_component(router.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(router, target, anchor);
    			current = true;
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(router.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(router.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(router, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_default_slot.name,
    		type: "slot",
    		source: "(13:0) <MainLayout pending_notification={true} on:ntoggle={(ev) => console.log(ev)}>",
    		ctx
    	});

    	return block;
    }

    function create_fragment$2(ctx) {
    	let bigmodal;
    	let t0;
    	let notificationmodal;
    	let t1;
    	let mainlayout;
    	let current;
    	bigmodal = new Big_modal({ $$inline: true });
    	notificationmodal = new Modal({ $$inline: true });

    	mainlayout = new Main({
    			props: {
    				pending_notification: true,
    				$$slots: { default: [create_default_slot] },
    				$$scope: { ctx }
    			},
    			$$inline: true
    		});

    	mainlayout.$on("ntoggle", /*ntoggle_handler*/ ctx[0]);

    	const block = {
    		c: function create() {
    			create_component(bigmodal.$$.fragment);
    			t0 = space();
    			create_component(notificationmodal.$$.fragment);
    			t1 = space();
    			create_component(mainlayout.$$.fragment);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			mount_component(bigmodal, target, anchor);
    			insert_dev(target, t0, anchor);
    			mount_component(notificationmodal, target, anchor);
    			insert_dev(target, t1, anchor);
    			mount_component(mainlayout, target, anchor);
    			current = true;
    		},
    		p: function update(ctx, [dirty]) {
    			const mainlayout_changes = {};

    			if (dirty & /*$$scope*/ 4) {
    				mainlayout_changes.$$scope = { dirty, ctx };
    			}

    			mainlayout.$set(mainlayout_changes);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(bigmodal.$$.fragment, local);
    			transition_in(notificationmodal.$$.fragment, local);
    			transition_in(mainlayout.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(bigmodal.$$.fragment, local);
    			transition_out(notificationmodal.$$.fragment, local);
    			transition_out(mainlayout.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(bigmodal, detaching);
    			if (detaching) detach_dev(t0);
    			destroy_component(notificationmodal, detaching);
    			if (detaching) detach_dev(t1);
    			destroy_component(mainlayout, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$2.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$2($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Portal', slots, []);
    	routes.set(page_routes);
    	let app;
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console_1.warn(`<Portal> was created with unknown prop '${key}'`);
    	});

    	const ntoggle_handler = ev => console.log(ev);

    	$$self.$capture_state = () => ({
    		routes,
    		Router,
    		page_routes,
    		MainLayout: Main,
    		BigModal: Big_modal,
    		NotificationModal: Modal,
    		app
    	});

    	$$self.$inject_state = $$props => {
    		if ('app' in $$props) app = $$props.app;
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [ntoggle_handler];
    }

    class Portal extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$2, create_fragment$2, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Portal",
    			options,
    			id: create_fragment$2.name
    		});
    	}
    }

    /* entries/portal/index.svelte generated by Svelte v3.48.0 */

    // (11:0) {#if loading}
    function create_if_block(ctx) {
    	let portal;
    	let current;
    	portal = new Portal({ $$inline: true });

    	const block = {
    		c: function create() {
    			create_component(portal.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(portal, target, anchor);
    			current = true;
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(portal.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(portal.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(portal, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block.name,
    		type: "if",
    		source: "(11:0) {#if loading}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$1(ctx) {
    	let t;
    	let tailwind;
    	let current;
    	let if_block = /*loading*/ ctx[0] && create_if_block(ctx);
    	tailwind = new Tailwind({ $$inline: true });

    	const block = {
    		c: function create() {
    			if (if_block) if_block.c();
    			t = space();
    			create_component(tailwind.$$.fragment);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			if (if_block) if_block.m(target, anchor);
    			insert_dev(target, t, anchor);
    			mount_component(tailwind, target, anchor);
    			current = true;
    		},
    		p: noop,
    		i: function intro(local) {
    			if (current) return;
    			transition_in(if_block);
    			transition_in(tailwind.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(if_block);
    			transition_out(tailwind.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (if_block) if_block.d(detaching);
    			if (detaching) detach_dev(t);
    			destroy_component(tailwind, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$1.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$1($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Portal', slots, []);
    	let loading = true;
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Portal> was created with unknown prop '${key}'`);
    	});

    	$$self.$capture_state = () => ({ Tailwind, Portal, loading });

    	$$self.$inject_state = $$props => {
    		if ('loading' in $$props) $$invalidate(0, loading = $$props.loading);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [loading];
    }

    class Portal_1 extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$1, create_fragment$1, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Portal_1",
    			options,
    			id: create_fragment$1.name
    		});
    	}
    }

    /* entries/playground/index.svelte generated by Svelte v3.48.0 */

    function create_fragment(ctx) {
    	let portal;
    	let current;
    	portal = new Portal_1({ $$inline: true });

    	const block = {
    		c: function create() {
    			create_component(portal.$$.fragment);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			mount_component(portal, target, anchor);
    			current = true;
    		},
    		p: noop,
    		i: function intro(local) {
    			if (current) return;
    			transition_in(portal.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(portal.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(portal, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Playground', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Playground> was created with unknown prop '${key}'`);
    	});

    	$$self.$capture_state = () => ({ Portal: Portal_1 });
    	return [];
    }

    class Playground extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance, create_fragment, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Playground",
    			options,
    			id: create_fragment.name
    		});
    	}
    }

    const __svelte_app__ = new Playground({
        target: document.body,
        props: {},
    });

    return __svelte_app__;

})();
//# sourceMappingURL=playground.js.map

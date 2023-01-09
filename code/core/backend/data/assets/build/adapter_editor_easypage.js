(function () {
    'use strict';

    function noop() { }
    const identity = x => x;
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

    const is_client = typeof window !== 'undefined';
    let now = is_client
        ? () => window.performance.now()
        : () => Date.now();
    let raf = is_client ? cb => requestAnimationFrame(cb) : noop;

    const tasks = new Set();
    function run_tasks(now) {
        tasks.forEach(task => {
            if (!task.c(now)) {
                tasks.delete(task);
                task.f();
            }
        });
        if (tasks.size !== 0)
            raf(run_tasks);
    }
    /**
     * Creates a new task that runs on each raf frame
     * until it returns a falsy value or is aborted
     */
    function loop(callback) {
        let task;
        if (tasks.size === 0)
            raf(run_tasks);
        return {
            promise: new Promise(fulfill => {
                tasks.add(task = { c: callback, f: fulfill });
            }),
            abort() {
                tasks.delete(task);
            }
        };
    }
    function append(target, node) {
        target.appendChild(node);
    }
    function get_root_for_style(node) {
        if (!node)
            return document;
        const root = node.getRootNode ? node.getRootNode() : node.ownerDocument;
        if (root && root.host) {
            return root;
        }
        return node.ownerDocument;
    }
    function append_empty_stylesheet(node) {
        const style_element = element('style');
        append_stylesheet(get_root_for_style(node), style_element);
        return style_element.sheet;
    }
    function append_stylesheet(node, style) {
        append(node.head || node, style);
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
    function attr(node, attribute, value) {
        if (value == null)
            node.removeAttribute(attribute);
        else if (node.getAttribute(attribute) !== value)
            node.setAttribute(attribute, value);
    }
    function set_svg_attributes(node, attributes) {
        for (const key in attributes) {
            attr(node, key, attributes[key]);
        }
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
    function set_input_value(input, value) {
        input.value = value == null ? '' : value;
    }
    function set_style(node, key, value, important) {
        if (value === null) {
            node.style.removeProperty(key);
        }
        else {
            node.style.setProperty(key, value, important ? 'important' : '');
        }
    }
    function custom_event(type, detail, { bubbles = false, cancelable = false } = {}) {
        const e = document.createEvent('CustomEvent');
        e.initCustomEvent(type, bubbles, cancelable, detail);
        return e;
    }

    // we need to store the information for multiple documents because a Svelte application could also contain iframes
    // https://github.com/sveltejs/svelte/issues/3624
    const managed_styles = new Map();
    let active = 0;
    // https://github.com/darkskyapp/string-hash/blob/master/index.js
    function hash(str) {
        let hash = 5381;
        let i = str.length;
        while (i--)
            hash = ((hash << 5) - hash) ^ str.charCodeAt(i);
        return hash >>> 0;
    }
    function create_style_information(doc, node) {
        const info = { stylesheet: append_empty_stylesheet(node), rules: {} };
        managed_styles.set(doc, info);
        return info;
    }
    function create_rule(node, a, b, duration, delay, ease, fn, uid = 0) {
        const step = 16.666 / duration;
        let keyframes = '{\n';
        for (let p = 0; p <= 1; p += step) {
            const t = a + (b - a) * ease(p);
            keyframes += p * 100 + `%{${fn(t, 1 - t)}}\n`;
        }
        const rule = keyframes + `100% {${fn(b, 1 - b)}}\n}`;
        const name = `__svelte_${hash(rule)}_${uid}`;
        const doc = get_root_for_style(node);
        const { stylesheet, rules } = managed_styles.get(doc) || create_style_information(doc, node);
        if (!rules[name]) {
            rules[name] = true;
            stylesheet.insertRule(`@keyframes ${name} ${rule}`, stylesheet.cssRules.length);
        }
        const animation = node.style.animation || '';
        node.style.animation = `${animation ? `${animation}, ` : ''}${name} ${duration}ms linear ${delay}ms 1 both`;
        active += 1;
        return name;
    }
    function delete_rule(node, name) {
        const previous = (node.style.animation || '').split(', ');
        const next = previous.filter(name
            ? anim => anim.indexOf(name) < 0 // remove specific animation
            : anim => anim.indexOf('__svelte') === -1 // remove all Svelte animations
        );
        const deleted = previous.length - next.length;
        if (deleted) {
            node.style.animation = next.join(', ');
            active -= deleted;
            if (!active)
                clear_rules();
        }
    }
    function clear_rules() {
        raf(() => {
            if (active)
                return;
            managed_styles.forEach(info => {
                const { stylesheet } = info;
                let i = stylesheet.cssRules.length;
                while (i--)
                    stylesheet.deleteRule(i);
                info.rules = {};
            });
            managed_styles.clear();
        });
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
    function onMount(fn) {
        get_current_component().$$.on_mount.push(fn);
    }
    function onDestroy(fn) {
        get_current_component().$$.on_destroy.push(fn);
    }
    function setContext(key, context) {
        get_current_component().$$.context.set(key, context);
        return context;
    }
    function getContext(key) {
        return get_current_component().$$.context.get(key);
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
    function add_flush_callback(fn) {
        flush_callbacks.push(fn);
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

    let promise;
    function wait() {
        if (!promise) {
            promise = Promise.resolve();
            promise.then(() => {
                promise = null;
            });
        }
        return promise;
    }
    function dispatch(node, direction, kind) {
        node.dispatchEvent(custom_event(`${direction ? 'intro' : 'outro'}${kind}`));
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
    const null_transition = { duration: 0 };
    function create_in_transition(node, fn, params) {
        let config = fn(node, params);
        let running = false;
        let animation_name;
        let task;
        let uid = 0;
        function cleanup() {
            if (animation_name)
                delete_rule(node, animation_name);
        }
        function go() {
            const { delay = 0, duration = 300, easing = identity, tick = noop, css } = config || null_transition;
            if (css)
                animation_name = create_rule(node, 0, 1, duration, delay, easing, css, uid++);
            tick(0, 1);
            const start_time = now() + delay;
            const end_time = start_time + duration;
            if (task)
                task.abort();
            running = true;
            add_render_callback(() => dispatch(node, true, 'start'));
            task = loop(now => {
                if (running) {
                    if (now >= end_time) {
                        tick(1, 0);
                        dispatch(node, true, 'end');
                        cleanup();
                        return running = false;
                    }
                    if (now >= start_time) {
                        const t = easing((now - start_time) / duration);
                        tick(t, 1 - t);
                    }
                }
                return running;
            });
        }
        let started = false;
        return {
            start() {
                if (started)
                    return;
                started = true;
                delete_rule(node);
                if (is_function(config)) {
                    config = config();
                    wait().then(go);
                }
                else {
                    go();
                }
            },
            invalidate() {
                started = false;
            },
            end() {
                if (running) {
                    cleanup();
                    running = false;
                }
            }
        };
    }
    function create_out_transition(node, fn, params) {
        let config = fn(node, params);
        let running = true;
        let animation_name;
        const group = outros;
        group.r += 1;
        function go() {
            const { delay = 0, duration = 300, easing = identity, tick = noop, css } = config || null_transition;
            if (css)
                animation_name = create_rule(node, 1, 0, duration, delay, easing, css);
            const start_time = now() + delay;
            const end_time = start_time + duration;
            add_render_callback(() => dispatch(node, false, 'start'));
            loop(now => {
                if (running) {
                    if (now >= end_time) {
                        tick(0, 1);
                        dispatch(node, false, 'end');
                        if (!--group.r) {
                            // this will result in `end()` being called,
                            // so we don't need to clean up here
                            run_all(group.c);
                        }
                        return false;
                    }
                    if (now >= start_time) {
                        const t = easing((now - start_time) / duration);
                        tick(1 - t, t);
                    }
                }
                return running;
            });
        }
        if (is_function(config)) {
            wait().then(() => {
                // @ts-ignore
                config = config();
                go();
            });
        }
        else {
            go();
        }
        return {
            end(reset) {
                if (reset && config.tick) {
                    config.tick(1, 0);
                }
                if (running) {
                    if (animation_name)
                        delete_rule(node, animation_name);
                    running = false;
                }
            }
        };
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

    function bind(component, name, callback) {
        const index = component.$$.props[name];
        if (index !== undefined) {
            component.$$.bound[index] = callback;
            callback(component.$$.ctx[index]);
        }
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

    function getDefaultExportFromCjs (x) {
    	return x && x.__esModule && Object.prototype.hasOwnProperty.call(x, 'default') ? x['default'] : x;
    }

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
    let params = derived(match, $ => $.params || {});
    let matches = derived(match, $ => $.matches || []); // parents of active route and itself
    let components = derived(matches, $ => $.map(e => e.$$component).filter(e => e));// components to use in <Router/>

    /* node_modules/svelte-hash-router/src/components/Router.svelte generated by Svelte v3.48.0 */

    function create_fragment$d(ctx) {
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
    		id: create_fragment$d.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    let level = 0;

    function instance$d($$self, $$props, $$invalidate) {
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
    		init(this, options, instance$d, create_fragment$d, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Router",
    			options,
    			id: create_fragment$d.name
    		});
    	}
    }

    /* entries/portal/layout/modal.svelte generated by Svelte v3.48.0 */

    const file$8 = "entries/portal/layout/modal.svelte";

    // (28:2) {#if _show_big && current_big_compo}
    function create_if_block_1$2(ctx) {
    	let modal_big_wrapper;
    	let div0;
    	let svg;
    	let path;
    	let t;
    	let div1;
    	let switch_instance;
    	let current;
    	let mounted;
    	let dispose;
    	const switch_instance_spread_levels = [/*big_props*/ ctx[3]];
    	var switch_value = /*current_big_compo*/ ctx[2];

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
    			modal_big_wrapper = element("modal-big-wrapper");
    			div0 = element("div");
    			svg = svg_element("svg");
    			path = svg_element("path");
    			t = space();
    			div1 = element("div");
    			if (switch_instance) create_component(switch_instance.$$.fragment);
    			attr_dev(path, "d", "M14.53 4.53l-1.06-1.06L9 7.94 4.53 3.47 3.47 4.53 7.94 9l-4.47 4.47 1.06 1.06L9 10.06l4.47 4.47 1.06-1.06L10.06 9z");
    			add_location(path, file$8, 39, 10, 1025);
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg, "width", "18");
    			attr_dev(svg, "height", "18");
    			attr_dev(svg, "viewBox", "0 0 18 18");
    			add_location(svg, file$8, 33, 8, 882);
    			attr_dev(div0, "class", "modal-close absolute top-0 right-0 cursor-pointer flex flex-col items-center mt-4 mr-4 bg-white rounded-lg border-4 z-50");
    			add_location(div0, file$8, 29, 6, 695);
    			attr_dev(div1, "class", "sm:w-3/4 lg:w-1/2 3xl:w-1/3 modal-big-section svelte-ohvc1q");
    			add_location(div1, file$8, 45, 6, 1210);
    			set_custom_element_data(modal_big_wrapper, "class", "svelte-ohvc1q");
    			add_location(modal_big_wrapper, file$8, 28, 4, 669);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, modal_big_wrapper, anchor);
    			append_dev(modal_big_wrapper, div0);
    			append_dev(div0, svg);
    			append_dev(svg, path);
    			append_dev(modal_big_wrapper, t);
    			append_dev(modal_big_wrapper, div1);

    			if (switch_instance) {
    				mount_component(switch_instance, div1, null);
    			}

    			current = true;

    			if (!mounted) {
    				dispose = listen_dev(div0, "click", /*close_big*/ ctx[0], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			const switch_instance_changes = (dirty & /*big_props*/ 8)
    			? get_spread_update(switch_instance_spread_levels, [get_spread_object(/*big_props*/ ctx[3])])
    			: {};

    			if (switch_value !== (switch_value = /*current_big_compo*/ ctx[2])) {
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
    					mount_component(switch_instance, div1, null);
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
    			if (detaching) detach_dev(modal_big_wrapper);
    			if (switch_instance) destroy_component(switch_instance);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_1$2.name,
    		type: "if",
    		source: "(28:2) {#if _show_big && current_big_compo}",
    		ctx
    	});

    	return block;
    }

    // (27:0) {#key _show_big}
    function create_key_block_1(ctx) {
    	let if_block_anchor;
    	let current;
    	let if_block = /*_show_big*/ ctx[7] && /*current_big_compo*/ ctx[2] && create_if_block_1$2(ctx);

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
    			if (/*_show_big*/ ctx[7] && /*current_big_compo*/ ctx[2]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);

    					if (dirty & /*_show_big, current_big_compo*/ 132) {
    						transition_in(if_block, 1);
    					}
    				} else {
    					if_block = create_if_block_1$2(ctx);
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
    		id: create_key_block_1.name,
    		type: "key",
    		source: "(27:0) {#key _show_big}",
    		ctx
    	});

    	return block;
    }

    // (54:2) {#if _show_small && current_small_compo}
    function create_if_block$4(ctx) {
    	let modal_small_wrapper;
    	let div;
    	let svg;
    	let path;
    	let t;
    	let modal_small_section;
    	let switch_instance;
    	let current;
    	let mounted;
    	let dispose;
    	const switch_instance_spread_levels = [/*small_props*/ ctx[5]];
    	var switch_value = /*current_small_compo*/ ctx[4];

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
    			modal_small_wrapper = element("modal-small-wrapper");
    			div = element("div");
    			svg = svg_element("svg");
    			path = svg_element("path");
    			t = space();
    			modal_small_section = element("modal-small-section");
    			if (switch_instance) create_component(switch_instance.$$.fragment);
    			attr_dev(path, "d", "M14.53 4.53l-1.06-1.06L9 7.94 4.53 3.47 3.47 4.53 7.94 9l-4.47 4.47 1.06 1.06L9 10.06l4.47 4.47 1.06-1.06L10.06 9z");
    			add_location(path, file$8, 65, 10, 1819);
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg, "width", "18");
    			attr_dev(svg, "height", "18");
    			attr_dev(svg, "viewBox", "0 0 18 18");
    			add_location(svg, file$8, 59, 8, 1676);
    			attr_dev(div, "class", "modal-close absolute top-0 right-0 cursor-pointer flex flex-col items-center mt-4 mr-4 bg-white rounded-lg border-4 z-50");
    			add_location(div, file$8, 55, 6, 1487);
    			set_custom_element_data(modal_small_section, "class", "svelte-ohvc1q");
    			add_location(modal_small_section, file$8, 71, 6, 2004);
    			set_custom_element_data(modal_small_wrapper, "class", "svelte-ohvc1q");
    			add_location(modal_small_wrapper, file$8, 54, 4, 1459);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, modal_small_wrapper, anchor);
    			append_dev(modal_small_wrapper, div);
    			append_dev(div, svg);
    			append_dev(svg, path);
    			append_dev(modal_small_wrapper, t);
    			append_dev(modal_small_wrapper, modal_small_section);

    			if (switch_instance) {
    				mount_component(switch_instance, modal_small_section, null);
    			}

    			current = true;

    			if (!mounted) {
    				dispose = listen_dev(div, "click", /*close_small*/ ctx[1], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			const switch_instance_changes = (dirty & /*small_props*/ 32)
    			? get_spread_update(switch_instance_spread_levels, [get_spread_object(/*small_props*/ ctx[5])])
    			: {};

    			if (switch_value !== (switch_value = /*current_small_compo*/ ctx[4])) {
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
    					mount_component(switch_instance, modal_small_section, null);
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
    			if (detaching) detach_dev(modal_small_wrapper);
    			if (switch_instance) destroy_component(switch_instance);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$4.name,
    		type: "if",
    		source: "(54:2) {#if _show_small && current_small_compo}",
    		ctx
    	});

    	return block;
    }

    // (53:0) {#key _show_small}
    function create_key_block(ctx) {
    	let if_block_anchor;
    	let current;
    	let if_block = /*_show_small*/ ctx[6] && /*current_small_compo*/ ctx[4] && create_if_block$4(ctx);

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
    			if (/*_show_small*/ ctx[6] && /*current_small_compo*/ ctx[4]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);

    					if (dirty & /*_show_small, current_small_compo*/ 80) {
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
    		source: "(53:0) {#key _show_small}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$c(ctx) {
    	let previous_key = /*_show_big*/ ctx[7];
    	let t0;
    	let previous_key_1 = /*_show_small*/ ctx[6];
    	let t1;
    	let current;
    	let key_block0 = create_key_block_1(ctx);
    	let key_block1 = create_key_block(ctx);
    	const default_slot_template = /*#slots*/ ctx[11].default;
    	const default_slot = create_slot(default_slot_template, ctx, /*$$scope*/ ctx[10], null);

    	const block = {
    		c: function create() {
    			key_block0.c();
    			t0 = space();
    			key_block1.c();
    			t1 = space();
    			if (default_slot) default_slot.c();
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			key_block0.m(target, anchor);
    			insert_dev(target, t0, anchor);
    			key_block1.m(target, anchor);
    			insert_dev(target, t1, anchor);

    			if (default_slot) {
    				default_slot.m(target, anchor);
    			}

    			current = true;
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*_show_big*/ 128 && safe_not_equal(previous_key, previous_key = /*_show_big*/ ctx[7])) {
    				group_outros();
    				transition_out(key_block0, 1, 1, noop);
    				check_outros();
    				key_block0 = create_key_block_1(ctx);
    				key_block0.c();
    				transition_in(key_block0, 1);
    				key_block0.m(t0.parentNode, t0);
    			} else {
    				key_block0.p(ctx, dirty);
    			}

    			if (dirty & /*_show_small*/ 64 && safe_not_equal(previous_key_1, previous_key_1 = /*_show_small*/ ctx[6])) {
    				group_outros();
    				transition_out(key_block1, 1, 1, noop);
    				check_outros();
    				key_block1 = create_key_block(ctx);
    				key_block1.c();
    				transition_in(key_block1, 1);
    				key_block1.m(t1.parentNode, t1);
    			} else {
    				key_block1.p(ctx, dirty);
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
    			transition_in(key_block0);
    			transition_in(key_block1);
    			transition_in(default_slot, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(key_block0);
    			transition_out(key_block1);
    			transition_out(default_slot, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			key_block0.d(detaching);
    			if (detaching) detach_dev(t0);
    			key_block1.d(detaching);
    			if (detaching) detach_dev(t1);
    			if (default_slot) default_slot.d(detaching);
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

    function instance$c($$self, $$props, $$invalidate) {
    	let _show_big;
    	let _show_small;
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Modal', slots, ['default']);
    	let current_big_compo;
    	let big_props = {};
    	let current_small_compo;
    	let small_props = {};

    	const show_big = (_compo, _props) => {
    		$$invalidate(2, current_big_compo = _compo);
    		$$invalidate(3, big_props = _props);
    		$$invalidate(7, _show_big = true);
    	};

    	const close_big = () => {
    		$$invalidate(7, _show_big = false);
    		$$invalidate(2, current_big_compo = null);
    	};

    	const show_small = (_compo, _props) => {
    		$$invalidate(4, current_small_compo = _compo);
    		$$invalidate(5, small_props = _props);
    		$$invalidate(6, _show_small = true);
    	};

    	const close_small = () => {
    		$$invalidate(6, _show_small = false);
    		$$invalidate(4, current_small_compo = null);
    	};

    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Modal> was created with unknown prop '${key}'`);
    	});

    	$$self.$$set = $$props => {
    		if ('$$scope' in $$props) $$invalidate(10, $$scope = $$props.$$scope);
    	};

    	$$self.$capture_state = () => ({
    		current_big_compo,
    		big_props,
    		current_small_compo,
    		small_props,
    		show_big,
    		close_big,
    		show_small,
    		close_small,
    		_show_small,
    		_show_big
    	});

    	$$self.$inject_state = $$props => {
    		if ('current_big_compo' in $$props) $$invalidate(2, current_big_compo = $$props.current_big_compo);
    		if ('big_props' in $$props) $$invalidate(3, big_props = $$props.big_props);
    		if ('current_small_compo' in $$props) $$invalidate(4, current_small_compo = $$props.current_small_compo);
    		if ('small_props' in $$props) $$invalidate(5, small_props = $$props.small_props);
    		if ('_show_small' in $$props) $$invalidate(6, _show_small = $$props._show_small);
    		if ('_show_big' in $$props) $$invalidate(7, _show_big = $$props._show_big);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	$$invalidate(7, _show_big = true);
    	$$invalidate(6, _show_small = true);

    	return [
    		close_big,
    		close_small,
    		current_big_compo,
    		big_props,
    		current_small_compo,
    		small_props,
    		_show_small,
    		_show_big,
    		show_big,
    		show_small,
    		$$scope,
    		slots
    	];
    }

    class Modal extends SvelteComponentDev {
    	constructor(options) {
    		super(options);

    		init(this, options, instance$c, create_fragment$c, safe_not_equal, {
    			show_big: 8,
    			close_big: 0,
    			show_small: 9,
    			close_small: 1
    		});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Modal",
    			options,
    			id: create_fragment$c.name
    		});
    	}

    	get show_big() {
    		return this.$$.ctx[8];
    	}

    	set show_big(value) {
    		throw new Error("<Modal>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get close_big() {
    		return this.$$.ctx[0];
    	}

    	set close_big(value) {
    		throw new Error("<Modal>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get show_small() {
    		return this.$$.ctx[9];
    	}

    	set show_small(value) {
    		throw new Error("<Modal>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get close_small() {
    		return this.$$.ctx[1];
    	}

    	set close_small(value) {
    		throw new Error("<Modal>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    function cubicOut(t) {
        const f = t - 1.0;
        return f * f * f + 1.0;
    }

    function scale$1(node, { delay = 0, duration = 400, easing = cubicOut, start = 0, opacity = 0 } = {}) {
        const style = getComputedStyle(node);
        const target_opacity = +style.opacity;
        const transform = style.transform === 'none' ? '' : style.transform;
        const sd = 1 - start;
        const od = target_opacity * (1 - opacity);
        return {
            delay,
            duration,
            easing,
            css: (_t, u) => `
			transform: ${transform} scale(${1 - (sd * u)});
			opacity: ${target_opacity - (od * u)}
		`
        };
    }

    /* entries/xcompo/autotable/_dropdown.svelte generated by Svelte v3.48.0 */
    const file$7 = "entries/xcompo/autotable/_dropdown.svelte";

    // (42:4) {#if show}
    function create_if_block$3(ctx) {
    	let div;
    	let div_intro;
    	let div_outro;
    	let current;
    	const default_slot_template = /*#slots*/ ctx[3].default;
    	const default_slot = create_slot(default_slot_template, ctx, /*$$scope*/ ctx[2], null);

    	const block = {
    		c: function create() {
    			div = element("div");
    			if (default_slot) default_slot.c();
    			attr_dev(div, "class", "origin-top-right absolute right-0 w-48 py-2 mt-1 bg-white z-50 rounded shadow-md flex flex-col");
    			add_location(div, file$7, 42, 6, 1110);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);

    			if (default_slot) {
    				default_slot.m(div, null);
    			}

    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			if (default_slot) {
    				if (default_slot.p && (!current || dirty & /*$$scope*/ 4)) {
    					update_slot_base(
    						default_slot,
    						default_slot_template,
    						ctx,
    						/*$$scope*/ ctx[2],
    						!current
    						? get_all_dirty_from_scope(/*$$scope*/ ctx[2])
    						: get_slot_changes(default_slot_template, /*$$scope*/ ctx[2], dirty, null),
    						null
    					);
    				}
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(default_slot, local);

    			add_render_callback(() => {
    				if (div_outro) div_outro.end(1);
    				div_intro = create_in_transition(div, scale$1, { duration: 100, start: 0.95 });
    				div_intro.start();
    			});

    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(default_slot, local);
    			if (div_intro) div_intro.invalidate();
    			div_outro = create_out_transition(div, scale$1, { duration: 75, start: 0.95 });
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			if (default_slot) default_slot.d(detaching);
    			if (detaching && div_outro) div_outro.end();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$3.name,
    		type: "if",
    		source: "(42:4) {#if show}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$b(ctx) {
    	let div1;
    	let div0;
    	let button;
    	let t1;
    	let current;
    	let mounted;
    	let dispose;
    	let if_block = /*show*/ ctx[0] && create_if_block$3(ctx);

    	const block = {
    		c: function create() {
    			div1 = element("div");
    			div0 = element("div");
    			button = element("button");
    			button.textContent = "Options";
    			t1 = space();
    			if (if_block) if_block.c();
    			attr_dev(button, "class", "menu focus:outline-none focus:shadow-solid p-2 shadow border rounded");
    			add_location(button, file$7, 34, 4, 925);
    			add_location(div0, file$7, 33, 2, 915);
    			attr_dev(div1, "class", "relative");
    			add_location(div1, file$7, 32, 0, 870);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div1, anchor);
    			append_dev(div1, div0);
    			append_dev(div0, button);
    			append_dev(div0, t1);
    			if (if_block) if_block.m(div0, null);
    			/*div1_binding*/ ctx[5](div1);
    			current = true;

    			if (!mounted) {
    				dispose = listen_dev(button, "click", /*click_handler*/ ctx[4], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			if (/*show*/ ctx[0]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);

    					if (dirty & /*show*/ 1) {
    						transition_in(if_block, 1);
    					}
    				} else {
    					if_block = create_if_block$3(ctx);
    					if_block.c();
    					transition_in(if_block, 1);
    					if_block.m(div0, null);
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
    			if (detaching) detach_dev(div1);
    			if (if_block) if_block.d();
    			/*div1_binding*/ ctx[5](null);
    			mounted = false;
    			dispose();
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

    function instance$b($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Dropdown', slots, ['default']);
    	let show = false; // menu state
    	let menuRef = null; // menu wrapper DOM reference

    	onMount(() => {
    		const handleOutsideClick = event => {
    			if (show && !menuRef.contains(event.target)) {
    				$$invalidate(0, show = false);
    			}
    		};

    		const handleEscape = event => {
    			if (show && event.key === "Escape") {
    				$$invalidate(0, show = false);
    			}
    		};

    		// add events when element is added to the DOM
    		document.addEventListener("click", handleOutsideClick, false);

    		document.addEventListener("keyup", handleEscape, false);

    		// remove events when element is removed from the DOM
    		return () => {
    			document.removeEventListener("click", handleOutsideClick, false);
    			document.removeEventListener("keyup", handleEscape, false);
    		};
    	});

    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Dropdown> was created with unknown prop '${key}'`);
    	});

    	const click_handler = () => $$invalidate(0, show = !show);

    	function div1_binding($$value) {
    		binding_callbacks[$$value ? 'unshift' : 'push'](() => {
    			menuRef = $$value;
    			$$invalidate(1, menuRef);
    		});
    	}

    	$$self.$$set = $$props => {
    		if ('$$scope' in $$props) $$invalidate(2, $$scope = $$props.$$scope);
    	};

    	$$self.$capture_state = () => ({ onMount, scale: scale$1, show, menuRef });

    	$$self.$inject_state = $$props => {
    		if ('show' in $$props) $$invalidate(0, show = $$props.show);
    		if ('menuRef' in $$props) $$invalidate(1, menuRef = $$props.menuRef);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [show, menuRef, $$scope, slots, click_handler, div1_binding];
    }

    class Dropdown extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$b, create_fragment$b, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Dropdown",
    			options,
    			id: create_fragment$b.name
    		});
    	}
    }

    var iconSet = { "academic-cap": { "outline": "<path d=\"M12 14l9-5-9-5-9 5 9 5z\"/><path d=\"M12 14l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14z\"/><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 14l9-5-9-5-9 5 9 5zm0 0l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14zm-4 6v-7.5l4-2.222\"/>", "solid": "<path d=\"M10.394 2.08a1 1 0 00-.788 0l-7 3a1 1 0 000 1.84L5.25 8.051a.999.999 0 01.356-.257l4-1.714a1 1 0 11.788 1.838L7.667 9.088l1.94.831a1 1 0 00.787 0l7-3a1 1 0 000-1.838l-7-3zM3.31 9.397L5 10.12v4.102a8.969 8.969 0 00-1.05-.174 1 1 0 01-.89-.89 11.115 11.115 0 01.25-3.762zM9.3 16.573A9.026 9.026 0 007 14.935v-3.957l1.818.78a3 3 0 002.364 0l5.508-2.361a11.026 11.026 0 01.25 3.762 1 1 0 01-.89.89 8.968 8.968 0 00-5.35 2.524 1 1 0 01-1.4 0zM6 18a1 1 0 001-1v-2.065a8.935 8.935 0 00-2-.712V17a1 1 0 001 1z\"/>" }, "adjustments": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4\"/>", "solid": "<path d=\"M5 4a1 1 0 00-2 0v7.268a2 2 0 000 3.464V16a1 1 0 102 0v-1.268a2 2 0 000-3.464V4zM11 4a1 1 0 10-2 0v1.268a2 2 0 000 3.464V16a1 1 0 102 0V8.732a2 2 0 000-3.464V4zM16 3a1 1 0 011 1v7.268a2 2 0 010 3.464V16a1 1 0 11-2 0v-1.268a2 2 0 010-3.464V4a1 1 0 011-1z\"/>" }, "annotation": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M18 13V5a2 2 0 00-2-2H4a2 2 0 00-2 2v8a2 2 0 002 2h3l3 3 3-3h3a2 2 0 002-2zM5 7a1 1 0 011-1h8a1 1 0 110 2H6a1 1 0 01-1-1zm1 3a1 1 0 100 2h3a1 1 0 100-2H6z\" clip-rule=\"evenodd\"/>" }, "archive": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4\"/>", "solid": "<path d=\"M4 3a2 2 0 100 4h12a2 2 0 100-4H4z\"/><path fill-rule=\"evenodd\" d=\"M3 8h14v7a2 2 0 01-2 2H5a2 2 0 01-2-2V8zm5 3a1 1 0 011-1h2a1 1 0 110 2H9a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/>" }, "arrow-circle-down": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 13l-3 3m0 0l-3-3m3 3V8m0 13a9 9 0 110-18 9 9 0 010 18z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zm1-11a1 1 0 10-2 0v3.586L7.707 9.293a1 1 0 00-1.414 1.414l3 3a1 1 0 001.414 0l3-3a1 1 0 00-1.414-1.414L11 10.586V7z\" clip-rule=\"evenodd\"/>" }, "arrow-circle-left": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M11 15l-3-3m0 0l3-3m-3 3h8M3 12a9 9 0 1118 0 9 9 0 01-18 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zm.707-10.293a1 1 0 00-1.414-1.414l-3 3a1 1 0 000 1.414l3 3a1 1 0 001.414-1.414L9.414 11H13a1 1 0 100-2H9.414l1.293-1.293z\" clip-rule=\"evenodd\"/>" }, "arrow-circle-right": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 9l3 3m0 0l-3 3m3-3H8m13 0a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-8.707l-3-3a1 1 0 00-1.414 1.414L10.586 9H7a1 1 0 100 2h3.586l-1.293 1.293a1 1 0 101.414 1.414l3-3a1 1 0 000-1.414z\" clip-rule=\"evenodd\"/>" }, "arrow-circle-up": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 11l3-3m0 0l3 3m-3-3v8m0-13a9 9 0 110 18 9 9 0 010-18z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-8.707l-3-3a1 1 0 00-1.414 0l-3 3a1 1 0 001.414 1.414L9 9.414V13a1 1 0 102 0V9.414l1.293 1.293a1 1 0 001.414-1.414z\" clip-rule=\"evenodd\"/>" }, "arrow-down": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M19 14l-7 7m0 0l-7-7m7 7V3\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M16.707 10.293a1 1 0 010 1.414l-6 6a1 1 0 01-1.414 0l-6-6a1 1 0 111.414-1.414L9 14.586V3a1 1 0 012 0v11.586l4.293-4.293a1 1 0 011.414 0z\" clip-rule=\"evenodd\"/>" }, "arrow-left": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10 19l-7-7m0 0l7-7m-7 7h18\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M9.707 16.707a1 1 0 01-1.414 0l-6-6a1 1 0 010-1.414l6-6a1 1 0 011.414 1.414L5.414 9H17a1 1 0 110 2H5.414l4.293 4.293a1 1 0 010 1.414z\" clip-rule=\"evenodd\"/>" }, "arrow-narrow-down": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M16 17l-4 4m0 0l-4-4m4 4V3\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M14.707 12.293a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 111.414-1.414L9 14.586V3a1 1 0 012 0v11.586l2.293-2.293a1 1 0 011.414 0z\" clip-rule=\"evenodd\"/>" }, "arrow-narrow-left": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 16l-4-4m0 0l4-4m-4 4h18\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M7.707 14.707a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 1.414L5.414 9H17a1 1 0 110 2H5.414l2.293 2.293a1 1 0 010 1.414z\" clip-rule=\"evenodd\"/>" }, "arrow-narrow-right": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M17 8l4 4m0 0l-4 4m4-4H3\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M12.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-2.293-2.293a1 1 0 010-1.414z\" clip-rule=\"evenodd\"/>" }, "arrow-narrow-up": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 7l4-4m0 0l4 4m-4-4v18\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5.293 7.707a1 1 0 010-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 01-1.414 1.414L11 5.414V17a1 1 0 11-2 0V5.414L6.707 7.707a1 1 0 01-1.414 0z\" clip-rule=\"evenodd\"/>" }, "arrow-right": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M14 5l7 7m0 0l-7 7m7-7H3\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z\" clip-rule=\"evenodd\"/>" }, "arrow-sm-down": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M17 13l-5 5m0 0l-5-5m5 5V6\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M14.707 10.293a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 111.414-1.414L9 12.586V5a1 1 0 012 0v7.586l2.293-2.293a1 1 0 011.414 0z\" clip-rule=\"evenodd\"/>" }, "arrow-sm-left": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M11 17l-5-5m0 0l5-5m-5 5h12\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M9.707 14.707a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 1.414L7.414 9H15a1 1 0 110 2H7.414l2.293 2.293a1 1 0 010 1.414z\" clip-rule=\"evenodd\"/>" }, "arrow-sm-right": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 7l5 5m0 0l-5 5m5-5H6\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L12.586 11H5a1 1 0 110-2h7.586l-2.293-2.293a1 1 0 010-1.414z\" clip-rule=\"evenodd\"/>" }, "arrow-sm-up": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 11l5-5m0 0l5 5m-5-5v12\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5.293 9.707a1 1 0 010-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 01-1.414 1.414L11 7.414V15a1 1 0 11-2 0V7.414L6.707 9.707a1 1 0 01-1.414 0z\" clip-rule=\"evenodd\"/>" }, "arrow-up": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5 10l7-7m0 0l7 7m-7-7v18\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3.293 9.707a1 1 0 010-1.414l6-6a1 1 0 011.414 0l6 6a1 1 0 01-1.414 1.414L11 5.414V17a1 1 0 11-2 0V5.414L4.707 9.707a1 1 0 01-1.414 0z\" clip-rule=\"evenodd\"/>" }, "arrows-expand": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 8V4m0 0h4M4 4l5 5m11-1V4m0 0h-4m4 0l-5 5M4 16v4m0 0h4m-4 0l5-5m11 5l-5-5m5 5v-4m0 4h-4\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 4a1 1 0 011-1h4a1 1 0 010 2H6.414l2.293 2.293a1 1 0 01-1.414 1.414L5 6.414V8a1 1 0 01-2 0V4zm9 1a1 1 0 110-2h4a1 1 0 011 1v4a1 1 0 11-2 0V6.414l-2.293 2.293a1 1 0 11-1.414-1.414L13.586 5H12zm-9 7a1 1 0 112 0v1.586l2.293-2.293a1 1 0 011.414 1.414L6.414 15H8a1 1 0 110 2H4a1 1 0 01-1-1v-4zm13-1a1 1 0 011 1v4a1 1 0 01-1 1h-4a1 1 0 110-2h1.586l-2.293-2.293a1 1 0 011.414-1.414L15 13.586V12a1 1 0 011-1z\" clip-rule=\"evenodd\"/>" }, "at-symbol": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M14.243 5.757a6 6 0 10-.986 9.284 1 1 0 111.087 1.678A8 8 0 1118 10a3 3 0 01-4.8 2.401A4 4 0 1114 10a1 1 0 102 0c0-1.537-.586-3.07-1.757-4.243zM12 10a2 2 0 10-4 0 2 2 0 004 0z\" clip-rule=\"evenodd\"/>" }, "backspace": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2M3 12l6.414 6.414a2 2 0 001.414.586H19a2 2 0 002-2V7a2 2 0 00-2-2h-8.172a2 2 0 00-1.414.586L3 12z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M6.707 4.879A3 3 0 018.828 4H15a3 3 0 013 3v6a3 3 0 01-3 3H8.828a3 3 0 01-2.12-.879l-4.415-4.414a1 1 0 010-1.414l4.414-4.414zm4 2.414a1 1 0 00-1.414 1.414L10.586 10l-1.293 1.293a1 1 0 101.414 1.414L12 11.414l1.293 1.293a1 1 0 001.414-1.414L13.414 10l1.293-1.293a1 1 0 00-1.414-1.414L12 8.586l-1.293-1.293z\" clip-rule=\"evenodd\"/>" }, "badge-check": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M6.267 3.455a3.066 3.066 0 001.745-.723 3.066 3.066 0 013.976 0 3.066 3.066 0 001.745.723 3.066 3.066 0 012.812 2.812c.051.643.304 1.254.723 1.745a3.066 3.066 0 010 3.976 3.066 3.066 0 00-.723 1.745 3.066 3.066 0 01-2.812 2.812 3.066 3.066 0 00-1.745.723 3.066 3.066 0 01-3.976 0 3.066 3.066 0 00-1.745-.723 3.066 3.066 0 01-2.812-2.812 3.066 3.066 0 00-.723-1.745 3.066 3.066 0 010-3.976 3.066 3.066 0 00.723-1.745 3.066 3.066 0 012.812-2.812zm7.44 5.252a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z\" clip-rule=\"evenodd\"/>" }, "ban": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728A9 9 0 015.636 5.636m12.728 12.728L5.636 5.636\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M13.477 14.89A6 6 0 015.11 6.524l8.367 8.368zm1.414-1.414L6.524 5.11a6 6 0 018.367 8.367zM18 10a8 8 0 11-16 0 8 8 0 0116 0z\" clip-rule=\"evenodd\"/>" }, "beaker": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M7 2a1 1 0 00-.707 1.707L7 4.414v3.758a1 1 0 01-.293.707l-4 4C.817 14.769 2.156 18 4.828 18h10.343c2.673 0 4.012-3.231 2.122-5.121l-4-4A1 1 0 0113 8.172V4.414l.707-.707A1 1 0 0013 2H7zm2 6.172V4h2v4.172a3 3 0 00.879 2.12l1.027 1.028a4 4 0 00-2.171.102l-.47.156a4 4 0 01-2.53 0l-.563-.187a1.993 1.993 0 00-.114-.035l1.063-1.063A3 3 0 009 8.172z\" clip-rule=\"evenodd\"/>" }, "bell": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9\"/>", "solid": "<path d=\"M10 2a6 6 0 00-6 6v3.586l-.707.707A1 1 0 004 14h12a1 1 0 00.707-1.707L16 11.586V8a6 6 0 00-6-6zM10 18a3 3 0 01-3-3h6a3 3 0 01-3 3z\"/>" }, "book-open": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253\"/>", "solid": "<path d=\"M9 4.804A7.968 7.968 0 005.5 4c-1.255 0-2.443.29-3.5.804v10A7.969 7.969 0 015.5 14c1.669 0 3.218.51 4.5 1.385A7.962 7.962 0 0114.5 14c1.255 0 2.443.29 3.5.804v-10A7.968 7.968 0 0014.5 4c-1.255 0-2.443.29-3.5.804V12a1 1 0 11-2 0V4.804z\"/>" }, "bookmark-alt": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M16 4v12l-4-2-4 2V4M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 5a2 2 0 012-2h10a2 2 0 012 2v10a2 2 0 01-2 2H5a2 2 0 01-2-2V5zm11 1H6v8l4-2 4 2V6z\" clip-rule=\"evenodd\"/>" }, "bookmark": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z\"/>", "solid": "<path d=\"M5 4a2 2 0 012-2h6a2 2 0 012 2v14l-5-2.5L5 18V4z\"/>" }, "briefcase": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M6 6V5a3 3 0 013-3h2a3 3 0 013 3v1h2a2 2 0 012 2v3.57A22.952 22.952 0 0110 13a22.95 22.95 0 01-8-1.43V8a2 2 0 012-2h2zm2-1a1 1 0 011-1h2a1 1 0 011 1v1H8V5zm1 5a1 1 0 011-1h.01a1 1 0 110 2H10a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/><path d=\"M2 13.692V16a2 2 0 002 2h12a2 2 0 002-2v-2.308A24.974 24.974 0 0110 15c-2.796 0-5.487-.46-8-1.308z\"/>" }, "cake": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M21 15.546c-.523 0-1.046.151-1.5.454a2.704 2.704 0 01-3 0 2.704 2.704 0 00-3 0 2.704 2.704 0 01-3 0 2.704 2.704 0 00-3 0 2.704 2.704 0 01-3 0 2.701 2.701 0 00-1.5-.454M9 6v2m3-2v2m3-2v2M9 3h.01M12 3h.01M15 3h.01M21 21v-7a2 2 0 00-2-2H5a2 2 0 00-2 2v7h18zm-3-9v-2a2 2 0 00-2-2H8a2 2 0 00-2 2v2h12z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M6 3a1 1 0 011-1h.01a1 1 0 010 2H7a1 1 0 01-1-1zm2 3a1 1 0 00-2 0v1a2 2 0 00-2 2v1a2 2 0 00-2 2v.683a3.7 3.7 0 011.055.485 1.704 1.704 0 001.89 0 3.704 3.704 0 014.11 0 1.704 1.704 0 001.89 0 3.704 3.704 0 014.11 0 1.704 1.704 0 001.89 0A3.7 3.7 0 0118 12.683V12a2 2 0 00-2-2V9a2 2 0 00-2-2V6a1 1 0 10-2 0v1h-1V6a1 1 0 10-2 0v1H8V6zm10 8.868a3.704 3.704 0 01-4.055-.036 1.704 1.704 0 00-1.89 0 3.704 3.704 0 01-4.11 0 1.704 1.704 0 00-1.89 0A3.704 3.704 0 012 14.868V17a1 1 0 001 1h14a1 1 0 001-1v-2.132zM9 3a1 1 0 011-1h.01a1 1 0 110 2H10a1 1 0 01-1-1zm3 0a1 1 0 011-1h.01a1 1 0 110 2H13a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/>" }, "calculator": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M6 2a2 2 0 00-2 2v12a2 2 0 002 2h8a2 2 0 002-2V4a2 2 0 00-2-2H6zm1 2a1 1 0 000 2h6a1 1 0 100-2H7zm6 7a1 1 0 011 1v3a1 1 0 11-2 0v-3a1 1 0 011-1zm-3 3a1 1 0 100 2h.01a1 1 0 100-2H10zm-4 1a1 1 0 011-1h.01a1 1 0 110 2H7a1 1 0 01-1-1zm1-4a1 1 0 100 2h.01a1 1 0 100-2H7zm2 1a1 1 0 011-1h.01a1 1 0 110 2H10a1 1 0 01-1-1zm4-4a1 1 0 100 2h.01a1 1 0 100-2H13zM9 9a1 1 0 011-1h.01a1 1 0 110 2H10a1 1 0 01-1-1zM7 8a1 1 0 000 2h.01a1 1 0 000-2H7z\" clip-rule=\"evenodd\"/>" }, "calendar": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z\" clip-rule=\"evenodd\"/>" }, "camera": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z\"/><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 13a3 3 0 11-6 0 3 3 0 016 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M4 5a2 2 0 00-2 2v8a2 2 0 002 2h12a2 2 0 002-2V7a2 2 0 00-2-2h-1.586a1 1 0 01-.707-.293l-1.121-1.121A2 2 0 0011.172 3H8.828a2 2 0 00-1.414.586L6.293 4.707A1 1 0 015.586 5H4zm6 9a3 3 0 100-6 3 3 0 000 6z\" clip-rule=\"evenodd\"/>" }, "cash": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M4 4a2 2 0 00-2 2v4a2 2 0 002 2V6h10a2 2 0 00-2-2H4zm2 6a2 2 0 012-2h8a2 2 0 012 2v4a2 2 0 01-2 2H8a2 2 0 01-2-2v-4zm6 4a2 2 0 100-4 2 2 0 000 4z\" clip-rule=\"evenodd\"/>" }, "chart-bar": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z\"/>", "solid": "<path d=\"M2 11a1 1 0 011-1h2a1 1 0 011 1v5a1 1 0 01-1 1H3a1 1 0 01-1-1v-5zM8 7a1 1 0 011-1h2a1 1 0 011 1v9a1 1 0 01-1 1H9a1 1 0 01-1-1V7zM14 4a1 1 0 011-1h2a1 1 0 011 1v12a1 1 0 01-1 1h-2a1 1 0 01-1-1V4z\"/>" }, "chart-pie": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M11 3.055A9.001 9.001 0 1020.945 13H11V3.055z\"/><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M20.488 9H15V3.512A9.025 9.025 0 0120.488 9z\"/>", "solid": "<path d=\"M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z\"/><path d=\"M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z\"/>" }, "chart-square-bar": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M16 8v8m-4-5v5m-4-2v2m-2 4h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5 3a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2V5a2 2 0 00-2-2H5zm9 4a1 1 0 10-2 0v6a1 1 0 102 0V7zm-3 2a1 1 0 10-2 0v4a1 1 0 102 0V9zm-3 3a1 1 0 10-2 0v1a1 1 0 102 0v-1z\" clip-rule=\"evenodd\"/>" }, "chat-alt-2": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M17 8h2a2 2 0 012 2v6a2 2 0 01-2 2h-2v4l-4-4H9a1.994 1.994 0 01-1.414-.586m0 0L11 14h4a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2v4l.586-.586z\"/>", "solid": "<path d=\"M2 5a2 2 0 012-2h7a2 2 0 012 2v4a2 2 0 01-2 2H9l-3 3v-3H4a2 2 0 01-2-2V5z\"/><path d=\"M15 7v2a4 4 0 01-4 4H9.828l-1.766 1.767c.28.149.599.233.938.233h2l3 3v-3h2a2 2 0 002-2V9a2 2 0 00-2-2h-1z\"/>" }, "chat-alt": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M18 5v8a2 2 0 01-2 2h-5l-5 4v-4H4a2 2 0 01-2-2V5a2 2 0 012-2h12a2 2 0 012 2zM7 8H5v2h2V8zm2 0h2v2H9V8zm6 0h-2v2h2V8z\" clip-rule=\"evenodd\"/>" }, "chat": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M18 10c0 3.866-3.582 7-8 7a8.841 8.841 0 01-4.083-.98L2 17l1.338-3.123C2.493 12.767 2 11.434 2 10c0-3.866 3.582-7 8-7s8 3.134 8 7zM7 9H5v2h2V9zm8 0h-2v2h2V9zM9 9h2v2H9V9z\" clip-rule=\"evenodd\"/>" }, "check-circle": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z\" clip-rule=\"evenodd\"/>" }, "check": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5 13l4 4L19 7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z\" clip-rule=\"evenodd\"/>" }, "chevron-double-down": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M19 13l-7 7-7-7m14-8l-7 7-7-7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M15.707 4.293a1 1 0 010 1.414l-5 5a1 1 0 01-1.414 0l-5-5a1 1 0 011.414-1.414L10 8.586l4.293-4.293a1 1 0 011.414 0zm0 6a1 1 0 010 1.414l-5 5a1 1 0 01-1.414 0l-5-5a1 1 0 111.414-1.414L10 14.586l4.293-4.293a1 1 0 011.414 0z\" clip-rule=\"evenodd\"/>" }, "chevron-double-left": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M11 19l-7-7 7-7m8 14l-7-7 7-7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M15.707 15.707a1 1 0 01-1.414 0l-5-5a1 1 0 010-1.414l5-5a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 010 1.414zm-6 0a1 1 0 01-1.414 0l-5-5a1 1 0 010-1.414l5-5a1 1 0 011.414 1.414L5.414 10l4.293 4.293a1 1 0 010 1.414z\" clip-rule=\"evenodd\"/>" }, "chevron-double-right": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 5l7 7-7 7M5 5l7 7-7 7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10.293 15.707a1 1 0 010-1.414L14.586 10l-4.293-4.293a1 1 0 111.414-1.414l5 5a1 1 0 010 1.414l-5 5a1 1 0 01-1.414 0z\" clip-rule=\"evenodd\"/><path fill-rule=\"evenodd\" d=\"M4.293 15.707a1 1 0 010-1.414L8.586 10 4.293 5.707a1 1 0 011.414-1.414l5 5a1 1 0 010 1.414l-5 5a1 1 0 01-1.414 0z\" clip-rule=\"evenodd\"/>" }, "chevron-double-up": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5 11l7-7 7 7M5 19l7-7 7 7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M4.293 15.707a1 1 0 010-1.414l5-5a1 1 0 011.414 0l5 5a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414 0zm0-6a1 1 0 010-1.414l5-5a1 1 0 011.414 0l5 5a1 1 0 01-1.414 1.414L10 5.414 5.707 9.707a1 1 0 01-1.414 0z\" clip-rule=\"evenodd\"/>" }, "chevron-down": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M19 9l-7 7-7-7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z\" clip-rule=\"evenodd\"/>" }, "chevron-left": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 19l-7-7 7-7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z\" clip-rule=\"evenodd\"/>" }, "chevron-right": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 5l7 7-7 7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z\" clip-rule=\"evenodd\"/>" }, "chevron-up": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5 15l7-7 7 7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M14.707 12.707a1 1 0 01-1.414 0L10 9.414l-3.293 3.293a1 1 0 01-1.414-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 010 1.414z\" clip-rule=\"evenodd\"/>" }, "chip": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 3v2m6-2v2M9 19v2m6-2v2M5 9H3m2 6H3m18-6h-2m2 6h-2M7 19h10a2 2 0 002-2V7a2 2 0 00-2-2H7a2 2 0 00-2 2v10a2 2 0 002 2zM9 9h6v6H9V9z\"/>", "solid": "<path d=\"M13 7H7v6h6V7z\"/><path fill-rule=\"evenodd\" d=\"M7 2a1 1 0 012 0v1h2V2a1 1 0 112 0v1h2a2 2 0 012 2v2h1a1 1 0 110 2h-1v2h1a1 1 0 110 2h-1v2a2 2 0 01-2 2h-2v1a1 1 0 11-2 0v-1H9v1a1 1 0 11-2 0v-1H5a2 2 0 01-2-2v-2H2a1 1 0 110-2h1V9H2a1 1 0 010-2h1V5a2 2 0 012-2h2V2zM5 5h10v10H5V5z\" clip-rule=\"evenodd\"/>" }, "clipboard-check": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4\"/>", "solid": "<path d=\"M9 2a1 1 0 000 2h2a1 1 0 100-2H9z\"/><path fill-rule=\"evenodd\" d=\"M4 5a2 2 0 012-2 3 3 0 003 3h2a3 3 0 003-3 2 2 0 012 2v11a2 2 0 01-2 2H6a2 2 0 01-2-2V5zm9.707 5.707a1 1 0 00-1.414-1.414L9 12.586l-1.293-1.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z\" clip-rule=\"evenodd\"/>" }, "clipboard-copy": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3\"/>", "solid": "<path d=\"M8 2a1 1 0 000 2h2a1 1 0 100-2H8z\"/><path d=\"M3 5a2 2 0 012-2 3 3 0 003 3h2a3 3 0 003-3 2 2 0 012 2v6h-4.586l1.293-1.293a1 1 0 00-1.414-1.414l-3 3a1 1 0 000 1.414l3 3a1 1 0 001.414-1.414L10.414 13H15v3a2 2 0 01-2 2H5a2 2 0 01-2-2V5zM15 11h2a1 1 0 110 2h-2v-2z\"/>" }, "clipboard-list": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01\"/>", "solid": "<path d=\"M9 2a1 1 0 000 2h2a1 1 0 100-2H9z\"/><path fill-rule=\"evenodd\" d=\"M4 5a2 2 0 012-2 3 3 0 003 3h2a3 3 0 003-3 2 2 0 012 2v11a2 2 0 01-2 2H6a2 2 0 01-2-2V5zm3 4a1 1 0 000 2h.01a1 1 0 100-2H7zm3 0a1 1 0 000 2h3a1 1 0 100-2h-3zm-3 4a1 1 0 100 2h.01a1 1 0 100-2H7zm3 0a1 1 0 100 2h3a1 1 0 100-2h-3z\" clip-rule=\"evenodd\"/>" }, "clipboard": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2\"/>", "solid": "<path d=\"M8 3a1 1 0 011-1h2a1 1 0 110 2H9a1 1 0 01-1-1z\"/><path d=\"M6 3a2 2 0 00-2 2v11a2 2 0 002 2h8a2 2 0 002-2V5a2 2 0 00-2-2 3 3 0 01-3 3H9a3 3 0 01-3-3z\"/>" }, "clock": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z\" clip-rule=\"evenodd\"/>" }, "cloud-download": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M2 9.5A3.5 3.5 0 005.5 13H9v2.586l-1.293-1.293a1 1 0 00-1.414 1.414l3 3a1 1 0 001.414 0l3-3a1 1 0 00-1.414-1.414L11 15.586V13h2.5a4.5 4.5 0 10-.616-8.958 4.002 4.002 0 10-7.753 1.977A3.5 3.5 0 002 9.5zm9 3.5H9V8a1 1 0 012 0v5z\" clip-rule=\"evenodd\"/>" }, "cloud-upload": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12\"/>", "solid": "<path d=\"M5.5 13a3.5 3.5 0 01-.369-6.98 4 4 0 117.753-1.977A4.5 4.5 0 1113.5 13H11V9.413l1.293 1.293a1 1 0 001.414-1.414l-3-3a1 1 0 00-1.414 0l-3 3a1 1 0 001.414 1.414L9 9.414V13H5.5z\"/><path d=\"M9 13h2v5a1 1 0 11-2 0v-5z\"/>" }, "cloud": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 15a4 4 0 004 4h9a5 5 0 10-.1-9.999 5.002 5.002 0 10-9.78 2.096A4.001 4.001 0 003 15z\"/>", "solid": "<path d=\"M5.5 16a3.5 3.5 0 01-.369-6.98 4 4 0 117.753-1.977A4.5 4.5 0 1113.5 16h-8z\"/>" }, "code": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M12.316 3.051a1 1 0 01.633 1.265l-4 12a1 1 0 11-1.898-.632l4-12a1 1 0 011.265-.633zM5.707 6.293a1 1 0 010 1.414L3.414 10l2.293 2.293a1 1 0 11-1.414 1.414l-3-3a1 1 0 010-1.414l3-3a1 1 0 011.414 0zm8.586 0a1 1 0 011.414 0l3 3a1 1 0 010 1.414l-3 3a1 1 0 11-1.414-1.414L16.586 10l-2.293-2.293a1 1 0 010-1.414z\" clip-rule=\"evenodd\"/>" }, "cog": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z\"/><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 12a3 3 0 11-6 0 3 3 0 016 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z\" clip-rule=\"evenodd\"/>" }, "collection": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10\"/>", "solid": "<path d=\"M7 3a1 1 0 000 2h6a1 1 0 100-2H7zM4 7a1 1 0 011-1h10a1 1 0 110 2H5a1 1 0 01-1-1zM2 11a2 2 0 012-2h12a2 2 0 012 2v4a2 2 0 01-2 2H4a2 2 0 01-2-2v-4z\"/>" }, "color-swatch": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M4 2a2 2 0 00-2 2v11a3 3 0 106 0V4a2 2 0 00-2-2H4zm1 14a1 1 0 100-2 1 1 0 000 2zm5-1.757l4.9-4.9a2 2 0 000-2.828L13.485 5.1a2 2 0 00-2.828 0L10 5.757v8.486zM16 18H9.071l6-6H16a2 2 0 012 2v2a2 2 0 01-2 2z\" clip-rule=\"evenodd\"/>" }, "credit-card": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z\"/>", "solid": "<path d=\"M4 4a2 2 0 00-2 2v1h16V6a2 2 0 00-2-2H4z\"/><path fill-rule=\"evenodd\" d=\"M18 9H2v5a2 2 0 002 2h12a2 2 0 002-2V9zM4 13a1 1 0 011-1h1a1 1 0 110 2H5a1 1 0 01-1-1zm5-1a1 1 0 100 2h1a1 1 0 100-2H9z\" clip-rule=\"evenodd\"/>" }, "cube-transparent": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M14 10l-2 1m0 0l-2-1m2 1v2.5M20 7l-2 1m2-1l-2-1m2 1v2.5M14 4l-2-1-2 1M4 7l2-1M4 7l2 1M4 7v2.5M12 21l-2-1m2 1l2-1m-2 1v-2.5M6 18l-2-1v-2.5M18 18l2-1v-2.5\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M9.504 1.132a1 1 0 01.992 0l1.75 1a1 1 0 11-.992 1.736L10 3.152l-1.254.716a1 1 0 11-.992-1.736l1.75-1zM5.618 4.504a1 1 0 01-.372 1.364L5.016 6l.23.132a1 1 0 11-.992 1.736L4 7.723V8a1 1 0 01-2 0V6a.996.996 0 01.52-.878l1.734-.99a1 1 0 011.364.372zm8.764 0a1 1 0 011.364-.372l1.733.99A1.002 1.002 0 0118 6v2a1 1 0 11-2 0v-.277l-.254.145a1 1 0 11-.992-1.736l.23-.132-.23-.132a1 1 0 01-.372-1.364zm-7 4a1 1 0 011.364-.372L10 8.848l1.254-.716a1 1 0 11.992 1.736L11 10.58V12a1 1 0 11-2 0v-1.42l-1.246-.712a1 1 0 01-.372-1.364zM3 11a1 1 0 011 1v1.42l1.246.712a1 1 0 11-.992 1.736l-1.75-1A1 1 0 012 14v-2a1 1 0 011-1zm14 0a1 1 0 011 1v2a1 1 0 01-.504.868l-1.75 1a1 1 0 11-.992-1.736L16 13.42V12a1 1 0 011-1zm-9.618 5.504a1 1 0 011.364-.372l.254.145V16a1 1 0 112 0v.277l.254-.145a1 1 0 11.992 1.736l-1.735.992a.995.995 0 01-1.022 0l-1.735-.992a1 1 0 01-.372-1.364z\" clip-rule=\"evenodd\"/>" }, "cube": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4\"/>", "solid": "<path d=\"M11 17a1 1 0 001.447.894l4-2A1 1 0 0017 15V9.236a1 1 0 00-1.447-.894l-4 2a1 1 0 00-.553.894V17zM15.211 6.276a1 1 0 000-1.788l-4.764-2.382a1 1 0 00-.894 0L4.789 4.488a1 1 0 000 1.788l4.764 2.382a1 1 0 00.894 0l4.764-2.382zM4.447 8.342A1 1 0 003 9.236V15a1 1 0 00.553.894l4 2A1 1 0 009 17v-5.764a1 1 0 00-.553-.894l-4-2z\"/>" }, "currency-bangladeshi": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M11 11V9a2 2 0 00-2-2m2 4v4a2 2 0 104 0v-1m-4-3H9m2 0h4m6 1a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zM7 4a1 1 0 000 2 1 1 0 011 1v1H7a1 1 0 000 2h1v3a3 3 0 106 0v-1a1 1 0 10-2 0v1a1 1 0 11-2 0v-3h3a1 1 0 100-2h-3V7a3 3 0 00-3-3z\" clip-rule=\"evenodd\"/>" }, "currency-dollar": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path d=\"M8.433 7.418c.155-.103.346-.196.567-.267v1.698a2.305 2.305 0 01-.567-.267C8.07 8.34 8 8.114 8 8c0-.114.07-.34.433-.582zM11 12.849v-1.698c.22.071.412.164.567.267.364.243.433.468.433.582 0 .114-.07.34-.433.582a2.305 2.305 0 01-.567.267z\"/><path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zm1-13a1 1 0 10-2 0v.092a4.535 4.535 0 00-1.676.662C6.602 6.234 6 7.009 6 8c0 .99.602 1.765 1.324 2.246.48.32 1.054.545 1.676.662v1.941c-.391-.127-.68-.317-.843-.504a1 1 0 10-1.51 1.31c.562.649 1.413 1.076 2.353 1.253V15a1 1 0 102 0v-.092a4.535 4.535 0 001.676-.662C13.398 13.766 14 12.991 14 12c0-.99-.602-1.765-1.324-2.246A4.535 4.535 0 0011 9.092V7.151c.391.127.68.317.843.504a1 1 0 101.511-1.31c-.563-.649-1.413-1.076-2.354-1.253V5z\" clip-rule=\"evenodd\"/>" }, "currency-euro": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M14.121 15.536c-1.171 1.952-3.07 1.952-4.242 0-1.172-1.953-1.172-5.119 0-7.072 1.171-1.952 3.07-1.952 4.242 0M8 10.5h4m-4 3h4m9-1.5a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zM8.736 6.979C9.208 6.193 9.696 6 10 6c.304 0 .792.193 1.264.979a1 1 0 001.715-1.029C12.279 4.784 11.232 4 10 4s-2.279.784-2.979 1.95c-.285.475-.507 1-.67 1.55H6a1 1 0 000 2h.013a9.358 9.358 0 000 1H6a1 1 0 100 2h.351c.163.55.385 1.075.67 1.55C7.721 15.216 8.768 16 10 16s2.279-.784 2.979-1.95a1 1 0 10-1.715-1.029c-.472.786-.96.979-1.264.979-.304 0-.792-.193-1.264-.979a4.265 4.265 0 01-.264-.521H10a1 1 0 100-2H8.017a7.36 7.36 0 010-1H10a1 1 0 100-2H8.472c.08-.185.167-.36.264-.521z\" clip-rule=\"evenodd\"/>" }, "currency-pound": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 9a2 2 0 10-4 0v5a2 2 0 01-2 2h6m-6-4h4m8 0a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zm1-14a3 3 0 00-3 3v2H7a1 1 0 000 2h1v1a1 1 0 01-1 1 1 1 0 100 2h6a1 1 0 100-2H9.83c.11-.313.17-.65.17-1v-1h1a1 1 0 100-2h-1V7a1 1 0 112 0 1 1 0 102 0 3 3 0 00-3-3z\" clip-rule=\"evenodd\"/>" }, "currency-rupee": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 8h6m-5 0a3 3 0 110 6H9l3 3m-3-6h6m6 1a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zM7 5a1 1 0 100 2h1a2 2 0 011.732 1H7a1 1 0 100 2h2.732A2 2 0 018 11H7a1 1 0 00-.707 1.707l3 3a1 1 0 001.414-1.414l-1.483-1.484A4.008 4.008 0 0011.874 10H13a1 1 0 100-2h-1.126a3.976 3.976 0 00-.41-1H13a1 1 0 100-2H7z\" clip-rule=\"evenodd\"/>" }, "currency-yen": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 8l3 5m0 0l3-5m-3 5v4m-3-5h6m-6 3h6m6-3a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zM7.858 5.485a1 1 0 00-1.715 1.03L7.633 9H7a1 1 0 100 2h1.834l.166.277V12H7a1 1 0 100 2h2v1a1 1 0 102 0v-1h2a1 1 0 100-2h-2v-.723l.166-.277H13a1 1 0 100-2h-.634l1.492-2.486a1 1 0 10-1.716-1.029L10.034 9h-.068L7.858 5.485z\" clip-rule=\"evenodd\"/>" }, "cursor-click": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 15l-2 5L9 9l11 4-5 2zm0 0l5 5M7.188 2.239l.777 2.897M5.136 7.965l-2.898-.777M13.95 4.05l-2.122 2.122m-5.657 5.656l-2.12 2.122\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M6.672 1.911a1 1 0 10-1.932.518l.259.966a1 1 0 001.932-.518l-.26-.966zM2.429 4.74a1 1 0 10-.517 1.932l.966.259a1 1 0 00.517-1.932l-.966-.26zm8.814-.569a1 1 0 00-1.415-1.414l-.707.707a1 1 0 101.415 1.415l.707-.708zm-7.071 7.072l.707-.707A1 1 0 003.465 9.12l-.708.707a1 1 0 001.415 1.415zm3.2-5.171a1 1 0 00-1.3 1.3l4 10a1 1 0 001.823.075l1.38-2.759 3.018 3.02a1 1 0 001.414-1.415l-3.019-3.02 2.76-1.379a1 1 0 00-.076-1.822l-10-4z\" clip-rule=\"evenodd\"/>" }, "database": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4m0 5c0 2.21-3.582 4-8 4s-8-1.79-8-4\"/>", "solid": "<path d=\"M3 12v3c0 1.657 3.134 3 7 3s7-1.343 7-3v-3c0 1.657-3.134 3-7 3s-7-1.343-7-3z\"/><path d=\"M3 7v3c0 1.657 3.134 3 7 3s7-1.343 7-3V7c0 1.657-3.134 3-7 3S3 8.657 3 7z\"/><path d=\"M17 5c0 1.657-3.134 3-7 3S3 6.657 3 5s3.134-3 7-3 7 1.343 7 3z\"/>" }, "desktop-computer": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 5a2 2 0 012-2h10a2 2 0 012 2v8a2 2 0 01-2 2h-2.22l.123.489.804.804A1 1 0 0113 18H7a1 1 0 01-.707-1.707l.804-.804L7.22 15H5a2 2 0 01-2-2V5zm5.771 7H5V5h10v7H8.771z\" clip-rule=\"evenodd\"/>" }, "device-mobile": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M7 2a2 2 0 00-2 2v12a2 2 0 002 2h6a2 2 0 002-2V4a2 2 0 00-2-2H7zm3 14a1 1 0 100-2 1 1 0 000 2z\" clip-rule=\"evenodd\"/>" }, "device-tablet": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 18h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M6 2a2 2 0 00-2 2v12a2 2 0 002 2h8a2 2 0 002-2V4a2 2 0 00-2-2H6zm4 14a1 1 0 100-2 1 1 0 000 2z\" clip-rule=\"evenodd\"/>" }, "document-add": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 13h6m-3-3v6m5 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M6 2a2 2 0 00-2 2v12a2 2 0 002 2h8a2 2 0 002-2V7.414A2 2 0 0015.414 6L12 2.586A2 2 0 0010.586 2H6zm5 6a1 1 0 10-2 0v2H7a1 1 0 100 2h2v2a1 1 0 102 0v-2h2a1 1 0 100-2h-2V8z\" clip-rule=\"evenodd\"/>" }, "document-download": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M6 2a2 2 0 00-2 2v12a2 2 0 002 2h8a2 2 0 002-2V7.414A2 2 0 0015.414 6L12 2.586A2 2 0 0010.586 2H6zm5 6a1 1 0 10-2 0v3.586l-1.293-1.293a1 1 0 10-1.414 1.414l3 3a1 1 0 001.414 0l3-3a1 1 0 00-1.414-1.414L11 11.586V8z\" clip-rule=\"evenodd\"/>" }, "document-duplicate": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 7v8a2 2 0 002 2h6M8 7V5a2 2 0 012-2h4.586a1 1 0 01.707.293l4.414 4.414a1 1 0 01.293.707V15a2 2 0 01-2 2h-2M8 7H6a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2v-2\"/>", "solid": "<path d=\"M9 2a2 2 0 00-2 2v8a2 2 0 002 2h6a2 2 0 002-2V6.414A2 2 0 0016.414 5L14 2.586A2 2 0 0012.586 2H9z\"/><path d=\"M3 8a2 2 0 012-2v10h8a2 2 0 01-2 2H5a2 2 0 01-2-2V8z\"/>" }, "document-remove": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 13h6m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M6 2a2 2 0 00-2 2v12a2 2 0 002 2h8a2 2 0 002-2V7.414A2 2 0 0015.414 6L12 2.586A2 2 0 0010.586 2H6zm1 8a1 1 0 100 2h6a1 1 0 100-2H7z\" clip-rule=\"evenodd\"/>" }, "document-report": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M6 2a2 2 0 00-2 2v12a2 2 0 002 2h8a2 2 0 002-2V7.414A2 2 0 0015.414 6L12 2.586A2 2 0 0010.586 2H6zm2 10a1 1 0 10-2 0v3a1 1 0 102 0v-3zm2-3a1 1 0 011 1v5a1 1 0 11-2 0v-5a1 1 0 011-1zm4-1a1 1 0 10-2 0v7a1 1 0 102 0V8z\" clip-rule=\"evenodd\"/>" }, "document-search": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10 21h7a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v11m0 5l4.879-4.879m0 0a3 3 0 104.243-4.242 3 3 0 00-4.243 4.242z\"/>", "solid": "<path d=\"M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2h-1.528A6 6 0 004 9.528V4z\"/><path fill-rule=\"evenodd\" d=\"M8 10a4 4 0 00-3.446 6.032l-1.261 1.26a1 1 0 101.414 1.415l1.261-1.261A4 4 0 108 10zm-2 4a2 2 0 114 0 2 2 0 01-4 0z\" clip-rule=\"evenodd\"/>" }, "document-text": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4zm2 6a1 1 0 011-1h6a1 1 0 110 2H7a1 1 0 01-1-1zm1 3a1 1 0 100 2h6a1 1 0 100-2H7z\" clip-rule=\"evenodd\"/>" }, "document": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4z\" clip-rule=\"evenodd\"/>" }, "dots-circle-horizontal": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 12h.01M12 12h.01M16 12h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zM7 9H5v2h2V9zm8 0h-2v2h2V9zM9 9h2v2H9V9z\" clip-rule=\"evenodd\"/>" }, "dots-horizontal": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5 12h.01M12 12h.01M19 12h.01M6 12a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0z\"/>", "solid": "<path d=\"M6 10a2 2 0 11-4 0 2 2 0 014 0zM12 10a2 2 0 11-4 0 2 2 0 014 0zM16 12a2 2 0 100-4 2 2 0 000 4z\"/>" }, "dots-vertical": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z\"/>", "solid": "<path d=\"M10 6a2 2 0 110-4 2 2 0 010 4zM10 12a2 2 0 110-4 2 2 0 010 4zM10 18a2 2 0 110-4 2 2 0 010 4z\"/>" }, "download": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm3.293-7.707a1 1 0 011.414 0L9 10.586V3a1 1 0 112 0v7.586l1.293-1.293a1 1 0 111.414 1.414l-3 3a1 1 0 01-1.414 0l-3-3a1 1 0 010-1.414z\" clip-rule=\"evenodd\"/>" }, "duplicate": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z\"/>", "solid": "<path d=\"M7 9a2 2 0 012-2h6a2 2 0 012 2v6a2 2 0 01-2 2H9a2 2 0 01-2-2V9z\"/><path d=\"M5 3a2 2 0 00-2 2v6a2 2 0 002 2V5h8a2 2 0 00-2-2H5z\"/>" }, "emoji-happy": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M14.828 14.828a4 4 0 01-5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zM7 9a1 1 0 100-2 1 1 0 000 2zm7-1a1 1 0 11-2 0 1 1 0 012 0zm-.464 5.535a1 1 0 10-1.415-1.414 3 3 0 01-4.242 0 1 1 0 00-1.415 1.414 5 5 0 007.072 0z\" clip-rule=\"evenodd\"/>" }, "emoji-sad": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zM7 9a1 1 0 100-2 1 1 0 000 2zm7-1a1 1 0 11-2 0 1 1 0 012 0zm-7.536 5.879a1 1 0 001.415 0 3 3 0 014.242 0 1 1 0 001.415-1.415 5 5 0 00-7.072 0 1 1 0 000 1.415z\" clip-rule=\"evenodd\"/>" }, "exclamation-circle": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z\" clip-rule=\"evenodd\"/>" }, "exclamation": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z\" clip-rule=\"evenodd\"/>" }, "external-link": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14\"/>", "solid": "<path d=\"M11 3a1 1 0 100 2h2.586l-6.293 6.293a1 1 0 101.414 1.414L15 6.414V9a1 1 0 102 0V4a1 1 0 00-1-1h-5z\"/><path d=\"M5 5a2 2 0 00-2 2v8a2 2 0 002 2h8a2 2 0 002-2v-3a1 1 0 10-2 0v3H5V7h3a1 1 0 000-2H5z\"/>" }, "eye-off": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3.707 2.293a1 1 0 00-1.414 1.414l14 14a1 1 0 001.414-1.414l-1.473-1.473A10.014 10.014 0 0019.542 10C18.268 5.943 14.478 3 10 3a9.958 9.958 0 00-4.512 1.074l-1.78-1.781zm4.261 4.26l1.514 1.515a2.003 2.003 0 012.45 2.45l1.514 1.514a4 4 0 00-5.478-5.478z\" clip-rule=\"evenodd\"/><path d=\"M12.454 16.697L9.75 13.992a4 4 0 01-3.742-3.741L2.335 6.578A9.98 9.98 0 00.458 10c1.274 4.057 5.065 7 9.542 7 .847 0 1.669-.105 2.454-.303z\"/>" }, "eye": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 12a3 3 0 11-6 0 3 3 0 016 0z\"/><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z\"/>", "solid": "<path d=\"M10 12a2 2 0 100-4 2 2 0 000 4z\"/><path fill-rule=\"evenodd\" d=\"M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z\" clip-rule=\"evenodd\"/>" }, "fast-forward": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M11.933 12.8a1 1 0 000-1.6L6.6 7.2A1 1 0 005 8v8a1 1 0 001.6.8l5.333-4zM19.933 12.8a1 1 0 000-1.6l-5.333-4A1 1 0 0013 8v8a1 1 0 001.6.8l5.333-4z\"/>", "solid": "<path d=\"M4.555 5.168A1 1 0 003 6v8a1 1 0 001.555.832L10 11.202V14a1 1 0 001.555.832l6-4a1 1 0 000-1.664l-6-4A1 1 0 0010 6v2.798l-5.445-3.63z\"/>" }, "film": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 4v16M17 4v16M3 8h4m10 0h4M3 12h18M3 16h4m10 0h4M4 20h16a1 1 0 001-1V5a1 1 0 00-1-1H4a1 1 0 00-1 1v14a1 1 0 001 1z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M4 3a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V5a2 2 0 00-2-2H4zm3 2h6v4H7V5zm8 8v2h1v-2h-1zm-2-2H7v4h6v-4zm2 0h1V9h-1v2zm1-4V5h-1v2h1zM5 5v2H4V5h1zm0 4H4v2h1V9zm-1 4h1v2H4v-2z\" clip-rule=\"evenodd\"/>" }, "filter": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 3a1 1 0 011-1h12a1 1 0 011 1v3a1 1 0 01-.293.707L12 11.414V15a1 1 0 01-.293.707l-2 2A1 1 0 018 17v-5.586L3.293 6.707A1 1 0 013 6V3z\" clip-rule=\"evenodd\"/>" }, "finger-print": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 11c0 3.517-1.009 6.799-2.753 9.571m-3.44-2.04l.054-.09A13.916 13.916 0 008 11a4 4 0 118 0c0 1.017-.07 2.019-.203 3m-2.118 6.844A21.88 21.88 0 0015.171 17m3.839 1.132c.645-2.266.99-4.659.99-7.132A8 8 0 008 4.07M3 15.364c.64-1.319 1-2.8 1-4.364 0-1.457.39-2.823 1.07-4\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M6.625 2.655A9 9 0 0119 11a1 1 0 11-2 0 7 7 0 00-9.625-6.492 1 1 0 11-.75-1.853zM4.662 4.959A1 1 0 014.75 6.37 6.97 6.97 0 003 11a1 1 0 11-2 0 8.97 8.97 0 012.25-5.953 1 1 0 011.412-.088z\" clip-rule=\"evenodd\"/><path fill-rule=\"evenodd\" d=\"M5 11a5 5 0 1110 0 1 1 0 11-2 0 3 3 0 10-6 0c0 1.677-.345 3.276-.968 4.729a1 1 0 11-1.838-.789A9.964 9.964 0 005 11zm8.921 2.012a1 1 0 01.831 1.145 19.86 19.86 0 01-.545 2.436 1 1 0 11-1.92-.558c.207-.713.371-1.445.49-2.192a1 1 0 011.144-.83z\" clip-rule=\"evenodd\"/><path fill-rule=\"evenodd\" d=\"M10 10a1 1 0 011 1c0 2.236-.46 4.368-1.29 6.304a1 1 0 01-1.838-.789A13.952 13.952 0 009 11a1 1 0 011-1z\" clip-rule=\"evenodd\"/>" }, "fire": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M17.657 18.657A8 8 0 016.343 7.343S7 9 9 10c0-2 .5-5 2.986-7C14 5 16.09 5.777 17.656 7.343A7.975 7.975 0 0120 13a7.975 7.975 0 01-2.343 5.657z\"/><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9.879 16.121A3 3 0 1012.015 11L11 14H9c0 .768.293 1.536.879 2.121z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M12.395 2.553a1 1 0 00-1.45-.385c-.345.23-.614.558-.822.88-.214.33-.403.713-.57 1.116-.334.804-.614 1.768-.84 2.734a31.365 31.365 0 00-.613 3.58 2.64 2.64 0 01-.945-1.067c-.328-.68-.398-1.534-.398-2.654A1 1 0 005.05 6.05 6.981 6.981 0 003 11a7 7 0 1011.95-4.95c-.592-.591-.98-.985-1.348-1.467-.363-.476-.724-1.063-1.207-2.03zM12.12 15.12A3 3 0 017 13s.879.5 2.5.5c0-1 .5-4 1.25-4.5.5 1 .786 1.293 1.371 1.879A2.99 2.99 0 0113 13a2.99 2.99 0 01-.879 2.121z\" clip-rule=\"evenodd\"/>" }, "flag": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 21v-4m0 0V5a2 2 0 012-2h6.5l1 1H21l-3 6 3 6h-8.5l-1-1H5a2 2 0 00-2 2zm9-13.5V9\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 6a3 3 0 013-3h10a1 1 0 01.8 1.6L14.25 8l2.55 3.4A1 1 0 0116 13H6a1 1 0 00-1 1v3a1 1 0 11-2 0V6z\" clip-rule=\"evenodd\"/>" }, "folder-add": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 13h6m-3-3v6m-9 1V7a2 2 0 012-2h6l2 2h6a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2z\"/>", "solid": "<path d=\"M2 6a2 2 0 012-2h5l2 2h5a2 2 0 012 2v6a2 2 0 01-2 2H4a2 2 0 01-2-2V6z\"/><path stroke=\"#fff\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 11h4m-2-2v4\"/>" }, "folder-download": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 10v6m0 0l-3-3m3 3l3-3M3 17V7a2 2 0 012-2h6l2 2h6a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2z\"/>", "solid": "<path d=\"M2 6a2 2 0 012-2h5l2 2h5a2 2 0 012 2v6a2 2 0 01-2 2H4a2 2 0 01-2-2V6z\"/><path stroke=\"#fff\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10 9v4m0 0l-2-2m2 2l2-2\"/>" }, "folder-open": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5 19a2 2 0 01-2-2V7a2 2 0 012-2h4l2 2h4a2 2 0 012 2v1M5 19h14a2 2 0 002-2v-5a2 2 0 00-2-2H9a2 2 0 00-2 2v5a2 2 0 01-2 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M2 6a2 2 0 012-2h4l2 2h4a2 2 0 012 2v1H8a3 3 0 00-3 3v1.5a1.5 1.5 0 01-3 0V6z\" clip-rule=\"evenodd\"/><path d=\"M6 12a2 2 0 012-2h8a2 2 0 012 2v2a2 2 0 01-2 2H2h2a2 2 0 002-2v-2z\"/>" }, "folder-remove": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 13h6M3 17V7a2 2 0 012-2h6l2 2h6a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2z\"/>", "solid": "<path d=\"M2 6a2 2 0 012-2h5l2 2h5a2 2 0 012 2v6a2 2 0 01-2 2H4a2 2 0 01-2-2V6z\"/><path stroke=\"#fff\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 11h4\"/>" }, "folder": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z\"/>", "solid": "<path d=\"M2 6a2 2 0 012-2h5l2 2h5a2 2 0 012 2v6a2 2 0 01-2 2H4a2 2 0 01-2-2V6z\"/>" }, "gift": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 8v13m0-13V6a2 2 0 112 2h-2zm0 0V5.5A2.5 2.5 0 109.5 8H12zm-7 4h14M5 12a2 2 0 110-4h14a2 2 0 110 4M5 12v7a2 2 0 002 2h10a2 2 0 002-2v-7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5 5a3 3 0 015-2.236A3 3 0 0114.83 6H16a2 2 0 110 4h-5V9a1 1 0 10-2 0v1H4a2 2 0 110-4h1.17C5.06 5.687 5 5.35 5 5zm4 1V5a1 1 0 10-1 1h1zm3 0a1 1 0 10-1-1v1h1z\" clip-rule=\"evenodd\"/><path d=\"M9 11H3v5a2 2 0 002 2h4v-7zM11 18h4a2 2 0 002-2v-5h-6v7z\"/>" }, "globe-alt": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M4.083 9h1.946c.089-1.546.383-2.97.837-4.118A6.004 6.004 0 004.083 9zM10 2a8 8 0 100 16 8 8 0 000-16zm0 2c-.076 0-.232.032-.465.262-.238.234-.497.623-.737 1.182-.389.907-.673 2.142-.766 3.556h3.936c-.093-1.414-.377-2.649-.766-3.556-.24-.56-.5-.948-.737-1.182C10.232 4.032 10.076 4 10 4zm3.971 5c-.089-1.546-.383-2.97-.837-4.118A6.004 6.004 0 0115.917 9h-1.946zm-2.003 2H8.032c.093 1.414.377 2.649.766 3.556.24.56.5.948.737 1.182.233.23.389.262.465.262.076 0 .232-.032.465-.262.238-.234.498-.623.737-1.182.389-.907.673-2.142.766-3.556zm1.166 4.118c.454-1.147.748-2.572.837-4.118h1.946a6.004 6.004 0 01-2.783 4.118zm-6.268 0C6.412 13.97 6.118 12.546 6.03 11H4.083a6.004 6.004 0 002.783 4.118z\" clip-rule=\"evenodd\"/>" }, "globe": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3.055 11H5a2 2 0 012 2v1a2 2 0 002 2 2 2 0 012 2v2.945M8 3.935V5.5A2.5 2.5 0 0010.5 8h.5a2 2 0 012 2 2 2 0 104 0 2 2 0 012-2h1.064M15 20.488V18a2 2 0 012-2h3.064M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zM4.332 8.027a6.012 6.012 0 011.912-2.706C6.512 5.73 6.974 6 7.5 6A1.5 1.5 0 019 7.5V8a2 2 0 004 0 2 2 0 011.523-1.943A5.977 5.977 0 0116 10c0 .34-.028.675-.083 1H15a2 2 0 00-2 2v2.197A5.973 5.973 0 0110 16v-2a2 2 0 00-2-2 2 2 0 01-2-2 2 2 0 00-1.668-1.973z\" clip-rule=\"evenodd\"/>" }, "hand": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 11.5V14m0-2.5v-6a1.5 1.5 0 113 0m-3 6a1.5 1.5 0 00-3 0v2a7.5 7.5 0 0015 0v-5a1.5 1.5 0 00-3 0m-6-3V11m0-5.5v-1a1.5 1.5 0 013 0v1m0 0V11m0-5.5a1.5 1.5 0 013 0v3m0 0V11\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M9 3a1 1 0 012 0v5.5a.5.5 0 001 0V4a1 1 0 112 0v4.5a.5.5 0 001 0V6a1 1 0 112 0v5a7 7 0 11-14 0V9a1 1 0 012 0v2.5a.5.5 0 001 0V4a1 1 0 012 0v4.5a.5.5 0 001 0V3z\" clip-rule=\"evenodd\"/>" }, "hashtag": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 20l4-16m2 16l4-16M6 9h14M4 15h14\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M9.243 3.03a1 1 0 01.727 1.213L9.53 6h2.94l.56-2.243a1 1 0 111.94.486L14.53 6H17a1 1 0 110 2h-2.97l-1 4H15a1 1 0 110 2h-2.47l-.56 2.242a1 1 0 11-1.94-.485L10.47 14H7.53l-.56 2.242a1 1 0 11-1.94-.485L5.47 14H3a1 1 0 110-2h2.97l1-4H5a1 1 0 110-2h2.47l.56-2.243a1 1 0 011.213-.727zM9.03 8l-1 4h2.938l1-4H9.031z\" clip-rule=\"evenodd\"/>" }, "heart": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z\" clip-rule=\"evenodd\"/>" }, "home": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6\"/>", "solid": "<path d=\"M10.707 2.293a1 1 0 00-1.414 0l-7 7a1 1 0 001.414 1.414L4 10.414V17a1 1 0 001 1h2a1 1 0 001-1v-2a1 1 0 011-1h2a1 1 0 011 1v2a1 1 0 001 1h2a1 1 0 001-1v-6.586l.293.293a1 1 0 001.414-1.414l-7-7z\"/>" }, "identification": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10 6H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V8a2 2 0 00-2-2h-5m-4 0V5a2 2 0 114 0v1m-4 0a2 2 0 104 0m-5 8a2 2 0 100-4 2 2 0 000 4zm0 0c1.306 0 2.417.835 2.83 2M9 14a3.001 3.001 0 00-2.83 2M15 11h3m-3 4h2\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 2a1 1 0 00-1 1v1a1 1 0 002 0V3a1 1 0 00-1-1zM4 4h3a3 3 0 006 0h3a2 2 0 012 2v9a2 2 0 01-2 2H4a2 2 0 01-2-2V6a2 2 0 012-2zm2.5 7a1.5 1.5 0 100-3 1.5 1.5 0 000 3zm2.45 4a2.5 2.5 0 10-4.9 0h4.9zM12 9a1 1 0 100 2h3a1 1 0 100-2h-3zm-1 4a1 1 0 011-1h2a1 1 0 110 2h-2a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/>" }, "inbox-in": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 4H6a2 2 0 00-2 2v12a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-2m-4-1v8m0 0l3-3m-3 3L9 8m-5 5h2.586a1 1 0 01.707.293l2.414 2.414a1 1 0 00.707.293h3.172a1 1 0 00.707-.293l2.414-2.414a1 1 0 01.707-.293H20\"/>", "solid": "<path d=\"M8.707 7.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l2-2a1 1 0 00-1.414-1.414L11 7.586V3a1 1 0 10-2 0v4.586l-.293-.293z\"/><path d=\"M3 5a2 2 0 012-2h1a1 1 0 010 2H5v7h2l1 2h4l1-2h2V5h-1a1 1 0 110-2h1a2 2 0 012 2v10a2 2 0 01-2 2H5a2 2 0 01-2-2V5z\"/>" }, "inbox": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5 3a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2V5a2 2 0 00-2-2H5zm0 2h10v7h-2l-1 2H8l-1-2H5V5z\" clip-rule=\"evenodd\"/>" }, "information-circle": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z\" clip-rule=\"evenodd\"/>" }, "key": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M18 8a6 6 0 01-7.743 5.743L10 14l-1 1-1 1H6v2H2v-4l4.257-4.257A6 6 0 1118 8zm-6-4a1 1 0 100 2 2 2 0 012 2 1 1 0 102 0 4 4 0 00-4-4z\" clip-rule=\"evenodd\"/>" }, "library": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 14v3m4-3v3m4-3v3M3 21h18M3 10h18M3 7l9-4 9 4M4 10h16v11H4V10z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10.496 2.132a1 1 0 00-.992 0l-7 4A1 1 0 003 8v7a1 1 0 100 2h14a1 1 0 100-2V8a1 1 0 00.496-1.868l-7-4zM6 9a1 1 0 00-1 1v3a1 1 0 102 0v-3a1 1 0 00-1-1zm3 1a1 1 0 012 0v3a1 1 0 11-2 0v-3zm5-1a1 1 0 00-1 1v3a1 1 0 102 0v-3a1 1 0 00-1-1z\" clip-rule=\"evenodd\"/>" }, "light-bulb": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z\"/>", "solid": "<path d=\"M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.476.859h4.002z\"/>" }, "lightning-bolt": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 10V3L4 14h7v7l9-11h-7z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z\" clip-rule=\"evenodd\"/>" }, "link": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M12.586 4.586a2 2 0 112.828 2.828l-3 3a2 2 0 01-2.828 0 1 1 0 00-1.414 1.414 4 4 0 005.656 0l3-3a4 4 0 00-5.656-5.656l-1.5 1.5a1 1 0 101.414 1.414l1.5-1.5zm-5 5a2 2 0 012.828 0 1 1 0 101.414-1.414 4 4 0 00-5.656 0l-3 3a4 4 0 105.656 5.656l1.5-1.5a1 1 0 10-1.414-1.414l-1.5 1.5a2 2 0 11-2.828-2.828l3-3z\" clip-rule=\"evenodd\"/>" }, "location-marker": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z\"/><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 11a3 3 0 11-6 0 3 3 0 016 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z\" clip-rule=\"evenodd\"/>" }, "lock-closed": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z\" clip-rule=\"evenodd\"/>" }, "lock-open": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 11V7a4 4 0 118 0m-4 8v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2z\"/>", "solid": "<path d=\"M10 2a5 5 0 00-5 5v2a2 2 0 00-2 2v5a2 2 0 002 2h10a2 2 0 002-2v-5a2 2 0 00-2-2H7V7a3 3 0 015.905-.75 1 1 0 001.937-.5A5.002 5.002 0 0010 2z\"/>" }, "login": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 3a1 1 0 011 1v12a1 1 0 11-2 0V4a1 1 0 011-1zm7.707 3.293a1 1 0 010 1.414L9.414 9H17a1 1 0 110 2H9.414l1.293 1.293a1 1 0 01-1.414 1.414l-3-3a1 1 0 010-1.414l3-3a1 1 0 011.414 0z\" clip-rule=\"evenodd\"/>" }, "logout": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 3a1 1 0 00-1 1v12a1 1 0 102 0V4a1 1 0 00-1-1zm10.293 9.293a1 1 0 001.414 1.414l3-3a1 1 0 000-1.414l-3-3a1 1 0 10-1.414 1.414L14.586 9H7a1 1 0 100 2h7.586l-1.293 1.293z\" clip-rule=\"evenodd\"/>" }, "mail-open": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 19v-8.93a2 2 0 01.89-1.664l7-4.666a2 2 0 012.22 0l7 4.666A2 2 0 0121 10.07V19M3 19a2 2 0 002 2h14a2 2 0 002-2M3 19l6.75-4.5M21 19l-6.75-4.5M3 10l6.75 4.5M21 10l-6.75 4.5m0 0l-1.14.76a2 2 0 01-2.22 0l-1.14-.76\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M2.94 6.412A2 2 0 002 8.108V16a2 2 0 002 2h12a2 2 0 002-2V8.108a2 2 0 00-.94-1.696l-6-3.75a2 2 0 00-2.12 0l-6 3.75zm2.615 2.423a1 1 0 10-1.11 1.664l5 3.333a1 1 0 001.11 0l5-3.333a1 1 0 00-1.11-1.664L10 11.798 5.555 8.835z\" clip-rule=\"evenodd\"/>" }, "mail": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z\"/>", "solid": "<path d=\"M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z\"/><path d=\"M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z\"/>" }, "map": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M12 1.586l-4 4v12.828l4-4V1.586zM3.707 3.293A1 1 0 002 4v10a1 1 0 00.293.707L6 18.414V5.586L3.707 3.293zM17.707 5.293L14 1.586v12.828l2.293 2.293A1 1 0 0018 16V6a1 1 0 00-.293-.707z\" clip-rule=\"evenodd\"/>" }, "menu-alt-1": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h8m-8 6h16\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h6a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/>" }, "menu-alt-2": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16M4 18h7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h6a1 1 0 110 2H4a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/>" }, "menu-alt-3": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16m-7 6h7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM9 15a1 1 0 011-1h6a1 1 0 110 2h-6a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/>" }, "menu-alt-4": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 8h16M4 16h16\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 7a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 13a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/>" }, "menu": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16M4 18h16\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/>" }, "microphone": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M19 11a7 7 0 01-7 7m0 0a7 7 0 01-7-7m7 7v4m0 0H8m4 0h4m-4-8a3 3 0 01-3-3V5a3 3 0 116 0v6a3 3 0 01-3 3z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M7 4a3 3 0 016 0v4a3 3 0 11-6 0V4zm4 10.93A7.001 7.001 0 0017 8a1 1 0 10-2 0A5 5 0 015 8a1 1 0 00-2 0 7.001 7.001 0 006 6.93V17H6a1 1 0 100 2h8a1 1 0 100-2h-3v-2.07z\" clip-rule=\"evenodd\"/>" }, "minus-circle": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 12H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zM7 9a1 1 0 000 2h6a1 1 0 100-2H7z\" clip-rule=\"evenodd\"/>" }, "minus-sm": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M18 12H6\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5 10a1 1 0 011-1h8a1 1 0 110 2H6a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/>" }, "minus": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M20 12H4\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/>" }, "moon": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z\"/>", "solid": "<path d=\"M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z\"/>" }, "music-note": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 19V6l12-3v13M9 19c0 1.105-1.343 2-3 2s-3-.895-3-2 1.343-2 3-2 3 .895 3 2zm12-3c0 1.105-1.343 2-3 2s-3-.895-3-2 1.343-2 3-2 3 .895 3 2zM9 10l12-3\"/>", "solid": "<path d=\"M18 3a1 1 0 00-1.196-.98l-10 2A1 1 0 006 5v9.114A4.369 4.369 0 005 14c-1.657 0-3 .895-3 2s1.343 2 3 2 3-.895 3-2V7.82l8-1.6v5.894A4.37 4.37 0 0015 12c-1.657 0-3 .895-3 2s1.343 2 3 2 3-.895 3-2V3z\"/>" }, "newspaper": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9M7 16h6M7 8h6v4H7V8z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M2 5a2 2 0 012-2h8a2 2 0 012 2v10a2 2 0 002 2H4a2 2 0 01-2-2V5zm3 1h6v4H5V6zm6 6H5v2h6v-2z\" clip-rule=\"evenodd\"/><path d=\"M15 7h1a2 2 0 012 2v5.5a1.5 1.5 0 01-3 0V7z\"/>" }, "office-building": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M4 4a2 2 0 012-2h8a2 2 0 012 2v12a1 1 0 110 2h-3a1 1 0 01-1-1v-2a1 1 0 00-1-1H9a1 1 0 00-1 1v2a1 1 0 01-1 1H4a1 1 0 110-2V4zm3 1h2v2H7V5zm2 4H7v2h2V9zm2-4h2v2h-2V5zm2 4h-2v2h2V9z\" clip-rule=\"evenodd\"/>" }, "paper-airplane": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 19l9 2-9-18-9 18 9-2zm0 0v-8\"/>", "solid": "<path d=\"M10.894 2.553a1 1 0 00-1.788 0l-7 14a1 1 0 001.169 1.409l5-1.429A1 1 0 009 15.571V11a1 1 0 112 0v4.571a1 1 0 00.725.962l5 1.428a1 1 0 001.17-1.408l-7-14z\"/>" }, "paper-clip": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M8 4a3 3 0 00-3 3v4a5 5 0 0010 0V7a1 1 0 112 0v4a7 7 0 11-14 0V7a5 5 0 0110 0v4a3 3 0 11-6 0V7a1 1 0 012 0v4a1 1 0 102 0V7a3 3 0 00-3-3z\" clip-rule=\"evenodd\"/>" }, "pause": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10 9v6m4-6v6m7-3a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M18 10a8 8 0 11-16 0 8 8 0 0116 0zM7 8a1 1 0 012 0v4a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v4a1 1 0 102 0V8a1 1 0 00-1-1z\" clip-rule=\"evenodd\"/>" }, "pencil-alt": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z\"/>", "solid": "<path d=\"M17.414 2.586a2 2 0 00-2.828 0L7 10.172V13h2.828l7.586-7.586a2 2 0 000-2.828z\"/><path fill-rule=\"evenodd\" d=\"M2 6a2 2 0 012-2h4a1 1 0 010 2H4v10h10v-4a1 1 0 112 0v4a2 2 0 01-2 2H4a2 2 0 01-2-2V6z\" clip-rule=\"evenodd\"/>" }, "pencil": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z\"/>", "solid": "<path d=\"M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z\"/>" }, "phone-incoming": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M21 3l-6 6m0 0V4m0 5h5M5 3a2 2 0 00-2 2v1c0 8.284 6.716 15 15 15h1a2 2 0 002-2v-3.28a1 1 0 00-.684-.948l-4.493-1.498a1 1 0 00-1.21.502l-1.13 2.257a11.042 11.042 0 01-5.516-5.517l2.257-1.128a1 1 0 00.502-1.21L9.228 3.683A1 1 0 008.279 3H5z\"/>", "solid": "<path d=\"M14.414 7l3.293-3.293a1 1 0 00-1.414-1.414L13 5.586V4a1 1 0 10-2 0v4.003a.996.996 0 00.617.921A.997.997 0 0012 9h4a1 1 0 100-2h-1.586z\"/><path d=\"M2 3a1 1 0 011-1h2.153a1 1 0 01.986.836l.74 4.435a1 1 0 01-.54 1.06l-1.548.773a11.037 11.037 0 006.105 6.105l.774-1.548a1 1 0 011.059-.54l4.435.74a1 1 0 01.836.986V17a1 1 0 01-1 1h-2C7.82 18 2 12.18 2 5V3z\"/>" }, "phone-missed-call": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M16 8l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2M5 3a2 2 0 00-2 2v1c0 8.284 6.716 15 15 15h1a2 2 0 002-2v-3.28a1 1 0 00-.684-.948l-4.493-1.498a1 1 0 00-1.21.502l-1.13 2.257a11.042 11.042 0 01-5.516-5.517l2.257-1.128a1 1 0 00.502-1.21L9.228 3.683A1 1 0 008.279 3H5z\"/>", "solid": "<path d=\"M2 3a1 1 0 011-1h2.153a1 1 0 01.986.836l.74 4.435a1 1 0 01-.54 1.06l-1.548.773a11.037 11.037 0 006.105 6.105l.774-1.548a1 1 0 011.059-.54l4.435.74a1 1 0 01.836.986V17a1 1 0 01-1 1h-2C7.82 18 2 12.18 2 5V3z\"/><path d=\"M16.707 3.293a1 1 0 010 1.414L15.414 6l1.293 1.293a1 1 0 01-1.414 1.414L14 7.414l-1.293 1.293a1 1 0 11-1.414-1.414L12.586 6l-1.293-1.293a1 1 0 011.414-1.414L14 4.586l1.293-1.293a1 1 0 011.414 0z\"/>" }, "phone-outgoing": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M16 3h5m0 0v5m0-5l-6 6M5 3a2 2 0 00-2 2v1c0 8.284 6.716 15 15 15h1a2 2 0 002-2v-3.28a1 1 0 00-.684-.948l-4.493-1.498a1 1 0 00-1.21.502l-1.13 2.257a11.042 11.042 0 01-5.516-5.517l2.257-1.128a1 1 0 00.502-1.21L9.228 3.683A1 1 0 008.279 3H5z\"/>", "solid": "<path d=\"M17.924 2.617a.997.997 0 00-.215-.322l-.004-.004A.997.997 0 0017 2h-4a1 1 0 100 2h1.586l-3.293 3.293a1 1 0 001.414 1.414L16 5.414V7a1 1 0 102 0V3a.997.997 0 00-.076-.383z\"/><path d=\"M2 3a1 1 0 011-1h2.153a1 1 0 01.986.836l.74 4.435a1 1 0 01-.54 1.06l-1.548.773a11.037 11.037 0 006.105 6.105l.774-1.548a1 1 0 011.059-.54l4.435.74a1 1 0 01.836.986V17a1 1 0 01-1 1h-2C7.82 18 2 12.18 2 5V3z\"/>" }, "phone": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z\"/>", "solid": "<path d=\"M2 3a1 1 0 011-1h2.153a1 1 0 01.986.836l.74 4.435a1 1 0 01-.54 1.06l-1.548.773a11.037 11.037 0 006.105 6.105l.774-1.548a1 1 0 011.059-.54l4.435.74a1 1 0 01.836.986V17a1 1 0 01-1 1h-2C7.82 18 2 12.18 2 5V3z\"/>" }, "photograph": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M4 3a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V5a2 2 0 00-2-2H4zm12 12H4l4-8 3 6 2-4 3 6z\" clip-rule=\"evenodd\"/>" }, "play": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z\"/><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z\" clip-rule=\"evenodd\"/>" }, "plus-circle": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 9v3m0 0v3m0-3h3m-3 0H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zm1-11a1 1 0 10-2 0v2H7a1 1 0 100 2h2v2a1 1 0 102 0v-2h2a1 1 0 100-2h-2V7z\" clip-rule=\"evenodd\"/>" }, "plus-sm": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 6v6m0 0v6m0-6h6m-6 0H6\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z\" clip-rule=\"evenodd\"/>" }, "plus": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 4v16m8-8H4\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z\" clip-rule=\"evenodd\"/>" }, "presentation-chart-bar": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 13v-1m4 1v-3m4 3V8M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11 4a1 1 0 10-2 0v4a1 1 0 102 0V7zm-3 1a1 1 0 10-2 0v3a1 1 0 102 0V8zM8 9a1 1 0 00-2 0v2a1 1 0 102 0V9z\" clip-rule=\"evenodd\"/>" }, "presentation-chart-line": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11.707 4.707a1 1 0 00-1.414-1.414L10 9.586 8.707 8.293a1 1 0 00-1.414 0l-2 2a1 1 0 101.414 1.414L8 10.414l1.293 1.293a1 1 0 001.414 0l4-4z\" clip-rule=\"evenodd\"/>" }, "printer": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5 4v3H4a2 2 0 00-2 2v3a2 2 0 002 2h1v2a2 2 0 002 2h6a2 2 0 002-2v-2h1a2 2 0 002-2V9a2 2 0 00-2-2h-1V4a2 2 0 00-2-2H7a2 2 0 00-2 2zm8 0H7v3h6V4zm0 8H7v4h6v-4z\" clip-rule=\"evenodd\"/>" }, "puzzle": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M11 4a2 2 0 114 0v1a1 1 0 001 1h3a1 1 0 011 1v3a1 1 0 01-1 1h-1a2 2 0 100 4h1a1 1 0 011 1v3a1 1 0 01-1 1h-3a1 1 0 01-1-1v-1a2 2 0 10-4 0v1a1 1 0 01-1 1H7a1 1 0 01-1-1v-3a1 1 0 00-1-1H4a2 2 0 110-4h1a1 1 0 001-1V7a1 1 0 011-1h3a1 1 0 001-1V4z\"/>", "solid": "<path d=\"M10 3.5a1.5 1.5 0 013 0V4a1 1 0 001 1h3a1 1 0 011 1v3a1 1 0 01-1 1h-.5a1.5 1.5 0 000 3h.5a1 1 0 011 1v3a1 1 0 01-1 1h-3a1 1 0 01-1-1v-.5a1.5 1.5 0 00-3 0v.5a1 1 0 01-1 1H6a1 1 0 01-1-1v-3a1 1 0 00-1-1h-.5a1.5 1.5 0 010-3H4a1 1 0 001-1V6a1 1 0 011-1h3a1 1 0 001-1v-.5z\"/>" }, "qrcode": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 4v1m6 11h2m-6 0h-2v4m0-11v3m0 0h.01M12 12h4.01M16 20h4M4 12h4m12 0h.01M5 8h2a1 1 0 001-1V5a1 1 0 00-1-1H5a1 1 0 00-1 1v2a1 1 0 001 1zm12 0h2a1 1 0 001-1V5a1 1 0 00-1-1h-2a1 1 0 00-1 1v2a1 1 0 001 1zM5 20h2a1 1 0 001-1v-2a1 1 0 00-1-1H5a1 1 0 00-1 1v2a1 1 0 001 1z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 4a1 1 0 011-1h3a1 1 0 011 1v3a1 1 0 01-1 1H4a1 1 0 01-1-1V4zm2 2V5h1v1H5zM3 13a1 1 0 011-1h3a1 1 0 011 1v3a1 1 0 01-1 1H4a1 1 0 01-1-1v-3zm2 2v-1h1v1H5zM13 3a1 1 0 00-1 1v3a1 1 0 001 1h3a1 1 0 001-1V4a1 1 0 00-1-1h-3zm1 2v1h1V5h-1z\" clip-rule=\"evenodd\"/><path d=\"M11 4a1 1 0 10-2 0v1a1 1 0 002 0V4zM10 7a1 1 0 011 1v1h2a1 1 0 110 2h-3a1 1 0 01-1-1V8a1 1 0 011-1zM16 9a1 1 0 100 2 1 1 0 000-2zM9 13a1 1 0 011-1h1a1 1 0 110 2v2a1 1 0 11-2 0v-3zM7 11a1 1 0 100-2H4a1 1 0 100 2h3zM17 13a1 1 0 01-1 1h-2a1 1 0 110-2h2a1 1 0 011 1zM16 17a1 1 0 100-2h-3a1 1 0 100 2h3z\"/>" }, "question-mark-circle": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-8-3a1 1 0 00-.867.5 1 1 0 11-1.731-1A3 3 0 0113 8a3.001 3.001 0 01-2 2.83V11a1 1 0 11-2 0v-1a1 1 0 011-1 1 1 0 100-2zm0 8a1 1 0 100-2 1 1 0 000 2z\" clip-rule=\"evenodd\"/>" }, "receipt-refund": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M16 15v-1a4 4 0 00-4-4H8m0 0l3 3m-3-3l3-3m9 14V5a2 2 0 00-2-2H6a2 2 0 00-2 2v16l4-2 4 2 4-2 4 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5 2a2 2 0 00-2 2v14l3.5-2 3.5 2 3.5-2 3.5 2V4a2 2 0 00-2-2H5zm4.707 3.707a1 1 0 00-1.414-1.414l-3 3a1 1 0 000 1.414l3 3a1 1 0 001.414-1.414L8.414 9H10a3 3 0 013 3v1a1 1 0 102 0v-1a5 5 0 00-5-5H8.414l1.293-1.293z\" clip-rule=\"evenodd\"/>" }, "receipt-tax": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 14l6-6m-5.5.5h.01m4.99 5h.01M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16l3.5-2 3.5 2 3.5-2 3.5 2zM10 8.5a.5.5 0 11-1 0 .5.5 0 011 0zm5 5a.5.5 0 11-1 0 .5.5 0 011 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5 2a2 2 0 00-2 2v14l3.5-2 3.5 2 3.5-2 3.5 2V4a2 2 0 00-2-2H5zm2.5 3a1.5 1.5 0 100 3 1.5 1.5 0 000-3zm6.207.293a1 1 0 00-1.414 0l-6 6a1 1 0 101.414 1.414l6-6a1 1 0 000-1.414zM12.5 10a1.5 1.5 0 100 3 1.5 1.5 0 000-3z\" clip-rule=\"evenodd\"/>" }, "refresh": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z\" clip-rule=\"evenodd\"/>" }, "reply": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M7.707 3.293a1 1 0 010 1.414L5.414 7H11a7 7 0 017 7v2a1 1 0 11-2 0v-2a5 5 0 00-5-5H5.414l2.293 2.293a1 1 0 11-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z\" clip-rule=\"evenodd\"/>" }, "rewind": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12.066 11.2a1 1 0 000 1.6l5.334 4A1 1 0 0019 16V8a1 1 0 00-1.6-.8l-5.333 4zM4.066 11.2a1 1 0 000 1.6l5.334 4A1 1 0 0011 16V8a1 1 0 00-1.6-.8l-5.334 4z\"/>", "solid": "<path d=\"M8.445 14.832A1 1 0 0010 14v-2.798l5.445 3.63A1 1 0 0017 14V6a1 1 0 00-1.555-.832L10 8.798V6a1 1 0 00-1.555-.832l-6 4a1 1 0 000 1.664l6 4z\"/>" }, "rss": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M6 5c7.18 0 13 5.82 13 13M6 11a7 7 0 017 7m-6 0a1 1 0 11-2 0 1 1 0 012 0z\"/>", "solid": "<path d=\"M5 3a1 1 0 000 2c5.523 0 10 4.477 10 10a1 1 0 102 0C17 8.373 11.627 3 5 3z\"/><path d=\"M4 9a1 1 0 011-1 7 7 0 017 7 1 1 0 11-2 0 5 5 0 00-5-5 1 1 0 01-1-1zM3 15a2 2 0 114 0 2 2 0 01-4 0z\"/>" }, "save-as": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M17 16v2a2 2 0 01-2 2H5a2 2 0 01-2-2v-7a2 2 0 012-2h2m3-4H9a2 2 0 00-2 2v7a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-1m-1 4l-3 3m0 0l-3-3m3 3V3\"/>", "solid": "<path d=\"M9.707 7.293a1 1 0 00-1.414 1.414l3 3a1 1 0 001.414 0l3-3a1 1 0 00-1.414-1.414L13 8.586V5h3a2 2 0 012 2v5a2 2 0 01-2 2H8a2 2 0 01-2-2V7a2 2 0 012-2h3v3.586L9.707 7.293zM11 3a1 1 0 112 0v2h-2V3z\"/><path d=\"M4 9a2 2 0 00-2 2v5a2 2 0 002 2h8a2 2 0 002-2H4V9z\"/>" }, "save": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4\"/>", "solid": "<path d=\"M7.707 10.293a1 1 0 10-1.414 1.414l3 3a1 1 0 001.414 0l3-3a1 1 0 00-1.414-1.414L11 11.586V6h5a2 2 0 012 2v7a2 2 0 01-2 2H4a2 2 0 01-2-2V8a2 2 0 012-2h5v5.586l-1.293-1.293zM9 4a1 1 0 012 0v2H9V4z\"/>" }, "scale": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 6l3 1m0 0l-3 9a5.002 5.002 0 006.001 0M6 7l3 9M6 7l6-2m6 2l3-1m-3 1l-3 9a5.002 5.002 0 006.001 0M18 7l3 9m-3-9l-6-2m0-2v2m0 16V5m0 16H9m3 0h3\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 2a1 1 0 011 1v1.323l3.954 1.582 1.599-.8a1 1 0 01.894 1.79l-1.233.616 1.738 5.42a1 1 0 01-.285 1.05A3.989 3.989 0 0115 15a3.989 3.989 0 01-2.667-1.019 1 1 0 01-.285-1.05l1.715-5.349L11 6.477V16h2a1 1 0 110 2H7a1 1 0 110-2h2V6.477L6.237 7.582l1.715 5.349a1 1 0 01-.285 1.05A3.989 3.989 0 015 15a3.989 3.989 0 01-2.667-1.019 1 1 0 01-.285-1.05l1.738-5.42-1.233-.617a1 1 0 01.894-1.788l1.599.799L9 4.323V3a1 1 0 011-1zm-5 8.274l-.818 2.552c.25.112.526.174.818.174.292 0 .569-.062.818-.174L5 10.274zm10 0l-.818 2.552c.25.112.526.174.818.174.292 0 .569-.062.818-.174L15 10.274z\" clip-rule=\"evenodd\"/>" }, "scissors": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M14.121 14.121L19 19m-7-7l7-7m-7 7l-2.879 2.879M12 12L9.121 9.121m0 5.758a3 3 0 10-4.243 4.243 3 3 0 004.243-4.243zm0-5.758a3 3 0 10-4.243-4.243 3 3 0 004.243 4.243z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5.5 2a3.5 3.5 0 101.665 6.58L8.585 10l-1.42 1.42a3.5 3.5 0 101.414 1.414l8.128-8.127a1 1 0 00-1.414-1.414L10 8.586l-1.42-1.42A3.5 3.5 0 005.5 2zM4 5.5a1.5 1.5 0 113 0 1.5 1.5 0 01-3 0zm0 9a1.5 1.5 0 113 0 1.5 1.5 0 01-3 0z\" clip-rule=\"evenodd\"/><path d=\"M12.828 11.414a1 1 0 00-1.414 1.414l3.879 3.88a1 1 0 001.414-1.415l-3.879-3.879z\"/>" }, "search-circle": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 16l2.879-2.879m0 0a3 3 0 104.243-4.242 3 3 0 00-4.243 4.242zM21 12a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path d=\"M9 9a2 2 0 114 0 2 2 0 01-4 0z\"/><path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zm1-13a4 4 0 00-3.446 6.032l-2.261 2.26a1 1 0 101.414 1.415l2.261-2.261A4 4 0 1011 5z\" clip-rule=\"evenodd\"/>" }, "search": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z\" clip-rule=\"evenodd\"/>" }, "selector": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 9l4-4 4 4m0 6l-4 4-4-4\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 3a1 1 0 01.707.293l3 3a1 1 0 01-1.414 1.414L10 5.414 7.707 7.707a1 1 0 01-1.414-1.414l3-3A1 1 0 0110 3zm-3.707 9.293a1 1 0 011.414 0L10 14.586l2.293-2.293a1 1 0 011.414 1.414l-3 3a1 1 0 01-1.414 0l-3-3a1 1 0 010-1.414z\" clip-rule=\"evenodd\"/>" }, "server": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M2 5a2 2 0 012-2h12a2 2 0 012 2v2a2 2 0 01-2 2H4a2 2 0 01-2-2V5zm14 1a1 1 0 11-2 0 1 1 0 012 0zM2 13a2 2 0 012-2h12a2 2 0 012 2v2a2 2 0 01-2 2H4a2 2 0 01-2-2v-2zm14 1a1 1 0 11-2 0 1 1 0 012 0z\" clip-rule=\"evenodd\"/>" }, "share": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z\"/>", "solid": "<path d=\"M15 8a3 3 0 10-2.977-2.63l-4.94 2.47a3 3 0 100 4.319l4.94 2.47a3 3 0 10.895-1.789l-4.94-2.47a3.027 3.027 0 000-.74l4.94-2.47C13.456 7.68 14.19 8 15 8z\"/>" }, "shield-check": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M2.166 4.999A11.954 11.954 0 0010 1.944 11.954 11.954 0 0017.834 5c.11.65.166 1.32.166 2.001 0 5.225-3.34 9.67-8 11.317C5.34 16.67 2 12.225 2 7c0-.682.057-1.35.166-2.001zm11.541 3.708a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z\" clip-rule=\"evenodd\"/>" }, "shield-exclamation": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M20.618 5.984A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016zM12 9v2m0 4h.01\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 1.944A11.954 11.954 0 012.166 5C2.056 5.649 2 6.319 2 7c0 5.225 3.34 9.67 8 11.317C14.66 16.67 18 12.225 18 7c0-.682-.057-1.35-.166-2.001A11.954 11.954 0 0110 1.944zM11 14a1 1 0 11-2 0 1 1 0 012 0zm0-7a1 1 0 10-2 0v3a1 1 0 102 0V7z\" clip-rule=\"evenodd\"/>" }, "shopping-bag": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 2a4 4 0 00-4 4v1H5a1 1 0 00-.994.89l-1 9A1 1 0 004 18h12a1 1 0 00.994-1.11l-1-9A1 1 0 0015 7h-1V6a4 4 0 00-4-4zm2 5V6a2 2 0 10-4 0v1h4zm-6 3a1 1 0 112 0 1 1 0 01-2 0zm7-1a1 1 0 100 2 1 1 0 000-2z\" clip-rule=\"evenodd\"/>" }, "shopping-cart": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z\"/>", "solid": "<path d=\"M3 1a1 1 0 000 2h1.22l.305 1.222a.997.997 0 00.01.042l1.358 5.43-.893.892C3.74 11.846 4.632 14 6.414 14H15a1 1 0 000-2H6.414l1-1H14a1 1 0 00.894-.553l3-6A1 1 0 0017 3H6.28l-.31-1.243A1 1 0 005 1H3zM16 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zM6.5 18a1.5 1.5 0 100-3 1.5 1.5 0 000 3z\"/>" }, "sort-ascending": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 4h13M3 8h9m-9 4h6m4 0l4-4m0 0l4 4m-4-4v12\"/>", "solid": "<path d=\"M3 3a1 1 0 000 2h11a1 1 0 100-2H3zM3 7a1 1 0 000 2h5a1 1 0 000-2H3zM3 11a1 1 0 100 2h4a1 1 0 100-2H3zM13 16a1 1 0 102 0v-5.586l1.293 1.293a1 1 0 001.414-1.414l-3-3a1 1 0 00-1.414 0l-3 3a1 1 0 101.414 1.414L13 10.414V16z\"/>" }, "sort-descending": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 4h13M3 8h9m-9 4h9m5-4v12m0 0l-4-4m4 4l4-4\"/>", "solid": "<path d=\"M3 3a1 1 0 000 2h11a1 1 0 100-2H3zM3 7a1 1 0 000 2h7a1 1 0 100-2H3zM3 11a1 1 0 100 2h4a1 1 0 100-2H3zM15 8a1 1 0 10-2 0v5.586l-1.293-1.293a1 1 0 00-1.414 1.414l3 3a1 1 0 001.414 0l3-3a1 1 0 00-1.414-1.414L15 13.586V8z\"/>" }, "sparkles": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5 2a1 1 0 011 1v1h1a1 1 0 010 2H6v1a1 1 0 01-2 0V6H3a1 1 0 010-2h1V3a1 1 0 011-1zm0 10a1 1 0 011 1v1h1a1 1 0 110 2H6v1a1 1 0 11-2 0v-1H3a1 1 0 110-2h1v-1a1 1 0 011-1zM12 2a1 1 0 01.967.744L14.146 7.2 17.5 9.134a1 1 0 010 1.732l-3.354 1.935-1.18 4.455a1 1 0 01-1.933 0L9.854 12.8 6.5 10.866a1 1 0 010-1.732l3.354-1.935 1.18-4.455A1 1 0 0112 2z\" clip-rule=\"evenodd\"/>" }, "speakerphone": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M11 5.882V19.24a1.76 1.76 0 01-3.417.592l-2.147-6.15M18 13a3 3 0 100-6M5.436 13.683A4.001 4.001 0 017 6h1.832c4.1 0 7.625-1.234 9.168-3v14c-1.543-1.766-5.067-3-9.168-3H7a3.988 3.988 0 01-1.564-.317z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M18 3a1 1 0 00-1.447-.894L8.763 6H5a3 3 0 000 6h.28l1.771 5.316A1 1 0 008 18h1a1 1 0 001-1v-4.382l6.553 3.276A1 1 0 0018 15V3z\" clip-rule=\"evenodd\"/>" }, "star": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z\"/>", "solid": "<path d=\"M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z\"/>" }, "status-offline": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M18.364 5.636a9 9 0 010 12.728m0 0l-2.829-2.829m2.829 2.829L21 21M15.536 8.464a5 5 0 010 7.072m0 0l-2.829-2.829m-4.243 2.829a4.978 4.978 0 01-1.414-2.83m-1.414 5.658a9 9 0 01-2.167-9.238m7.824 2.167a1 1 0 111.414 1.414m-1.414-1.414L3 3m8.293 8.293l1.414 1.414\"/>", "solid": "<path d=\"M3.707 2.293a1 1 0 00-1.414 1.414l6.921 6.922c.05.062.105.118.168.167l6.91 6.911a1 1 0 001.415-1.414l-.675-.675a9.001 9.001 0 00-.668-11.982A1 1 0 1014.95 5.05a7.002 7.002 0 01.657 9.143l-1.435-1.435a5.002 5.002 0 00-.636-6.294A1 1 0 0012.12 7.88c.924.923 1.12 2.3.587 3.415l-1.992-1.992a.922.922 0 00-.018-.018l-6.99-6.991zM3.238 8.187a1 1 0 00-1.933-.516c-.8 3-.025 6.336 2.331 8.693a1 1 0 001.414-1.415 6.997 6.997 0 01-1.812-6.762zM7.4 11.5a1 1 0 10-1.73 1c.214.371.48.72.795 1.035a1 1 0 001.414-1.414c-.191-.191-.35-.4-.478-.622z\"/>" }, "status-online": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5.636 18.364a9 9 0 010-12.728m12.728 0a9 9 0 010 12.728m-9.9-2.829a5 5 0 010-7.07m7.072 0a5 5 0 010 7.07M13 12a1 1 0 11-2 0 1 1 0 012 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5.05 3.636a1 1 0 010 1.414 7 7 0 000 9.9 1 1 0 11-1.414 1.414 9 9 0 010-12.728 1 1 0 011.414 0zm9.9 0a1 1 0 011.414 0 9 9 0 010 12.728 1 1 0 11-1.414-1.414 7 7 0 000-9.9 1 1 0 010-1.414zM7.879 6.464a1 1 0 010 1.414 3 3 0 000 4.243 1 1 0 11-1.415 1.414 5 5 0 010-7.07 1 1 0 011.415 0zm4.242 0a1 1 0 011.415 0 5 5 0 010 7.072 1 1 0 01-1.415-1.415 3 3 0 000-4.242 1 1 0 010-1.415zM10 9a1 1 0 011 1v.01a1 1 0 11-2 0V10a1 1 0 011-1z\" clip-rule=\"evenodd\"/>" }, "stop": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"/><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 10a1 1 0 011-1h4a1 1 0 011 1v4a1 1 0 01-1 1h-4a1 1 0 01-1-1v-4z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zM8 7a1 1 0 00-1 1v4a1 1 0 001 1h4a1 1 0 001-1V8a1 1 0 00-1-1H8z\" clip-rule=\"evenodd\"/>" }, "sun": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 2a1 1 0 011 1v1a1 1 0 11-2 0V3a1 1 0 011-1zm4 8a4 4 0 11-8 0 4 4 0 018 0zm-.464 4.95l.707.707a1 1 0 001.414-1.414l-.707-.707a1 1 0 00-1.414 1.414zm2.12-10.607a1 1 0 010 1.414l-.706.707a1 1 0 11-1.414-1.414l.707-.707a1 1 0 011.414 0zM17 11a1 1 0 100-2h-1a1 1 0 100 2h1zm-7 4a1 1 0 011 1v1a1 1 0 11-2 0v-1a1 1 0 011-1zM5.05 6.464A1 1 0 106.465 5.05l-.708-.707a1 1 0 00-1.414 1.414l.707.707zm1.414 8.486l-.707.707a1 1 0 01-1.414-1.414l.707-.707a1 1 0 011.414 1.414zM4 11a1 1 0 100-2H3a1 1 0 000 2h1z\" clip-rule=\"evenodd\"/>" }, "support": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M18.364 5.636l-3.536 3.536m0 5.656l3.536 3.536M9.172 9.172L5.636 5.636m3.536 9.192l-3.536 3.536M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-5 0a4 4 0 11-8 0 4 4 0 018 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-2 0c0 .993-.241 1.929-.668 2.754l-1.524-1.525a3.997 3.997 0 00.078-2.183l1.562-1.562C15.802 8.249 16 9.1 16 10zm-5.165 3.913l1.58 1.58A5.98 5.98 0 0110 16a5.976 5.976 0 01-2.516-.552l1.562-1.562a4.006 4.006 0 001.789.027zm-4.677-2.796a4.002 4.002 0 01-.041-2.08l-.08.08-1.53-1.533A5.98 5.98 0 004 10c0 .954.223 1.856.619 2.657l1.54-1.54zm1.088-6.45A5.974 5.974 0 0110 4c.954 0 1.856.223 2.657.619l-1.54 1.54a4.002 4.002 0 00-2.346.033L7.246 4.668zM12 10a2 2 0 11-4 0 2 2 0 014 0z\" clip-rule=\"evenodd\"/>" }, "switch-horizontal": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4\"/>", "solid": "<path d=\"M8 5a1 1 0 100 2h5.586l-1.293 1.293a1 1 0 001.414 1.414l3-3a1 1 0 000-1.414l-3-3a1 1 0 10-1.414 1.414L13.586 5H8zM12 15a1 1 0 100-2H6.414l1.293-1.293a1 1 0 10-1.414-1.414l-3 3a1 1 0 000 1.414l3 3a1 1 0 001.414-1.414L6.414 15H12z\"/>" }, "switch-vertical": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4\"/>", "solid": "<path d=\"M5 12a1 1 0 102 0V6.414l1.293 1.293a1 1 0 001.414-1.414l-3-3a1 1 0 00-1.414 0l-3 3a1 1 0 001.414 1.414L5 6.414V12zM15 8a1 1 0 10-2 0v5.586l-1.293-1.293a1 1 0 00-1.414 1.414l3 3a1 1 0 001.414 0l3-3a1 1 0 00-1.414-1.414L15 13.586V8z\"/>" }, "table": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 10h18M3 14h18m-9-4v8m-7 0h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M5 4a3 3 0 00-3 3v6a3 3 0 003 3h10a3 3 0 003-3V7a3 3 0 00-3-3H5zm-1 9v-1h5v2H5a1 1 0 01-1-1zm7 1h4a1 1 0 001-1v-1h-5v2zm0-4h5V8h-5v2zM9 8H4v2h5V8z\" clip-rule=\"evenodd\"/>" }, "tag": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M17.707 9.293a1 1 0 010 1.414l-7 7a1 1 0 01-1.414 0l-7-7A.997.997 0 012 10V5a3 3 0 013-3h5c.256 0 .512.098.707.293l7 7zM5 6a1 1 0 100-2 1 1 0 000 2z\" clip-rule=\"evenodd\"/>" }, "template": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 5a1 1 0 011-1h14a1 1 0 011 1v2a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM4 13a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1v-6zM16 13a1 1 0 011-1h2a1 1 0 011 1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-6z\"/>", "solid": "<path d=\"M3 4a1 1 0 011-1h12a1 1 0 011 1v2a1 1 0 01-1 1H4a1 1 0 01-1-1V4zM3 10a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H4a1 1 0 01-1-1v-6zM14 9a1 1 0 00-1 1v6a1 1 0 001 1h2a1 1 0 001-1v-6a1 1 0 00-1-1h-2z\"/>" }, "terminal": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M2 5a2 2 0 012-2h12a2 2 0 012 2v10a2 2 0 01-2 2H4a2 2 0 01-2-2V5zm3.293 1.293a1 1 0 011.414 0l3 3a1 1 0 010 1.414l-3 3a1 1 0 01-1.414-1.414L7.586 10 5.293 7.707a1 1 0 010-1.414zM11 12a1 1 0 100 2h3a1 1 0 100-2h-3z\" clip-rule=\"evenodd\"/>" }, "thumb-down": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10 14H5.236a2 2 0 01-1.789-2.894l3.5-7A2 2 0 018.736 3h4.018a2 2 0 01.485.06l3.76.94m-7 10v5a2 2 0 002 2h.096c.5 0 .905-.405.905-.904 0-.715.211-1.413.608-2.008L17 13V4m-7 10h2m5-10h2a2 2 0 012 2v6a2 2 0 01-2 2h-2.5\"/>", "solid": "<path d=\"M18 9.5a1.5 1.5 0 11-3 0v-6a1.5 1.5 0 013 0v6zM14 9.667v-5.43a2 2 0 00-1.105-1.79l-.05-.025A4 4 0 0011.055 2H5.64a2 2 0 00-1.962 1.608l-1.2 6A2 2 0 004.44 12H8v4a2 2 0 002 2 1 1 0 001-1v-.667a4 4 0 01.8-2.4l1.4-1.866a4 4 0 00.8-2.4z\"/>" }, "thumb-up": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M14 10h4.764a2 2 0 011.789 2.894l-3.5 7A2 2 0 0115.263 21h-4.017c-.163 0-.326-.02-.485-.06L7 20m7-10V5a2 2 0 00-2-2h-.095c-.5 0-.905.405-.905.905 0 .714-.211 1.412-.608 2.006L7 11v9m7-10h-2M7 20H5a2 2 0 01-2-2v-6a2 2 0 012-2h2.5\"/>", "solid": "<path d=\"M2 10.5a1.5 1.5 0 113 0v6a1.5 1.5 0 01-3 0v-6zM6 10.333v5.43a2 2 0 001.106 1.79l.05.025A4 4 0 008.943 18h5.416a2 2 0 001.962-1.608l1.2-6A2 2 0 0015.56 8H12V4a2 2 0 00-2-2 1 1 0 00-1 1v.667a4 4 0 01-.8 2.4L6.8 7.933a4 4 0 00-.8 2.4z\"/>" }, "ticket": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 5v2m0 4v2m0 4v2M5 5a2 2 0 00-2 2v3a2 2 0 110 4v3a2 2 0 002 2h14a2 2 0 002-2v-3a2 2 0 110-4V7a2 2 0 00-2-2H5z\"/>", "solid": "<path d=\"M2 6a2 2 0 012-2h12a2 2 0 012 2v2a2 2 0 100 4v2a2 2 0 01-2 2H4a2 2 0 01-2-2v-2a2 2 0 100-4V6z\"/>" }, "translate": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 5h12M9 3v2m1.048 9.5A18.022 18.022 0 016.412 9m6.088 9h7M11 21l5-10 5 10M12.751 5C11.783 10.77 8.07 15.61 3 18.129\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M7 2a1 1 0 011 1v1h3a1 1 0 110 2H9.578a18.87 18.87 0 01-1.724 4.78c.29.354.596.696.914 1.026a1 1 0 11-1.44 1.389c-.188-.196-.373-.396-.554-.6a19.098 19.098 0 01-3.107 3.567 1 1 0 01-1.334-1.49 17.087 17.087 0 003.13-3.733 18.992 18.992 0 01-1.487-2.494 1 1 0 111.79-.89c.234.47.489.928.764 1.372.417-.934.752-1.913.997-2.927H3a1 1 0 110-2h3V3a1 1 0 011-1zm6 6a1 1 0 01.894.553l2.991 5.982a.869.869 0 01.02.037l.99 1.98a1 1 0 11-1.79.895L15.383 16h-4.764l-.724 1.447a1 1 0 11-1.788-.894l.99-1.98.019-.038 2.99-5.982A1 1 0 0113 8zm-1.382 6h2.764L13 11.236 11.618 14z\" clip-rule=\"evenodd\"/>" }, "trash": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z\" clip-rule=\"evenodd\"/>" }, "trending-down": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 17h8m0 0V9m0 8l-8-8-4 4-6-6\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M12 13a1 1 0 100 2h5a1 1 0 001-1V9a1 1 0 10-2 0v2.586l-4.293-4.293a1 1 0 00-1.414 0L8 9.586 3.707 5.293a1 1 0 00-1.414 1.414l5 5a1 1 0 001.414 0L11 9.414 14.586 13H12z\" clip-rule=\"evenodd\"/>" }, "trending-up": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 7h8m0 0v8m0-8l-8 8-4-4-6 6\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M12 7a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0V8.414l-4.293 4.293a1 1 0 01-1.414 0L8 10.414l-4.293 4.293a1 1 0 01-1.414-1.414l5-5a1 1 0 011.414 0L11 10.586 14.586 7H12z\" clip-rule=\"evenodd\"/>" }, "truck": { "outline": "<path d=\"M9 17a2 2 0 11-4 0 2 2 0 014 0zM19 17a2 2 0 11-4 0 2 2 0 014 0z\"/><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 16V6a1 1 0 00-1-1H4a1 1 0 00-1 1v10a1 1 0 001 1h1m8-1a1 1 0 01-1 1H9m4-1V8a1 1 0 011-1h2.586a1 1 0 01.707.293l3.414 3.414a1 1 0 01.293.707V16a1 1 0 01-1 1h-1m-6-1a1 1 0 001 1h1M5 17a2 2 0 104 0m-4 0a2 2 0 114 0m6 0a2 2 0 104 0m-4 0a2 2 0 114 0\"/>", "solid": "<path d=\"M8 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zM15 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0z\"/><path d=\"M3 4a1 1 0 00-1 1v10a1 1 0 001 1h1.05a2.5 2.5 0 014.9 0H10a1 1 0 001-1V5a1 1 0 00-1-1H3zM14 7a1 1 0 00-1 1v6.05A2.5 2.5 0 0115.95 16H17a1 1 0 001-1v-5a1 1 0 00-.293-.707l-2-2A1 1 0 0015 7h-1z\"/>" }, "upload": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM6.293 6.707a1 1 0 010-1.414l3-3a1 1 0 011.414 0l3 3a1 1 0 01-1.414 1.414L11 5.414V13a1 1 0 11-2 0V5.414L7.707 6.707a1 1 0 01-1.414 0z\" clip-rule=\"evenodd\"/>" }, "user-add": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z\"/>", "solid": "<path d=\"M8 9a3 3 0 100-6 3 3 0 000 6zM8 11a6 6 0 016 6H2a6 6 0 016-6zM16 7a1 1 0 10-2 0v1h-1a1 1 0 100 2h1v1a1 1 0 102 0v-1h1a1 1 0 100-2h-1V7z\"/>" }, "user-circle": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5.121 17.804A13.937 13.937 0 0112 16c2.5 0 4.847.655 6.879 1.804M15 10a3 3 0 11-6 0 3 3 0 016 0zm6 2a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-6-3a2 2 0 11-4 0 2 2 0 014 0zm-2 4a5 5 0 00-4.546 2.916A5.986 5.986 0 0010 16a5.986 5.986 0 004.546-2.084A5 5 0 0010 11z\" clip-rule=\"evenodd\"/>" }, "user-group": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z\"/>", "solid": "<path d=\"M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3zM6 8a2 2 0 11-4 0 2 2 0 014 0zM16 18v-3a5.972 5.972 0 00-.75-2.906A3.005 3.005 0 0119 15v3h-3zM4.75 12.094A5.973 5.973 0 004 15v3H1v-3a3 3 0 013.75-2.906z\"/>" }, "user-remove": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 7a4 4 0 11-8 0 4 4 0 018 0zM9 14a6 6 0 00-6 6v1h12v-1a6 6 0 00-6-6zM21 12h-6\"/>", "solid": "<path d=\"M11 6a3 3 0 11-6 0 3 3 0 016 0zM14 17a6 6 0 00-12 0h12zM13 8a1 1 0 100 2h4a1 1 0 100-2h-4z\"/>" }, "user": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z\" clip-rule=\"evenodd\"/>" }, "users": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z\"/>", "solid": "<path d=\"M9 6a3 3 0 11-6 0 3 3 0 016 0zM17 6a3 3 0 11-6 0 3 3 0 016 0zM12.93 17c.046-.327.07-.66.07-1a6.97 6.97 0 00-1.5-4.33A5 5 0 0119 16v1h-6.07zM6 11a5 5 0 015 5v1H1v-1a5 5 0 015-5z\"/>" }, "variable": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4.871 4A17.926 17.926 0 003 12c0 2.874.673 5.59 1.871 8m14.13 0a17.926 17.926 0 001.87-8c0-2.874-.673-5.59-1.87-8M9 9h1.246a1 1 0 01.961.725l1.586 5.55a1 1 0 00.961.725H15m1-7h-.08a2 2 0 00-1.519.698L9.6 15.302A2 2 0 018.08 16H8\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M4.649 3.084A1 1 0 015.163 4.4 13.95 13.95 0 004 10c0 1.993.416 3.886 1.164 5.6a1 1 0 01-1.832.8A15.95 15.95 0 012 10c0-2.274.475-4.44 1.332-6.4a1 1 0 011.317-.516zM12.96 7a3 3 0 00-2.342 1.126l-.328.41-.111-.279A2 2 0 008.323 7H8a1 1 0 000 2h.323l.532 1.33-1.035 1.295a1 1 0 01-.781.375H7a1 1 0 100 2h.039a3 3 0 002.342-1.126l.328-.41.111.279A2 2 0 0011.677 14H12a1 1 0 100-2h-.323l-.532-1.33 1.035-1.295A1 1 0 0112.961 9H13a1 1 0 100-2h-.039zm1.874-2.6a1 1 0 011.833-.8A15.95 15.95 0 0118 10c0 2.274-.475 4.44-1.332 6.4a1 1 0 11-1.832-.8A13.949 13.949 0 0016 10c0-1.993-.416-3.886-1.165-5.6z\" clip-rule=\"evenodd\"/>" }, "video-camera": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z\"/>", "solid": "<path d=\"M2 6a2 2 0 012-2h6a2 2 0 012 2v8a2 2 0 01-2 2H4a2 2 0 01-2-2V6zM14.553 7.106A1 1 0 0014 8v4a1 1 0 00.553.894l2 1A1 1 0 0018 13V7a1 1 0 00-1.447-.894l-2 1z\"/>" }, "view-boards": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 17V7m0 10a2 2 0 01-2 2H5a2 2 0 01-2-2V7a2 2 0 012-2h2a2 2 0 012 2m0 10a2 2 0 002 2h2a2 2 0 002-2M9 7a2 2 0 012-2h2a2 2 0 012 2m0 10V7m0 10a2 2 0 002 2h2a2 2 0 002-2V7a2 2 0 00-2-2h-2a2 2 0 00-2 2\"/>", "solid": "<path d=\"M2 4a1 1 0 011-1h2a1 1 0 011 1v12a1 1 0 01-1 1H3a1 1 0 01-1-1V4zM8 4a1 1 0 011-1h2a1 1 0 011 1v12a1 1 0 01-1 1H9a1 1 0 01-1-1V4zM15 3a1 1 0 00-1 1v12a1 1 0 001 1h2a1 1 0 001-1V4a1 1 0 00-1-1h-2z\"/>" }, "view-grid-add": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M17 14v6m-3-3h6M6 10h2a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v2a2 2 0 002 2zm10 0h2a2 2 0 002-2V6a2 2 0 00-2-2h-2a2 2 0 00-2 2v2a2 2 0 002 2zM6 20h2a2 2 0 002-2v-2a2 2 0 00-2-2H6a2 2 0 00-2 2v2a2 2 0 002 2z\"/>", "solid": "<path d=\"M5 3a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V5a2 2 0 00-2-2H5zM5 11a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2v-2a2 2 0 00-2-2H5zM11 5a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V5zM14 11a1 1 0 011 1v1h1a1 1 0 110 2h-1v1a1 1 0 11-2 0v-1h-1a1 1 0 110-2h1v-1a1 1 0 011-1z\"/>" }, "view-grid": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z\"/>", "solid": "<path d=\"M5 3a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V5a2 2 0 00-2-2H5zM5 11a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2v-2a2 2 0 00-2-2H5zM11 5a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V5zM11 13a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z\"/>" }, "view-list": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 10h16M4 14h16M4 18h16\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M3 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/>" }, "volume-off": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z\" clip-rule=\"evenodd\"/><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M17 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M9.383 3.076A1 1 0 0110 4v12a1 1 0 01-1.707.707L4.586 13H2a1 1 0 01-1-1V8a1 1 0 011-1h2.586l3.707-3.707a1 1 0 011.09-.217zM12.293 7.293a1 1 0 011.414 0L15 8.586l1.293-1.293a1 1 0 111.414 1.414L16.414 10l1.293 1.293a1 1 0 01-1.414 1.414L15 11.414l-1.293 1.293a1 1 0 01-1.414-1.414L13.586 10l-1.293-1.293a1 1 0 010-1.414z\" clip-rule=\"evenodd\"/>" }, "volume-up": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15.536 8.464a5 5 0 010 7.072m2.828-9.9a9 9 0 010 12.728M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M9.383 3.076A1 1 0 0110 4v12a1 1 0 01-1.707.707L4.586 13H2a1 1 0 01-1-1V8a1 1 0 011-1h2.586l3.707-3.707a1 1 0 011.09-.217zM14.657 2.929a1 1 0 011.414 0A9.972 9.972 0 0119 10a9.972 9.972 0 01-2.929 7.071 1 1 0 01-1.414-1.414A7.971 7.971 0 0017 10c0-2.21-.894-4.208-2.343-5.657a1 1 0 010-1.414zm-2.829 2.828a1 1 0 011.415 0A5.983 5.983 0 0115 10a5.984 5.984 0 01-1.757 4.243 1 1 0 01-1.415-1.415A3.984 3.984 0 0013 10a3.983 3.983 0 00-1.172-2.828 1 1 0 010-1.415z\" clip-rule=\"evenodd\"/>" }, "wifi": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8.111 16.404a5.5 5.5 0 017.778 0M12 20h.01m-7.08-7.071c3.904-3.905 10.236-3.905 14.141 0M1.394 9.393c5.857-5.857 15.355-5.857 21.213 0\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M17.778 8.222c-4.296-4.296-11.26-4.296-15.556 0A1 1 0 01.808 6.808c5.076-5.077 13.308-5.077 18.384 0a1 1 0 01-1.414 1.414zM14.95 11.05a7 7 0 00-9.9 0 1 1 0 01-1.414-1.414 9 9 0 0112.728 0 1 1 0 01-1.414 1.414zM12.12 13.88a3 3 0 00-4.242 0 1 1 0 01-1.415-1.415 5 5 0 017.072 0 1 1 0 01-1.415 1.415zM9 16a1 1 0 011-1h.01a1 1 0 110 2H10a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/>" }, "x-circle": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z\" clip-rule=\"evenodd\"/>" }, "x": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M6 18L18 6M6 6l12 12\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z\" clip-rule=\"evenodd\"/>" }, "zoom-in": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7\"/>", "solid": "<path d=\"M5 8a1 1 0 011-1h1V6a1 1 0 012 0v1h1a1 1 0 110 2H9v1a1 1 0 11-2 0V9H6a1 1 0 01-1-1z\"/><path fill-rule=\"evenodd\" d=\"M2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8zm6-4a4 4 0 100 8 4 4 0 000-8z\" clip-rule=\"evenodd\"/>" }, "zoom-out": { "outline": "<path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM13 10H7\"/>", "solid": "<path fill-rule=\"evenodd\" d=\"M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z\" clip-rule=\"evenodd\"/><path fill-rule=\"evenodd\" d=\"M5 8a1 1 0 011-1h4a1 1 0 110 2H6a1 1 0 01-1-1z\" clip-rule=\"evenodd\"/>" } };

    /* node_modules/@krowten/svelte-heroicons/Icon.svelte generated by Svelte v3.48.0 */
    const file$6 = "node_modules/@krowten/svelte-heroicons/Icon.svelte";

    function create_fragment$a(ctx) {
    	let svg;
    	let raw_value = iconSet[/*name*/ ctx[0]][/*solid*/ ctx[1] ? 'solid' : 'outline'] + "";

    	let svg_levels = [
    		{ xmlns: "http://www.w3.org/2000/svg" },
    		{ "aria-hidden": "true" },
    		{ viewBox: /*viewBox*/ ctx[2] },
    		{ stroke: /*stroke*/ ctx[3] },
    		{ fill: /*fill*/ ctx[4] },
    		/*$$props*/ ctx[5]
    	];

    	let svg_data = {};

    	for (let i = 0; i < svg_levels.length; i += 1) {
    		svg_data = assign(svg_data, svg_levels[i]);
    	}

    	const block = {
    		c: function create() {
    			svg = svg_element("svg");
    			set_svg_attributes(svg, svg_data);
    			add_location(svg, file$6, 8, 0, 236);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, svg, anchor);
    			svg.innerHTML = raw_value;
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*name, solid*/ 3 && raw_value !== (raw_value = iconSet[/*name*/ ctx[0]][/*solid*/ ctx[1] ? 'solid' : 'outline'] + "")) svg.innerHTML = raw_value;
    			set_svg_attributes(svg, svg_data = get_spread_update(svg_levels, [
    				{ xmlns: "http://www.w3.org/2000/svg" },
    				{ "aria-hidden": "true" },
    				dirty & /*viewBox*/ 4 && { viewBox: /*viewBox*/ ctx[2] },
    				dirty & /*stroke*/ 8 && { stroke: /*stroke*/ ctx[3] },
    				dirty & /*fill*/ 16 && { fill: /*fill*/ ctx[4] },
    				dirty & /*$$props*/ 32 && /*$$props*/ ctx[5]
    			]));
    		},
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(svg);
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
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Icon', slots, []);
    	let { name } = $$props;
    	let { solid = false } = $$props;
    	let { viewBox = '0 0 24 24' } = $$props;
    	let { stroke = !solid ? 'currentColor' : 'none' } = $$props;
    	let { fill = solid ? 'currentColor' : 'none' } = $$props;

    	$$self.$$set = $$new_props => {
    		$$invalidate(5, $$props = assign(assign({}, $$props), exclude_internal_props($$new_props)));
    		if ('name' in $$new_props) $$invalidate(0, name = $$new_props.name);
    		if ('solid' in $$new_props) $$invalidate(1, solid = $$new_props.solid);
    		if ('viewBox' in $$new_props) $$invalidate(2, viewBox = $$new_props.viewBox);
    		if ('stroke' in $$new_props) $$invalidate(3, stroke = $$new_props.stroke);
    		if ('fill' in $$new_props) $$invalidate(4, fill = $$new_props.fill);
    	};

    	$$self.$capture_state = () => ({
    		iconSet,
    		name,
    		solid,
    		viewBox,
    		stroke,
    		fill
    	});

    	$$self.$inject_state = $$new_props => {
    		$$invalidate(5, $$props = assign(assign({}, $$props), $$new_props));
    		if ('name' in $$props) $$invalidate(0, name = $$new_props.name);
    		if ('solid' in $$props) $$invalidate(1, solid = $$new_props.solid);
    		if ('viewBox' in $$props) $$invalidate(2, viewBox = $$new_props.viewBox);
    		if ('stroke' in $$props) $$invalidate(3, stroke = $$new_props.stroke);
    		if ('fill' in $$props) $$invalidate(4, fill = $$new_props.fill);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	$$props = exclude_internal_props($$props);
    	return [name, solid, viewBox, stroke, fill, $$props];
    }

    class Icon extends SvelteComponentDev {
    	constructor(options) {
    		super(options);

    		init(this, options, instance$a, create_fragment$a, safe_not_equal, {
    			name: 0,
    			solid: 1,
    			viewBox: 2,
    			stroke: 3,
    			fill: 4
    		});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Icon",
    			options,
    			id: create_fragment$a.name
    		});

    		const { ctx } = this.$$;
    		const props = options.props || {};

    		if (/*name*/ ctx[0] === undefined && !('name' in props)) {
    			console.warn("<Icon> was created without expected prop 'name'");
    		}
    	}

    	get name() {
    		throw new Error("<Icon>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set name(value) {
    		throw new Error("<Icon>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get solid() {
    		throw new Error("<Icon>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set solid(value) {
    		throw new Error("<Icon>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get viewBox() {
    		throw new Error("<Icon>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set viewBox(value) {
    		throw new Error("<Icon>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get stroke() {
    		throw new Error("<Icon>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set stroke(value) {
    		throw new Error("<Icon>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get fill() {
    		throw new Error("<Icon>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set fill(value) {
    		throw new Error("<Icon>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/xcompo/autotable/autotable.svelte generated by Svelte v3.48.0 */

    const { console: console_1$1 } = globals;
    const file$5 = "entries/xcompo/autotable/autotable.svelte";

    function get_each_context(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[12] = list[i];
    	return child_ctx;
    }

    function get_each_context_1(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[15] = list[i];
    	return child_ctx;
    }

    function get_each_context_2(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[15] = list[i];
    	return child_ctx;
    }

    function get_each_context_3(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[20] = list[i];
    	return child_ctx;
    }

    function get_each_context_4(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[20] = list[i];
    	return child_ctx;
    }

    // (46:8) {#each key_names as key_name}
    function create_each_block_4(ctx) {
    	let th;
    	let t_value = /*key_name*/ ctx[20][1] + "";
    	let t;

    	const block = {
    		c: function create() {
    			th = element("th");
    			t = text(t_value);
    			attr_dev(th, "class", "px-2 py-1");
    			set_style(th, "background-color", "#f8f8f8");
    			add_location(th, file$5, 46, 10, 1188);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, th, anchor);
    			append_dev(th, t);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*key_names*/ 1 && t_value !== (t_value = /*key_name*/ ctx[20][1] + "")) set_data_dev(t, t_value);
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(th);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block_4.name,
    		type: "each",
    		source: "(46:8) {#each key_names as key_name}",
    		ctx
    	});

    	return block;
    }

    // (59:10) {#each key_names as key_name}
    function create_each_block_3(ctx) {
    	let td;
    	let span;
    	let t_value = (/*data*/ ctx[12][/*key_name*/ ctx[20][0]] || "") + "";
    	let t;
    	let span_style_value;

    	const block = {
    		c: function create() {
    			td = element("td");
    			span = element("span");
    			t = text(t_value);
    			attr_dev(span, "class", "p-1 rounded-lg");
    			attr_dev(span, "style", span_style_value = /*color_it*/ ctx[6](/*key_name*/ ctx[20][0], /*data*/ ctx[12][/*key_name*/ ctx[20][0]] || ""));
    			add_location(span, file$5, 60, 14, 1679);
    			attr_dev(td, "class", "px-3 py-1");
    			add_location(td, file$5, 59, 12, 1642);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, td, anchor);
    			append_dev(td, span);
    			append_dev(span, t);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*datas, key_names*/ 3 && t_value !== (t_value = (/*data*/ ctx[12][/*key_name*/ ctx[20][0]] || "") + "")) set_data_dev(t, t_value);

    			if (dirty & /*key_names, datas*/ 3 && span_style_value !== (span_style_value = /*color_it*/ ctx[6](/*key_name*/ ctx[20][0], /*data*/ ctx[12][/*key_name*/ ctx[20][0]] || ""))) {
    				attr_dev(span, "style", span_style_value);
    			}
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(td);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block_3.name,
    		type: "each",
    		source: "(59:10) {#each key_names as key_name}",
    		ctx
    	});

    	return block;
    }

    // (77:18) {#if action["icon"]}
    function create_if_block_2(ctx) {
    	let icon;
    	let current;

    	icon = new Icon({
    			props: {
    				name: /*action*/ ctx[15]["icon"],
    				class: "h-5 w-5"
    			},
    			$$inline: true
    		});

    	const block = {
    		c: function create() {
    			create_component(icon.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(icon, target, anchor);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			const icon_changes = {};
    			if (dirty & /*extern_actions*/ 16) icon_changes.name = /*action*/ ctx[15]["icon"];
    			icon.$set(icon_changes);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(icon.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(icon.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(icon, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_2.name,
    		type: "if",
    		source: "(77:18) {#if action[\\\"icon\\\"]}",
    		ctx
    	});

    	return block;
    }

    // (71:14) {#each extern_actions as action}
    function create_each_block_2(ctx) {
    	let button;
    	let t0;
    	let t1_value = /*action*/ ctx[15].Name + "";
    	let t1;
    	let button_class_value;
    	let current;
    	let mounted;
    	let dispose;
    	let if_block = /*action*/ ctx[15]["icon"] && create_if_block_2(ctx);

    	function click_handler() {
    		return /*click_handler*/ ctx[9](/*action*/ ctx[15], /*data*/ ctx[12]);
    	}

    	const block = {
    		c: function create() {
    			button = element("button");
    			if (if_block) if_block.c();
    			t0 = space();
    			t1 = text(t1_value);
    			attr_dev(button, "class", button_class_value = "flex p-1 m-1 text-sm font-semibold text-white rounded transform hover:scale-110 " + (/*action*/ ctx[15].Class || 'bg-blue-400'));
    			add_location(button, file$5, 71, 16, 2033);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, button, anchor);
    			if (if_block) if_block.m(button, null);
    			append_dev(button, t0);
    			append_dev(button, t1);
    			current = true;

    			if (!mounted) {
    				dispose = listen_dev(button, "click", click_handler, false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(new_ctx, dirty) {
    			ctx = new_ctx;

    			if (/*action*/ ctx[15]["icon"]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);

    					if (dirty & /*extern_actions*/ 16) {
    						transition_in(if_block, 1);
    					}
    				} else {
    					if_block = create_if_block_2(ctx);
    					if_block.c();
    					transition_in(if_block, 1);
    					if_block.m(button, t0);
    				}
    			} else if (if_block) {
    				group_outros();

    				transition_out(if_block, 1, 1, () => {
    					if_block = null;
    				});

    				check_outros();
    			}

    			if ((!current || dirty & /*extern_actions*/ 16) && t1_value !== (t1_value = /*action*/ ctx[15].Name + "")) set_data_dev(t1, t1_value);

    			if (!current || dirty & /*extern_actions*/ 16 && button_class_value !== (button_class_value = "flex p-1 m-1 text-sm font-semibold text-white rounded transform hover:scale-110 " + (/*action*/ ctx[15].Class || 'bg-blue-400'))) {
    				attr_dev(button, "class", button_class_value);
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
    			if (detaching) detach_dev(button);
    			if (if_block) if_block.d();
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block_2.name,
    		type: "each",
    		source: "(71:14) {#each extern_actions as action}",
    		ctx
    	});

    	return block;
    }

    // (85:14) {#if show_drop}
    function create_if_block$2(ctx) {
    	let dropdown;
    	let current;

    	dropdown = new Dropdown({
    			props: {
    				$$slots: { default: [create_default_slot$1] },
    				$$scope: { ctx }
    			},
    			$$inline: true
    		});

    	const block = {
    		c: function create() {
    			create_component(dropdown.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(dropdown, target, anchor);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			const dropdown_changes = {};

    			if (dirty & /*$$scope, drop_actions, datas, action_key*/ 33554470) {
    				dropdown_changes.$$scope = { dirty, ctx };
    			}

    			dropdown.$set(dropdown_changes);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(dropdown.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(dropdown.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(dropdown, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$2.name,
    		type: "if",
    		source: "(85:14) {#if show_drop}",
    		ctx
    	});

    	return block;
    }

    // (94:22) {#if action["icon"]}
    function create_if_block_1$1(ctx) {
    	let icon;
    	let current;

    	icon = new Icon({
    			props: {
    				name: /*action*/ ctx[15]["icon"],
    				class: "h-5 w-5"
    			},
    			$$inline: true
    		});

    	const block = {
    		c: function create() {
    			create_component(icon.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(icon, target, anchor);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			const icon_changes = {};
    			if (dirty & /*drop_actions*/ 32) icon_changes.name = /*action*/ ctx[15]["icon"];
    			icon.$set(icon_changes);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(icon.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(icon.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(icon, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_1$1.name,
    		type: "if",
    		source: "(94:22) {#if action[\\\"icon\\\"]}",
    		ctx
    	});

    	return block;
    }

    // (87:18) {#each drop_actions as action}
    function create_each_block_1(ctx) {
    	let button;
    	let t0;
    	let span;
    	let t1_value = /*action*/ ctx[15].Name + "";
    	let t1;
    	let t2;
    	let current;
    	let mounted;
    	let dispose;
    	let if_block = /*action*/ ctx[15]["icon"] && create_if_block_1$1(ctx);

    	function click_handler_1() {
    		return /*click_handler_1*/ ctx[10](/*action*/ ctx[15], /*data*/ ctx[12]);
    	}

    	const block = {
    		c: function create() {
    			button = element("button");
    			if (if_block) if_block.c();
    			t0 = space();
    			span = element("span");
    			t1 = text(t1_value);
    			t2 = space();
    			add_location(span, file$5, 97, 22, 3091);
    			attr_dev(button, "class", "flex justify-between rounded-sm px-4 py-2 text-sm capitalize text-gray-700 hover:bg-blue-500 hover:text-white");
    			add_location(button, file$5, 87, 20, 2628);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, button, anchor);
    			if (if_block) if_block.m(button, null);
    			append_dev(button, t0);
    			append_dev(button, span);
    			append_dev(span, t1);
    			append_dev(button, t2);
    			current = true;

    			if (!mounted) {
    				dispose = listen_dev(button, "click", click_handler_1, false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(new_ctx, dirty) {
    			ctx = new_ctx;

    			if (/*action*/ ctx[15]["icon"]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);

    					if (dirty & /*drop_actions*/ 32) {
    						transition_in(if_block, 1);
    					}
    				} else {
    					if_block = create_if_block_1$1(ctx);
    					if_block.c();
    					transition_in(if_block, 1);
    					if_block.m(button, t0);
    				}
    			} else if (if_block) {
    				group_outros();

    				transition_out(if_block, 1, 1, () => {
    					if_block = null;
    				});

    				check_outros();
    			}

    			if ((!current || dirty & /*drop_actions*/ 32) && t1_value !== (t1_value = /*action*/ ctx[15].Name + "")) set_data_dev(t1, t1_value);
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
    			if (detaching) detach_dev(button);
    			if (if_block) if_block.d();
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block_1.name,
    		type: "each",
    		source: "(87:18) {#each drop_actions as action}",
    		ctx
    	});

    	return block;
    }

    // (86:16) <Dropdown>
    function create_default_slot$1(ctx) {
    	let each_1_anchor;
    	let current;
    	let each_value_1 = /*drop_actions*/ ctx[5];
    	validate_each_argument(each_value_1);
    	let each_blocks = [];

    	for (let i = 0; i < each_value_1.length; i += 1) {
    		each_blocks[i] = create_each_block_1(get_each_context_1(ctx, each_value_1, i));
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
    			if (dirty & /*drop_actions, datas, action_key*/ 38) {
    				each_value_1 = /*drop_actions*/ ctx[5];
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
    						each_blocks[i].m(each_1_anchor.parentNode, each_1_anchor);
    					}
    				}

    				group_outros();

    				for (i = each_value_1.length; i < each_blocks.length; i += 1) {
    					out(i);
    				}

    				check_outros();
    			}
    		},
    		i: function intro(local) {
    			if (current) return;

    			for (let i = 0; i < each_value_1.length; i += 1) {
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
    		id: create_default_slot$1.name,
    		type: "slot",
    		source: "(86:16) <Dropdown>",
    		ctx
    	});

    	return block;
    }

    // (55:6) {#each datas as data}
    function create_each_block(ctx) {
    	let tr;
    	let t0;
    	let td;
    	let div;
    	let t1;
    	let t2;
    	let current;
    	let each_value_3 = /*key_names*/ ctx[0];
    	validate_each_argument(each_value_3);
    	let each_blocks_1 = [];

    	for (let i = 0; i < each_value_3.length; i += 1) {
    		each_blocks_1[i] = create_each_block_3(get_each_context_3(ctx, each_value_3, i));
    	}

    	let each_value_2 = /*extern_actions*/ ctx[4];
    	validate_each_argument(each_value_2);
    	let each_blocks = [];

    	for (let i = 0; i < each_value_2.length; i += 1) {
    		each_blocks[i] = create_each_block_2(get_each_context_2(ctx, each_value_2, i));
    	}

    	const out = i => transition_out(each_blocks[i], 1, 1, () => {
    		each_blocks[i] = null;
    	});

    	let if_block = /*show_drop*/ ctx[3] && create_if_block$2(ctx);

    	const block = {
    		c: function create() {
    			tr = element("tr");

    			for (let i = 0; i < each_blocks_1.length; i += 1) {
    				each_blocks_1[i].c();
    			}

    			t0 = space();
    			td = element("td");
    			div = element("div");

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].c();
    			}

    			t1 = space();
    			if (if_block) if_block.c();
    			t2 = space();
    			attr_dev(div, "class", "flex flex-row");
    			add_location(div, file$5, 69, 12, 1942);
    			attr_dev(td, "class", "px-3 py-1");
    			add_location(td, file$5, 68, 10, 1907);
    			attr_dev(tr, "class", "hover:bg-gray-100 border-b border-gray-200 py-10 text-gray-700");
    			add_location(tr, file$5, 55, 8, 1495);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, tr, anchor);

    			for (let i = 0; i < each_blocks_1.length; i += 1) {
    				each_blocks_1[i].m(tr, null);
    			}

    			append_dev(tr, t0);
    			append_dev(tr, td);
    			append_dev(td, div);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].m(div, null);
    			}

    			append_dev(div, t1);
    			if (if_block) if_block.m(div, null);
    			append_dev(tr, t2);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*color_it, key_names, datas*/ 67) {
    				each_value_3 = /*key_names*/ ctx[0];
    				validate_each_argument(each_value_3);
    				let i;

    				for (i = 0; i < each_value_3.length; i += 1) {
    					const child_ctx = get_each_context_3(ctx, each_value_3, i);

    					if (each_blocks_1[i]) {
    						each_blocks_1[i].p(child_ctx, dirty);
    					} else {
    						each_blocks_1[i] = create_each_block_3(child_ctx);
    						each_blocks_1[i].c();
    						each_blocks_1[i].m(tr, t0);
    					}
    				}

    				for (; i < each_blocks_1.length; i += 1) {
    					each_blocks_1[i].d(1);
    				}

    				each_blocks_1.length = each_value_3.length;
    			}

    			if (dirty & /*extern_actions, datas, action_key*/ 22) {
    				each_value_2 = /*extern_actions*/ ctx[4];
    				validate_each_argument(each_value_2);
    				let i;

    				for (i = 0; i < each_value_2.length; i += 1) {
    					const child_ctx = get_each_context_2(ctx, each_value_2, i);

    					if (each_blocks[i]) {
    						each_blocks[i].p(child_ctx, dirty);
    						transition_in(each_blocks[i], 1);
    					} else {
    						each_blocks[i] = create_each_block_2(child_ctx);
    						each_blocks[i].c();
    						transition_in(each_blocks[i], 1);
    						each_blocks[i].m(div, t1);
    					}
    				}

    				group_outros();

    				for (i = each_value_2.length; i < each_blocks.length; i += 1) {
    					out(i);
    				}

    				check_outros();
    			}

    			if (/*show_drop*/ ctx[3]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);

    					if (dirty & /*show_drop*/ 8) {
    						transition_in(if_block, 1);
    					}
    				} else {
    					if_block = create_if_block$2(ctx);
    					if_block.c();
    					transition_in(if_block, 1);
    					if_block.m(div, null);
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

    			for (let i = 0; i < each_value_2.length; i += 1) {
    				transition_in(each_blocks[i]);
    			}

    			transition_in(if_block);
    			current = true;
    		},
    		o: function outro(local) {
    			each_blocks = each_blocks.filter(Boolean);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				transition_out(each_blocks[i]);
    			}

    			transition_out(if_block);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(tr);
    			destroy_each(each_blocks_1, detaching);
    			destroy_each(each_blocks, detaching);
    			if (if_block) if_block.d();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block.name,
    		type: "each",
    		source: "(55:6) {#each datas as data}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$9(ctx) {
    	let div;
    	let table;
    	let thead;
    	let tr;
    	let t0;
    	let th;
    	let t2;
    	let tbody;
    	let current;
    	let each_value_4 = /*key_names*/ ctx[0];
    	validate_each_argument(each_value_4);
    	let each_blocks_1 = [];

    	for (let i = 0; i < each_value_4.length; i += 1) {
    		each_blocks_1[i] = create_each_block_4(get_each_context_4(ctx, each_value_4, i));
    	}

    	let each_value = /*datas*/ ctx[1];
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
    			div = element("div");
    			table = element("table");
    			thead = element("thead");
    			tr = element("tr");

    			for (let i = 0; i < each_blocks_1.length; i += 1) {
    				each_blocks_1[i].c();
    			}

    			t0 = space();
    			th = element("th");
    			th.textContent = "Actions";
    			t2 = space();
    			tbody = element("tbody");

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].c();
    			}

    			attr_dev(th, "class", "px-2 py-2");
    			set_style(th, "background-color", "#f8f8f8");
    			add_location(th, file$5, 50, 8, 1310);
    			attr_dev(tr, "class", "rounded-lg text-sm font-medium text-gray-700 text-left");
    			set_style(tr, "font-size", "0.9674rem");
    			add_location(tr, file$5, 41, 6, 1020);
    			add_location(thead, file$5, 40, 4, 1006);
    			attr_dev(tbody, "class", "text-sm font-normal text-gray-700");
    			add_location(tbody, file$5, 53, 4, 1409);
    			attr_dev(table, "class", "table-auto border-collapse w-full bg-white shadow rounded-xl");
    			add_location(table, file$5, 39, 2, 925);
    			attr_dev(div, "class", "p-2 overflow-visible");
    			add_location(div, file$5, 38, 0, 888);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, table);
    			append_dev(table, thead);
    			append_dev(thead, tr);

    			for (let i = 0; i < each_blocks_1.length; i += 1) {
    				each_blocks_1[i].m(tr, null);
    			}

    			append_dev(tr, t0);
    			append_dev(tr, th);
    			append_dev(table, t2);
    			append_dev(table, tbody);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].m(tbody, null);
    			}

    			current = true;
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*key_names*/ 1) {
    				each_value_4 = /*key_names*/ ctx[0];
    				validate_each_argument(each_value_4);
    				let i;

    				for (i = 0; i < each_value_4.length; i += 1) {
    					const child_ctx = get_each_context_4(ctx, each_value_4, i);

    					if (each_blocks_1[i]) {
    						each_blocks_1[i].p(child_ctx, dirty);
    					} else {
    						each_blocks_1[i] = create_each_block_4(child_ctx);
    						each_blocks_1[i].c();
    						each_blocks_1[i].m(tr, t0);
    					}
    				}

    				for (; i < each_blocks_1.length; i += 1) {
    					each_blocks_1[i].d(1);
    				}

    				each_blocks_1.length = each_value_4.length;
    			}

    			if (dirty & /*drop_actions, datas, action_key, show_drop, extern_actions, key_names, color_it*/ 127) {
    				each_value = /*datas*/ ctx[1];
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
    						each_blocks[i].m(tbody, null);
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
    			if (detaching) detach_dev(div);
    			destroy_each(each_blocks_1, detaching);
    			destroy_each(each_blocks, detaching);
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
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Autotable', slots, []);
    	let { actions = [] } = $$props;
    	let { key_names = [] } = $$props;
    	let { datas = [] } = $$props;
    	let { action_key = "" } = $$props;
    	let { color = [] } = $$props;
    	let { show_drop = false } = $$props;
    	let extern_actions = [];
    	let drop_actions = [];

    	if (!show_drop) {
    		extern_actions = actions;
    	} else {
    		extern_actions = actions.filter(v => !v["drop"]);
    		drop_actions = actions.filter(v => !!v["drop"]);
    	}

    	const hashCode = str => {
    		let hash = 77;

    		for (var i = 0; i < str.length; i++) {
    			hash = str.charCodeAt(i) + ((hash << 5) - hash);
    		}

    		return hash;
    	};

    	const color_it = (key, str) => {
    		console.log(key, str);

    		if (!color.includes(key)) {
    			return "";
    		}

    		return `background: hsl(${hashCode(str) % 360}, 100%, 80%)`;
    	};

    	const writable_props = ['actions', 'key_names', 'datas', 'action_key', 'color', 'show_drop'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console_1$1.warn(`<Autotable> was created with unknown prop '${key}'`);
    	});

    	const click_handler = (action, data) => action.Action(data[action_key], data);

    	const click_handler_1 = (action, data) => {
    		action.Action(data[action_key], data);
    	};

    	$$self.$$set = $$props => {
    		if ('actions' in $$props) $$invalidate(7, actions = $$props.actions);
    		if ('key_names' in $$props) $$invalidate(0, key_names = $$props.key_names);
    		if ('datas' in $$props) $$invalidate(1, datas = $$props.datas);
    		if ('action_key' in $$props) $$invalidate(2, action_key = $$props.action_key);
    		if ('color' in $$props) $$invalidate(8, color = $$props.color);
    		if ('show_drop' in $$props) $$invalidate(3, show_drop = $$props.show_drop);
    	};

    	$$self.$capture_state = () => ({
    		Dropdown,
    		Icon,
    		actions,
    		key_names,
    		datas,
    		action_key,
    		color,
    		show_drop,
    		extern_actions,
    		drop_actions,
    		hashCode,
    		color_it
    	});

    	$$self.$inject_state = $$props => {
    		if ('actions' in $$props) $$invalidate(7, actions = $$props.actions);
    		if ('key_names' in $$props) $$invalidate(0, key_names = $$props.key_names);
    		if ('datas' in $$props) $$invalidate(1, datas = $$props.datas);
    		if ('action_key' in $$props) $$invalidate(2, action_key = $$props.action_key);
    		if ('color' in $$props) $$invalidate(8, color = $$props.color);
    		if ('show_drop' in $$props) $$invalidate(3, show_drop = $$props.show_drop);
    		if ('extern_actions' in $$props) $$invalidate(4, extern_actions = $$props.extern_actions);
    		if ('drop_actions' in $$props) $$invalidate(5, drop_actions = $$props.drop_actions);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [
    		key_names,
    		datas,
    		action_key,
    		show_drop,
    		extern_actions,
    		drop_actions,
    		color_it,
    		actions,
    		color,
    		click_handler,
    		click_handler_1
    	];
    }

    class Autotable extends SvelteComponentDev {
    	constructor(options) {
    		super(options);

    		init(this, options, instance$9, create_fragment$9, safe_not_equal, {
    			actions: 7,
    			key_names: 0,
    			datas: 1,
    			action_key: 2,
    			color: 8,
    			show_drop: 3
    		});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Autotable",
    			options,
    			id: create_fragment$9.name
    		});
    	}

    	get actions() {
    		throw new Error("<Autotable>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set actions(value) {
    		throw new Error("<Autotable>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get key_names() {
    		throw new Error("<Autotable>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set key_names(value) {
    		throw new Error("<Autotable>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get datas() {
    		throw new Error("<Autotable>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set datas(value) {
    		throw new Error("<Autotable>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get action_key() {
    		throw new Error("<Autotable>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set action_key(value) {
    		throw new Error("<Autotable>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get color() {
    		throw new Error("<Autotable>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set color(value) {
    		throw new Error("<Autotable>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get show_drop() {
    		throw new Error("<Autotable>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set show_drop(value) {
    		throw new Error("<Autotable>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/xcompo/common/loading_spinner.svelte generated by Svelte v3.48.0 */

    const file$4 = "entries/xcompo/common/loading_spinner.svelte";

    function create_fragment$8(ctx) {
    	let div6;
    	let div5;
    	let div4;
    	let div0;
    	let t0;
    	let div1;
    	let t1;
    	let div2;
    	let t2;
    	let div3;

    	const block = {
    		c: function create() {
    			div6 = element("div");
    			div5 = element("div");
    			div4 = element("div");
    			div0 = element("div");
    			t0 = space();
    			div1 = element("div");
    			t1 = space();
    			div2 = element("div");
    			t2 = space();
    			div3 = element("div");
    			attr_dev(div0, "class", "svelte-i51e3c");
    			add_location(div0, file$4, 3, 6, 108);
    			attr_dev(div1, "class", "svelte-i51e3c");
    			add_location(div1, file$4, 4, 6, 122);
    			attr_dev(div2, "class", "svelte-i51e3c");
    			add_location(div2, file$4, 5, 6, 136);
    			attr_dev(div3, "class", "svelte-i51e3c");
    			add_location(div3, file$4, 6, 6, 150);
    			attr_dev(div4, "class", "lds-ring svelte-i51e3c");
    			add_location(div4, file$4, 2, 4, 79);
    			attr_dev(div5, "class", "p-10 flex justify-center");
    			add_location(div5, file$4, 1, 2, 36);
    			attr_dev(div6, "class", "w-screen h-40 mt-14");
    			add_location(div6, file$4, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div6, anchor);
    			append_dev(div6, div5);
    			append_dev(div5, div4);
    			append_dev(div4, div0);
    			append_dev(div4, t0);
    			append_dev(div4, div1);
    			append_dev(div4, t1);
    			append_dev(div4, div2);
    			append_dev(div4, t2);
    			append_dev(div4, div3);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div6);
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
    	validate_slots('Loading_spinner', slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Loading_spinner> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Loading_spinner extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$8, create_fragment$8, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Loading_spinner",
    			options,
    			id: create_fragment$8.name
    		});
    	}
    }

    /* entries/adapter_editor_easypage/page/_panels/new_page.svelte generated by Svelte v3.48.0 */

    const file$3 = "entries/adapter_editor_easypage/page/_panels/new_page.svelte";

    function create_fragment$7(ctx) {
    	let div2;
    	let p;
    	let t0;
    	let t1;
    	let div0;
    	let label0;
    	let t3;
    	let input0;
    	let t4;
    	let div1;
    	let label1;
    	let t6;
    	let input1;
    	let t7;
    	let button;
    	let mounted;
    	let dispose;

    	const block = {
    		c: function create() {
    			div2 = element("div");
    			p = element("p");
    			t0 = text(/*message*/ ctx[2]);
    			t1 = space();
    			div0 = element("div");
    			label0 = element("label");
    			label0.textContent = "Slug";
    			t3 = space();
    			input0 = element("input");
    			t4 = space();
    			div1 = element("div");
    			label1 = element("label");
    			label1.textContent = "Name";
    			t6 = space();
    			input1 = element("input");
    			t7 = space();
    			button = element("button");
    			button.textContent = "Create";
    			attr_dev(p, "class", "text-red-500");
    			add_location(p, file$3, 21, 2, 404);
    			attr_dev(label0, "class", "block mb-2 text-sm font-bold text-gray-700");
    			attr_dev(label0, "for", "slug");
    			add_location(label0, file$3, 24, 4, 468);
    			attr_dev(input0, "class", "w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline");
    			attr_dev(input0, "id", "slug");
    			attr_dev(input0, "type", "text");
    			attr_dev(input0, "placeholder", "Slug");
    			add_location(input0, file$3, 27, 4, 566);
    			attr_dev(div0, "class", "mb-4");
    			add_location(div0, file$3, 23, 2, 445);
    			attr_dev(label1, "class", "block mb-2 text-sm font-bold text-gray-700");
    			attr_dev(label1, "for", "name");
    			add_location(label1, file$3, 36, 4, 842);
    			attr_dev(input1, "class", "w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline");
    			attr_dev(input1, "id", "name");
    			attr_dev(input1, "type", "text");
    			attr_dev(input1, "placeholder", "name");
    			add_location(input1, file$3, 39, 4, 940);
    			attr_dev(div1, "class", "mb-4");
    			add_location(div1, file$3, 35, 2, 819);
    			attr_dev(button, "class", "w-full px-4 py-2 font-bold text-white bg-blue-500 rounded-full hover:bg-blue-700 focus:outline-none focus:shadow-outline");
    			attr_dev(button, "type", "button");
    			add_location(button, file$3, 48, 2, 1194);
    			attr_dev(div2, "class", "flex flex-col");
    			add_location(div2, file$3, 20, 0, 374);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div2, anchor);
    			append_dev(div2, p);
    			append_dev(p, t0);
    			append_dev(div2, t1);
    			append_dev(div2, div0);
    			append_dev(div0, label0);
    			append_dev(div0, t3);
    			append_dev(div0, input0);
    			set_input_value(input0, /*slug*/ ctx[0]);
    			append_dev(div2, t4);
    			append_dev(div2, div1);
    			append_dev(div1, label1);
    			append_dev(div1, t6);
    			append_dev(div1, input1);
    			set_input_value(input1, /*name*/ ctx[1]);
    			append_dev(div2, t7);
    			append_dev(div2, button);

    			if (!mounted) {
    				dispose = [
    					listen_dev(input0, "input", /*input0_input_handler*/ ctx[5]),
    					listen_dev(input1, "input", /*input1_input_handler*/ ctx[6]),
    					listen_dev(button, "click", /*create*/ ctx[3], false, false, false)
    				];

    				mounted = true;
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*message*/ 4) set_data_dev(t0, /*message*/ ctx[2]);

    			if (dirty & /*slug*/ 1 && input0.value !== /*slug*/ ctx[0]) {
    				set_input_value(input0, /*slug*/ ctx[0]);
    			}

    			if (dirty & /*name*/ 2 && input1.value !== /*name*/ ctx[1]) {
    				set_input_value(input1, /*name*/ ctx[1]);
    			}
    		},
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div2);
    			mounted = false;
    			run_all(dispose);
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
    	validate_slots('New_page', slots, []);
    	let { onSave } = $$props;
    	const validateSlug = v => (/^[a-z](-?[a-z])*$/).test(v);
    	let slug = "";
    	let name = "";
    	let message = "";

    	const create = () => {
    		if (!validateSlug(slug)) {
    			$$invalidate(2, message = "Invalid slug");
    		}

    		if (!name) {
    			$$invalidate(2, message = "Invalid name");
    		}

    		$$invalidate(2, message = "");
    		onSave({ slug, name });
    	};

    	const writable_props = ['onSave'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<New_page> was created with unknown prop '${key}'`);
    	});

    	function input0_input_handler() {
    		slug = this.value;
    		$$invalidate(0, slug);
    	}

    	function input1_input_handler() {
    		name = this.value;
    		$$invalidate(1, name);
    	}

    	$$self.$$set = $$props => {
    		if ('onSave' in $$props) $$invalidate(4, onSave = $$props.onSave);
    	};

    	$$self.$capture_state = () => ({
    		onSave,
    		validateSlug,
    		slug,
    		name,
    		message,
    		create
    	});

    	$$self.$inject_state = $$props => {
    		if ('onSave' in $$props) $$invalidate(4, onSave = $$props.onSave);
    		if ('slug' in $$props) $$invalidate(0, slug = $$props.slug);
    		if ('name' in $$props) $$invalidate(1, name = $$props.name);
    		if ('message' in $$props) $$invalidate(2, message = $$props.message);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [
    		slug,
    		name,
    		message,
    		create,
    		onSave,
    		input0_input_handler,
    		input1_input_handler
    	];
    }

    class New_page extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$7, create_fragment$7, safe_not_equal, { onSave: 4 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "New_page",
    			options,
    			id: create_fragment$7.name
    		});

    		const { ctx } = this.$$;
    		const props = options.props || {};

    		if (/*onSave*/ ctx[4] === undefined && !('onSave' in props)) {
    			console.warn("<New_page> was created without expected prop 'onSave'");
    		}
    	}

    	get onSave() {
    		throw new Error("<New_page>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set onSave(value) {
    		throw new Error("<New_page>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/adapter_editor_easypage/page/_layout.svelte generated by Svelte v3.48.0 */
    const file$2 = "entries/adapter_editor_easypage/page/_layout.svelte";

    function create_fragment$6(ctx) {
    	let div3;
    	let div2;
    	let div0;
    	let h3;
    	let t1;
    	let div1;
    	let button0;
    	let t3;
    	let button1;
    	let t5;
    	let current;
    	let mounted;
    	let dispose;
    	const default_slot_template = /*#slots*/ ctx[4].default;
    	const default_slot = create_slot(default_slot_template, ctx, /*$$scope*/ ctx[3], null);

    	const block = {
    		c: function create() {
    			div3 = element("div");
    			div2 = element("div");
    			div0 = element("div");
    			h3 = element("h3");
    			h3.textContent = "Easypage Adapter Editor";
    			t1 = space();
    			div1 = element("div");
    			button0 = element("button");
    			button0.textContent = "↻";
    			t3 = space();
    			button1 = element("button");
    			button1.textContent = "+";
    			t5 = space();
    			if (default_slot) default_slot.c();
    			attr_dev(h3, "class", "text-white mr-4");
    			add_location(h3, file$2, 13, 6, 433);
    			attr_dev(div0, "class", "flex");
    			add_location(div0, file$2, 12, 4, 408);
    			attr_dev(button0, "class", "bg-blue-200 hover:bg-blue-500 rounded font-bold text-white text-sm px-2");
    			add_location(button0, file$2, 17, 6, 537);
    			attr_dev(button1, "class", "bg-blue-200 hover:bg-blue-500 rounded font-bold text-white text-sm px-2");
    			add_location(button1, file$2, 22, 6, 696);
    			attr_dev(div1, "class", "flex gap-2");
    			add_location(div1, file$2, 16, 4, 506);
    			attr_dev(div2, "class", "flex p-2 justify-between bg-blue-400");
    			add_location(div2, file$2, 11, 2, 353);
    			attr_dev(div3, "class", "bg-blue-50 w-full h-screen font-sans flex flex-col");
    			add_location(div3, file$2, 10, 0, 286);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div3, anchor);
    			append_dev(div3, div2);
    			append_dev(div2, div0);
    			append_dev(div0, h3);
    			append_dev(div2, t1);
    			append_dev(div2, div1);
    			append_dev(div1, button0);
    			append_dev(div1, t3);
    			append_dev(div1, button1);
    			append_dev(div3, t5);

    			if (default_slot) {
    				default_slot.m(div3, null);
    			}

    			current = true;

    			if (!mounted) {
    				dispose = [
    					listen_dev(
    						button0,
    						"click",
    						function () {
    							if (is_function(/*onRefresh*/ ctx[0])) /*onRefresh*/ ctx[0].apply(this, arguments);
    						},
    						false,
    						false,
    						false
    					),
    					listen_dev(button1, "click", /*addPage*/ ctx[1], false, false, false)
    				];

    				mounted = true;
    			}
    		},
    		p: function update(new_ctx, [dirty]) {
    			ctx = new_ctx;

    			if (default_slot) {
    				if (default_slot.p && (!current || dirty & /*$$scope*/ 8)) {
    					update_slot_base(
    						default_slot,
    						default_slot_template,
    						ctx,
    						/*$$scope*/ ctx[3],
    						!current
    						? get_all_dirty_from_scope(/*$$scope*/ ctx[3])
    						: get_slot_changes(default_slot_template, /*$$scope*/ ctx[3], dirty, null),
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
    			if (detaching) detach_dev(div3);
    			if (default_slot) default_slot.d(detaching);
    			mounted = false;
    			run_all(dispose);
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
    	validate_slots('Layout', slots, ['default']);
    	let { onSave } = $$props;
    	let { onRefresh } = $$props;
    	const service = getContext("__easypage_service__");

    	const addPage = () => {
    		service.modal.small_open(New_page, { onSave });
    	};

    	const writable_props = ['onSave', 'onRefresh'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Layout> was created with unknown prop '${key}'`);
    	});

    	$$self.$$set = $$props => {
    		if ('onSave' in $$props) $$invalidate(2, onSave = $$props.onSave);
    		if ('onRefresh' in $$props) $$invalidate(0, onRefresh = $$props.onRefresh);
    		if ('$$scope' in $$props) $$invalidate(3, $$scope = $$props.$$scope);
    	};

    	$$self.$capture_state = () => ({
    		getContext,
    		NewPage: New_page,
    		onSave,
    		onRefresh,
    		service,
    		addPage
    	});

    	$$self.$inject_state = $$props => {
    		if ('onSave' in $$props) $$invalidate(2, onSave = $$props.onSave);
    		if ('onRefresh' in $$props) $$invalidate(0, onRefresh = $$props.onRefresh);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [onRefresh, addPage, onSave, $$scope, slots];
    }

    class Layout extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$6, create_fragment$6, safe_not_equal, { onSave: 2, onRefresh: 0 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Layout",
    			options,
    			id: create_fragment$6.name
    		});

    		const { ctx } = this.$$;
    		const props = options.props || {};

    		if (/*onSave*/ ctx[2] === undefined && !('onSave' in props)) {
    			console.warn("<Layout> was created without expected prop 'onSave'");
    		}

    		if (/*onRefresh*/ ctx[0] === undefined && !('onRefresh' in props)) {
    			console.warn("<Layout> was created without expected prop 'onRefresh'");
    		}
    	}

    	get onSave() {
    		throw new Error("<Layout>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set onSave(value) {
    		throw new Error("<Layout>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get onRefresh() {
    		throw new Error("<Layout>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set onRefresh(value) {
    		throw new Error("<Layout>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/adapter_editor_easypage/page/_panels/link.svelte generated by Svelte v3.48.0 */

    const file$1 = "entries/adapter_editor_easypage/page/_panels/link.svelte";

    // (16:2) {#if _needs_subdomain_fill}
    function create_if_block_1(ctx) {
    	let div;
    	let label;
    	let t1;
    	let input;
    	let mounted;
    	let dispose;

    	const block = {
    		c: function create() {
    			div = element("div");
    			label = element("label");
    			label.textContent = "Sub Domain Fill";
    			t1 = space();
    			input = element("input");
    			attr_dev(label, "class", "block mb-2 text-sm font-bold text-gray-700");
    			attr_dev(label, "for", "subd");
    			add_location(label, file$1, 17, 6, 514);
    			attr_dev(input, "class", "w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline");
    			attr_dev(input, "id", "subd");
    			attr_dev(input, "type", "text");
    			attr_dev(input, "placeholder", "subdomain");
    			add_location(input, file$1, 20, 6, 629);
    			attr_dev(div, "class", "mb-4");
    			add_location(div, file$1, 16, 4, 489);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, label);
    			append_dev(div, t1);
    			append_dev(div, input);
    			set_input_value(input, /*subdomain_fill*/ ctx[0]);

    			if (!mounted) {
    				dispose = listen_dev(input, "input", /*input_input_handler*/ ctx[7]);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*subdomain_fill*/ 1 && input.value !== /*subdomain_fill*/ ctx[0]) {
    				set_input_value(input, /*subdomain_fill*/ ctx[0]);
    			}
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_1.name,
    		type: "if",
    		source: "(16:2) {#if _needs_subdomain_fill}",
    		ctx
    	});

    	return block;
    }

    // (31:2) {#if _show_button}
    function create_if_block$1(ctx) {
    	let button;
    	let mounted;
    	let dispose;

    	const block = {
    		c: function create() {
    			button = element("button");
    			button.textContent = "Go";
    			attr_dev(button, "class", "w-full px-4 py-2 font-bold text-white bg-blue-500 rounded-full hover:bg-blue-700 focus:outline-none focus:shadow-outline");
    			attr_dev(button, "type", "button");
    			add_location(button, file$1, 31, 4, 943);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, button, anchor);

    			if (!mounted) {
    				dispose = listen_dev(button, "click", /*openLink*/ ctx[3], false, false, false);
    				mounted = true;
    			}
    		},
    		p: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(button);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$1.name,
    		type: "if",
    		source: "(31:2) {#if _show_button}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$5(ctx) {
    	let div;
    	let t;
    	let if_block0 = /*_needs_subdomain_fill*/ ctx[1] && create_if_block_1(ctx);
    	let if_block1 = /*_show_button*/ ctx[2] && create_if_block$1(ctx);

    	const block = {
    		c: function create() {
    			div = element("div");
    			if (if_block0) if_block0.c();
    			t = space();
    			if (if_block1) if_block1.c();
    			attr_dev(div, "class", "flex flex-col");
    			add_location(div, file$1, 14, 0, 427);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			if (if_block0) if_block0.m(div, null);
    			append_dev(div, t);
    			if (if_block1) if_block1.m(div, null);
    		},
    		p: function update(ctx, [dirty]) {
    			if (/*_needs_subdomain_fill*/ ctx[1]) {
    				if (if_block0) {
    					if_block0.p(ctx, dirty);
    				} else {
    					if_block0 = create_if_block_1(ctx);
    					if_block0.c();
    					if_block0.m(div, t);
    				}
    			} else if (if_block0) {
    				if_block0.d(1);
    				if_block0 = null;
    			}

    			if (/*_show_button*/ ctx[2]) {
    				if (if_block1) {
    					if_block1.p(ctx, dirty);
    				} else {
    					if_block1 = create_if_block$1(ctx);
    					if_block1.c();
    					if_block1.m(div, null);
    				}
    			} else if (if_block1) {
    				if_block1.d(1);
    				if_block1 = null;
    			}
    		},
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			if (if_block0) if_block0.d();
    			if (if_block1) if_block1.d();
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
    	let _needs_subdomain_fill;
    	let _show_button;
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Link', slots, []);
    	let { domain } = $$props;
    	let { slug } = $$props;
    	let { service } = $$props;
    	let subdomain_fill = "";
    	const openLink = () => window.open(`${domain}/${slug}`.replace("*", subdomain_fill), "_blank");

    	if (!_needs_subdomain_fill) {
    		openLink();
    		service.modal.small_close();
    	}

    	const writable_props = ['domain', 'slug', 'service'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Link> was created with unknown prop '${key}'`);
    	});

    	function input_input_handler() {
    		subdomain_fill = this.value;
    		$$invalidate(0, subdomain_fill);
    	}

    	$$self.$$set = $$props => {
    		if ('domain' in $$props) $$invalidate(4, domain = $$props.domain);
    		if ('slug' in $$props) $$invalidate(5, slug = $$props.slug);
    		if ('service' in $$props) $$invalidate(6, service = $$props.service);
    	};

    	$$self.$capture_state = () => ({
    		domain,
    		slug,
    		service,
    		subdomain_fill,
    		openLink,
    		_needs_subdomain_fill,
    		_show_button
    	});

    	$$self.$inject_state = $$props => {
    		if ('domain' in $$props) $$invalidate(4, domain = $$props.domain);
    		if ('slug' in $$props) $$invalidate(5, slug = $$props.slug);
    		if ('service' in $$props) $$invalidate(6, service = $$props.service);
    		if ('subdomain_fill' in $$props) $$invalidate(0, subdomain_fill = $$props.subdomain_fill);
    		if ('_needs_subdomain_fill' in $$props) $$invalidate(1, _needs_subdomain_fill = $$props._needs_subdomain_fill);
    		if ('_show_button' in $$props) $$invalidate(2, _show_button = $$props._show_button);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	$$self.$$.update = () => {
    		if ($$self.$$.dirty & /*domain*/ 16) {
    			$$invalidate(1, _needs_subdomain_fill = domain.includes("*"));
    		}

    		if ($$self.$$.dirty & /*_needs_subdomain_fill, subdomain_fill*/ 3) {
    			$$invalidate(2, _show_button = !_needs_subdomain_fill || _needs_subdomain_fill && subdomain_fill);
    		}
    	};

    	return [
    		subdomain_fill,
    		_needs_subdomain_fill,
    		_show_button,
    		openLink,
    		domain,
    		slug,
    		service,
    		input_input_handler
    	];
    }

    class Link extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$5, create_fragment$5, safe_not_equal, { domain: 4, slug: 5, service: 6 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Link",
    			options,
    			id: create_fragment$5.name
    		});

    		const { ctx } = this.$$;
    		const props = options.props || {};

    		if (/*domain*/ ctx[4] === undefined && !('domain' in props)) {
    			console.warn("<Link> was created without expected prop 'domain'");
    		}

    		if (/*slug*/ ctx[5] === undefined && !('slug' in props)) {
    			console.warn("<Link> was created without expected prop 'slug'");
    		}

    		if (/*service*/ ctx[6] === undefined && !('service' in props)) {
    			console.warn("<Link> was created without expected prop 'service'");
    		}
    	}

    	get domain() {
    		throw new Error("<Link>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set domain(value) {
    		throw new Error("<Link>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get slug() {
    		throw new Error("<Link>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set slug(value) {
    		throw new Error("<Link>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get service() {
    		throw new Error("<Link>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set service(value) {
    		throw new Error("<Link>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/adapter_editor_easypage/page/start.svelte generated by Svelte v3.48.0 */

    const { console: console_1 } = globals;

    // (25:0) {:else}
    function create_else_block(ctx) {
    	let layout;
    	let current;

    	layout = new Layout({
    			props: {
    				onRefresh: /*func_3*/ ctx[7],
    				onSave: /*func_4*/ ctx[8],
    				$$slots: { default: [create_default_slot] },
    				$$scope: { ctx }
    			},
    			$$inline: true
    		});

    	const block = {
    		c: function create() {
    			create_component(layout.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(layout, target, anchor);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			const layout_changes = {};
    			if (dirty & /*datas*/ 2) layout_changes.onSave = /*func_4*/ ctx[8];

    			if (dirty & /*$$scope, loading, datas*/ 1027) {
    				layout_changes.$$scope = { dirty, ctx };
    			}

    			layout.$set(layout_changes);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(layout.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(layout.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(layout, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_else_block.name,
    		type: "else",
    		source: "(25:0) {:else}",
    		ctx
    	});

    	return block;
    }

    // (23:0) {#if loading}
    function create_if_block(ctx) {
    	let loadingspinner;
    	let current;
    	loadingspinner = new Loading_spinner({ $$inline: true });

    	const block = {
    		c: function create() {
    			create_component(loadingspinner.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(loadingspinner, target, anchor);
    			current = true;
    		},
    		p: noop,
    		i: function intro(local) {
    			if (current) return;
    			transition_in(loadingspinner.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(loadingspinner.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(loadingspinner, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block.name,
    		type: "if",
    		source: "(23:0) {#if loading}",
    		ctx
    	});

    	return block;
    }

    // (26:2) <Layout     onRefresh={() => {       load();     }}     onSave={async (data) => {       const resp = await service.updatePages([...datas, data]);       if (!resp.ok) {         console.log(resp);         return;       }       service.modal.small_close();        load();     }}   >
    function create_default_slot(ctx) {
    	let autotable;
    	let current;

    	autotable = new Autotable({
    			props: {
    				action_key: "slug",
    				actions: [
    					{
    						Name: "Visit",
    						Action: /*func*/ ctx[4],
    						Class: "bg-green-400",
    						icon: "link"
    					},
    					{
    						Name: "Edit",
    						Action: /*func_1*/ ctx[5],
    						icon: "pencil-alt"
    					},
    					{
    						Name: "Delete",
    						Class: "bg-red-400",
    						Action: /*func_2*/ ctx[6],
    						icon: "trash"
    					}
    				],
    				datas: /*datas*/ ctx[1],
    				key_names: [["slug", "Slug"], ["name", "Name"]]
    			},
    			$$inline: true
    		});

    	const block = {
    		c: function create() {
    			create_component(autotable.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(autotable, target, anchor);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			const autotable_changes = {};

    			if (dirty & /*loading, datas*/ 3) autotable_changes.actions = [
    				{
    					Name: "Visit",
    					Action: /*func*/ ctx[4],
    					Class: "bg-green-400",
    					icon: "link"
    				},
    				{
    					Name: "Edit",
    					Action: /*func_1*/ ctx[5],
    					icon: "pencil-alt"
    				},
    				{
    					Name: "Delete",
    					Class: "bg-red-400",
    					Action: /*func_2*/ ctx[6],
    					icon: "trash"
    				}
    			];

    			if (dirty & /*datas*/ 2) autotable_changes.datas = /*datas*/ ctx[1];
    			autotable.$set(autotable_changes);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(autotable.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(autotable.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(autotable, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_default_slot.name,
    		type: "slot",
    		source: "(26:2) <Layout     onRefresh={() => {       load();     }}     onSave={async (data) => {       const resp = await service.updatePages([...datas, data]);       if (!resp.ok) {         console.log(resp);         return;       }       service.modal.small_close();        load();     }}   >",
    		ctx
    	});

    	return block;
    }

    function create_fragment$4(ctx) {
    	let current_block_type_index;
    	let if_block;
    	let if_block_anchor;
    	let current;
    	const if_block_creators = [create_if_block, create_else_block];
    	const if_blocks = [];

    	function select_block_type(ctx, dirty) {
    		if (/*loading*/ ctx[0]) return 0;
    		return 1;
    	}

    	current_block_type_index = select_block_type(ctx);
    	if_block = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx);

    	const block = {
    		c: function create() {
    			if_block.c();
    			if_block_anchor = empty();
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			if_blocks[current_block_type_index].m(target, anchor);
    			insert_dev(target, if_block_anchor, anchor);
    			current = true;
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
    				if_block.m(if_block_anchor.parentNode, if_block_anchor);
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
    			if_blocks[current_block_type_index].d(detaching);
    			if (detaching) detach_dev(if_block_anchor);
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
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Start', slots, []);
    	const service = getContext("__easypage_service__");
    	let loading = true;
    	let message = "";
    	let datas = [];

    	const load = async () => {
    		$$invalidate(0, loading = true);
    		const resp = await service.load();

    		if (!resp.ok) {
    			message = resp.data;
    			console.log("Err", resp.data);
    			return;
    		}

    		$$invalidate(1, datas = resp.data["pages"] || []);
    		$$invalidate(0, loading = false);
    	};

    	load();
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console_1.warn(`<Start> was created with unknown prop '${key}'`);
    	});

    	const func = id => {
    		service.modal.small_open(Link, {
    			domain: "http://test.temphia.local:4000", // fixme => replace this
    			slug: id,
    			service
    		});
    	};

    	const func_1 = id => {
    		location.hash = `/page/${id}`;
    	};

    	const func_2 = async id => {
    		$$invalidate(0, loading = true);
    		const newDatas = datas.filter(v => v["slug"] !== id);
    		await service.updatePages(newDatas);
    		await service.deletePageData(id);
    		load();
    	};

    	const func_3 = () => {
    		load();
    	};

    	const func_4 = async data => {
    		const resp = await service.updatePages([...datas, data]);

    		if (!resp.ok) {
    			console.log(resp);
    			return;
    		}

    		service.modal.small_close();
    		load();
    	};

    	$$self.$capture_state = () => ({
    		getContext,
    		Autotable,
    		LoadingSpinner: Loading_spinner,
    		Layout,
    		Link,
    		service,
    		loading,
    		message,
    		datas,
    		load
    	});

    	$$self.$inject_state = $$props => {
    		if ('loading' in $$props) $$invalidate(0, loading = $$props.loading);
    		if ('message' in $$props) message = $$props.message;
    		if ('datas' in $$props) $$invalidate(1, datas = $$props.datas);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [loading, datas, service, load, func, func_1, func_2, func_3, func_4];
    }

    class Start extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$4, create_fragment$4, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Start",
    			options,
    			id: create_fragment$4.name
    		});
    	}
    }

    var grapes_min = createCommonjsModule(function (module, exports) {
    /*! grapesjs - 0.20.3 */

    });

    var grapejs = /*@__PURE__*/getDefaultExportFromCjs(grapes_min);

    var dist$6 = createCommonjsModule(function (module, exports) {
    /*! grapesjs-preset-webpage - 1.0.2 */
    !function(e,n){module.exports=n();}('undefined'!=typeof globalThis?globalThis:'undefined'!=typeof window?window:commonjsGlobal,(()=>(()=>{var e={d:(n,t)=>{for(var o in t)e.o(t,o)&&!e.o(n,o)&&Object.defineProperty(n,o,{enumerable:!0,get:t[o]});},o:(e,n)=>Object.prototype.hasOwnProperty.call(e,n),r:e=>{'undefined'!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:'Module'}),Object.defineProperty(e,'__esModule',{value:!0});}},n={};e.r(n),e.d(n,{default:()=>u});var t='gjs-open-import-webpage',o='set-device-desktop',r='set-device-tablet',a='set-device-mobile',i='canvas-clear',c=function(){return c=Object.assign||function(e){for(var n,t=1,o=arguments.length;t<o;t++)for(var r in n=arguments[t])Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r]);return e},c.apply(this,arguments)};const l=function(e,n){var l=e.Commands,s=n.textCleanCanvas;l.add(t,function(e,n){var o=e.getConfig('stylePrefix'),r=e.Modal,a=document.createElement('div'),i=n.modalImportLabel,l=n.modalImportContent,s=e.CodeManager.getViewer('CodeMirror').clone(),d=s.editor,u=document.createElement('button');return u.type='button',u.innerHTML=n.modalImportButton,u.className="".concat(o,"btn-prim ").concat(o,"btn-import"),u.onclick=function(n){e.Css.clear(),e.setComponents(d.getValue().trim()),r.close();},s.set(c({codeName:'htmlmixed',theme:'hopscotch',readOnly:0},n.importViewerOptions)),{run:function(e){if(!d){var c=document.createElement('textarea');if(i){var p=document.createElement('div');p.className="".concat(o,"import-label"),p.innerHTML=i,a.appendChild(p);}a.appendChild(c),a.appendChild(u),s.init(c),d=s.editor;}r.setTitle(n.modalImportTitle),r.setContent(a);var m='function'==typeof l?l(e):l;s.setContent(m||''),r.open().onceClose((function(){return e.stopCommand(t)})),d.refresh();},stop:function(){r.close();}}}(e,n)),l.add(o,{run:function(e){return e.setDevice('Desktop')},stop:function(){}}),l.add(r,{run:function(e){return e.setDevice('Tablet')},stop:function(){}}),l.add(a,{run:function(e){return e.setDevice('Mobile portrait')},stop:function(){}}),l.add(i,(function(e){return confirm(s)&&e.runCommand('core:canvas-clear')}));};var s=function(){return s=Object.assign||function(e){for(var n,t=1,o=arguments.length;t<o;t++)for(var r in n=arguments[t])Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r]);return e},s.apply(this,arguments)};var d=function(){return d=Object.assign||function(e){for(var n,t=1,o=arguments.length;t<o;t++)for(var r in n=arguments[t])Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r]);return e},d.apply(this,arguments)};const u=function(e,n){void 0===n&&(n={});var c=d({blocks:['link-block','quote','text-basic'],block:function(){return {}},modalImportTitle:'Import',modalImportButton:'Import',modalImportLabel:'',modalImportContent:'',importViewerOptions:{},textCleanCanvas:'Are you sure you want to clear the canvas?',showStylesOnChange:!0,useCustomTheme:!0},n);if(c.useCustomTheme&&'undefined'!=typeof window){var u='gjs-',p='';[['one','#463a3c'],['two','#b9a5a6'],['three','#804f7b'],['four','#d97aa6']].forEach((function(e){var n=e[0],t=e[1];p+="\n        .".concat(u).concat(n,"-bg {\n          background-color: ").concat(t,";\n        }\n\n        .").concat(u).concat(n,"-color {\n          color: ").concat(t,";\n        }\n\n        .").concat(u).concat(n,"-color-h:hover {\n          color: ").concat(t,";\n        }\n      ");}));var m=document.createElement('style');m.innerText=p,document.head.appendChild(m);}!function(e,n){var t=function(t,o){n.blocks.indexOf(t)>=0&&e.Blocks.add(t,s(s({select:!0,category:'Basic'},o),n.block(t)));};t('link-block',{label:'Link Block',media:"<svg viewBox=\"0 0 24 24\">\n      <path fill=\"currentColor\" d=\"M3.9,12C3.9,10.29 5.29,8.9 7,8.9H11V7H7A5,5 0 0,0 2,12A5,5 0 0,0 7,17H11V15.1H7C5.29,15.1 3.9,13.71 3.9,12M8,13H16V11H8V13M17,7H13V8.9H17C18.71,8.9 20.1,10.29 20.1,12C20.1,13.71 18.71,15.1 17,15.1H13V17H17A5,5 0 0,0 22,12A5,5 0 0,0 17,7Z\"></path>\n    </svg>",content:{type:'link',editable:!1,droppable:!0,style:{display:'inline-block',padding:'5px','min-height':'50px','min-width':'50px'}}}),t('quote',{label:'Quote',media:"<svg viewBox=\"0 0 24 24\">\n        <path fill=\"currentColor\" d=\"M14,17H17L19,13V7H13V13H16M6,17H9L11,13V7H5V13H8L6,17Z\" />\n    </svg>",content:"<blockquote class=\"quote\">\n        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore ipsum dolor sit\n      </blockquote>"}),t('text-basic',{label:'Text section',media:"<svg viewBox=\"0 0 24 24\">\n        <path fill=\"currentColor\" d=\"M21,6V8H3V6H21M3,18H12V16H3V18M3,13H21V11H3V13Z\" />\n    </svg>",content:"<section class=\"bdg-sect\">\n      <h1 class=\"heading\">Insert title here</h1>\n      <p class=\"paragraph\">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua</p>\n      </section>"});}(e,c),l(e,c),function(e,n){var c=e.Panels,l=e.getConfig(),s='sw-visibility',d='export-template',u='open-sm',p='open-tm',m='open-layers',v='open-blocks',H='fullscreen',C='preview',f='style="display: block; max-width:22px"';l.showDevices=!1,c.getPanels().reset([{id:'commands',buttons:[{}]},{id:'devices-c',buttons:[{id:o,command:o,active:!0,label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n            <path fill=\"currentColor\" d=\"M21,16H3V4H21M21,2H3C1.89,2 1,2.89 1,4V16A2,2 0 0,0 3,18H10V20H8V22H16V20H14V18H21A2,2 0 0,0 23,16V4C23,2.89 22.1,2 21,2Z\" />\n        </svg>")},{id:r,command:r,label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n            <path fill=\"currentColor\" d=\"M19,18H5V6H19M21,4H3C1.89,4 1,4.89 1,6V18A2,2 0 0,0 3,20H21A2,2 0 0,0 23,18V6C23,4.89 22.1,4 21,4Z\" />\n        </svg>")},{id:a,command:a,label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n            <path fill=\"currentColor\" d=\"M17,19H7V5H17M17,1H7C5.89,1 5,1.89 5,3V21A2,2 0 0,0 7,23H17A2,2 0 0,0 19,21V3C19,1.89 18.1,1 17,1Z\" />\n        </svg>")}]},{id:'options',buttons:[{id:s,command:s,context:s,label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n        <path fill=\"currentColor\" d=\"M15,5H17V3H15M15,21H17V19H15M11,5H13V3H11M19,5H21V3H19M19,9H21V7H19M19,21H21V19H19M19,13H21V11H19M19,17H21V15H19M3,5H5V3H3M3,9H5V7H3M3,13H5V11H3M3,17H5V15H3M3,21H5V19H3M11,21H13V19H11M7,21H9V19H7M7,5H9V3H7V5Z\" />\n    </svg>")},{id:C,context:C,command:function(){return e.runCommand(C)},label:"<svg ".concat(f," viewBox=\"0 0 24 24\"><path fill=\"currentColor\" d=\"M12,9A3,3 0 0,0 9,12A3,3 0 0,0 12,15A3,3 0 0,0 15,12A3,3 0 0,0 12,9M12,17A5,5 0 0,1 7,12A5,5 0 0,1 12,7A5,5 0 0,1 17,12A5,5 0 0,1 12,17M12,4.5C7,4.5 2.73,7.61 1,12C2.73,16.39 7,19.5 12,19.5C17,19.5 21.27,16.39 23,12C21.27,7.61 17,4.5 12,4.5Z\"></path></svg>")},{id:H,command:H,context:H,label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n            <path fill=\"currentColor\" d=\"M5,5H10V7H7V10H5V5M14,5H19V10H17V7H14V5M17,14H19V19H14V17H17V14M10,17V19H5V14H7V17H10Z\" />\n        </svg>")},{id:d,command:function(){return e.runCommand(d)},label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n            <path fill=\"currentColor\" d=\"M12.89,3L14.85,3.4L11.11,21L9.15,20.6L12.89,3M19.59,12L16,8.41V5.58L22.42,12L16,18.41V15.58L19.59,12M1.58,12L8,5.58V8.41L4.41,12L8,15.58V18.41L1.58,12Z\" />\n        </svg>")},{id:'undo',command:function(){return e.runCommand('core:undo')},label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n            <path fill=\"currentColor\" d=\"M20 13.5C20 17.09 17.09 20 13.5 20H6V18H13.5C16 18 18 16 18 13.5S16 9 13.5 9H7.83L10.91 12.09L9.5 13.5L4 8L9.5 2.5L10.92 3.91L7.83 7H13.5C17.09 7 20 9.91 20 13.5Z\" />\n        </svg>")},{id:'redo',command:function(){return e.runCommand('core:redo')},label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n            <path fill=\"currentColor\" d=\"M10.5 18H18V20H10.5C6.91 20 4 17.09 4 13.5S6.91 7 10.5 7H16.17L13.08 3.91L14.5 2.5L20 8L14.5 13.5L13.09 12.09L16.17 9H10.5C8 9 6 11 6 13.5S8 18 10.5 18Z\" />\n        </svg>")},{id:t,command:function(){return e.runCommand(t)},label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n            <path fill=\"currentColor\" d=\"M5,20H19V18H5M19,9H15V3H9V9H5L12,16L19,9Z\" />\n        </svg>")},{id:i,command:function(){return e.runCommand(i)},label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n              <path fill=\"currentColor\" d=\"M19,4H15.5L14.5,3H9.5L8.5,4H5V6H19M6,19A2,2 0 0,0 8,21H16A2,2 0 0,0 18,19V7H6V19Z\" />\n          </svg>")}]},{id:'views',buttons:[{id:u,command:u,active:!0,label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n            <path fill=\"currentColor\" d=\"M20.71,4.63L19.37,3.29C19,2.9 18.35,2.9 17.96,3.29L9,12.25L11.75,15L20.71,6.04C21.1,5.65 21.1,5 20.71,4.63M7,14A3,3 0 0,0 4,17C4,18.31 2.84,19 2,19C2.92,20.22 4.5,21 6,21A4,4 0 0,0 10,17A3,3 0 0,0 7,14Z\" />\n        </svg>")},{id:p,command:p,label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n          <path fill=\"currentColor\" d=\"M12,15.5A3.5,3.5 0 0,1 8.5,12A3.5,3.5 0 0,1 12,8.5A3.5,3.5 0 0,1 15.5,12A3.5,3.5 0 0,1 12,15.5M19.43,12.97C19.47,12.65 19.5,12.33 19.5,12C19.5,11.67 19.47,11.34 19.43,11L21.54,9.37C21.73,9.22 21.78,8.95 21.66,8.73L19.66,5.27C19.54,5.05 19.27,4.96 19.05,5.05L16.56,6.05C16.04,5.66 15.5,5.32 14.87,5.07L14.5,2.42C14.46,2.18 14.25,2 14,2H10C9.75,2 9.54,2.18 9.5,2.42L9.13,5.07C8.5,5.32 7.96,5.66 7.44,6.05L4.95,5.05C4.73,4.96 4.46,5.05 4.34,5.27L2.34,8.73C2.21,8.95 2.27,9.22 2.46,9.37L4.57,11C4.53,11.34 4.5,11.67 4.5,12C4.5,12.33 4.53,12.65 4.57,12.97L2.46,14.63C2.27,14.78 2.21,15.05 2.34,15.27L4.34,18.73C4.46,18.95 4.73,19.03 4.95,18.95L7.44,17.94C7.96,18.34 8.5,18.68 9.13,18.93L9.5,21.58C9.54,21.82 9.75,22 10,22H14C14.25,22 14.46,21.82 14.5,21.58L14.87,18.93C15.5,18.67 16.04,18.34 16.56,17.94L19.05,18.95C19.27,19.03 19.54,18.95 19.66,18.73L21.66,15.27C21.78,15.05 21.73,14.78 21.54,14.63L19.43,12.97Z\" />\n      </svg>")},{id:m,command:m,label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n          <path fill=\"currentColor\" d=\"M12,16L19.36,10.27L21,9L12,2L3,9L4.63,10.27M12,18.54L4.62,12.81L3,14.07L12,21.07L21,14.07L19.37,12.8L12,18.54Z\" />\n      </svg>")},{id:v,command:v,label:"<svg ".concat(f," viewBox=\"0 0 24 24\">\n          <path fill=\"currentColor\" d=\"M17,13H13V17H11V13H7V11H11V7H13V11H17M19,3H5C3.89,3 3,3.89 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V5C21,3.89 20.1,3 19,3Z\" />\n      </svg>")}]}]);var g=c.getButton('views',v);e.on('load',(function(){return null==g?void 0:g.set('active',!0)})),n.showStylesOnChange&&e.on('component:selected',(function(){var n=c.getButton('views',u),t=c.getButton('views',m);t&&t.get('active')||!e.getSelected()||null==n||n.set('active',!0);}));}(e,c);};return n})()));

    });

    var webpagePlugin = /*@__PURE__*/getDefaultExportFromCjs(dist$6);

    var dist$5 = createCommonjsModule(function (module, exports) {
    /*! grapesjs-blocks-basic - 1.0.1 */
    !function(n,a){module.exports=a();}('undefined'!=typeof globalThis?globalThis:'undefined'!=typeof window?window:commonjsGlobal,(()=>(()=>{var n={d:(a,e)=>{for(var t in e)n.o(e,t)&&!n.o(a,t)&&Object.defineProperty(a,t,{enumerable:!0,get:e[t]});},o:(n,a)=>Object.prototype.hasOwnProperty.call(n,a),r:n=>{'undefined'!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(n,Symbol.toStringTag,{value:'Module'}),Object.defineProperty(n,'__esModule',{value:!0});}},a={};n.r(a),n.d(a,{default:()=>l});var e=function(){return e=Object.assign||function(n){for(var a,e=1,t=arguments.length;e<t;e++)for(var l in a=arguments[e])Object.prototype.hasOwnProperty.call(a,l)&&(n[l]=a[l]);return n},e.apply(this,arguments)};var t=function(){return t=Object.assign||function(n){for(var a,e=1,t=arguments.length;e<t;e++)for(var l in a=arguments[e])Object.prototype.hasOwnProperty.call(a,l)&&(n[l]=a[l]);return n},t.apply(this,arguments)};const l=function(n,a){void 0===a&&(a={}),function(n,a){var t=n.BlockManager,l=a.category,o=a.blocks,c=a.stylePrefix,i=a.flexGrid,d=a.rowHeight,r=a.addBasicStyle,s="".concat(c,"row"),v="".concat(c,"cell"),m=i?"\n    .".concat(s," {\n      display: flex;\n      justify-content: flex-start;\n      align-items: stretch;\n      flex-wrap: nowrap;\n      padding: 10px;\n    }\n    @media (max-width: 768px) {\n      .").concat(s," {\n        flex-wrap: wrap;\n      }\n    }"):"\n    .".concat(s," {\n      display: table;\n      padding: 10px;\n      width: 100%;\n    }\n    @media (max-width: 768px) {\n      .").concat(c,"cell, .").concat(c,"cell30, .").concat(c,"cell70 {\n        width: 100%;\n        display: block;\n      }\n    }"),p=i?"\n    .".concat(v," {\n      min-height: ").concat(d,"px;\n      flex-grow: 1;\n      flex-basis: 100%;\n    }"):"\n    .".concat(v," {\n      width: 8%;\n      display: table-cell;\n      height: ").concat(d,"px;\n    }"),u="\n  .".concat(c,"cell30 {\n    width: 30%;\n  }"),b="\n  .".concat(c,"cell70 {\n    width: 70%;\n  }"),g=1,h={tl:0,tc:0,tr:0,cl:0,cr:0,bl:0,br:0,minDim:g},f=e(e({},h),{cr:1,bc:0,currentUnit:1,minDim:g,step:.2});i&&(f.keyWidth='flex-basis');var y={class:s,'data-gjs-droppable':".".concat(v),'data-gjs-resizable':h,'data-gjs-name':'Row'},C={class:v,'data-gjs-draggable':".".concat(s),'data-gjs-resizable':f,'data-gjs-name':'Cell'};i&&(C['data-gjs-unstylable']=['width'],C['data-gjs-stylable-require']=['flex-basis']);var x=[".".concat(s),".".concat(v)];n.on('selector:add',(function(n){return x.indexOf(n.getFullName())>=0&&n.set('private',1)}));var V=function(n){var a=[];for(var e in n){var t=n[e];t=t instanceof Array||t instanceof Object?JSON.stringify(t):t,a.push("".concat(e,"=").concat("'".concat(t,"'")));}return a.length?" ".concat(a.join(' ')):''},w=function(n){return o.indexOf(n)>=0},L=V(y),H=V(C),M={category:l,select:!0};w('column1')&&t.add('column1',e(e({},M),{label:a.labelColumn1,media:"<svg viewBox=\"0 0 24 24\">\n        <path fill=\"currentColor\" d=\"M2 20h20V4H2v16Zm-1 0V4a1 1 0 0 1 1-1h20a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1Z\"/>\n      </svg>",content:"<div ".concat(L,">\n        <div ").concat(H,"></div>\n      </div>\n      ").concat(r?"<style>\n          ".concat(m,"\n          ").concat(p,"\n        </style>"):'')})),w('column2')&&t.add('column2',e(e({},M),{label:a.labelColumn2,media:"<svg viewBox=\"0 0 23 24\">\n        <path fill=\"currentColor\" d=\"M2 20h8V4H2v16Zm-1 0V4a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1ZM13 20h8V4h-8v16Zm-1 0V4a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1Z\"/>\n      </svg>",content:"<div ".concat(L,">\n        <div ").concat(H,"></div>\n        <div ").concat(H,"></div>\n      </div>\n      ").concat(r?"<style>\n          ".concat(m,"\n          ").concat(p,"\n        </style>"):'')})),w('column3')&&t.add('column3',e(e({},M),{label:a.labelColumn3,media:"<svg viewBox=\"0 0 23 24\">\n        <path fill=\"currentColor\" d=\"M2 20h4V4H2v16Zm-1 0V4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1ZM17 20h4V4h-4v16Zm-1 0V4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1ZM9.5 20h4V4h-4v16Zm-1 0V4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1Z\"/>\n      </svg>",content:"<div ".concat(L,">\n        <div ").concat(H,"></div>\n        <div ").concat(H,"></div>\n        <div ").concat(H,"></div>\n      </div>\n      ").concat(r?"<style>\n          ".concat(m,"\n          ").concat(p,"\n        </style>"):'')})),w('column3-7')&&t.add('column3-7',e(e({},M),{label:a.labelColumn37,media:"<svg viewBox=\"0 0 24 24\">\n        <path fill=\"currentColor\" d=\"M2 20h5V4H2v16Zm-1 0V4a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1ZM10 20h12V4H10v16Zm-1 0V4a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H10a1 1 0 0 1-1-1Z\"/>\n      </svg>",content:"<div ".concat(L,">\n        <div ").concat(H," style='").concat(i?'flex-basis':'width',": 30%;'></div>\n        <div ").concat(H," style='").concat(i?'flex-basis':'width',": 70%;'></div>\n      </div>\n      ").concat(r?"<style>\n          ".concat(m,"\n          ").concat(p,"\n          ").concat(u,"\n          ").concat(b,"\n        </style>"):'')})),w('text')&&t.add('text',e(e({},M),{activate:!0,label:a.labelText,media:"<svg viewBox=\"0 0 24 24\">\n        <path fill=\"currentColor\" d=\"M18.5,4L19.66,8.35L18.7,8.61C18.25,7.74 17.79,6.87 17.26,6.43C16.73,6 16.11,6 15.5,6H13V16.5C13,17 13,17.5 13.33,17.75C13.67,18 14.33,18 15,18V19H9V18C9.67,18 10.33,18 10.67,17.75C11,17.5 11,17 11,16.5V6H8.5C7.89,6 7.27,6 6.74,6.43C6.21,6.87 5.75,7.74 5.3,8.61L4.34,8.35L5.5,4H18.5Z\" />\n      </svg>",content:{type:'text',content:'Insert your text here',style:{padding:'10px'}}})),w('link')&&t.add('link',e(e({},M),{label:a.labelLink,media:"<svg viewBox=\"0 0 24 24\">\n        <path fill=\"currentColor\" d=\"M3.9,12C3.9,10.29 5.29,8.9 7,8.9H11V7H7A5,5 0 0,0 2,12A5,5 0 0,0 7,17H11V15.1H7C5.29,15.1 3.9,13.71 3.9,12M8,13H16V11H8V13M17,7H13V8.9H17C18.71,8.9 20.1,10.29 20.1,12C20.1,13.71 18.71,15.1 17,15.1H13V17H17A5,5 0 0,0 22,12A5,5 0 0,0 17,7Z\" />\n      </svg>",content:{type:'link',content:'Link',style:{color:'#d983a6'}}})),w('image')&&t.add('image',e(e({},M),{activate:!0,label:a.labelImage,media:"<svg viewBox=\"0 0 24 24\">\n        <path fill=\"currentColor\" d=\"M21,3H3C2,3 1,4 1,5V19A2,2 0 0,0 3,21H21C22,21 23,20 23,19V5C23,4 22,3 21,3M5,17L8.5,12.5L11,15.5L14.5,11L19,17H5Z\" />\n      </svg>",content:{style:{color:'black'},type:'image'}})),w('video')&&t.add('video',e(e({},M),{label:a.labelVideo,media:"<svg viewBox=\"0 0 24 24\">\n        <path fill=\"currentColor\" d=\"M10,15L15.19,12L10,9V15M21.56,7.17C21.69,7.64 21.78,8.27 21.84,9.07C21.91,9.87 21.94,10.56 21.94,11.16L22,12C22,14.19 21.84,15.8 21.56,16.83C21.31,17.73 20.73,18.31 19.83,18.56C19.36,18.69 18.5,18.78 17.18,18.84C15.88,18.91 14.69,18.94 13.59,18.94L12,19C7.81,19 5.2,18.84 4.17,18.56C3.27,18.31 2.69,17.73 2.44,16.83C2.31,16.36 2.22,15.73 2.16,14.93C2.09,14.13 2.06,13.44 2.06,12.84L2,12C2,9.81 2.16,8.2 2.44,7.17C2.69,6.27 3.27,5.69 4.17,5.44C4.64,5.31 5.5,5.22 6.82,5.16C8.12,5.09 9.31,5.06 10.41,5.06L12,5C16.19,5 18.8,5.16 19.83,5.44C20.73,5.69 21.31,6.27 21.56,7.17Z\" />\n      </svg>",content:{type:'video',src:'img/video2.webm',style:{height:'350px',width:'615px'}}})),w('map')&&t.add('map',e(e({},M),{label:a.labelMap,media:"<svg viewBox=\"0 0 24 24\">\n        <path fill=\"currentColor\" d=\"M20.5,3L20.34,3.03L15,5.1L9,3L3.36,4.9C3.15,4.97 3,5.15 3,5.38V20.5A0.5,0.5 0 0,0 3.5,21L3.66,20.97L9,18.9L15,21L20.64,19.1C20.85,19.03 21,18.85 21,18.62V3.5A0.5,0.5 0 0,0 20.5,3M10,5.47L14,6.87V18.53L10,17.13V5.47M5,6.46L8,5.45V17.15L5,18.31V6.46M19,17.54L16,18.55V6.86L19,5.7V17.54Z\" />\n      </svg>",content:{type:'map',style:{height:'350px'}}}));}(n,t({blocks:['column1','column2','column3','column3-7','text','link','image','video','map'],flexGrid:!1,stylePrefix:'gjs-',addBasicStyle:!0,category:'Basic',labelColumn1:'1 Column',labelColumn2:'2 Columns',labelColumn3:'3 Columns',labelColumn37:'2 Columns 3/7',labelText:'Text',labelLink:'Link',labelImage:'Image',labelVideo:'Video',labelMap:'Map',rowHeight:75},a));};return a})()));

    });

    var basicPlugin = /*@__PURE__*/getDefaultExportFromCjs(dist$5);

    var dist$4 = createCommonjsModule(function (module, exports) {
    /*! grapesjs-plugin-forms - 2.0.5 */
    !function(e,t){module.exports=t();}('undefined'!=typeof globalThis?globalThis:'undefined'!=typeof window?window:commonjsGlobal,(()=>(()=>{var e={d:(t,n)=>{for(var o in n)e.o(n,o)&&!e.o(t,o)&&Object.defineProperty(t,o,{enumerable:!0,get:n[o]});},o:(e,t)=>Object.prototype.hasOwnProperty.call(e,t),r:e=>{'undefined'!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:'Module'}),Object.defineProperty(e,'__esModule',{value:!0});}},t={};e.r(t),e.d(t,{default:()=>m});var n='form',o='input',a='textarea',i='select',s='checkbox',r='radio',p='button',c='label',l='option';var d=function(){return d=Object.assign||function(e){for(var t,n=1,o=arguments.length;n<o;n++)for(var a in t=arguments[n])Object.prototype.hasOwnProperty.call(t,a)&&(e[a]=t[a]);return e},d.apply(this,arguments)};var u=function(){return u=Object.assign||function(e){for(var t,n=1,o=arguments.length;n<o;n++)for(var a in t=arguments[n])Object.prototype.hasOwnProperty.call(t,a)&&(e[a]=t[a]);return e},u.apply(this,arguments)};const m=function(e,t){void 0===t&&(t={});var m=u({blocks:['form','input','textarea','select','button','label','checkbox','radio'],category:{id:'forms',label:'Forms'},block:function(){return {}}},t);!function(e){var t=e.Components,d={name:'name'},u={name:'placeholder'},m={type:'checkbox',name:'required'},h=function(e,t){return {type:l,content:t,attributes:{value:e}}},v=function(t){e.Commands.isActive('preview')||t.preventDefault();};t.addType(n,{isComponent:function(e){return 'FORM'==e.tagName},model:{defaults:{tagName:'form',droppable:':not(form)',draggable:':not(form)',attributes:{method:'get'},traits:[{type:'select',name:'method',options:[{value:'get',name:'GET'},{value:'post',name:'POST'}]},{name:'action'}]}},view:{events:{submit:function(e){return e.preventDefault()}}}}),t.addType(o,{isComponent:function(e){return 'INPUT'==e.tagName},model:{defaults:{tagName:'input',droppable:!1,highlightable:!1,attributes:{type:'text'},traits:[d,u,{type:'select',name:'type',options:[{value:'text'},{value:'email'},{value:'password'},{value:'number'}]},m]}},extendFnView:['updateAttributes'],view:{updateAttributes:function(){this.el.setAttribute('autocomplete','off');}}}),t.addType(a,{extend:o,isComponent:function(e){return 'TEXTAREA'==e.tagName},model:{defaults:{tagName:'textarea',attributes:{},traits:[d,u,m]}}}),t.addType(l,{isComponent:function(e){return 'OPTION'==e.tagName},model:{defaults:{tagName:'option',layerable:!1,droppable:!1,draggable:!1,highlightable:!1}}}),t.addType(i,{extend:o,isComponent:function(e){return 'SELECT'==e.tagName},model:{defaults:{tagName:'select',components:[h('opt1','Option 1'),h('opt2','Option 2')],traits:[d,{name:'options',type:'select-options'},m]}},view:{events:{mousedown:v}}}),t.addType(s,{extend:o,isComponent:function(e){return 'INPUT'==e.tagName&&'checkbox'==e.type},model:{defaults:{copyable:!1,attributes:{type:'checkbox'},traits:[{name:'id'},d,{name:'value'},m,{type:'checkbox',name:'checked'}]}},view:{events:{click:v},init:function(){this.listenTo(this.model,'change:attributes:checked',this.handleChecked);},handleChecked:function(){var e;this.el.checked=!!(null===(e=this.model.get('attributes'))||void 0===e?void 0:e.checked);}}}),t.addType(r,{extend:s,isComponent:function(e){return 'INPUT'==e.tagName&&'radio'==e.type},model:{defaults:{attributes:{type:'radio'}}}}),t.addType(p,{extend:o,isComponent:function(e){return 'BUTTON'==e.tagName},model:{defaults:{tagName:'button',attributes:{type:'button'},text:'Send',traits:[{name:'text',changeProp:!0},{type:'select',name:'type',options:[{value:'button'},{value:'submit'},{value:'reset'}]}]},init:function(){var e=this.components(),t=1===e.length&&e.models[0],n=t&&t.is('textnode')&&t.get('content')||'',o=n||this.get('text');this.set('text',o),this.on('change:text',this.__onTextChange),o!==n&&this.__onTextChange();},__onTextChange:function(){this.components(this.get('text'));}},view:{events:{click:v}}}),t.addType(c,{extend:'text',isComponent:function(e){return 'LABEL'==e.tagName},model:{defaults:{tagName:'label',components:'Label',traits:[{name:'for'}]}}});}(e),function(e){e.TraitManager.addType('select-options',{events:{keyup:'onChange'},onValueChange:function(){for(var e=this.model,t=this.target,n=e.get('value').trim().split('\n'),o=[],a=0;a<n.length;a++){var i=n[a].split('::');o.push({type:l,components:i[1]||i[0],attributes:{value:i[0]}});}t.components().reset(o),t.view.render();},getInputEl:function(){if(!this.$input){for(var e=[],t=this.target.components(),n=0;n<t.length;n++){var o=t.models[n],a=o.get('attributes').value||'',i=o.components().models[0],s=i&&i.get('content')||'';e.push("".concat(a,"::").concat(s));}this.$input=document.createElement('textarea'),this.$input.value=e.join("\n");}return this.$input}});}(e),function(e,t){var l=t,u=e.BlockManager,m=function(e,n){var o;(null===(o=l.blocks)||void 0===o?void 0:o.indexOf(e))>=0&&u.add(e,d(d(d({},n),{category:l.category,select:!0}),t.block(e)));};m(n,{label:'Form',media:'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M22 5.5c0-.3-.5-.5-1.3-.5H3.4c-.8 0-1.3.2-1.3.5v3c0 .3.5.5 1.3.5h17.4c.8 0 1.3-.2 1.3-.5v-3zM21 8H3V6h18v2zM22 10.5c0-.3-.5-.5-1.3-.5H3.4c-.8 0-1.3.2-1.3.5v3c0 .3.5.5 1.3.5h17.4c.8 0 1.3-.2 1.3-.5v-3zM21 13H3v-2h18v2z"/><rect width="10" height="3" x="2" y="15" rx=".5"/></svg>',content:{type:n,components:[{components:[{type:c,components:'Name'},{type:o}]},{components:[{type:c,components:'Email'},{type:o,attributes:{type:'email'}}]},{components:[{type:c,components:'Gender'},{type:s,attributes:{value:'M'}},{type:c,components:'M'},{type:s,attributes:{value:'F'}},{type:c,components:'F'}]},{components:[{type:c,components:'Message'},{type:a}]},{components:[{type:p}]}]}}),m(o,{label:'Input',media:'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M22 9c0-.6-.5-1-1.3-1H3.4C2.5 8 2 8.4 2 9v6c0 .6.5 1 1.3 1h17.4c.8 0 1.3-.4 1.3-1V9zm-1 6H3V9h18v6z"/><path d="M4 10h1v4H4z"/></svg>',content:{type:o}}),m(a,{label:'Textarea',media:'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M22 7.5c0-.9-.5-1.5-1.3-1.5H3.4C2.5 6 2 6.6 2 7.5v9c0 .9.5 1.5 1.3 1.5h17.4c.8 0 1.3-.6 1.3-1.5v-9zM21 17H3V7h18v10z"/><path d="M4 8h1v4H4zM19 7h1v10h-1zM20 8h1v1h-1zM20 15h1v1h-1z"/></svg>',content:{type:a}}),m(i,{label:'Select',media:'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M22 9c0-.6-.5-1-1.3-1H3.4C2.5 8 2 8.4 2 9v6c0 .6.5 1 1.3 1h17.4c.8 0 1.3-.4 1.3-1V9zm-1 6H3V9h18v6z"/><path d="M18.5 13l1.5-2h-3zM4 11.5h11v1H4z"/></svg>',content:{type:i}}),m(p,{label:'Button',media:'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M22 9c0-.6-.5-1-1.3-1H3.4C2.5 8 2 8.4 2 9v6c0 .6.5 1 1.3 1h17.4c.8 0 1.3-.4 1.3-1V9zm-1 6H3V9h18v6z"/><path d="M4 11.5h16v1H4z"/></svg>',content:{type:p}}),m(c,{label:'Label',media:'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M22 11.9c0-.6-.5-.9-1.3-.9H3.4c-.8 0-1.3.3-1.3.9V17c0 .5.5.9 1.3.9h17.4c.8 0 1.3-.4 1.3-.9V12zM21 17H3v-5h18v5z"/><rect width="14" height="5" x="2" y="5" rx=".5"/><path d="M4 13h1v3H4z"/></svg>',content:{type:c}}),m(s,{label:'Checkbox',media:'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M10 17l-5-5 1.41-1.42L10 14.17l7.59-7.59L19 8m0-5H5c-1.11 0-2 .89-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5a2 2 0 0 0-2-2z"></path></svg>',content:{type:s}}),m(r,{label:'Radio',media:'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M12 20c-4.42 0-8-3.58-8-8s3.58-8 8-8 8 3.58 8 8-3.58 8-8 8m0-18C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2m0 5c-2.76 0-5 2.24-5 5s2.24 5 5 5 5-2.24 5-5-2.24-5-5-5z"></path></svg>',content:{type:r}});}(e,m);};return t})()));

    });

    var gjsForms = /*@__PURE__*/getDefaultExportFromCjs(dist$4);

    var dist$3 = createCommonjsModule(function (module, exports) {
    /*! grapesjs-navbar - 1.0.1 */
    !function(n,t){module.exports=t();}('undefined'!=typeof globalThis?globalThis:'undefined'!=typeof window?window:commonjsGlobal,(()=>(()=>{var n={d:(t,e)=>{for(var a in e)n.o(e,a)&&!n.o(t,a)&&Object.defineProperty(t,a,{enumerable:!0,get:e[a]});},o:(n,t)=>Object.prototype.hasOwnProperty.call(n,t),r:n=>{'undefined'!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(n,Symbol.toStringTag,{value:'Module'}),Object.defineProperty(n,'__esModule',{value:!0});}},t={};n.r(t),n.d(t,{default:()=>i});var e=function(){return e=Object.assign||function(n){for(var t,e=1,a=arguments.length;e<a;e++)for(var o in t=arguments[e])Object.prototype.hasOwnProperty.call(t,o)&&(n[o]=t[o]);return n},e.apply(this,arguments)};const a=function(n,t){var e=n.Components,a=t.id,o=t.label,i=t.classPrefix,r="".concat(a,"-container"),l="".concat(a,"-nav-menu"),c="".concat(a,"-nav-menu-link"),s="".concat(a,"-burger-menu"),d="".concat(a,"-burger-menu-line");e.addType(a,{model:{defaults:{droppable:!1,name:o,attributes:{class:i},components:{type:r},styles:(t.style||"\n          .".concat(i," {\n            background-color: #222;\n            color: #ddd;\n            min-height: 50px;\n            width: 100%;\n          }\n\n          .").concat(i,"-container {\n            max-width: 950px;\n            margin: 0 auto;\n            width: 95%;\n          }\n\n          .").concat(i,"-items-c {\n            display: inline-block;\n            float: right;\n          }\n\n          .").concat(i,"-container::after {\n            content: \"\";\n            clear: both;\n            display: block;\n          }\n\n          .").concat(i,"-brand {\n            vertical-align: top;\n            display: inline-block;\n            padding: 5px;\n            min-height: 50px;\n            min-width: 50px;\n            color: inherit;\n            text-decoration: none;\n          }\n\n          .").concat(i,"-menu {\n            padding: 10px 0;\n            display: block;\n            float: right;\n            margin: 0;\n          }\n\n          .").concat(i,"-menu-link {\n            margin: 0;\n            color: inherit;\n            text-decoration: none;\n            display: inline-block;\n            padding: 10px 15px;\n          }\n\n          .").concat(i,"-burger {\n            margin: 10px 0;\n            width: 45px;\n            padding: 5px 10px;\n            display: none;\n            float: right;\n            cursor: pointer;\n          }\n\n          .").concat(i,"-burger-line {\n            padding: 1px;\n            background-color: white;\n            margin: 5px 0;\n          }\n\n          @media (max-width: 768px) {\n            .").concat(i,"-items-c {\n              display: none;\n              width: 100%;\n            }\n\n            .").concat(i,"-burger {\n              display: block;\n            }\n\n            .").concat(i,"-menu {\n              width: 100%;\n            }\n\n            .").concat(i,"-menu-link {\n              display: block;\n            }\n          }\n        "))+t.styleAdditional}}}),e.addType(r,{model:{defaults:{attributes:{class:"".concat(i,"-container"),'data-gjs':'navbar'},name:'Navbar Container',droppable:!1,draggable:!1,removable:!1,copyable:!1,highlightable:!1,components:[{type:'link',attributes:{class:"".concat(i,"-brand"),href:'/'}},{type:s},{attributes:{class:"".concat(i,"-items-c"),'data-gjs':'navbar-items'},components:{type:l}}]}}}),e.addType(l,{model:{defaults:{name:'Navbar Menu',tagName:'nav',attributes:{class:"".concat(i,"-menu")},components:[{type:c,components:'Home'},{type:c,components:'About'},{type:c,components:'Contact'}]}}}),e.addType(c,{extend:'link',model:{defaults:{name:'Menu link',draggable:"[data-gjs-type=\"".concat(l,"\"]"),attributes:{class:"".concat(i,"-menu-link")}}}}),e.addType(s,{model:{defaults:{name:'Burger',draggable:!1,droppable:!1,copyable:!1,removable:!1,script:function(){var n,t=this,e='gjs-collapse',a='max-height',o=0,i=function(){var n=document.createElement('void'),t={transition:'transitionend',OTransition:'oTransitionEnd',MozTransition:'transitionend',WebkitTransition:'webkitTransitionEnd'};for(var e in t)if(void 0!==n.style[e])return t[e]}(),r=function(n){o=1;var t=function(n){var t=window.getComputedStyle(n),e=t.display,o=parseInt(t[a]);if('none'!==e&&0!==o)return n.offsetHeight;n.style.height='auto',n.style.display='block',n.style.position='absolute',n.style.visibility='hidden';var i=n.offsetHeight;return n.style.height='',n.style.display='',n.style.position='',n.style.visibility='',i}(n),e=n.style;e.display='block',e.transition="".concat(a," 0.25s ease-in-out"),e.overflowY='hidden',''==e[a]&&(e[a]=0),0==parseInt(e[a])?(e[a]='0',setTimeout((function(){e[a]=t+'px';}),10)):e[a]='0';};e in t||t.addEventListener('click',(function(e){if(e.preventDefault(),!o){var l=t.closest("[data-gjs=navbar]"),c=null==l?void 0:l.querySelector("[data-gjs=navbar-items]");c&&r(c),n||(null==c||c.addEventListener(i,(function(){o=0;var n=c.style;0==parseInt(n[a])&&(n.display='',n[a]='');})),n=1);}})),t[e]=1;},attributes:{class:"".concat(i,"-burger")},components:[{type:d},{type:d},{type:d}]}}}),e.addType(d,{model:{defaults:{name:'Burger Line',droppable:!1,draggable:!1,highlightable:!1,attributes:{class:"".concat(i,"-burger-line")}}}});};var o=function(){return o=Object.assign||function(n){for(var t,e=1,a=arguments.length;e<a;e++)for(var o in t=arguments[e])Object.prototype.hasOwnProperty.call(t,o)&&(n[o]=t[o]);return n},o.apply(this,arguments)};const i=function(n,t){void 0===t&&(t={});var i=o({id:'navbar',label:'Navbar',block:{},style:'',styleAdditional:'',classPrefix:'navbar'},t);!function(n,t){var a=t.block,o=t.label,i=t.id;a&&n.Blocks.add(i,e({media:"<svg viewBox=\"0 0 24 24\">\n        <path d=\"M22 9c0-.6-.5-1-1.25-1H3.25C2.5 8 2 8.4 2 9v6c0 .6.5 1 1.25 1h17.5c.75 0 1.25-.4 1.25-1V9Zm-1 6H3V9h18v6Z\"/><path d=\"M15 10h5v1h-5zM15 13h5v1h-5zM15 11.5h5v1h-5z\"/>\n      </svg>",label:o,category:'Extra',select:!0,content:{type:i}},a));}(n,i),a(n,i);};return t})()));

    });

    var navPlugin = /*@__PURE__*/getDefaultExportFromCjs(dist$3);

    var dist$2 = createCommonjsModule(function (module, exports) {
    /*! grapesjs-custom-code - 1.0.1 */
    !function(e,t){module.exports=t();}('undefined'!=typeof globalThis?globalThis:'undefined'!=typeof window?window:commonjsGlobal,(()=>(()=>{var e={d:(t,o)=>{for(var n in o)e.o(o,n)&&!e.o(t,n)&&Object.defineProperty(t,n,{enumerable:!0,get:o[n]});},o:(e,t)=>Object.prototype.hasOwnProperty.call(e,t),r:e=>{'undefined'!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:'Module'}),Object.defineProperty(e,'__esModule',{value:!0});}},t={};e.r(t),e.d(t,{default:()=>d});var o='custom-code-plugin__code',n='custom-code',i='custom-code:open-modal',r=function(){return r=Object.assign||function(e){for(var t,o=1,n=arguments.length;o<n;o++)for(var i in t=arguments[o])Object.prototype.hasOwnProperty.call(t,i)&&(e[i]=t[i]);return e},r.apply(this,arguments)};var s=function(){return s=Object.assign||function(e){for(var t,o=1,n=arguments.length;o<n;o++)for(var i in t=arguments[o])Object.prototype.hasOwnProperty.call(t,i)&&(e[i]=t[i]);return e},s.apply(this,arguments)};var a=function(){return a=Object.assign||function(e){for(var t,o=1,n=arguments.length;o<n;o++)for(var i in t=arguments[o])Object.prototype.hasOwnProperty.call(t,i)&&(e[i]=t[i]);return e},a.apply(this,arguments)};var c=function(){return c=Object.assign||function(e){for(var t,o=1,n=arguments.length;o<n;o++)for(var i in t=arguments[o])Object.prototype.hasOwnProperty.call(t,i)&&(e[i]=t[i]);return e},c.apply(this,arguments)};const d=function(e,t){void 0===t&&(t={});var d=c({blockCustomCode:{},propsCustomCode:{},toolbarBtnCustomCode:{},placeholderScript:"<div style=\"pointer-events: none; padding: 10px;\">\n      <svg viewBox=\"0 0 24 24\" style=\"height: 30px; vertical-align: middle;\">\n        <path d=\"M13 14h-2v-4h2m0 8h-2v-2h2M1 21h22L12 2 1 21z\"></path>\n        </svg>\n      Custom code with <i>&lt;script&gt;</i> can't be rendered on the canvas\n    </div>",modalTitle:'Insert your code',codeViewOptions:{},buttonLabel:'Save',commandCustomCode:{}},t);!function(e,t){void 0===t&&(t={});var s,a=e.Components,c=t.toolbarBtnCustomCode;a.addType('script',{view:{onRender:function(){var e=this.model,t=this.el;e.closestType(n)&&(t.innerHTML='');}}}),a.addType(n,{model:{defaults:r({name:'Custom Code',editable:!0,components:{tagName:'span',components:{type:'textnode',content:'Insert here your custom code'}}},t.propsCustomCode),init:function(){this.on("change:".concat(o),this.onCustomCodeChange);var e=this.get(o);!this.components().length&&this.components(e);var t=this.get('toolbar'),n='custom-code';c&&!t.filter((function(e){return e.id===n})).length&&t.unshift(r({id:n,command:i,label:"<svg viewBox=\"0 0 24 24\">\n              <path d=\"M14.6 16.6l4.6-4.6-4.6-4.6L16 6l6 6-6 6-1.4-1.4m-5.2 0L4.8 12l4.6-4.6L8 6l-6 6 6 6 1.4-1.4z\"></path>\n            </svg>"},c));},onCustomCodeChange:function(){this.components(this.get(o));}},view:{events:{dblclick:'onActive'},init:function(){this.listenTo(this.model.components(),'add remove reset',this.onComponentsChange),this.onComponentsChange();},onComponentsChange:function(){var e=this;s&&clearInterval(s),s=setTimeout((function(){var n=e,i=n.model,r=n.el,s=!0;(i.get(o)||'').indexOf('<script')>=0&&t.placeholderScript&&(r.innerHTML=t.placeholderScript,s=!1),i.set({droppable:s});}),0);},onActive:function(){var e=this.model;this.em.get('Commands').run(i,{target:e});}}});}(e,d),function(e,t){var o=(void 0===t?{}:t).blockCustomCode,i=e.Blocks;o&&i.add(n,s({label:'Custom Code',media:"\n      <svg viewBox=\"0 0 24 24\">\n        <path d=\"M14.6 16.6l4.6-4.6-4.6-4.6L16 6l6 6-6 6-1.4-1.4m-5.2 0L4.8 12l4.6-4.6L8 6l-6 6 6 6 1.4-1.4z\"></path>\n      </svg>\n    ",category:'Extra',activate:!0,select:!0,content:{type:n}},o));}(e,d),function(e,t){void 0===t&&(t={});var n=t.modalTitle,r=t.codeViewOptions,s=t.commandCustomCode,c=function(e,t){t instanceof HTMLElement?e.appendChild(t):t&&e.insertAdjacentHTML('beforeend',t);};e.Commands.add(i,a({keyCustomCode:o,run:function(e,t,o){void 0===o&&(o={});var n=o.target||e.getSelected();this.target=n,(null==n?void 0:n.get('editable'))&&this.showCustomCode(n,o);},stop:function(e){e.Modal.close();},showCustomCode:function(t,r){var s=r.title||n,a=t.get(o)||'',c=this.getContent();e.Modal.open({title:s,content:c}).onceClose((function(){return e.stopCommand(i)})),this.getCodeViewer().setContent(a);},getPreContent:function(){},getPostContent:function(){},getContent:function(){var t=this.getCodeViewer(),o=document.createElement('div'),n=e.getConfig('stylePrefix');return o.className="".concat(n,"custom-code"),c(o,this.getPreContent()),o.appendChild(t.getElement()),c(o,this.getPostContent()),c(o,this.getContentActions()),t.refresh(),setTimeout((function(){return t.focus()}),0),o},getContentActions:function(){var o=this,n=document.createElement('button');n.setAttribute('type','button');var i=e.getConfig('stylePrefix');return n.innerHTML=t.buttonLabel,n.className="".concat(i,"btn-prim ").concat(i,"btn-import__custom-code"),n.onclick=function(){return o.handleSave()},n},handleSave:function(){var t=this.target,n=this.getCodeViewer().getContent();t.set(o,n),e.Modal.close();},getCodeViewer:function(){return this.codeViewer||(this.codeViewer=e.CodeManager.createViewer(a({codeName:'htmlmixed',theme:'hopscotch',readOnly:0},r))),this.codeViewer}},s));}(e,d);};return t})()));

    });

    var customCodePlugin = /*@__PURE__*/getDefaultExportFromCjs(dist$2);

    var require$$0 = grapes_min;

    var grapesjsBlocksFlexbox_min = createCommonjsModule(function (module, exports) {
    /*! grapesjs-blocks-flexbox - 0.1.1 */
    !function(e,t){module.exports=t(require$$0);}(commonjsGlobal,function(e){return function(e){function t(r){if(n[r])return n[r].exports;var a=n[r]={i:r,l:!1,exports:{}};return e[r].call(a.exports,a,a.exports,t),a.l=!0,a.exports}var n={};return t.m=e,t.c=n,t.d=function(e,n,r){t.o(e,n)||Object.defineProperty(e,n,{configurable:!1,enumerable:!0,get:r});},t.n=function(e){var n=e&&e.__esModule?function(){return e.default}:function(){return e};return t.d(n,"a",n),n},t.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},t.p="",t(t.s=0)}([function(e,t,n){function r(e){return e&&e.__esModule?e:{default:e}}Object.defineProperty(t,"__esModule",{value:!0});var a=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r]);}return e},o=n(1),l=r(o),s=n(2),i=r(s);t.default=l.default.plugins.add("gjs-blocks-flexbox",function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{},n={flexboxBlock:{},stylePrefix:"",labelRow:"Row",labelColumn:"Column"},r=a({},t,n);(0, i.default)(e,r);});},function(t,n){t.exports=e;},function(e,t,n){Object.defineProperty(t,"__esModule",{value:!0});var r=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r]);}return e};t.default=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{},n=e.BlockManager,a=(t.blocks,t.stylePrefix),o=a+"row",l=a+"cell",s=t.labelRow,i=t.labelColumn,c=function(e){var t=[];for(var n in e){var r=e[n],a=r instanceof Array||r instanceof Object;r=a?JSON.stringify(r):r,t.push(n+"="+(a?"'"+r+"'":'"'+r+'"'));}return t.length?" "+t.join(" "):""},f={tl:0,tc:0,tr:0,cl:0,cr:0,bl:0,br:0,minDim:1},u=r({},f,{cr:1,bc:0,keyWidth:"flex-basis",currentUnit:1,minDim:1,step:.2}),d={class:o,"data-gjs-droppable":"."+l,"data-gjs-resizable":f,"data-gjs-custom-name":s},p={class:l,"data-gjs-draggable":"."+o,"data-gjs-resizable":u,"data-gjs-custom-name":i,"data-gjs-unstylable":["width"],"data-gjs-stylable-require":["flex-basis"]},b=["."+o,"."+l];e.on("selector:add",function(e){return b.indexOf(e.getFullName())>=0&&e.set("private",1)});var x=c(d),g=c(p),j="\n    ."+o+" {\n      display: flex;\n      justify-content: flex-start;\n      align-items: stretch;\n      flex-wrap: nowrap;\n      padding: 10px;\n    }\n    @media (max-width: 768px) {\n      ."+o+" {\n        flex-wrap: wrap;\n      }\n    }\n    ",v="\n    ."+l+" {\n      min-height: 75px;\n      flex-grow: 1;\n      flex-basis: 100%;\n    }";n.add("flexbox",r({label:"Flexbox",category:"Basic",attributes:{class:"gjs-fonts gjs-f-b2"},content:"\n        <div "+x+">\n          <div "+g+"></div>\n          <div "+g+"></div>\n        </div>\n        <style>\n          "+j+"\n          "+v+"\n        </style>\n        "},t.flexboxBlock));};}])});
    });

    var blkFlexboxPlugin = /*@__PURE__*/getDefaultExportFromCjs(grapesjsBlocksFlexbox_min);

    var grapesjsStyleGradient_min = createCommonjsModule(function (module, exports) {
    /*! grapesjs-style-gradient - 2.0.14 */
    !function(e,t){module.exports=t();}(window,(function(){return function(e){var t={};function n(r){if(t[r])return t[r].exports;var o=t[r]={i:r,l:!1,exports:{}};return e[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=e,n.c=t,n.d=function(e,t,r){n.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:r});},n.r=function(e){'undefined'!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:'Module'}),Object.defineProperty(e,'__esModule',{value:!0});},n.t=function(e,t){if(1&t&&(e=n(e)),8&t)return e;if(4&t&&'object'==typeof e&&e&&e.__esModule)return e;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,'default',{enumerable:!0,value:e}),2&t&&'string'!=typeof e)for(var o in e)n.d(r,o,function(t){return e[t]}.bind(null,o));return r},n.n=function(e){var t=e&&e.__esModule?function(){return e['default']}:function(){return e};return n.d(t,'a',t),t},n.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},n.p="",n(n.s=3)}([function(e,t){e.exports=function(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e};},function(e,t){function n(t){return "function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?e.exports=n=function(e){return typeof e}:e.exports=n=function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e},n(t)}e.exports=n;},function(e,t,n){(e.exports=function(e){function t(r){if(n[r])return n[r].exports;var o=n[r]={i:r,l:!1,exports:{}};return e[r].call(o.exports,o,o.exports,t),o.l=!0,o.exports}var n={};return t.m=e,t.c=n,t.d=function(e,n,r){t.o(e,n)||Object.defineProperty(e,n,{configurable:!1,enumerable:!0,get:r});},t.n=function(e){var n=e&&e.__esModule?function(){return e.default}:function(){return e};return t.d(n,"a",n),n},t.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},t.p="",t(t.s=1)}([function(e,t,n){Object.defineProperty(t,"__esModule",{value:!0}),t.on=function(e,t,n){t=t.split(/\s+/);for(var r=0;r<t.length;++r)e.addEventListener(t[r],n);},t.off=function(e,t,n){t=t.split(/\s+/);for(var r=0;r<t.length;++r)e.removeEventListener(t[r],n);},t.isFunction=function(e){return "function"==typeof e},t.isDef=function(e){return void 0!==e},t.getPointerEvent=function(e){return e.touches&&e.touches[0]||e};},function(e,t,n){var r=function(e){return e&&e.__esModule?e:{default:e}}(n(2));e.exports=function(e){return new r.default(e)};},function(e,t,n){function r(e){return e&&e.__esModule?e:{default:e}}function o(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}function i(e,t){if(!e)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return !t||"object"!=typeof t&&"function"!=typeof t?e:t}Object.defineProperty(t,"__esModule",{value:!0});var a=function(){function e(e,t){for(var n=0;n<t.length;n++){var r=t[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),Object.defineProperty(e,r.key,r);}}return function(t,n,r){return n&&e(t.prototype,n),r&&e(t,r),t}}(),l=r(n(3)),c=r(n(4)),u=n(0),s=function(e,t){return e.position-t.position},f=function(e){return e+"-gradient("},d=function(e){function t(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};o(this,t);var n=i(this,(t.__proto__||Object.getPrototypeOf(t)).call(this));e=Object.assign({},e);var r={pfx:"grp",el:".grp",colorEl:"",min:0,max:100,direction:"90deg",type:"linear",height:"30px",width:"100%",emptyColor:"#000",onValuePos:function(e){return parseInt(e)}};for(var a in r)a in e||(e[a]=r[a]);var l=e.el;if(!((l="string"==typeof l?document.querySelector(l):l)instanceof HTMLElement))throw "Element not found, given "+l;return n.el=l,n.handlers=[],n.options=e,n.on("handler:color:change",(function(e,t){return n.change(t)})),n.on("handler:position:change",(function(e,t){return n.change(t)})),n.on("handler:remove",(function(e){return n.change(1)})),n.on("handler:add",(function(e){return n.change(1)})),n.render(),n}return function(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function, not "+typeof t);e.prototype=Object.create(t&&t.prototype,{constructor:{value:e,enumerable:!1,writable:!0,configurable:!0}}),t&&(Object.setPrototypeOf?Object.setPrototypeOf(e,t):e.__proto__=t);}(t,e),a(t,[{key:"destroy",value:function(){var e=this;this.clear(),this.e={},["el","handlers","options","colorPicker"].forEach((function(t){return e[t]=0})),["previewEl","wrapperEl","sandEl"].forEach((function(t){var n=e[t];n&&n.parentNode&&n.parentNode.removeChild(n),delete e[t];}));}},{key:"setColorPicker",value:function(e){this.colorPicker=e;}},{key:"getValue",value:function(e,t){var n=this.getColorValue(),r=e||this.getType(),o=["top","left","bottom","right","center"],i=t||this.getDirection();return ["linear","repeating-linear"].indexOf(r)>=0&&o.indexOf(i)>=0&&(i="center"===i?"to right":"to "+i),["radial","repeating-radial"].indexOf(r)>=0&&o.indexOf(i)>=0&&(i="circle at "+i),n?r+"-gradient("+i+", "+n+")":""}},{key:"getSafeValue",value:function(e,t){var n=this.previewEl,r=this.getValue(e,t);if(!this.sandEl&&(this.sandEl=document.createElement("div")),!n||!r)return "";for(var o=this.sandEl.style,i=[r].concat(function(e){if(Array.isArray(e)){for(var t=0,n=Array(e.length);t<e.length;t++)n[t]=e[t];return n}return Array.from(e)}(this.getPrefixedValues(e,t))),a=void 0,l=0;l<i.length&&(a=i[l],o.backgroundImage=a,o.backgroundImage!=a);l++);return o.backgroundImage}},{key:"setValue",value:function(){var e=this,t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"",n=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{},r=this.type,o=this.direction,i=t.indexOf("(")+1,a=t.lastIndexOf(")"),l=t.substring(i,a),c=l.split(/,(?![^(]*\)) /);if(this.clear(n),l){c.length>2&&(o=c.shift());var u=void 0;["repeating-linear","repeating-radial","linear","radial"].forEach((function(e){t.indexOf(f(e))>-1&&!u&&(u=1,r=e);})),this.setDirection(o,n),this.setType(r,n),c.forEach((function(t){var r=t.split(" "),o=parseFloat(r.pop()),i=r.join("");e.addHandler(o,i,0,n);})),this.updatePreview();}else this.updatePreview();}},{key:"getColorValue",value:function(){var e=this.handlers;return e.sort(s),(e=1==e.length?[e[0],e[0]]:e).map((function(e){return e.getValue()})).join(", ")}},{key:"getPrefixedValues",value:function(e,t){var n=this.getValue(e,t);return ["-moz-","-webkit-","-o-","-ms-"].map((function(e){return ""+e+n}))}},{key:"change",value:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:1,t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};this.updatePreview(),!t.silent&&this.emit("change",e);}},{key:"setDirection",value:function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};this.options.direction=e;var n=t.complete,r=void 0===n?1:n;this.change(r,t);}},{key:"getDirection",value:function(){return this.options.direction}},{key:"setType",value:function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};this.options.type=e;var n=t.complete,r=void 0===n?1:n;this.change(r,t);}},{key:"getType",value:function(){return this.options.type}},{key:"addHandler",value:function(e,t){var n=arguments.length>2&&void 0!==arguments[2]?arguments[2]:1,r=arguments.length>3&&void 0!==arguments[3]?arguments[3]:{},o=new c.default(this,e,t,n,r);return !r.silent&&this.emit("handler:add",o),o}},{key:"getHandler",value:function(e){return this.handlers[e]}},{key:"getHandlers",value:function(){return this.handlers}},{key:"clear",value:function(){for(var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},t=this.handlers,n=t.length-1;n>=0;n--)t[n].remove(e);}},{key:"getSelected",value:function(){for(var e=this.getHandlers(),t=0;t<e.length;t++){var n=e[t];if(n.isSelected())return n}return null}},{key:"updatePreview",value:function(){var e=this.previewEl;e&&(e.style.backgroundImage=this.getValue("linear","to right"));}},{key:"initEvents",value:function(){var e=this,t=this.previewEl;t&&(0, u.on)(t,"click",(function(n){var r=e.options,o=r.min,i=r.max,a={w:t.clientWidth,h:t.clientHeight},l=n.offsetX-t.clientLeft,c=n.offsetY-t.clientTop,u=l/a.w*100;if(!(u>i||u<o||c>a.h||c<0)){var s=document.createElement("canvas"),f=s.getContext("2d");s.width=a.w,s.height=a.h;var d=f.createLinearGradient(0,0,a.w,a.h);e.getHandlers().forEach((function(e){return d.addColorStop(e.position/100,e.color)})),f.fillStyle=d,f.fillRect(0,0,s.width,s.height),s.style.background="black";var p=s.getContext("2d").getImageData(l,c,1,1).data,h="rgba("+p[0]+", "+p[1]+", "+p[2]+", "+p[3]+")",v="rgba(0, 0, 0, 0)"==h?r.emptyColor:h;e.addHandler(u,v);}}));}},{key:"render",value:function(){var e=this.options,t=this.el,n=e.pfx,r=e.height,o=e.width;if(t){var i=n+"-wrapper",a=n+"-preview";t.innerHTML='\n      <div class="'+i+'">\n        <div class="'+a+'"></div>\n      </div>\n    ';var l=t.querySelector("."+i),c=t.querySelector("."+a),u=l.style;u.position="relative",this.wrapperEl=l,this.previewEl=c,r&&(u.height=r),o&&(u.width=o),this.initEvents(),this.updatePreview();}}}]),t}(l.default);t.default=d;},function(e,t,n){Object.defineProperty(t,"__esModule",{value:!0});var r=function(){function e(e,t){for(var n=0;n<t.length;n++){var r=t[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),Object.defineProperty(e,r.key,r);}}return function(t,n,r){return n&&e(t.prototype,n),r&&e(t,r),t}}(),o=function(){function e(){!function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}(this,e);}return r(e,[{key:"on",value:function(e,t,n){var r=this.e||(this.e={});return (r[e]||(r[e]=[])).push({fn:t,ctx:n}),this}},{key:"once",value:function(e,t,n){function r(){o.off(e,r),t.apply(n,arguments);}var o=this;return r._=t,this.on(e,r,n)}},{key:"emit",value:function(e){for(var t=[].slice.call(arguments,1),n=((this.e||(this.e={}))[e]||[]).slice(),r=0,o=n.length;r<o;r++)n[r].fn.apply(n[r].ctx,t);return this}},{key:"off",value:function(e,t){var n=this.e||(this.e={}),r=n[e],o=[];if(r&&t)for(var i=0,a=r.length;i<a;i++)r[i].fn!==t&&r[i].fn._!==t&&o.push(r[i]);return o.length?n[e]=o:delete n[e],this}}]),e}();t.default=o;},function(e,t,n){function r(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}Object.defineProperty(t,"__esModule",{value:!0});var o=function(){function e(e,t){for(var n=0;n<t.length;n++){var r=t[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),Object.defineProperty(e,r.key,r);}}return function(t,n,r){return n&&e(t.prototype,n),r&&e(t,r),t}}(),i=n(0),a=function(){function e(t){var n=arguments.length>1&&void 0!==arguments[1]?arguments[1]:0,o=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"black",i=arguments.length>3&&void 0!==arguments[3]?arguments[3]:1,a=arguments.length>4&&void 0!==arguments[4]?arguments[4]:{};r(this,e),t.getHandlers().push(this),this.gp=t,this.position=n,this.color=o,this.selected=0,this.render(),i&&this.select(a);}return o(e,[{key:"toJSON",value:function(){return {position:this.position,selected:this.selected,color:this.color}}},{key:"setColor",value:function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:1;this.color=e,this.emit("handler:color:change",this,t);}},{key:"setPosition",value:function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:1,n=this.getEl();this.position=e,n&&(n.style.left=e+"%"),this.emit("handler:position:change",this,t);}},{key:"getColor",value:function(){return this.color}},{key:"getPosition",value:function(){var e=this.position,t=this.gp.options.onValuePos;return (0, i.isFunction)(t)?t(e):e}},{key:"isSelected",value:function(){return !!this.selected}},{key:"getValue",value:function(){return this.getColor()+" "+this.getPosition()+"%"}},{key:"select",value:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},t=this.getEl(),n=this.gp.getHandlers();!e.keepSelect&&n.forEach((function(e){return e.deselect()})),this.selected=1;var r=this.getSelectedCls();t&&(t.className+=" "+r),this.emit("handler:select",this);}},{key:"deselect",value:function(){var e=this.getEl();this.selected=0;var t=this.getSelectedCls();e&&(e.className=e.className.replace(t,"").trim()),this.emit("handler:deselect",this);}},{key:"getSelectedCls",value:function(){return this.gp.options.pfx+"-handler-selected"}},{key:"remove",value:function(){var e=this,t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},n=this.cpFn,r=this.getEl(),o=this.gp.getHandlers(),a=o.splice(o.indexOf(this),1)[0];return r&&r.parentNode.removeChild(r),!t.silent&&this.emit("handler:remove",a),(0, i.isFunction)(n)&&n(this),["el","gp"].forEach((function(t){return e[t]=0})),a}},{key:"getEl",value:function(){return this.el}},{key:"initEvents",value:function(){var e=this,t=this.getEl(),n=this.gp.previewEl,r=this.gp.options,o=r.min,a=r.max,l=t.querySelector("[data-toggle=handler-close]"),c=t.querySelector("[data-toggle=handler-color-c]"),u=t.querySelector("[data-toggle=handler-color-wrap]"),s=t.querySelector("[data-toggle=handler-color]"),f=t.querySelector("[data-toggle=handler-drag]"),d=function(t){var n=arguments.length>1&&void 0!==arguments[1]?arguments[1]:1,r=t.target.value;e.setColor(r,n),u&&(u.style.backgroundColor=r);};if(c&&(0, i.on)(c,"click",(function(e){return e.stopPropagation()})),l&&(0, i.on)(l,"click",(function(t){t.stopPropagation(),e.remove();})),s&&((0, i.on)(s,"change",d),(0, i.on)(s,"input",(function(e){return d(e,0)}))),f){var p=0,h=0,v=0,g={},y={},m={},b=function(t){var n=(0, i.getPointerEvent)(t);v=1,m.x=n.clientX-y.x,m.y=n.clientY-y.y,p=100*m.x,p/=g.w,p=(p=(p=h+p)<o?o:p)>a?a:p,e.setPosition(p,0),e.emit("handler:drag",e,p),(0, i.isDef)(t.button)&&0===t.which&&k(t);},k=function t(n){(0, i.off)(document,"touchmove mousemove",b),(0, i.off)(document,"touchend mouseup",t),v&&(v=0,e.setPosition(p),e.emit("handler:drag:end",e,p));};(0, i.on)(f,"touchstart mousedown",(function(t){if(!(0, i.isDef)(t.button)||0===t.button){e.select();var r=(0, i.getPointerEvent)(t);h=e.position,g.w=n.clientWidth,g.h=n.clientHeight,y.x=r.clientX,y.y=r.clientY,(0, i.on)(document,"touchmove mousemove",b),(0, i.on)(document,"touchend mouseup",k),e.emit("handler:drag:start",e);}})),(0, i.on)(f,"click",(function(e){return e.stopPropagation()}));}}},{key:"emit",value:function(){var e;(e=this.gp).emit.apply(e,arguments);}},{key:"render",value:function(){var e=this.gp,t=e.options,n=e.previewEl,r=e.colorPicker,o=t.pfx,i=t.colorEl,a=this.getColor();if(n){var l=document.createElement("div"),c=l.style,u=o+"-handler";return l.className=u,l.innerHTML='\n      <div class="'+u+'-close-c">\n        <div class="'+u+'-close" data-toggle="handler-close">&Cross;</div>\n      </div>\n      <div class="'+u+'-drag" data-toggle="handler-drag"></div>\n      <div class="'+u+'-cp-c" data-toggle="handler-color-c">\n        '+(i||'\n          <div class="'+u+'-cp-wrap" data-toggle="handler-color-wrap" style="background-color: '+a+'">\n            <input type="color" data-toggle="handler-color" value="'+a+'">\n          </div>')+"\n      </div>\n    ",c.position="absolute",c.top=0,c.left=this.position+"%",n.appendChild(l),this.el=l,this.initEvents(),this.cpFn=r&&r(this),l}}}]),e}();t.default=a;}]));},function(e,t,n){n.r(t);var r=n(0),o=n.n(r),i=n(1),a=n.n(i),l=n(2),c=n.n(l);function u(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r);}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?u(Object(n),!0).forEach((function(t){o()(e,t,n[t]);})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):u(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t));}));}return e}var f,d,p='data-cp',h=function(e){return (1==e.getAlpha()?e.toHexString():e.toRgbString()).replace(/ /g,'')},v=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{},n=e.StyleManager,r=t.selectEdgeStops,o=t.colorPicker,i={},l=['top','right','bottom','left'],u=function(e){i=e||{fromTarget:1,avoidStore:1},setTimeout((function(){return i={}}));};n.addType('gradient',{view:{events:{'click [data-clear-style]':'clear'},templateInput:function(){return ''},setValue:function(e){var t=this.gp;if(t){var n=this.model.getDefaultValue();e=e||n,u(),t.setValue(e);var o=t.getDirection(),i=l.filter((function(e){return o.indexOf(e)>=0}))[0]||o;t.setDirection(i),d&&d.setValue(t.getType()),f&&f.setValue(i);var a=t.getHandlers();r&&[a[0],a[a.length-1]].filter((function(e){return e})).map((function(e){return e.select({keepSelect:1})}));}},destroy:function(){var e=this.gp;e&&e.destroy();},onRender:function(){var r=this,v=this.ppfx,g=this.em,y=this.model,m=s(s({},t),y.get('gradientConfig')||{}),b=m.onCustomInputChange,k=document.createElement('div'),w=o&&"<div class=\"grp-handler-cp-wrap\">\n          <div class=\"".concat(v,"field-colorp-c\">\n            <div class=\"").concat(v,"checker-bg\"></div>\n            <div class=\"").concat(v,"field-color-picker\" ").concat(p,"></div>\n          </div>\n        </div>"),O=new c.a(s({el:k,colorEl:w},m.grapickOpts)),P=this.el.querySelector(".".concat(v,"fields"));P.style.flexWrap='wrap',P.appendChild(k.children[0]),this.gp=O,O.on('change',(function(e){y.setValueFromInput(O.getValue(),e,i);})),[['inputDirection','select','setDirection',{name:'Direction',property:'__gradient-direction',defaults:'right',options:l.map((function(e){return {value:e}}))}],['inputType','select','setType',{name:'Type',defaults:'linear',property:'__gradient-type',options:[{value:'radial'},{value:'linear'},{value:'repeating-radial'},{value:'repeating-linear'}]}]].forEach((function(e){var t=e[0],o=m[t];if(o){var i=y.parent,l=e[1],c='object'==a()(o)?o:{},p=n.createType(c.type||l,{model:s(s({},e[3]),c),view:{propTarget:r.propTarget}});i&&(p.model.parent=i),p.render(),p.model.on('change:value',(function(t,n){var r=arguments.length>2&&void 0!==arguments[2]?arguments[2]:{};u(r),O.el&&O[e[2]](t.getFullValue()||t.getDefaultValue(),{complete:!r.avoidStore}),b({model:t,input:e,inputDirection:f,inputType:d,opts:r});})),P.appendChild(p.el),'inputDirection'==t&&(f=p),'inputType'==t&&(d=p);}})),'default'==o&&(o=function(t){var n=t.getEl().querySelector("[".concat(p,"]")),r=n.style;r.backgroundColor=t.getColor();var o=g&&g.getConfig()||{},i=o.colorPicker||{},a=o.el,l=function(e){var n=arguments.length>1&&void 0!==arguments[1]?arguments[1]:1,o=h(e);r.backgroundColor=o,t.setColor(o,n);},c={color:t.getColor(),change:function(e){l(e);},move:function(e){l(e,0);}},u=g&&g.initBaseColorPicker;u?u(n,c):e.$(n).spectrum(s(s({containerClassName:"".concat(v,"one-bg ").concat(v,"two-color"),appendTo:a||'body',maxSelectionSize:8,showPalette:!0,palette:[],showAlpha:!0,chooseText:'Ok',cancelText:'⨯'},c),i));},O.on('handler:remove',(function(t){var n=t.getEl().querySelector("[".concat(p,"]")),r=e.$(n);r.spectrum&&r.spectrum('destroy');}))),o&&O.setColorPicker(o);}}});};function g(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r);}return n}function y(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?g(Object(n),!0).forEach((function(t){o()(e,t,n[t]);})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):g(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t));}));}return e}t["default"]=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{},n={grapickOpts:{},colorPicker:'',inputDirection:1,inputType:1,selectEdgeStops:1,onCustomInputChange:function(){return 0}},r=y(y({},n),t);v(e,r);};}])}));

    });

    var stgrPlugin = /*@__PURE__*/getDefaultExportFromCjs(grapesjsStyleGradient_min);

    var dist$1 = createCommonjsModule(function (module, exports) {
    /*! grapesjs-style-filter - 1.0.1 */
    !function(e,t){module.exports=t();}('undefined'!=typeof globalThis?globalThis:'undefined'!=typeof window?window:commonjsGlobal,(function(){return (()=>{var e={d:(t,r)=>{for(var n in r)e.o(r,n)&&!e.o(t,n)&&Object.defineProperty(t,n,{enumerable:!0,get:r[n]});},o:(e,t)=>Object.prototype.hasOwnProperty.call(e,t),r:e=>{'undefined'!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:'Module'}),Object.defineProperty(e,'__esModule',{value:!0});}},t={};function r(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function n(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n);}return r}e.r(t),e.d(t,{default:()=>o});const o=function(e){e.Styles.addBuiltIn('filter',{type:'stack',layerSeparator:' ',fromStyle:function(e,t){var r=t.property,n=e[t.name]||'',o=r.getLayerSeparator();return n?n.split(o).map((function(e){var t=r.__parseFn(e);return {name:t.name,value:t.value}})):[]},toStyle:function(e,t){return r({},t.name,"".concat(e.name,"(").concat(e.value,")"))},properties:[{property:'name',name:'Type',type:'select',default:'sepia',full:!0,options:[{id:'blur',propValue:{min:0,units:['px','em','rem','vw','vh']}},{id:'brightness',propValue:{min:0,units:['%']}},{id:'contrast',propValue:{min:0,units:['%']}},{id:'grayscale',propValue:{min:0,max:100,units:['%']}},{id:'hue-rotate',propValue:{min:0,max:360,units:['deg','rad','grad']}},{id:'invert',propValue:{min:0,max:100,units:['%']}},{id:'saturate',propValue:{min:0,units:['%']}},{id:'sepia',propValue:{min:0,max:100,units:['%']}}],onChange:function(e){var t=e.property;if(e.to.value){var o=function(e){for(var t=1;t<arguments.length;t++){var o=null!=arguments[t]?arguments[t]:{};t%2?n(Object(o),!0).forEach((function(t){r(e,t,o[t]);})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(o)):n(Object(o)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(o,t));}));}return e}({},t.getOption().propValue||{}),i=t.getParent().getProperty('value'),a=i.getUnit();(!a||(null==o?void 0:o.units.indexOf(a))<0)&&(o.unit=(null==o?void 0:o.units[0])||''),i.up(o);}}},{property:'value',type:'slider',default:'0',full:!0}]});};return t})()}));

    });

    var styleFilter = /*@__PURE__*/getDefaultExportFromCjs(dist$1);

    var grapesjsTabs_min = createCommonjsModule(function (module, exports) {
    /*! grapesjs-tabs - 1.0.6 */
    !function(t,e){module.exports=e();}(window,(function(){return function(t){var e={};function n(r){if(e[r])return e[r].exports;var o=e[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=t,n.c=e,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r});},n.r=function(t){'undefined'!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:'Module'}),Object.defineProperty(t,'__esModule',{value:!0});},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&'object'==typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,'default',{enumerable:!0,value:t}),2&e&&'string'!=typeof t)for(var o in t)n.d(r,o,function(e){return t[e]}.bind(null,o));return r},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,'a',e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="",n(n.s=3)}([function(t,e){t.exports=function(t,e,n){return e in t?Object.defineProperty(t,e,{value:n,enumerable:!0,configurable:!0,writable:!0}):t[e]=n,t};},function(t,e,n){var r=n(2);t.exports=function(t,e){if(null==t)return {};var n,o,a=r(t,e);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(t);for(o=0;o<c.length;o++)n=c[o],e.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(t,n)&&(a[n]=t[n]);}return a};},function(t,e){t.exports=function(t,e){if(null==t)return {};var n,r,o={},a=Object.keys(t);for(r=0;r<a.length;r++)n=a[r],e.indexOf(n)>=0||(o[n]=t[n]);return o};},function(t,e,n){n.r(e);var r=n(0),o=n.n(r),a=n(1),c=n.n(a);function i(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(t);e&&(r=r.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,r);}return n}function s(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?i(Object(n),!0).forEach((function(e){o()(t,e,n[e]);})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e));}));}return t}var b=function(t,e){var n=e.defaultModel,r=e.typeTabs,a=e.selectorTab,i=(e.editor,c()(e,["defaultModel","typeTabs","selectorTab","editor"])),b=[{full:1,type:'button',label:!1,text:'Style Active',command:function(t){var e=t.Panels.getButton('views','open-sm');e&&e.set('active',1);var n=".".concat(i.classTab,".").concat(i.classTabActive);t.StyleManager.setTarget(n,{targetIsClass:1});}}];t.addType(i.typeTab,{model:{defaults:s({name:'Tab',draggable:"[data-gjs-type=\"".concat(i.typeTabContainer,"\"]"),attributes:{role:"tab"},components:i.templateTab,classes:i.classTab,traits:b},i.tabProps),init:function(){this.on('removed',this.__onRemove);},__initTab:function(){if(!this.tabContent){var t=this.getTabContent();if(!t){var e,n=(t=this.getTabsType().getContentsType().append({type:i.typeTabContent,components:i.templateTabContent(this)})[0]).getId(),r=this.getId();t.addAttributes({id:n,'aria-labelledby':r,hidden:!0}),this.addAttributes((e={},o()(e,a,n),o()(e,"id",r),e)),this.tabContent=t;}this.tabContent=t;}},__onRemove:function(){var t=this.getTabContent();t&&t.remove(),this.getTabsType().trigger('rerender');},getTabsType:function(){return this.closestType(r)},getTabContent:function(){var t=this.getAttributes()[a],e=this.getTabsType();if(e&&t)return e.findContents().filter((function(e){return e.getId()==t}))[0]},clone:function(){var t=n.prototype.clone.apply(this,arguments);return t.addAttributes(o()({},a,'')),t}}});};function p(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(t);e&&(r=r.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,r);}return n}function l(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?p(Object(n),!0).forEach((function(e){o()(t,e,n[e]);})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):p(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e));}));}return t}var u=function(t,e){var n=e.typeTab,r=e.typeTabContent,o=e.typeTabContents,a=e.typeTabContainer,i=e.style,s=c()(e,["typeTab","typeTabContent","typeTabContents","typeTabContainer","style"]),b=s.typeTabs,p=[1,2,3].map((function(t){return {type:n}}));t.addType(b,{model:{defaults:l({name:'Tabs',classactive:s.classTabActive,selectortab:s.selectorTab,'script-props':['classactive','selectortab'],script:function(t){var e,n,r=this,o=t.classactive,a=t.selectortab,c=window,i=c.history,s=c._isEditor,b='[role=tab]',p=document,l=p.body,u=p.location,f=l.matchesSelector||l.webkitMatchesSelector||l.mozMatchesSelector||l.msMatchesSelector,y=function(t,e){for(var n=t||[],r=0;r<n.length;r++)e(n[r],r);},d=function(t){return t.getAttribute(a)},O=function(t,e){return t.querySelector(e)},g=function(){return r.querySelectorAll(b)},j=function(t,e){return !s&&(t.tabIndex=e)},h=function(t){y(g(),(function(t){t.className=t.className.replace(o,'').trim(),t.ariaSelected='false',j(t,'-1');})),y(r.querySelectorAll("[role=tabpanel]"),(function(t){return t.hidden=!0})),t.className+=' '+o,t.ariaSelected='true',j(t,'0');var e=d(t),n=e&&O(r,"#".concat(e));n&&(n.hidden=!1);},v=O(r,".".concat(o).concat(b));(v=v||(n=(u.hash||'').replace('#',''))&&O(r,(e=a,"".concat(b,"[").concat(e,"=").concat(n,"]")))||O(r,b))&&h(v),r.addEventListener('click',(function(t){var e=t.target,n=f.call(e,b);if(n||(e=function(t){var e;return y(g(),(function(n){e||n.contains(t)&&(e=n);})),e}(e))&&(n=1),n&&!t.__trg&&e.className.indexOf(o)<0){t.preventDefault(),t.__trg=1,h(e);var r=d(e);try{i&&i.pushState(null,null,"#".concat(r));}catch(t){}}}));},traits:[{full:1,type:'button',label:!1,text:'Add Tab',command:function(t){var e=t.getSelected();e&&e.addTab();}}],components:[{type:a,components:p},{type:o},i&&"<style>".concat(i(s),"</style>")]},s.tabsProps),init:function(){this.findTabs().map(this.__onTab),this.listenTo(this.getTabContainerType().components(),'add',this.__onTab);},__onTab:function(t,e){var n=arguments.length>2&&void 0!==arguments[2]?arguments[2]:{};!n.avoidStore&&!n.temporary&&t.__initTab&&t.__initTab();},getTabContainerType:function(){return this.findType(a)[0]},getContentsType:function(){return this.findType(o)[0]||this},findTabs:function(){return this.findType(n)},findContents:function(){return this.findType(r)},addTab:function(t){this.getTabContainerType().append({type:n,components:t});}}});};function f(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(t);e&&(r=r.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,r);}return n}function y(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?f(Object(n),!0).forEach((function(e){o()(t,e,n[e]);})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):f(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e));}));}return t}var d=function(t,e){t.addType(e.typeTabContent,{model:{defaults:y({name:'Tab Content',draggable:!1,copyable:!1,removable:!1,highlightable:!1,attributes:{role:"tabpanel"},classes:e.classTabContent,traits:[]},e.tabContentProps)}});};function O(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(t);e&&(r=r.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,r);}return n}function g(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?O(Object(n),!0).forEach((function(e){o()(t,e,n[e]);})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):O(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e));}));}return t}var j=function(t,e){t.addType(e.typeTabContents,{model:{defaults:g({name:'Tab Contents',draggable:!1,droppable:!1,copyable:!1,removable:!1,classes:e.classTabContents},e.tabContentsProps)}});};function h(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(t);e&&(r=r.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,r);}return n}function v(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?h(Object(n),!0).forEach((function(e){o()(t,e,n[e]);})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):h(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e));}));}return t}var T=function(t,e){t.addType(e.typeTabContainer,{model:{defaults:v({name:'Tab Container',draggable:"[data-gjs-type=\"".concat(e.typeTabs,"\"]"),droppable:"[data-gjs-type=\"".concat(e.typeTab,"\"]"),copyable:!1,removable:!1,highlightable:!1,attributes:{role:"tablist"},classes:e.classTabContainer},e.tabContainerProps)}});};function m(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(t);e&&(r=r.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,r);}return n}function P(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?m(Object(n),!0).forEach((function(e){o()(t,e,n[e]);})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):m(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e));}));}return t}var w=function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{},n=t.DomComponents,r=P(P({},e),{},{defaultModel:n.getType('default').model,editor:t});[b,u,d,j,T].map((function(t){return t(n,r)}));};function C(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(t);e&&(r=r.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,r);}return n}var S=function(t,e){var n=e.tabsBlock,r=e.typeTabs,a=t.BlockManager;n&&a.add(r,function(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?C(Object(n),!0).forEach((function(e){o()(t,e,n[e]);})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):C(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e));}));}return t}({media:"\n      <svg viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\">\n        <path d=\"M22 9.3c0-.8-.5-1.3-1.3-1.3H3.4C2.5 8 2 8.5 2 9.3v7.4c0 .8.5 1.3 1.3 1.3h17.4c.8 0 1.3-.5 1.3-1.3V9.3zM21 17H3V9h18v8z\" fill-rule=\"nonzero\"/><rect x=\"3\" y=\"5\" width=\"4\" height=\"2\" rx=\".5\"/><rect x=\"8\" y=\"5\" width=\"4\" height=\"2\" rx=\".5\"/><rect x=\"13\" y=\"5\" width=\"4\" height=\"2\" rx=\".5\"/>\n      </svg>\n    ",label:'Tabs',content:{type:r}},n));},D={tabsBlock:{},tabsProps:{},tabContainerProps:{},tabProps:{},tabContentProps:{},tabContentsProps:{},classTab:'tab',classTabContainer:'tab-container',classTabActive:'tab-active',classTabContent:'tab-content',classTabContents:'tab-contents',selectorTab:'aria-controls',typeTabs:'tabs',typeTabContainer:'tab-container',typeTab:'tab',typeTabContent:'tab-content',typeTabContents:'tab-contents',templateTab:function(t){return '<span data-gjs-highlightable="false">Tab</span>'},templateTabContent:function(t){return '<div>Tab Content</div>'},style:function(t){return "\n        .".concat(t.classTab," {\n            padding: 7px 14px;\n            display: inline-block;\n            border-radius: 3px;\n            margin-right: 10px;\n        }\n\n        .").concat(t.classTab,":focus {\n            outline: none;\n        }\n\n        .").concat(t.classTab,".").concat(t.classTabActive," {\n            background-color: #0d94e6;\n            color: white;\n        }\n\n        .").concat(t.classTabContainer," {\n            display: inline-block;\n        }\n\n        .").concat(t.classTabContent," {\n            animation: fadeEffect 1s;\n        }\n\n        .").concat(t.classTabContents," {\n            min-height: 100px;\n            padding: 10px;\n        }\n\n        @keyframes fadeEffect {\n            from {opacity: 0;}\n            to {opacity: 1;}\n        }\n    ")}};function x(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(t);e&&(r=r.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,r);}return n}function _(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?x(Object(n),!0).forEach((function(e){o()(t,e,n[e]);})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):x(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e));}));}return t}e.default=function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{},n=_(_({},D),e);w(t,n),S(t,n);};}])}));

    });

    var tabPlugin = /*@__PURE__*/getDefaultExportFromCjs(grapesjsTabs_min);

    var dist = createCommonjsModule(function (module, exports) {
    /*! grapesjs-tooltip - 0.1.7 */
    !function(t,e){module.exports=e();}('undefined'!=typeof globalThis?globalThis:'undefined'!=typeof window?window:commonjsGlobal,(()=>(()=>{var t={d:(e,n)=>{for(var o in n)t.o(n,o)&&!t.o(e,o)&&Object.defineProperty(e,o,{enumerable:!0,get:n[o]});},o:(t,e)=>Object.prototype.hasOwnProperty.call(t,e),r:t=>{'undefined'!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:'Module'}),Object.defineProperty(t,'__esModule',{value:!0});}},e={};t.r(e),t.d(e,{default:()=>o});var n=function(){return n=Object.assign||function(t){for(var e,n=1,o=arguments.length;n<o;n++)for(var a in e=arguments[n])Object.prototype.hasOwnProperty.call(e,a)&&(t[a]=e[a]);return t},n.apply(this,arguments)};const o=function(t,e){var o;void 0===e&&(e={});var a=n({id:'tooltip',labelTooltip:'Tooltip',blockTooltip:{},propsTooltip:{},extendTraits:function(t){return t},attrTooltip:'data-tooltip',classTooltip:'tooltip-component',style:'',styleAdditional:'',privateClasses:!0,stylableTooltip:['background-color','padding','padding-top','padding-right','padding-bottom','padding-left','font-family','font-size','font-weight','letter-spacing','color','line-height','text-align','border-radius','border-top-left-radius','border-top-right-radius','border-bottom-left-radius','border-bottom-right-radius','border','border-width','border-style','border-color'],showTooltipOnStyle:!0},e),r=a.propsTooltip,l=a.classTooltip,i=a.style,c=a.styleAdditional,s=a.privateClasses,p=a.stylableTooltip,d=a.showTooltipOnStyle,f=a.blockTooltip,m=a.extendTraits,u=a.id,b=a.labelTooltip,h=a.attrTooltip;f&&t.BlockManager.add(u,n({media:"<svg viewBox=\"0 0 24 24\">\n          <path d=\"M4 2h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2h-4l-4 4-4-4H4c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2m0 2v12h4.83L12 19.17 15.17 16H20V4H4z\"></path>\n        </svg>",label:b,category:'Extra',select:!0,content:{type:u}},f));var v="".concat(l,"__body"),g="".concat(l,"--empty"),y="".concat(h,"-visible"),w="".concat(h,"-pos"),x="".concat(h,"-length");s&&t.SelectorManager.getAll().add([{private:1,name:l},{private:1,name:v},{private:1,name:g}]),t.Components.addType(u,{isComponent:function(t){var e;return null===(e=t.hasAttribute)||void 0===e?void 0:e.call(t,h)},model:{defaults:n({name:b,classes:[l],attributes:(o={},o[h]=b,o),styles:(i||"\n          .".concat(l," {\n            position: relative;\n            display: inline-block;\n            vertical-align: top;\n          }\n\n          .").concat(g," {\n            width: 50px;\n            height: 50px;\n          }\n\n          .").concat(v,",\n          [").concat(h,"]::after {\n            font-family: Helvetica, sans-serif;\n            background: rgba(55, 61, 73, 0.95);\n            border-radius: 3px;\n            bottom: 100%;\n            color: #fff;\n            content: attr(").concat(h,");\n            display: block;\n            font-size: 12px;\n            left: 50%;\n            line-height: normal;\n            max-width: 32rem;\n            opacity: 0;\n            overflow: hidden;\n            padding: 8px 16px;\n            pointer-events: none;\n            position: absolute;\n            text-overflow: ellipsis;\n            transform: translate(-50%, 0);\n            transition: opacity 0.25s, transform 0.25s;\n            white-space: nowrap;\n            box-sizing: border-box;\n            z-index: 10;\n          }\n\n          [").concat(y,"=true]::after,\n          [").concat(h,"]:focus::after,\n          [").concat(h,"]:hover::after {\n            opacity: 1;\n            transform: translate(-50%, -0.5rem);\n          }\n\n          [").concat(w,"=right]::after {\n            bottom: 50%;\n            left: 100%;\n            transform: translate(0, 50%);\n          }\n\n          [").concat(w,"=right]:focus::after,\n          [").concat(w,"=right]:hover::after,\n          [").concat(y,"=true][").concat(w,"=right]::after {\n            transform: translate(0.5rem, 50%);\n          }\n\n          [").concat(w,"=bottom]::after {\n            bottom: auto;\n            top: 100%;\n            transform: translate(-50%, 0);\n          }\n\n          [").concat(w,"=bottom]:focus::after,\n          [").concat(w,"=bottom]:hover::after,\n          [").concat(y,"=true][").concat(w,"=bottom]::after {\n            transform: translate(-50%, 0.5rem);\n          }\n\n          [").concat(w,"=left]::after {\n            bottom: 50%;\n            left: auto;\n            right: 100%;\n            transform: translate(0, 50%);\n          }\n\n          [").concat(w,"=left]:focus::after,\n          [").concat(w,"=left]:hover::after,\n          [").concat(y,"=true][").concat(w,"=left]::after {\n            transform: translate(-0.5rem, 50%);\n          }\n\n          [").concat(x,"=small]::after {\n            white-space: normal;\n            width: 80px;\n          }\n\n          [").concat(x,"=medium]::after {\n            white-space: normal;\n            width: 150px;\n          }\n\n          [").concat(x,"=large]::after {\n            white-space: normal;\n            width: 300px;\n          }\n\n          [").concat(x,"=fit]::after {\n            white-space: normal;\n            width: 100%;\n          }\n\n          // IE 11 bugfix\n          button[").concat(h,"] {\n            overflow: visible;\n          }\n        "))+c,traits:m([{name:h,label:'Text'},{name:w,label:'Position',type:'select',options:[{value:'top',name:'Top'},{value:'right',name:'Right'},{value:'bottom',name:'Bottom'},{value:'left',name:'Left'}]},{name:x,label:'Length',type:'select',options:[{value:'',name:'One line'},{value:'small',name:'Small'},{value:'medium',name:'Medium'},{value:'large',name:'Large'},{value:'fit',name:'Fit'}]},{name:y,label:'Visible',type:'checkbox',valueTrue:'true'},{name:'style-tooltip',labelButton:'Style tooltip',type:'button',full:!0,command:function(t){var e,n=t.Panels.getButton('views','open-sm');null==n||n.set('active',!0);var o=t.Css.getRules(".".concat(v))[0];if(o.set('stylable',p),t.StyleManager.select(o),d){var a=t.getSelected();(null==a?void 0:a.is(u))&&(a.addAttributes(((e={})[y]='true',e)),t.once('style:target',(function(){var t;a.addAttributes(((t={})[y]='false',t));})));}}}])},r),init:function(){this.listenTo(this.components(),'add remove',this.checkEmpty),this.checkEmpty();},checkEmpty:function(){this[!this.components().length?'addClass':'removeClass']("".concat(g));}}});};return e})()));

    });

    var toolTipPlugin = /*@__PURE__*/getDefaultExportFromCjs(dist);

    var grapesjsTailwind_min = createCommonjsModule(function (module, exports) {
    /*! grapesjs-tailwind - 1.0.8 */
    !function(t,e){module.exports=e();}('undefined'!=typeof globalThis?globalThis:'undefined'!=typeof window?window:commonjsGlobal,(function(){return function(t){var e={};function l(i){if(e[i])return e[i].exports;var r=e[i]={i:i,l:!1,exports:{}};return t[i].call(r.exports,r,r.exports,l),r.l=!0,r.exports}return l.m=t,l.c=e,l.d=function(t,e,i){l.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:i});},l.r=function(t){'undefined'!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:'Module'}),Object.defineProperty(t,'__esModule',{value:!0});},l.t=function(t,e){if(1&e&&(t=l(t)),8&e)return t;if(4&e&&'object'==typeof t&&t&&t.__esModule)return t;var i=Object.create(null);if(l.r(i),Object.defineProperty(i,'default',{enumerable:!0,value:t}),2&e&&'string'!=typeof t)for(var r in t)l.d(i,r,function(e){return t[e]}.bind(null,r));return i},l.n=function(t){var e=t&&t.__esModule?function(){return t['default']}:function(){return t};return l.d(e,'a',e),e},l.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},l.p="",l(l.s=4)}([function(t,e,l){t.exports=l(3);},function(t,e){function l(t,e,l,i,r,a,s){try{var o=t[a](s),n=o.value;}catch(t){return void l(t)}o.done?e(n):Promise.resolve(n).then(i,r);}t.exports=function(t){return function(){var e=this,i=arguments;return new Promise((function(r,a){var s=t.apply(e,i);function o(t){l(s,r,a,o,n,"next",t);}function n(t){l(s,r,a,o,n,"throw",t);}o(void 0);}))}},t.exports["default"]=t.exports,t.exports.__esModule=!0;},function(t,e){t.exports=function(t,e,l){return e in t?Object.defineProperty(t,e,{value:l,enumerable:!0,configurable:!0,writable:!0}):t[e]=l,t},t.exports["default"]=t.exports,t.exports.__esModule=!0;},function(t,e,l){var i=function(t){var e,l=Object.prototype,i=l.hasOwnProperty,r="function"==typeof Symbol?Symbol:{},a=r.iterator||"@@iterator",s=r.asyncIterator||"@@asyncIterator",o=r.toStringTag||"@@toStringTag";function n(t,e,l){return Object.defineProperty(t,e,{value:l,enumerable:!0,configurable:!0,writable:!0}),t[e]}try{n({},"");}catch(t){n=function(t,e,l){return t[e]=l};}function c(t,e,l,i){var r=e&&e.prototype instanceof m?e:m,a=Object.create(r.prototype),s=new j(i||[]);return a._invoke=function(t,e,l){var i=h;return function(r,a){if(i===g)throw new Error("Generator is already running");if(i===f){if("throw"===r)throw a;return B()}for(l.method=r,l.arg=a;1;){var s=l.delegate;if(s){var o=E(s,l);if(o){if(o===p)continue;return o}}if("next"===l.method)l.sent=l._sent=l.arg;else if("throw"===l.method){if(i===h)throw i=f,l.arg;l.dispatchException(l.arg);}else "return"===l.method&&l.abrupt("return",l.arg);i=g;var n=d(t,e,l);if("normal"===n.type){if(i=l.done?f:x,n.arg===p)continue;return {value:n.arg,done:l.done}}"throw"===n.type&&(i=f,l.method="throw",l.arg=n.arg);}}}(t,l,s),a}function d(t,e,l){try{return {type:"normal",arg:t.call(e,l)}}catch(t){return {type:"throw",arg:t}}}t.wrap=c;var h="suspendedStart",x="suspendedYield",g="executing",f="completed",p={};function m(){}function u(){}function v(){}var w={};n(w,a,(function(){return this}));var y=Object.getPrototypeOf,b=y&&y(y(z([])));b&&b!==l&&i.call(b,a)&&(w=b);var k=v.prototype=m.prototype=Object.create(w);function A(t){["next","throw","return"].forEach((function(e){n(t,e,(function(t){return this._invoke(e,t)}));}));}function C(t,e){var l;this._invoke=function(r,a){function s(){return new e((function(l,s){!function l(r,a,s,o){var n=d(t[r],t,a);if("throw"!==n.type){var c=n.arg,h=c.value;return h&&"object"==typeof h&&i.call(h,"__await")?e.resolve(h.__await).then((function(t){l("next",t,s,o);}),(function(t){l("throw",t,s,o);})):e.resolve(h).then((function(t){c.value=t,s(c);}),(function(t){return l("throw",t,s,o)}))}o(n.arg);}(r,a,l,s);}))}return l=l?l.then(s,s):s()};}function E(t,l){var i=t.iterator[l.method];if(i===e){if(l.delegate=null,"throw"===l.method){if(t.iterator["return"]&&(l.method="return",l.arg=e,E(t,l),"throw"===l.method))return p;l.method="throw",l.arg=new TypeError("The iterator does not provide a 'throw' method");}return p}var r=d(i,t.iterator,l.arg);if("throw"===r.type)return l.method="throw",l.arg=r.arg,l.delegate=null,p;var a=r.arg;return a?a.done?(l[t.resultName]=a.value,l.next=t.nextLoc,"return"!==l.method&&(l.method="next",l.arg=e),l.delegate=null,p):a:(l.method="throw",l.arg=new TypeError("iterator result is not an object"),l.delegate=null,p)}function M(t){var e={tryLoc:t[0]};1 in t&&(e.catchLoc=t[1]),2 in t&&(e.finallyLoc=t[2],e.afterLoc=t[3]),this.tryEntries.push(e);}function F(t){var e=t.completion||{};e.type="normal",delete e.arg,t.completion=e;}function j(t){this.tryEntries=[{tryLoc:"root"}],t.forEach(M,this),this.reset(!0);}function z(t){if(t){var l=t[a];if(l)return l.call(t);if("function"==typeof t.next)return t;if(!isNaN(t.length)){var r=-1,s=function l(){for(;++r<t.length;)if(i.call(t,r))return l.value=t[r],l.done=!1,l;return l.value=e,l.done=!0,l};return s.next=s}}return {next:B}}function B(){return {value:e,done:!0}}return u.prototype=v,n(k,"constructor",v),n(v,"constructor",u),u.displayName=n(v,o,"GeneratorFunction"),t.isGeneratorFunction=function(t){var e="function"==typeof t&&t.constructor;return !!e&&(e===u||"GeneratorFunction"===(e.displayName||e.name))},t.mark=function(t){return Object.setPrototypeOf?Object.setPrototypeOf(t,v):(t.__proto__=v,n(t,o,"GeneratorFunction")),t.prototype=Object.create(k),t},t.awrap=function(t){return {__await:t}},A(C.prototype),n(C.prototype,s,(function(){return this})),t.AsyncIterator=C,t.async=function(e,l,i,r,a){void 0===a&&(a=Promise);var s=new C(c(e,l,i,r),a);return t.isGeneratorFunction(l)?s:s.next().then((function(t){return t.done?t.value:s.next()}))},A(k),n(k,o,"Generator"),n(k,a,(function(){return this})),n(k,"toString",(function(){return "[object Generator]"})),t.keys=function(t){var e=[];for(var l in t)e.push(l);return e.reverse(),function l(){for(;e.length;){var i=e.pop();if(i in t)return l.value=i,l.done=!1,l}return l.done=!0,l}},t.values=z,j.prototype={constructor:j,reset:function(t){if(this.prev=0,this.next=0,this.sent=this._sent=e,this.done=!1,this.delegate=null,this.method="next",this.arg=e,this.tryEntries.forEach(F),!t)for(var l in this)"t"===l.charAt(0)&&i.call(this,l)&&!isNaN(+l.slice(1))&&(this[l]=e);},stop:function(){this.done=!0;var t=this.tryEntries[0].completion;if("throw"===t.type)throw t.arg;return this.rval},dispatchException:function(t){if(this.done)throw t;var l=this;function r(i,r){return o.type="throw",o.arg=t,l.next=i,r&&(l.method="next",l.arg=e),!!r}for(var a=this.tryEntries.length-1;a>=0;--a){var s=this.tryEntries[a],o=s.completion;if("root"===s.tryLoc)return r("end");if(s.tryLoc<=this.prev){var n=i.call(s,"catchLoc"),c=i.call(s,"finallyLoc");if(n&&c){if(this.prev<s.catchLoc)return r(s.catchLoc,!0);if(this.prev<s.finallyLoc)return r(s.finallyLoc)}else if(n){if(this.prev<s.catchLoc)return r(s.catchLoc,!0)}else {if(!c)throw new Error("try statement without catch or finally");if(this.prev<s.finallyLoc)return r(s.finallyLoc)}}}},abrupt:function(t,e){for(var l=this.tryEntries.length-1;l>=0;--l){var r=this.tryEntries[l];if(r.tryLoc<=this.prev&&i.call(r,"finallyLoc")&&this.prev<r.finallyLoc){var a=r;break}}a&&("break"===t||"continue"===t)&&a.tryLoc<=e&&e<=a.finallyLoc&&(a=null);var s=a?a.completion:{};return s.type=t,s.arg=e,a?(this.method="next",this.next=a.finallyLoc,p):this.complete(s)},complete:function(t,e){if("throw"===t.type)throw t.arg;return "break"===t.type||"continue"===t.type?this.next=t.arg:"return"===t.type?(this.rval=this.arg=t.arg,this.method="return",this.next="end"):"normal"===t.type&&e&&(this.next=e),p},finish:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var l=this.tryEntries[e];if(l.finallyLoc===t)return this.complete(l.completion,l.afterLoc),F(l),p}},catch:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var l=this.tryEntries[e];if(l.tryLoc===t){var i=l.completion;if("throw"===i.type){var r=i.arg;F(l);}return r}}throw new Error("illegal catch attempt")},delegateYield:function(t,l,i){return this.delegate={iterator:z(t),resultName:l,nextLoc:i},"next"===this.method&&(this.arg=e),p}},t}(t.exports);try{regeneratorRuntime=i;}catch(t){"object"==typeof globalThis?globalThis.regeneratorRuntime=i:Function("r","regeneratorRuntime = r")(i);}},function(t,e,l){l.r(e);var i=l(1),r=l.n(i),a=l(2),s=l.n(a),o=l(0),n=l.n(o),c=[{id:'blog-block-1',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="20" y="43" width="68" height="63" rx="2" fill="#E2E8F0"></rect><path d="M29 73a1 1 0 011-1h48a1 1 0 110 2H30a1 1 0 01-1-1zM33 78a1 1 0 011-1h40a1 1 0 110 2H34a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M48 83a1 1 0 011-1h11a1 1 0 110 2H49a1 1 0 01-1-1z" fill="#6366F1"></path><path d="M37 67.5a1.5 1.5 0 011.5-1.5h32a1.5 1.5 0 010 3h-32a1.5 1.5 0 01-1.5-1.5z" fill="#4A5568"></path><rect x="99" y="43" width="68" height="63" rx="2" fill="#E2E8F0"></rect><path d="M108 73a1 1 0 011-1h48a1 1 0 010 2h-48a1 1 0 01-1-1zM112 78a1 1 0 011-1h40a1 1 0 010 2h-40a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M127 83a1 1 0 011-1h11a1 1 0 010 2h-11a1 1 0 01-1-1z" fill="#6366F1"></path><path d="M116 67.5a1.5 1.5 0 011.5-1.5h32a1.5 1.5 0 010 3h-32a1.5 1.5 0 01-1.5-1.5z" fill="#4A5568"></path><rect x="178" y="43" width="68" height="63" rx="2" fill="#E2E8F0"></rect><path d="M187 73a1 1 0 011-1h48a1 1 0 010 2h-48a1 1 0 01-1-1zM191 78a1 1 0 011-1h40a1 1 0 010 2h-40a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M206 83a1 1 0 011-1h11a1 1 0 010 2h-11a1 1 0 01-1-1z" fill="#6366F1"></path><path d="M195 67.5a1.5 1.5 0 011.5-1.5h32a1.5 1.5 0 010 3h-32a1.5 1.5 0 01-1.5-1.5z" fill="#4A5568"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-wrap -m-4\"><div class=\"p-4 lg:w-1/3\"><div class=\"h-full bg-gray-100 bg-opacity-75 px-8 pt-16 pb-24 rounded-lg overflow-hidden text-center relative\"><h2 class=\"tracking-widest text-xs title-font font-medium text-gray-400 mb-1\">CATEGORY</h2><h1 class=\"title-font sm:text-2xl text-xl font-medium text-gray-900 mb-3\">Raclette Blueberry Nextious Level</h1><p class=\"leading-relaxed mb-3\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p><a class=\"text-indigo-500 inline-flex items-center\">Learn More<svg class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path><path d=\"M12 5l7 7-7 7\"></path></svg></a><div class=\"text-center mt-2 leading-none flex justify-center absolute bottom-0 left-0 w-full py-4\"><span class=\"text-gray-400 mr-3 inline-flex items-center leading-none text-sm pr-3 py-1 border-r-2 border-gray-200\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z\"></path><circle cx=\"12\" cy=\"12\" r=\"3\"></circle></svg>1.2K</span><span class=\"text-gray-400 inline-flex items-center leading-none text-sm\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg>6</span></div></div></div><div class=\"p-4 lg:w-1/3\"><div class=\"h-full bg-gray-100 bg-opacity-75 px-8 pt-16 pb-24 rounded-lg overflow-hidden text-center relative\"><h2 class=\"tracking-widest text-xs title-font font-medium text-gray-400 mb-1\">CATEGORY</h2><h1 class=\"title-font sm:text-2xl text-xl font-medium text-gray-900 mb-3\">Ennui Snackwave Thundercats</h1><p class=\"leading-relaxed mb-3\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p><a class=\"text-indigo-500 inline-flex items-center\">Learn More<svg class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path><path d=\"M12 5l7 7-7 7\"></path></svg></a><div class=\"text-center mt-2 leading-none flex justify-center absolute bottom-0 left-0 w-full py-4\"><span class=\"text-gray-400 mr-3 inline-flex items-center leading-none text-sm pr-3 py-1 border-r-2 border-gray-200\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z\"></path><circle cx=\"12\" cy=\"12\" r=\"3\"></circle></svg>1.2K</span><span class=\"text-gray-400 inline-flex items-center leading-none text-sm\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg>6</span></div></div></div><div class=\"p-4 lg:w-1/3\"><div class=\"h-full bg-gray-100 bg-opacity-75 px-8 pt-16 pb-24 rounded-lg overflow-hidden text-center relative\"><h2 class=\"tracking-widest text-xs title-font font-medium text-gray-400 mb-1\">CATEGORY</h2><h1 class=\"title-font sm:text-2xl text-xl font-medium text-gray-900 mb-3\">Selvage Poke Waistcoat Godard</h1><p class=\"leading-relaxed mb-3\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p><a class=\"text-indigo-500 inline-flex items-center\">Learn More<svg class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path><path d=\"M12 5l7 7-7 7\"></path></svg></a><div class=\"text-center mt-2 leading-none flex justify-center absolute bottom-0 left-0 w-full py-4\"><span class=\"text-gray-400 mr-3 inline-flex items-center leading-none text-sm pr-3 py-1 border-r-2 border-gray-200\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z\"></path><circle cx=\"12\" cy=\"12\" r=\"3\"></circle></svg>1.2K</span><span class=\"text-gray-400 inline-flex items-center leading-none text-sm\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg>6</span></div></div></div></div></div></section>",category:'Blog'},{id:'blog-block-2',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="20.5" y="43.5" width="67" height="62" rx="1.5" fill="#FFFFFF" stroke="#CBD5E0"></rect><path d="M48.556 69h10.888c.86 0 1.556-.696 1.556-1.556V56.556c0-.86-.696-1.556-1.556-1.556H48.556c-.86 0-1.556.696-1.556 1.556v10.888c0 .86.696 1.556 1.556 1.556zm0 0l8.555-8.556L61 64.334m-8.556-5.056a1.167 1.167 0 11-2.333 0 1.167 1.167 0 012.333 0z" stroke="#A0AEC0" stroke-width="1.2px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><path d="M26 88a1 1 0 011-1h48a1 1 0 110 2H27a1 1 0 01-1-1zM26 93a1 1 0 011-1h40a1 1 0 110 2H27a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M26 98a1 1 0 011-1h11a1 1 0 110 2H27a1 1 0 01-1-1z" fill="#6366F1"></path><path d="M26 82.5a1.5 1.5 0 011.5-1.5h32a1.5 1.5 0 010 3h-32a1.5 1.5 0 01-1.5-1.5z" fill="#4A5568"></path><rect x="99.5" y="43.5" width="67" height="62" rx="1.5" fill="#FFFFFF" stroke="#CBD5E0"></rect><path d="M127.556 69h10.888c.86 0 1.556-.696 1.556-1.556V56.556c0-.86-.696-1.556-1.556-1.556h-10.888c-.86 0-1.556.696-1.556 1.556v10.888c0 .86.696 1.556 1.556 1.556zm0 0l8.555-8.556 3.889 3.89m-8.556-5.056a1.166 1.166 0 11-2.333 0 1.166 1.166 0 012.333 0z" stroke="#A0AEC0" stroke-width="1.2px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><path d="M105 88a1 1 0 011-1h48a1 1 0 010 2h-48a1 1 0 01-1-1zM105 93a1 1 0 011-1h40a1 1 0 010 2h-40a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M105 98a1 1 0 011-1h11a1 1 0 010 2h-11a1 1 0 01-1-1z" fill="#6366F1"></path><path d="M105 82.5a1.5 1.5 0 011.5-1.5h32a1.5 1.5 0 010 3h-32a1.5 1.5 0 01-1.5-1.5z" fill="#4A5568"></path><rect x="178.5" y="43.5" width="67" height="62" rx="1.5" fill="#FFFFFF" stroke="#CBD5E0"></rect><path d="M206.556 69h10.888c.86 0 1.556-.696 1.556-1.556V56.556c0-.86-.696-1.556-1.556-1.556h-10.888c-.86 0-1.556.696-1.556 1.556v10.888c0 .86.696 1.556 1.556 1.556zm0 0l8.555-8.556 3.889 3.89m-8.556-5.056a1.166 1.166 0 11-2.333 0 1.166 1.166 0 012.333 0z" stroke="#A0AEC0" stroke-width="1.2px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><path d="M184 88a1 1 0 011-1h48a1 1 0 010 2h-48a1 1 0 01-1-1zM184 93a1 1 0 011-1h40a1 1 0 010 2h-40a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M184 98a1 1 0 011-1h11a1 1 0 010 2h-11a1 1 0 01-1-1z" fill="#6366F1"></path><path d="M184 82.5a1.5 1.5 0 011.5-1.5h32a1.5 1.5 0 010 3h-32a1.5 1.5 0 01-1.5-1.5z" fill="#4A5568"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-wrap -m-4\"><div class=\"p-4 md:w-1/3\"><div class=\"h-full border-2 border-gray-200 border-opacity-60 rounded-lg overflow-hidden\"><img class=\"lg:h-48 md:h-36 w-full object-cover object-center\" src=\"https://dummyimage.com/720x400\" alt=\"blog\"><div class=\"p-6\"><h2 class=\"tracking-widest text-xs title-font font-medium text-gray-400 mb-1\">CATEGORY</h2><h1 class=\"title-font text-lg font-medium text-gray-900 mb-3\">The Catalyzer</h1><p class=\"leading-relaxed mb-3\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p><div class=\"flex items-center flex-wrap \"><a class=\"text-indigo-500 inline-flex items-center md:mb-2 lg:mb-0\">Learn More<svg class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path><path d=\"M12 5l7 7-7 7\"></path></svg></a><span class=\"text-gray-400 mr-3 inline-flex items-center lg:ml-auto md:ml-0 ml-auto leading-none text-sm pr-3 py-1 border-r-2 border-gray-200\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z\"></path><circle cx=\"12\" cy=\"12\" r=\"3\"></circle></svg>1.2K</span><span class=\"text-gray-400 inline-flex items-center leading-none text-sm\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg>6</span></div></div></div></div><div class=\"p-4 md:w-1/3\"><div class=\"h-full border-2 border-gray-200 border-opacity-60 rounded-lg overflow-hidden\"><img class=\"lg:h-48 md:h-36 w-full object-cover object-center\" src=\"https://dummyimage.com/721x401\" alt=\"blog\"><div class=\"p-6\"><h2 class=\"tracking-widest text-xs title-font font-medium text-gray-400 mb-1\">CATEGORY</h2><h1 class=\"title-font text-lg font-medium text-gray-900 mb-3\">The 400 Blows</h1><p class=\"leading-relaxed mb-3\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p><div class=\"flex items-center flex-wrap\"><a class=\"text-indigo-500 inline-flex items-center md:mb-2 lg:mb-0\">Learn More<svg class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path><path d=\"M12 5l7 7-7 7\"></path></svg></a><span class=\"text-gray-400 mr-3 inline-flex items-center lg:ml-auto md:ml-0 ml-auto leading-none text-sm pr-3 py-1 border-r-2 border-gray-200\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z\"></path><circle cx=\"12\" cy=\"12\" r=\"3\"></circle></svg>1.2K</span><span class=\"text-gray-400 inline-flex items-center leading-none text-sm\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg>6</span></div></div></div></div><div class=\"p-4 md:w-1/3\"><div class=\"h-full border-2 border-gray-200 border-opacity-60 rounded-lg overflow-hidden\"><img class=\"lg:h-48 md:h-36 w-full object-cover object-center\" src=\"https://dummyimage.com/722x402\" alt=\"blog\"><div class=\"p-6\"><h2 class=\"tracking-widest text-xs title-font font-medium text-gray-400 mb-1\">CATEGORY</h2><h1 class=\"title-font text-lg font-medium text-gray-900 mb-3\">Shooting Stars</h1><p class=\"leading-relaxed mb-3\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p><div class=\"flex items-center flex-wrap \"><a class=\"text-indigo-500 inline-flex items-center md:mb-2 lg:mb-0\">Learn More<svg class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path><path d=\"M12 5l7 7-7 7\"></path></svg></a><span class=\"text-gray-400 mr-3 inline-flex items-center lg:ml-auto md:ml-0 ml-auto leading-none text-sm pr-3 py-1 border-r-2 border-gray-200\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z\"></path><circle cx=\"12\" cy=\"12\" r=\"3\"></circle></svg>1.2K</span><span class=\"text-gray-400 inline-flex items-center leading-none text-sm\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg>6</span></div></div></div></div></div></div></section>",category:'Blog'},{id:'blog-block-3',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="20" y="48" width="70" height="5" rx="2.5" fill="#4A5568"></rect><rect x="20" y="36" width="34" height="8" rx="2" fill="#EBF4FF"></rect><rect x="20" y="58" width="92" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="66" width="82" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="74" width="68" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="82" width="18" height="4" rx="2" fill="#6366F1"></rect><rect x="23" y="39" width="28" height="2" rx="1" fill="#6366F1"></rect><path d="M121.5 93a.5.5 0 010 1h-101a.5.5 0 010-1h101z" fill="#E2E8F0"></path><circle cx="27.5" cy="107.5" r="7.5" fill="#E2E8F0"></circle><path d="M39 110a1 1 0 011-1h19a1 1 0 010 2H40a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="39" y="103" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="144" y="48" width="70" height="5" rx="2.5" fill="#4A5568"></rect><rect x="144" y="36" width="34" height="8" rx="2" fill="#EBF4FF"></rect><rect x="144" y="58" width="92" height="4" rx="2" fill="#A0AEC0"></rect><rect x="144" y="66" width="82" height="4" rx="2" fill="#A0AEC0"></rect><rect x="144" y="74" width="68" height="4" rx="2" fill="#A0AEC0"></rect><rect x="144" y="82" width="18" height="4" rx="2" fill="#6366F1"></rect><rect x="147" y="39" width="28" height="2" rx="1" fill="#6366F1"></rect><path d="M245.5 93a.5.5 0 010 1h-101a.5.5 0 010-1h101z" fill="#E2E8F0"></path><circle cx="151.5" cy="107.5" r="7.5" fill="#E2E8F0"></circle><path d="M163 110a1 1 0 011-1h19a1 1 0 010 2h-19a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="163" y="103" width="35" height="3" rx="1.5" fill="#4A5568"></rect></svg>',content:"<section class=\"text-gray-600 body-font overflow-hidden\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-wrap -m-12\"><div class=\"p-12 md:w-1/2 flex flex-col items-start\"><span class=\"inline-block py-1 px-2 rounded bg-indigo-50 text-indigo-500 text-xs font-medium tracking-widest\">CATEGORY</span><h2 class=\"sm:text-3xl text-2xl title-font font-medium text-gray-900 mt-4 mb-4\">Roof party normcore before they sold out, cornhole vape</h2><p class=\"leading-relaxed mb-8\">Live-edge letterpress cliche, salvia fanny pack humblebrag narwhal portland. VHS man braid palo santo hoodie brunch trust fund. Bitters hashtag waistcoat fashion axe chia unicorn. Plaid fixie chambray 90's, slow-carb etsy tumeric. Cray pug you probably haven't heard of them hexagon kickstarter craft beer pork chic.</p><div class=\"flex items-center flex-wrap pb-4 mb-4 border-b-2 border-gray-100 mt-auto w-full\"><a class=\"text-indigo-500 inline-flex items-center\">Learn More<svg class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path><path d=\"M12 5l7 7-7 7\"></path></svg></a><span class=\"text-gray-400 mr-3 inline-flex items-center ml-auto leading-none text-sm pr-3 py-1 border-r-2 border-gray-200\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z\"></path><circle cx=\"12\" cy=\"12\" r=\"3\"></circle></svg>1.2K</span><span class=\"text-gray-400 inline-flex items-center leading-none text-sm\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg>6</span></div><a class=\"inline-flex items-center\"><img alt=\"blog\" src=\"https://dummyimage.com/104x104\" class=\"w-12 h-12 rounded-full flex-shrink-0 object-cover object-center\"><span class=\"flex-grow flex flex-col pl-4\"><span class=\"title-font font-medium text-gray-900\">Holden Caulfield</span><span class=\"text-gray-400 text-xs tracking-widest mt-0.5\">UI DEVELOPER</span></span></a></div><div class=\"p-12 md:w-1/2 flex flex-col items-start\"><span class=\"inline-block py-1 px-2 rounded bg-indigo-50 text-indigo-500 text-xs font-medium tracking-widest\">CATEGORY</span><h2 class=\"sm:text-3xl text-2xl title-font font-medium text-gray-900 mt-4 mb-4\">Pinterest DIY dreamcatcher gentrify single-origin coffee</h2><p class=\"leading-relaxed mb-8\">Live-edge letterpress cliche, salvia fanny pack humblebrag narwhal portland. VHS man braid palo santo hoodie brunch trust fund. Bitters hashtag waistcoat fashion axe chia unicorn. Plaid fixie chambray 90's, slow-carb etsy tumeric.</p><div class=\"flex items-center flex-wrap pb-4 mb-4 border-b-2 border-gray-100 mt-auto w-full\"><a class=\"text-indigo-500 inline-flex items-center\">Learn More<svg class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path><path d=\"M12 5l7 7-7 7\"></path></svg></a><span class=\"text-gray-400 mr-3 inline-flex items-center ml-auto leading-none text-sm pr-3 py-1 border-r-2 border-gray-200\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z\"></path><circle cx=\"12\" cy=\"12\" r=\"3\"></circle></svg>1.2K</span><span class=\"text-gray-400 inline-flex items-center leading-none text-sm\"><svg class=\"w-4 h-4 mr-1\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg>6</span></div><a class=\"inline-flex items-center\"><img alt=\"blog\" src=\"https://dummyimage.com/103x103\" class=\"w-12 h-12 rounded-full flex-shrink-0 object-cover object-center\"><span class=\"flex-grow flex flex-col pl-4\"><span class=\"title-font font-medium text-gray-900\">Alper Kamu</span><span class=\"text-gray-400 text-xs tracking-widest mt-0.5\">DESIGNER</span></span></a></div></div></div></section>",category:'Blog'},{id:'blog-block-4',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="84" y="20" width="70" height="5" rx="2.5" fill="#4A5568"></rect><rect x="84" y="29" width="145" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="20" y="26" width="12" height="2" rx="1" fill="#A0AEC0"></rect><rect x="84" y="35" width="129" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="84" y="41" width="18" height="3" rx="1.5" fill="#6366F1"></rect><path d="M245.5 53a.5.5 0 010 1h-225a.5.5 0 010-1h225zM245.5 96a.5.5 0 010 1h-225a.5.5 0 010-1h225z" fill="#E2E8F0"></path><rect x="20" y="20" width="23" height="3" rx="1.5" fill="#4A5568"></rect><rect x="84" y="63" width="70" height="5" rx="2.5" fill="#4A5568"></rect><rect x="84" y="72" width="145" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="20" y="69" width="12" height="2" rx="1" fill="#A0AEC0"></rect><rect x="84" y="78" width="129" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="84" y="84" width="18" height="3" rx="1.5" fill="#6366F1"></rect><rect x="20" y="63" width="23" height="3" rx="1.5" fill="#4A5568"></rect><rect x="84" y="106" width="70" height="5" rx="2.5" fill="#4A5568"></rect><rect x="84" y="115" width="145" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="20" y="112" width="12" height="2" rx="1" fill="#A0AEC0"></rect><rect x="84" y="121" width="129" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="84" y="127" width="18" height="3" rx="1.5" fill="#6366F1"></rect><rect x="20" y="106" width="23" height="3" rx="1.5" fill="#4A5568"></rect></svg>',content:"<section class=\"text-gray-600 body-font overflow-hidden\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"-my-8 divide-y-2 divide-gray-100\"><div class=\"py-8 flex flex-wrap md:flex-nowrap\"><div class=\"md:w-64 md:mb-0 mb-6 flex-shrink-0 flex flex-col\"><span class=\"font-semibold title-font text-gray-700\">CATEGORY</span><span class=\"mt-1 text-gray-500 text-sm\">12 Jun 2019</span></div><div class=\"md:flex-grow\"><h2 class=\"text-2xl font-medium text-gray-900 title-font mb-2\">Bitters hashtag waistcoat fashion axe chia unicorn</h2><p class=\"leading-relaxed\">Glossier echo park pug, church-key sartorial biodiesel vexillologist pop-up snackwave ramps cornhole. Marfa 3 wolf moon party messenger bag selfies, poke vaporware kombucha lumbersexual pork belly polaroid hoodie portland craft beer.</p><a class=\"text-indigo-500 inline-flex items-center mt-4\">Learn More<svg class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path><path d=\"M12 5l7 7-7 7\"></path></svg></a></div></div><div class=\"py-8 flex flex-wrap md:flex-nowrap\"><div class=\"md:w-64 md:mb-0 mb-6 flex-shrink-0 flex flex-col\"><span class=\"font-semibold title-font text-gray-700\">CATEGORY</span><span class=\"mt-1 text-gray-500 text-sm\">12 Jun 2019</span></div><div class=\"md:flex-grow\"><h2 class=\"text-2xl font-medium text-gray-900 title-font mb-2\">Meditation bushwick direct trade taxidermy shaman</h2><p class=\"leading-relaxed\">Glossier echo park pug, church-key sartorial biodiesel vexillologist pop-up snackwave ramps cornhole. Marfa 3 wolf moon party messenger bag selfies, poke vaporware kombucha lumbersexual pork belly polaroid hoodie portland craft beer.</p><a class=\"text-indigo-500 inline-flex items-center mt-4\">Learn More<svg class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path><path d=\"M12 5l7 7-7 7\"></path></svg></a></div></div><div class=\"py-8 flex flex-wrap md:flex-nowrap\"><div class=\"md:w-64 md:mb-0 mb-6 flex-shrink-0 flex flex-col\"><span class=\"font-semibold title-font text-gray-700\">CATEGORY</span><span class=\"text-sm text-gray-500\">12 Jun 2019</span></div><div class=\"md:flex-grow\"><h2 class=\"text-2xl font-medium text-gray-900 title-font mb-2\">Woke master cleanse drinking vinegar salvia</h2><p class=\"leading-relaxed\">Glossier echo park pug, church-key sartorial biodiesel vexillologist pop-up snackwave ramps cornhole. Marfa 3 wolf moon party messenger bag selfies, poke vaporware kombucha lumbersexual pork belly polaroid hoodie portland craft beer.</p><a class=\"text-indigo-500 inline-flex items-center mt-4\">Learn More<svg class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path><path d=\"M12 5l7 7-7 7\"></path></svg></a></div></div></div></div></section>",category:'Blog'},{id:'blog-block-5',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M39 70.5a1.5 1.5 0 011.5-1.5h46a1.5 1.5 0 010 3h-46a1.5 1.5 0 01-1.5-1.5zM39 76.5a1.5 1.5 0 011.5-1.5h40a1.5 1.5 0 010 3h-40a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M39 58.5a1.5 1.5 0 011.5-1.5h14a1.5 1.5 0 010 3h-14a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><path d="M39 64.5a1.5 1.5 0 011.5-1.5h32a1.5 1.5 0 010 3h-32a1.5 1.5 0 01-1.5-1.5z" fill="#4A5568"></path><path d="M20 60a3 3 0 013-3h6a3 3 0 110 6h-6a3 3 0 01-3-3z" fill="#E2E8F0"></path><circle cx="44" cy="88" r="5" fill="#E2E8F0"></circle><path d="M51 87.5a1.5 1.5 0 011.5-1.5h18a1.5 1.5 0 010 3h-18a1.5 1.5 0 01-1.5-1.5z" fill="#4A5568"></path><path d="M118 70.5a1.5 1.5 0 011.5-1.5h46a1.5 1.5 0 010 3h-46a1.5 1.5 0 01-1.5-1.5zM118 76.5a1.5 1.5 0 011.5-1.5h40a1.5 1.5 0 010 3h-40a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M118 58.5a1.5 1.5 0 011.5-1.5h14a1.5 1.5 0 010 3h-14a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><path d="M118 64.5a1.5 1.5 0 011.5-1.5h32a1.5 1.5 0 010 3h-32a1.5 1.5 0 01-1.5-1.5z" fill="#4A5568"></path><path d="M99 60a3 3 0 013-3h6a3 3 0 110 6h-6a3 3 0 01-3-3z" fill="#E2E8F0"></path><circle cx="123" cy="88" r="5" fill="#E2E8F0"></circle><path d="M130 87.5a1.5 1.5 0 011.5-1.5h18a1.5 1.5 0 010 3h-18a1.5 1.5 0 01-1.5-1.5z" fill="#4A5568"></path><path d="M197 70.5a1.5 1.5 0 011.5-1.5h46a1.5 1.5 0 010 3h-46a1.5 1.5 0 01-1.5-1.5zM197 76.5a1.5 1.5 0 011.5-1.5h40a1.5 1.5 0 010 3h-40a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M197 58.5a1.5 1.5 0 011.5-1.5h14a1.5 1.5 0 010 3h-14a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><path d="M197 64.5a1.5 1.5 0 011.5-1.5h32a1.5 1.5 0 010 3h-32a1.5 1.5 0 01-1.5-1.5z" fill="#4A5568"></path><path d="M178 60a3 3 0 013-3h6a3 3 0 110 6h-6a3 3 0 01-3-3z" fill="#E2E8F0"></path><circle cx="202" cy="88" r="5" fill="#E2E8F0"></circle><path d="M209 87.5a1.5 1.5 0 011.5-1.5h18a1.5 1.5 0 010 3h-18a1.5 1.5 0 01-1.5-1.5z" fill="#4A5568"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-wrap -mx-4 -my-8\"><div class=\"py-8 px-4 lg:w-1/3\"><div class=\"h-full flex items-start\"><div class=\"w-12 flex-shrink-0 flex flex-col text-center leading-none\"><span class=\"text-gray-500 pb-2 mb-2 border-b-2 border-gray-200\">Jul</span><span class=\"font-medium text-lg text-gray-800 title-font leading-none\">18</span></div><div class=\"flex-grow pl-6\"><h2 class=\"tracking-widest text-xs title-font font-medium text-indigo-500 mb-1\">CATEGORY</h2><h1 class=\"title-font text-xl font-medium text-gray-900 mb-3\">The 400 Blows</h1><p class=\"leading-relaxed mb-5\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p><a class=\"inline-flex items-center\"><img alt=\"blog\" src=\"https://dummyimage.com/103x103\" class=\"w-8 h-8 rounded-full flex-shrink-0 object-cover object-center\"><span class=\"flex-grow flex flex-col pl-3\"><span class=\"title-font font-medium text-gray-900\">Alper Kamu</span></span></a></div></div></div><div class=\"py-8 px-4 lg:w-1/3\"><div class=\"h-full flex items-start\"><div class=\"w-12 flex-shrink-0 flex flex-col text-center leading-none\"><span class=\"text-gray-500 pb-2 mb-2 border-b-2 border-gray-200\">Jul</span><span class=\"font-medium text-lg text-gray-800 title-font leading-none\">18</span></div><div class=\"flex-grow pl-6\"><h2 class=\"tracking-widest text-xs title-font font-medium text-indigo-500 mb-1\">CATEGORY</h2><h1 class=\"title-font text-xl font-medium text-gray-900 mb-3\">Shooting Stars</h1><p class=\"leading-relaxed mb-5\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p><a class=\"inline-flex items-center\"><img alt=\"blog\" src=\"https://dummyimage.com/102x102\" class=\"w-8 h-8 rounded-full flex-shrink-0 object-cover object-center\"><span class=\"flex-grow flex flex-col pl-3\"><span class=\"title-font font-medium text-gray-900\">Holden Caulfield</span></span></a></div></div></div><div class=\"py-8 px-4 lg:w-1/3\"><div class=\"h-full flex items-start\"><div class=\"w-12 flex-shrink-0 flex flex-col text-center leading-none\"><span class=\"text-gray-500 pb-2 mb-2 border-b-2 border-gray-200\">Jul</span><span class=\"font-medium text-lg text-gray-800 title-font leading-none\">18</span></div><div class=\"flex-grow pl-6\"><h2 class=\"tracking-widest text-xs title-font font-medium text-indigo-500 mb-1\">CATEGORY</h2><h1 class=\"title-font text-xl font-medium text-gray-900 mb-3\">Neptune</h1><p class=\"leading-relaxed mb-5\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p><a class=\"inline-flex items-center\"><img alt=\"blog\" src=\"https://dummyimage.com/101x101\" class=\"w-8 h-8 rounded-full flex-shrink-0 object-cover object-center\"><span class=\"flex-grow flex flex-col pl-3\"><span class=\"title-font font-medium text-gray-900\">Henry Letham</span></span></a></div></div></div></div></div></section>",category:'Blog'},{id:'contact-block-1',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="153" y="30" width="93" height="90" rx="2" fill="#E2E8F0"></rect><rect x="162" y="101" width="75" height="10" rx="2" fill="#6366F1"></rect><rect x="162" y="66" width="75" height="30" rx="2" fill="#CBD5E0"></rect><rect x="162" y="51" width="75" height="10" rx="2" fill="#CBD5E0"></rect><rect x="162" y="40" width="40" height="4" rx="2" fill="#4A5568"></rect><path d="M89 71.682C89 81.546 76.5 90 76.5 90S64 81.546 64 71.682c0-3.364 1.317-6.59 3.661-8.968A12.41 12.41 0 0176.5 59a12.41 12.41 0 018.839 3.714A12.776 12.776 0 0189 71.682z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><path d="M76.5 75.91c2.301 0 4.167-1.894 4.167-4.228 0-2.335-1.866-4.228-4.167-4.228-2.301 0-4.167 1.893-4.167 4.228 0 2.334 1.866 4.227 4.167 4.227z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path></svg>',content:"<section class=\"text-gray-600 body-font relative\"><div class=\"absolute inset-0 bg-gray-300\"><iframe width=\"100%\" height=\"100%\" frameborder=\"0\" marginheight=\"0\" marginwidth=\"0\" title=\"map\" scrolling=\"no\" src=\"https://maps.google.com/maps?width=100%25&amp;height=600&amp;hl=en&amp;q=cyprus&amp;ie=UTF8&amp;t=&amp;z=14&amp;iwloc=B&amp;output=embed\" style=\"filter: grayscale(1) contrast(1.2) opacity(0.4);\"></iframe></div><div class=\"container px-5 py-24 mx-auto flex\"><div class=\"lg:w-1/3 md:w-1/2 bg-white rounded-lg p-8 flex flex-col md:ml-auto w-full mt-10 md:mt-0 relative z-10 shadow-md\"><form style=\"margin: 0;\"><h2 class=\"text-gray-900 text-lg mb-1 font-medium title-font\">Feedback</h2><p class=\"leading-relaxed mb-5 text-gray-600\">Post-ironic portland shabby chic echo park, banjo fashion axe</p><div class=\"relative mb-4\"><label for=\"email\" class=\"leading-7 text-sm text-gray-600\">Email</label><input type=\"email\" id=\"email\" name=\"email\" class=\"w-full bg-white rounded border border-gray-300 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div><div class=\"relative mb-4\"><label for=\"message\" class=\"leading-7 text-sm text-gray-600\">Message</label><textarea id=\"message\" name=\"message\" class=\"w-full bg-white rounded border border-gray-300 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 h-32 text-base outline-none text-gray-700 py-1 px-3 resize-none leading-6 transition-colors duration-200 ease-in-out\" required></textarea></div><button type=\"submit\" class=\"text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button><p class=\"text-xs text-gray-500 mt-3\">Chicharrones blog helvetica normcore iceland tousled brook viral artisan.</p></form></div></div></section>",category:'Contact'},{id:'contact-block-2',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="20" y="30" width="127" height="90" rx="2" fill="#E2E8F0"></rect><rect x="30" y="71" width="107" height="39" rx="2" fill="#FFFFFF"></rect><rect x="35" y="76" width="24" height="3" rx="1.5" fill="#4A5568"></rect><rect x="35" y="83" width="37" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="35" y="89" width="40" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="86" y="76" width="24" height="3" rx="1.5" fill="#4A5568"></rect><rect x="86" y="83" width="32" height="3" rx="1.5" fill="#6366F1"></rect><rect x="86" y="95" width="20" height="3" rx="1.5" fill="#4A5568"></rect><rect x="86" y="102" width="32" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="162" y="101" width="84" height="10" rx="2" fill="#6366F1"></rect><rect x="162" y="66" width="84" height="30" rx="2" fill="#CBD5E0"></rect><rect x="162" y="51" width="84" height="10" rx="2" fill="#CBD5E0"></rect><rect x="162" y="40" width="44.8" height="4" rx="2" fill="#4A5568"></rect><path d="M89 49.136C89 53.91 83 58 83 58s-6-4.09-6-8.864a6.21 6.21 0 011.757-4.339A5.933 5.933 0 0183 43c1.591 0 3.117.647 4.243 1.797A6.208 6.208 0 0189 49.137z" stroke="#A0AEC0" stroke-width="1.6px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><path d="M83 51.182c1.105 0 2-.916 2-2.046s-.895-2.045-2-2.045-2 .916-2 2.045c0 1.13.895 2.046 2 2.046z" stroke="#A0AEC0" stroke-width="1.6px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path></svg>',content:"<section class=\"text-gray-600 body-font relative\"><div class=\"container px-5 py-24 mx-auto flex sm:flex-nowrap flex-wrap\"><div class=\"lg:w-2/3 md:w-1/2 bg-gray-300 rounded-lg overflow-hidden sm:mr-10 p-10 flex items-end justify-start relative\"><iframe width=\"100%\" height=\"100%\" class=\"absolute inset-0\" frameborder=\"0\" title=\"map\" marginheight=\"0\" marginwidth=\"0\" scrolling=\"no\" src=\"https://maps.google.com/maps?width=100%25&height=600&hl=en&q=cyprus&ie=UTF8&t=&z=14&iwloc=B&output=embed\" style=\"filter: grayscale(1) contrast(1.2) opacity(0.4);\"></iframe><div class=\"bg-white relative flex flex-wrap py-6 rounded shadow-md\"><div class=\"lg:w-1/2 px-6\"><h2 class=\"title-font font-semibold text-gray-900 tracking-widest text-xs\">ADDRESS</h2><p class=\"mt-1\">Photo booth tattooed prism, portland taiyaki hoodie neutra typewriter</p></div><div class=\"lg:w-1/2 px-6 mt-4 lg:mt-0\"><h2 class=\"title-font font-semibold text-gray-900 tracking-widest text-xs\">EMAIL</h2><a class=\"text-indigo-500 leading-relaxed\">example@email.com</a><h2 class=\"title-font font-semibold text-gray-900 tracking-widest text-xs mt-4\">PHONE</h2><p class=\"leading-relaxed\">123-456-7890</p></div></div></div><div class=\"lg:w-1/3 md:w-1/2 bg-white flex flex-col md:ml-auto w-full md:py-8 mt-8 md:mt-0\"><form style=\"margin: 0;\"><h2 class=\"text-gray-900 text-lg mb-1 font-medium title-font\">Feedback</h2><p class=\"leading-relaxed mb-5 text-gray-600\">Post-ironic portland shabby chic echo park, banjo fashion axe</p><div class=\"relative mb-4\"><label for=\"name\" class=\"leading-7 text-sm text-gray-600\">Name</label><input type=\"text\" id=\"name\" name=\"name\" class=\"w-full bg-white rounded border border-gray-300 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div><div class=\"relative mb-4\"><label for=\"email\" class=\"leading-7 text-sm text-gray-600\">Email</label><input type=\"email\" id=\"email\" name=\"email\" class=\"w-full bg-white rounded border border-gray-300 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div><div class=\"relative mb-4\"><label for=\"message\" class=\"leading-7 text-sm text-gray-600\">Message</label><textarea id=\"message\" name=\"message\" class=\"w-full bg-white rounded border border-gray-300 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 h-32 text-base outline-none text-gray-700 py-1 px-3 resize-none leading-6 transition-colors duration-200 ease-in-out\" required></textarea></div><button type=\"submit\" class=\"text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button><p class=\"text-xs text-gray-500 mt-3\">Chicharrones blog helvetica normcore iceland tousled brook viral artisan.</p></form></div></div></section>",category:'Contact'},{id:'contact-block-3',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="113" y="117" width="40" height="10" rx="2" fill="#6366F1"></rect><rect x="63" y="81" width="140" height="30" rx="2" fill="#CBD5E0"></rect><rect x="63" y="65" width="66" height="10" rx="2" fill="#CBD5E0"></rect><rect x="135" y="65" width="68" height="10" rx="2" fill="#CBD5E0"></rect><rect x="90" y="24" width="86" height="5" rx="2.5" fill="#4A5568"></rect><rect x="80" y="33" width="106" height="4" rx="2" fill="#A0AEC0"></rect><rect x="85" y="41" width="97" height="4" rx="2" fill="#A0AEC0"></rect></svg>',content:"<section class=\"text-gray-600 body-font relative\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-col text-center w-full mb-12\"><h1 class=\"sm:text-3xl text-2xl font-medium title-font mb-4 text-gray-900\">Contact Us</h1><p class=\"lg:w-2/3 mx-auto leading-relaxed text-base\">Whatever cardigan tote bag tumblr hexagon brooklyn asymmetrical gentrify.</p></div><div class=\"lg:w-1/2 md:w-2/3 mx-auto\"><form style=\"margin: 0;\"><div class=\"flex flex-wrap -m-2\"><div class=\"p-2 w-1/2\"><div class=\"relative\"><label for=\"name\" class=\"leading-7 text-sm text-gray-600\">Name</label><input type=\"text\" id=\"name\" name=\"name\" class=\"w-full bg-gray-100 bg-opacity-50 rounded border border-gray-300 focus:border-indigo-500 focus:bg-white focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div></div><div class=\"p-2 w-1/2\"><div class=\"relative\"><label for=\"email\" class=\"leading-7 text-sm text-gray-600\">Email</label><input type=\"email\" id=\"email\" name=\"email\" class=\"w-full bg-gray-100 bg-opacity-50 rounded border border-gray-300 focus:border-indigo-500 focus:bg-white focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div></div><div class=\"p-2 w-full\"><div class=\"relative\"><label for=\"message\" class=\"leading-7 text-sm text-gray-600\">Message</label><textarea id=\"message\" name=\"message\" class=\"w-full bg-gray-100 bg-opacity-50 rounded border border-gray-300 focus:border-indigo-500 focus:bg-white focus:ring-2 focus:ring-indigo-200 h-32 text-base outline-none text-gray-700 py-1 px-3 resize-none leading-6 transition-colors duration-200 ease-in-out\" required></textarea></div></div><div class=\"p-2 w-full\"><button type=\"submit\" class=\"flex mx-auto text-white bg-indigo-500 border-0 py-2 px-8 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button></div><div class=\"p-2 w-full pt-8 mt-8 border-t border-gray-200 text-center\"><a class=\"text-indigo-500\">example@email.com</a><p class=\"leading-normal my-5\">49 Smith St.<br>Saint Cloud, MN 56301</p><span class=\"inline-flex\"><a class=\"text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-4 text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-4 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><rect width=\"20\" height=\"20\" x=\"2\" y=\"2\" rx=\"5\" ry=\"5\"></rect><path d=\"M16 11.37A4 4 0 1112.63 8 4 4 0 0116 11.37zm1.5-4.87h.01\"></path></svg></a><a class=\"ml-4 text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg></a></span></div></div></div></form></div></section>",category:'Contact'},{id:'content-block-1',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="113" y="120" width="40" height="10" rx="2" fill="#6366F1"></rect><rect x="81" y="31" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="96" y="20" width="74" height="5" rx="2.5" fill="#4A5568"></rect><rect x="85" y="39" width="97.365" height="4" rx="2" fill="#A0AEC0"></rect><rect x="26" y="73" width="28" height="3" rx="1.5" fill="#4A5568"></rect><rect x="26" y="79" width="34" height="2" rx="1" fill="#A0AEC0"></rect><rect x="26" y="84" width="38" height="2" rx="1" fill="#A0AEC0"></rect><rect x="26" y="89" width="24" height="2" rx="1" fill="#6366F1"></rect><rect x="20" y="62" width="1" height="39" rx="0.5" fill="#CBD5E0"></rect><rect x="86" y="73" width="28" height="3" rx="1.5" fill="#4A5568"></rect><rect x="86" y="79" width="34" height="2" rx="1" fill="#A0AEC0"></rect><rect x="86" y="84" width="38" height="2" rx="1" fill="#A0AEC0"></rect><rect x="86" y="89" width="24" height="2" rx="1" fill="#6366F1"></rect><rect x="80" y="62" width="1" height="39" rx="0.5" fill="#CBD5E0"></rect><rect x="146.136" y="73" width="28.636" height="3" rx="1.5" fill="#4A5568"></rect><rect x="146.136" y="79" width="34.773" height="2" rx="1" fill="#A0AEC0"></rect><rect x="146.136" y="84" width="38.864" height="2" rx="1" fill="#A0AEC0"></rect><rect x="146.136" y="89" width="24.546" height="2" rx="1" fill="#6366F1"></rect><rect x="140" y="62" width="1.023" height="39" rx="0.511" fill="#CBD5E0"></rect><rect x="207.136" y="73" width="28.636" height="3" rx="1.5" fill="#4A5568"></rect><rect x="207.136" y="79" width="34.773" height="2" rx="1" fill="#A0AEC0"></rect><rect x="207.136" y="84" width="38.864" height="2" rx="1" fill="#A0AEC0"></rect><rect x="207.136" y="89" width="24.546" height="2" rx="1" fill="#6366F1"></rect><rect x="201" y="62" width="1.023" height="39" rx="0.511" fill="#CBD5E0"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-col text-center w-full mb-20\"><h2 class=\"text-xs text-indigo-500 tracking-widest font-medium title-font mb-1\">ROOF PARTY POLAROID</h2><h1 class=\"sm:text-3xl text-2xl font-medium title-font mb-4 text-gray-900\">Master Cleanse Reliac Heirloom</h1><p class=\"lg:w-2/3 mx-auto leading-relaxed text-base\">Whatever cardigan tote bag tumblr hexagon brooklyn asymmetrical gentrify, subway tile poke farm-to-table. Franzen you probably haven't heard of them man bun deep jianbing selfies heirloom prism food truck ugh squid celiac humblebrag.</p></div><div class=\"flex flex-wrap\"><div class=\"xl:w-1/4 lg:w-1/2 md:w-full px-8 py-6 border-l-2 border-gray-200 border-opacity-60\"><h2 class=\"text-lg sm:text-xl text-gray-900 font-medium title-font mb-2\">Shooting Stars</h2><p class=\"leading-relaxed text-base mb-4\">Fingerstache flexitarian street art 8-bit waistcoat. Distillery hexagon disrupt edison bulbche.</p><a class=\"text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div><div class=\"xl:w-1/4 lg:w-1/2 md:w-full px-8 py-6 border-l-2 border-gray-200 border-opacity-60\"><h2 class=\"text-lg sm:text-xl text-gray-900 font-medium title-font mb-2\">The Catalyzer</h2><p class=\"leading-relaxed text-base mb-4\">Fingerstache flexitarian street art 8-bit waistcoat. Distillery hexagon disrupt edison bulbche.</p><a class=\"text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div><div class=\"xl:w-1/4 lg:w-1/2 md:w-full px-8 py-6 border-l-2 border-gray-200 border-opacity-60\"><h2 class=\"text-lg sm:text-xl text-gray-900 font-medium title-font mb-2\">Neptune</h2><p class=\"leading-relaxed text-base mb-4\">Fingerstache flexitarian street art 8-bit waistcoat. Distillery hexagon disrupt edison bulbche.</p><a class=\"text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div><div class=\"xl:w-1/4 lg:w-1/2 md:w-full px-8 py-6 border-l-2 border-gray-200 border-opacity-60\"><h2 class=\"text-lg sm:text-xl text-gray-900 font-medium title-font mb-2\">Melanchole</h2><p class=\"leading-relaxed text-base mb-4\">Fingerstache flexitarian street art 8-bit waistcoat. Distillery hexagon disrupt edison bulbche.</p><a class=\"text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div><button class=\"flex mx-auto mt-16 text-white bg-indigo-500 border-0 py-2 px-8 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button></div></section>",category:'Content'},{id:'content-block-2',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="142" y="32" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="32" width="74" height="5" rx="2.5" fill="#4A5568"></rect><rect x="142" y="40" width="77" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="74" width="50" height="44" rx="2" fill="#E2E8F0"></rect><path d="M40.333 95h9.334c.736 0 1.333-.597 1.333-1.333v-9.334c0-.736-.597-1.333-1.333-1.333h-9.334c-.736 0-1.333.597-1.333 1.333v9.334c0 .736.597 1.333 1.333 1.333zm0 0l7.334-7.333L51 91m-7.333-4.333a1 1 0 11-2 0 1 1 0 012 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="24" y="108" width="28" height="2" rx="1" fill="#4A5568"></rect><rect x="24" y="112" width="34" height="2" rx="1" fill="#A0AEC0"></rect><rect x="24" y="104" width="10" height="2" rx="1" fill="#6366F1"></rect><rect x="79" y="74" width="50" height="44" rx="2" fill="#E2E8F0"></rect><path d="M99.333 95h9.334c.736 0 1.333-.597 1.333-1.333v-9.334c0-.736-.597-1.333-1.333-1.333h-9.334c-.736 0-1.333.597-1.333 1.333v9.334c0 .736.597 1.333 1.333 1.333zm0 0l7.334-7.333L110 91m-7.333-4.333a1 1 0 11-2 0 1 1 0 012 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="83" y="108" width="28" height="2" rx="1" fill="#4A5568"></rect><rect x="83" y="112" width="34" height="2" rx="1" fill="#A0AEC0"></rect><rect x="83" y="104" width="10" height="2" rx="1" fill="#6366F1"></rect><rect x="138" y="74" width="50" height="44" rx="2" fill="#E2E8F0"></rect><path d="M158.333 95h9.334c.736 0 1.333-.597 1.333-1.333v-9.334c0-.736-.597-1.333-1.333-1.333h-9.334c-.736 0-1.333.597-1.333 1.333v9.334c0 .736.597 1.333 1.333 1.333zm0 0l7.334-7.333L169 91m-7.333-4.333a1 1 0 11-2 0 1 1 0 012 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="142" y="108" width="28" height="2" rx="1" fill="#4A5568"></rect><rect x="142" y="112" width="34" height="2" rx="1" fill="#A0AEC0"></rect><rect x="142" y="104" width="10" height="2" rx="1" fill="#6366F1"></rect><rect x="197" y="74" width="50" height="44" rx="2" fill="#E2E8F0"></rect><path d="M217.333 95h9.334c.736 0 1.333-.597 1.333-1.333v-9.334c0-.736-.597-1.333-1.333-1.333h-9.334c-.736 0-1.333.597-1.333 1.333v9.334c0 .736.597 1.333 1.333 1.333zm0 0l7.334-7.333L228 91m-7.333-4.333a1 1 0 11-2 0 1 1 0 012 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="201" y="108" width="28" height="2" rx="1" fill="#4A5568"></rect><rect x="201" y="112" width="34" height="2" rx="1" fill="#A0AEC0"></rect><rect x="201" y="104" width="10" height="2" rx="1" fill="#6366F1"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-wrap w-full mb-20\"><div class=\"lg:w-1/2 w-full mb-6 lg:mb-0\"><h1 class=\"sm:text-3xl text-2xl font-medium title-font mb-2 text-gray-900\">Pitchfork Kickstarter Taxidermy</h1><div class=\"h-1 w-20 bg-indigo-500 rounded\"></div></div><p class=\"lg:w-1/2 w-full leading-relaxed text-gray-500\">Whatever cardigan tote bag tumblr hexagon brooklyn asymmetrical gentrify, subway tile poke farm-to-table. Franzen you probably haven't heard of them man bun deep jianbing selfies heirloom prism food truck ugh squid celiac humblebrag.</p></div><div class=\"flex flex-wrap -m-4\"><div class=\"xl:w-1/4 md:w-1/2 p-4\"><div class=\"bg-gray-100 p-6 rounded-lg\"><img class=\"h-40 rounded w-full object-cover object-center mb-6\" src=\"https://dummyimage.com/720x400\" alt=\"content\"><h3 class=\"tracking-widest text-indigo-500 text-xs font-medium title-font\">SUBTITLE</h3><h2 class=\"text-lg text-gray-900 font-medium title-font mb-4\">Chichen Itza</h2><p class=\"leading-relaxed text-base\">Fingerstache flexitarian street art 8-bit waistcoat. Distillery hexagon disrupt edison bulbche.</p></div></div><div class=\"xl:w-1/4 md:w-1/2 p-4\"><div class=\"bg-gray-100 p-6 rounded-lg\"><img class=\"h-40 rounded w-full object-cover object-center mb-6\" src=\"https://dummyimage.com/721x401\" alt=\"content\"><h3 class=\"tracking-widest text-indigo-500 text-xs font-medium title-font\">SUBTITLE</h3><h2 class=\"text-lg text-gray-900 font-medium title-font mb-4\">Colosseum Roma</h2><p class=\"leading-relaxed text-base\">Fingerstache flexitarian street art 8-bit waistcoat. Distillery hexagon disrupt edison bulbche.</p></div></div><div class=\"xl:w-1/4 md:w-1/2 p-4\"><div class=\"bg-gray-100 p-6 rounded-lg\"><img class=\"h-40 rounded w-full object-cover object-center mb-6\" src=\"https://dummyimage.com/722x402\" alt=\"content\"><h3 class=\"tracking-widest text-indigo-500 text-xs font-medium title-font\">SUBTITLE</h3><h2 class=\"text-lg text-gray-900 font-medium title-font mb-4\">Great Pyramid of Giza</h2><p class=\"leading-relaxed text-base\">Fingerstache flexitarian street art 8-bit waistcoat. Distillery hexagon disrupt edison bulbche.</p></div></div><div class=\"xl:w-1/4 md:w-1/2 p-4\"><div class=\"bg-gray-100 p-6 rounded-lg\"><img class=\"h-40 rounded w-full object-cover object-center mb-6\" src=\"https://dummyimage.com/723x403\" alt=\"content\"><h3 class=\"tracking-widest text-indigo-500 text-xs font-medium title-font\">SUBTITLE</h3><h2 class=\"text-lg text-gray-900 font-medium title-font mb-4\">San Francisco</h2><p class=\"leading-relaxed text-base\">Fingerstache flexitarian street art 8-bit waistcoat. Distillery hexagon disrupt edison bulbche.</p></div></div></div></div></section>",category:'Content'},{id:'content-block-3',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="113" y="120" width="40" height="10" rx="2" fill="#6366F1"></rect><rect x="81" y="31" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="96" y="20" width="74" height="5" rx="2.5" fill="#4A5568"></rect><rect x="20.5" y="51.5" width="69" height="23" rx="1.5" fill="#FFFFFF" stroke="#CBD5E0"></rect><rect x="24" y="65" width="28" height="2" rx="1" fill="#4A5568"></rect><rect x="24" y="69" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="28" cy="59" r="4" fill="#C3DAFE"></circle><rect x="98.5" y="51.5" width="69" height="23" rx="1.5" fill="#FFFFFF" stroke="#CBD5E0"></rect><rect x="102" y="65" width="28" height="2" rx="1" fill="#4A5568"></rect><rect x="102" y="69" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="106" cy="59" r="4" fill="#C3DAFE"></circle><rect x="176.5" y="51.5" width="69" height="23" rx="1.5" fill="#FFFFFF" stroke="#CBD5E0"></rect><rect x="180" y="65" width="28" height="2" rx="1" fill="#4A5568"></rect><rect x="180" y="69" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="184" cy="59" r="4" fill="#C3DAFE"></circle><rect x="20.5" y="81.5" width="69" height="23" rx="1.5" fill="#FFFFFF" stroke="#CBD5E0"></rect><rect x="24" y="95" width="28" height="2" rx="1" fill="#4A5568"></rect><rect x="24" y="99" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="28" cy="89" r="4" fill="#C3DAFE"></circle><rect x="98.5" y="81.5" width="69" height="23" rx="1.5" fill="#FFFFFF" stroke="#CBD5E0"></rect><rect x="102" y="95" width="28" height="2" rx="1" fill="#4A5568"></rect><rect x="102" y="99" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="106" cy="89" r="4" fill="#C3DAFE"></circle><rect x="176.5" y="81.5" width="69" height="23" rx="1.5" fill="#FFFFFF" stroke="#CBD5E0"></rect><rect x="180" y="95" width="28" height="2" rx="1" fill="#4A5568"></rect><rect x="180" y="99" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="184" cy="89" r="4" fill="#C3DAFE"></circle></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-wrap w-full mb-20 flex-col items-center text-center\"><h1 class=\"sm:text-3xl text-2xl font-medium title-font mb-2 text-gray-900\">Pitchfork Kickstarter Taxidermy</h1><p class=\"lg:w-1/2 w-full leading-relaxed text-gray-500\">Whatever cardigan tote bag tumblr hexagon brooklyn asymmetrical gentrify, subway tile poke farm-to-table.</p></div><div class=\"flex flex-wrap -m-4\"><div class=\"xl:w-1/3 md:w-1/2 p-4\"><div class=\"border border-gray-200 p-6 rounded-lg\"><div class=\"w-10 h-10 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-4\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-6 h-6\" viewBox=\"0 0 24 24\"><path d=\"M22 12h-4l-3 9L9 3l-3 9H2\"></path></svg></div><h2 class=\"text-lg text-gray-900 font-medium title-font mb-2\">Shooting Stars</h2><p class=\"leading-relaxed text-base\">Fingerstache flexitarian street art 8-bit waist co, subway tile poke farm.</p></div></div><div class=\"xl:w-1/3 md:w-1/2 p-4\"><div class=\"border border-gray-200 p-6 rounded-lg\"><div class=\"w-10 h-10 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-4\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-6 h-6\" viewBox=\"0 0 24 24\"><circle cx=\"6\" cy=\"6\" r=\"3\"></circle><circle cx=\"6\" cy=\"18\" r=\"3\"></circle><path d=\"M20 4L8.12 15.88M14.47 14.48L20 20M8.12 8.12L12 12\"></path></svg></div><h2 class=\"text-lg text-gray-900 font-medium title-font mb-2\">The Catalyzer</h2><p class=\"leading-relaxed text-base\">Fingerstache flexitarian street art 8-bit waist co, subway tile poke farm.</p></div></div><div class=\"xl:w-1/3 md:w-1/2 p-4\"><div class=\"border border-gray-200 p-6 rounded-lg\"><div class=\"w-10 h-10 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-4\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-6 h-6\" viewBox=\"0 0 24 24\"><path d=\"M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2\"></path><circle cx=\"12\" cy=\"7\" r=\"4\"></circle></svg></div><h2 class=\"text-lg text-gray-900 font-medium title-font mb-2\">Neptune</h2><p class=\"leading-relaxed text-base\">Fingerstache flexitarian street art 8-bit waist co, subway tile poke farm.</p></div></div><div class=\"xl:w-1/3 md:w-1/2 p-4\"><div class=\"border border-gray-200 p-6 rounded-lg\"><div class=\"w-10 h-10 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-4\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-6 h-6\" viewBox=\"0 0 24 24\"><path d=\"M4 15s1-1 4-1 5 2 8 2 4-1 4-1V3s-1 1-4 1-5-2-8-2-4 1-4 1zM4 22v-7\"></path></svg></div><h2 class=\"text-lg text-gray-900 font-medium title-font mb-2\">Melanchole</h2><p class=\"leading-relaxed text-base\">Fingerstache flexitarian street art 8-bit waist co, subway tile poke farm.</p></div></div><div class=\"xl:w-1/3 md:w-1/2 p-4\"><div class=\"border border-gray-200 p-6 rounded-lg\"><div class=\"w-10 h-10 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-4\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-6 h-6\" viewBox=\"0 0 24 24\"><path d=\"M21 12.79A9 9 0 1111.21 3 7 7 0 0021 12.79z\"></path></svg></div><h2 class=\"text-lg text-gray-900 font-medium title-font mb-2\">Bunker</h2><p class=\"leading-relaxed text-base\">Fingerstache flexitarian street art 8-bit waist co, subway tile poke farm.</p></div></div><div class=\"xl:w-1/3 md:w-1/2 p-4\"><div class=\"border border-gray-200 p-6 rounded-lg\"><div class=\"w-10 h-10 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-4\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-6 h-6\" viewBox=\"0 0 24 24\"><path d=\"M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z\"></path></svg></div><h2 class=\"text-lg text-gray-900 font-medium title-font mb-2\">Ramona Falls</h2><p class=\"leading-relaxed text-base\">Fingerstache flexitarian street art 8-bit waist co, subway tile poke farm.</p></div></div></div><button class=\"flex mx-auto mt-16 text-white bg-indigo-500 border-0 py-2 px-8 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button></div></section>",category:'Content'},{id:'content-block-4',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="20" y="59" width="70" height="5" rx="2.5" fill="#4A5568"></rect><rect x="20" y="70" width="92" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="78" width="79" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="88" width="24" height="4" rx="2" fill="#6366F1"></rect><rect x="144" y="65" width="40" height="4" rx="2" fill="#4A5568"></rect><rect x="144" y="74" width="22" height="2" rx="1" fill="#A0AEC0"></rect><rect x="144" y="79" width="28" height="2" rx="1" fill="#A0AEC0"></rect><rect x="144" y="84" width="19" height="2" rx="1" fill="#A0AEC0"></rect><rect x="180" y="74" width="18" height="2" rx="1" fill="#A0AEC0"></rect><rect x="180" y="79" width="24" height="2" rx="1" fill="#A0AEC0"></rect><rect x="180" y="84" width="24" height="2" rx="1" fill="#A0AEC0"></rect><rect x="212" y="74" width="18" height="2" rx="1" fill="#A0AEC0"></rect><rect x="212" y="79" width="24" height="2" rx="1" fill="#A0AEC0"></rect><rect x="212" y="84" width="24" height="2" rx="1" fill="#A0AEC0"></rect><path d="M128 44.5a.5.5 0 011 0v62a.5.5 0 01-1 0v-62z" fill="#CBD5E0"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container flex flex-wrap px-5 py-24 mx-auto items-center\"><div class=\"md:w-1/2 md:pr-12 md:py-8 md:border-r md:border-b-0 mb-10 md:mb-0 pb-10 border-b border-gray-200\"><h1 class=\"sm:text-3xl text-2xl font-medium title-font mb-2 text-gray-900\">Pitchfork Kickstarter Taxidermy</h1><p class=\"leading-relaxed text-base\">Locavore cardigan small batch roof party blue bottle blog meggings sartorial jean shorts kickstarter migas sriracha church-key synth succulents. Actually taiyaki neutra, distillery gastropub pok pok ugh.</p><a class=\"text-indigo-500 inline-flex items-center mt-4\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div><div class=\"flex flex-col md:w-1/2 md:pl-12\"><h2 class=\"title-font font-semibold text-gray-800 tracking-wider text-sm mb-3\">CATEGORIES</h2><nav class=\"flex flex-wrap list-none -mb-1\"><li class=\"lg:w-1/3 mb-1 w-1/2\"><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li class=\"lg:w-1/3 mb-1 w-1/2\"><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li class=\"lg:w-1/3 mb-1 w-1/2\"><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li class=\"lg:w-1/3 mb-1 w-1/2\"><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li><li class=\"lg:w-1/3 mb-1 w-1/2\"><a class=\"text-gray-600 hover:text-gray-800\">Fifth Link</a></li><li class=\"lg:w-1/3 mb-1 w-1/2\"><a class=\"text-gray-600 hover:text-gray-800\">Sixth Link</a></li><li class=\"lg:w-1/3 mb-1 w-1/2\"><a class=\"text-gray-600 hover:text-gray-800\">Seventh Link</a></li><li class=\"lg:w-1/3 mb-1 w-1/2\"><a class=\"text-gray-600 hover:text-gray-800\">Eighth Link</a></li></nav></div></div></section>",category:'Content'},{id:'content-block-5',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M128 44.5a.5.5 0 011 0v62a.5.5 0 01-1 0v-62z" fill="#CBD5E0"></path><rect x="20" y="69" width="70" height="5" rx="2.5" fill="#4A5568"></rect><rect x="20" y="78" width="92" height="5" rx="2.5" fill="#4A5568"></rect><path d="M144 60a2 2 0 012-2h75a2 2 0 110 4h-75a2 2 0 01-2-2zM144 68a2 2 0 012-2h88a2 2 0 110 4h-88a2 2 0 01-2-2zM144 76a2 2 0 012-2h60a2 2 0 110 4h-60a2 2 0 01-2-2z" fill="#A0AEC0"></path><path d="M190 89a2 2 0 012-2h20a2 2 0 110 4h-20a2 2 0 01-2-2z" fill="#6366F1"></path><rect x="144" y="84" width="40" height="10" rx="2" fill="#6366F1"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto flex flex-wrap\"><h2 class=\"sm:text-3xl text-2xl text-gray-900 font-medium title-font mb-2 md:w-2/5\">Kickstarter Actually Pinterest Brunch Bitters Occupy</h2><div class=\"md:w-3/5 md:pl-6\"><p class=\"leading-relaxed text-base\">Taxidermy bushwick celiac master cleanse microdosing seitan. Fashion axe four dollar toast truffaut, direct trade kombucha brunch williamsburg keffiyeh gastropub tousled squid meh taiyaki drinking vinegar tacos.</p><div class=\"flex md:mt-4 mt-6\"><button class=\"inline-flex text-white bg-indigo-500 border-0 py-1 px-4 focus:outline-none hover:bg-indigo-600 rounded\">Button</button><a class=\"text-indigo-500 inline-flex items-center ml-4\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div></div></section>",category:'Content'},{id:'content-block-6',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M93 67.5a.5.5 0 011 0v62a.5.5 0 01-1 0v-62z" fill="#CBD5E0"></path><path d="M106 83a2 2 0 012-2h103.337a2 2 0 110 4H108a2 2 0 01-2-2zM106 107a2 2 0 012-2h95a2 2 0 110 4h-95a2 2 0 01-2-2zM106 91a2 2 0 012-2h121a2 2 0 110 4H108a2 2 0 01-2-2zM106 99a2 2 0 012-2h82.957a2 2 0 010 4H108a2 2 0 01-2-2z" fill="#A0AEC0"></path><path d="M106 115a2 2 0 012-2h20a2 2 0 110 4h-20a2 2 0 01-2-2z" fill="#6366F1"></path><path d="M45 104a2 2 0 012-2h20a2 2 0 110 4H47a2 2 0 01-2-2z" fill="#4A5568"></path><rect x="37" y="110" width="40" height="2" rx="1" fill="#A0AEC0"></rect><rect x="35" y="120" width="44" height="2" rx="1" fill="#A0AEC0"></rect><path d="M33 116a1 1 0 011-1h45a1 1 0 010 2H34a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M122.889 47h20.222A2.889 2.889 0 00146 44.111V23.89a2.889 2.889 0 00-2.889-2.89h-20.222A2.889 2.889 0 00120 23.889V44.11a2.889 2.889 0 002.889 2.89zm0 0l15.889-15.889L146 38.333m-15.889-9.389a2.167 2.167 0 11-4.333 0 2.167 2.167 0 014.333 0z" stroke="#A0AEC0" stroke-width="2px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><circle cx="56.5" cy="85.5" r="10.5" fill="#E2E8F0"></circle></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto flex flex-col\"><div class=\"lg:w-4/6 mx-auto\"><div class=\"rounded-lg h-64 overflow-hidden\"><img alt=\"content\" class=\"object-cover object-center h-full w-full\" src=\"https://dummyimage.com/1200x500\"></div><div class=\"flex flex-col sm:flex-row mt-10\"><div class=\"sm:w-1/3 text-center sm:pr-8 sm:py-8\"><div class=\"w-20 h-20 rounded-full inline-flex items-center justify-center bg-gray-200 text-gray-400\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-10 h-10\" viewBox=\"0 0 24 24\"><path d=\"M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2\"></path><circle cx=\"12\" cy=\"7\" r=\"4\"></circle></svg></div><div class=\"flex flex-col items-center text-center justify-center\"><h2 class=\"font-medium title-font mt-4 text-gray-900 text-lg\">Phoebe Caulfield</h2><div class=\"w-12 h-1 bg-indigo-500 rounded mt-2 mb-4\"></div><p class=\"text-base\">Raclette knausgaard hella meggs normcore williamsburg enamel pin sartorial venmo tbh hot chicken gentrify portland.</p></div></div><div class=\"sm:w-2/3 sm:pl-8 sm:py-8 sm:border-l border-gray-200 sm:border-t-0 border-t mt-4 pt-4 sm:mt-0 text-center sm:text-left\"><p class=\"leading-relaxed text-lg mb-4\">Meggings portland fingerstache lyft, post-ironic fixie man bun banh mi umami everyday carry hexagon locavore direct trade art party. Locavore small batch listicle gastropub farm-to-table lumbersexual salvia messenger bag. Coloring book flannel truffaut craft beer drinking vinegar sartorial, disrupt fashion axe normcore meh butcher. Portland 90's scenester vexillologist forage post-ironic asymmetrical, chartreuse disrupt butcher paleo intelligentsia pabst before they sold out four loko. 3 wolf moon brooklyn.</p><a class=\"text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div></div></div></section>",category:'Content'},{id:'content-block-7',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M41.692 86a2 2 0 012-2H114.4a2 2 0 010 4H43.692a2 2 0 01-2-2z" fill="#A0AEC0"></path><rect x="59" y="104" width="40" height="10" rx="2" fill="#6366F1"></rect><path d="M35 94a2 2 0 012-2h83a2 2 0 110 4H37a2 2 0 01-2-2z" fill="#A0AEC0"></path><path d="M68.889 63H89.11A2.889 2.889 0 0092 60.111V39.89A2.889 2.889 0 0089.111 37H68.89A2.889 2.889 0 0066 39.889V60.11A2.889 2.889 0 0068.889 63zm0 0l15.889-15.889L92 54.333m-15.889-9.389a2.167 2.167 0 11-4.333 0 2.167 2.167 0 014.333 0z" stroke="#A0AEC0" stroke-width="2px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="56" y="73" width="46" height="5" rx="2.5" fill="#4A5568"></rect><path d="M150.692 86a2 2 0 012-2h70.707a2 2 0 010 4h-70.707a2 2 0 01-2-2z" fill="#A0AEC0"></path><rect x="168" y="104" width="40" height="10" rx="2" fill="#6366F1"></rect><path d="M144 94a2 2 0 012-2h83a2 2 0 110 4h-83a2 2 0 01-2-2z" fill="#A0AEC0"></path><path d="M177.889 63h20.222A2.889 2.889 0 00201 60.111V39.89a2.889 2.889 0 00-2.889-2.89h-20.222A2.889 2.889 0 00175 39.889V60.11a2.889 2.889 0 002.889 2.89zm0 0l15.889-15.889L201 54.333m-15.889-9.389a2.167 2.167 0 11-4.333 0 2.167 2.167 0 014.333 0z" stroke="#A0AEC0" stroke-width="2px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="165" y="73" width="46" height="5" rx="2.5" fill="#4A5568"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-wrap -mx-4 -mb-10 text-center\"><div class=\"sm:w-1/2 mb-10 px-4\"><div class=\"rounded-lg h-64 overflow-hidden\"><img alt=\"content\" class=\"object-cover object-center h-full w-full\" src=\"https://dummyimage.com/1201x501\"></div><h2 class=\"title-font text-2xl font-medium text-gray-900 mt-6 mb-3\">Buy YouTube Videos</h2><p class=\"leading-relaxed text-base\">Williamsburg occupy sustainable snackwave gochujang. Pinterest cornhole brunch, slow-carb neutra irony.</p><button class=\"flex mx-auto mt-6 text-white bg-indigo-500 border-0 py-2 px-5 focus:outline-none hover:bg-indigo-600 rounded\">Button</button></div><div class=\"sm:w-1/2 mb-10 px-4\"><div class=\"rounded-lg h-64 overflow-hidden\"><img alt=\"content\" class=\"object-cover object-center h-full w-full\" src=\"https://dummyimage.com/1202x502\"></div><h2 class=\"title-font text-2xl font-medium text-gray-900 mt-6 mb-3\">The Catalyzer</h2><p class=\"leading-relaxed text-base\">Williamsburg occupy sustainable snackwave gochujang. Pinterest cornhole brunch, slow-carb neutra irony.</p><button class=\"flex mx-auto mt-6 text-white bg-indigo-500 border-0 py-2 px-5 focus:outline-none hover:bg-indigo-600 rounded\">Button</button></div></div></div></section>",category:'Content'},{id:'content-block-8',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M20 111.5a1.5 1.5 0 011.5-1.5h64a1.5 1.5 0 010 3h-64a1.5 1.5 0 01-1.5-1.5zM20 118.5a1.5 1.5 0 011.5-1.5h51a1.5 1.5 0 010 3h-51a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M20 125.5a1.5 1.5 0 011.5-1.5h18a1.5 1.5 0 010 3h-18a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><path d="M45.444 87h17.112A2.444 2.444 0 0065 84.556V67.444A2.444 2.444 0 0062.556 65H45.444A2.444 2.444 0 0043 67.444v17.112A2.444 2.444 0 0045.444 87zm0 0L58.89 73.556l6.11 6.11m-13.444-7.944a1.833 1.833 0 11-3.667 0 1.833 1.833 0 013.667 0z" stroke="#A0AEC0" stroke-width="1.8px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="20" y="101" width="56" height="4" rx="2" fill="#4A5568"></rect><path d="M100 111.5a1.5 1.5 0 011.5-1.5h64a1.5 1.5 0 010 3h-64a1.5 1.5 0 01-1.5-1.5zM100 118.5a1.5 1.5 0 011.5-1.5h51a1.5 1.5 0 010 3h-51a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M100 125.5a1.5 1.5 0 011.5-1.5h18a1.5 1.5 0 010 3h-18a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><path d="M125.444 87h17.112A2.444 2.444 0 00145 84.556V67.444A2.444 2.444 0 00142.556 65h-17.112A2.444 2.444 0 00123 67.444v17.112A2.444 2.444 0 00125.444 87zm0 0l13.445-13.444 6.111 6.11m-13.444-7.944a1.834 1.834 0 11-3.667 0 1.834 1.834 0 013.667 0z" stroke="#A0AEC0" stroke-width="1.8px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="100" y="101" width="56" height="4" rx="2" fill="#4A5568"></rect><path d="M180 111.5a1.5 1.5 0 011.5-1.5h64a1.5 1.5 0 010 3h-64a1.5 1.5 0 01-1.5-1.5zM180 118.5a1.5 1.5 0 011.5-1.5h51a1.5 1.5 0 010 3h-51a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M180 125.5a1.5 1.5 0 011.5-1.5h18a1.5 1.5 0 010 3h-18a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><path d="M205.444 87h17.112A2.444 2.444 0 00225 84.556V67.444A2.444 2.444 0 00222.556 65h-17.112A2.444 2.444 0 00203 67.444v17.112A2.444 2.444 0 00205.444 87zm0 0l13.445-13.444 6.111 6.11m-13.444-7.944a1.834 1.834 0 11-3.667 0 1.834 1.834 0 013.667 0z" stroke="#A0AEC0" stroke-width="1.8px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="180" y="101" width="56" height="4" rx="2" fill="#4A5568"></rect><rect x="142" y="23" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="23" width="74" height="5" rx="2.5" fill="#4A5568"></rect><rect x="142" y="31" width="77" height="4" rx="2" fill="#A0AEC0"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-col\"><div class=\"h-1 bg-gray-200 rounded overflow-hidden\"><div class=\"w-24 h-full bg-indigo-500\"></div></div><div class=\"flex flex-wrap sm:flex-row flex-col py-6 mb-12\"><h1 class=\"sm:w-2/5 text-gray-900 font-medium title-font text-2xl mb-2 sm:mb-0\">Space The Final Frontier</h1><p class=\"sm:w-3/5 leading-relaxed text-base sm:pl-10 pl-0\">Street art subway tile salvia four dollar toast bitters selfies quinoa yuccie synth meditation iPhone intelligentsia prism tofu. Viral gochujang bitters dreamcatcher.</p></div></div><div class=\"flex flex-wrap sm:-m-4 -mx-4 -mb-10 -mt-4\"><div class=\"p-4 md:w-1/3 sm:mb-0 mb-6\"><div class=\"rounded-lg h-64 overflow-hidden\"><img alt=\"content\" class=\"object-cover object-center h-full w-full\" src=\"https://dummyimage.com/1203x503\"></div><h2 class=\"text-xl font-medium title-font text-gray-900 mt-5\">Shooting Stars</h2><p class=\"text-base leading-relaxed mt-2\">Swag shoivdigoitch literally meditation subway tile tumblr cold-pressed. Gastropub street art beard dreamcatcher neutra, ethical XOXO lumbersexual.</p><a class=\"text-indigo-500 inline-flex items-center mt-3\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div><div class=\"p-4 md:w-1/3 sm:mb-0 mb-6\"><div class=\"rounded-lg h-64 overflow-hidden\"><img alt=\"content\" class=\"object-cover object-center h-full w-full\" src=\"https://dummyimage.com/1204x504\"></div><h2 class=\"text-xl font-medium title-font text-gray-900 mt-5\">The Catalyzer</h2><p class=\"text-base leading-relaxed mt-2\">Swag shoivdigoitch literally meditation subway tile tumblr cold-pressed. Gastropub street art beard dreamcatcher neutra, ethical XOXO lumbersexual.</p><a class=\"text-indigo-500 inline-flex items-center mt-3\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div><div class=\"p-4 md:w-1/3 sm:mb-0 mb-6\"><div class=\"rounded-lg h-64 overflow-hidden\"><img alt=\"content\" class=\"object-cover object-center h-full w-full\" src=\"https://dummyimage.com/1205x505\"></div><h2 class=\"text-xl font-medium title-font text-gray-900 mt-5\">The 400 Blows</h2><p class=\"text-base leading-relaxed mt-2\">Swag shoivdigoitch literally meditation subway tile tumblr cold-pressed. Gastropub street art beard dreamcatcher neutra, ethical XOXO lumbersexual.</p><a class=\"text-indigo-500 inline-flex items-center mt-3\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div></div></section>",category:'Content'},{id:'cta-block-1',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="20" y="68" width="119" height="5" rx="2.5" fill="#4A5568"></rect><rect x="20" y="77" width="92" height="5" rx="2.5" fill="#4A5568"></rect><rect x="206" y="70" width="40" height="10" rx="2" fill="#6366F1"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"lg:w-2/3 flex flex-col sm:flex-row sm:items-center items-start mx-auto\"><h1 class=\"flex-grow sm:pr-16 text-2xl font-medium title-font text-gray-900\">Slow-carb next level shoindxgoitch ethical authentic, scenester sriracha forage.</h1><button class=\"flex-shrink-0 text-white bg-indigo-500 border-0 py-2 px-8 focus:outline-none hover:bg-indigo-600 rounded text-lg mt-10 sm:mt-0\">Button</button></div></div></section>",category:'CTA'},{id:'cta-block-2',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="151" y="40" width="93" height="70" rx="2" fill="#E2E8F0"></rect><rect x="20" y="61" width="86" height="5" rx="2.5" fill="#4A5568"></rect><rect x="20" y="70" width="66" height="5" rx="2.5" fill="#4A5568"></rect><rect x="20" y="79" width="106" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="87" width="97" height="4" rx="2" fill="#A0AEC0"></rect><rect x="160" y="91" width="75" height="10" rx="2" fill="#6366F1"></rect><rect x="160" y="76" width="75" height="10" rx="2" fill="#CBD5E0"></rect><rect x="160" y="61" width="75" height="10" rx="2" fill="#CBD5E0"></rect><rect x="160" y="50" width="40" height="4" rx="2" fill="#4A5568"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto flex flex-wrap items-center\"><div class=\"lg:w-3/5 md:w-1/2 md:pr-16 lg:pr-0 pr-0\"><h1 class=\"title-font font-medium text-3xl text-gray-900\">Slow-carb next level shoindcgoitch ethical authentic, poko scenester</h1><p class=\"leading-relaxed mt-4\">Poke slow-carb mixtape knausgaard, typewriter street art gentrify hammock starladder roathse. Craies vegan tousled etsy austin.</p></div><div class=\"lg:w-2/6 md:w-1/2 bg-gray-100 rounded-lg p-8 flex flex-col md:ml-auto w-full mt-10 md:mt-0\"><form style=\"margin: 0;\"><h2 class=\"text-gray-900 text-lg font-medium title-font mb-5\">Sign Up</h2><div class=\"relative mb-4\"><label for=\"full-name\" class=\"leading-7 text-sm text-gray-600\">Full Name</label><input type=\"text\" id=\"full-name\" name=\"full-name\" class=\"w-full bg-white rounded border border-gray-300 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div><div class=\"relative mb-4\"><label for=\"email\" class=\"leading-7 text-sm text-gray-600\">Email</label><input type=\"email\" id=\"email\" name=\"email\" class=\"w-full bg-white rounded border border-gray-300 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div><button type=\"submit\" class=\"text-white bg-indigo-500 border-0 py-2 px-8 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button><p class=\"text-xs text-gray-500 mt-3\">Literally you probably haven't heard of them jean shorts.</p></form></div></div></section>",category:'CTA'},{id:'cta-block-3',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="90" y="48" width="86" height="5" rx="2.5" fill="#4A5568"></rect><rect x="80" y="57" width="106" height="4" rx="2" fill="#A0AEC0"></rect><rect x="85" y="65" width="97" height="4" rx="2" fill="#A0AEC0"></rect><rect x="183" y="92" width="44" height="10" rx="2" fill="#6366F1"></rect><rect x="111" y="92" width="66" height="10" rx="2" fill="#CBD5E0"></rect><rect x="39" y="92" width="66" height="10" rx="2" fill="#CBD5E0"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-col text-center w-full mb-12\"><h1 class=\"sm:text-3xl text-2xl font-medium title-font mb-4 text-gray-900\">Master Cleanse Reliac Heirloom</h1><p class=\"lg:w-2/3 mx-auto leading-relaxed text-base\">Whatever cardigan tote bag tumblr hexagon brooklyn asymmetrical gentrify, subway tile poke farm-to-table. Franzen you probably haven't heard of them man bun deep.</p></div><form style=\"margin: 0;\"><div class=\"flex lg:w-2/3 w-full sm:flex-row flex-col mx-auto px-8 sm:space-x-4 sm:space-y-0 space-y-4 sm:px-0 items-end\"><div class=\"relative flex-grow w-full\"><label for=\"full-name\" class=\"leading-7 text-sm text-gray-600\">Full Name</label><input type=\"text\" id=\"full-name\" name=\"full-name\" class=\"w-full bg-gray-100 bg-opacity-50 rounded border border-gray-300 focus:border-indigo-500 focus:bg-transparent focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div><div class=\"relative flex-grow w-full\"><label for=\"email\" class=\"leading-7 text-sm text-gray-600\">Email</label><input type=\"email\" id=\"email\" name=\"email\" class=\"w-full bg-gray-100 bg-opacity-50 rounded border border-gray-300 focus:border-indigo-500 focus:bg-transparent focus:ring-2 focus:ring-indigo-200 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div><button type=\"submit\" class=\"text-white bg-indigo-500 border-0 py-2 px-8 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button></div></form></div></section>",category:'CTA'},{id:'cta-block-4',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="20" y="68" width="26" height="5" rx="2.5" fill="#6366F1"></rect><rect x="20" y="77" width="92" height="5" rx="2.5" fill="#4A5568"></rect><rect x="206" y="70" width="40" height="10" rx="2" fill="#CBD5E0"></rect><rect x="160" y="70" width="40" height="10" rx="2" fill="#CBD5E0"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto flex items-center md:flex-row flex-col\"><div class=\"flex flex-col md:pr-10 md:mb-0 mb-6 pr-0 w-full md:w-auto md:text-left text-center\"><h2 class=\"text-xs text-indigo-500 tracking-widest font-medium title-font mb-1\">ROOF PARTY POLAROID</h2><h1 class=\"md:text-3xl text-2xl font-medium title-font text-gray-900\">Master Cleanse Reliac Heirloom</h1></div><div class=\"flex md:ml-auto md:mr-0 mx-auto items-center flex-shrink-0 space-x-4\"><button class=\"bg-gray-100 inline-flex py-3 px-5 rounded-lg items-center hover:bg-gray-200 focus:outline-none\"><svg fill=\"currentColor\" class=\"w-6 h-6\" viewBox=\"0 0 512 512\"><path d=\"M99.617 8.057a50.191 50.191 0 00-38.815-6.713l230.932 230.933 74.846-74.846L99.617 8.057zM32.139 20.116c-6.441 8.563-10.148 19.077-10.148 30.199v411.358c0 11.123 3.708 21.636 10.148 30.199l235.877-235.877L32.139 20.116zM464.261 212.087l-67.266-37.637-81.544 81.544 81.548 81.548 67.273-37.64c16.117-9.03 25.738-25.442 25.738-43.908s-9.621-34.877-25.749-43.907zM291.733 279.711L60.815 510.629c3.786.891 7.639 1.371 11.492 1.371a50.275 50.275 0 0027.31-8.07l266.965-149.372-74.849-74.847z\"></path></svg><span class=\"ml-4 flex items-start flex-col leading-none\"><span class=\"text-xs text-gray-600 mb-1\">GET IT ON</span><span class=\"title-font font-medium\">Google Play</span></span></button><button class=\"bg-gray-100 inline-flex py-3 px-5 rounded-lg items-center hover:bg-gray-200 focus:outline-none\"><svg fill=\"currentColor\" class=\"w-6 h-6\" viewBox=\"0 0 305 305\"><path d=\"M40.74 112.12c-25.79 44.74-9.4 112.65 19.12 153.82C74.09 286.52 88.5 305 108.24 305c.37 0 .74 0 1.13-.02 9.27-.37 15.97-3.23 22.45-5.99 7.27-3.1 14.8-6.3 26.6-6.3 11.22 0 18.39 3.1 25.31 6.1 6.83 2.95 13.87 6 24.26 5.81 22.23-.41 35.88-20.35 47.92-37.94a168.18 168.18 0 0021-43l.09-.28a2.5 2.5 0 00-1.33-3.06l-.18-.08c-3.92-1.6-38.26-16.84-38.62-58.36-.34-33.74 25.76-51.6 31-54.84l.24-.15a2.5 2.5 0 00.7-3.51c-18-26.37-45.62-30.34-56.73-30.82a50.04 50.04 0 00-4.95-.24c-13.06 0-25.56 4.93-35.61 8.9-6.94 2.73-12.93 5.09-17.06 5.09-4.64 0-10.67-2.4-17.65-5.16-9.33-3.7-19.9-7.9-31.1-7.9l-.79.01c-26.03.38-50.62 15.27-64.18 38.86z\"></path><path d=\"M212.1 0c-15.76.64-34.67 10.35-45.97 23.58-9.6 11.13-19 29.68-16.52 48.38a2.5 2.5 0 002.29 2.17c1.06.08 2.15.12 3.23.12 15.41 0 32.04-8.52 43.4-22.25 11.94-14.5 17.99-33.1 16.16-49.77A2.52 2.52 0 00212.1 0z\"></path></svg><span class=\"ml-4 flex items-start flex-col leading-none\"><span class=\"text-xs text-gray-600 mb-1\">Download on the</span><span class=\"title-font font-medium\">App Store</span></span></button></div></div></section>",category:'CTA'},{id:'commerce-block-1',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M40.333 48h9.334c.736 0 1.333-.597 1.333-1.333v-9.334c0-.736-.597-1.333-1.333-1.333h-9.334c-.736 0-1.333.597-1.333 1.333v9.334c0 .736.597 1.333 1.333 1.333zm0 0l7.334-7.333L51 44m-7.333-4.333a1 1 0 11-2 0 1 1 0 012 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="20" y="61" width="41" height="2" rx="1" fill="#4A5568"></rect><rect x="20" y="65" width="13" height="2" rx="1" fill="#A0AEC0"></rect><rect x="20" y="57" width="17" height="2" rx="1" fill="#CBD5E0"></rect><path d="M99.333 48h9.334c.736 0 1.333-.597 1.333-1.333v-9.334c0-.736-.597-1.333-1.333-1.333h-9.334c-.736 0-1.333.597-1.333 1.333v9.334c0 .736.597 1.333 1.333 1.333zm0 0l7.334-7.333L110 44m-7.333-4.333a1 1 0 11-2 0 1 1 0 012 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="79" y="61" width="41" height="2" rx="1" fill="#4A5568"></rect><rect x="79" y="65" width="13" height="2" rx="1" fill="#A0AEC0"></rect><rect x="79" y="57" width="17" height="2" rx="1" fill="#CBD5E0"></rect><path d="M158.333 48h9.334c.736 0 1.333-.597 1.333-1.333v-9.334c0-.736-.597-1.333-1.333-1.333h-9.334c-.736 0-1.333.597-1.333 1.333v9.334c0 .736.597 1.333 1.333 1.333zm0 0l7.334-7.333L169 44m-7.333-4.333a1 1 0 11-2 0 1 1 0 012 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="138" y="61" width="41" height="2" rx="1" fill="#4A5568"></rect><rect x="138" y="65" width="13" height="2" rx="1" fill="#A0AEC0"></rect><rect x="138" y="57" width="17" height="2" rx="1" fill="#CBD5E0"></rect><path d="M217.333 48h9.334c.736 0 1.333-.597 1.333-1.333v-9.334c0-.736-.597-1.333-1.333-1.333h-9.334c-.736 0-1.333.597-1.333 1.333v9.334c0 .736.597 1.333 1.333 1.333zm0 0l7.334-7.333L228 44m-7.333-4.333a1 1 0 11-2 0 1 1 0 012 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="197" y="61" width="41" height="2" rx="1" fill="#4A5568"></rect><rect x="197" y="65" width="13" height="2" rx="1" fill="#A0AEC0"></rect><rect x="197" y="57" width="17" height="2" rx="1" fill="#CBD5E0"></rect><path d="M40.333 94h9.334c.736 0 1.333-.597 1.333-1.333v-9.334c0-.736-.597-1.333-1.333-1.333h-9.334c-.736 0-1.333.597-1.333 1.333v9.334c0 .736.597 1.333 1.333 1.333zm0 0l7.334-7.333L51 90m-7.333-4.333a1 1 0 11-2 0 1 1 0 012 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="20" y="107" width="41" height="2" rx="1" fill="#4A5568"></rect><rect x="20" y="111" width="13" height="2" rx="1" fill="#A0AEC0"></rect><rect x="20" y="103" width="17" height="2" rx="1" fill="#CBD5E0"></rect><path d="M99.333 94h9.334c.736 0 1.333-.597 1.333-1.333v-9.334c0-.736-.597-1.333-1.333-1.333h-9.334c-.736 0-1.333.597-1.333 1.333v9.334c0 .736.597 1.333 1.333 1.333zm0 0l7.334-7.333L110 90m-7.333-4.333a1 1 0 11-2 0 1 1 0 012 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="79" y="107" width="41" height="2" rx="1" fill="#4A5568"></rect><rect x="79" y="111" width="13" height="2" rx="1" fill="#A0AEC0"></rect><rect x="79" y="103" width="17" height="2" rx="1" fill="#CBD5E0"></rect><path d="M158.333 94h9.334c.736 0 1.333-.597 1.333-1.333v-9.334c0-.736-.597-1.333-1.333-1.333h-9.334c-.736 0-1.333.597-1.333 1.333v9.334c0 .736.597 1.333 1.333 1.333zm0 0l7.334-7.333L169 90m-7.333-4.333a1 1 0 11-2 0 1 1 0 012 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="138" y="107" width="41" height="2" rx="1" fill="#4A5568"></rect><rect x="138" y="111" width="13" height="2" rx="1" fill="#A0AEC0"></rect><rect x="138" y="103" width="17" height="2" rx="1" fill="#CBD5E0"></rect><path d="M217.333 94h9.334c.736 0 1.333-.597 1.333-1.333v-9.334c0-.736-.597-1.333-1.333-1.333h-9.334c-.736 0-1.333.597-1.333 1.333v9.334c0 .736.597 1.333 1.333 1.333zm0 0l7.334-7.333L228 90m-7.333-4.333a1 1 0 11-2 0 1 1 0 012 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="197" y="107" width="41" height="2" rx="1" fill="#4A5568"></rect><rect x="197" y="111" width="13" height="2" rx="1" fill="#A0AEC0"></rect><rect x="197" y="103" width="17" height="2" rx="1" fill="#CBD5E0"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-wrap -m-4\"><div class=\"lg:w-1/4 md:w-1/2 p-4 w-full\"><a class=\"block relative h-48 rounded overflow-hidden\"><img alt=\"ecommerce\" class=\"object-cover object-center w-full h-full block\" src=\"https://dummyimage.com/420x260\"></a><div class=\"mt-4\"><h3 class=\"text-gray-500 text-xs tracking-widest title-font mb-1\">CATEGORY</h3><h2 class=\"text-gray-900 title-font text-lg font-medium\">The Catalyzer</h2><p class=\"mt-1\">$16.00</p></div></div><div class=\"lg:w-1/4 md:w-1/2 p-4 w-full\"><a class=\"block relative h-48 rounded overflow-hidden\"><img alt=\"ecommerce\" class=\"object-cover object-center w-full h-full block\" src=\"https://dummyimage.com/421x261\"></a><div class=\"mt-4\"><h3 class=\"text-gray-500 text-xs tracking-widest title-font mb-1\">CATEGORY</h3><h2 class=\"text-gray-900 title-font text-lg font-medium\">Shooting Stars</h2><p class=\"mt-1\">$21.15</p></div></div><div class=\"lg:w-1/4 md:w-1/2 p-4 w-full\"><a class=\"block relative h-48 rounded overflow-hidden\"><img alt=\"ecommerce\" class=\"object-cover object-center w-full h-full block\" src=\"https://dummyimage.com/422x262\"></a><div class=\"mt-4\"><h3 class=\"text-gray-500 text-xs tracking-widest title-font mb-1\">CATEGORY</h3><h2 class=\"text-gray-900 title-font text-lg font-medium\">Neptune</h2><p class=\"mt-1\">$12.00</p></div></div><div class=\"lg:w-1/4 md:w-1/2 p-4 w-full\"><a class=\"block relative h-48 rounded overflow-hidden\"><img alt=\"ecommerce\" class=\"object-cover object-center w-full h-full block\" src=\"https://dummyimage.com/423x263\"></a><div class=\"mt-4\"><h3 class=\"text-gray-500 text-xs tracking-widest title-font mb-1\">CATEGORY</h3><h2 class=\"text-gray-900 title-font text-lg font-medium\">The 400 Blows</h2><p class=\"mt-1\">$18.40</p></div></div><div class=\"lg:w-1/4 md:w-1/2 p-4 w-full\"><a class=\"block relative h-48 rounded overflow-hidden\"><img alt=\"ecommerce\" class=\"object-cover object-center w-full h-full block\" src=\"https://dummyimage.com/424x264\"></a><div class=\"mt-4\"><h3 class=\"text-gray-500 text-xs tracking-widest title-font mb-1\">CATEGORY</h3><h2 class=\"text-gray-900 title-font text-lg font-medium\">The Catalyzer</h2><p class=\"mt-1\">$16.00</p></div></div><div class=\"lg:w-1/4 md:w-1/2 p-4 w-full\"><a class=\"block relative h-48 rounded overflow-hidden\"><img alt=\"ecommerce\" class=\"object-cover object-center w-full h-full block\" src=\"https://dummyimage.com/425x265\"></a><div class=\"mt-4\"><h3 class=\"text-gray-500 text-xs tracking-widest title-font mb-1\">CATEGORY</h3><h2 class=\"text-gray-900 title-font text-lg font-medium\">Shooting Stars</h2><p class=\"mt-1\">$21.15</p></div></div><div class=\"lg:w-1/4 md:w-1/2 p-4 w-full\"><a class=\"block relative h-48 rounded overflow-hidden\"><img alt=\"ecommerce\" class=\"object-cover object-center w-full h-full block\" src=\"https://dummyimage.com/427x267\"></a><div class=\"mt-4\"><h3 class=\"text-gray-500 text-xs tracking-widest title-font mb-1\">CATEGORY</h3><h2 class=\"text-gray-900 title-font text-lg font-medium\">Neptune</h2><p class=\"mt-1\">$12.00</p></div></div><div class=\"lg:w-1/4 md:w-1/2 p-4 w-full\"><a class=\"block relative h-48 rounded overflow-hidden\"><img alt=\"ecommerce\" class=\"object-cover object-center w-full h-full block\" src=\"https://dummyimage.com/428x268\"></a><div class=\"mt-4\"><h3 class=\"text-gray-500 text-xs tracking-widest title-font mb-1\">CATEGORY</h3><h2 class=\"text-gray-900 title-font text-lg font-medium\">The 400 Blows</h2><p class=\"mt-1\">$18.40</p></div></div></div></div></section>",category:'Commerce'},{id:'commerce-block-2',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M52.792 91h26.544a3.785 3.785 0 003.792-3.778V60.778A3.785 3.785 0 0079.336 57H52.792A3.785 3.785 0 0049 60.778v26.444A3.785 3.785 0 0052.792 91zm0 0l20.856-20.778 9.48 9.445M62.272 67.389a2.839 2.839 0 01-2.844 2.833 2.839 2.839 0 01-2.844-2.833 2.839 2.839 0 012.844-2.833 2.839 2.839 0 012.844 2.833z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="133" y="42" width="74" height="5" rx="2.5" fill="#4A5568"></rect><rect x="133" y="105" width="26" height="5" rx="2.5" fill="#4A5568"></rect><rect x="133" y="35" width="38" height="3" rx="1.5" fill="#CBD5E0"></rect><rect x="133" y="51" width="26" height="3" rx="1.5" fill="#6366F1"></rect><rect x="133" y="64" width="92" height="4" rx="2" fill="#A0AEC0"></rect><circle cx="241" cy="108" r="5" fill="#E2E8F0"></circle><path d="M245.5 94a.5.5 0 010 1h-112a.5.5 0 010-1h112z" fill="#E2E8F0"></path><rect x="200" y="103" width="31" height="10" rx="2" fill="#6366F1"></rect><rect x="133" y="72" width="82" height="4" rx="2" fill="#A0AEC0"></rect><rect x="133" y="80" width="68" height="4" rx="2" fill="#A0AEC0"></rect></svg>',content:"<section class=\"text-gray-600 body-font overflow-hidden\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"lg:w-4/5 mx-auto flex flex-wrap\"><img alt=\"ecommerce\" class=\"lg:w-1/2 w-full lg:h-auto h-64 object-cover object-center rounded\" src=\"https://dummyimage.com/400x400\"><div class=\"lg:w-1/2 w-full lg:pl-10 lg:py-6 mt-6 lg:mt-0\"><form style=\"margin: 0;\"><h2 class=\"text-sm title-font text-gray-500 tracking-widest\">BRAND NAME</h2><h1 class=\"text-gray-900 text-3xl title-font font-medium mb-1\">The Catcher in the Rye</h1><div class=\"flex mb-4\"><span class=\"flex items-center\"><svg fill=\"currentColor\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 text-indigo-500\" viewBox=\"0 0 24 24\"><path d=\"M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z\"></path></svg><svg fill=\"currentColor\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 text-indigo-500\" viewBox=\"0 0 24 24\"><path d=\"M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z\"></path></svg><svg fill=\"currentColor\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 text-indigo-500\" viewBox=\"0 0 24 24\"><path d=\"M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z\"></path></svg><svg fill=\"currentColor\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 text-indigo-500\" viewBox=\"0 0 24 24\"><path d=\"M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z\"></path></svg><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 text-indigo-500\" viewBox=\"0 0 24 24\"><path d=\"M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z\"></path></svg><span class=\"text-gray-600 ml-3\">4 Reviews</span></span><span class=\"flex ml-3 pl-3 py-2 border-l-2 border-gray-200 space-x-2s\"><a class=\"text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg></a></span></div><p class=\"leading-relaxed\">Fam locavore kickstarter distillery. Mixtape chillwave tumeric sriracha taximy chia microdosing tilde DIY. XOXO fam indxgo juiceramps cornhole raw denim forage brooklyn. Everyday carry +1 seitan poutine tumeric. Gastropub blue bottle austin listicle pour-over, neutra jean shorts keytar banjo tattooed umami cardigan.</p><div class=\"flex mt-6 items-center pb-5 border-b-2 border-gray-100 mb-5\"><div class=\"flex items-center\"><span class=\"mr-3\">Color</span><div class=\"relative\"><select class=\"rounded border appearance-none border-gray-300 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-200 focus:border-indigo-500 text-base pl-3 pr-10\" required><option>⚫️</option><option>⚪️</option><option>🔵</option><option>🟣</option></select><span class=\"absolute right-0 top-0 h-full w-10 text-center text-gray-600 pointer-events-none flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4\" viewBox=\"0 0 24 24\"><path d=\"M6 9l6 6 6-6\"></path></svg></span></div></div><div class=\"flex ml-6 items-center\"><span class=\"mr-3\">Size</span><div class=\"relative\"><select class=\"rounded border appearance-none border-gray-300 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-200 focus:border-indigo-500 text-base pl-3 pr-10\" required><option>SM</option><option>M</option><option>L</option><option>XL</option></select><span class=\"absolute right-0 top-0 h-full w-10 text-center text-gray-600 pointer-events-none flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4\" viewBox=\"0 0 24 24\"><path d=\"M6 9l6 6 6-6\"></path></svg></span></div></div></div><div class=\"flex\"><span class=\"title-font font-medium text-2xl text-gray-900\">$58.00</span><button type=\"submit\" class=\"flex ml-auto text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded\">Button</button><button class=\"rounded-full w-10 h-10 bg-gray-200 p-0 border-0 inline-flex items-center justify-center text-gray-500 ml-4\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M20.84 4.61a5.5 5.5 0 00-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 00-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 000-7.78z\"></path></svg></button></div></form></div></div></div></section>",category:'Commerce'},{id:'commerce-block-3',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="20" y="37" width="74" height="5" rx="2.5" fill="#4A5568"></rect><rect x="20" y="112" width="26" height="5" rx="2.5" fill="#4A5568"></rect><rect x="20" y="30" width="38" height="3" rx="1.5" fill="#CBD5E0"></rect><rect x="20" y="55" width="26" height="3" rx="1.5" fill="#6366F1"></rect><path d="M56 56.5a1.5 1.5 0 011.5-1.5h23a1.5 1.5 0 010 3h-23a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><rect x="92" y="55" width="26" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="20" y="73" width="92" height="4" rx="2" fill="#A0AEC0"></rect><path d="M132.5 102a.5.5 0 010 1h-112a.5.5 0 010-1h112zM133 64v2H20v-2h113z" fill="#E2E8F0"></path><rect x="87" y="110" width="31" height="10" rx="2" fill="#6366F1"></rect><rect x="20" y="81" width="82" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="89" width="68" height="4" rx="2" fill="#A0AEC0"></rect><path fill="#6366F1" d="M20 64h32v2H20z"></path><path d="M187.792 92h26.544a3.785 3.785 0 003.792-3.778V61.778A3.785 3.785 0 00214.336 58h-26.544A3.785 3.785 0 00184 61.778v26.444A3.785 3.785 0 00187.792 92zm0 0l20.856-20.778 9.48 9.445m-20.856-12.278a2.838 2.838 0 01-2.844 2.833 2.838 2.838 0 01-2.844-2.833 2.838 2.838 0 012.844-2.833 2.838 2.838 0 012.844 2.833z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><circle cx="128" cy="114" r="5" fill="#E2E8F0"></circle></svg>',content:"<section class=\"text-gray-600 body-font overflow-hidden\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"lg:w-4/5 mx-auto flex flex-wrap\"><div class=\"lg:w-1/2 w-full lg:pr-10 lg:py-6 mb-6 lg:mb-0\"><form style=\"margin: 0;\"><h2 class=\"text-sm title-font text-gray-500 tracking-widest\">BRAND NAME</h2><h1 class=\"text-gray-900 text-3xl title-font font-medium mb-4\">Animated Night Hill Illustrations</h1><p class=\"leading-relaxed mb-4\">Fam locavore kickstarter distillery. Mixtape chillwave tumeric sriracha taximy chia microdosing tilde DIY. XOXO fam inxigo juiceramps cornhole raw denim forage brooklyn. Everyday carry +1 seitan poutine tumeric. Gastropub blue bottle austin listicle pour-over, neutra jean.</p><div class=\"flex border-t border-gray-200 py-2\"><span class=\"inline-flex items-center text-gray-500\">Color</span><div class=\"ml-auto relative\"><select class=\"rounded border appearance-none border-gray-300 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-200 focus:border-indigo-500 text-base pl-3 pr-10\" required><option>⚫️</option><option>⚪️</option><option>🔵</option><option>🟣</option></select><span class=\"absolute right-0 top-0 h-full w-10 text-center text-gray-600 pointer-events-none flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4\" viewBox=\"0 0 24 24\"><path d=\"M6 9l6 6 6-6\"></path></svg></span></div></div><div class=\"flex border-t border-gray-200 py-2\"><span class=\"inline-flex items-center text-gray-500\">Size</span><div class=\"ml-auto relative\"><select class=\"rounded border appearance-none border-gray-300 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-200 focus:border-indigo-500 text-base pl-3 pr-10\" required><option>Small</option><option>Medium</option><option>Large</option><option>X-Large</option></select><span class=\"absolute right-0 top-0 h-full w-10 text-center text-gray-600 pointer-events-none flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4\" viewBox=\"0 0 24 24\"><path d=\"M6 9l6 6 6-6\"></path></svg></span></div></div><div class=\"flex border-t border-b mb-6 border-gray-200 py-2\"><span class=\"inline-flex items-center text-gray-500\">Quantity</span><div class=\"ml-auto relative\"><select class=\"rounded border appearance-none border-gray-300 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-200 focus:border-indigo-500 text-base pl-3 pr-10\" required><option>1</option><option>2</option><option>3</option><option>4</option><option>5</option></select><span class=\"absolute right-0 top-0 h-full w-10 text-center text-gray-600 pointer-events-none flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4\" viewBox=\"0 0 24 24\"><path d=\"M6 9l6 6 6-6\"></path></svg></span></div></div><div class=\"flex\"><span class=\"title-font font-medium text-2xl text-gray-900\">$58.00</span><button class=\"flex ml-auto text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded\">Button</button><button class=\"rounded-full w-10 h-10 bg-gray-200 p-0 border-0 inline-flex items-center justify-center text-gray-500 ml-4\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M20.84 4.61a5.5 5.5 0 00-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 00-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 000-7.78z\"></path></svg></button></div></form></div><img alt=\"ecommerce\" class=\"lg:w-1/2 w-full lg:h-auto h-64 object-cover object-center rounded\" src=\"https://dummyimage.com/400x400\"></div></div></section>",category:'Commerce'},{id:'feature-block-1',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="96" y="39" width="74" height="5" rx="2.5" fill="#4A5568"></rect><circle cx="26" cy="90" r="6" fill="#C3DAFE"></circle><rect x="82" y="48" width="102" height="5" rx="2.5" fill="#4A5568"></rect><path d="M38 94.5a1.5 1.5 0 011.5-1.5h48a1.5 1.5 0 010 3h-48a1.5 1.5 0 01-1.5-1.5zM38 101.5a1.5 1.5 0 011.5-1.5h38a1.5 1.5 0 010 3h-38a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M38 108.5a1.5 1.5 0 011.5-1.5h13a1.5 1.5 0 010 3h-13a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><rect x="38" y="84" width="43" height="4" rx="2" fill="#4A5568"></rect><circle cx="105" cy="90" r="6" fill="#C3DAFE"></circle><path d="M117 94.5a1.5 1.5 0 011.5-1.5h48a1.5 1.5 0 010 3h-48a1.5 1.5 0 01-1.5-1.5zM117 101.5a1.5 1.5 0 011.5-1.5h38a1.5 1.5 0 010 3h-38a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M117 108.5a1.5 1.5 0 011.5-1.5h13a1.5 1.5 0 010 3h-13a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><rect x="117" y="84" width="43" height="4" rx="2" fill="#4A5568"></rect><circle cx="184" cy="90" r="6" fill="#C3DAFE"></circle><path d="M196 94.5a1.5 1.5 0 011.5-1.5h48a1.5 1.5 0 010 3h-48a1.5 1.5 0 01-1.5-1.5zM196 101.5a1.5 1.5 0 011.5-1.5h38a1.5 1.5 0 010 3h-38a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M196 108.5a1.5 1.5 0 011.5-1.5h13a1.5 1.5 0 010 3h-13a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><rect x="196" y="84" width="43" height="4" rx="2" fill="#4A5568"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><h1 class=\"sm:text-3xl text-2xl font-medium title-font text-center text-gray-900 mb-20\">Raw Denim Heirloom Man Braid<br class=\"hidden sm:block\">Selfies Wayfarers</h1><div class=\"flex flex-wrap sm:-m-4 -mx-4 -mb-10 -mt-4 md:space-y-0 space-y-6\"><div class=\"p-4 md:w-1/3 flex\"><div class=\"w-12 h-12 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-4 flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-6 h-6\" viewBox=\"0 0 24 24\"><path d=\"M22 12h-4l-3 9L9 3l-3 9H2\"></path></svg></div><div class=\"flex-grow pl-6\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-2\">Shooting Stars</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine, ramps microdosing banh mi pug VHS try-hard ugh iceland kickstarter tumblr live-edge tilde.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div><div class=\"p-4 md:w-1/3 flex\"><div class=\"w-12 h-12 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-4 flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-6 h-6\" viewBox=\"0 0 24 24\"><circle cx=\"6\" cy=\"6\" r=\"3\"></circle><circle cx=\"6\" cy=\"18\" r=\"3\"></circle><path d=\"M20 4L8.12 15.88M14.47 14.48L20 20M8.12 8.12L12 12\"></path></svg></div><div class=\"flex-grow pl-6\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-2\">The Catalyzer</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine, ramps microdosing banh mi pug VHS try-hard ugh iceland kickstarter tumblr live-edge tilde.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div><div class=\"p-4 md:w-1/3 flex\"><div class=\"w-12 h-12 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-4 flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-6 h-6\" viewBox=\"0 0 24 24\"><path d=\"M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2\"></path><circle cx=\"12\" cy=\"7\" r=\"4\"></circle></svg></div><div class=\"flex-grow pl-6\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-2\">Neptune</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine, ramps microdosing banh mi pug VHS try-hard ugh iceland kickstarter tumblr live-edge tilde.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div></div></div></section>",category:'Features'},{id:'feature-block-2',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><circle cx="56" cy="61" r="8" fill="#C3DAFE"></circle><path d="M20 87.5a1.5 1.5 0 011.5-1.5h65a1.5 1.5 0 010 3h-65a1.5 1.5 0 01-1.5-1.5zM27 94.5a1.5 1.5 0 011.5-1.5h51a1.5 1.5 0 010 3h-51a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M43 101.5a1.5 1.5 0 011.5-1.5h20a1.5 1.5 0 010 3h-20a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><rect x="25" y="77" width="58" height="4" rx="2" fill="#4A5568"></rect><circle cx="135" cy="61" r="8" fill="#C3DAFE"></circle><path d="M99 87.5a1.5 1.5 0 011.5-1.5h65a1.5 1.5 0 010 3h-65a1.5 1.5 0 01-1.5-1.5zM106 94.5a1.5 1.5 0 011.5-1.5h51a1.5 1.5 0 010 3h-51a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M122 101.5a1.5 1.5 0 011.5-1.5h20a1.5 1.5 0 010 3h-20a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><rect x="104" y="77" width="58" height="4" rx="2" fill="#4A5568"></rect><circle cx="214" cy="61" r="8" fill="#C3DAFE"></circle><path d="M178 87.5a1.5 1.5 0 011.5-1.5h65a1.5 1.5 0 010 3h-65a1.5 1.5 0 01-1.5-1.5zM185 94.5a1.5 1.5 0 011.5-1.5h51a1.5 1.5 0 010 3h-51a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M201 101.5a1.5 1.5 0 011.5-1.5h20a1.5 1.5 0 010 3h-20a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><rect x="183" y="77" width="58" height="4" rx="2" fill="#4A5568"></rect><rect x="81" y="32" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="96" y="21" width="74" height="5" rx="2.5" fill="#4A5568"></rect><rect x="113" y="120" width="40" height="10" rx="2" fill="#6366F1"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"text-center mb-20\"><h1 class=\"sm:text-3xl text-2xl font-medium title-font text-gray-900 mb-4\">Raw Denim Heirloom Man Braid</h1><p class=\"text-base leading-relaxed xl:w-2/4 lg:w-3/4 mx-auto text-gray-500s\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine, ramps microdosing banh mi pug.</p><div class=\"flex mt-6 justify-center\"><div class=\"w-16 h-1 rounded-full bg-indigo-500 inline-flex\"></div></div></div><div class=\"flex flex-wrap sm:-m-4 -mx-4 -mb-10 -mt-4 md:space-y-0 space-y-6\"><div class=\"p-4 md:w-1/3 flex flex-col text-center items-center\"><div class=\"w-20 h-20 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-5 flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-10 h-10\" viewBox=\"0 0 24 24\"><path d=\"M22 12h-4l-3 9L9 3l-3 9H2\"></path></svg></div><div class=\"flex-grow\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-3\">Shooting Stars</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine, ramps microdosing banh mi pug VHS try-hard.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div><div class=\"p-4 md:w-1/3 flex flex-col text-center items-center\"><div class=\"w-20 h-20 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-5 flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-10 h-10\" viewBox=\"0 0 24 24\"><circle cx=\"6\" cy=\"6\" r=\"3\"></circle><circle cx=\"6\" cy=\"18\" r=\"3\"></circle><path d=\"M20 4L8.12 15.88M14.47 14.48L20 20M8.12 8.12L12 12\"></path></svg></div><div class=\"flex-grow\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-3\">The Catalyzer</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine, ramps microdosing banh mi pug VHS try-hard.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div><div class=\"p-4 md:w-1/3 flex flex-col text-center items-center\"><div class=\"w-20 h-20 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-5 flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-10 h-10\" viewBox=\"0 0 24 24\"><path d=\"M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2\"></path><circle cx=\"12\" cy=\"7\" r=\"4\"></circle></svg></div><div class=\"flex-grow\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-3\">Neptune</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine, ramps microdosing banh mi pug VHS try-hard.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div></div><button class=\"flex mx-auto mt-16 text-white bg-indigo-500 border-0 py-2 px-8 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button></div></section>",category:'Features'},{id:'feature-block-3',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><circle cx="140" cy="26" r="6" fill="#C3DAFE"></circle><path d="M134 44a1 1 0 011-1h77a1 1 0 010 2h-77a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M134 49a1 1 0 011-1h21a1 1 0 010 2h-21a1 1 0 01-1-1z" fill="#6366F1"></path><rect x="134" y="37" width="58" height="3" rx="1.5" fill="#4A5568"></rect><circle cx="140" cy="66" r="6" fill="#C3DAFE"></circle><path d="M134 84a1 1 0 011-1h77a1 1 0 010 2h-77a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M134 89a1 1 0 011-1h21a1 1 0 010 2h-21a1 1 0 01-1-1z" fill="#6366F1"></path><rect x="134" y="77" width="58" height="3" rx="1.5" fill="#4A5568"></rect><circle cx="140" cy="106" r="6" fill="#C3DAFE"></circle><path d="M134 124a1 1 0 011-1h77a1 1 0 010 2h-77a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M134 129a1 1 0 011-1h21a1 1 0 010 2h-21a1 1 0 01-1-1z" fill="#6366F1"></path><rect x="134" y="117" width="58" height="3" rx="1.5" fill="#4A5568"></rect><path d="M63.792 92h26.544a3.785 3.785 0 003.792-3.778V61.778A3.785 3.785 0 0090.336 58H63.792A3.785 3.785 0 0060 61.778v26.444A3.785 3.785 0 0063.792 92zm0 0l20.856-20.778 9.48 9.445M73.272 68.389a2.839 2.839 0 01-2.844 2.833 2.839 2.839 0 01-2.844-2.833 2.839 2.839 0 012.844-2.833 2.839 2.839 0 012.844 2.833z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto flex flex-wrap\"><div class=\"lg:w-1/2 w-full mb-10 lg:mb-0 rounded-lg overflow-hidden\"><img alt=\"feature\" class=\"object-cover object-center h-full w-full\" src=\"https://dummyimage.com/460x500\"></div><div class=\"flex flex-col flex-wrap lg:py-6 -mb-10 lg:w-1/2 lg:pl-12 lg:text-left text-center\"><div class=\"flex flex-col mb-10 lg:items-start items-center\"><div class=\"w-12 h-12 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-5\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-6 h-6\" viewBox=\"0 0 24 24\"><path d=\"M22 12h-4l-3 9L9 3l-3 9H2\"></path></svg></div><div class=\"flex-grow\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-3\">Shooting Stars</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div><div class=\"flex flex-col mb-10 lg:items-start items-center\"><div class=\"w-12 h-12 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-5\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-6 h-6\" viewBox=\"0 0 24 24\"><circle cx=\"6\" cy=\"6\" r=\"3\"></circle><circle cx=\"6\" cy=\"18\" r=\"3\"></circle><path d=\"M20 4L8.12 15.88M14.47 14.48L20 20M8.12 8.12L12 12\"></path></svg></div><div class=\"flex-grow\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-3\">The Catalyzer</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div><div class=\"flex flex-col mb-10 lg:items-start items-center\"><div class=\"w-12 h-12 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-5\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-6 h-6\" viewBox=\"0 0 24 24\"><path d=\"M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2\"></path><circle cx=\"12\" cy=\"7\" r=\"4\"></circle></svg></div><div class=\"flex-grow\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-3\">Neptune</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div></div></div></section>",category:'Features'},{id:'feature-block-4',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="20.5" y="54.5" width="107" height="41" rx="1.5" fill="#FFFFFF" stroke="#CBD5E0"></rect><circle cx="32" cy="68" r="6" fill="#C3DAFE"></circle><path d="M44 72.5a1.5 1.5 0 011.5-1.5h68a1.5 1.5 0 010 3h-68a1.5 1.5 0 01-1.5-1.5zM44 79.5a1.5 1.5 0 011.5-1.5h54a1.5 1.5 0 010 3h-54a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M44 86.5a1.5 1.5 0 011.5-1.5h19a1.5 1.5 0 010 3h-19a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><rect x="44" y="62" width="32" height="4" rx="2" fill="#4A5568"></rect><rect x="138.5" y="54.5" width="107" height="41" rx="1.5" fill="#FFFFFF" stroke="#CBD5E0"></rect><circle cx="150" cy="68" r="6" fill="#C3DAFE"></circle><path d="M162 72.5a1.5 1.5 0 011.5-1.5h68a1.5 1.5 0 010 3h-68a1.5 1.5 0 01-1.5-1.5zM162 79.5a1.5 1.5 0 011.5-1.5h54a1.5 1.5 0 010 3h-54a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M162 86.5a1.5 1.5 0 011.5-1.5h19a1.5 1.5 0 010 3h-19a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><rect x="162" y="62" width="32" height="4" rx="2" fill="#4A5568"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto flex flex-wrap\"><div class=\"flex flex-wrap -m-4\"><div class=\"p-4 lg:w-1/2 md:w-full\"><div class=\"flex border-2 rounded-lg border-gray-200 border-opacity-50 p-8 sm:flex-row flex-col\"><div class=\"w-16 h-16 sm:mr-8 sm:mb-0 mb-4 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-8 h-8\" viewBox=\"0 0 24 24\"><path d=\"M22 12h-4l-3 9L9 3l-3 9H2\"></path></svg></div><div class=\"flex-grow\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-3\">Shooting Stars</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div></div><div class=\"p-4 lg:w-1/2 md:w-full\"><div class=\"flex border-2 rounded-lg border-gray-200 border-opacity-50 p-8 sm:flex-row flex-col\"><div class=\"w-16 h-16 sm:mr-8 sm:mb-0 mb-4 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-10 h-10\" viewBox=\"0 0 24 24\"><path d=\"M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2\"></path><circle cx=\"12\" cy=\"7\" r=\"4\"></circle></svg></div><div class=\"flex-grow\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-3\">The Catalyzer</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div></div></div></div></section>",category:'Features'},{id:'feature-block-5',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="20" y="78" width="69" height="31" rx="2" fill="#E2E8F0"></rect><rect x="113" y="41" width="40" height="3" rx="1.5" fill="#6366F1"></rect><circle cx="28" cy="86" r="3" fill="#6366F1"></circle><rect x="98" y="49" width="70" height="5" rx="2.5" fill="#4A5568"></rect><path d="M25 93a1 1 0 011-1h54a1 1 0 110 2H26a1 1 0 01-1-1zM25 98a1 1 0 011-1h44a1 1 0 110 2H26a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M25 103a1 1 0 011-1h11a1 1 0 010 2H26a1 1 0 01-1-1z" fill="#6366F1"></path><rect x="34" y="84.5" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="99" y="78" width="69" height="31" rx="2" fill="#E2E8F0"></rect><circle cx="107" cy="86" r="3" fill="#6366F1"></circle><path d="M104 93a1 1 0 011-1h54a1 1 0 010 2h-54a1 1 0 01-1-1zM104 98a1 1 0 011-1h44a1 1 0 010 2h-44a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M104 103a1 1 0 011-1h11a1 1 0 010 2h-11a1 1 0 01-1-1z" fill="#6366F1"></path><rect x="113" y="84.5" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="178" y="78" width="69" height="31" rx="2" fill="#E2E8F0"></rect><circle cx="186" cy="86" r="3" fill="#6366F1"></circle><path d="M183 93a1 1 0 011-1h54a1 1 0 010 2h-54a1 1 0 01-1-1zM183 98a1 1 0 011-1h44a1 1 0 010 2h-44a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M183 103a1 1 0 011-1h11a1 1 0 010 2h-11a1 1 0 01-1-1z" fill="#6366F1"></path><rect x="192" y="84.5" width="35" height="3" rx="1.5" fill="#4A5568"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-col text-center w-full mb-20\"><h2 class=\"text-xs text-indigo-500 tracking-widest font-medium title-font mb-1\">ROOF PARTY POLAROID</h2><h1 class=\"sm:text-3xl text-2xl font-medium title-font text-gray-900\">Master Cleanse Reliac Heirloom</h1></div><div class=\"flex flex-wrap -m-4\"><div class=\"p-4 md:w-1/3\"><div class=\"flex rounded-lg h-full bg-gray-100 p-8 flex-col\"><div class=\"flex items-center mb-3\"><div class=\"w-8 h-8 mr-3 inline-flex items-center justify-center rounded-full bg-indigo-500 text-white flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M22 12h-4l-3 9L9 3l-3 9H2\"></path></svg></div><h2 class=\"text-gray-900 text-lg title-font font-medium\">Shooting Stars</h2></div><div class=\"flex-grow\"><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div></div><div class=\"p-4 md:w-1/3\"><div class=\"flex rounded-lg h-full bg-gray-100 p-8 flex-col\"><div class=\"flex items-center mb-3\"><div class=\"w-8 h-8 mr-3 inline-flex items-center justify-center rounded-full bg-indigo-500 text-white flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2\"></path><circle cx=\"12\" cy=\"7\" r=\"4\"></circle></svg></div><h2 class=\"text-gray-900 text-lg title-font font-medium\">The Catalyzer</h2></div><div class=\"flex-grow\"><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div></div><div class=\"p-4 md:w-1/3\"><div class=\"flex rounded-lg h-full bg-gray-100 p-8 flex-col\"><div class=\"flex items-center mb-3\"><div class=\"w-8 h-8 mr-3 inline-flex items-center justify-center rounded-full bg-indigo-500 text-white flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><circle cx=\"6\" cy=\"6\" r=\"3\"></circle><circle cx=\"6\" cy=\"18\" r=\"3\"></circle><path d=\"M20 4L8.12 15.88M14.47 14.48L20 20M8.12 8.12L12 12\"></path></svg></div><h2 class=\"text-gray-900 text-lg title-font font-medium\">Neptune</h2></div><div class=\"flex-grow\"><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div></div></div></div></section>",category:'Features'},{id:'feature-block-6',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><circle cx="88" cy="28" r="8" fill="#C3DAFE"></circle><path d="M102 28a1 1 0 011-1h68a1 1 0 010 2h-68a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M102 33a1 1 0 011-1h14a1 1 0 010 2h-14a1 1 0 01-1-1z" fill="#6366F1"></path><rect x="102" y="21" width="40" height="3" rx="1.5" fill="#4A5568"></rect><circle cx="88" cy="98" r="8" fill="#C3DAFE"></circle><path d="M102 98a1 1 0 011-1h68a1 1 0 010 2h-68a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M102 103a1 1 0 011-1h14a1 1 0 010 2h-14a1 1 0 01-1-1z" fill="#6366F1"></path><rect x="102" y="91" width="40" height="3" rx="1.5" fill="#4A5568"></rect><circle cx="178" cy="63" r="8" fill="#C3DAFE"></circle><path d="M80 63a1 1 0 011-1h68a1 1 0 010 2H81a1 1 0 01-1-1z" fill="#A0AEC0"></path><path d="M80 68a1 1 0 011-1h14a1 1 0 110 2H81a1 1 0 01-1-1z" fill="#6366F1"></path><rect x="80" y="56" width="40" height="3" rx="1.5" fill="#4A5568"></rect><rect x="113" y="120" width="40" height="10" rx="2" fill="#6366F1"></rect><path d="M185.5 45a.5.5 0 010 1h-105a.5.5 0 010-1h105zM185.5 80a.5.5 0 010 1h-105a.5.5 0 010-1h105z" fill="#CBD5E0"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex items-center lg:w-3/5 mx-auto border-b pb-10 mb-10 border-gray-200 sm:flex-row flex-col\"><div class=\"sm:w-32 sm:h-32 h-20 w-20 sm:mr-10 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"sm:w-16 sm:h-16 w-10 h-10\" viewBox=\"0 0 24 24\"><path d=\"M22 12h-4l-3 9L9 3l-3 9H2\"></path></svg></div><div class=\"flex-grow sm:text-left text-center mt-6 sm:mt-0\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-2\">Shooting Stars</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div><div class=\"flex items-center lg:w-3/5 mx-auto border-b pb-10 mb-10 border-gray-200 sm:flex-row flex-col\"><div class=\"flex-grow sm:text-left text-center mt-6 sm:mt-0\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-2\">The Catalyzer</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div><div class=\"sm:w-32 sm:order-none order-first sm:h-32 h-20 w-20 sm:ml-10 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"sm:w-16 sm:h-16 w-10 h-10\" viewBox=\"0 0 24 24\"><circle cx=\"6\" cy=\"6\" r=\"3\"></circle><circle cx=\"6\" cy=\"18\" r=\"3\"></circle><path d=\"M20 4L8.12 15.88M14.47 14.48L20 20M8.12 8.12L12 12\"></path></svg></div></div><div class=\"flex items-center lg:w-3/5 mx-auto sm:flex-row flex-col\"><div class=\"sm:w-32 sm:h-32 h-20 w-20 sm:mr-10 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"sm:w-16 sm:h-16 w-10 h-10\" viewBox=\"0 0 24 24\"><path d=\"M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2\"></path><circle cx=\"12\" cy=\"7\" r=\"4\"></circle></svg></div><div class=\"flex-grow sm:text-left text-center mt-6 sm:mt-0\"><h2 class=\"text-gray-900 text-lg title-font font-medium mb-2\">The 400 Blows</h2><p class=\"leading-relaxed text-base\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div><button class=\"flex mx-auto mt-20 text-white bg-indigo-500 border-0 py-2 px-8 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button></div></section>",category:'Features'},{id:'feature-block-7',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M70 61a1 1 0 011-1h58a1 1 0 011 1v8a1 1 0 01-1 1H71a1 1 0 01-1-1v-8z" fill="#E2E8F0"></path><path d="M80 65a1 1 0 011-1h40a1 1 0 010 2H81a1 1 0 01-1-1z" fill="#4A5568"></path><path d="M77 65a2 2 0 11-4 0 2 2 0 014 0z" fill="#6366F1"></path><path d="M136 61a1 1 0 011-1h58a1 1 0 011 1v8a1 1 0 01-1 1h-58a1 1 0 01-1-1v-8z" fill="#E2E8F0"></path><path d="M146 65a1 1 0 011-1h40a1 1 0 010 2h-40a1 1 0 01-1-1z" fill="#4A5568"></path><path d="M143 65a2 2 0 11-4 0 2 2 0 014 0z" fill="#6366F1"></path><path d="M70 77a1 1 0 011-1h58a1 1 0 011 1v8a1 1 0 01-1 1H71a1 1 0 01-1-1v-8z" fill="#E2E8F0"></path><path d="M80 81a1 1 0 011-1h40a1 1 0 010 2H81a1 1 0 01-1-1z" fill="#4A5568"></path><path d="M77 81a2 2 0 11-4 0 2 2 0 014 0z" fill="#6366F1"></path><path d="M136 77a1 1 0 011-1h58a1 1 0 011 1v8a1 1 0 01-1 1h-58a1 1 0 01-1-1v-8z" fill="#E2E8F0"></path><path d="M146 81a1 1 0 011-1h40a1 1 0 010 2h-40a1 1 0 01-1-1z" fill="#4A5568"></path><path d="M143 81a2 2 0 11-4 0 2 2 0 014 0z" fill="#6366F1"></path><path d="M70 93a1 1 0 011-1h58a1 1 0 011 1v8a1 1 0 01-1 1H71a1 1 0 01-1-1v-8z" fill="#E2E8F0"></path><path d="M80 97a1 1 0 011-1h40a1 1 0 010 2H81a1 1 0 01-1-1z" fill="#4A5568"></path><path d="M77 97a2 2 0 11-4 0 2 2 0 014 0z" fill="#6366F1"></path><path d="M136 93a1 1 0 011-1h58a1 1 0 011 1v8a1 1 0 01-1 1h-58a1 1 0 01-1-1v-8z" fill="#E2E8F0"></path><path d="M146 97a1 1 0 011-1h40a1 1 0 010 2h-40a1 1 0 01-1-1z" fill="#4A5568"></path><path d="M143 97a2 2 0 11-4 0 2 2 0 014 0z" fill="#6366F1"></path><rect x="81" y="31" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="81" y="31" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="96" y="20" width="74" height="5" rx="2.5" fill="#4A5568"></rect><rect x="97" y="39" width="73" height="4" rx="2" fill="#A0AEC0"></rect><rect x="113" y="120" width="40" height="10" rx="2" fill="#6366F1"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"text-center mb-20\"><h1 class=\"sm:text-3xl text-2xl font-medium text-center title-font text-gray-900 mb-4\">Raw Denim Heirloom Man Braid</h1><p class=\"text-base leading-relaxed xl:w-2/4 lg:w-3/4 mx-auto\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine, ramps microdosing banh mi pug.</p></div><div class=\"flex flex-wrap lg:w-4/5 sm:mx-auto sm:mb-2 -mx-2\"><div class=\"p-2 sm:w-1/2 w-full\"><div class=\"bg-gray-100 rounded flex p-4 h-full items-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"text-indigo-500 w-6 h-6 flex-shrink-0 mr-4\" viewBox=\"0 0 24 24\"><path d=\"M22 11.08V12a10 10 0 11-5.93-9.14\"></path><path d=\"M22 4L12 14.01l-3-3\"></path></svg><span class=\"title-font font-medium\">Authentic Cliche Forage</span></div></div><div class=\"p-2 sm:w-1/2 w-full\"><div class=\"bg-gray-100 rounded flex p-4 h-full items-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"text-indigo-500 w-6 h-6 flex-shrink-0 mr-4\" viewBox=\"0 0 24 24\"><path d=\"M22 11.08V12a10 10 0 11-5.93-9.14\"></path><path d=\"M22 4L12 14.01l-3-3\"></path></svg><span class=\"title-font font-medium\">Kinfolk Chips Snackwave</span></div></div><div class=\"p-2 sm:w-1/2 w-full\"><div class=\"bg-gray-100 rounded flex p-4 h-full items-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"text-indigo-500 w-6 h-6 flex-shrink-0 mr-4\" viewBox=\"0 0 24 24\"><path d=\"M22 11.08V12a10 10 0 11-5.93-9.14\"></path><path d=\"M22 4L12 14.01l-3-3\"></path></svg><span class=\"title-font font-medium\">Coloring Book Ethical</span></div></div><div class=\"p-2 sm:w-1/2 w-full\"><div class=\"bg-gray-100 rounded flex p-4 h-full items-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"text-indigo-500 w-6 h-6 flex-shrink-0 mr-4\" viewBox=\"0 0 24 24\"><path d=\"M22 11.08V12a10 10 0 11-5.93-9.14\"></path><path d=\"M22 4L12 14.01l-3-3\"></path></svg><span class=\"title-font font-medium\">Typewriter Polaroid Cray</span></div></div><div class=\"p-2 sm:w-1/2 w-full\"><div class=\"bg-gray-100 rounded flex p-4 h-full items-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"text-indigo-500 w-6 h-6 flex-shrink-0 mr-4\" viewBox=\"0 0 24 24\"><path d=\"M22 11.08V12a10 10 0 11-5.93-9.14\"></path><path d=\"M22 4L12 14.01l-3-3\"></path></svg><span class=\"title-font font-medium\">Pack Truffaut Blue</span></div></div><div class=\"p-2 sm:w-1/2 w-full\"><div class=\"bg-gray-100 rounded flex p-4 h-full items-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"text-indigo-500 w-6 h-6 flex-shrink-0 mr-4\" viewBox=\"0 0 24 24\"><path d=\"M22 11.08V12a10 10 0 11-5.93-9.14\"></path><path d=\"M22 4L12 14.01l-3-3\"></path></svg><span class=\"title-font font-medium\">The Catcher In The Rye</span></div></div></div><button class=\"flex mx-auto mt-16 text-white bg-indigo-500 border-0 py-2 px-8 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button></div></section>",category:'Features'},{id:'feature-block-8',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="81" y="31" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="81" y="31" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="96" y="39" width="73" height="4" rx="2" fill="#A0AEC0"></rect><rect x="96" y="20" width="74" height="5" rx="2.5" fill="#4A5568"></rect><rect x="28" y="62" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="35" y="70" width="23" height="2" rx="1" fill="#A0AEC0"></rect><path d="M32 71a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="35" y="77" width="16" height="2" rx="1" fill="#A0AEC0"></rect><path d="M32 78a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="35" y="84" width="28" height="2" rx="1" fill="#A0AEC0"></rect><path d="M32 85a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="35" y="91" width="23" height="2" rx="1" fill="#A0AEC0"></rect><path d="M32 92a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="35" y="98" width="21" height="2" rx="1" fill="#A0AEC0"></rect><path d="M32 99a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="85" y="62" width="33" height="3" rx="1.5" fill="#4A5568"></rect><rect x="92" y="70" width="23" height="2" rx="1" fill="#A0AEC0"></rect><path d="M89 71a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="92" y="77" width="16" height="2" rx="1" fill="#A0AEC0"></rect><path d="M89 78a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="92" y="84" width="28" height="2" rx="1" fill="#A0AEC0"></rect><path d="M89 85a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="92" y="91" width="23" height="2" rx="1" fill="#A0AEC0"></rect><path d="M89 92a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="92" y="98" width="21" height="2" rx="1" fill="#A0AEC0"></rect><path d="M89 99a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="142" y="62" width="28" height="3" rx="1.5" fill="#4A5568"></rect><rect x="149" y="70" width="23" height="2" rx="1" fill="#A0AEC0"></rect><path d="M146 71a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="149" y="77" width="16" height="2" rx="1" fill="#A0AEC0"></rect><path d="M146 78a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="149" y="84" width="28" height="2" rx="1" fill="#A0AEC0"></rect><path d="M146 85a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="149" y="91" width="23" height="2" rx="1" fill="#A0AEC0"></rect><path d="M146 92a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="149" y="98" width="21" height="2" rx="1" fill="#A0AEC0"></rect><path d="M146 99a2 2 0 11-3.999.001A2 2 0 01146 99z" fill="#C3DAFE"></path><rect x="199" y="62" width="39" height="3" rx="1.5" fill="#4A5568"></rect><rect x="206" y="70" width="23" height="2" rx="1" fill="#A0AEC0"></rect><path d="M203 71a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="206" y="77" width="16" height="2" rx="1" fill="#A0AEC0"></rect><path d="M203 78a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="206" y="84" width="28" height="2" rx="1" fill="#A0AEC0"></rect><path d="M203 85a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="206" y="91" width="23" height="2" rx="1" fill="#A0AEC0"></rect><path d="M203 92a2 2 0 11-4 0 2 2 0 014 0z" fill="#C3DAFE"></path><rect x="206" y="98" width="21" height="2" rx="1" fill="#A0AEC0"></rect><path d="M203 99a2 2 0 11-3.999.001A2 2 0 01203 99z" fill="#C3DAFE"></path><rect x="113" y="120" width="40" height="10" rx="2" fill="#6366F1"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"text-center mb-20\"><h1 class=\"sm:text-3xl text-2xl font-medium text-center title-font text-gray-900 mb-4\">Raw Denim Heirloom Man Braid</h1><p class=\"text-base leading-relaxed xl:w-2/4 lg:w-3/4 mx-auto\">Blue bottle crucifix vinyl post-ironic four dollar toast vegan taxidermy. Gastropub indxgo juice poutine, ramps microdosing banh mi pug.</p></div><div class=\"flex flex-wrap -m-4\"><div class=\"p-4 lg:w-1/4 sm:w-1/2 w-full\"><h2 class=\"font-medium title-font tracking-widest text-gray-900 mb-4 text-sm text-center sm:text-left\">SHOOTING STARS</h2><nav class=\"flex flex-col sm:items-start sm:text-left text-center items-center -mb-1 space-y-2.5\"><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>First Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Second Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Third Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Fourth Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Fifth Link</a></nav></div><div class=\"p-4 lg:w-1/4 sm:w-1/2 w-full\"><h2 class=\"font-medium title-font tracking-widest text-gray-900 mb-4 text-sm text-center sm:text-left\">THE 400 BLOWS</h2><nav class=\"flex flex-col sm:items-start sm:text-left text-center items-center -mb-1 space-y-2.5\"><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>First Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Second Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Third Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Fourth Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Fifth Link</a></nav></div><div class=\"p-4 lg:w-1/4 sm:w-1/2 w-full\"><h2 class=\"font-medium title-font tracking-widest text-gray-900 mb-4 text-sm text-center sm:text-left\">THE CATALYZER</h2><nav class=\"flex flex-col sm:items-start sm:text-left text-center items-center -mb-1 space-y-2.5\"><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>First Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Second Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Third Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Fourth Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Fifth Link</a></nav></div><div class=\"p-4 lg:w-1/4 sm:w-1/2 w-full\"><h2 class=\"font-medium title-font tracking-widest text-gray-900 mb-4 text-sm text-center sm:text-left\">NEPTUNE</h2><nav class=\"flex flex-col sm:items-start sm:text-left text-center items-center -mb-1 space-y-2.5\"><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>First Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Second Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Third Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Fourth Link</a><a><span class=\"bg-indigo-100 text-indigo-500 w-4 h-4 mr-2 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"3\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Fifth Link</a></nav></div></div><button class=\"flex mx-auto mt-16 text-white bg-indigo-500 border-0 py-2 px-8 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button></div></section>",category:'Features'},{id:'footer-block-1',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="61" y="85" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="61" y="93" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="61" y="100" width="16" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="61" y="107" width="28" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="61" y="114" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="111" y="85" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="111" y="93" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="111" y="100" width="16" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="111" y="107" width="28" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="111" y="114" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="161" y="85" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="161" y="93" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="161" y="107" width="28" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="161" y="100" width="16" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="161" y="114" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="211" y="85" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="211" y="93" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="211" y="100" width="16" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="211" y="107" width="28" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="211" y="114" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><path fill="#E2E8F0" d="M0 131h266v19H0z"></path><rect x="20" y="139" width="41" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="222" y="139" width="25" height="3" rx="1.5" fill="#A0AEC0"></rect><circle cx="29" cy="94" r="9" fill="#6366F1"></circle></svg>',content:"<footer class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto flex md:items-center lg:items-start md:flex-row md:flex-nowrap flex-wrap flex-col\"><div class=\"w-64 flex-shrink-0 md:mx-0 mx-auto text-center md:text-left\"><a class=\"flex title-font font-medium items-center md:justify-start justify-center text-gray-900\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-10 h-10 text-white p-2 bg-indigo-500 rounded-full\" viewBox=\"0 0 24 24\"><path d=\"M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5\"></path></svg><span class=\"ml-3 text-xl\">Tailblocks</span></a><p class=\"mt-2 text-sm text-gray-500\">Air plant banjo lyft occupy retro adaptogen indego</p></div><div class=\"flex-grow flex flex-wrap md:pl-20 -mb-10 md:mt-0 mt-10 md:text-left text-center\"><div class=\"lg:w-1/4 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/4 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/4 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/4 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div></div></div><div class=\"bg-gray-100\"><div class=\"container mx-auto py-4 px-5 flex flex-wrap flex-col sm:flex-row\"><p class=\"text-gray-500 text-sm text-center sm:text-left\">© 2020 Tailblocks —<a href=\"https://twitter.com/knyttneve\" rel=\"noopener noreferrer\" class=\"text-gray-600 ml-1\" target=\"_blank\">@knyttneve</a></p><span class=\"inline-flex sm:ml-auto sm:mt-0 mt-2 justify-center sm:justify-start\"><a class=\"text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><rect width=\"20\" height=\"20\" x=\"2\" y=\"2\" rx=\"5\" ry=\"5\"></rect><path d=\"M16 11.37A4 4 0 1112.63 8 4 4 0 0116 11.37zm1.5-4.87h.01\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"currentColor\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"0\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path stroke=\"none\" d=\"M16 8a6 6 0 016 6v7h-4v-7a2 2 0 00-2-2 2 2 0 00-2 2v7h-4v-7a6 6 0 016-6zM2 9h4v12H2z\"></path><circle cx=\"4\" cy=\"4\" r=\"2\" stroke=\"none\"></circle></svg></a></span></div></div></footer>",category:'Footer'},{id:'footer-block-2',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="21" y="85" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="21" y="93" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="21" y="100" width="16" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="21" y="107" width="28" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="21" y="114" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="71" y="85" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="71" y="93" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="71" y="100" width="16" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="71" y="107" width="28" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="71" y="114" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="121" y="85" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="121" y="93" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="121" y="107" width="28" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="121" y="100" width="16" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="121" y="114" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="171" y="85" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="171" y="93" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="171" y="100" width="16" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="171" y="107" width="28" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="171" y="114" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><path fill="#E2E8F0" d="M0 131h266v19H0z"></path><rect x="20" y="139" width="41" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="222" y="139" width="25" height="3" rx="1.5" fill="#A0AEC0"></rect><circle cx="237" cy="94" r="9" fill="#6366F1"></circle></svg>',content:"<footer class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto flex md:items-center lg:items-start md:flex-row md:flex-nowrap flex-wrap flex-col\"><div class=\"w-64 flex-shrink-0 md:mx-0 mx-auto text-center md:text-left md:mt-0 mt-10\"><a class=\"flex title-font font-medium items-center md:justify-start justify-center text-gray-900\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-10 h-10 text-white p-2 bg-indigo-500 rounded-full\" viewBox=\"0 0 24 24\"><path d=\"M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5\"></path></svg><span class=\"ml-3 text-xl\">Tailblocks</span></a><p class=\"mt-2 text-sm text-gray-500\">Air plant banjo lyft occupy retro adaptogen indego</p></div><div class=\"flex-grow flex flex-wrap md:pr-20 -mb-10 md:text-left text-center order-first\"><div class=\"lg:w-1/4 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/4 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/4 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/4 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div></div></div><div class=\"bg-gray-100\"><div class=\"container mx-auto py-4 px-5 flex flex-wrap flex-col sm:flex-row\"><p class=\"text-gray-500 text-sm text-center sm:text-left\">© 2020 Tailblocks —<a href=\"https://twitter.com/knyttneve\" rel=\"noopener noreferrer\" class=\"text-gray-600 ml-1\" target=\"_blank\">@knyttneve</a></p><span class=\"inline-flex sm:ml-auto sm:mt-0 mt-2 justify-center sm:justify-start\"><a class=\"text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><rect width=\"20\" height=\"20\" x=\"2\" y=\"2\" rx=\"5\" ry=\"5\"></rect><path d=\"M16 11.37A4 4 0 1112.63 8 4 4 0 0116 11.37zm1.5-4.87h.01\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"currentColor\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"0\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path stroke=\"none\" d=\"M16 8a6 6 0 016 6v7h-4v-7a2 2 0 00-2-2 2 2 0 00-2 2v7h-4v-7a6 6 0 016-6zM2 9h4v12H2z\"></path><circle cx=\"4\" cy=\"4\" r=\"2\" stroke=\"none\"></circle></svg></a></span></div></div></footer>",category:'Footer'},{id:'footer-block-3',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="67" y="57" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="67" y="65" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="67" y="72" width="16" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="67" y="79" width="28" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="67" y="86" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="20" y="57" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="20" y="65" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="20" y="72" width="16" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="20" y="79" width="28" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="20" y="86" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="115" y="57" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="115" y="65" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="115" y="72" width="16" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="115" y="79" width="28" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="115" y="86" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="163" y="57" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="163" y="65" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="163" y="79" width="28" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="163" y="72" width="16" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="163" y="86" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="211" y="57" width="35" height="3" rx="1.5" fill="#4A5568"></rect><rect x="211" y="65" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="211" y="72" width="16" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="211" y="79" width="28" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="211" y="86" width="23" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="199" y="117" width="48" height="3" rx="1.5" fill="#A0AEC0"></rect><path stroke="#E2E8F0" d="M266 103.5H0" fill="none"></path><path d="M79 114a2 2 0 012-2h25a2 2 0 012 2v6a2 2 0 01-2 2H81a2 2 0 01-2-2v-6z" fill="#6366F1"></path><rect x="20" y="112" width="55" height="10" rx="2" fill="#CBD5E0"></rect><path fill="#E2E8F0" d="M0 131h266v19H0z"></path><rect x="20" y="139" width="41" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="222" y="139" width="25" height="3" rx="1.5" fill="#A0AEC0"></rect></svg>',content:"<footer class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-wrap md:text-left text-center -mb-10 -mx-4\"><div class=\"lg:w-1/6 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/6 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/6 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/6 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/6 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/6 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div></div></div><div class=\"border-t border-gray-200\"><div class=\"container px-5 py-8 flex flex-wrap mx-auto items-center\"><form style=\"margin: 0;\"><div class=\"flex md:flex-nowrap flex-wrap justify-center items-end md:justify-start\"><div class=\"relative sm:w-64 w-40 sm:mr-4 mr-2\"><label for=\"footer-field\" class=\"leading-7 text-sm text-gray-600\">Placeholder</label><input type=\"text\" id=\"footer-field\" name=\"footer-field\" class=\"w-full bg-gray-100 bg-opacity-50 rounded border border-gray-300 focus:ring-2 focus:bg-transparent focus:ring-indigo-200 focus:border-indigo-500 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div><button type=\"submit\" class=\"inline-flex text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded\">Button</button><p class=\"text-gray-500 text-sm md:ml-6 md:mt-0 mt-2 sm:text-left text-center\">Bitters chicharrones fanny pack<br class=\"lg:block hidden\">waistcoat green juice</p></div></form><span class=\"inline-flex lg:ml-auto lg:mt-0 mt-6 w-full justify-center md:justify-start md:w-auto\"><a class=\"text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><rect width=\"20\" height=\"20\" x=\"2\" y=\"2\" rx=\"5\" ry=\"5\"></rect><path d=\"M16 11.37A4 4 0 1112.63 8 4 4 0 0116 11.37zm1.5-4.87h.01\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"currentColor\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"0\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path stroke=\"none\" d=\"M16 8a6 6 0 016 6v7h-4v-7a2 2 0 00-2-2 2 2 0 00-2 2v7h-4v-7a6 6 0 016-6zM2 9h4v12H2z\"></path><circle cx=\"4\" cy=\"4\" r=\"2\" stroke=\"none\"></circle></svg></a></span></div></div><div class=\"bg-gray-100\"><div class=\"container mx-auto py-4 px-5 flex flex-wrap flex-col sm:flex-row\"><p class=\"text-gray-500 text-sm text-center sm:text-left\">© 2020 Tailblocks —<a href=\"https://twitter.com/knyttneve\" class=\"text-gray-600 ml-1\" target=\"_blank\" rel=\"noopener noreferrer\">@knyttneve</a></p><span class=\"sm:ml-auto sm:mt-0 mt-2 sm:w-auto w-full sm:text-left text-center text-gray-500 text-sm\">Enamel pin tousled raclette tacos irony</span></div></div></footer>",category:'Footer'},{id:'footer-block-4',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path stroke="#E2E8F0" d="M266 112.5H0" fill="none"></path><circle cx="29" cy="131" r="9" fill="#6366F1"></circle><rect x="213" y="129" width="31" height="4" rx="2" fill="#A0AEC0"></rect><rect x="53" y="129" width="45" height="4" rx="2" fill="#A0AEC0"></rect><path fill="#CBD5E0" d="M45 120h1v22h-1z"></path></svg>',content:"<footer class=\"text-gray-600 body-font\"><div class=\"container px-5 py-8 mx-auto flex items-center sm:flex-row flex-col\"><a class=\"flex title-font font-medium items-center md:justify-start justify-center text-gray-900\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-10 h-10 text-white p-2 bg-indigo-500 rounded-full\" viewBox=\"0 0 24 24\"><path d=\"M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5\"></path></svg><span class=\"ml-3 text-xl\">Tailblocks</span></a><p class=\"text-sm text-gray-500 sm:ml-4 sm:pl-4 sm:border-l-2 sm:border-gray-200 sm:py-2 sm:mt-0 mt-4\">© 2020 Tailblocks —<a href=\"https://twitter.com/knyttneve\" class=\"text-gray-600 ml-1\" rel=\"noopener noreferrer\" target=\"_blank\">@knyttneve</a></p><span class=\"inline-flex sm:ml-auto sm:mt-0 mt-4 justify-center sm:justify-start\"><a class=\"text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><rect width=\"20\" height=\"20\" x=\"2\" y=\"2\" rx=\"5\" ry=\"5\"></rect><path d=\"M16 11.37A4 4 0 1112.63 8 4 4 0 0116 11.37zm1.5-4.87h.01\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"currentColor\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"0\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path stroke=\"none\" d=\"M16 8a6 6 0 016 6v7h-4v-7a2 2 0 00-2-2 2 2 0 00-2 2v7h-4v-7a6 6 0 016-6zM2 9h4v12H2z\"></path><circle cx=\"4\" cy=\"4\" r=\"2\" stroke=\"none\"></circle></svg></a></span></div></footer>",category:'Footer'},{id:'footer-block-5',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="63.385" y="76" width="32.308" height="3" rx="1.5" fill="#4A5568"></rect><rect x="63.385" y="84" width="21.231" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="63.385" y="91" width="14.769" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="63.385" y="98" width="25.846" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="63.385" y="105" width="21.231" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="20" y="76" width="32.308" height="3" rx="1.5" fill="#4A5568"></rect><rect x="20" y="84" width="21.231" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="20" y="91" width="14.769" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="20" y="98" width="25.846" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="20" y="105" width="21.231" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="107.692" y="76" width="32.308" height="3" rx="1.5" fill="#4A5568"></rect><rect x="163" y="76" width="32.308" height="3" rx="1.5" fill="#4A5568"></rect><rect x="107.692" y="84" width="21.231" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="107.692" y="91" width="14.769" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="107.692" y="98" width="25.846" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="107.692" y="105" width="21.231" height="3" rx="1.5" fill="#A0AEC0"></rect><path d="M217 86a2 2 0 012-2h25a2 2 0 012 2v6a2 2 0 01-2 2h-25a2 2 0 01-2-2v-6z" fill="#6366F1"></path><rect x="163" y="84" width="50" height="10" rx="2" fill="#CBD5E0"></rect><path fill="#E2E8F0" d="M0 119h266v31H0z"></path><circle cx="28.5" cy="134.5" r="8.5" fill="#6366F1"></circle><rect x="45" y="133" width="30" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="198" y="133" width="48" height="3" rx="1.5" fill="#A0AEC0"></rect></svg>',content:"<footer class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-wrap md:text-left text-center order-first\"><div class=\"lg:w-1/4 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/4 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/4 md:w-1/2 w-full px-4\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">CATEGORIES</h2><nav class=\"list-none mb-10\"><li><a class=\"text-gray-600 hover:text-gray-800\">First Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Second Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Third Link</a></li><li><a class=\"text-gray-600 hover:text-gray-800\">Fourth Link</a></li></nav></div><div class=\"lg:w-1/4 md:w-1/2 w-full px-4\"><form style=\"margin: 0;\"><h2 class=\"title-font font-medium text-gray-900 tracking-widest text-sm mb-3\">SUBSCRIBE</h2><div class=\"flex xl:flex-nowrap md:flex-nowrap lg:flex-wrap flex-wrap justify-center items-end md:justify-start\"><div class=\"relative w-40 sm:w-auto xl:mr-4 lg:mr-0 sm:mr-4 mr-2\"><label for=\"footer-field\" class=\"leading-7 text-sm text-gray-600\">Placeholder</label><input type=\"text\" id=\"footer-field\" name=\"footer-field\" class=\"w-full bg-gray-100 bg-opacity-50 rounded border border-gray-300 focus:bg-transparent focus:ring-2 focus:ring-indigo-200 focus:border-indigo-500 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div><button type=\"submit\" class=\"lg:mt-2 xl:mt-0 flex-shrink-0 inline-flex text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded\">Button</button></div><p class=\"text-gray-500 text-sm mt-2 md:text-left text-center\">Bitters chicharrones fanny pack<br class=\"lg:block hidden\">waistcoat green juice</p></form></div></div></div><div class=\"bg-gray-100\"><div class=\"container px-5 py-6 mx-auto flex items-center sm:flex-row flex-col\"><a class=\"flex title-font font-medium items-center md:justify-start justify-center text-gray-900\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-10 h-10 text-white p-2 bg-indigo-500 rounded-full\" viewBox=\"0 0 24 24\"><path d=\"M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5\"></path></svg><span class=\"ml-3 text-xl\">Tailblocks</span></a><p class=\"text-sm text-gray-500 sm:ml-6 sm:mt-0 mt-4\">© 2020 Tailblocks —<a href=\"https://twitter.com/knyttneve\" rel=\"noopener noreferrer\" class=\"text-gray-600 ml-1\" target=\"_blank\">@knyttneve</a></p><span class=\"inline-flex sm:ml-auto sm:mt-0 mt-4 justify-center sm:justify-start\"><a class=\"text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><rect width=\"20\" height=\"20\" x=\"2\" y=\"2\" rx=\"5\" ry=\"5\"></rect><path d=\"M16 11.37A4 4 0 1112.63 8 4 4 0 0116 11.37zm1.5-4.87h.01\"></path></svg></a><a class=\"ml-3 text-gray-500\"><svg fill=\"currentColor\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"0\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path stroke=\"none\" d=\"M16 8a6 6 0 016 6v7h-4v-7a2 2 0 00-2-2 2 2 0 00-2 2v7h-4v-7a6 6 0 016-6zM2 9h4v12H2z\"></path><circle cx=\"4\" cy=\"4\" r=\"2\" stroke=\"none\"></circle></svg></a></span></div></div></footer>",category:'Footer'},{id:'gallery-block-1',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="142" y="32" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="32" width="74" height="5" rx="2.5" fill="#4A5568"></rect><rect x="142" y="40" width="77" height="4" rx="2" fill="#A0AEC0"></rect><path fill="#E2E8F0" d="M20 61h55v27H20zM20 91h111v39H20z"></path><path d="M70.556 118h10.888c.86 0 1.556-.696 1.556-1.556v-10.888c0-.86-.696-1.556-1.556-1.556H70.556c-.86 0-1.556.696-1.556 1.556v10.888c0 .86.696 1.556 1.556 1.556zm0 0l8.555-8.556L83 113.333m-8.556-5.055a1.166 1.166 0 11-2.332 0 1.166 1.166 0 012.332 0z" stroke="#A0AEC0" stroke-width="1.2px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><path fill="#E2E8F0" d="M78 61h53v27H78z"></path><path d="M101.111 79h7.778A1.11 1.11 0 00110 77.889V70.11a1.11 1.11 0 00-1.111-1.11h-7.778A1.11 1.11 0 00100 70.111v7.778A1.11 1.11 0 00101.111 79zm0 0l6.111-6.111L110 75.667m-6.111-3.611a.833.833 0 11-1.666 0 .833.833 0 011.666 0zM44.111 79h7.778c.613 0 1.111-.498 1.111-1.111V70.11c0-.614-.498-1.111-1.111-1.111H44.11c-.613 0-1.111.498-1.111 1.111v7.778c0 .614.498 1.111 1.111 1.111zm0 0l6.111-6.111L53 75.667m-6.111-3.611a.833.833 0 11-1.667 0 .833.833 0 011.667 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><path fill="#E2E8F0" d="M134 103h56v27h-56z"></path><path d="M158.111 122h7.778a1.11 1.11 0 001.111-1.111v-7.778a1.11 1.11 0 00-1.111-1.111h-7.778a1.11 1.11 0 00-1.111 1.111v7.778a1.11 1.11 0 001.111 1.111zm0 0l6.111-6.111 2.778 2.778m-6.111-3.611a.833.833 0 11-1.666 0 .833.833 0 011.666 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><path fill="#E2E8F0" d="M134 61h112v39H134zM193 103h53v27h-53z"></path><path d="M215.111 122h7.778a1.11 1.11 0 001.111-1.111v-7.778a1.11 1.11 0 00-1.111-1.111h-7.778a1.11 1.11 0 00-1.111 1.111v7.778a1.11 1.11 0 001.111 1.111zm0 0l6.111-6.111 2.778 2.778m-6.111-3.611a.833.833 0 11-1.666 0 .833.833 0 011.666 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><path d="M184.556 87h10.888c.86 0 1.556-.696 1.556-1.556V74.556c0-.86-.696-1.556-1.556-1.556h-10.888c-.86 0-1.556.696-1.556 1.556v10.888c0 .86.696 1.556 1.556 1.556zm0 0l8.555-8.556 3.889 3.89m-8.556-5.056a1.166 1.166 0 11-2.333 0 1.166 1.166 0 012.333 0z" stroke="#A0AEC0" stroke-width="1.2px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto flex flex-wrap\"><div class=\"flex w-full mb-20 flex-wrap\"><h1 class=\"sm:text-3xl text-2xl font-medium title-font text-gray-900 lg:w-1/3 lg:mb-0 mb-4\">Master Cleanse Reliac Heirloom</h1><p class=\"lg:pl-6 lg:w-2/3 mx-auto leading-relaxed text-base\">Whatever cardigan tote bag tumblr hexagon brooklyn asymmetrical gentrify, subway tile poke farm-to-table. Franzen you probably haven't heard of them man bun deep jianbing selfies heirloom.</p></div><div class=\"flex flex-wrap md:-m-2 -m-1\"><div class=\"flex flex-wrap w-1/2\"><div class=\"md:p-2 p-1 w-1/2\"><img alt=\"gallery\" class=\"w-full object-cover h-full object-center block\" src=\"https://dummyimage.com/500x300\"></div><div class=\"md:p-2 p-1 w-1/2\"><img alt=\"gallery\" class=\"w-full object-cover h-full object-center block\" src=\"https://dummyimage.com/501x301\"></div><div class=\"md:p-2 p-1 w-full\"><img alt=\"gallery\" class=\"w-full h-full object-cover object-center block\" src=\"https://dummyimage.com/600x360\"></div></div><div class=\"flex flex-wrap w-1/2\"><div class=\"md:p-2 p-1 w-full\"><img alt=\"gallery\" class=\"w-full h-full object-cover object-center block\" src=\"https://dummyimage.com/601x361\"></div><div class=\"md:p-2 p-1 w-1/2\"><img alt=\"gallery\" class=\"w-full object-cover h-full object-center block\" src=\"https://dummyimage.com/502x302\"></div><div class=\"md:p-2 p-1 w-1/2\"><img alt=\"gallery\" class=\"w-full object-cover h-full object-center block\" src=\"https://dummyimage.com/503x303\"></div></div></div></div></section>",category:'Gallery'},{id:'gallery-block-2',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path fill="#E2E8F0" d="M71 84h62v38H71zM71 27h127v54H71zM136 84h62v38h-62z"></path><path d="M93 54.5a1.5 1.5 0 011.5-1.5h77a1.5 1.5 0 010 3h-77a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M123 61.5a1.5 1.5 0 011.5-1.5h17a1.5 1.5 0 010 3h-17a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><rect x="108" y="45" width="50" height="4" rx="2" fill="#4A5568"></rect><path d="M81 103.5a1.5 1.5 0 011.5-1.5h39a1.5 1.5 0 010 3h-39a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M97 110.5a1.5 1.5 0 011.5-1.5h7a1.5 1.5 0 010 3h-7a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><rect x="89" y="94" width="26" height="4" rx="2" fill="#4A5568"></rect><path d="M146 103.5a1.5 1.5 0 011.5-1.5h39a1.5 1.5 0 010 3h-39a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M162 110.5a1.5 1.5 0 011.5-1.5h7a1.5 1.5 0 010 3h-7a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><rect x="154" y="94" width="26" height="4" rx="2" fill="#4A5568"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto flex flex-wrap\"><div class=\"lg:w-2/3 mx-auto\"><div class=\"flex flex-wrap w-full bg-gray-100 py-32 px-10 relative mb-4\"><img alt=\"gallery\" class=\"w-full object-cover h-full object-center block opacity-25 absolute inset-0\" src=\"https://dummyimage.com/820x340\"><div class=\"text-center relative z-10 w-full\"><h2 class=\"text-2xl text-gray-900 font-medium title-font mb-2\">Shooting Stars</h2><p class=\"leading-relaxed\">Skateboard +1 mustache fixie paleo lumbersexual.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div><div class=\"flex flex-wrap -mx-2\"><div class=\"px-2 w-1/2\"><div class=\"flex flex-wrap w-full bg-gray-100 sm:py-24 py-16 sm:px-10 px-6 relative\"><img alt=\"gallery\" class=\"w-full object-cover h-full object-center block opacity-25 absolute inset-0\" src=\"https://dummyimage.com/542x460\"><div class=\"text-center relative z-10 w-full\"><h2 class=\"text-xl text-gray-900 font-medium title-font mb-2\">Shooting Stars</h2><p class=\"leading-relaxed\">Skateboard +1 mustache fixie paleo lumbersexual.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div></div><div class=\"px-2 w-1/2\"><div class=\"flex flex-wrap w-full bg-gray-100 sm:py-24 py-16 sm:px-10 px-6 relative\"><img alt=\"gallery\" class=\"w-full object-cover h-full object-center block opacity-25 absolute inset-0\" src=\"https://dummyimage.com/542x420\"><div class=\"text-center relative z-10 w-full\"><h2 class=\"text-xl text-gray-900 font-medium title-font mb-2\">Shooting Stars</h2><p class=\"leading-relaxed\">Skateboard +1 mustache fixie paleo lumbersexual.</p><a class=\"mt-3 text-indigo-500 inline-flex items-center\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a></div></div></div></div></div></div></section>",category:'Gallery'},{id:'gallery-block-3',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path fill="#E2E8F0" d="M20 61h72v32H20zM97 61h72v32H97zM174 61h72v32h-72zM20 98h72v32H20zM97 98h72v32H97zM174 98h72v32h-72z"></path><rect x="81" y="31" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="81" y="31" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="96" y="20" width="74" height="5" rx="2.5" fill="#4A5568"></rect><rect x="97" y="39" width="73" height="4" rx="2" fill="#A0AEC0"></rect><path d="M50.556 84h10.888c.86 0 1.556-.696 1.556-1.556V71.556c0-.86-.696-1.556-1.556-1.556H50.556c-.86 0-1.556.696-1.556 1.556v10.888c0 .86.696 1.556 1.556 1.556zm0 0l8.555-8.556L63 79.334m-8.556-5.056a1.167 1.167 0 11-2.333 0 1.167 1.167 0 012.333 0zM127.556 84h10.888c.86 0 1.556-.696 1.556-1.556V71.556c0-.86-.696-1.556-1.556-1.556h-10.888c-.86 0-1.556.696-1.556 1.556v10.888c0 .86.696 1.556 1.556 1.556zm0 0l8.555-8.556 3.889 3.89m-8.556-5.056a1.166 1.166 0 11-2.333 0 1.166 1.166 0 012.333 0zM204.556 84h10.888c.86 0 1.556-.696 1.556-1.556V71.556c0-.86-.696-1.556-1.556-1.556h-10.888c-.86 0-1.556.696-1.556 1.556v10.888c0 .86.696 1.556 1.556 1.556zm0 0l8.555-8.556 3.889 3.89m-8.556-5.056a1.166 1.166 0 11-2.333 0 1.166 1.166 0 012.333 0zM204.556 121h10.888c.86 0 1.556-.696 1.556-1.556v-10.888c0-.86-.696-1.556-1.556-1.556h-10.888c-.86 0-1.556.696-1.556 1.556v10.888c0 .86.696 1.556 1.556 1.556zm0 0l8.555-8.556 3.889 3.889m-8.556-5.055a1.166 1.166 0 11-2.332 0 1.166 1.166 0 012.332 0zM127.556 121h10.888c.86 0 1.556-.696 1.556-1.556v-10.888c0-.86-.696-1.556-1.556-1.556h-10.888c-.86 0-1.556.696-1.556 1.556v10.888c0 .86.696 1.556 1.556 1.556zm0 0l8.555-8.556 3.889 3.889m-8.556-5.055a1.166 1.166 0 11-2.332 0 1.166 1.166 0 012.332 0zM50.556 121h10.888c.86 0 1.556-.696 1.556-1.556v-10.888c0-.86-.696-1.556-1.556-1.556H50.556c-.86 0-1.556.696-1.556 1.556v10.888c0 .86.696 1.556 1.556 1.556zm0 0l8.555-8.556L63 116.333m-8.556-5.055a1.166 1.166 0 11-2.332 0 1.166 1.166 0 012.332 0z" stroke="#A0AEC0" stroke-width="1.2px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-col text-center w-full mb-20\"><h1 class=\"sm:text-3xl text-2xl font-medium title-font mb-4 text-gray-900\">Master Cleanse Reliac Heirloom</h1><p class=\"lg:w-2/3 mx-auto leading-relaxed text-base\">Whatever cardigan tote bag tumblr hexagon brooklyn asymmetrical gentrify, subway tile poke farm-to-table. Franzen you probably haven't heard of them man bun deep jianbing selfies heirloom.</p></div><div class=\"flex flex-wrap -m-4\"><div class=\"lg:w-1/3 sm:w-1/2 p-4\"><div class=\"flex relative\"><img alt=\"gallery\" class=\"absolute inset-0 w-full h-full object-cover object-center\" src=\"https://dummyimage.com/600x360\"><div class=\"px-8 py-10 relative z-10 w-full border-4 border-gray-200 bg-white opacity-0 hover:opacity-100\"><h2 class=\"tracking-widest text-sm title-font font-medium text-indigo-500 mb-1\">THE SUBTITLE</h2><h1 class=\"title-font text-lg font-medium text-gray-900 mb-3\">Shooting Stars</h1><p class=\"leading-relaxed\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p></div></div></div><div class=\"lg:w-1/3 sm:w-1/2 p-4\"><div class=\"flex relative\"><img alt=\"gallery\" class=\"absolute inset-0 w-full h-full object-cover object-center\" src=\"https://dummyimage.com/601x361\"><div class=\"px-8 py-10 relative z-10 w-full border-4 border-gray-200 bg-white opacity-0 hover:opacity-100\"><h2 class=\"tracking-widest text-sm title-font font-medium text-indigo-500 mb-1\">THE SUBTITLE</h2><h1 class=\"title-font text-lg font-medium text-gray-900 mb-3\">The Catalyzer</h1><p class=\"leading-relaxed\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p></div></div></div><div class=\"lg:w-1/3 sm:w-1/2 p-4\"><div class=\"flex relative\"><img alt=\"gallery\" class=\"absolute inset-0 w-full h-full object-cover object-center\" src=\"https://dummyimage.com/603x363\"><div class=\"px-8 py-10 relative z-10 w-full border-4 border-gray-200 bg-white opacity-0 hover:opacity-100\"><h2 class=\"tracking-widest text-sm title-font font-medium text-indigo-500 mb-1\">THE SUBTITLE</h2><h1 class=\"title-font text-lg font-medium text-gray-900 mb-3\">The 400 Blows</h1><p class=\"leading-relaxed\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p></div></div></div><div class=\"lg:w-1/3 sm:w-1/2 p-4\"><div class=\"flex relative\"><img alt=\"gallery\" class=\"absolute inset-0 w-full h-full object-cover object-center\" src=\"https://dummyimage.com/602x362\"><div class=\"px-8 py-10 relative z-10 w-full border-4 border-gray-200 bg-white opacity-0 hover:opacity-100\"><h2 class=\"tracking-widest text-sm title-font font-medium text-indigo-500 mb-1\">THE SUBTITLE</h2><h1 class=\"title-font text-lg font-medium text-gray-900 mb-3\">Neptune</h1><p class=\"leading-relaxed\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p></div></div></div><div class=\"lg:w-1/3 sm:w-1/2 p-4\"><div class=\"flex relative\"><img alt=\"gallery\" class=\"absolute inset-0 w-full h-full object-cover object-center\" src=\"https://dummyimage.com/605x365\"><div class=\"px-8 py-10 relative z-10 w-full border-4 border-gray-200 bg-white opacity-0 hover:opacity-100\"><h2 class=\"tracking-widest text-sm title-font font-medium text-indigo-500 mb-1\">THE SUBTITLE</h2><h1 class=\"title-font text-lg font-medium text-gray-900 mb-3\">Holden Caulfield</h1><p class=\"leading-relaxed\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p></div></div></div><div class=\"lg:w-1/3 sm:w-1/2 p-4\"><div class=\"flex relative\"><img alt=\"gallery\" class=\"absolute inset-0 w-full h-full object-cover object-center\" src=\"https://dummyimage.com/606x366\"><div class=\"px-8 py-10 relative z-10 w-full border-4 border-gray-200 bg-white opacity-0 hover:opacity-100\"><h2 class=\"tracking-widest text-sm title-font font-medium text-indigo-500 mb-1\">THE SUBTITLE</h2><h1 class=\"title-font text-lg font-medium text-gray-900 mb-3\">Alper Kamu</h1><p class=\"leading-relaxed\">Photo booth fam kinfolk cold-pressed sriracha leggings jianbing microdosing tousled waistcoat.</p></div></div></div></div></div></section>",category:'Gallery'},{id:'header-block-1',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path stroke="#E2E8F0" d="M266 38.5H0" fill="none"></path><rect x="217" y="14" width="29" height="10" rx="2" fill="#CBD5E0"></rect><circle cx="29" cy="19" r="9" fill="#6366F1"></circle><rect x="150.132" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect><rect x="171.264" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect><rect x="192.396" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect><rect x="129" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect></svg>',content:"<header class=\"text-gray-600 body-font\"><div class=\"container mx-auto flex flex-wrap p-5 flex-col md:flex-row items-center\"><a class=\"flex title-font font-medium items-center text-gray-900 mb-4 md:mb-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-10 h-10 text-white p-2 bg-indigo-500 rounded-full\" viewBox=\"0 0 24 24\"><path d=\"M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5\"></path></svg><span class=\"ml-3 text-xl\">Tailblocks</span></a><nav class=\"md:ml-auto flex flex-wrap items-center text-base justify-center\"><a class=\"mr-5 hover:text-gray-900\">First Link</a><a class=\"mr-5 hover:text-gray-900\">Second Link</a><a class=\"mr-5 hover:text-gray-900\">Third Link</a><a class=\"mr-5 hover:text-gray-900\">Fourth Link</a></nav><button class=\"inline-flex items-center bg-gray-100 border-0 py-1 px-3 focus:outline-none hover:bg-gray-200 rounded text-base mt-4 md:mt-0\">Button<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-1\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></button></div></header>",category:'Header'},{id:'header-block-2',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path stroke="#E2E8F0" d="M266 38.5H0" fill="none"></path><rect x="141" y="14" width="29" height="10" rx="2" fill="#CBD5E0"></rect><circle cx="29" cy="19" r="9" fill="#6366F1"></circle><rect x="74.132" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect><rect x="95.264" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect><rect x="116.396" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect><rect x="53" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect><path fill="#CBD5E0" d="M45 8h1v22h-1z"></path></svg>',content:"<header class=\"text-gray-600 body-font\"><div class=\"container mx-auto flex flex-wrap p-5 flex-col md:flex-row items-center\"><a class=\"flex title-font font-medium items-center text-gray-900 mb-4 md:mb-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-10 h-10 text-white p-2 bg-indigo-500 rounded-full\" viewBox=\"0 0 24 24\"><path d=\"M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5\"></path></svg><span class=\"ml-3 text-xl\">Tailblocks</span></a><nav class=\"md:mr-auto md:ml-4 md:py-1 md:pl-4 md:border-l md:border-gray-400\tflex flex-wrap items-center text-base justify-center\"><a class=\"mr-5 hover:text-gray-900\">First Link</a><a class=\"mr-5 hover:text-gray-900\">Second Link</a><a class=\"mr-5 hover:text-gray-900\">Third Link</a><a class=\"mr-5 hover:text-gray-900\">Fourth Link</a></nav><button class=\"inline-flex items-center bg-gray-100 border-0 py-1 px-3 focus:outline-none hover:bg-gray-200 rounded text-base mt-4 md:mt-0\">Button<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-1\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></button></div></header>",category:'Header'},{id:'header-block-3',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path stroke="#E2E8F0" d="M266 38.5H0" fill="none"></path><rect x="217" y="14" width="29" height="10" rx="2" fill="#CBD5E0"></rect><circle cx="133" cy="19" r="9" fill="#6366F1"></circle><rect x="62.264" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect><rect x="41.132" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect><rect x="83.396" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect><rect x="20" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect></svg>',content:"<header class=\"text-gray-600 body-font\"><div class=\"container mx-auto flex flex-wrap p-5 flex-col md:flex-row items-center\"><nav class=\"flex lg:w-2/5 flex-wrap items-center text-base md:ml-auto\"><a class=\"mr-5 hover:text-gray-900\">First Link</a><a class=\"mr-5 hover:text-gray-900\">Second Link</a><a class=\"mr-5 hover:text-gray-900\">Third Link</a><a class=\"hover:text-gray-900\">Fourth Link</a></nav><a class=\"flex order-first lg:order-none lg:w-1/5 title-font font-medium items-center text-gray-900 lg:items-center lg:justify-center mb-4 md:mb-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-10 h-10 text-white p-2 bg-indigo-500 rounded-full\" viewBox=\"0 0 24 24\"><path d=\"M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5\"></path></svg><span class=\"ml-3 text-xl\">Tailblocks</span></a><div class=\"lg:w-2/5 inline-flex lg:justify-end ml-5 lg:ml-0\"><button class=\"inline-flex items-center bg-gray-100 border-0 py-1 px-3 focus:outline-none hover:bg-gray-200 rounded text-base mt-4 md:mt-0\">Button<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-1\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></button></div></div></header>",category:'Header'},{id:'header-block-4',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path stroke="#E2E8F0" d="M266 38.5H0" fill="none"></path><rect x="217" y="14" width="29" height="10" rx="2" fill="#CBD5E0"></rect><circle cx="29" cy="19" r="9" fill="#6366F1"></circle><rect x="129.264" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect><rect x="108.132" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect><rect x="150.396" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect><rect x="87" y="17" width="16.604" height="4" rx="2" fill="#4A5568"></rect></svg>',content:"<header class=\"text-gray-600 body-font\"><div class=\"container mx-auto flex flex-wrap p-5 flex-col md:flex-row items-center\"><a class=\"flex title-font font-medium items-center text-gray-900 mb-4 md:mb-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-10 h-10 text-white p-2 bg-indigo-500 rounded-full\" viewBox=\"0 0 24 24\"><path d=\"M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5\"></path></svg><span class=\"ml-3 text-xl\">Tailblocks</span></a><nav class=\"md:ml-auto md:mr-auto flex flex-wrap items-center text-base justify-center\"><a class=\"mr-5 hover:text-gray-900\">First Link</a><a class=\"mr-5 hover:text-gray-900\">Second Link</a><a class=\"mr-5 hover:text-gray-900\">Third Link</a><a class=\"mr-5 hover:text-gray-900\">Fourth Link</a></nav><button class=\"inline-flex items-center bg-gray-100 border-0 py-1 px-3 focus:outline-none hover:bg-gray-200 rounded text-base mt-4 md:mt-0\">Button<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-1\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></button></div></header>",category:'Header'},{id:'hero-block-1',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="133" y="86" width="29" height="10" rx="2" fill="#6366F1"></rect><rect x="168" y="86" width="29" height="10" rx="2" fill="#CBD5E0"></rect><rect x="133" y="64" width="104" height="4" rx="2" fill="#A0AEC0"></rect><rect x="133" y="53" width="72" height="5" rx="2.5" fill="#4A5568"></rect><rect x="133" y="72" width="97" height="4" rx="2" fill="#A0AEC0"></rect><path d="M62.778 92h26.444A3.778 3.778 0 0093 88.222V61.778A3.778 3.778 0 0089.222 58H62.778A3.778 3.778 0 0059 61.778v26.444A3.778 3.778 0 0062.778 92zm0 0l20.778-20.778L93 80.667M72.222 68.389a2.833 2.833 0 11-5.666 0 2.833 2.833 0 015.666 0z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container mx-auto flex px-5 py-24 md:flex-row flex-col items-center\"><div class=\"lg:max-w-lg lg:w-full md:w-1/2 w-5/6 mb-10 md:mb-0\"><img class=\"object-cover object-center rounded\" alt=\"hero\" src=\"https://dummyimage.com/720x600\"></div><div class=\"lg:flex-grow md:w-1/2 lg:pl-24 md:pl-16 flex flex-col md:items-start md:text-left items-center text-center\"><h1 class=\"title-font sm:text-4xl text-3xl mb-4 font-medium text-gray-900\">Before they sold out<br class=\"hidden lg:inline-block\">readymade gluten</h1><p class=\"mb-8 leading-relaxed\">Copper mug try-hard pitchfork pour-over freegan heirloom neutra air plant cold-pressed tacos poke beard tote bag. Heirloom echo park mlkshk tote bag selvage hot chicken authentic tumeric truffaut hexagon try-hard chambray.</p><div class=\"flex justify-center\"><button class=\"inline-flex text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button><button class=\"ml-4 inline-flex text-gray-700 bg-gray-100 border-0 py-2 px-6 focus:outline-none hover:bg-gray-200 rounded text-lg\">Button</button></div></div></div></section>",category:'Hero'},{id:'hero-block-2',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="136" y="114" width="29" height="10" rx="2" fill="#CBD5E0"></rect><rect x="101" y="114" width="29" height="10" rx="2" fill="#6366F1"></rect><rect x="81" y="92" width="104" height="4" rx="2" fill="#A0AEC0"></rect><rect x="97" y="81" width="72" height="5" rx="2.5" fill="#4A5568"></rect><rect x="85" y="100" width="97" height="4" rx="2" fill="#A0AEC0"></rect><path d="M119.778 61h26.444A3.778 3.778 0 00150 57.222V30.778A3.778 3.778 0 00146.222 27h-26.444A3.778 3.778 0 00116 30.778v26.444A3.778 3.778 0 00119.778 61zm0 0l20.778-20.778L150 49.667m-20.778-12.278a2.833 2.833 0 11-5.666 0 2.833 2.833 0 015.666 0z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container mx-auto flex px-5 py-24 items-center justify-center flex-col\"><img class=\"lg:w-2/6 md:w-3/6 w-5/6 mb-10 object-cover object-center rounded\" alt=\"hero\" src=\"https://dummyimage.com/720x600\"><div class=\"text-center lg:w-2/3 w-full\"><h1 class=\"title-font sm:text-4xl text-3xl mb-4 font-medium text-gray-900\">Microdosing synth tattooed vexillologist</h1><p class=\"mb-8 leading-relaxed\">Meggings kinfolk echo park stumptown DIY, kale chips beard jianbing tousled. Chambray dreamcatcher trust fund, kitsch vice godard disrupt ramps hexagon mustache umami snackwave tilde chillwave ugh. Pour-over meditation PBR&B pickled ennui celiac mlkshk freegan photo booth af fingerstache pitchfork.</p><div class=\"flex justify-center\"><button class=\"inline-flex text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button><button class=\"ml-4 inline-flex text-gray-700 bg-gray-100 border-0 py-2 px-6 focus:outline-none hover:bg-gray-200 rounded text-lg\">Button</button></div></div></div></section>",category:'Hero'},{id:'hero-block-3',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="20" y="86" width="29" height="10" rx="2" fill="#6366F1"></rect><rect x="55" y="86" width="29" height="10" rx="2" fill="#CBD5E0"></rect><rect x="20" y="64" width="104" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="53" width="72" height="5" rx="2.5" fill="#4A5568"></rect><rect x="20" y="72" width="97" height="4" rx="2" fill="#A0AEC0"></rect><path d="M176.778 92h26.444A3.778 3.778 0 00207 88.222V61.778A3.778 3.778 0 00203.222 58h-26.444A3.778 3.778 0 00173 61.778v26.444A3.778 3.778 0 00176.778 92zm0 0l20.778-20.778L207 80.667m-20.778-12.278a2.833 2.833 0 11-5.666 0 2.833 2.833 0 015.666 0z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container mx-auto flex px-5 py-24 md:flex-row flex-col items-center\"><div class=\"lg:flex-grow md:w-1/2 lg:pr-24 md:pr-16 flex flex-col md:items-start md:text-left mb-16 md:mb-0 items-center text-center\"><h1 class=\"title-font sm:text-4xl text-3xl mb-4 font-medium text-gray-900\">Before they sold out<br class=\"hidden lg:inline-block\">readymade gluten</h1><p class=\"mb-8 leading-relaxed\">Copper mug try-hard pitchfork pour-over freegan heirloom neutra air plant cold-pressed tacos poke beard tote bag. Heirloom echo park mlkshk tote bag selvage hot chicken authentic tumeric truffaut hexagon try-hard chambray.</p><div class=\"flex justify-center\"><button class=\"inline-flex text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button><button class=\"ml-4 inline-flex text-gray-700 bg-gray-100 border-0 py-2 px-6 focus:outline-none hover:bg-gray-200 rounded text-lg\">Button</button></div></div><div class=\"lg:max-w-lg lg:w-full md:w-1/2 w-5/6\"><img class=\"object-cover object-center rounded\" alt=\"hero\" src=\"https://dummyimage.com/720x600\"></div></div></section>",category:'Hero'},{id:'hero-block-4',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M79 88a2 2 0 012-2h25a2 2 0 012 2v6a2 2 0 01-2 2H81a2 2 0 01-2-2v-6z" fill="#6366F1"></path><rect x="20" y="86" width="55" height="10" rx="2" fill="#CBD5E0"></rect><rect x="20" y="64" width="104" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="53" width="72" height="5" rx="2.5" fill="#4A5568"></rect><rect x="20" y="72" width="97" height="4" rx="2" fill="#A0AEC0"></rect><path d="M176.778 92h26.444A3.778 3.778 0 00207 88.222V61.778A3.778 3.778 0 00203.222 58h-26.444A3.778 3.778 0 00173 61.778v26.444A3.778 3.778 0 00176.778 92zm0 0l20.778-20.778L207 80.667m-20.778-12.278a2.833 2.833 0 11-5.666 0 2.833 2.833 0 015.666 0z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container mx-auto flex px-5 py-24 md:flex-row flex-col items-center\"><div class=\"lg:flex-grow md:w-1/2 lg:pr-24 md:pr-16 flex flex-col md:items-start md:text-left mb-16 md:mb-0 items-center text-center\"><h1 class=\"title-font sm:text-4xl text-3xl mb-4 font-medium text-gray-900\">Knausgaard typewriter readymade marfa</h1><p class=\"mb-8 leading-relaxed\">Chillwave portland ugh, knausgaard fam polaroid iPhone. Man braid swag typewriter affogato, hella selvage wolf narwhal dreamcatcher.</p><form style=\"margin: 0;\"><div class=\"flex w-full md:justify-start justify-center items-end\"><div class=\"relative mr-4 md:w-full lg:w-full xl:w-1/2 w-2/4\"><label for=\"hero-field\" class=\"leading-7 text-sm text-gray-600\">Placeholder</label><input type=\"text\" id=\"hero-field\" name=\"hero-field\" class=\"w-full bg-gray-100 rounded border bg-opacity-50 border-gray-300 focus:ring-2 focus:ring-indigo-200 focus:bg-transparent focus:border-indigo-500 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div><button type=\"submit\" class=\"inline-flex text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button></div></form><p class=\"text-sm mt-2 text-gray-500 mb-8 w-full\">Neutra shabby chic ramps, viral fixie.</p><div class=\"flex lg:flex-row md:flex-col\"><button class=\"bg-gray-100 inline-flex py-3 px-5 rounded-lg items-center hover:bg-gray-200 focus:outline-none\"><svg fill=\"currentColor\" class=\"w-6 h-6\" viewBox=\"0 0 512 512\"><path d=\"M99.617 8.057a50.191 50.191 0 00-38.815-6.713l230.932 230.933 74.846-74.846L99.617 8.057zM32.139 20.116c-6.441 8.563-10.148 19.077-10.148 30.199v411.358c0 11.123 3.708 21.636 10.148 30.199l235.877-235.877L32.139 20.116zM464.261 212.087l-67.266-37.637-81.544 81.544 81.548 81.548 67.273-37.64c16.117-9.03 25.738-25.442 25.738-43.908s-9.621-34.877-25.749-43.907zM291.733 279.711L60.815 510.629c3.786.891 7.639 1.371 11.492 1.371a50.275 50.275 0 0027.31-8.07l266.965-149.372-74.849-74.847z\"></path></svg><span class=\"ml-4 flex items-start flex-col leading-none\"><span class=\"text-xs text-gray-600 mb-1\">GET IT ON</span><span class=\"title-font font-medium\">Google Play</span></span></button><button class=\"bg-gray-100 inline-flex py-3 px-5 rounded-lg items-center lg:ml-4 md:ml-0 ml-4 md:mt-4 mt-0 lg:mt-0 hover:bg-gray-200 focus:outline-none\"><svg fill=\"currentColor\" class=\"w-6 h-6\" viewBox=\"0 0 305 305\"><path d=\"M40.74 112.12c-25.79 44.74-9.4 112.65 19.12 153.82C74.09 286.52 88.5 305 108.24 305c.37 0 .74 0 1.13-.02 9.27-.37 15.97-3.23 22.45-5.99 7.27-3.1 14.8-6.3 26.6-6.3 11.22 0 18.39 3.1 25.31 6.1 6.83 2.95 13.87 6 24.26 5.81 22.23-.41 35.88-20.35 47.92-37.94a168.18 168.18 0 0021-43l.09-.28a2.5 2.5 0 00-1.33-3.06l-.18-.08c-3.92-1.6-38.26-16.84-38.62-58.36-.34-33.74 25.76-51.6 31-54.84l.24-.15a2.5 2.5 0 00.7-3.51c-18-26.37-45.62-30.34-56.73-30.82a50.04 50.04 0 00-4.95-.24c-13.06 0-25.56 4.93-35.61 8.9-6.94 2.73-12.93 5.09-17.06 5.09-4.64 0-10.67-2.4-17.65-5.16-9.33-3.7-19.9-7.9-31.1-7.9l-.79.01c-26.03.38-50.62 15.27-64.18 38.86z\"></path><path d=\"M212.1 0c-15.76.64-34.67 10.35-45.97 23.58-9.6 11.13-19 29.68-16.52 48.38a2.5 2.5 0 002.29 2.17c1.06.08 2.15.12 3.23.12 15.41 0 32.04-8.52 43.4-22.25 11.94-14.5 17.99-33.1 16.16-49.77A2.52 2.52 0 00212.1 0z\"></path></svg><span class=\"ml-4 flex items-start flex-col leading-none\"><span class=\"text-xs text-gray-600 mb-1\">Download on the</span><span class=\"title-font font-medium\">App Store</span></span></button></div></div><div class=\"lg:max-w-lg lg:w-full md:w-1/2 w-5/6\"><img class=\"object-cover object-center rounded\" alt=\"hero\" src=\"https://dummyimage.com/720x600\"></div></div></section>",category:'Hero'},{id:'hero-block-5',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M192 88a2 2 0 012-2h25a2 2 0 012 2v6a2 2 0 01-2 2h-25a2 2 0 01-2-2v-6z" fill="#6366F1"></path><rect x="133" y="86" width="55" height="10" rx="2" fill="#CBD5E0"></rect><rect x="133" y="64" width="104" height="4" rx="2" fill="#A0AEC0"></rect><rect x="133" y="53" width="72" height="5" rx="2.5" fill="#4A5568"></rect><rect x="133" y="72" width="97" height="4" rx="2" fill="#A0AEC0"></rect><path d="M62.778 92h26.444A3.778 3.778 0 0093 88.222V61.778A3.778 3.778 0 0089.222 58H62.778A3.778 3.778 0 0059 61.778v26.444A3.778 3.778 0 0062.778 92zm0 0l20.778-20.778L93 80.667M72.222 68.389a2.833 2.833 0 11-5.666 0 2.833 2.833 0 015.666 0z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container mx-auto flex px-5 py-24 md:flex-row flex-col items-center\"><div class=\"lg:max-w-lg lg:w-full md:w-1/2 w-5/6 mb-10 md:mb-0\"><img class=\"object-cover object-center rounded\" alt=\"hero\" src=\"https://dummyimage.com/720x600\"></div><div class=\"lg:flex-grow md:w-1/2 lg:pl-24 md:pl-16 flex flex-col md:items-start md:text-left items-center text-center\"><h1 class=\"title-font sm:text-4xl text-3xl mb-4 font-medium text-gray-900\">Knausgaard typewriter readymade marfa</h1><p class=\"mb-8 leading-relaxed\">Chillwave portland ugh, knausgaard fam polaroid iPhone. Man braid swag typewriter affogato, hella selvage wolf narwhal dreamcatcher.</p><form style=\"margin: 0;\"><div class=\"flex w-full md:justify-start justify-center items-end\"><div class=\"relative mr-4 lg:w-full xl:w-1/2 w-2/4\"><label for=\"hero-field\" class=\"leading-7 text-sm text-gray-600\">Placeholder</label><input type=\"text\" id=\"hero-field\" name=\"hero-field\" class=\"w-full bg-gray-100 bg-opacity-50 rounded border border-gray-300 focus:ring-2 focus:ring-indigo-200 focus:bg-transparent focus:border-indigo-500 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div><button type=\"submit\" class=\"inline-flex text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button></div></form><p class=\"text-sm mt-2 text-gray-500 mb-8 w-full\">Neutra shabby chic ramps, viral fixie.</p><div class=\"flex lg:flex-row md:flex-col\"><button class=\"bg-gray-100 inline-flex py-3 px-5 rounded-lg items-center hover:bg-gray-200 focus:outline-none\"><svg fill=\"currentColor\" class=\"w-6 h-6\" viewBox=\"0 0 512 512\"><path d=\"M99.617 8.057a50.191 50.191 0 00-38.815-6.713l230.932 230.933 74.846-74.846L99.617 8.057zM32.139 20.116c-6.441 8.563-10.148 19.077-10.148 30.199v411.358c0 11.123 3.708 21.636 10.148 30.199l235.877-235.877L32.139 20.116zM464.261 212.087l-67.266-37.637-81.544 81.544 81.548 81.548 67.273-37.64c16.117-9.03 25.738-25.442 25.738-43.908s-9.621-34.877-25.749-43.907zM291.733 279.711L60.815 510.629c3.786.891 7.639 1.371 11.492 1.371a50.275 50.275 0 0027.31-8.07l266.965-149.372-74.849-74.847z\"></path></svg><span class=\"ml-4 flex items-start flex-col leading-none\"><span class=\"text-xs text-gray-600 mb-1\">GET IT ON</span><span class=\"title-font font-medium\">Google Play</span></span></button><button class=\"bg-gray-100 inline-flex py-3 px-5 rounded-lg items-center lg:ml-4 md:ml-0 ml-4 md:mt-4 mt-0 lg:mt-0 hover:bg-gray-200 focus:outline-none\"><svg fill=\"currentColor\" class=\"w-6 h-6\" viewBox=\"0 0 305 305\"><path d=\"M40.74 112.12c-25.79 44.74-9.4 112.65 19.12 153.82C74.09 286.52 88.5 305 108.24 305c.37 0 .74 0 1.13-.02 9.27-.37 15.97-3.23 22.45-5.99 7.27-3.1 14.8-6.3 26.6-6.3 11.22 0 18.39 3.1 25.31 6.1 6.83 2.95 13.87 6 24.26 5.81 22.23-.41 35.88-20.35 47.92-37.94a168.18 168.18 0 0021-43l.09-.28a2.5 2.5 0 00-1.33-3.06l-.18-.08c-3.92-1.6-38.26-16.84-38.62-58.36-.34-33.74 25.76-51.6 31-54.84l.24-.15a2.5 2.5 0 00.7-3.51c-18-26.37-45.62-30.34-56.73-30.82a50.04 50.04 0 00-4.95-.24c-13.06 0-25.56 4.93-35.61 8.9-6.94 2.73-12.93 5.09-17.06 5.09-4.64 0-10.67-2.4-17.65-5.16-9.33-3.7-19.9-7.9-31.1-7.9l-.79.01c-26.03.38-50.62 15.27-64.18 38.86z\"></path><path d=\"M212.1 0c-15.76.64-34.67 10.35-45.97 23.58-9.6 11.13-19 29.68-16.52 48.38a2.5 2.5 0 002.29 2.17c1.06.08 2.15.12 3.23.12 15.41 0 32.04-8.52 43.4-22.25 11.94-14.5 17.99-33.1 16.16-49.77A2.52 2.52 0 00212.1 0z\"></path></svg><span class=\"ml-4 flex items-start flex-col leading-none\"><span class=\"text-xs text-gray-600 mb-1\">Download on the</span><span class=\"title-font font-medium\">App Store</span></span></button></div></div></div></section>",category:'Hero'},{id:'hero-block-6',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="81" y="92" width="104" height="4" rx="2" fill="#A0AEC0"></rect><rect x="97" y="81" width="72" height="5" rx="2.5" fill="#4A5568"></rect><path d="M148 116a2 2 0 012-2h25a2 2 0 012 2v6a2 2 0 01-2 2h-25a2 2 0 01-2-2v-6z" fill="#6366F1"></path><rect x="89" y="114" width="55" height="10" rx="2" fill="#CBD5E0"></rect><rect x="85" y="100" width="97" height="4" rx="2" fill="#A0AEC0"></rect><path d="M119.778 61h26.444A3.778 3.778 0 00150 57.222V30.778A3.778 3.778 0 00146.222 27h-26.444A3.778 3.778 0 00116 30.778v26.444A3.778 3.778 0 00119.778 61zm0 0l20.778-20.778L150 49.667m-20.778-12.278a2.833 2.833 0 11-5.666 0 2.833 2.833 0 015.666 0z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container mx-auto flex flex-col px-5 py-24 justify-center items-center\"><img class=\"lg:w-2/6 md:w-3/6 w-5/6 mb-10 object-cover object-center rounded\" alt=\"hero\" src=\"https://dummyimage.com/720x600\"><div class=\"w-full md:w-2/3 flex flex-col mb-16 items-center text-center\"><h1 class=\"title-font sm:text-4xl text-3xl mb-4 font-medium text-gray-900\">Knausgaard typewriter readymade marfa</h1><p class=\"mb-8 leading-relaxed\">Kickstarter biodiesel roof party wayfarers cold-pressed. Palo santo live-edge tumeric scenester copper mug flexitarian. Prism vice offal plaid everyday carry. Gluten-free chia VHS squid listicle artisan.</p><form style=\"margin: 0;\"><div class=\"flex w-full justify-center items-end\"><div class=\"relative mr-4 lg:w-full xl:w-1/2 w-2/4 md:w-full text-left\"><label for=\"hero-field\" class=\"leading-7 text-sm text-gray-600\">Placeholder</label><input type=\"text\" id=\"hero-field\" name=\"hero-field\" class=\"w-full bg-gray-100 bg-opacity-50 rounded focus:ring-2 focus:ring-indigo-200 focus:bg-transparent border border-gray-300 focus:border-indigo-500 text-base outline-none text-gray-700 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out\" required></div><button type=\"submit\" class=\"inline-flex text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded text-lg\">Button</button></div></form><p class=\"text-sm mt-2 text-gray-500 mb-8 w-full\">Neutra shabby chic ramps, viral fixie.</p><div class=\"flex\"><button class=\"bg-gray-100 inline-flex py-3 px-5 rounded-lg items-center hover:bg-gray-200 focus:outline-none\"><svg fill=\"currentColor\" class=\"w-6 h-6\" viewBox=\"0 0 512 512\"><path d=\"M99.617 8.057a50.191 50.191 0 00-38.815-6.713l230.932 230.933 74.846-74.846L99.617 8.057zM32.139 20.116c-6.441 8.563-10.148 19.077-10.148 30.199v411.358c0 11.123 3.708 21.636 10.148 30.199l235.877-235.877L32.139 20.116zM464.261 212.087l-67.266-37.637-81.544 81.544 81.548 81.548 67.273-37.64c16.117-9.03 25.738-25.442 25.738-43.908s-9.621-34.877-25.749-43.907zM291.733 279.711L60.815 510.629c3.786.891 7.639 1.371 11.492 1.371a50.275 50.275 0 0027.31-8.07l266.965-149.372-74.849-74.847z\"></path></svg><span class=\"ml-4 flex items-start flex-col leading-none\"><span class=\"text-xs text-gray-600 mb-1\">GET IT ON</span><span class=\"title-font font-medium\">Google Play</span></span></button><button class=\"bg-gray-100 inline-flex py-3 px-5 rounded-lg items-center ml-4 hover:bg-gray-200 focus:outline-none\"><svg fill=\"currentColor\" class=\"w-6 h-6\" viewBox=\"0 0 305 305\"><path d=\"M40.74 112.12c-25.79 44.74-9.4 112.65 19.12 153.82C74.09 286.52 88.5 305 108.24 305c.37 0 .74 0 1.13-.02 9.27-.37 15.97-3.23 22.45-5.99 7.27-3.1 14.8-6.3 26.6-6.3 11.22 0 18.39 3.1 25.31 6.1 6.83 2.95 13.87 6 24.26 5.81 22.23-.41 35.88-20.35 47.92-37.94a168.18 168.18 0 0021-43l.09-.28a2.5 2.5 0 00-1.33-3.06l-.18-.08c-3.92-1.6-38.26-16.84-38.62-58.36-.34-33.74 25.76-51.6 31-54.84l.24-.15a2.5 2.5 0 00.7-3.51c-18-26.37-45.62-30.34-56.73-30.82a50.04 50.04 0 00-4.95-.24c-13.06 0-25.56 4.93-35.61 8.9-6.94 2.73-12.93 5.09-17.06 5.09-4.64 0-10.67-2.4-17.65-5.16-9.33-3.7-19.9-7.9-31.1-7.9l-.79.01c-26.03.38-50.62 15.27-64.18 38.86z\"></path><path d=\"M212.1 0c-15.76.64-34.67 10.35-45.97 23.58-9.6 11.13-19 29.68-16.52 48.38a2.5 2.5 0 002.29 2.17c1.06.08 2.15.12 3.23.12 15.41 0 32.04-8.52 43.4-22.25 11.94-14.5 17.99-33.1 16.16-49.77A2.52 2.52 0 00212.1 0z\"></path></svg><span class=\"ml-4 flex items-start flex-col leading-none\"><span class=\"text-xs text-gray-600 mb-1\">Download on the</span><span class=\"title-font font-medium\">App Store</span></span></button></div></div></div></section>",category:'Hero'},{id:'pricing-block-1',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="96" y="22" width="74" height="5" rx="2.5" fill="#4A5568"></rect><rect x="87" y="31" width="92" height="4" rx="2" fill="#A0AEC0"></rect><path d="M65.5 83a.5.5 0 010 1h-41a.5.5 0 010-1h41z" fill="#E2E8F0"></path><rect x="20.5" y="69.5" width="49" height="58" rx="1.5" stroke="#CBD5E0" fill="none"></rect><circle cx="26" cy="89" r="2" fill="#A0AEC0"></circle><rect x="24" y="77" width="18" height="2" rx="1" fill="#4A5568"></rect><rect x="30" y="88" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="26" cy="95" r="2" fill="#A0AEC0"></circle><rect x="30" y="94" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="26" cy="101" r="2" fill="#A0AEC0"></circle><rect x="30" y="100" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="26" cy="107" r="2" fill="#A0AEC0"></circle><rect x="30" y="106" width="34" height="2" rx="1" fill="#A0AEC0"></rect><rect x="24" y="73" width="10" height="2" rx="1" fill="#4A5568"></rect><path d="M24 116.5a2 2 0 012-2h38a2 2 0 012 2v4a2 2 0 01-2 2H26a2 2 0 01-2-2v-4z" fill="#A0AEC0"></path><path d="M124.5 83a.5.5 0 010 1h-41a.5.5 0 010-1h41z" fill="#E2E8F0"></path><rect x="79.5" y="69.5" width="49" height="58" rx="1.5" stroke="#6366F1" fill="none"></rect><circle cx="85" cy="89" r="2" fill="#A0AEC0"></circle><rect x="83" y="77" width="18" height="2" rx="1" fill="#4A5568"></rect><rect x="89" y="88" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="85" cy="95" r="2" fill="#A0AEC0"></circle><rect x="89" y="94" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="85" cy="101" r="2" fill="#A0AEC0"></circle><rect x="89" y="100" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="85" cy="107" r="2" fill="#A0AEC0"></circle><rect x="89" y="106" width="34" height="2" rx="1" fill="#A0AEC0"></rect><rect x="83" y="73" width="10" height="2" rx="1" fill="#4A5568"></rect><path d="M183.5 83a.5.5 0 010 1h-41a.5.5 0 010-1h41z" fill="#E2E8F0"></path><rect x="138.5" y="69.5" width="49" height="58" rx="1.5" stroke="#CBD5E0" fill="none"></rect><circle cx="144" cy="89" r="2" fill="#A0AEC0"></circle><rect x="142" y="77" width="18" height="2" rx="1" fill="#4A5568"></rect><rect x="148" y="88" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="144" cy="95" r="2" fill="#A0AEC0"></circle><rect x="148" y="94" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="144" cy="101" r="2" fill="#A0AEC0"></circle><rect x="148" y="100" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="144" cy="107" r="2" fill="#A0AEC0"></circle><rect x="148" y="106" width="34" height="2" rx="1" fill="#A0AEC0"></rect><rect x="142" y="73" width="10" height="2" rx="1" fill="#4A5568"></rect><path d="M142 116.5a2 2 0 012-2h38a2 2 0 012 2v4a2 2 0 01-2 2h-38a2 2 0 01-2-2v-4z" fill="#A0AEC0"></path><path d="M83 116.5a2 2 0 012-2h38a2 2 0 012 2v4a2 2 0 01-2 2H85a2 2 0 01-2-2v-4z" fill="#6366F1"></path><path d="M242.5 83a.5.5 0 010 1h-41a.5.5 0 010-1h41z" fill="#E2E8F0"></path><rect x="197.5" y="69.5" width="49" height="58" rx="1.5" stroke="#CBD5E0" fill="none"></rect><circle cx="203" cy="89" r="2" fill="#A0AEC0"></circle><rect x="201" y="77" width="18" height="2" rx="1" fill="#4A5568"></rect><rect x="207" y="88" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="203" cy="95" r="2" fill="#A0AEC0"></circle><rect x="207" y="94" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="203" cy="101" r="2" fill="#A0AEC0"></circle><rect x="207" y="100" width="34" height="2" rx="1" fill="#A0AEC0"></rect><circle cx="203" cy="107" r="2" fill="#A0AEC0"></circle><rect x="207" y="106" width="34" height="2" rx="1" fill="#A0AEC0"></rect><rect x="201" y="73" width="10" height="2" rx="1" fill="#4A5568"></rect><path d="M201 116.5a2 2 0 012-2h38a2 2 0 012 2v4a2 2 0 01-2 2h-38a2 2 0 01-2-2v-4z" fill="#A0AEC0"></path><path d="M118 43.5h30a1.5 1.5 0 011.5 1.5v4a1.5 1.5 0 01-1.5 1.5h-30a1.5 1.5 0 01-1.5-1.5v-4a1.5 1.5 0 011.5-1.5z" stroke="#6366F1" fill="none"></path><path fill="#6366F1" d="M117 43h16v7h-16z"></path></svg>',content:"<section class=\"text-gray-600 body-font overflow-hidden\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-col text-center w-full mb-20\"><h1 class=\"sm:text-4xl text-3xl font-medium title-font mb-2 text-gray-900\">Pricing</h1><p class=\"lg:w-2/3 mx-auto leading-relaxed text-base text-gray-500\">Whatever cardigan tote bag tumblr hexagon brooklyn asymmetrical.</p></div><div class=\"flex flex-wrap -m-4\"><div class=\"p-4 xl:w-1/4 md:w-1/2 w-full\"><div class=\"h-full p-6 rounded-lg border-2 border-gray-300 flex flex-col relative overflow-hidden\"><h2 class=\"text-sm tracking-widest title-font mb-1 font-medium\">START</h2><h1 class=\"text-5xl text-gray-900 pb-4 mb-4 border-b border-gray-200 leading-none\">Free</h1><p class=\"flex items-center text-gray-600 mb-2\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Vexillologist pitchfork</p><p class=\"flex items-center text-gray-600 mb-2\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Tumeric plaid portland</p><p class=\"flex items-center text-gray-600 mb-6\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Mixtape chillwave tumeric</p><button class=\"flex items-center mt-auto text-white bg-gray-400 border-0 py-2 px-4 w-full focus:outline-none hover:bg-gray-500 rounded\">Button<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-auto\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></button><p class=\"text-xs text-gray-500 mt-3\">Literally you probably haven't heard of them jean shorts.</p></div></div><div class=\"p-4 xl:w-1/4 md:w-1/2 w-full\"><div class=\"h-full p-6 rounded-lg border-2 border-indigo-500 flex flex-col relative overflow-hidden\"><span class=\"bg-indigo-500 text-white px-3 py-1 tracking-widest text-xs absolute right-0 top-0 rounded-bl\">POPULAR</span><h2 class=\"text-sm tracking-widest title-font mb-1 font-medium\">PRO</h2><h1 class=\"text-5xl text-gray-900 leading-none flex items-center pb-4 mb-4 border-b border-gray-200\"><span>$38</span><span class=\"text-lg ml-1 font-normal text-gray-500\">/mo</span></h1><p class=\"flex items-center text-gray-600 mb-2\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Vexillologist pitchfork</p><p class=\"flex items-center text-gray-600 mb-2\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Tumeric plaid portland</p><p class=\"flex items-center text-gray-600 mb-2\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Hexagon neutra unicorn</p><p class=\"flex items-center text-gray-600 mb-6\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Mixtape chillwave tumeric</p><button class=\"flex items-center mt-auto text-white bg-indigo-500 border-0 py-2 px-4 w-full focus:outline-none hover:bg-indigo-600 rounded\">Button<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-auto\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></button><p class=\"text-xs text-gray-500 mt-3\">Literally you probably haven't heard of them jean shorts.</p></div></div><div class=\"p-4 xl:w-1/4 md:w-1/2 w-full\"><div class=\"h-full p-6 rounded-lg border-2 border-gray-300 flex flex-col relative overflow-hidden\"><h2 class=\"text-sm tracking-widest title-font mb-1 font-medium\">BUSINESS</h2><h1 class=\"text-5xl text-gray-900 leading-none flex items-center pb-4 mb-4 border-b border-gray-200\"><span>$56</span><span class=\"text-lg ml-1 font-normal text-gray-500\">/mo</span></h1><p class=\"flex items-center text-gray-600 mb-2\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Vexillologist pitchfork</p><p class=\"flex items-center text-gray-600 mb-2\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Tumeric plaid portland</p><p class=\"flex items-center text-gray-600 mb-2\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Hexagon neutra unicorn</p><p class=\"flex items-center text-gray-600 mb-2\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Vexillologist pitchfork</p><p class=\"flex items-center text-gray-600 mb-6\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Mixtape chillwave tumeric</p><button class=\"flex items-center mt-auto text-white bg-gray-400 border-0 py-2 px-4 w-full focus:outline-none hover:bg-gray-500 rounded\">Button<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-auto\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></button><p class=\"text-xs text-gray-500 mt-3\">Literally you probably haven't heard of them jean shorts.</p></div></div><div class=\"p-4 xl:w-1/4 md:w-1/2 w-full\"><div class=\"h-full p-6 rounded-lg border-2 border-gray-300 flex flex-col relative overflow-hidden\"><h2 class=\"text-sm tracking-widest title-font mb-1 font-medium\">SPECIAL</h2><h1 class=\"text-5xl text-gray-900 leading-none flex items-center pb-4 mb-4 border-b border-gray-200\"><span>$72</span><span class=\"text-lg ml-1 font-normal text-gray-500\">/mo</span></h1><p class=\"flex items-center text-gray-600 mb-2\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Vexillologist pitchfork</p><p class=\"flex items-center text-gray-600 mb-2\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Tumeric plaid portland</p><p class=\"flex items-center text-gray-600 mb-2\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Hexagon neutra unicorn</p><p class=\"flex items-center text-gray-600 mb-2\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Vexillologist pitchfork</p><p class=\"flex items-center text-gray-600 mb-6\"><span class=\"w-4 h-4 mr-2 inline-flex items-center justify-center bg-gray-400 text-white rounded-full flex-shrink-0\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2.5\" class=\"w-3 h-3\" viewBox=\"0 0 24 24\"><path d=\"M20 6L9 17l-5-5\"></path></svg></span>Mixtape chillwave tumeric</p><button class=\"flex items-center mt-auto text-white bg-gray-400 border-0 py-2 px-4 w-full focus:outline-none hover:bg-gray-500 rounded\">Button<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-auto\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></button><p class=\"text-xs text-gray-500 mt-3\">Literally you probably haven't heard of them jean shorts.</p></div></div></div></div></section>",category:'Pricing'},{id:'pricing-block-2',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="110" y="20" width="46" height="5" rx="2.5" fill="#4A5568"></rect><rect x="87" y="29" width="92" height="4" rx="2" fill="#A0AEC0"></rect><rect x="106" y="37" width="55" height="4" rx="2" fill="#A0AEC0"></rect><rect x="50" y="57" width="167" height="11" rx="1" fill="#E2E8F0"></rect><rect x="55" y="61" width="39" height="3" rx="1.5" fill="#4A5568"></rect><rect x="55" y="121" width="19" height="3" rx="1.5" fill="#6366F1"></rect><rect x="108" y="61" width="24" height="3" rx="1.5" fill="#4A5568"></rect><rect x="145" y="61" width="30" height="3" rx="1.5" fill="#4A5568"></rect><rect x="188" y="61" width="20" height="3" rx="1.5" fill="#4A5568"></rect><rect x="55" y="74" width="26" height="2" rx="1" fill="#A0AEC0"></rect><rect x="108" y="74" width="12" height="2" rx="1" fill="#A0AEC0"></rect><path d="M216.5 82a.5.5 0 010 1h-166a.5.5 0 010-1h166z" fill="#E2E8F0"></path><rect x="145" y="74" width="13" height="2" rx="1" fill="#A0AEC0"></rect><rect x="188" y="74" width="20" height="2" rx="1" fill="#A0AEC0"></rect><rect x="55" y="89" width="39" height="2" rx="1" fill="#A0AEC0"></rect><rect x="108" y="89" width="17" height="2" rx="1" fill="#A0AEC0"></rect><path d="M216.5 97a.5.5 0 010 1h-166a.5.5 0 010-1h166z" fill="#E2E8F0"></path><rect x="145" y="89" width="18" height="2" rx="1" fill="#A0AEC0"></rect><rect x="188" y="89" width="13" height="2" rx="1" fill="#A0AEC0"></rect><rect x="55" y="104" width="33" height="2" rx="1" fill="#A0AEC0"></rect><rect x="108" y="104" width="14" height="2" rx="1" fill="#A0AEC0"></rect><path d="M216.5 112a.5.5 0 010 1h-166a.5.5 0 010-1h166z" fill="#E2E8F0"></path><rect x="182" y="119" width="31" height="10" rx="2" fill="#6366F1"></rect><rect x="145" y="104" width="18" height="2" rx="1" fill="#A0AEC0"></rect><rect x="188" y="104" width="20" height="2" rx="1" fill="#A0AEC0"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-col text-center w-full mb-20\"><h1 class=\"sm:text-4xl text-3xl font-medium title-font mb-2 text-gray-900\">Pricing</h1><p class=\"lg:w-2/3 mx-auto leading-relaxed text-base\">Banh mi cornhole echo park skateboard authentic crucifix neutra tilde lyft biodiesel artisan direct trade mumblecore 3 wolf moon twee</p></div><div class=\"lg:w-2/3 w-full mx-auto overflow-auto\"><table class=\"table-auto w-full text-left whitespace-no-wrap\"><thead><tr><th class=\"px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-100 rounded-tl rounded-bl\">Plan</th><th class=\"px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-100\">Speed</th><th class=\"px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-100\">Storage</th><th class=\"px-4 py-3 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-100\">Price</th><th class=\"w-10 title-font tracking-wider font-medium text-gray-900 text-sm bg-gray-100 rounded-tr rounded-br\"></th></tr></thead><tbody><tr><td class=\"px-4 py-3\">Start</td><td class=\"px-4 py-3\">5 Mb/s</td><td class=\"px-4 py-3\">15 GB</td><td class=\"px-4 py-3 text-lg text-gray-900\">Free</td><td class=\"w-10 text-center\"><input name=\"plan\" type=\"radio\"></td></tr><tr><td class=\"border-t-2 border-gray-200 px-4 py-3\">Pro</td><td class=\"border-t-2 border-gray-200 px-4 py-3\">25 Mb/s</td><td class=\"border-t-2 border-gray-200 px-4 py-3\">25 GB</td><td class=\"border-t-2 border-gray-200 px-4 py-3 text-lg text-gray-900\">$24</td><td class=\"border-t-2 border-gray-200 w-10 text-center\"><input name=\"plan\" type=\"radio\"></td></tr><tr><td class=\"border-t-2 border-gray-200 px-4 py-3\">Business</td><td class=\"border-t-2 border-gray-200 px-4 py-3\">36 Mb/s</td><td class=\"border-t-2 border-gray-200 px-4 py-3\">40 GB</td><td class=\"border-t-2 border-gray-200 px-4 py-3 text-lg text-gray-900\">$50</td><td class=\"border-t-2 border-gray-200 w-10 text-center\"><input name=\"plan\" type=\"radio\"></td></tr><tr><td class=\"border-t-2 border-b-2 border-gray-200 px-4 py-3\">Exclusive</td><td class=\"border-t-2 border-b-2 border-gray-200 px-4 py-3\">48 Mb/s</td><td class=\"border-t-2 border-b-2 border-gray-200 px-4 py-3\">120 GB</td><td class=\"border-t-2 border-b-2 border-gray-200 px-4 py-3 text-lg text-gray-900\">$72</td><td class=\"border-t-2 border-b-2 border-gray-200 w-10 text-center\"><input name=\"plan\" type=\"radio\"></td></tr></tbody></table></div><div class=\"flex pl-4 mt-4 lg:w-2/3 w-full mx-auto\"><a class=\"text-indigo-500 inline-flex items-center md:mb-2 lg:mb-0\">Learn More<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-4 h-4 ml-2\" viewBox=\"0 0 24 24\"><path d=\"M5 12h14M12 5l7 7-7 7\"></path></svg></a><button class=\"flex ml-auto text-white bg-indigo-500 border-0 py-2 px-6 focus:outline-none hover:bg-indigo-600 rounded\">Button</button></div></div></section>",category:'Pricing'},{id:'statistic-block-1',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="45" y="66" width="26" height="10" rx="5" fill="#4A5568"></rect><rect x="43" y="80" width="30" height="4" rx="2" fill="#A0AEC0"></rect><rect x="95" y="66" width="26" height="10" rx="5" fill="#4A5568"></rect><rect x="93" y="80" width="30" height="4" rx="2" fill="#A0AEC0"></rect><rect x="145" y="66" width="26" height="10" rx="5" fill="#4A5568"></rect><rect x="143" y="80" width="30" height="4" rx="2" fill="#A0AEC0"></rect><rect x="195" y="66" width="26" height="10" rx="5" fill="#4A5568"></rect><rect x="193" y="80" width="30" height="4" rx="2" fill="#A0AEC0"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-wrap -m-4 text-center\"><div class=\"p-4 sm:w-1/4 w-1/2\"><h2 class=\"title-font font-medium sm:text-4xl text-3xl text-gray-900\">2.7K</h2><p class=\"leading-relaxed\">Users</p></div><div class=\"p-4 sm:w-1/4 w-1/2\"><h2 class=\"title-font font-medium sm:text-4xl text-3xl text-gray-900\">1.8K</h2><p class=\"leading-relaxed\">Subscribes</p></div><div class=\"p-4 sm:w-1/4 w-1/2\"><h2 class=\"title-font font-medium sm:text-4xl text-3xl text-gray-900\">35</h2><p class=\"leading-relaxed\">Downloads</p></div><div class=\"p-4 sm:w-1/4 w-1/2\"><h2 class=\"title-font font-medium sm:text-4xl text-3xl text-gray-900\">4</h2><p class=\"leading-relaxed\">Products</p></div></div></div></section>",category:'Statistics'},{id:'statistic-block-2',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M175.792 92h26.544a3.785 3.785 0 003.792-3.778V61.778A3.785 3.785 0 00202.336 58h-26.544A3.785 3.785 0 00172 61.778v26.444A3.785 3.785 0 00175.792 92zm0 0l20.856-20.778 9.48 9.445m-20.856-12.278a2.838 2.838 0 01-2.844 2.833 2.838 2.838 0 01-2.844-2.833 2.838 2.838 0 012.844-2.833 2.838 2.838 0 012.844 2.833z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="20" y="46" width="70" height="5" rx="2.5" fill="#4A5568"></rect><rect x="20" y="57" width="98" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="65" width="87" height="4" rx="2" fill="#A0AEC0"></rect><rect x="20" y="88" width="17" height="7" rx="3.5" fill="#4A5568"></rect><rect x="20" y="99" width="20" height="4" rx="2" fill="#A0AEC0"></rect><rect x="46" y="88" width="17" height="7" rx="3.5" fill="#4A5568"></rect><rect x="46" y="99" width="20" height="4" rx="2" fill="#A0AEC0"></rect><rect x="72" y="88" width="17" height="7" rx="3.5" fill="#4A5568"></rect><rect x="72" y="99" width="20" height="4" rx="2" fill="#A0AEC0"></rect><rect x="98" y="88" width="17" height="7" rx="3.5" fill="#4A5568"></rect><rect x="98" y="99" width="20" height="4" rx="2" fill="#A0AEC0"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto flex flex-wrap\"><div class=\"flex flex-wrap -mx-4 mt-auto mb-auto lg:w-1/2 sm:w-2/3 content-start sm:pr-10\"><div class=\"w-full sm:p-4 px-4 mb-6\"><h1 class=\"title-font font-medium text-xl mb-2 text-gray-900\">Moon hashtag pop-up try-hard offal truffaut</h1><div class=\"leading-relaxed\">Pour-over craft beer pug drinking vinegar live-edge gastropub, keytar neutra sustainable fingerstache kickstarter.</div></div><div class=\"p-4 sm:w-1/2 lg:w-1/4 w-1/2\"><h2 class=\"title-font font-medium text-3xl text-gray-900\">2.7K</h2><p class=\"leading-relaxed\">Users</p></div><div class=\"p-4 sm:w-1/2 lg:w-1/4 w-1/2\"><h2 class=\"title-font font-medium text-3xl text-gray-900\">1.8K</h2><p class=\"leading-relaxed\">Subscribes</p></div><div class=\"p-4 sm:w-1/2 lg:w-1/4 w-1/2\"><h2 class=\"title-font font-medium text-3xl text-gray-900\">35</h2><p class=\"leading-relaxed\">Downloads</p></div><div class=\"p-4 sm:w-1/2 lg:w-1/4 w-1/2\"><h2 class=\"title-font font-medium text-3xl text-gray-900\">4</h2><p class=\"leading-relaxed\">Products</p></div></div><div class=\"lg:w-1/2 sm:w-1/3 w-full rounded-lg overflow-hidden mt-6 sm:mt-0\"><img class=\"object-cover object-center w-full h-full\" src=\"https://dummyimage.com/600x300\" alt=\"stats\"></div></div></section>",category:'Statistics'},{id:'statistic-block-3',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M21 64.5h48a.5.5 0 01.5.5v33a.5.5 0 01-.5.5H21a.5.5 0 01-.5-.5V65a.5.5 0 01.5-.5z" stroke="#CBD5E0" fill="none"></path><path d="M50 75a5 5 0 11-10 0 5 5 0 0110 0z" fill="#6366F1"></path><path d="M39 92a1 1 0 011-1h10a1 1 0 110 2H40a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="35" y="84" width="20" height="4" rx="2" fill="#4A5568"></rect><path d="M80 64.5h48a.5.5 0 01.5.5v33a.5.5 0 01-.5.5H80a.5.5 0 01-.5-.5V65a.5.5 0 01.5-.5z" stroke="#CBD5E0" fill="none"></path><path d="M109 75a5 5 0 11-10 0 5 5 0 0110 0z" fill="#6366F1"></path><path d="M98 92a1 1 0 011-1h10a1 1 0 010 2H99a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="94" y="84" width="20" height="4" rx="2" fill="#4A5568"></rect><path d="M139 64.5h48a.5.5 0 01.5.5v33a.5.5 0 01-.5.5h-48a.5.5 0 01-.5-.5V65a.5.5 0 01.5-.5z" stroke="#CBD5E0" fill="none"></path><path d="M168 75a5 5 0 11-10 0 5 5 0 0110 0z" fill="#6366F1"></path><path d="M157 92a1 1 0 011-1h10a1 1 0 010 2h-10a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="153" y="84" width="20" height="4" rx="2" fill="#4A5568"></rect><path d="M198 64.5h48a.5.5 0 01.5.5v33a.5.5 0 01-.5.5h-48a.5.5 0 01-.5-.5V65a.5.5 0 01.5-.5z" stroke="#CBD5E0" fill="none"></path><path d="M227 75a5 5 0 11-10 0 5 5 0 0110 0z" fill="#6366F1"></path><path d="M216 92a1 1 0 011-1h10a1 1 0 010 2h-10a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="212" y="84" width="20" height="4" rx="2" fill="#4A5568"></rect><rect x="81" y="36" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="81" y="36" width="104.391" height="4" rx="2" fill="#A0AEC0"></rect><rect x="96" y="25" width="74" height="5" rx="2.5" fill="#4A5568"></rect><rect x="97" y="44" width="73" height="4" rx="2" fill="#A0AEC0"></rect><rect x="113" y="115" width="40" height="10" rx="2" fill="#6366F1"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-col text-center w-full mb-20\"><h1 class=\"sm:text-3xl text-2xl font-medium title-font mb-4 text-gray-900\">Master Cleanse Reliac Heirloom</h1><p class=\"lg:w-2/3 mx-auto leading-relaxed text-base\">Whatever cardigan tote bag tumblr hexagon brooklyn asymmetrical gentrify, subway tile poke farm-to-table. Franzen you probably haven't heard of them man bun deep jianbing selfies heirloom prism food truck ugh squid celiac humblebrag.</p></div><div class=\"flex flex-wrap -m-4 text-center\"><div class=\"p-4 md:w-1/4 sm:w-1/2 w-full\"><div class=\"border-2 border-gray-200 px-4 py-6 rounded-lg\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"text-indigo-500 w-12 h-12 mb-3 inline-block\" viewBox=\"0 0 24 24\"><path d=\"M8 17l4 4 4-4m-4-5v9\"></path><path d=\"M20.88 18.09A5 5 0 0018 9h-1.26A8 8 0 103 16.29\"></path></svg><h2 class=\"title-font font-medium text-3xl text-gray-900\">2.7K</h2><p class=\"leading-relaxed\">Downloads</p></div></div><div class=\"p-4 md:w-1/4 sm:w-1/2 w-full\"><div class=\"border-2 border-gray-200 px-4 py-6 rounded-lg\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"text-indigo-500 w-12 h-12 mb-3 inline-block\" viewBox=\"0 0 24 24\"><path d=\"M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2\"></path><circle cx=\"9\" cy=\"7\" r=\"4\"></circle><path d=\"M23 21v-2a4 4 0 00-3-3.87m-4-12a4 4 0 010 7.75\"></path></svg><h2 class=\"title-font font-medium text-3xl text-gray-900\">1.3K</h2><p class=\"leading-relaxed\">Users</p></div></div><div class=\"p-4 md:w-1/4 sm:w-1/2 w-full\"><div class=\"border-2 border-gray-200 px-4 py-6 rounded-lg\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"text-indigo-500 w-12 h-12 mb-3 inline-block\" viewBox=\"0 0 24 24\"><path d=\"M3 18v-6a9 9 0 0118 0v6\"></path><path d=\"M21 19a2 2 0 01-2 2h-1a2 2 0 01-2-2v-3a2 2 0 012-2h3zM3 19a2 2 0 002 2h1a2 2 0 002-2v-3a2 2 0 00-2-2H3z\"></path></svg><h2 class=\"title-font font-medium text-3xl text-gray-900\">74</h2><p class=\"leading-relaxed\">Files</p></div></div><div class=\"p-4 md:w-1/4 sm:w-1/2 w-full\"><div class=\"border-2 border-gray-200 px-4 py-6 rounded-lg\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"text-indigo-500 w-12 h-12 mb-3 inline-block\" viewBox=\"0 0 24 24\"><path d=\"M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z\"></path></svg><h2 class=\"title-font font-medium text-3xl text-gray-900\">46</h2><p class=\"leading-relaxed\">Places</p></div></div></div></div></section>",category:'Statistics'},{id:'step-block-1',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M24 28.5a.5.5 0 011 0v94a.5.5 0 01-1 0v-94z" fill="#CBD5E0"></path><path d="M29 30.5a4.5 4.5 0 11-9 0 4.5 4.5 0 019 0z" fill="#6366F1"></path><path d="M34 35a1 1 0 011-1h54a1 1 0 110 2H35a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="34" y="28" width="35" height="3" rx="1.5" fill="#4A5568"></rect><path d="M29 52.5a4.5 4.5 0 11-9 0 4.5 4.5 0 019 0z" fill="#6366F1"></path><path d="M34 57a1 1 0 011-1h54a1 1 0 110 2H35a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="34" y="50" width="35" height="3" rx="1.5" fill="#4A5568"></rect><path d="M29 74.5a4.5 4.5 0 11-9 0 4.5 4.5 0 019 0z" fill="#6366F1"></path><path d="M34 79a1 1 0 011-1h54a1 1 0 110 2H35a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="34" y="72" width="35" height="3" rx="1.5" fill="#4A5568"></rect><path d="M29 96.5a4.5 4.5 0 11-9 0 4.5 4.5 0 019 0z" fill="#6366F1"></path><path d="M34 101a1 1 0 011-1h54a1 1 0 010 2H35a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="34" y="94" width="35" height="3" rx="1.5" fill="#4A5568"></rect><path d="M29 118.5a4.5 4.5 0 11-9 0 4.5 4.5 0 019 0z" fill="#6366F1"></path><path d="M34 123a1 1 0 011-1h54a1 1 0 010 2H35a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="34" y="116" width="35" height="3" rx="1.5" fill="#4A5568"></rect><path d="M175.792 89h26.544a3.785 3.785 0 003.792-3.778V58.778A3.785 3.785 0 00202.336 55h-26.544A3.785 3.785 0 00172 58.778v26.444A3.785 3.785 0 00175.792 89zm0 0l20.856-20.778 9.48 9.445m-20.856-12.278a2.838 2.838 0 01-2.844 2.833 2.838 2.838 0 01-2.844-2.833 2.838 2.838 0 012.844-2.833 2.838 2.838 0 012.844 2.833z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path></svg>',content:"\n<section class=\"text-gray-600 body-font\">\n<div class=\"container px-5 py-24 mx-auto flex flex-wrap\">\n<div class=\"flex flex-wrap w-full\">\n<div class=\"lg:w-2/5 md:w-1/2 md:pr-10 md:py-6\">\n<div class=\"flex relative pb-12\">\n<div class=\"h-full w-10 absolute inset-0 flex items-center justify-center\">\n<div class=\"h-full w-1 bg-gray-200 pointer-events-none\"></div>\n</div>\n<div class=\"flex-shrink-0 w-10 h-10 rounded-full bg-indigo-500 inline-flex items-center justify-center text-white relative z-10\">\n<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\">\n<path d=\"M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z\"></path>\n</svg>\n</div>\n<div class=\"flex-grow pl-4\">\n<h2 class=\"font-medium title-font text-sm text-gray-900 mb-1 tracking-wider\">STEP 1</h2>\n<p class=\"leading-relaxed\">VHS cornhole pop-up, try-hard 8-bit iceland helvetica. Kinfolk bespoke try-hard cliche palo santo offal.</p>\n</div>\n</div>\n<div class=\"flex relative pb-12\">\n<div class=\"h-full w-10 absolute inset-0 flex items-center justify-center\">\n<div class=\"h-full w-1 bg-gray-200 pointer-events-none\"></div>\n</div>\n<div class=\"flex-shrink-0 w-10 h-10 rounded-full bg-indigo-500 inline-flex items-center justify-center text-white relative z-10\">\n<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\">\n<path d=\"M22 12h-4l-3 9L9 3l-3 9H2\"></path>\n</svg>\n</div>\n<div class=\"flex-grow pl-4\">\n<h2 class=\"font-medium title-font text-sm text-gray-900 mb-1 tracking-wider\">STEP 2</h2>\n<p class=\"leading-relaxed\">Vice migas literally kitsch +1 pok pok. Truffaut hot chicken slow-carb health goth, vape typewriter.</p>\n</div>\n</div>\n<div class=\"flex relative pb-12\">\n<div class=\"h-full w-10 absolute inset-0 flex items-center justify-center\">\n<div class=\"h-full w-1 bg-gray-200 pointer-events-none\"></div>\n</div>\n<div class=\"flex-shrink-0 w-10 h-10 rounded-full bg-indigo-500 inline-flex items-center justify-center text-white relative z-10\">\n<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\">\n<circle cx=\"12\" cy=\"5\" r=\"3\"></circle>\n<path d=\"M12 22V8M5 12H2a10 10 0 0020 0h-3\"></path>\n</svg>\n</div>\n<div class=\"flex-grow pl-4\">\n<h2 class=\"font-medium title-font text-sm text-gray-900 mb-1 tracking-wider\">STEP 3</h2>\n<p class=\"leading-relaxed\">Coloring book nar whal glossier master cleanse umami. Salvia +1 master cleanse blog taiyaki.</p>\n</div>\n</div>\n<div class=\"flex relative pb-12\">\n<div class=\"h-full w-10 absolute inset-0 flex items-center justify-center\">\n<div class=\"h-full w-1 bg-gray-200 pointer-events-none\"></div>\n</div>\n<div class=\"flex-shrink-0 w-10 h-10 rounded-full bg-indigo-500 inline-flex items-center justify-center text-white relative z-10\">\n<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\">\n<path d=\"M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2\"></path>\n<circle cx=\"12\" cy=\"7\" r=\"4\"></circle>\n</svg>\n</div>\n<div class=\"flex-grow pl-4\">\n<h2 class=\"font-medium title-font text-sm text-gray-900 mb-1 tracking-wider\">STEP 4</h2>\n<p class=\"leading-relaxed\">VHS cornhole pop-up, try-hard 8-bit iceland helvetica. Kinfolk bespoke try-hard cliche palo santo offal.</p>\n</div>\n</div>\n<div class=\"flex relative\">\n<div class=\"flex-shrink-0 w-10 h-10 rounded-full bg-indigo-500 inline-flex items-center justify-center text-white relative z-10\">\n<svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\">\n<path d=\"M22 11.08V12a10 10 0 11-5.93-9.14\"></path>\n<path d=\"M22 4L12 14.01l-3-3\"></path>\n</svg>\n</div>\n<div class=\"flex-grow pl-4\">\n<h2 class=\"font-medium title-font text-sm text-gray-900 mb-1 tracking-wider\">FINISH</h2>\n<p class=\"leading-relaxed\">Pitchfork ugh tattooed scenester echo park gastropub whatever cold-pressed retro.</p>\n</div>\n</div>\n</div>\n<img class=\"lg:w-3/5 md:w-1/2 object-cover object-center rounded-lg md:mt-0 mt-12\" src=\"https://dummyimage.com/1200x500\" alt=\"step\">\n</div>\n</div>\n</section>\n",category:'Steps'},{id:'step-block-3',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M24 28.5a.5.5 0 011 0v93a.5.5 0 01-1 0v-93z" fill="#CBD5E0"></path><path d="M26 40.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0z" fill="#6366F1"></path><path d="M47 43a1 1 0 011-1h54a1 1 0 010 2H48a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="47" y="36" width="35" height="3" rx="1.5" fill="#4A5568"></rect><circle cx="36.5" cy="40.5" r="6.5" fill="#C3DAFE"></circle><path d="M26 63.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0z" fill="#6366F1"></path><path d="M47 66a1 1 0 011-1h54a1 1 0 010 2H48a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="47" y="59" width="35" height="3" rx="1.5" fill="#4A5568"></rect><circle cx="36.5" cy="63.5" r="6.5" fill="#C3DAFE"></circle><path d="M26 86.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0z" fill="#6366F1"></path><path d="M47 89a1 1 0 011-1h54a1 1 0 010 2H48a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="47" y="82" width="35" height="3" rx="1.5" fill="#4A5568"></rect><circle cx="36.5" cy="86.5" r="6.5" fill="#C3DAFE"></circle><path d="M26 109.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0z" fill="#6366F1"></path><path d="M47 112a1 1 0 011-1h54a1 1 0 010 2H48a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="47" y="105" width="35" height="3" rx="1.5" fill="#4A5568"></rect><circle cx="36.5" cy="109.5" r="6.5" fill="#C3DAFE"></circle><path d="M175.792 89h26.544a3.785 3.785 0 003.792-3.778V58.778A3.785 3.785 0 00202.336 55h-26.544A3.785 3.785 0 00172 58.778v26.444A3.785 3.785 0 00175.792 89zm0 0l20.856-20.778 9.48 9.445m-20.856-12.278a2.838 2.838 0 01-2.844 2.833 2.838 2.838 0 01-2.844-2.833 2.838 2.838 0 012.844-2.833 2.838 2.838 0 012.844 2.833z" stroke="#A0AEC0" stroke-width="3px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto flex flex-wrap\"><div class=\"flex relative pt-10 pb-20 sm:items-center md:w-2/3 mx-auto\"><div class=\"h-full w-6 absolute inset-0 flex items-center justify-center\"><div class=\"h-full w-1 bg-gray-200 pointer-events-none\"></div></div><div class=\"flex-shrink-0 w-6 h-6 rounded-full mt-10 sm:mt-0 inline-flex items-center justify-center bg-indigo-500 text-white relative z-10 title-font font-medium text-sm\">1</div><div class=\"flex-grow md:pl-8 pl-6 flex sm:items-center items-start flex-col sm:flex-row\"><div class=\"flex-shrink-0 w-24 h-24 bg-indigo-100 text-indigo-500 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-12 h-12\" viewBox=\"0 0 24 24\"><path d=\"M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z\"></path></svg></div><div class=\"flex-grow sm:pl-6 mt-6 sm:mt-0\"><h2 class=\"font-medium title-font text-gray-900 mb-1 text-xl\">Shooting Stars</h2><p class=\"leading-relaxed\">VHS cornhole pop-up, try-hard 8-bit iceland helvetica. Kinfolk bespoke try-hard cliche palo santo offal.</p></div></div></div><div class=\"flex relative pb-20 sm:items-center md:w-2/3 mx-auto\"><div class=\"h-full w-6 absolute inset-0 flex items-center justify-center\"><div class=\"h-full w-1 bg-gray-200 pointer-events-none\"></div></div><div class=\"flex-shrink-0 w-6 h-6 rounded-full mt-10 sm:mt-0 inline-flex items-center justify-center bg-indigo-500 text-white relative z-10 title-font font-medium text-sm\">2</div><div class=\"flex-grow md:pl-8 pl-6 flex sm:items-center items-start flex-col sm:flex-row\"><div class=\"flex-shrink-0 w-24 h-24 bg-indigo-100 text-indigo-500 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-12 h-12\" viewBox=\"0 0 24 24\"><path d=\"M22 12h-4l-3 9L9 3l-3 9H2\"></path></svg></div><div class=\"flex-grow sm:pl-6 mt-6 sm:mt-0\"><h2 class=\"font-medium title-font text-gray-900 mb-1 text-xl\">The Catalyzer</h2><p class=\"leading-relaxed\">VHS cornhole pop-up, try-hard 8-bit iceland helvetica. Kinfolk bespoke try-hard cliche palo santo offal.</p></div></div></div><div class=\"flex relative pb-20 sm:items-center md:w-2/3 mx-auto\"><div class=\"h-full w-6 absolute inset-0 flex items-center justify-center\"><div class=\"h-full w-1 bg-gray-200 pointer-events-none\"></div></div><div class=\"flex-shrink-0 w-6 h-6 rounded-full mt-10 sm:mt-0 inline-flex items-center justify-center bg-indigo-500 text-white relative z-10 title-font font-medium text-sm\">3</div><div class=\"flex-grow md:pl-8 pl-6 flex sm:items-center items-start flex-col sm:flex-row\"><div class=\"flex-shrink-0 w-24 h-24 bg-indigo-100 text-indigo-500 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-12 h-12\" viewBox=\"0 0 24 24\"><circle cx=\"12\" cy=\"5\" r=\"3\"></circle><path d=\"M12 22V8M5 12H2a10 10 0 0020 0h-3\"></path></svg></div><div class=\"flex-grow sm:pl-6 mt-6 sm:mt-0\"><h2 class=\"font-medium title-font text-gray-900 mb-1 text-xl\">The 400 Blows</h2><p class=\"leading-relaxed\">VHS cornhole pop-up, try-hard 8-bit iceland helvetica. Kinfolk bespoke try-hard cliche palo santo offal.</p></div></div></div><div class=\"flex relative pb-10 sm:items-center md:w-2/3 mx-auto\"><div class=\"h-full w-6 absolute inset-0 flex items-center justify-center\"><div class=\"h-full w-1 bg-gray-200 pointer-events-none\"></div></div><div class=\"flex-shrink-0 w-6 h-6 rounded-full mt-10 sm:mt-0 inline-flex items-center justify-center bg-indigo-500 text-white relative z-10 title-font font-medium text-sm\">4</div><div class=\"flex-grow md:pl-8 pl-6 flex sm:items-center items-start flex-col sm:flex-row\"><div class=\"flex-shrink-0 w-24 h-24 bg-indigo-100 text-indigo-500 rounded-full inline-flex items-center justify-center\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-12 h-12\" viewBox=\"0 0 24 24\"><path d=\"M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2\"></path><circle cx=\"12\" cy=\"7\" r=\"4\"></circle></svg></div><div class=\"flex-grow sm:pl-6 mt-6 sm:mt-0\"><h2 class=\"font-medium title-font text-gray-900 mb-1 text-xl\">Neptune</h2><p class=\"leading-relaxed\">VHS cornhole pop-up, try-hard 8-bit iceland helvetica. Kinfolk bespoke try-hard cliche palo santo offal.</p></div></div></div></div></section>",category:'Steps'},{id:'team-block-1',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="110" y="20" width="46" height="5" rx="2.5" fill="#4A5568"></rect><rect x="87" y="29" width="92" height="4" rx="2" fill="#A0AEC0"></rect><rect x="44" y="52" width="39" height="3" rx="1.5" fill="#4A5568"></rect><rect x="44" y="58" width="20" height="3" rx="1.5" fill="#A0AEC0"></rect><circle cx="31.5" cy="56.5" r="7.5" fill="#E2E8F0"></circle><rect x="20.5" y="45.5" width="69" height="22" rx="1.5" stroke="#CBD5E0" fill="none"></rect><rect x="122" y="52" width="39" height="3" rx="1.5" fill="#4A5568"></rect><rect x="122" y="58" width="20" height="3" rx="1.5" fill="#A0AEC0"></rect><circle cx="109.5" cy="56.5" r="7.5" fill="#E2E8F0"></circle><rect x="98.5" y="45.5" width="69" height="22" rx="1.5" stroke="#CBD5E0" fill="none"></rect><rect x="200" y="52" width="39" height="3" rx="1.5" fill="#4A5568"></rect><rect x="200" y="58" width="20" height="3" rx="1.5" fill="#A0AEC0"></rect><circle cx="187.5" cy="56.5" r="7.5" fill="#E2E8F0"></circle><rect x="176.5" y="45.5" width="69" height="22" rx="1.5" stroke="#CBD5E0" fill="none"></rect><rect x="44" y="83" width="39" height="3" rx="1.5" fill="#4A5568"></rect><rect x="44" y="89" width="20" height="3" rx="1.5" fill="#A0AEC0"></rect><circle cx="31.5" cy="87.5" r="7.5" fill="#E2E8F0"></circle><rect x="20.5" y="76.5" width="69" height="22" rx="1.5" stroke="#CBD5E0" fill="none"></rect><rect x="122" y="83" width="39" height="3" rx="1.5" fill="#4A5568"></rect><rect x="122" y="89" width="20" height="3" rx="1.5" fill="#A0AEC0"></rect><circle cx="109.5" cy="87.5" r="7.5" fill="#E2E8F0"></circle><rect x="98.5" y="76.5" width="69" height="22" rx="1.5" stroke="#CBD5E0" fill="none"></rect><rect x="200" y="83" width="39" height="3" rx="1.5" fill="#4A5568"></rect><rect x="200" y="89" width="20" height="3" rx="1.5" fill="#A0AEC0"></rect><circle cx="187.5" cy="87.5" r="7.5" fill="#E2E8F0"></circle><rect x="176.5" y="76.5" width="69" height="22" rx="1.5" stroke="#CBD5E0" fill="none"></rect><rect x="44" y="114" width="39" height="3" rx="1.5" fill="#4A5568"></rect><rect x="44" y="120" width="20" height="3" rx="1.5" fill="#A0AEC0"></rect><circle cx="31.5" cy="118.5" r="7.5" fill="#E2E8F0"></circle><rect x="20.5" y="107.5" width="69" height="22" rx="1.5" stroke="#CBD5E0" fill="none"></rect><rect x="122" y="114" width="39" height="3" rx="1.5" fill="#4A5568"></rect><rect x="122" y="120" width="20" height="3" rx="1.5" fill="#A0AEC0"></rect><circle cx="109.5" cy="118.5" r="7.5" fill="#E2E8F0"></circle><rect x="98.5" y="107.5" width="69" height="22" rx="1.5" stroke="#CBD5E0" fill="none"></rect><rect x="200" y="114" width="39" height="3" rx="1.5" fill="#4A5568"></rect><rect x="200" y="120" width="20" height="3" rx="1.5" fill="#A0AEC0"></rect><circle cx="187.5" cy="118.5" r="7.5" fill="#E2E8F0"></circle><rect x="176.5" y="107.5" width="69" height="22" rx="1.5" stroke="#CBD5E0" fill="none"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-col text-center w-full mb-20\"><h1 class=\"sm:text-3xl text-2xl font-medium title-font mb-4 text-gray-900\">Our Team</h1><p class=\"lg:w-2/3 mx-auto leading-relaxed text-base\">Whatever cardigan tote bag tumblr hexagon brooklyn asymmetrical gentrify, subway tile poke farm-to-table. Franzen you probably haven't heard of them.</p></div><div class=\"flex flex-wrap -m-2\"><div class=\"p-2 lg:w-1/3 md:w-1/2 w-full\"><div class=\"h-full flex items-center border-gray-200 border p-4 rounded-lg\"><img alt=\"team\" class=\"w-16 h-16 bg-gray-100 object-cover object-center flex-shrink-0 rounded-full mr-4\" src=\"https://dummyimage.com/80x80\"><div class=\"flex-grow\"><h2 class=\"text-gray-900 title-font font-medium\">Holden Caulfield</h2><p class=\"text-gray-500\">UI Designer</p></div></div></div><div class=\"p-2 lg:w-1/3 md:w-1/2 w-full\"><div class=\"h-full flex items-center border-gray-200 border p-4 rounded-lg\"><img alt=\"team\" class=\"w-16 h-16 bg-gray-100 object-cover object-center flex-shrink-0 rounded-full mr-4\" src=\"https://dummyimage.com/84x84\"><div class=\"flex-grow\"><h2 class=\"text-gray-900 title-font font-medium\">Henry Letham</h2><p class=\"text-gray-500\">CTO</p></div></div></div><div class=\"p-2 lg:w-1/3 md:w-1/2 w-full\"><div class=\"h-full flex items-center border-gray-200 border p-4 rounded-lg\"><img alt=\"team\" class=\"w-16 h-16 bg-gray-100 object-cover object-center flex-shrink-0 rounded-full mr-4\" src=\"https://dummyimage.com/88x88\"><div class=\"flex-grow\"><h2 class=\"text-gray-900 title-font font-medium\">Oskar Blinde</h2><p class=\"text-gray-500\">Founder</p></div></div></div><div class=\"p-2 lg:w-1/3 md:w-1/2 w-full\"><div class=\"h-full flex items-center border-gray-200 border p-4 rounded-lg\"><img alt=\"team\" class=\"w-16 h-16 bg-gray-100 object-cover object-center flex-shrink-0 rounded-full mr-4\" src=\"https://dummyimage.com/90x90\"><div class=\"flex-grow\"><h2 class=\"text-gray-900 title-font font-medium\">John Doe</h2><p class=\"text-gray-500\">DevOps</p></div></div></div><div class=\"p-2 lg:w-1/3 md:w-1/2 w-full\"><div class=\"h-full flex items-center border-gray-200 border p-4 rounded-lg\"><img alt=\"team\" class=\"w-16 h-16 bg-gray-100 object-cover object-center flex-shrink-0 rounded-full mr-4\" src=\"https://dummyimage.com/94x94\"><div class=\"flex-grow\"><h2 class=\"text-gray-900 title-font font-medium\">Martin Eden</h2><p class=\"text-gray-500\">Software Engineer</p></div></div></div><div class=\"p-2 lg:w-1/3 md:w-1/2 w-full\"><div class=\"h-full flex items-center border-gray-200 border p-4 rounded-lg\"><img alt=\"team\" class=\"w-16 h-16 bg-gray-100 object-cover object-center flex-shrink-0 rounded-full mr-4\" src=\"https://dummyimage.com/98x98\"><div class=\"flex-grow\"><h2 class=\"text-gray-900 title-font font-medium\">Boris Kitua</h2><p class=\"text-gray-500\">UX Researcher</p></div></div></div><div class=\"p-2 lg:w-1/3 md:w-1/2 w-full\"><div class=\"h-full flex items-center border-gray-200 border p-4 rounded-lg\"><img alt=\"team\" class=\"w-16 h-16 bg-gray-100 object-cover object-center flex-shrink-0 rounded-full mr-4\" src=\"https://dummyimage.com/100x90\"><div class=\"flex-grow\"><h2 class=\"text-gray-900 title-font font-medium\">Atticus Finch</h2><p class=\"text-gray-500\">QA Engineer</p></div></div></div><div class=\"p-2 lg:w-1/3 md:w-1/2 w-full\"><div class=\"h-full flex items-center border-gray-200 border p-4 rounded-lg\"><img alt=\"team\" class=\"w-16 h-16 bg-gray-100 object-cover object-center flex-shrink-0 rounded-full mr-4\" src=\"https://dummyimage.com/104x94\"><div class=\"flex-grow\"><h2 class=\"text-gray-900 title-font font-medium\">Alper Kamu</h2><p class=\"text-gray-500\">System</p></div></div></div><div class=\"p-2 lg:w-1/3 md:w-1/2 w-full\"><div class=\"h-full flex items-center border-gray-200 border p-4 rounded-lg\"><img alt=\"team\" class=\"w-16 h-16 bg-gray-100 object-cover object-center flex-shrink-0 rounded-full mr-4\" src=\"https://dummyimage.com/108x98\"><div class=\"flex-grow\"><h2 class=\"text-gray-900 title-font font-medium\">Rodrigo Monchi</h2><p class=\"text-gray-500\">Product Manager</p></div></div></div></div></div></section>",category:'Team'},{id:'team-block-2',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="110" y="28" width="46" height="5" rx="2.5" fill="#4A5568"></rect><rect x="87" y="37" width="92" height="4" rx="2" fill="#A0AEC0"></rect><path d="M36.111 83H50.89A2.111 2.111 0 0053 80.889V66.11A2.111 2.111 0 0050.889 64H36.11A2.111 2.111 0 0034 66.111V80.89a2.11 2.11 0 002.111 2.111zm0 0l11.611-11.611L53 76.667m-11.611-6.861a1.583 1.583 0 11-3.167 0 1.583 1.583 0 013.167 0z" stroke="#A0AEC0" stroke-width="1.7px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="66" y="65" width="39" height="3" rx="1.5" fill="#4A5568"></rect><path d="M66 73.5a1.5 1.5 0 011.5-1.5h48a1.5 1.5 0 010 3h-48a1.5 1.5 0 01-1.5-1.5zm0 7a1.5 1.5 0 011.5-1.5h53a1.5 1.5 0 010 3h-53a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M147.111 83h14.778A2.111 2.111 0 00164 80.889V66.11a2.111 2.111 0 00-2.111-2.11h-14.778A2.111 2.111 0 00145 66.111V80.89a2.11 2.11 0 002.111 2.111zm0 0l11.611-11.611L164 76.667m-11.611-6.861a1.583 1.583 0 11-3.167 0 1.583 1.583 0 013.167 0z" stroke="#A0AEC0" stroke-width="1.7px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="177" y="65" width="39" height="3" rx="1.5" fill="#4A5568"></rect><path d="M177 73.5a1.5 1.5 0 011.5-1.5h48a1.5 1.5 0 010 3h-48a1.5 1.5 0 01-1.5-1.5zm0 7a1.5 1.5 0 011.5-1.5h53a1.5 1.5 0 010 3h-53a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M36.111 121H50.89a2.111 2.111 0 002.11-2.111v-14.778A2.111 2.111 0 0050.889 102H36.11a2.111 2.111 0 00-2.11 2.111v14.778A2.11 2.11 0 0036.111 121zm0 0l11.611-11.611L53 114.667m-11.611-6.861a1.583 1.583 0 11-3.167 0 1.583 1.583 0 013.167 0z" stroke="#A0AEC0" stroke-width="1.7px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="66" y="103" width="39" height="3" rx="1.5" fill="#4A5568"></rect><path d="M66 111.5a1.5 1.5 0 011.5-1.5h48a1.5 1.5 0 010 3h-48a1.5 1.5 0 01-1.5-1.5zm0 7a1.5 1.5 0 011.5-1.5h53a1.5 1.5 0 010 3h-53a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path><path d="M147.111 121h14.778a2.11 2.11 0 002.111-2.111v-14.778a2.11 2.11 0 00-2.111-2.111h-14.778a2.11 2.11 0 00-2.111 2.111v14.778a2.11 2.11 0 002.111 2.111zm0 0l11.611-11.611 5.278 5.278m-11.611-6.861a1.583 1.583 0 11-3.167 0 1.583 1.583 0 013.167 0z" stroke="#A0AEC0" stroke-width="1.7px" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="177" y="103" width="39" height="3" rx="1.5" fill="#4A5568"></rect><path d="M177 111.5a1.5 1.5 0 011.5-1.5h48a1.5 1.5 0 010 3h-48a1.5 1.5 0 01-1.5-1.5zm0 7a1.5 1.5 0 011.5-1.5h53a1.5 1.5 0 010 3h-53a1.5 1.5 0 01-1.5-1.5z" fill="#A0AEC0"></path></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-col text-center w-full mb-20\"><h1 class=\"text-2xl font-medium title-font mb-4 text-gray-900 tracking-widest\">OUR TEAM</h1><p class=\"lg:w-2/3 mx-auto leading-relaxed text-base\">Whatever cardigan tote bag tumblr hexagon brooklyn asymmetrical gentrify, subway tile poke farm-to-table. Franzen you probably haven't heard of them.</p></div><div class=\"flex flex-wrap -m-4\"><div class=\"p-4 lg:w-1/2\"><div class=\"h-full flex sm:flex-row flex-col items-center sm:justify-start justify-center text-center sm:text-left\"><img alt=\"team\" class=\"flex-shrink-0 rounded-lg w-48 h-48 object-cover object-center sm:mb-0 mb-4\" src=\"https://dummyimage.com/200x200\"><div class=\"flex-grow sm:pl-8\"><h2 class=\"title-font font-medium text-lg text-gray-900\">Holden Caulfield</h2><h3 class=\"text-gray-500 mb-3\">UI Developer</h3><p class=\"mb-4\">DIY tote bag drinking vinegar cronut adaptogen squid fanny pack vaporware.</p><span class=\"inline-flex\"><a class=\"text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg></a></span></div></div></div><div class=\"p-4 lg:w-1/2\"><div class=\"h-full flex sm:flex-row flex-col items-center sm:justify-start justify-center text-center sm:text-left\"><img alt=\"team\" class=\"flex-shrink-0 rounded-lg w-48 h-48 object-cover object-center sm:mb-0 mb-4\" src=\"https://dummyimage.com/201x201\"><div class=\"flex-grow sm:pl-8\"><h2 class=\"title-font font-medium text-lg text-gray-900\">Alper Kamu</h2><h3 class=\"text-gray-500 mb-3\">Designer</h3><p class=\"mb-4\">DIY tote bag drinking vinegar cronut adaptogen squid fanny pack vaporware.</p><span class=\"inline-flex\"><a class=\"text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg></a></span></div></div></div><div class=\"p-4 lg:w-1/2\"><div class=\"h-full flex sm:flex-row flex-col items-center sm:justify-start justify-center text-center sm:text-left\"><img alt=\"team\" class=\"flex-shrink-0 rounded-lg w-48 h-48 object-cover object-center sm:mb-0 mb-4\" src=\"https://dummyimage.com/204x204\"><div class=\"flex-grow sm:pl-8\"><h2 class=\"title-font font-medium text-lg text-gray-900\">Atticus Finch</h2><h3 class=\"text-gray-500 mb-3\">UI Developer</h3><p class=\"mb-4\">DIY tote bag drinking vinegar cronut adaptogen squid fanny pack vaporware.</p><span class=\"inline-flex\"><a class=\"text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg></a></span></div></div></div><div class=\"p-4 lg:w-1/2\"><div class=\"h-full flex sm:flex-row flex-col items-center sm:justify-start justify-center text-center sm:text-left\"><img alt=\"team\" class=\"flex-shrink-0 rounded-lg w-48 h-48 object-cover object-center sm:mb-0 mb-4\" src=\"https://dummyimage.com/206x206\"><div class=\"flex-grow sm:pl-8\"><h2 class=\"title-font font-medium text-lg text-gray-900\">Henry Letham</h2><h3 class=\"text-gray-500 mb-3\">Designer</h3><p class=\"mb-4\">DIY tote bag drinking vinegar cronut adaptogen squid fanny pack vaporware.</p><span class=\"inline-flex\"><a class=\"text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg></a></span></div></div></div></div></div></section>",category:'Team'},{id:'team-block-3',class:'',label:'<svg fill="none" viewBox="0 0 266 150" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="110" y="20" width="46" height="5" rx="2.5" fill="#4A5568"></rect><rect x="87" y="29" width="92" height="4" rx="2" fill="#A0AEC0"></rect><path d="M39.444 66h10.112c.797 0 1.444-.647 1.444-1.444V54.444c0-.797-.647-1.444-1.444-1.444H39.444c-.797 0-1.444.647-1.444 1.444v10.112c0 .797.647 1.444 1.444 1.444zm0 0l7.945-7.944L51 61.666m-7.944-4.694a1.083 1.083 0 11-2.167 0 1.083 1.083 0 012.167 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="24" y="78" width="41" height="2" rx="1" fill="#A0AEC0"></rect><rect x="29" y="82" width="32" height="2" rx="1" fill="#A0AEC0"></rect><rect x="36" y="74" width="17" height="2" rx="1" fill="#4A5568"></rect><path d="M98.444 66h10.112c.797 0 1.444-.647 1.444-1.444V54.444c0-.797-.647-1.444-1.444-1.444H98.444c-.797 0-1.444.647-1.444 1.444v10.112c0 .797.647 1.444 1.444 1.444zm0 0l7.945-7.944 3.611 3.61m-7.944-4.694a1.084 1.084 0 11-2.167 0 1.084 1.084 0 012.167 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="83" y="78" width="41" height="2" rx="1" fill="#A0AEC0"></rect><rect x="88" y="82" width="32" height="2" rx="1" fill="#A0AEC0"></rect><rect x="95" y="74" width="17" height="2" rx="1" fill="#4A5568"></rect><path d="M157.444 66h10.112c.797 0 1.444-.647 1.444-1.444V54.444c0-.797-.647-1.444-1.444-1.444h-10.112c-.797 0-1.444.647-1.444 1.444v10.112c0 .797.647 1.444 1.444 1.444zm0 0l7.945-7.944 3.611 3.61m-7.944-4.694a1.084 1.084 0 11-2.167 0 1.084 1.084 0 012.167 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="142" y="78" width="41" height="2" rx="1" fill="#A0AEC0"></rect><rect x="147" y="82" width="32" height="2" rx="1" fill="#A0AEC0"></rect><rect x="154" y="74" width="17" height="2" rx="1" fill="#4A5568"></rect><path d="M216.444 66h10.112c.797 0 1.444-.647 1.444-1.444V54.444c0-.797-.647-1.444-1.444-1.444h-10.112c-.797 0-1.444.647-1.444 1.444v10.112c0 .797.647 1.444 1.444 1.444zm0 0l7.945-7.944 3.611 3.61m-7.944-4.694a1.084 1.084 0 11-2.167 0 1.084 1.084 0 012.167 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="201" y="78" width="41" height="2" rx="1" fill="#A0AEC0"></rect><rect x="206" y="82" width="32" height="2" rx="1" fill="#A0AEC0"></rect><rect x="213" y="74" width="17" height="2" rx="1" fill="#4A5568"></rect><path d="M39.444 112h10.112c.797 0 1.444-.647 1.444-1.444v-10.112c0-.797-.647-1.444-1.444-1.444H39.444c-.797 0-1.444.647-1.444 1.444v10.112c0 .797.647 1.444 1.444 1.444zm0 0l7.945-7.944L51 107.667m-7.944-4.695a1.084 1.084 0 11-2.167.001 1.084 1.084 0 012.167-.001z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="24" y="124" width="41" height="2" rx="1" fill="#A0AEC0"></rect><rect x="29" y="128" width="32" height="2" rx="1" fill="#A0AEC0"></rect><rect x="36" y="120" width="17" height="2" rx="1" fill="#4A5568"></rect><path d="M98.444 112h10.112c.797 0 1.444-.647 1.444-1.444v-10.112c0-.797-.647-1.444-1.444-1.444H98.444c-.797 0-1.444.647-1.444 1.444v10.112c0 .797.647 1.444 1.444 1.444zm0 0l7.945-7.944 3.611 3.611m-7.944-4.695a1.084 1.084 0 11-2.167 0 1.084 1.084 0 012.167 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="83" y="124" width="41" height="2" rx="1" fill="#A0AEC0"></rect><rect x="88" y="128" width="32" height="2" rx="1" fill="#A0AEC0"></rect><rect x="95" y="120" width="17" height="2" rx="1" fill="#4A5568"></rect><path d="M157.444 112h10.112c.797 0 1.444-.647 1.444-1.444v-10.112c0-.797-.647-1.444-1.444-1.444h-10.112c-.797 0-1.444.647-1.444 1.444v10.112c0 .797.647 1.444 1.444 1.444zm0 0l7.945-7.944 3.611 3.611m-7.944-4.695a1.084 1.084 0 11-2.167 0 1.084 1.084 0 012.167 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="142" y="124" width="41" height="2" rx="1" fill="#A0AEC0"></rect><rect x="147" y="128" width="32" height="2" rx="1" fill="#A0AEC0"></rect><rect x="154" y="120" width="17" height="2" rx="1" fill="#4A5568"></rect><path d="M216.444 112h10.112c.797 0 1.444-.647 1.444-1.444v-10.112c0-.797-.647-1.444-1.444-1.444h-10.112c-.797 0-1.444.647-1.444 1.444v10.112c0 .797.647 1.444 1.444 1.444zm0 0l7.945-7.944 3.611 3.611m-7.944-4.695a1.084 1.084 0 11-2.167 0 1.084 1.084 0 012.167 0z" stroke="#A0AEC0" stroke-linecap="round" stroke-linejoin="round" fill="none"></path><rect x="201" y="124" width="41" height="2" rx="1" fill="#A0AEC0"></rect><rect x="206" y="128" width="32" height="2" rx="1" fill="#A0AEC0"></rect><rect x="213" y="120" width="17" height="2" rx="1" fill="#4A5568"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-col text-center w-full mb-20\"><h1 class=\"text-2xl font-medium title-font mb-4 text-gray-900\">OUR TEAM</h1><p class=\"lg:w-2/3 mx-auto leading-relaxed text-base\">Whatever cardigan tote bag tumblr hexagon brooklyn asymmetrical gentrify, subway tile poke farm-to-table. Franzen you probably haven't heard of them.</p></div><div class=\"flex flex-wrap -m-4\"><div class=\"p-4 lg:w-1/4 md:w-1/2\"><div class=\"h-full flex flex-col items-center text-center\"><img alt=\"team\" class=\"flex-shrink-0 rounded-lg w-full h-56 object-cover object-center mb-4\" src=\"https://dummyimage.com/200x200\"><div class=\"w-full\"><h2 class=\"title-font font-medium text-lg text-gray-900\">Alper Kamu</h2><h3 class=\"text-gray-500 mb-3\">UI Developer</h3><p class=\"mb-4\">DIY tote bag drinking vinegar cronut adaptogen squid fanny pack vaporware.</p><span class=\"inline-flex\"><a class=\"text-gray-500\"><svg fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg></a></span></div></div></div><div class=\"p-4 lg:w-1/4 md:w-1/2\"><div class=\"h-full flex flex-col items-center text-center\"><img alt=\"team\" class=\"flex-shrink-0 rounded-lg w-full h-56 object-cover object-center mb-4\" src=\"https://dummyimage.com/201x201\"><div class=\"w-full\"><h2 class=\"title-font font-medium text-lg text-gray-900\">Holden Caulfield</h2><h3 class=\"text-gray-500 mb-3\">UI Developer</h3><p class=\"mb-4\">DIY tote bag drinking vinegar cronut adaptogen squid fanny pack vaporware.</p><span class=\"inline-flex\"><a class=\"text-gray-500\"><svg fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg></a></span></div></div></div><div class=\"p-4 lg:w-1/4 md:w-1/2\"><div class=\"h-full flex flex-col items-center text-center\"><img alt=\"team\" class=\"flex-shrink-0 rounded-lg w-full h-56 object-cover object-center mb-4\" src=\"https://dummyimage.com/202x202\"><div class=\"w-full\"><h2 class=\"title-font font-medium text-lg text-gray-900\">Atticus Finch</h2><h3 class=\"text-gray-500 mb-3\">UI Developer</h3><p class=\"mb-4\">DIY tote bag drinking vinegar cronut adaptogen squid fanny pack vaporware.</p><span class=\"inline-flex\"><a class=\"text-gray-500\"><svg fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg></a></span></div></div></div><div class=\"p-4 lg:w-1/4 md:w-1/2\"><div class=\"h-full flex flex-col items-center text-center\"><img alt=\"team\" class=\"flex-shrink-0 rounded-lg w-full h-56 object-cover object-center mb-4\" src=\"https://dummyimage.com/203x203\"><div class=\"w-full\"><h2 class=\"title-font font-medium text-lg text-gray-900\">Henry Letham</h2><h3 class=\"text-gray-500 mb-3\">UI Developer</h3><p class=\"mb-4\">DIY tote bag drinking vinegar cronut adaptogen squid fanny pack vaporware.</p><span class=\"inline-flex\"><a class=\"text-gray-500\"><svg fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M18 2h-3a5 5 0 00-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 011-1h3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M23 3a10.9 10.9 0 01-3.14 1.53 4.48 4.48 0 00-7.86 3v1A10.66 10.66 0 013 4s-4 9 5 13a11.64 11.64 0 01-7 2c9 5 20 0 20-11.5a4.5 4.5 0 00-.08-.83A7.72 7.72 0 0023 3z\"></path></svg></a><a class=\"ml-2 text-gray-500\"><svg fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" class=\"w-5 h-5\" viewBox=\"0 0 24 24\"><path d=\"M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z\"></path></svg></a></span></div></div></div></div></div></section>",category:'Team'},{id:'testimonial-block-1',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><rect x="25" y="57" width="104" height="62" rx="1" fill="#E2E8F0"></rect><circle cx="39.5" cy="105.5" r="7.5" fill="#CBD5E0"></circle><path d="M52 108a1 1 0 011-1h24a1 1 0 010 2H53a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="52" y="101" width="43" height="3" rx="1.5" fill="#4A5568"></rect><rect x="32" y="75" width="76" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="32" y="81" width="88" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="32" y="87" width="83" height="3" rx="1.5" fill="#A0AEC0"></rect><path d="M38.641 63h-2.182a.354.354 0 00-.36.349v2.119c0 .192.161.349.36.349h1.044c-.014.554-.146.999-.398 1.333-.198.263-.498.481-.9.653a.344.344 0 00-.177.468l.258.53c.084.17.29.245.468.17.475-.2.876-.452 1.204-.758.399-.375.672-.797.82-1.268.148-.472.222-1.115.222-1.93v-1.666a.354.354 0 00-.359-.349zM32.761 68.97c.47-.199.869-.451 1.198-.757.403-.375.678-.796.826-1.264.148-.467.222-1.112.222-1.934v-1.666a.354.354 0 00-.359-.349h-2.183a.354.354 0 00-.359.349v2.119c0 .192.161.349.36.349h1.044c-.014.554-.146.999-.398 1.333-.198.263-.498.481-.9.653a.344.344 0 00-.177.468l.258.529c.083.17.29.245.468.17z" fill="#A0AEC0"></path><rect x="137" y="57" width="104" height="62" rx="1" fill="#E2E8F0"></rect><circle cx="151.5" cy="105.5" r="7.5" fill="#CBD5E0"></circle><path d="M164 108a1 1 0 011-1h24a1 1 0 010 2h-24a1 1 0 01-1-1z" fill="#A0AEC0"></path><rect x="164" y="101" width="43" height="3" rx="1.5" fill="#4A5568"></rect><rect x="144" y="75" width="76" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="144" y="81" width="88" height="3" rx="1.5" fill="#A0AEC0"></rect><rect x="144" y="87" width="83" height="3" rx="1.5" fill="#A0AEC0"></rect><path d="M150.641 63h-2.182a.354.354 0 00-.359.349v2.119c0 .192.16.349.359.349h1.044c-.014.554-.146.999-.398 1.333-.198.263-.498.481-.899.653a.344.344 0 00-.178.468l.258.53c.084.17.29.245.468.17.475-.2.876-.452 1.204-.758.399-.375.672-.797.82-1.268.148-.472.222-1.115.222-1.93v-1.666a.354.354 0 00-.359-.349zM144.761 68.97c.47-.199.869-.451 1.198-.757.403-.375.678-.796.826-1.264.148-.467.222-1.112.222-1.934v-1.666a.354.354 0 00-.359-.349h-2.183a.353.353 0 00-.358.349v2.119c0 .192.16.349.358.349h1.045c-.014.554-.146.999-.398 1.333-.198.263-.498.481-.899.653a.344.344 0 00-.178.468l.257.529c.084.17.291.245.469.17z" fill="#A0AEC0"></path><rect x="107" y="31" width="52" height="5" rx="2.5" fill="#4A5568"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><h1 class=\"text-3xl font-medium title-font text-gray-900 mb-12 text-center\">Testimonials</h1><div class=\"flex flex-wrap -m-4\"><div class=\"p-4 md:w-1/2 w-full\"><div class=\"h-full bg-gray-100 p-8 rounded\"><svg fill=\"currentColor\" class=\"block w-5 h-5 text-gray-400 mb-4\" viewBox=\"0 0 975.036 975.036\"><path d=\"M925.036 57.197h-304c-27.6 0-50 22.4-50 50v304c0 27.601 22.4 50 50 50h145.5c-1.9 79.601-20.4 143.3-55.4 191.2-27.6 37.8-69.399 69.1-125.3 93.8-25.7 11.3-36.8 41.7-24.8 67.101l36 76c11.6 24.399 40.3 35.1 65.1 24.399 66.2-28.6 122.101-64.8 167.7-108.8 55.601-53.7 93.7-114.3 114.3-181.9 20.601-67.6 30.9-159.8 30.9-276.8v-239c0-27.599-22.401-50-50-50zM106.036 913.497c65.4-28.5 121-64.699 166.9-108.6 56.1-53.7 94.4-114.1 115-181.2 20.6-67.1 30.899-159.6 30.899-277.5v-239c0-27.6-22.399-50-50-50h-304c-27.6 0-50 22.4-50 50v304c0 27.601 22.4 50 50 50h145.5c-1.9 79.601-20.4 143.3-55.4 191.2-27.6 37.8-69.4 69.1-125.3 93.8-25.7 11.3-36.8 41.7-24.8 67.101l35.9 75.8c11.601 24.399 40.501 35.2 65.301 24.399z\"></path></svg><p class=\"leading-relaxed mb-6\">Synth chartreuse iPhone lomo cray raw denim brunch everyday carry neutra before they sold out fixie 90's microdosing. Tacos pinterest fanny pack venmo, post-ironic heirloom try-hard pabst authentic iceland.</p><a class=\"inline-flex items-center\"><img alt=\"testimonial\" src=\"https://dummyimage.com/106x106\" class=\"w-12 h-12 rounded-full flex-shrink-0 object-cover object-center\"><span class=\"flex-grow flex flex-col pl-4\"><span class=\"title-font font-medium text-gray-900\">Holden Caulfield</span><span class=\"text-gray-500 text-sm\">UI DEVELOPER</span></span></a></div></div><div class=\"p-4 md:w-1/2 w-full\"><div class=\"h-full bg-gray-100 p-8 rounded\"><svg fill=\"currentColor\" class=\"block w-5 h-5 text-gray-400 mb-4\" viewBox=\"0 0 975.036 975.036\"><path d=\"M925.036 57.197h-304c-27.6 0-50 22.4-50 50v304c0 27.601 22.4 50 50 50h145.5c-1.9 79.601-20.4 143.3-55.4 191.2-27.6 37.8-69.399 69.1-125.3 93.8-25.7 11.3-36.8 41.7-24.8 67.101l36 76c11.6 24.399 40.3 35.1 65.1 24.399 66.2-28.6 122.101-64.8 167.7-108.8 55.601-53.7 93.7-114.3 114.3-181.9 20.601-67.6 30.9-159.8 30.9-276.8v-239c0-27.599-22.401-50-50-50zM106.036 913.497c65.4-28.5 121-64.699 166.9-108.6 56.1-53.7 94.4-114.1 115-181.2 20.6-67.1 30.899-159.6 30.899-277.5v-239c0-27.6-22.399-50-50-50h-304c-27.6 0-50 22.4-50 50v304c0 27.601 22.4 50 50 50h145.5c-1.9 79.601-20.4 143.3-55.4 191.2-27.6 37.8-69.4 69.1-125.3 93.8-25.7 11.3-36.8 41.7-24.8 67.101l35.9 75.8c11.601 24.399 40.501 35.2 65.301 24.399z\"></path></svg><p class=\"leading-relaxed mb-6\">Synth chartreuse iPhone lomo cray raw denim brunch everyday carry neutra before they sold out fixie 90's microdosing. Tacos pinterest fanny pack venmo, post-ironic heirloom try-hard pabst authentic iceland.</p><a class=\"inline-flex items-center\"><img alt=\"testimonial\" src=\"https://dummyimage.com/107x107\" class=\"w-12 h-12 rounded-full flex-shrink-0 object-cover object-center\"><span class=\"flex-grow flex flex-col pl-4\"><span class=\"title-font font-medium text-gray-900\">Alper Kamu</span><span class=\"text-gray-500 text-sm\">DESIGNER</span></span></a></div></div></div></div></section>",category:'Testimonials'},{id:'testimonial-block-2',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M139.282 32h-4.365a.708.708 0 00-.718.697v4.239c0 .385.322.697.718.697h2.089c-.027 1.11-.293 1.998-.795 2.666-.396.527-.997.963-1.799 1.308a.687.687 0 00-.356.935l.517 1.06c.166.34.578.49.934.34.951-.398 1.753-.903 2.408-1.517.798-.748 1.346-1.593 1.641-2.536.296-.943.444-2.228.444-3.86v-3.332a.708.708 0 00-.718-.697zM127.523 43.94c.939-.398 1.737-.903 2.396-1.515.805-.748 1.355-1.59 1.651-2.526.296-.936.444-2.226.444-3.87v-3.332a.708.708 0 00-.718-.697h-4.365a.708.708 0 00-.718.697v4.239c0 .385.322.697.718.697h2.089c-.027 1.11-.293 1.998-.795 2.666-.397.527-.997.963-1.799 1.308a.689.689 0 00-.357.935l.516 1.057c.166.34.581.491.938.34z" fill="#A0AEC0"></path><rect x="95" y="58" width="76" height="4" rx="2" fill="#A0AEC0"></rect><rect x="123" y="94" width="20" height="2" rx="1" fill="#6366F1"></rect><rect x="89" y="66" width="88" height="4" rx="2" fill="#A0AEC0"></rect><rect x="92" y="74" width="83" height="4" rx="2" fill="#A0AEC0"></rect><rect x="103" y="82" width="60" height="4" rx="2" fill="#A0AEC0"></rect><rect x="113" y="104" width="40" height="4" rx="2" fill="#4A5568"></rect><rect x="106" y="112" width="54" height="4" rx="2" fill="#CBD5E0"></rect></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"xl:w-1/2 lg:w-3/4 w-full mx-auto text-center\"><svg fill=\"currentColor\" class=\"inline-block w-8 h-8 text-gray-400 mb-8\" viewBox=\"0 0 975.036 975.036\"><path d=\"M925.036 57.197h-304c-27.6 0-50 22.4-50 50v304c0 27.601 22.4 50 50 50h145.5c-1.9 79.601-20.4 143.3-55.4 191.2-27.6 37.8-69.399 69.1-125.3 93.8-25.7 11.3-36.8 41.7-24.8 67.101l36 76c11.6 24.399 40.3 35.1 65.1 24.399 66.2-28.6 122.101-64.8 167.7-108.8 55.601-53.7 93.7-114.3 114.3-181.9 20.601-67.6 30.9-159.8 30.9-276.8v-239c0-27.599-22.401-50-50-50zM106.036 913.497c65.4-28.5 121-64.699 166.9-108.6 56.1-53.7 94.4-114.1 115-181.2 20.6-67.1 30.899-159.6 30.899-277.5v-239c0-27.6-22.399-50-50-50h-304c-27.6 0-50 22.4-50 50v304c0 27.601 22.4 50 50 50h145.5c-1.9 79.601-20.4 143.3-55.4 191.2-27.6 37.8-69.4 69.1-125.3 93.8-25.7 11.3-36.8 41.7-24.8 67.101l35.9 75.8c11.601 24.399 40.501 35.2 65.301 24.399z\"></path></svg><p class=\"leading-relaxed text-lg\">Edison bulb retro cloud bread echo park, helvetica stumptown taiyaki taxidermy 90's cronut +1 kinfolk. Single-origin coffee ennui shaman taiyaki vape DIY tote bag drinking vinegar cronut adaptogen squid fanny pack vaporware. Man bun next level coloring book skateboard four loko knausgaard. Kitsch keffiyeh master cleanse direct trade indigo juice before they sold out gentrify plaid gastropub normcore XOXO 90's pickled cindigo jean shorts. Slow-carb next level shoindigoitch ethical authentic, yr scenester sriracha forage franzen organic drinking vinegar.</p><span class=\"inline-block h-1 w-10 rounded bg-indigo-500 mt-8 mb-6\"></span><h2 class=\"text-gray-900 font-medium title-font tracking-wider text-sm\">HOLDEN CAULFIELD</h2><p class=\"text-gray-500\">Senior Product Designer</p></div></div></section>",category:'Testimonials'},{id:'testimonial-block-3',class:'',label:'<svg viewBox="0 0 266 150" fill="none" width="266"  height="150" ><path fill="#FFFFFF" d="M0 0h266v150H0z"></path><path d="M21 77a2 2 0 012-2h59a2 2 0 110 4H23a2 2 0 01-2-2zM26 85a2 2 0 012-2h48.92a2 2 0 110 4H28a2 2 0 01-2-2z" fill="#A0AEC0"></path><path d="M38 104a2 2 0 012-2h25a2 2 0 110 4H40a2 2 0 01-2-2z" fill="#4A5568"></path><path d="M26 69a2 2 0 012-2h48.92a2 2 0 110 4H28a2 2 0 01-2-2z" fill="#A0AEC0"></path><path d="M44 94.5a1.5 1.5 0 011.5-1.5h13.38a1.5 1.5 0 010 3H45.5a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><circle cx="53" cy="53" r="8" fill="#E2E8F0"></circle><path d="M102 77a2 2 0 012-2h59a2 2 0 110 4h-59a2 2 0 01-2-2zM107 85a2 2 0 012-2h48.92a2 2 0 110 4H109a2 2 0 01-2-2z" fill="#A0AEC0"></path><path d="M119 104a2 2 0 012-2h25a2 2 0 110 4h-25a2 2 0 01-2-2z" fill="#4A5568"></path><path d="M107 69a2 2 0 012-2h48.92a2 2 0 110 4H109a2 2 0 01-2-2z" fill="#A0AEC0"></path><path d="M125 94.5a1.5 1.5 0 011.5-1.5h13.38a1.5 1.5 0 010 3H126.5a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><circle cx="134" cy="53" r="8" fill="#E2E8F0"></circle><path d="M183 77a2 2 0 012-2h59a2 2 0 110 4h-59a2 2 0 01-2-2zM188 85a2 2 0 012-2h48.92a2 2 0 110 4H190a2 2 0 01-2-2z" fill="#A0AEC0"></path><path d="M200 104a2 2 0 012-2h25a2 2 0 110 4h-25a2 2 0 01-2-2z" fill="#4A5568"></path><path d="M188 69a2 2 0 012-2h48.92a2 2 0 110 4H190a2 2 0 01-2-2z" fill="#A0AEC0"></path><path d="M206 94.5a1.5 1.5 0 011.5-1.5h13.38a1.5 1.5 0 010 3H207.5a1.5 1.5 0 01-1.5-1.5z" fill="#6366F1"></path><circle cx="215" cy="53" r="8" fill="#E2E8F0"></circle></svg>',content:"<section class=\"text-gray-600 body-font\"><div class=\"container px-5 py-24 mx-auto\"><div class=\"flex flex-wrap -m-4\"><div class=\"lg:w-1/3 lg:mb-0 mb-6 p-4\"><div class=\"h-full text-center\"><img alt=\"testimonial\" class=\"w-20 h-20 mb-8 object-cover object-center rounded-full inline-block border-2 border-gray-200 bg-gray-100\" src=\"https://dummyimage.com/302x302\"><p class=\"leading-relaxed\">Edison bulb retro cloud bread echo park, helvetica stumptown taiyaki taxidermy 90's cronut +1 kinfolk. Single-origin coffee ennui shaman taiyaki vape DIY tote bag drinking vinegar cronut adaptogen squid fanny pack vaporware.</p><span class=\"inline-block h-1 w-10 rounded bg-indigo-500 mt-6 mb-4\"></span><h2 class=\"text-gray-900 font-medium title-font tracking-wider text-sm\">HOLDEN CAULFIELD</h2><p class=\"text-gray-500\">Senior Product Designer</p></div></div><div class=\"lg:w-1/3 lg:mb-0 mb-6 p-4\"><div class=\"h-full text-center\"><img alt=\"testimonial\" class=\"w-20 h-20 mb-8 object-cover object-center rounded-full inline-block border-2 border-gray-200 bg-gray-100\" src=\"https://dummyimage.com/300x300\"><p class=\"leading-relaxed\">Edison bulb retro cloud bread echo park, helvetica stumptown taiyaki taxidermy 90's cronut +1 kinfolk. Single-origin coffee ennui shaman taiyaki vape DIY tote bag drinking vinegar cronut adaptogen squid fanny pack vaporware.</p><span class=\"inline-block h-1 w-10 rounded bg-indigo-500 mt-6 mb-4\"></span><h2 class=\"text-gray-900 font-medium title-font tracking-wider text-sm\">ALPER KAMU</h2><p class=\"text-gray-500\">UI Develeoper</p></div></div><div class=\"lg:w-1/3 lg:mb-0 p-4\"><div class=\"h-full text-center\"><img alt=\"testimonial\" class=\"w-20 h-20 mb-8 object-cover object-center rounded-full inline-block border-2 border-gray-200 bg-gray-100\" src=\"https://dummyimage.com/305x305\"><p class=\"leading-relaxed\">Edison bulb retro cloud bread echo park, helvetica stumptown taiyaki taxidermy 90's cronut +1 kinfolk. Single-origin coffee ennui shaman taiyaki vape DIY tote bag drinking vinegar cronut adaptogen squid fanny pack vaporware.</p><span class=\"inline-block h-1 w-10 rounded bg-indigo-500 mt-6 mb-4\"></span><h2 class=\"text-gray-900 font-medium title-font tracking-wider text-sm\">HENRY LETHAM</h2><p class=\"text-gray-500\">CTO</p></div></div></div></div></section>",category:'Testimonials'}],d=function(t){var e=t.Blocks;c.forEach((function(t){e.add(t.id,{label:t.label,attributes:{class:t.class},content:t.content,category:{label:t.category,open:'Blog'===t.category}});}));},h=function(t){d(t);},x=[{name:'slate',color:'#64748b'},{name:'gray',color:'#6b7280'},{name:'zinc',color:'#71717a'},{name:'neutral',color:'#737373'},{name:'stone',color:'#78716c'},{name:'red',color:'#ef4444'},{name:'orange',color:'#f97316'},{name:'amber',color:'#f59e0b'},{name:'yellow',color:'#eab308'},{name:'lime',color:'#84cc16'},{name:'green',color:'#22c55e'},{name:'emerald',color:'#10b981'},{name:'teal',color:'#14b8a6'},{name:'cyan',color:'#06b6d4'},{name:'sky',color:'#0ea5e9'},{name:'blue',color:'#3b82f6'},{name:'indigo',color:'#6366f1'},{name:'violet',color:'#8b5cf6'},{name:'purple',color:'#a855f7'},{name:'fuchsia',color:'#d946ef'},{name:'pink',color:'#ec4899'},{name:'rose',color:'#f43f5e'}],g=new RegExp(/(bg|text|border|ring)-(red|yellow|green|blue|indigo|purple|green)-(\d\d\d)/,'g'),f=function(t){var e,l=t.Modal,i=t.getConfig().stylePrefix,r=document.createElement('div'),a=document.createElement('div');a.style.padding='40px 0px',a.style.display='flex',a.style.justifyContent='center',a.style.flexWrap='wrap',x.forEach((function(t){var l=document.createElement('button');l.className='change-theme-button',l.style.backgroundColor=t.color,l.onclick=function(){return e=t},a.appendChild(l);}));var s=document.createElement('div'),o=document.createElement('button');return o.innerHTML='Update',o.className=i+'btn-prim '+i+'btn-import',o.style.float='right',o.onclick=function(){p(t,e.name),l.close();},s.appendChild(o),r.appendChild(a),r.appendChild(s),r},p=function(t,e){(function t(e){var l=arguments.length>1&&void 0!==arguments[1]?arguments[1]:[];return l.push(e),e.components().each((function(e){return t(e,l)})),l})(t.DomComponents.getWrapper(),[]).forEach((function(t){var l,i,r=t.view.el;'string'==typeof(null===(l=r.className)||void 0===l?void 0:l.baseVal)&&null!==(i=r.className)&&void 0!==i&&i.baseVal.match(g)?(r.className.baseVal=r.className.baseVal.replace(g,"$1-".concat(e,"-$3")),t.replaceWith(r.outerHTML)):'string'==typeof r.className&&r.className.match(g)&&(r.className=r.className.replace(g,"$1-".concat(e,"-$3")),t.replaceWith(r.outerHTML));}));},m=function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{},l=t.Commands;l.add('open-update-theme',{run:function(l,i){(null==i?void 0:i.set)&&i.set('active',0);var r=t.Modal;r.setTitle(e.changeThemeText);var a=f(t);r.setContent(a),r.open();}}),l.add('get-tailwindCss',{run:function(t,l){var i=arguments.length>2&&void 0!==arguments[2]?arguments[2]:{};(null==l?void 0:l.set)&&l.set('active',0);var r=i.callback,a=void 0===r?function(t){return console.log(t)}:r,s=e.cover,o=t.Canvas.getDocument();o&&(o.head.querySelectorAll('style').forEach((function(t){t.innerText.includes('tailwind')&&(s+=t.innerText);})),a(s));}});},u={'grapesjs-tailwind':{}};function v(t,e){var l=Object.keys(t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(t);e&&(i=i.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),l.push.apply(l,i);}return l}function w(t){for(var e=1;e<arguments.length;e++){var l=null!=arguments[e]?arguments[e]:{};e%2?v(Object(l),!0).forEach((function(e){s()(t,e,l[e]);})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(l)):v(Object(l)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(l,e));}));}return t}e["default"]=function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{},l=w(w({},{i18n:{},tailwindPlayCdn:'https://cdn.tailwindcss.com',plugins:[],config:{},cover:"img.object-cover { filter: sepia(1) hue-rotate(190deg) opacity(.46) grayscale(.7) !important; }",changeThemeText:'Change Theme'}),e);h(t),m(t,l),t.I18n&&t.I18n.addMessages(w({en:u},l.i18n));var i=function(){var t=r()(n.a.mark((function t(e){var i,r,a,s,o,c,d,h,x;return n.a.wrap((function(t){for(;1;)switch(t.prev=t.next){case 0:if(i=e.view.getEl()){t.next=3;break}return t.abrupt("return");case 3:r=l.tailwindPlayCdn,a=l.plugins,s=l.config,o=l.cover,c=function(){i.contentWindow.tailwind.config=s;},(d=document.createElement('script')).src=r+(a.length?"?plugins=".concat(a.join()):''),d.onload=c,(h=document.createElement('style')).innerHTML=o,x=setInterval((function(){var t=i.contentDocument;'complete'===t.readyState&&(t.head.appendChild(d),t.head.appendChild(h),clearInterval(x));}),100);case 11:case"end":return t.stop()}}),t)})));return function(e){return t.apply(this,arguments)}}();t.Canvas.getModel()['on']('change:frames',(function(t,e){e.forEach((function(t){return t.once('loaded',(function(){return i(t)}))}));}));};}])}));

    });

    var twPlugin = /*@__PURE__*/getDefaultExportFromCjs(grapesjsTailwind_min);

    const easyPageStore = (service) => (editor) => {
        console.log("@grapejs_editor", editor);
        editor.Panels.addButton("options", {
            id: "the_save_button",
            className: "saveButton fa fa-floppy-o",
            command: (editor) => {
                editor.store({});
            },
            attributes: { title: "Save" },
            active: true,
        });
        editor.Panels.addButton("options", {
            id: "the_go_home",
            className: "goHome fa fa-home",
            command: async (editor) => {
                location.hash = "/";
            },
            attributes: { title: "Home" },
            active: true,
        });
        editor.Storage.add("easypage-store", {
            async load(options = {}) {
                const resp = await service.getPageData(options["page_slug"]);
                if (!resp.ok) {
                    console.log("Err", resp);
                    return {};
                }
                return JSON.parse(resp.data || "{}");
            },
            async store(data, options = {}) {
                data["gen_html"] = extractHtml(editor);
                const resp = await service.setPageData(options["page_slug"], JSON.stringify(data));
                if (!resp.ok) {
                    console.log("Err", resp);
                    return;
                }
            },
        });
    };
    const extractHtml = (editor) => {
        return editor.Pages.getAll().map((page) => {
            const component = page.getMainComponent();
            return {
                html: editor.getHtml({ component }),
                css: editor.getCss({ component }),
            };
        });
    };

    /* entries/adapter_editor_easypage/page/_builder/builder.svelte generated by Svelte v3.48.0 */
    const file = "entries/adapter_editor_easypage/page/_builder/builder.svelte";

    function create_fragment$3(ctx) {
    	let link;
    	let t0;
    	let div;

    	const block = {
    		c: function create() {
    			link = element("link");
    			t0 = space();
    			div = element("div");
    			div.textContent = "Site Builder";
    			attr_dev(link, "rel", "stylesheet");
    			attr_dev(link, "href", "https://unpkg.com/grapesjs/dist/css/grapes.min.css");
    			add_location(link, file, 51, 2, 1467);
    			add_location(div, file, 58, 0, 1651);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			append_dev(document.head, link);
    			insert_dev(target, t0, anchor);
    			insert_dev(target, div, anchor);
    			/*div_binding*/ ctx[3](div);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			detach_dev(link);
    			if (detaching) detach_dev(t0);
    			if (detaching) detach_dev(div);
    			/*div_binding*/ ctx[3](null);
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
    	validate_slots('Builder', slots, []);
    	let { page_slug } = $$props;
    	let { service } = $$props;
    	let rootElem;
    	let editor;

    	onMount(() => {
    		editor = grapejs.init({
    			container: rootElem,
    			plugins: [
    				webpagePlugin,
    				basicPlugin,
    				gjsForms,
    				navPlugin,
    				stgrPlugin,
    				styleFilter,
    				customCodePlugin,
    				blkFlexboxPlugin,
    				"grapesjs-lory-slider",
    				tabPlugin,
    				toolTipPlugin,
    				twPlugin,
    				easyPageStore(service)
    			],
    			pluginsOpts: {},
    			storageManager: {
    				type: "easypage-store",
    				stepsBeforeSave: 3,
    				options: { "easypage-store": { page_slug } }
    			}
    		});
    	});

    	const writable_props = ['page_slug', 'service'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Builder> was created with unknown prop '${key}'`);
    	});

    	function div_binding($$value) {
    		binding_callbacks[$$value ? 'unshift' : 'push'](() => {
    			rootElem = $$value;
    			$$invalidate(0, rootElem);
    		});
    	}

    	$$self.$$set = $$props => {
    		if ('page_slug' in $$props) $$invalidate(1, page_slug = $$props.page_slug);
    		if ('service' in $$props) $$invalidate(2, service = $$props.service);
    	};

    	$$self.$capture_state = () => ({
    		grapejs,
    		webpagePlugin,
    		basicPlugin,
    		gjsForms,
    		navPlugin,
    		customCodePlugin,
    		blkFlexboxPlugin,
    		stgrPlugin,
    		styleFilter,
    		tabPlugin,
    		toolTipPlugin,
    		twPlugin,
    		onMount,
    		easyPageStore,
    		page_slug,
    		service,
    		rootElem,
    		editor
    	});

    	$$self.$inject_state = $$props => {
    		if ('page_slug' in $$props) $$invalidate(1, page_slug = $$props.page_slug);
    		if ('service' in $$props) $$invalidate(2, service = $$props.service);
    		if ('rootElem' in $$props) $$invalidate(0, rootElem = $$props.rootElem);
    		if ('editor' in $$props) editor = $$props.editor;
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [rootElem, page_slug, service, div_binding];
    }

    class Builder extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$3, create_fragment$3, safe_not_equal, { page_slug: 1, service: 2 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Builder",
    			options,
    			id: create_fragment$3.name
    		});

    		const { ctx } = this.$$;
    		const props = options.props || {};

    		if (/*page_slug*/ ctx[1] === undefined && !('page_slug' in props)) {
    			console.warn("<Builder> was created without expected prop 'page_slug'");
    		}

    		if (/*service*/ ctx[2] === undefined && !('service' in props)) {
    			console.warn("<Builder> was created without expected prop 'service'");
    		}
    	}

    	get page_slug() {
    		throw new Error("<Builder>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set page_slug(value) {
    		throw new Error("<Builder>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get service() {
    		throw new Error("<Builder>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set service(value) {
    		throw new Error("<Builder>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/adapter_editor_easypage/page/page.svelte generated by Svelte v3.48.0 */

    function create_fragment$2(ctx) {
    	let builder;
    	let current;

    	builder = new Builder({
    			props: {
    				page_slug: /*pid*/ ctx[0],
    				service: /*service*/ ctx[1]
    			},
    			$$inline: true
    		});

    	const block = {
    		c: function create() {
    			create_component(builder.$$.fragment);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			mount_component(builder, target, anchor);
    			current = true;
    		},
    		p: function update(ctx, [dirty]) {
    			const builder_changes = {};
    			if (dirty & /*pid*/ 1) builder_changes.page_slug = /*pid*/ ctx[0];
    			builder.$set(builder_changes);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(builder.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(builder.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(builder, detaching);
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
    	let $params;
    	validate_store(params, 'params');
    	component_subscribe($$self, params, $$value => $$invalidate(2, $params = $$value));
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Page', slots, []);
    	let { pid = $params.pid } = $$props;
    	const service = getContext("__easypage_service__");
    	const writable_props = ['pid'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Page> was created with unknown prop '${key}'`);
    	});

    	$$self.$$set = $$props => {
    		if ('pid' in $$props) $$invalidate(0, pid = $$props.pid);
    	};

    	$$self.$capture_state = () => ({
    		Builder,
    		params,
    		getContext,
    		pid,
    		service,
    		$params
    	});

    	$$self.$inject_state = $$props => {
    		if ('pid' in $$props) $$invalidate(0, pid = $$props.pid);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [pid, service];
    }

    class Page extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$2, create_fragment$2, safe_not_equal, { pid: 0 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Page",
    			options,
    			id: create_fragment$2.name
    		});
    	}

    	get pid() {
    		throw new Error("<Page>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set pid(value) {
    		throw new Error("<Page>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/xcompo/common/_tailwind.svelte generated by Svelte v3.48.0 */

    function create_fragment$1(ctx) {
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
    		id: create_fragment$1.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$1($$self, $$props) {
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
    		init(this, options, instance$1, create_fragment$1, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Tailwind",
    			options,
    			id: create_fragment$1.name
    		});
    	}
    }

    class EasypageService {
        constructor(env) {
            this.load = () => {
                return this.api.perform_action("load", null);
            };
            this.updatePages = (data) => {
                return this.api.perform_action("update_pages", data);
            };
            this.getPageData = (slug) => {
                return this.api.perform_action("get_page_data", slug);
            };
            this.setPageData = (slug, data) => {
                return this.api.perform_action("set_page_data", {
                    slug,
                    data,
                });
            };
            this.deletePageData = (slug) => {
                return this.api.perform_action("delete_page_data", slug);
            };
            this.api = env.api;
            this.env = env;
        }
    }

    /* entries/adapter_editor_easypage/index.svelte generated by Svelte v3.48.0 */

    function create_fragment(ctx) {
    	let modal;
    	let updating_show_big;
    	let updating_close_big;
    	let updating_show_small;
    	let updating_close_small;
    	let t0;
    	let tailwind;
    	let t1;
    	let router;
    	let current;

    	function modal_show_big_binding(value) {
    		/*modal_show_big_binding*/ ctx[5](value);
    	}

    	function modal_close_big_binding(value) {
    		/*modal_close_big_binding*/ ctx[6](value);
    	}

    	function modal_show_small_binding(value) {
    		/*modal_show_small_binding*/ ctx[7](value);
    	}

    	function modal_close_small_binding(value) {
    		/*modal_close_small_binding*/ ctx[8](value);
    	}

    	let modal_props = {};

    	if (/*big_open*/ ctx[0] !== void 0) {
    		modal_props.show_big = /*big_open*/ ctx[0];
    	}

    	if (/*big_close*/ ctx[1] !== void 0) {
    		modal_props.close_big = /*big_close*/ ctx[1];
    	}

    	if (/*small_open*/ ctx[2] !== void 0) {
    		modal_props.show_small = /*small_open*/ ctx[2];
    	}

    	if (/*small_close*/ ctx[3] !== void 0) {
    		modal_props.close_small = /*small_close*/ ctx[3];
    	}

    	modal = new Modal({ props: modal_props, $$inline: true });
    	binding_callbacks.push(() => bind(modal, 'show_big', modal_show_big_binding));
    	binding_callbacks.push(() => bind(modal, 'close_big', modal_close_big_binding));
    	binding_callbacks.push(() => bind(modal, 'show_small', modal_show_small_binding));
    	binding_callbacks.push(() => bind(modal, 'close_small', modal_close_small_binding));
    	tailwind = new Tailwind({ $$inline: true });
    	router = new Router({ $$inline: true });

    	const block = {
    		c: function create() {
    			create_component(modal.$$.fragment);
    			t0 = space();
    			create_component(tailwind.$$.fragment);
    			t1 = space();
    			create_component(router.$$.fragment);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			mount_component(modal, target, anchor);
    			insert_dev(target, t0, anchor);
    			mount_component(tailwind, target, anchor);
    			insert_dev(target, t1, anchor);
    			mount_component(router, target, anchor);
    			current = true;
    		},
    		p: function update(ctx, [dirty]) {
    			const modal_changes = {};

    			if (!updating_show_big && dirty & /*big_open*/ 1) {
    				updating_show_big = true;
    				modal_changes.show_big = /*big_open*/ ctx[0];
    				add_flush_callback(() => updating_show_big = false);
    			}

    			if (!updating_close_big && dirty & /*big_close*/ 2) {
    				updating_close_big = true;
    				modal_changes.close_big = /*big_close*/ ctx[1];
    				add_flush_callback(() => updating_close_big = false);
    			}

    			if (!updating_show_small && dirty & /*small_open*/ 4) {
    				updating_show_small = true;
    				modal_changes.show_small = /*small_open*/ ctx[2];
    				add_flush_callback(() => updating_show_small = false);
    			}

    			if (!updating_close_small && dirty & /*small_close*/ 8) {
    				updating_close_small = true;
    				modal_changes.close_small = /*small_close*/ ctx[3];
    				add_flush_callback(() => updating_close_small = false);
    			}

    			modal.$set(modal_changes);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(modal.$$.fragment, local);
    			transition_in(tailwind.$$.fragment, local);
    			transition_in(router.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(modal.$$.fragment, local);
    			transition_out(tailwind.$$.fragment, local);
    			transition_out(router.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(modal, detaching);
    			if (detaching) detach_dev(t0);
    			destroy_component(tailwind, detaching);
    			if (detaching) detach_dev(t1);
    			destroy_component(router, detaching);
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
    	validate_slots('Adapter_editor_easypage', slots, []);
    	let { env } = $$props;
    	routes.set({ "/": Start, "/page/:pid": Page });
    	let big_open;
    	let big_close;
    	let small_open;
    	let small_close;
    	let service = new EasypageService(env);
    	setContext("__easypage_service__", service);

    	onMount(() => {
    		service.modal = {
    			big_open,
    			big_close,
    			small_open,
    			small_close
    		};
    	});

    	const writable_props = ['env'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Adapter_editor_easypage> was created with unknown prop '${key}'`);
    	});

    	function modal_show_big_binding(value) {
    		big_open = value;
    		$$invalidate(0, big_open);
    	}

    	function modal_close_big_binding(value) {
    		big_close = value;
    		$$invalidate(1, big_close);
    	}

    	function modal_show_small_binding(value) {
    		small_open = value;
    		$$invalidate(2, small_open);
    	}

    	function modal_close_small_binding(value) {
    		small_close = value;
    		$$invalidate(3, small_close);
    	}

    	$$self.$$set = $$props => {
    		if ('env' in $$props) $$invalidate(4, env = $$props.env);
    	};

    	$$self.$capture_state = () => ({
    		routes,
    		Router,
    		Start,
    		Page,
    		Tailwind,
    		EasypageService,
    		onMount,
    		setContext,
    		Modal,
    		env,
    		big_open,
    		big_close,
    		small_open,
    		small_close,
    		service
    	});

    	$$self.$inject_state = $$props => {
    		if ('env' in $$props) $$invalidate(4, env = $$props.env);
    		if ('big_open' in $$props) $$invalidate(0, big_open = $$props.big_open);
    		if ('big_close' in $$props) $$invalidate(1, big_close = $$props.big_close);
    		if ('small_open' in $$props) $$invalidate(2, small_open = $$props.small_open);
    		if ('small_close' in $$props) $$invalidate(3, small_close = $$props.small_close);
    		if ('service' in $$props) service = $$props.service;
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [
    		big_open,
    		big_close,
    		small_open,
    		small_close,
    		env,
    		modal_show_big_binding,
    		modal_close_big_binding,
    		modal_show_small_binding,
    		modal_close_small_binding
    	];
    }

    class Adapter_editor_easypage extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance, create_fragment, safe_not_equal, { env: 4 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Adapter_editor_easypage",
    			options,
    			id: create_fragment.name
    		});

    		const { ctx } = this.$$;
    		const props = options.props || {};

    		if (/*env*/ ctx[4] === undefined && !('env' in props)) {
    			console.warn("<Adapter_editor_easypage> was created without expected prop 'env'");
    		}
    	}

    	get env() {
    		throw new Error("<Adapter_editor_easypage>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set env(value) {
    		throw new Error("<Adapter_editor_easypage>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    const r = window["__registry__"];
    r.RegisterFactory("temphia.adapter_editor.loader", `easypage.main`, (opts) => {
        new Adapter_editor_easypage({
            target: opts.target,
            props: { env: opts.env },
        });
    });

})();
//# sourceMappingURL=adapter_editor_easypage.js.map
var app = (function () {
    'use strict';

    function noop() { }
    function assign(tar, src) {
        // @ts-ignore
        for (const k in src)
            tar[k] = src[k];
        return tar;
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
    let src_url_equal_anchor;
    function src_url_equal(element_src, url) {
        if (!src_url_equal_anchor) {
            src_url_equal_anchor = document.createElement('a');
        }
        src_url_equal_anchor.href = url;
        return element_src === src_url_equal_anchor.href;
    }
    function is_empty(obj) {
        return Object.keys(obj).length === 0;
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
        if (node.parentNode) {
            node.parentNode.removeChild(node);
        }
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
    function prevent_default(fn) {
        return function (event) {
            event.preventDefault();
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
    function children(element) {
        return Array.from(element.childNodes);
    }
    function set_data(text, data) {
        data = '' + data;
        if (text.data === data)
            return;
        text.data = data;
    }
    function set_input_value(input, value) {
        input.value = value == null ? '' : value;
    }
    function set_style(node, key, value, important) {
        if (value == null) {
            node.style.removeProperty(key);
        }
        else {
            node.style.setProperty(key, value, important ? 'important' : '');
        }
    }
    function construct_svelte_component(component, props) {
        return new component(props);
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
    /**
     * Schedules a callback to run immediately before the component is unmounted.
     *
     * Out of `onMount`, `beforeUpdate`, `afterUpdate` and `onDestroy`, this is the
     * only one that runs inside a server-side component.
     *
     * https://svelte.dev/docs#run-time-svelte-ondestroy
     */
    function onDestroy(fn) {
        get_current_component().$$.on_destroy.push(fn);
    }
    /**
     * Associates an arbitrary `context` object with the current component and the specified `key`
     * and returns that object. The context is then available to children of the component
     * (including slotted content) with `getContext`.
     *
     * Like lifecycle functions, this must be called during component initialisation.
     *
     * https://svelte.dev/docs#run-time-svelte-setcontext
     */
    function setContext(key, context) {
        get_current_component().$$.context.set(key, context);
        return context;
    }
    /**
     * Retrieves the context that belongs to the closest parent component with the specified `key`.
     * Must be called during component initialisation.
     *
     * https://svelte.dev/docs#run-time-svelte-getcontext
     */
    function getContext(key) {
        return get_current_component().$$.context.get(key);
    }

    const dirty_components = [];
    const binding_callbacks = [];
    let render_callbacks = [];
    const flush_callbacks = [];
    const resolved_promise = /* @__PURE__ */ Promise.resolve();
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
        // Do not reenter flush while dirty components are updated, as this can
        // result in an infinite loop. Instead, let the inner flush handle it.
        // Reentrancy is ok afterwards for bindings etc.
        if (flushidx !== 0) {
            return;
        }
        const saved_component = current_component;
        do {
            // first, call beforeUpdate functions
            // and update components
            try {
                while (flushidx < dirty_components.length) {
                    const component = dirty_components[flushidx];
                    flushidx++;
                    set_current_component(component);
                    update(component.$$);
                }
            }
            catch (e) {
                // reset dirty state to not end up in a deadlocked state and then rethrow
                dirty_components.length = 0;
                flushidx = 0;
                throw e;
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
    /**
     * Useful for example to execute remaining `afterUpdate` callbacks before executing `destroy`.
     */
    function flush_render_callbacks(fns) {
        const filtered = [];
        const targets = [];
        render_callbacks.forEach((c) => fns.indexOf(c) === -1 ? filtered.push(c) : targets.push(c));
        targets.forEach((c) => c());
        render_callbacks = filtered;
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
        else if (callback) {
            callback();
        }
    }

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
        const { fragment, after_update } = component.$$;
        fragment && fragment.m(target, anchor);
        if (!customElement) {
            // onMount happens before the initial afterUpdate
            add_render_callback(() => {
                const new_on_destroy = component.$$.on_mount.map(run).filter(is_function);
                // if the component was destroyed immediately
                // it will update the `$$.on_destroy` reference to `null`.
                // the destructured on_destroy may still reference to the old array
                if (component.$$.on_destroy) {
                    component.$$.on_destroy.push(...new_on_destroy);
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
            flush_render_callbacks($$.after_update);
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
            ctx: [],
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
            if (!is_function(callback)) {
                return noop;
            }
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

    const subscriber_queue = [];
    /**
     * Creates a `Readable` store that allows reading by subscription.
     * @param value initial value
     * @param {StartStopNotifier} [start]
     */
    function readable(value, start) {
        return {
            subscribe: writable(value, start).subscribe
        };
    }
    /**
     * Create a `Writable` store that allows both updating and reading by subscription.
     * @param {*=}value initial value
     * @param {StartStopNotifier=} start
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
                if (subscribers.size === 0 && stop) {
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
            let started = false;
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
                if (started) {
                    sync();
                }
            }, () => {
                pending |= (1 << i);
            }));
            started = true;
            sync();
            return function stop() {
                run_all(unsubscribers);
                cleanup();
                // We need to set this to false because callbacks can still happen despite having unsubscribed:
                // Callbacks might already be placed in the queue which doesn't know it should no longer
                // invoke this derived store.
                started = false;
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

    /* node_modules/svelte-hash-router/src/components/Router.svelte generated by Svelte v3.59.2 */

    function create_fragment$j(ctx) {
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

    		return { props: switch_instance_props };
    	}

    	if (switch_value) {
    		switch_instance = construct_svelte_component(switch_value, switch_props());
    	}

    	return {
    		c() {
    			if (switch_instance) create_component(switch_instance.$$.fragment);
    			switch_instance_anchor = empty();
    		},
    		m(target, anchor) {
    			if (switch_instance) mount_component(switch_instance, target, anchor);
    			insert(target, switch_instance_anchor, anchor);
    			current = true;
    		},
    		p(ctx, [dirty]) {
    			const switch_instance_changes = (dirty & /*$$props*/ 4)
    			? get_spread_update(switch_instance_spread_levels, [get_spread_object(/*$$props*/ ctx[2])])
    			: {};

    			if (dirty & /*$components*/ 1 && switch_value !== (switch_value = /*$components*/ ctx[0][/*i*/ ctx[1]])) {
    				if (switch_instance) {
    					group_outros();
    					const old_component = switch_instance;

    					transition_out(old_component.$$.fragment, 1, 0, () => {
    						destroy_component(old_component, 1);
    					});

    					check_outros();
    				}

    				if (switch_value) {
    					switch_instance = construct_svelte_component(switch_value, switch_props());
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
    		i(local) {
    			if (current) return;
    			if (switch_instance) transition_in(switch_instance.$$.fragment, local);
    			current = true;
    		},
    		o(local) {
    			if (switch_instance) transition_out(switch_instance.$$.fragment, local);
    			current = false;
    		},
    		d(detaching) {
    			if (detaching) detach(switch_instance_anchor);
    			if (switch_instance) destroy_component(switch_instance, detaching);
    		}
    	};
    }

    let level = 0;

    function instance$f($$self, $$props, $$invalidate) {
    	let $components;
    	component_subscribe($$self, components, $$value => $$invalidate(0, $components = $$value));
    	let i = level++;
    	onDestroy(() => level--);

    	$$self.$$set = $$new_props => {
    		$$invalidate(2, $$props = assign(assign({}, $$props), exclude_internal_props($$new_props)));
    	};

    	$$props = exclude_internal_props($$props);
    	return [$components, i, $$props];
    }

    class Router extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, instance$f, create_fragment$j, safe_not_equal, {});
    	}
    }

    /* src/entries/auth/pages/start/index.svelte generated by Svelte v3.59.2 */

    function instance$e($$self) {
    	const app = getContext("_auth_app_");

    	(async () => {
    		if (!app.site_manager.isLogged()) {
    			app.nav.goto_login_page();
    			return;
    		}

    		app.nav.goto_final_page();
    		return;
    	})();

    	return [];
    }

    class Start extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, instance$e, null, safe_not_equal, {});
    	}
    }

    class Http {
        constructor(baseURL, headers) {
            this.baseURL = baseURL;
            this.headers = headers;
        }
        replace_headers(headers) {
            this.headers = headers;
        }
        async get(path) {
            const resp = await fetch(`${this.baseURL}${path}`, {
                method: "GET",
                headers: this.headers,
            });
            if (resp.ok) {
                const text = await resp.text();
                try {
                    const data = JSON.parse(text);
                    return Promise.resolve({
                        ok: true,
                        data,
                        status: resp.status,
                    });
                }
                catch (error) {
                    return Promise.resolve({
                        ok: true,
                        data: text,
                        status: resp.status,
                    });
                }
            }
            return Promise.resolve({
                ok: false,
                data: await resp.text(),
                status: resp.status,
            });
        }
        async post(path, data) {
            return this.jsonMethod(path, "POST", data);
        }
        async patch(path, data) {
            return this.jsonMethod(path, "PATCH", data);
        }
        async put(path, data) {
            return this.jsonMethod(path, "PUT", data);
        }
        async jsonMethod(path, method, data) {
            const resp = await fetch(`${this.baseURL}${path}`, {
                method: method,
                headers: this.headers,
                body: JSON.stringify(data),
                mode: "cors",
            });
            if (resp.ok) {
                return Promise.resolve({
                    ok: true,
                    data: await resp.json(),
                    status: resp.status,
                });
            }
            return Promise.resolve({
                ok: false,
                data: await resp.text(),
                status: resp.status,
            });
        }
        async rawFetch(path, method, auth, build_path, data) {
            return await fetch(build_path ? `${this.baseURL}${path}` : path, {
                method: method,
                headers: auth ? { Authorization: this.headers["Authorization"] } : {},
                body: data,
            });
        }
        async postForm(path, auth, data) {
            return await fetch(`${this.baseURL}${path}`, {
                method: "POST",
                headers: auth ? { Authorization: this.headers["Authorization"] } : {},
                body: data,
            });
        }
        async patchForm(path, auth, data) {
            return await fetch(`${this.baseURL}${path}`, {
                method: "PATCH",
                headers: auth ? { Authorization: this.headers["Authorization"] } : {},
                body: data,
            });
        }
        async delete(path, data) {
            const resp = await fetch(`${this.baseURL}${path}`, {
                method: "DELETE",
                headers: this.headers,
                body: data ? JSON.stringify(data) : data,
            });
            if (resp.ok) {
                return Promise.resolve({
                    ok: true,
                    data: await resp.json(),
                    status: resp.status,
                });
            }
            return Promise.resolve({
                ok: false,
                data: await resp.text(),
                status: resp.status,
            });
        }
    }

    class AuthAPI {
        constructor(baseUrl, site_token) {
            this.list_methods = async (ugroup) => {
                return this.http.get(`/auth?ugroup=${ugroup}`);
            };
            this.login_next = async (data) => {
                return this.http.post("/auth/login/next", data);
            };
            this.login_submit = async (data) => {
                return this.http.post("/auth/login/submit", data);
            };
            this.altauth_generate = async (id, data) => {
                return this.http.post(`/auth/alt/${id}/generate`, data);
            };
            this.altauth_next = async (id, stage, data) => {
                return this.http.post(`/auth/alt/${id}/next/${stage}`, data);
            };
            this.altauth_submit = async (id, data) => {
                return this.http.post(`/auth/alt/${id}/submit`, data);
            };
            this.finish = async (data) => {
                return this.http.post("/auth/finish", data);
            };
            this.signup_next = async (data) => {
                return this.http.post("/auth/signup/next", data);
            };
            this.signup_submit = async (data) => {
                return this.http.post("/auth/signup/submit", data);
            };
            this.reset_submit = async (data) => {
                return this.http.post("/reset/submit", data);
            };
            this.reset_finish = async (data) => {
                return this.http.post("/reset/finish", data);
            };
            this.about = async (user_token) => {
                const http = new Http(this.http.baseURL, {
                    "Content-Type": "application/json",
                    Authorization: user_token,
                });
                return http.get("/auth/about");
            };
            this.http = new Http(baseUrl, {
                "Content-Type": "application/json",
                Authorization: site_token,
            });
        }
    }

    const AUTHED_KEY_PREFIX = "_temphia_authed_key_";
    class SiteUtils {
        constructor(site_token) {
            this._site_token = site_token;
        }
        isLogged() {
            return !!this.get();
        }
        gotoLoginPage() {
            window.location.pathname = "/z/auth";
        }
        setAuthedData(data) {
            const pdata = JSON.stringify(data);
            this.set(pdata);
        }
        getAuthedData() {
            const raw = this.get();
            const data = JSON.parse(raw);
            return data;
        }
        clearAuthedData() {
            localStorage.removeItem(this.key());
        }
        get() {
            return localStorage.getItem(this.key());
        }
        set(data) {
            localStorage.setItem(this.key(), data);
        }
        key() {
            // tenantify ?
            return AUTHED_KEY_PREFIX;
        }
    }

    //http://localhost:4000/z/api/:tenant_id/v2
    const apiURL = (tenant_id) => `${window.location.origin}/z/api/${tenant_id}/v2`;
    const authURL = (opts) => {
        if (!opts) {
            return `${window.location.origin}/z/auth`;
        }
        return `${window.location.origin}/z/auth?${opts.tenant_id ? "tenant_id=" + opts.tenant_id + "&" : ""}${opts.user_group ? "ugroup=" + opts.user_group : ""}`;
    };

    class AuthNav {
        constructor() {
            this.goto = (target, opts) => {
                this.nav_options = opts;
                window.location.hash = target;
            };
            this.goto_login_page = () => {
                this.goto("/login/");
            };
            this.goto_login_next_stage = (opts) => {
                this.goto("/login/next_stage", opts);
            };
            this.goto_alt_first_stage = (opts) => {
                this.goto("/alt/first_stage", opts);
            };
            this.goto_alt_second_stage = (opts) => {
                this.goto("/alt/second_stage", opts);
            };
            this.goto_prehook_page = (opts) => {
                this.goto("/prehook", opts);
            };
            this.goto_final_page = () => {
                this.goto("/final");
            };
            this.goto_error_page = (reason, opts) => {
                this.err_message = reason;
                this.goto("/error", opts);
            };
        }
    }

    class AuthService {
        constructor() {
            this.list_methods = async () => {
                const resp = await this.auth_api.list_methods(this.user_group);
                if (resp.status !== 200) {
                    return null;
                }
                return {
                    pass_auth: resp.data["pass_auth"],
                    open_signup: resp.data["open_signup"],
                    alt_auth_method: resp.data["alt_auth_method"],
                };
            };
            this.login_next = async (uid, pass) => {
                return this.auth_api.login_next({
                    user_ident: uid,
                    password: pass,
                    site_token: this._site_token,
                });
            };
            this.login_submit = async (ntoken) => {
                return this.auth_api.login_submit({
                    site_token: this._site_token,
                    next_token: ntoken,
                });
            };
            this.login_finish = async (pre_auth_token, proof) => {
                return this.auth_api.finish({
                    site_token: this._site_token,
                    preauthed_token: pre_auth_token,
                    proof_token: proof,
                });
            };
            this.generate_alt_auth = async (id) => {
                this.active_auth_id = id;
                return this.auth_api.altauth_generate(id, {
                    site_token: this._site_token,
                    user_group: this.user_group,
                });
            };
            this.alt_next_first = async (code, state) => {
                return this.auth_api.altauth_next(this.active_auth_id, "first", {
                    auth_code: code,
                    auth_state: state,
                    site_token: this._site_token,
                    user_group: this.user_group,
                });
            };
            this.alt_next_second = async (first_token, signup_data) => {
                return this.auth_api.altauth_next(this.active_auth_id, "second", {
                    site_token: this._site_token,
                    first_token: first_token,
                    signup_data: signup_data,
                });
            };
            this.submit_alt_auth = async (next_token) => {
                return this.auth_api.altauth_submit(this.active_auth_id, {
                    next_token: next_token,
                    site_token: this._site_token,
                });
            };
            // preauthed data
            this.get_preauthed_data = () => {
                if (this.nav.nav_options) {
                    return this.nav.nav_options;
                }
                const raw = localStorage.getItem("__pre_authed_data");
                return JSON.parse(raw);
            };
            this.save_preauthed_data = (data) => {
                localStorage.setItem("__pre_authed_data", JSON.stringify(data));
            };
            this.clear_preauthed_data = () => {
                localStorage.removeItem("__pre_authed_data");
            };
            // authed data
            this.save_authed_data = (user_token) => {
                this.site_manager.setAuthedData({
                    site_token: this._site_token,
                    tenant_id: this.tenant_id,
                    user_token: user_token,
                });
            };
            this.clear_authed_data = () => {
                this.site_manager.clearAuthedData();
            };
            this.about = async () => {
                const adata = this.site_manager.getAuthedData();
                const resp = await this.auth_api.about(adata.user_token);
                if (!resp.ok) {
                    return;
                }
                return resp.data;
            };
            const site_data = window["__temphia_site_data__"];
            this.site_manager = new SiteUtils(site_data.site_token);
            this.auth_api = new AuthAPI(apiURL(site_data.tenant_id), site_data.site_token);
            this.active_auth_id = 0;
            this.user_group = site_data.user_group;
            this.user_group_fixed = false;
            this.tenant_id = site_data.tenant_id;
            this.nav = new AuthNav();
            this._site_token = site_data.site_token;
        }
    }

    /* src/entries/auth/pages/login/_methods/_oauth.svelte generated by Svelte v3.59.2 */

    function create_fragment$i(ctx) {
    	let button;
    	let svg;
    	let path0;
    	let path1;
    	let t0;
    	let t1_value = /*method*/ ctx[0]["name"] + "";
    	let t1;
    	let mounted;
    	let dispose;

    	return {
    		c() {
    			button = element("button");
    			svg = svg_element("svg");
    			path0 = svg_element("path");
    			path1 = svg_element("path");
    			t0 = text("\n\n  Open ");
    			t1 = text(t1_value);
    			attr(path0, "d", "M11 3a1 1 0 100 2h2.586l-6.293 6.293a1 1 0 101.414 1.414L15 6.414V9a1 1 0 102 0V4a1 1 0 00-1-1h-5z");
    			attr(path1, "d", "M5 5a2 2 0 00-2 2v8a2 2 0 002 2h8a2 2 0 002-2v-3a1 1 0 10-2 0v3H5V7h3a1 1 0 000-2H5z");
    			attr(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr(svg, "class", "h-5 w-5");
    			attr(svg, "viewBox", "0 0 20 20");
    			attr(svg, "fill", "currentColor");
    			attr(button, "class", "p-2 text-white text-lg rounded bg-gray-600 flex justify-center");
    		},
    		m(target, anchor) {
    			insert(target, button, anchor);
    			append(button, svg);
    			append(svg, path0);
    			append(svg, path1);
    			append(button, t0);
    			append(button, t1);

    			if (!mounted) {
    				dispose = listen(button, "click", /*handleClick*/ ctx[1]);
    				mounted = true;
    			}
    		},
    		p(ctx, [dirty]) {
    			if (dirty & /*method*/ 1 && t1_value !== (t1_value = /*method*/ ctx[0]["name"] + "")) set_data(t1, t1_value);
    		},
    		i: noop,
    		o: noop,
    		d(detaching) {
    			if (detaching) detach(button);
    			mounted = false;
    			dispose();
    		}
    	};
    }

    function instance$d($$self, $$props, $$invalidate) {
    	let { app } = $$props;
    	let { method } = $$props;
    	let { data } = $$props;

    	const handleClick = ev => {
    		const redirect_uri = `${window.origin}/z/auth/oauth_redirect`;

    		let up = new URLSearchParams({
    				client_id: data["client_id"],
    				response_type: "code",
    				redirect_uri,
    				state: data["state_token"]
    			});

    		data["scopes"].forEach(s => {
    			up.set("scope", s);
    		});

    		const tabWindow = window.open(`${data["auth_url"]}?${up.toString()}`, "_blank");
    		console.log("@tabwindow =>", tabWindow);

    		const i = setInterval(
    			() => {
    				console.log("@checking .....");
    				const { location } = tabWindow;

    				try {
    					if (location.href.indexOf(redirect_uri) !== 0) return;
    					parseAndExtract(location.search);
    				} catch(error) {
    					if (error instanceof DOMException || error.message === "Permission denied") {
    						return;
    					}

    					if (!tabWindow.closed) tabWindow.close();
    				}

    				tabWindow.close();
    				clearInterval(i);
    			},
    			250
    		);
    	};

    	const parseAndExtract = async qstr => {
    		let up = new URLSearchParams(qstr);
    		const resp = await app.alt_next_first(up.get("code"), up.get("state"));

    		if (resp.status !== 200) {
    			console.log("Err", resp);
    			return;
    		}

    		app.nav.goto_alt_first_stage(resp.data);
    	};

    	$$self.$$set = $$props => {
    		if ('app' in $$props) $$invalidate(2, app = $$props.app);
    		if ('method' in $$props) $$invalidate(0, method = $$props.method);
    		if ('data' in $$props) $$invalidate(3, data = $$props.data);
    	};

    	return [method, handleClick, app, data];
    }

    class Oauth extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, instance$d, create_fragment$i, safe_not_equal, { app: 2, method: 0, data: 3 });
    	}
    }

    /* src/entries/auth/pages/login/_methods/alt_method.svelte generated by Svelte v3.59.2 */

    function create_if_block$9(ctx) {
    	let oauth;
    	let current;

    	oauth = new Oauth({
    			props: {
    				app: /*app*/ ctx[0],
    				data: /*data*/ ctx[2],
    				method: /*method*/ ctx[1]
    			}
    		});

    	return {
    		c() {
    			create_component(oauth.$$.fragment);
    		},
    		m(target, anchor) {
    			mount_component(oauth, target, anchor);
    			current = true;
    		},
    		p(ctx, dirty) {
    			const oauth_changes = {};
    			if (dirty & /*app*/ 1) oauth_changes.app = /*app*/ ctx[0];
    			if (dirty & /*data*/ 4) oauth_changes.data = /*data*/ ctx[2];
    			if (dirty & /*method*/ 2) oauth_changes.method = /*method*/ ctx[1];
    			oauth.$set(oauth_changes);
    		},
    		i(local) {
    			if (current) return;
    			transition_in(oauth.$$.fragment, local);
    			current = true;
    		},
    		o(local) {
    			transition_out(oauth.$$.fragment, local);
    			current = false;
    		},
    		d(detaching) {
    			destroy_component(oauth, detaching);
    		}
    	};
    }

    function create_fragment$h(ctx) {
    	let div;
    	let current;
    	let if_block = /*method*/ ctx[1]["type"] === "oauth" && create_if_block$9(ctx);

    	return {
    		c() {
    			div = element("div");
    			if (if_block) if_block.c();
    			attr(div, "class", "p-1 flex flex-col");
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    			if (if_block) if_block.m(div, null);
    			current = true;
    		},
    		p(ctx, [dirty]) {
    			if (/*method*/ ctx[1]["type"] === "oauth") {
    				if (if_block) {
    					if_block.p(ctx, dirty);

    					if (dirty & /*method*/ 2) {
    						transition_in(if_block, 1);
    					}
    				} else {
    					if_block = create_if_block$9(ctx);
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
    		i(local) {
    			if (current) return;
    			transition_in(if_block);
    			current = true;
    		},
    		o(local) {
    			transition_out(if_block);
    			current = false;
    		},
    		d(detaching) {
    			if (detaching) detach(div);
    			if (if_block) if_block.d();
    		}
    	};
    }

    function instance$c($$self, $$props, $$invalidate) {
    	let { app } = $$props;
    	let { method } = $$props;
    	let { data } = $$props;

    	$$self.$$set = $$props => {
    		if ('app' in $$props) $$invalidate(0, app = $$props.app);
    		if ('method' in $$props) $$invalidate(1, method = $$props.method);
    		if ('data' in $$props) $$invalidate(2, data = $$props.data);
    	};

    	return [app, method, data];
    }

    class Alt_method extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, instance$c, create_fragment$h, safe_not_equal, { app: 0, method: 1, data: 2 });
    	}
    }

    /* src/entries/auth/pages/login/_methods/password.svelte generated by Svelte v3.59.2 */

    function create_if_block$8(ctx) {
    	let p;
    	let t;

    	return {
    		c() {
    			p = element("p");
    			t = text(/*message*/ ctx[2]);
    		},
    		m(target, anchor) {
    			insert(target, p, anchor);
    			append(p, t);
    		},
    		p(ctx, dirty) {
    			if (dirty & /*message*/ 4) set_data(t, /*message*/ ctx[2]);
    		},
    		d(detaching) {
    			if (detaching) detach(p);
    		}
    	};
    }

    function create_fragment$g(ctx) {
    	let t0;
    	let div0;
    	let label0;
    	let t2;
    	let input0;
    	let t3;
    	let div2;
    	let div1;
    	let t7;
    	let input1;
    	let t8;
    	let div3;
    	let button;
    	let mounted;
    	let dispose;
    	let if_block = /*message*/ ctx[2] && create_if_block$8(ctx);

    	return {
    		c() {
    			if (if_block) if_block.c();
    			t0 = space();
    			div0 = element("div");
    			label0 = element("label");
    			label0.textContent = "Email Address / Username";
    			t2 = space();
    			input0 = element("input");
    			t3 = space();
    			div2 = element("div");
    			div1 = element("div");

    			div1.innerHTML = `<label class="block mb-2 text-sm font-medium text-gray-600" for="loggingPassword">Password</label> 
    <a href="#" class="text-xs text-gray-500 hover:underline">Forget Password?</a>`;

    			t7 = space();
    			input1 = element("input");
    			t8 = space();
    			div3 = element("div");
    			button = element("button");
    			button.textContent = "Login";
    			attr(label0, "class", "block mb-2 text-sm font-medium text-gray-600");
    			attr(label0, "for", "LoggingEmailAddress");
    			attr(input0, "id", "LoggingEmailAddress");
    			attr(input0, "class", "block w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-md focus:border-blue-500 focus:outline-none focus:ring");
    			attr(input0, "type", "email");
    			attr(div0, "class", "mt-4");
    			attr(div1, "class", "flex justify-between");
    			attr(input1, "class", "block w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-md focus:border-blue-500 focus:outline-none focus:ring");
    			attr(input1, "type", "password");
    			attr(div2, "class", "mt-4");
    			attr(button, "class", "w-full px-4 py-2 tracking-wide text-white font-semibold transition-colors duration-200 transform bg-blue-700 rounded hover:bg-blue-400");
    			attr(div3, "class", "mt-8");
    		},
    		m(target, anchor) {
    			if (if_block) if_block.m(target, anchor);
    			insert(target, t0, anchor);
    			insert(target, div0, anchor);
    			append(div0, label0);
    			append(div0, t2);
    			append(div0, input0);
    			set_input_value(input0, /*emailuser*/ ctx[0]);
    			insert(target, t3, anchor);
    			insert(target, div2, anchor);
    			append(div2, div1);
    			append(div2, t7);
    			append(div2, input1);
    			set_input_value(input1, /*password*/ ctx[1]);
    			insert(target, t8, anchor);
    			insert(target, div3, anchor);
    			append(div3, button);

    			if (!mounted) {
    				dispose = [
    					listen(input0, "input", /*input0_input_handler*/ ctx[5]),
    					listen(input1, "input", /*input1_input_handler*/ ctx[6]),
    					listen(button, "click", /*submit*/ ctx[3])
    				];

    				mounted = true;
    			}
    		},
    		p(ctx, [dirty]) {
    			if (/*message*/ ctx[2]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);
    				} else {
    					if_block = create_if_block$8(ctx);
    					if_block.c();
    					if_block.m(t0.parentNode, t0);
    				}
    			} else if (if_block) {
    				if_block.d(1);
    				if_block = null;
    			}

    			if (dirty & /*emailuser*/ 1 && input0.value !== /*emailuser*/ ctx[0]) {
    				set_input_value(input0, /*emailuser*/ ctx[0]);
    			}

    			if (dirty & /*password*/ 2 && input1.value !== /*password*/ ctx[1]) {
    				set_input_value(input1, /*password*/ ctx[1]);
    			}
    		},
    		i: noop,
    		o: noop,
    		d(detaching) {
    			if (if_block) if_block.d(detaching);
    			if (detaching) detach(t0);
    			if (detaching) detach(div0);
    			if (detaching) detach(t3);
    			if (detaching) detach(div2);
    			if (detaching) detach(t8);
    			if (detaching) detach(div3);
    			mounted = false;
    			run_all(dispose);
    		}
    	};
    }

    function instance$b($$self, $$props, $$invalidate) {
    	let { app } = $$props;
    	let emailuser;
    	let password;
    	let message;

    	const submit = async () => {
    		const resp = await app.login_next(emailuser, password);

    		if (resp.status !== 200) {
    			return;
    		}

    		if (resp.data["ok"]) {
    			app.nav.goto_login_next_stage(resp.data);
    		} else {
    			$$invalidate(2, message = resp.data["message"]);
    		}
    	};

    	function input0_input_handler() {
    		emailuser = this.value;
    		$$invalidate(0, emailuser);
    	}

    	function input1_input_handler() {
    		password = this.value;
    		$$invalidate(1, password);
    	}

    	$$self.$$set = $$props => {
    		if ('app' in $$props) $$invalidate(4, app = $$props.app);
    	};

    	return [
    		emailuser,
    		password,
    		message,
    		submit,
    		app,
    		input0_input_handler,
    		input1_input_handler
    	];
    }

    class Password extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, instance$b, create_fragment$g, safe_not_equal, { app: 4 });
    	}
    }

    var Icons = {
        "google": "https://icons.duckduckgo.com/ip3/google.com.ico"
    };

    /* src/entries/auth/pages/login/_inner.svelte generated by Svelte v3.59.2 */

    function get_each_context$1(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[10] = list[i];
    	return child_ctx;
    }

    // (67:0) {:else}
    function create_else_block$2(ctx) {
    	let altmethod;
    	let current;

    	altmethod = new Alt_method({
    			props: {
    				app: /*app*/ ctx[0],
    				data: /*data*/ ctx[5],
    				method: /*selected_method*/ ctx[4]
    			}
    		});

    	return {
    		c() {
    			create_component(altmethod.$$.fragment);
    		},
    		m(target, anchor) {
    			mount_component(altmethod, target, anchor);
    			current = true;
    		},
    		p(ctx, dirty) {
    			const altmethod_changes = {};
    			if (dirty & /*app*/ 1) altmethod_changes.app = /*app*/ ctx[0];
    			if (dirty & /*data*/ 32) altmethod_changes.data = /*data*/ ctx[5];
    			if (dirty & /*selected_method*/ 16) altmethod_changes.method = /*selected_method*/ ctx[4];
    			altmethod.$set(altmethod_changes);
    		},
    		i(local) {
    			if (current) return;
    			transition_in(altmethod.$$.fragment, local);
    			current = true;
    		},
    		o(local) {
    			transition_out(altmethod.$$.fragment, local);
    			current = false;
    		},
    		d(detaching) {
    			destroy_component(altmethod, detaching);
    		}
    	};
    }

    // (28:0) {#if !alt_auth_mode}
    function create_if_block$7(ctx) {
    	let t0;
    	let t1;
    	let div;
    	let t2;
    	let if_block2_anchor;
    	let current;
    	let if_block0 = /*password_mode*/ ctx[6] && create_if_block_4(ctx);
    	let if_block1 = /*alt_methods*/ ctx[1] && create_if_block_3();
    	let each_value = /*alt_methods*/ ctx[1];
    	let each_blocks = [];

    	for (let i = 0; i < each_value.length; i += 1) {
    		each_blocks[i] = create_each_block$1(get_each_context$1(ctx, each_value, i));
    	}

    	let if_block2 = /*opensignup*/ ctx[2] && create_if_block_1$1();

    	return {
    		c() {
    			if (if_block0) if_block0.c();
    			t0 = space();
    			if (if_block1) if_block1.c();
    			t1 = space();
    			div = element("div");

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].c();
    			}

    			t2 = space();
    			if (if_block2) if_block2.c();
    			if_block2_anchor = empty();
    			attr(div, "class", "p-4 flex flex-col border mt-2");
    		},
    		m(target, anchor) {
    			if (if_block0) if_block0.m(target, anchor);
    			insert(target, t0, anchor);
    			if (if_block1) if_block1.m(target, anchor);
    			insert(target, t1, anchor);
    			insert(target, div, anchor);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				if (each_blocks[i]) {
    					each_blocks[i].m(div, null);
    				}
    			}

    			insert(target, t2, anchor);
    			if (if_block2) if_block2.m(target, anchor);
    			insert(target, if_block2_anchor, anchor);
    			current = true;
    		},
    		p(ctx, dirty) {
    			if (/*password_mode*/ ctx[6]) if_block0.p(ctx, dirty);

    			if (/*alt_methods*/ ctx[1]) {
    				if (if_block1) ; else {
    					if_block1 = create_if_block_3();
    					if_block1.c();
    					if_block1.m(t1.parentNode, t1);
    				}
    			} else if (if_block1) {
    				if_block1.d(1);
    				if_block1 = null;
    			}

    			if (dirty & /*oauthNext, alt_methods, Icons*/ 130) {
    				each_value = /*alt_methods*/ ctx[1];
    				let i;

    				for (i = 0; i < each_value.length; i += 1) {
    					const child_ctx = get_each_context$1(ctx, each_value, i);

    					if (each_blocks[i]) {
    						each_blocks[i].p(child_ctx, dirty);
    					} else {
    						each_blocks[i] = create_each_block$1(child_ctx);
    						each_blocks[i].c();
    						each_blocks[i].m(div, null);
    					}
    				}

    				for (; i < each_blocks.length; i += 1) {
    					each_blocks[i].d(1);
    				}

    				each_blocks.length = each_value.length;
    			}

    			if (/*opensignup*/ ctx[2]) {
    				if (if_block2) ; else {
    					if_block2 = create_if_block_1$1();
    					if_block2.c();
    					if_block2.m(if_block2_anchor.parentNode, if_block2_anchor);
    				}
    			} else if (if_block2) {
    				if_block2.d(1);
    				if_block2 = null;
    			}
    		},
    		i(local) {
    			if (current) return;
    			transition_in(if_block0);
    			current = true;
    		},
    		o(local) {
    			transition_out(if_block0);
    			current = false;
    		},
    		d(detaching) {
    			if (if_block0) if_block0.d(detaching);
    			if (detaching) detach(t0);
    			if (if_block1) if_block1.d(detaching);
    			if (detaching) detach(t1);
    			if (detaching) detach(div);
    			destroy_each(each_blocks, detaching);
    			if (detaching) detach(t2);
    			if (if_block2) if_block2.d(detaching);
    			if (detaching) detach(if_block2_anchor);
    		}
    	};
    }

    // (29:2) {#if password_mode}
    function create_if_block_4(ctx) {
    	let password_1;
    	let current;
    	password_1 = new Password({ props: { app: /*app*/ ctx[0] } });

    	return {
    		c() {
    			create_component(password_1.$$.fragment);
    		},
    		m(target, anchor) {
    			mount_component(password_1, target, anchor);
    			current = true;
    		},
    		p(ctx, dirty) {
    			const password_1_changes = {};
    			if (dirty & /*app*/ 1) password_1_changes.app = /*app*/ ctx[0];
    			password_1.$set(password_1_changes);
    		},
    		i(local) {
    			if (current) return;
    			transition_in(password_1.$$.fragment, local);
    			current = true;
    		},
    		o(local) {
    			transition_out(password_1.$$.fragment, local);
    			current = false;
    		},
    		d(detaching) {
    			destroy_component(password_1, detaching);
    		}
    	};
    }

    // (33:2) {#if alt_methods}
    function create_if_block_3(ctx) {
    	let div;

    	return {
    		c() {
    			div = element("div");

    			div.innerHTML = `<hr class="w-full bg-gray-400"/> 
      <p class="text-base font-medium leading-4 px-2.5 text-gray-400">OR</p> 
      <hr class="w-full bg-gray-400 "/>`;

    			attr(div, "class", "w-full flex items-center justify-between py-5");
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    		},
    		d(detaching) {
    			if (detaching) detach(div);
    		}
    	};
    }

    // (43:6) {#if method.type === "oauth"}
    function create_if_block_2(ctx) {
    	let button;
    	let img;
    	let img_src_value;
    	let t0;
    	let t1_value = /*method*/ ctx[10].name + "";
    	let t1;
    	let t2;
    	let mounted;
    	let dispose;

    	function click_handler() {
    		return /*click_handler*/ ctx[9](/*method*/ ctx[10]);
    	}

    	return {
    		c() {
    			button = element("button");
    			img = element("img");
    			t0 = space();
    			t1 = text(t1_value);
    			t2 = space();
    			if (!src_url_equal(img.src, img_src_value = Icons[/*method*/ ctx[10]["provider"]] || "")) attr(img, "src", img_src_value);
    			attr(img, "alt", "");
    			attr(button, "class", "w-full p-2 text-gray-600 border rounded-lg shadow-md hover:bg-gray-200 flex justify-center gap-2");
    		},
    		m(target, anchor) {
    			insert(target, button, anchor);
    			append(button, img);
    			append(button, t0);
    			append(button, t1);
    			append(button, t2);

    			if (!mounted) {
    				dispose = listen(button, "click", click_handler);
    				mounted = true;
    			}
    		},
    		p(new_ctx, dirty) {
    			ctx = new_ctx;

    			if (dirty & /*alt_methods*/ 2 && !src_url_equal(img.src, img_src_value = Icons[/*method*/ ctx[10]["provider"]] || "")) {
    				attr(img, "src", img_src_value);
    			}

    			if (dirty & /*alt_methods*/ 2 && t1_value !== (t1_value = /*method*/ ctx[10].name + "")) set_data(t1, t1_value);
    		},
    		d(detaching) {
    			if (detaching) detach(button);
    			mounted = false;
    			dispose();
    		}
    	};
    }

    // (42:4) {#each alt_methods as method}
    function create_each_block$1(ctx) {
    	let if_block_anchor;
    	let if_block = /*method*/ ctx[10].type === "oauth" && create_if_block_2(ctx);

    	return {
    		c() {
    			if (if_block) if_block.c();
    			if_block_anchor = empty();
    		},
    		m(target, anchor) {
    			if (if_block) if_block.m(target, anchor);
    			insert(target, if_block_anchor, anchor);
    		},
    		p(ctx, dirty) {
    			if (/*method*/ ctx[10].type === "oauth") {
    				if (if_block) {
    					if_block.p(ctx, dirty);
    				} else {
    					if_block = create_if_block_2(ctx);
    					if_block.c();
    					if_block.m(if_block_anchor.parentNode, if_block_anchor);
    				}
    			} else if (if_block) {
    				if_block.d(1);
    				if_block = null;
    			}
    		},
    		d(detaching) {
    			if (if_block) if_block.d(detaching);
    			if (detaching) detach(if_block_anchor);
    		}
    	};
    }

    // (56:2) {#if opensignup}
    function create_if_block_1$1(ctx) {
    	let div;

    	return {
    		c() {
    			div = element("div");

    			div.innerHTML = `<span class="w-1/5 border-b md:w-1/4"></span> 

      <a href="#" class="text-xs text-gray-500 uppercase hover:underline">or sign up</a> 

      <span class="w-1/5 border-b md:w-1/4"></span>`;

    			attr(div, "class", "flex items-center justify-between mt-4");
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    		},
    		d(detaching) {
    			if (detaching) detach(div);
    		}
    	};
    }

    function create_fragment$f(ctx) {
    	let h2;
    	let t1;
    	let current_block_type_index;
    	let if_block;
    	let if_block_anchor;
    	let current;
    	const if_block_creators = [create_if_block$7, create_else_block$2];
    	const if_blocks = [];

    	function select_block_type(ctx, dirty) {
    		if (!/*alt_auth_mode*/ ctx[3]) return 0;
    		return 1;
    	}

    	current_block_type_index = select_block_type(ctx);
    	if_block = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx);

    	return {
    		c() {
    			h2 = element("h2");
    			h2.textContent = "Temphia User Login";
    			t1 = space();
    			if_block.c();
    			if_block_anchor = empty();
    			attr(h2, "class", "text-2xl font-semibold text-center text-gray-700 mb-5");
    		},
    		m(target, anchor) {
    			insert(target, h2, anchor);
    			insert(target, t1, anchor);
    			if_blocks[current_block_type_index].m(target, anchor);
    			insert(target, if_block_anchor, anchor);
    			current = true;
    		},
    		p(ctx, [dirty]) {
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
    		i(local) {
    			if (current) return;
    			transition_in(if_block);
    			current = true;
    		},
    		o(local) {
    			transition_out(if_block);
    			current = false;
    		},
    		d(detaching) {
    			if (detaching) detach(h2);
    			if (detaching) detach(t1);
    			if_blocks[current_block_type_index].d(detaching);
    			if (detaching) detach(if_block_anchor);
    		}
    	};
    }

    function instance$a($$self, $$props, $$invalidate) {
    	let { app } = $$props;
    	let { alt_methods = [] } = $$props;
    	let { password = false } = $$props;
    	let { opensignup = false } = $$props;
    	let password_mode = password;
    	let alt_auth_mode = false;
    	let selected_method;
    	let data;

    	const oauthNext = async method => {
    		const resp = await app.generate_alt_auth(Number(method["id"]));

    		if (resp.status !== 200) {
    			console.log("Error", resp);
    			return;
    		}

    		$$invalidate(4, selected_method = method);
    		$$invalidate(5, data = resp.data);
    		$$invalidate(3, alt_auth_mode = true);
    	};

    	const click_handler = method => oauthNext(method);

    	$$self.$$set = $$props => {
    		if ('app' in $$props) $$invalidate(0, app = $$props.app);
    		if ('alt_methods' in $$props) $$invalidate(1, alt_methods = $$props.alt_methods);
    		if ('password' in $$props) $$invalidate(8, password = $$props.password);
    		if ('opensignup' in $$props) $$invalidate(2, opensignup = $$props.opensignup);
    	};

    	return [
    		app,
    		alt_methods,
    		opensignup,
    		alt_auth_mode,
    		selected_method,
    		data,
    		password_mode,
    		oauthNext,
    		password,
    		click_handler
    	];
    }

    class Inner extends SvelteComponent {
    	constructor(options) {
    		super();

    		init(this, options, instance$a, create_fragment$f, safe_not_equal, {
    			app: 0,
    			alt_methods: 1,
    			password: 8,
    			opensignup: 2
    		});
    	}
    }

    /* src/entries/auth/pages/common/layout.svelte generated by Svelte v3.59.2 */

    function create_if_block$6(ctx) {
    	let div;
    	let p;
    	let t0;
    	let span;
    	let t2;
    	let button;
    	let mounted;
    	let dispose;

    	return {
    		c() {
    			div = element("div");
    			p = element("p");
    			t0 = text("Login as user group ");
    			span = element("span");
    			span.textContent = `${/*app*/ ctx[0].user_group}`;
    			t2 = space();
    			button = element("button");
    			button.textContent = "Change";
    			attr(span, "class", "p-0.5 rounded bg-slate-200 text-slate-700");
    			attr(button, "class", "text-blue-500");
    			attr(p, "class", "mt-2 text-sm text-gray-500");
    			attr(div, "class", "w-full h-8 mt-5");
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    			append(div, p);
    			append(p, t0);
    			append(p, span);
    			append(p, t2);
    			append(p, button);

    			if (!mounted) {
    				dispose = listen(button, "click", /*click_handler*/ ctx[3]);
    				mounted = true;
    			}
    		},
    		p: noop,
    		d(detaching) {
    			if (detaching) detach(div);
    			mounted = false;
    			dispose();
    		}
    	};
    }

    function create_fragment$e(ctx) {
    	let div4;
    	let div3;
    	let div2;
    	let div0;
    	let t0;
    	let div1;
    	let t1;
    	let current;
    	const default_slot_template = /*#slots*/ ctx[2].default;
    	const default_slot = create_slot(default_slot_template, ctx, /*$$scope*/ ctx[1], null);
    	let if_block = !/*app*/ ctx[0].user_group_fixed && create_if_block$6(ctx);

    	return {
    		c() {
    			div4 = element("div");
    			div3 = element("div");
    			div2 = element("div");
    			div0 = element("div");
    			div0.innerHTML = `<svg data-name="Layer 1" viewBox="0 0 744.85 747.08" xmlns="http://www.w3.org/2000/svg"><path transform="translate(-227.58 -76.461)" d="m299.2 705.81-6.56-25.872a335.97 335.97 0 0 0-35.643-12.788l-0.828 12.024-3.358-13.247c-15.021-4.2939-25.24-6.183-25.24-6.183s13.8 52.489 42.754 92.617l33.734 5.926-26.207 3.779a135.93 135.93 0 0 0 11.719 12.422c42.115 39.092 89.024 57.028 104.77 40.06s-5.625-62.412-47.74-101.5c-13.056-12.119-29.457-21.844-45.875-29.5z" fill="#f2f2f2" data-name="Path 1"></path><path transform="translate(-227.58 -76.461)" d="m361.59 677.71 7.758-25.538a335.94 335.94 0 0 0-23.9-29.371l-6.924 9.865 3.972-13.076c-10.641-11.436-18.412-18.335-18.412-18.335s-15.315 52.067-11.275 101.38l25.815 22.51-24.392-10.312a135.92 135.92 0 0 0 3.614 16.694c15.846 55.234 46.731 94.835 68.983 88.451s27.446-56.335 11.6-111.57c-4.912-17.123-13.926-33.926-24.023-48.965z" fill="#f2f2f2" data-name="Path 2"></path><path transform="translate(-227.58 -76.461)" d="m747.33 253.44h-4.092v-112.1a64.883 64.883 0 0 0-64.883-64.883h-237.51a64.883 64.883 0 0 0-64.883 64.883v615a64.883 64.883 0 0 0 64.883 64.883h237.51a64.883 64.883 0 0 0 64.882-64.883v-423.1h4.092z" fill="#e6e6e6" data-name="Path 22"></path><path transform="translate(-227.58 -76.461)" d="m680.97 93.336h-31a23.02 23.02 0 0 1-21.316 31.714h-136.06a23.02 23.02 0 0 1-21.314-31.714h-28.956a48.454 48.454 0 0 0-48.454 48.454v614.11a48.454 48.454 0 0 0 48.454 48.454h238.65a48.454 48.454 0 0 0 48.454-48.454v-614.11a48.454 48.454 0 0 0-48.454-48.453z" fill="#fff" data-name="Path 23"></path><path transform="translate(-227.58 -76.461)" d="m531.23 337.96a24.437 24.437 0 0 1 12.23-21.174 24.45 24.45 0 1 0 0 42.345 24.434 24.434 0 0 1-12.23-21.171z" fill="#ccc" data-name="Path 6"></path><path transform="translate(-227.58 -76.461)" d="m561.97 337.96a24.436 24.436 0 0 1 12.23-21.174 24.45 24.45 0 1 0 0 42.345 24.434 24.434 0 0 1-12.23-21.171z" fill="#ccc" data-name="Path 7"></path><circle cx="364.43" cy="261.5" r="24.45" fill="#6c63ff" data-name="Ellipse 1"></circle><path transform="translate(-227.58 -76.461)" d="m632.87 414.33h-142.5a5.123 5.123 0 0 1-5.117-5.117v-142.5a5.123 5.123 0 0 1 5.117-5.117h142.5a5.123 5.123 0 0 1 5.117 5.117v142.5a5.123 5.123 0 0 1-5.117 5.117zm-142.5-150.69a3.073 3.073 0 0 0-3.07 3.07v142.5a3.073 3.073 0 0 0 3.07 3.07h142.5a3.073 3.073 0 0 0 3.07-3.07v-142.5a3.073 3.073 0 0 0-3.07-3.07z" fill="#ccc" data-name="Path 8"></path><rect x="218.56" y="447.1" width="218.55" height="2.047" fill="#ccc" data-name="Rectangle 1"></rect><circle cx="225.46" cy="427.42" r="6.902" fill="#6c63ff" data-name="Ellipse 2"></circle><rect x="218.56" y="516.12" width="218.55" height="2.047" fill="#ccc" data-name="Rectangle 2"></rect><circle cx="225.46" cy="496.44" r="6.902" fill="#6c63ff" data-name="Ellipse 3"></circle><path transform="translate(-227.58 -76.461)" d="m660.69 671.17h-69.068a4.5049 4.5049 0 0 1-4.5-4.5v-24.208a4.5049 4.5049 0 0 1 4.5-4.5h69.068a4.5049 4.5049 0 0 1 4.5 4.5v24.208a4.5049 4.5049 0 0 1-4.5 4.5z" fill="#6c63ff"></path><circle cx="247.98" cy="427.42" r="6.902" fill="#6c63ff" data-name="Ellipse 7"></circle><circle cx="270.49" cy="427.42" r="6.902" fill="#6c63ff" data-name="Ellipse 8"></circle><circle cx="247.98" cy="496.44" r="6.902" fill="#6c63ff" data-name="Ellipse 9"></circle><circle cx="270.49" cy="496.44" r="6.902" fill="#6c63ff" data-name="Ellipse 10"></circle><path transform="translate(-227.58 -76.461)" d="m969.64 823.54h-717.99c-1.537 0-2.782-0.546-2.782-1.218s1.245-1.219 2.782-1.219h717.99c1.536 0 2.782 0.546 2.782 1.219s-1.246 1.218-2.782 1.218z" fill="#3f3d56" data-name="Path 88"></path><path transform="translate(-227.58 -76.461)" d="m792.25 565.92a10.094 10.094 0 0 1 1.4108 0.78731l44.852-19.143 1.6009-11.815 17.922-0.10956-1.0587 27.098-59.2 15.656a10.608 10.608 0 0 1-0.44749 1.2084 10.235 10.235 0 1 1-5.0795-13.682z" fill="#ffb8b8"></path><polygon points="636.98 735.02 624.72 735.02 618.89 687.73 636.98 687.73" fill="#ffb8b8"></polygon><path d="m615.96 731.52h23.644v14.887h-38.531a14.887 14.887 0 0 1 14.887-14.887z" fill="#2f2e41"></path><polygon points="684.66 731.56 672.46 732.76 662.02 686.27 680.02 684.5" fill="#ffb8b8"></polygon><path transform="translate(-303.01 15.291) rotate(-5.6253)" d="m891.69 806.13h23.644v14.887h-38.531a14.887 14.887 0 0 1 14.887-14.887z" fill="#2f2e41"></path><circle cx="640.39" cy="384.57" r="24.561" fill="#ffb8b8"></circle><path transform="translate(-227.58 -76.461)" d="m849.56 801.92a4.4709 4.4709 0 0 1-4.415-3.6973c-6.3457-35.226-27.088-150.41-27.584-153.6a1.4268 1.4268 0 0 1-0.01562-0.22168v-8.5879a1.489 1.489 0 0 1 0.27929-0.87207l2.7402-3.8379a1.4784 1.4784 0 0 1 1.1436-0.625c15.622-0.73242 66.784-2.8789 69.256 0.209 2.4824 3.1035 1.6055 12.507 1.4043 14.36l0.00977 0.19336 22.985 147a4.5124 4.5124 0 0 1-3.7148 5.1348l-14.356 2.3652a4.5213 4.5213 0 0 1-5.0254-3.0928c-4.4404-14.188-19.329-61.918-24.489-80.387a0.49922 0.49922 0 0 0-0.98047 0.13868c0.25781 17.605 0.88086 62.523 1.0957 78.037l0.02344 1.6709a4.5181 4.5181 0 0 1-4.0928 4.5361l-13.844 1.2578c-0.14066 0.01268-0.28131 0.01854-0.41995 0.01854z" fill="#2f2e41"></path><path transform="translate(-227.58 -76.461)" d="m852.38 495.25c-4.2863 2.548-6.8512 7.2304-8.3228 11.995a113.68 113.68 0 0 0-4.8844 27.159l-1.5555 27.6-19.255 73.17c16.689 14.121 26.315 10.912 48.78-0.63879s25.032 3.8512 25.032 3.8512l4.4924-62.258 6.4184-68.032a30.164 30.164 0 0 0-4.8614-4.6742 49.658 49.658 0 0 0-42.442-8.9954z" fill="#fff" data-name="Path 99"></path><path transform="translate(-227.58 -76.461)" d="m846.13 580.7a10.526 10.526 0 0 1 1.5006 0.70389l44.348-22.197 0.736-12.026 18.294-1.2613 0.98041 27.413-59.266 19.599a10.496 10.496 0 1 1-6.5933-12.232z" fill="#ffb8b8"></path><path transform="translate(-227.58 -76.461)" d="m902.77 508.41c10.912 3.8512 12.834 45.574 12.834 45.574-12.837-7.0604-28.241 4.4932-28.241 4.4932s-3.2092-10.912-7.0603-25.032a24.53 24.53 0 0 1 5.1344-23.106s6.4223-5.7819 17.334-1.9284z" fill="#fff" data-name="Path 101"></path><path transform="translate(-227.58 -76.461)" d="m889.99 467.53c-3.06-2.4484-7.2352 2.0017-7.2352 2.0017l-2.4484-22.033s-15.301 1.8329-25.094-0.61161-11.323 8.8751-11.323 8.8751a78.58 78.58 0 0 1-0.30582-13.771c0.61158-5.5084 8.5684-11.017 22.645-14.689s21.421 12.241 21.421 12.241c9.7936 4.8958 5.3994 30.436 2.3394 27.988z" fill="#2f2e41" data-name="Path 102"></path></svg>`;
    			t0 = space();
    			div1 = element("div");
    			if (default_slot) default_slot.c();
    			t1 = space();
    			if (if_block) if_block.c();
    			attr(div0, "class", "hidden md:block w-1/2 bg-indigo-500 py-10 px-10");
    			attr(div1, "class", "w-full md:w-1/2 py-10 px-5 md:px-10");
    			attr(div2, "class", "md:flex w-full");
    			attr(div3, "class", "bg-gray-50 text-gray-500 rounded-3xl shadow-xl w-full overflow-hidden");
    			set_style(div3, "max-width", "1000px");
    			attr(div4, "class", "min-w-screen min-h-screen bg-blue-200 flex items-center justify-center px-5 py-5");
    		},
    		m(target, anchor) {
    			insert(target, div4, anchor);
    			append(div4, div3);
    			append(div3, div2);
    			append(div2, div0);
    			append(div2, t0);
    			append(div2, div1);

    			if (default_slot) {
    				default_slot.m(div1, null);
    			}

    			append(div1, t1);
    			if (if_block) if_block.m(div1, null);
    			current = true;
    		},
    		p(ctx, [dirty]) {
    			if (default_slot) {
    				if (default_slot.p && (!current || dirty & /*$$scope*/ 2)) {
    					update_slot_base(
    						default_slot,
    						default_slot_template,
    						ctx,
    						/*$$scope*/ ctx[1],
    						!current
    						? get_all_dirty_from_scope(/*$$scope*/ ctx[1])
    						: get_slot_changes(default_slot_template, /*$$scope*/ ctx[1], dirty, null),
    						null
    					);
    				}
    			}

    			if (!/*app*/ ctx[0].user_group_fixed) if_block.p(ctx, dirty);
    		},
    		i(local) {
    			if (current) return;
    			transition_in(default_slot, local);
    			current = true;
    		},
    		o(local) {
    			transition_out(default_slot, local);
    			current = false;
    		},
    		d(detaching) {
    			if (detaching) detach(div4);
    			if (default_slot) default_slot.d(detaching);
    			if (if_block) if_block.d();
    		}
    	};
    }

    function instance$9($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	const app = getContext("_auth_app_");

    	const click_handler = () => {
    		window.location.href = authURL({
    			tenant_id: app.tenant_id,
    			user_group: prompt("Enter new user group you belong to.", app.user_group)
    		});

    		app.tenant_id;
    	};

    	$$self.$$set = $$props => {
    		if ('$$scope' in $$props) $$invalidate(1, $$scope = $$props.$$scope);
    	};

    	return [app, $$scope, slots, click_handler];
    }

    class Layout extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, instance$9, create_fragment$e, safe_not_equal, {});
    	}
    }

    /* src/entries/auth/pages/login/login.svelte generated by Svelte v3.59.2 */

    function create_else_block$1(ctx) {
    	let div;

    	return {
    		c() {
    			div = element("div");
    			div.textContent = "Loading..";
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d(detaching) {
    			if (detaching) detach(div);
    		}
    	};
    }

    // (19:2) {#if loaded}
    function create_if_block$5(ctx) {
    	let inner;
    	let current;

    	inner = new Inner({
    			props: {
    				app: /*app*/ ctx[4],
    				alt_methods: /*alt_methods*/ ctx[1],
    				opensignup: /*opensignup*/ ctx[3],
    				password: /*password*/ ctx[2]
    			}
    		});

    	return {
    		c() {
    			create_component(inner.$$.fragment);
    		},
    		m(target, anchor) {
    			mount_component(inner, target, anchor);
    			current = true;
    		},
    		p(ctx, dirty) {
    			const inner_changes = {};
    			if (dirty & /*alt_methods*/ 2) inner_changes.alt_methods = /*alt_methods*/ ctx[1];
    			if (dirty & /*opensignup*/ 8) inner_changes.opensignup = /*opensignup*/ ctx[3];
    			if (dirty & /*password*/ 4) inner_changes.password = /*password*/ ctx[2];
    			inner.$set(inner_changes);
    		},
    		i(local) {
    			if (current) return;
    			transition_in(inner.$$.fragment, local);
    			current = true;
    		},
    		o(local) {
    			transition_out(inner.$$.fragment, local);
    			current = false;
    		},
    		d(detaching) {
    			destroy_component(inner, detaching);
    		}
    	};
    }

    // (18:0) <AuthLayout>
    function create_default_slot$3(ctx) {
    	let current_block_type_index;
    	let if_block;
    	let if_block_anchor;
    	let current;
    	const if_block_creators = [create_if_block$5, create_else_block$1];
    	const if_blocks = [];

    	function select_block_type(ctx, dirty) {
    		if (/*loaded*/ ctx[0]) return 0;
    		return 1;
    	}

    	current_block_type_index = select_block_type(ctx);
    	if_block = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx);

    	return {
    		c() {
    			if_block.c();
    			if_block_anchor = empty();
    		},
    		m(target, anchor) {
    			if_blocks[current_block_type_index].m(target, anchor);
    			insert(target, if_block_anchor, anchor);
    			current = true;
    		},
    		p(ctx, dirty) {
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
    		i(local) {
    			if (current) return;
    			transition_in(if_block);
    			current = true;
    		},
    		o(local) {
    			transition_out(if_block);
    			current = false;
    		},
    		d(detaching) {
    			if_blocks[current_block_type_index].d(detaching);
    			if (detaching) detach(if_block_anchor);
    		}
    	};
    }

    function create_fragment$d(ctx) {
    	let authlayout;
    	let current;

    	authlayout = new Layout({
    			props: {
    				$$slots: { default: [create_default_slot$3] },
    				$$scope: { ctx }
    			}
    		});

    	return {
    		c() {
    			create_component(authlayout.$$.fragment);
    		},
    		m(target, anchor) {
    			mount_component(authlayout, target, anchor);
    			current = true;
    		},
    		p(ctx, [dirty]) {
    			const authlayout_changes = {};

    			if (dirty & /*$$scope, alt_methods, opensignup, password, loaded*/ 47) {
    				authlayout_changes.$$scope = { dirty, ctx };
    			}

    			authlayout.$set(authlayout_changes);
    		},
    		i(local) {
    			if (current) return;
    			transition_in(authlayout.$$.fragment, local);
    			current = true;
    		},
    		o(local) {
    			transition_out(authlayout.$$.fragment, local);
    			current = false;
    		},
    		d(detaching) {
    			destroy_component(authlayout, detaching);
    		}
    	};
    }

    function instance$8($$self, $$props, $$invalidate) {
    	const app = getContext("_auth_app_");
    	let loaded = false;
    	let alt_methods = [];
    	let password = false;
    	let opensignup = false;

    	(async () => {
    		const resp = await app.list_methods();
    		$$invalidate(1, alt_methods = resp.alt_auth_method);
    		$$invalidate(2, password = resp.pass_auth);
    		$$invalidate(3, opensignup = resp.open_signup);
    		$$invalidate(0, loaded = true);
    	})();

    	return [loaded, alt_methods, password, opensignup, app];
    }

    class Login extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, instance$8, create_fragment$d, safe_not_equal, {});
    	}
    }

    /* src/entries/auth/pages/login/nextstage/index.svelte generated by Svelte v3.59.2 */

    function create_if_block$4(ctx) {
    	let div;

    	return {
    		c() {
    			div = element("div");
    			div.textContent = "fixme";
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    		},
    		d(detaching) {
    			if (detaching) detach(div);
    		}
    	};
    }

    // (27:0) <Layout>
    function create_default_slot$2(ctx) {
    	let p;
    	let t0;
    	let t1;
    	let if_block_anchor;
    	let if_block = /*show*/ ctx[0] && create_if_block$4();

    	return {
    		c() {
    			p = element("p");
    			t0 = text(/*message*/ ctx[1]);
    			t1 = space();
    			if (if_block) if_block.c();
    			if_block_anchor = empty();
    		},
    		m(target, anchor) {
    			insert(target, p, anchor);
    			append(p, t0);
    			insert(target, t1, anchor);
    			if (if_block) if_block.m(target, anchor);
    			insert(target, if_block_anchor, anchor);
    		},
    		p(ctx, dirty) {
    			if (dirty & /*message*/ 2) set_data(t0, /*message*/ ctx[1]);

    			if (/*show*/ ctx[0]) {
    				if (if_block) ; else {
    					if_block = create_if_block$4();
    					if_block.c();
    					if_block.m(if_block_anchor.parentNode, if_block_anchor);
    				}
    			} else if (if_block) {
    				if_block.d(1);
    				if_block = null;
    			}
    		},
    		d(detaching) {
    			if (detaching) detach(p);
    			if (detaching) detach(t1);
    			if (if_block) if_block.d(detaching);
    			if (detaching) detach(if_block_anchor);
    		}
    	};
    }

    function create_fragment$c(ctx) {
    	let layout;
    	let current;

    	layout = new Layout({
    			props: {
    				$$slots: { default: [create_default_slot$2] },
    				$$scope: { ctx }
    			}
    		});

    	return {
    		c() {
    			create_component(layout.$$.fragment);
    		},
    		m(target, anchor) {
    			mount_component(layout, target, anchor);
    			current = true;
    		},
    		p(ctx, [dirty]) {
    			const layout_changes = {};

    			if (dirty & /*$$scope, show, message*/ 19) {
    				layout_changes.$$scope = { dirty, ctx };
    			}

    			layout.$set(layout_changes);
    		},
    		i(local) {
    			if (current) return;
    			transition_in(layout.$$.fragment, local);
    			current = true;
    		},
    		o(local) {
    			transition_out(layout.$$.fragment, local);
    			current = false;
    		},
    		d(detaching) {
    			destroy_component(layout, detaching);
    		}
    	};
    }

    function instance$7($$self, $$props, $$invalidate) {
    	const app = getContext("_auth_app_");
    	const opts = app.nav.nav_options;
    	let show = false;
    	let message = "";

    	(async () => {
    		if (!opts.email_verify && !opts.pass_change) {
    			const resp = await app.login_submit(opts.next_token);

    			if (resp.status !== 200) {
    				console.log("Err =>", resp);
    				return;
    			}

    			if (!resp.data["ok"]) {
    				$$invalidate(1, message = resp.data["message"]);
    				return;
    			}

    			app.save_preauthed_data(resp.data);
    			app.nav.goto_prehook_page(resp.data);
    		} else {
    			$$invalidate(0, show = true);
    		}
    	})();

    	return [show, message];
    }

    class Nextstage$1 extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, instance$7, create_fragment$c, safe_not_equal, {});
    	}
    }

    /* src/entries/auth/pages/common/new_user_info.svelte generated by Svelte v3.59.2 */

    function get_each_context(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[12] = list[i];
    	return child_ctx;
    }

    // (43:6) {#each user_id_hints as hint}
    function create_each_block(ctx) {
    	let a;
    	let t_value = /*hint*/ ctx[12] + "";
    	let t;
    	let mounted;
    	let dispose;

    	function click_handler() {
    		return /*click_handler*/ ctx[8](/*hint*/ ctx[12]);
    	}

    	return {
    		c() {
    			a = element("a");
    			t = text(t_value);
    			attr(a, "class", "text-blue-700");
    			attr(a, "href", "#");
    			attr(a, "target", "_blank");
    		},
    		m(target, anchor) {
    			insert(target, a, anchor);
    			append(a, t);

    			if (!mounted) {
    				dispose = listen(a, "click", prevent_default(click_handler));
    				mounted = true;
    			}
    		},
    		p(new_ctx, dirty) {
    			ctx = new_ctx;
    			if (dirty & /*user_id_hints*/ 16 && t_value !== (t_value = /*hint*/ ctx[12] + "")) set_data(t, t_value);
    		},
    		d(detaching) {
    			if (detaching) detach(a);
    			mounted = false;
    			dispose();
    		}
    	};
    }

    function create_fragment$b(ctx) {
    	let div5;
    	let h3;
    	let t1;
    	let div1;
    	let label0;
    	let t3;
    	let input0;
    	let t4;
    	let div0;
    	let t5;
    	let div2;
    	let label1;
    	let t7;
    	let input1;
    	let t8;
    	let div3;
    	let label2;
    	let t10;
    	let input2;
    	let t11;
    	let div4;
    	let label3;
    	let t13;
    	let textarea;
    	let t14;
    	let button;
    	let mounted;
    	let dispose;
    	let each_value = /*user_id_hints*/ ctx[4];
    	let each_blocks = [];

    	for (let i = 0; i < each_value.length; i += 1) {
    		each_blocks[i] = create_each_block(get_each_context(ctx, each_value, i));
    	}

    	return {
    		c() {
    			div5 = element("div");
    			h3 = element("h3");
    			h3.textContent = "Fill in your information.";
    			t1 = space();
    			div1 = element("div");
    			label0 = element("label");
    			label0.textContent = "User Id";
    			t3 = space();
    			input0 = element("input");
    			t4 = space();
    			div0 = element("div");

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].c();
    			}

    			t5 = space();
    			div2 = element("div");
    			label1 = element("label");
    			label1.textContent = "Email";
    			t7 = space();
    			input1 = element("input");
    			t8 = space();
    			div3 = element("div");
    			label2 = element("label");
    			label2.textContent = "Full Name";
    			t10 = space();
    			input2 = element("input");
    			t11 = space();
    			div4 = element("div");
    			label3 = element("label");
    			label3.textContent = "Bio";
    			t13 = space();
    			textarea = element("textarea");
    			t14 = space();
    			button = element("button");
    			button.textContent = "Next";
    			attr(h3, "class", "text-xl font-medium text-gray-900 mb-2");
    			attr(label0, "for", "userid");
    			attr(label0, "class", "text-sm font-medium text-gray-900 block mb-2");
    			attr(input0, "type", "text");
    			attr(input0, "name", "userid");
    			attr(input0, "id", "userid");
    			attr(input0, "class", "bg-gray-100 focus:bg-white border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 ");
    			attr(input0, "placeholder", "john12");
    			attr(div0, "class", "flex gap-2");
    			attr(label1, "for", "email");
    			attr(label1, "class", "text-sm font-medium text-gray-900 block mb-2 mt-4 ");
    			attr(input1, "type", "text");
    			attr(input1, "name", "email");
    			attr(input1, "id", "email");
    			input1.disabled = true;
    			attr(input1, "class", "bg-gray-100 focus:bg-white border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 ");
    			attr(label2, "for", "userid");
    			attr(label2, "class", "text-sm font-medium text-gray-900 block mb-2 mt-4 ");
    			attr(input2, "type", "text");
    			attr(input2, "name", "name");
    			attr(input2, "id", "name");
    			attr(input2, "class", "bg-gray-100 focus:bg-white border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 ");
    			attr(input2, "placeholder", "John Doe");
    			attr(label3, "for", "bio");
    			attr(label3, "class", "text-sm font-medium text-gray-900 block mb-2 mt-4");
    			attr(textarea, "name", "bio");
    			attr(textarea, "id", "bio");
    			attr(textarea, "class", "bg-gray-100 focus:bg-white border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 ");
    			attr(textarea, "placeholder", "i am ...");
    			attr(button, "class", "w-full px-4 py-2 tracking-wide text-white transition-colors duration-200 transform bg-blue-700 rounded hover:bg-blue-400 mt-4");
    		},
    		m(target, anchor) {
    			insert(target, div5, anchor);
    			append(div5, h3);
    			append(div5, t1);
    			append(div5, div1);
    			append(div1, label0);
    			append(div1, t3);
    			append(div1, input0);
    			set_input_value(input0, /*user_id*/ ctx[0]);
    			append(div1, t4);
    			append(div1, div0);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				if (each_blocks[i]) {
    					each_blocks[i].m(div0, null);
    				}
    			}

    			append(div5, t5);
    			append(div5, div2);
    			append(div2, label1);
    			append(div2, t7);
    			append(div2, input1);
    			set_input_value(input1, /*email*/ ctx[3]);
    			append(div5, t8);
    			append(div5, div3);
    			append(div3, label2);
    			append(div3, t10);
    			append(div3, input2);
    			set_input_value(input2, /*full_name*/ ctx[1]);
    			append(div5, t11);
    			append(div5, div4);
    			append(div4, label3);
    			append(div4, t13);
    			append(div4, textarea);
    			set_input_value(textarea, /*bio*/ ctx[2]);
    			append(div5, t14);
    			append(div5, button);

    			if (!mounted) {
    				dispose = [
    					listen(input0, "input", /*input0_input_handler*/ ctx[7]),
    					listen(input1, "input", /*input1_input_handler*/ ctx[9]),
    					listen(input2, "input", /*input2_input_handler*/ ctx[10]),
    					listen(textarea, "input", /*textarea_input_handler*/ ctx[11]),
    					listen(button, "click", /*applyNext*/ ctx[5])
    				];

    				mounted = true;
    			}
    		},
    		p(ctx, [dirty]) {
    			if (dirty & /*user_id*/ 1 && input0.value !== /*user_id*/ ctx[0]) {
    				set_input_value(input0, /*user_id*/ ctx[0]);
    			}

    			if (dirty & /*user_id, user_id_hints*/ 17) {
    				each_value = /*user_id_hints*/ ctx[4];
    				let i;

    				for (i = 0; i < each_value.length; i += 1) {
    					const child_ctx = get_each_context(ctx, each_value, i);

    					if (each_blocks[i]) {
    						each_blocks[i].p(child_ctx, dirty);
    					} else {
    						each_blocks[i] = create_each_block(child_ctx);
    						each_blocks[i].c();
    						each_blocks[i].m(div0, null);
    					}
    				}

    				for (; i < each_blocks.length; i += 1) {
    					each_blocks[i].d(1);
    				}

    				each_blocks.length = each_value.length;
    			}

    			if (dirty & /*email*/ 8 && input1.value !== /*email*/ ctx[3]) {
    				set_input_value(input1, /*email*/ ctx[3]);
    			}

    			if (dirty & /*full_name*/ 2 && input2.value !== /*full_name*/ ctx[1]) {
    				set_input_value(input2, /*full_name*/ ctx[1]);
    			}

    			if (dirty & /*bio*/ 4) {
    				set_input_value(textarea, /*bio*/ ctx[2]);
    			}
    		},
    		i: noop,
    		o: noop,
    		d(detaching) {
    			if (detaching) detach(div5);
    			destroy_each(each_blocks, detaching);
    			mounted = false;
    			run_all(dispose);
    		}
    	};
    }

    function instance$6($$self, $$props, $$invalidate) {
    	let { user_id_hints = ["john12", "john32", "ram13", "cena66", "philip122"] } = $$props;
    	let { user_id = "" } = $$props;
    	let { full_name = "" } = $$props;
    	let { bio = "" } = $$props;
    	let { email } = $$props;
    	let { onNext } = $$props;

    	const applyNext = () => {
    		onNext({ user_id, full_name, bio });
    	};

    	function input0_input_handler() {
    		user_id = this.value;
    		$$invalidate(0, user_id);
    	}

    	const click_handler = hint => {
    		$$invalidate(0, user_id = hint);
    	};

    	function input1_input_handler() {
    		email = this.value;
    		$$invalidate(3, email);
    	}

    	function input2_input_handler() {
    		full_name = this.value;
    		$$invalidate(1, full_name);
    	}

    	function textarea_input_handler() {
    		bio = this.value;
    		$$invalidate(2, bio);
    	}

    	$$self.$$set = $$props => {
    		if ('user_id_hints' in $$props) $$invalidate(4, user_id_hints = $$props.user_id_hints);
    		if ('user_id' in $$props) $$invalidate(0, user_id = $$props.user_id);
    		if ('full_name' in $$props) $$invalidate(1, full_name = $$props.full_name);
    		if ('bio' in $$props) $$invalidate(2, bio = $$props.bio);
    		if ('email' in $$props) $$invalidate(3, email = $$props.email);
    		if ('onNext' in $$props) $$invalidate(6, onNext = $$props.onNext);
    	};

    	return [
    		user_id,
    		full_name,
    		bio,
    		email,
    		user_id_hints,
    		applyNext,
    		onNext,
    		input0_input_handler,
    		click_handler,
    		input1_input_handler,
    		input2_input_handler,
    		textarea_input_handler
    	];
    }

    class New_user_info extends SvelteComponent {
    	constructor(options) {
    		super();

    		init(this, options, instance$6, create_fragment$b, safe_not_equal, {
    			user_id_hints: 4,
    			user_id: 0,
    			full_name: 1,
    			bio: 2,
    			email: 3,
    			onNext: 6
    		});
    	}
    }

    /* src/entries/auth/pages/alt/firststage/index.svelte generated by Svelte v3.59.2 */

    function create_if_block$3(ctx) {
    	let layout;
    	let current;

    	layout = new Layout({
    			props: {
    				$$slots: { default: [create_default_slot$1] },
    				$$scope: { ctx }
    			}
    		});

    	return {
    		c() {
    			create_component(layout.$$.fragment);
    		},
    		m(target, anchor) {
    			mount_component(layout, target, anchor);
    			current = true;
    		},
    		p(ctx, dirty) {
    			const layout_changes = {};

    			if (dirty & /*$$scope*/ 8) {
    				layout_changes.$$scope = { dirty, ctx };
    			}

    			layout.$set(layout_changes);
    		},
    		i(local) {
    			if (current) return;
    			transition_in(layout.$$.fragment, local);
    			current = true;
    		},
    		o(local) {
    			transition_out(layout.$$.fragment, local);
    			current = false;
    		},
    		d(detaching) {
    			destroy_component(layout, detaching);
    		}
    	};
    }

    // (23:2) <Layout>
    function create_default_slot$1(ctx) {
    	let newuserinfo;
    	let current;

    	newuserinfo = new New_user_info({
    			props: {
    				user_id_hints: /*opts*/ ctx[1].user_id_hints,
    				email: /*opts*/ ctx[1].email,
    				onNext: /*func*/ ctx[2]
    			}
    		});

    	return {
    		c() {
    			create_component(newuserinfo.$$.fragment);
    		},
    		m(target, anchor) {
    			mount_component(newuserinfo, target, anchor);
    			current = true;
    		},
    		p: noop,
    		i(local) {
    			if (current) return;
    			transition_in(newuserinfo.$$.fragment, local);
    			current = true;
    		},
    		o(local) {
    			transition_out(newuserinfo.$$.fragment, local);
    			current = false;
    		},
    		d(detaching) {
    			destroy_component(newuserinfo, detaching);
    		}
    	};
    }

    function create_fragment$a(ctx) {
    	let if_block_anchor;
    	let current;
    	let if_block = /*opts*/ ctx[1].new_user && create_if_block$3(ctx);

    	return {
    		c() {
    			if (if_block) if_block.c();
    			if_block_anchor = empty();
    		},
    		m(target, anchor) {
    			if (if_block) if_block.m(target, anchor);
    			insert(target, if_block_anchor, anchor);
    			current = true;
    		},
    		p(ctx, [dirty]) {
    			if (/*opts*/ ctx[1].new_user) if_block.p(ctx, dirty);
    		},
    		i(local) {
    			if (current) return;
    			transition_in(if_block);
    			current = true;
    		},
    		o(local) {
    			transition_out(if_block);
    			current = false;
    		},
    		d(detaching) {
    			if (if_block) if_block.d(detaching);
    			if (detaching) detach(if_block_anchor);
    		}
    	};
    }

    function instance$5($$self) {
    	const app = getContext("_auth_app_");
    	const opts = app.nav.nav_options;

    	if (!opts) {
    		app.nav.goto_login_page();
    	}

    	(async () => {
    		if (opts.new_user) {
    			return;
    		}

    		const resp = await app.alt_next_second(opts.first_token);

    		if (resp.status !== 200) {
    			app.nav.goto_error_page(resp.data);
    			return;
    		}

    		app.nav.goto_alt_second_stage(Object.assign({ email: opts.email }, resp.data));
    	})();

    	const func = async data => {
    		const resp = await app.alt_next_second(opts.first_token, data);

    		if (resp.status !== 200) {
    			app.nav.goto_error_page(resp.data);
    			return;
    		}

    		app.nav.goto_alt_second_stage({ email: opts.email, ...resp.data });
    	};

    	return [app, opts, func];
    }

    class Firststage extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, instance$5, create_fragment$a, safe_not_equal, {});
    	}
    }

    /* src/entries/auth/pages/alt/secondstage/index.svelte generated by Svelte v3.59.2 */

    function create_if_block$2(ctx) {
    	let layout;
    	let current;

    	layout = new Layout({
    			props: {
    				$$slots: { default: [create_default_slot] },
    				$$scope: { ctx }
    			}
    		});

    	return {
    		c() {
    			create_component(layout.$$.fragment);
    		},
    		m(target, anchor) {
    			mount_component(layout, target, anchor);
    			current = true;
    		},
    		p(ctx, dirty) {
    			const layout_changes = {};

    			if (dirty & /*$$scope, code*/ 17) {
    				layout_changes.$$scope = { dirty, ctx };
    			}

    			layout.$set(layout_changes);
    		},
    		i(local) {
    			if (current) return;
    			transition_in(layout.$$.fragment, local);
    			current = true;
    		},
    		o(local) {
    			transition_out(layout.$$.fragment, local);
    			current = false;
    		},
    		d(detaching) {
    			destroy_component(layout, detaching);
    		}
    	};
    }

    // (23:2) <Layout>
    function create_default_slot(ctx) {
    	let div;
    	let label;
    	let t0;
    	let span;
    	let t2;
    	let input;
    	let t3;
    	let button;
    	let mounted;
    	let dispose;

    	return {
    		c() {
    			div = element("div");
    			label = element("label");
    			t0 = text("Verify your email: ");
    			span = element("span");
    			span.textContent = `${/*opts*/ ctx[1].email}`;
    			t2 = space();
    			input = element("input");
    			t3 = space();
    			button = element("button");
    			button.textContent = "Verify";
    			attr(span, "class", "text-blue-700");
    			attr(label, "for", "code");
    			attr(label, "class", "text-sm font-medium text-gray-900 block mb-2 mt-4 ");
    			attr(input, "type", "text");
    			attr(input, "name", "code");
    			attr(input, "id", "code");
    			attr(input, "class", "bg-gray-100 h-20 focus:bg-white border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 ");
    			attr(button, "class", "w-full px-4 py-2 tracking-wide text-white transition-colors duration-200 transform bg-blue-700 rounded hover:bg-blue-400 mt-4");
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    			append(div, label);
    			append(label, t0);
    			append(label, span);
    			append(div, t2);
    			append(div, input);
    			set_input_value(input, /*code*/ ctx[0]);
    			insert(target, t3, anchor);
    			insert(target, button, anchor);

    			if (!mounted) {
    				dispose = listen(input, "input", /*input_input_handler*/ ctx[2]);
    				mounted = true;
    			}
    		},
    		p(ctx, dirty) {
    			if (dirty & /*code*/ 1 && input.value !== /*code*/ ctx[0]) {
    				set_input_value(input, /*code*/ ctx[0]);
    			}
    		},
    		d(detaching) {
    			if (detaching) detach(div);
    			if (detaching) detach(t3);
    			if (detaching) detach(button);
    			mounted = false;
    			dispose();
    		}
    	};
    }

    function create_fragment$9(ctx) {
    	let if_block_anchor;
    	let current;
    	let if_block = /*opts*/ ctx[1].email_verify && create_if_block$2(ctx);

    	return {
    		c() {
    			if (if_block) if_block.c();
    			if_block_anchor = empty();
    		},
    		m(target, anchor) {
    			if (if_block) if_block.m(target, anchor);
    			insert(target, if_block_anchor, anchor);
    			current = true;
    		},
    		p(ctx, [dirty]) {
    			if (/*opts*/ ctx[1].email_verify) if_block.p(ctx, dirty);
    		},
    		i(local) {
    			if (current) return;
    			transition_in(if_block);
    			current = true;
    		},
    		o(local) {
    			transition_out(if_block);
    			current = false;
    		},
    		d(detaching) {
    			if (if_block) if_block.d(detaching);
    			if (detaching) detach(if_block_anchor);
    		}
    	};
    }

    function instance$4($$self, $$props, $$invalidate) {
    	const app = getContext("_auth_app_");
    	const opts = app.nav.nav_options;

    	if (!opts) {
    		app.nav.goto_login_page();
    	}

    	let code = "";

    	(async () => {
    		if (!opts.email_verify) {
    			const resp = await app.submit_alt_auth(opts.next_token);

    			if (resp.status !== 200) {
    				app.nav.goto_error_page(resp.data);
    				return;
    			}

    			app.save_preauthed_data(resp.data);
    			app.nav.goto_prehook_page(resp.data);
    		}
    	})();

    	function input_input_handler() {
    		code = this.value;
    		$$invalidate(0, code);
    	}

    	return [code, opts, input_input_handler];
    }

    class Secondstage extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, instance$4, create_fragment$9, safe_not_equal, {});
    	}
    }

    /* src/entries/auth/pages/signup/index.svelte generated by Svelte v3.59.2 */

    function create_fragment$8(ctx) {
    	let div;

    	return {
    		c() {
    			div = element("div");
    			div.textContent = "Signup";
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d(detaching) {
    			if (detaching) detach(div);
    		}
    	};
    }

    class Signup extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, null, create_fragment$8, safe_not_equal, {});
    	}
    }

    /* src/entries/auth/pages/signup/nextstage/index.svelte generated by Svelte v3.59.2 */

    function create_fragment$7(ctx) {
    	let div;

    	return {
    		c() {
    			div = element("div");
    			div.textContent = "Next Stage";
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d(detaching) {
    			if (detaching) detach(div);
    		}
    	};
    }

    class Nextstage extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, null, create_fragment$7, safe_not_equal, {});
    	}
    }

    /* src/entries/auth/pages/reset/index.svelte generated by Svelte v3.59.2 */

    function create_fragment$6(ctx) {
    	let div;

    	return {
    		c() {
    			div = element("div");
    			div.textContent = "Reset";
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d(detaching) {
    			if (detaching) detach(div);
    		}
    	};
    }

    class Reset extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, null, create_fragment$6, safe_not_equal, {});
    	}
    }

    /* src/entries/auth/pages/reset/finish/index.svelte generated by Svelte v3.59.2 */

    function create_fragment$5(ctx) {
    	let div;

    	return {
    		c() {
    			div = element("div");
    			div.textContent = "Reset Finish";
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d(detaching) {
    			if (detaching) detach(div);
    		}
    	};
    }

    class Finish extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, null, create_fragment$5, safe_not_equal, {});
    	}
    }

    /* src/entries/auth/pages/prehook/index.svelte generated by Svelte v3.59.2 */

    function create_fragment$4(ctx) {
    	let div;

    	return {
    		c() {
    			div = element("div");
    			div.textContent = "Pre Hook Page";
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d(detaching) {
    			if (detaching) detach(div);
    		}
    	};
    }

    function instance$3($$self) {
    	const app = getContext("_auth_app_");
    	const opts = app.get_preauthed_data();
    	console.log("@opts =>", opts);

    	(async () => {
    		if (!opts.has_exec_hook) {
    			const resp = await app.login_finish(opts.preauthed_token);

    			if (resp.status !== 200) {
    				console.log("Err =>", resp);
    				return;
    			}

    			app.save_authed_data(resp.data["user_token"]);
    			app.nav.goto_final_page();
    			return;
    		}

    		console.log("TODO RUN HOOK");
    	})();

    	return [];
    }

    class Prehook extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, instance$3, create_fragment$4, safe_not_equal, {});
    	}
    }

    /* src/entries/auth/pages/common/user_card.svelte generated by Svelte v3.59.2 */

    function create_if_block_1(ctx) {
    	let span;
    	let t0;
    	let a;
    	let t1;
    	let t2;
    	let t3;
    	let t4;
    	let button;
    	let mounted;
    	let dispose;

    	return {
    		c() {
    			span = element("span");
    			t0 = text(", Redirecting ");
    			a = element("a");
    			t1 = text("in ");
    			t2 = text(/*seconds*/ ctx[8]);
    			t3 = text(" seconds");
    			t4 = space();
    			button = element("button");
    			button.textContent = "Cancel";
    			attr(a, "href", /*return_url*/ ctx[6]);
    			attr(span, "class", "text-md");
    			attr(button, "class", "text-blue-600");
    		},
    		m(target, anchor) {
    			insert(target, span, anchor);
    			append(span, t0);
    			append(span, a);
    			append(a, t1);
    			append(a, t2);
    			append(a, t3);
    			insert(target, t4, anchor);
    			insert(target, button, anchor);

    			if (!mounted) {
    				dispose = listen(button, "click", /*cancel*/ ctx[11]);
    				mounted = true;
    			}
    		},
    		p(ctx, dirty) {
    			if (dirty & /*seconds*/ 256) set_data(t2, /*seconds*/ ctx[8]);

    			if (dirty & /*return_url*/ 64) {
    				attr(a, "href", /*return_url*/ ctx[6]);
    			}
    		},
    		d(detaching) {
    			if (detaching) detach(span);
    			if (detaching) detach(t4);
    			if (detaching) detach(button);
    			mounted = false;
    			dispose();
    		}
    	};
    }

    // (84:8) {#if return_url}
    function create_if_block$1(ctx) {
    	let button;
    	let mounted;
    	let dispose;

    	return {
    		c() {
    			button = element("button");
    			button.textContent = "Go Back";
    			attr(button, "class", "p-2 bg-blue-400 hover:bg-blue-600 text-white font-semibold rounded");
    		},
    		m(target, anchor) {
    			insert(target, button, anchor);

    			if (!mounted) {
    				dispose = listen(button, "click", /*click_handler*/ ctx[13]);
    				mounted = true;
    			}
    		},
    		p: noop,
    		d(detaching) {
    			if (detaching) detach(button);
    			mounted = false;
    			dispose();
    		}
    	};
    }

    function create_fragment$3(ctx) {
    	let div6;
    	let div5;
    	let div4;
    	let div0;
    	let t0;
    	let t1;
    	let div1;
    	let img;
    	let img_src_value;
    	let t2;
    	let h1;
    	let t3;
    	let t4;
    	let h3;
    	let t5;
    	let t6;
    	let p;
    	let t7;
    	let t8;
    	let ul;
    	let li0;
    	let span0;
    	let t10;
    	let span1;
    	let t11;
    	let t12;
    	let li1;
    	let span2;
    	let t14;
    	let span3;
    	let t15_value = /*tenant_name*/ ctx[0] + " [" + /*tenant_id*/ ctx[1] + "]" + "";
    	let t15;
    	let t16;
    	let div2;
    	let t17;
    	let button0;
    	let t18;
    	let t19;
    	let t20;
    	let t21;
    	let div3;
    	let button1;
    	let t23;
    	let button2;
    	let mounted;
    	let dispose;
    	let if_block0 = /*show_timeout*/ ctx[7] && create_if_block_1(ctx);
    	let if_block1 = /*return_url*/ ctx[6] && create_if_block$1(ctx);

    	return {
    		c() {
    			div6 = element("div");
    			div5 = element("div");
    			div4 = element("div");
    			div0 = element("div");
    			t0 = text("You are Logged in\n\n        ");
    			if (if_block0) if_block0.c();
    			t1 = space();
    			div1 = element("div");
    			img = element("img");
    			t2 = space();
    			h1 = element("h1");
    			t3 = text(/*full_name*/ ctx[3]);
    			t4 = space();
    			h3 = element("h3");
    			t5 = text(/*group_name*/ ctx[4]);
    			t6 = space();
    			p = element("p");
    			t7 = text(/*bio*/ ctx[5]);
    			t8 = space();
    			ul = element("ul");
    			li0 = element("li");
    			span0 = element("span");
    			span0.textContent = "User Id";
    			t10 = space();
    			span1 = element("span");
    			t11 = text(/*user_id*/ ctx[2]);
    			t12 = space();
    			li1 = element("li");
    			span2 = element("span");
    			span2.textContent = "Organization";
    			t14 = space();
    			span3 = element("span");
    			t15 = text(t15_value);
    			t16 = space();
    			div2 = element("div");
    			if (if_block1) if_block1.c();
    			t17 = space();
    			button0 = element("button");
    			t18 = text("Go to Portal ( ");
    			t19 = text(/*seconds*/ ctx[8]);
    			t20 = text(" )");
    			t21 = space();
    			div3 = element("div");
    			button1 = element("button");
    			button1.textContent = "Home";
    			t23 = space();
    			button2 = element("button");
    			button2.textContent = "Logout";
    			attr(div0, "class", "rounded bg-green-500 text-white w-full p-1 mb-2");
    			attr(img, "class", "h-auto w-12 mx-auto rounded-full border");
    			if (!src_url_equal(img.src, img_src_value = `${apiURL(/*tenant_id*/ ctx[1])}/user_profile_image/${/*user_id*/ ctx[2]}`)) attr(img, "src", img_src_value);
    			attr(img, "alt", "");
    			attr(div1, "class", "image overflow-hidden");
    			attr(h1, "class", "text-gray-900 font-bold text-xl leading-8 my-1");
    			attr(h3, "class", "text-gray-600 font-lg text-semibold leading-6");
    			attr(p, "class", "text-sm text-gray-500 hover:text-gray-600 leading-6");
    			attr(span1, "class", "ml-auto bg-gray-300 rounded p-1");
    			attr(li0, "class", "flex items-center py-3");
    			attr(span3, "class", "ml-auto bg-gray-300 rounded p-1");
    			attr(li1, "class", "flex items-center py-3");
    			attr(ul, "class", "bg-gray-100 text-gray-600 hover:text-gray-700 hover:shadow py-2 px-3 mt-3 divide-y rounded shadow-sm");
    			attr(button0, "class", "p-2 bg-blue-400 hover:bg-green-600 text-white font-semibold rounded");
    			attr(div2, "class", "flex flex-col gap-1 mt-1");
    			attr(button1, "class", "p-2 bg-blue-400 hover:bg-blue-600 text-white font-semibold rounded");
    			attr(button2, "class", "p-2 bg-red-400 hover:bg-red-600 text-white font-semibold rounded");
    			attr(div3, "class", "flex gap-1 mt-1 justify-between justify-items-stretch");
    			attr(div4, "class", "bg-white border rounded p-4");
    			attr(div5, "class", "my-5 mx-auto border");
    			set_style(div5, "max-width", "500px");
    			attr(div6, "class", "w-screen h-screen bg-gradient-to-r from-teal-400 to-yellow-200 py-10");
    		},
    		m(target, anchor) {
    			insert(target, div6, anchor);
    			append(div6, div5);
    			append(div5, div4);
    			append(div4, div0);
    			append(div0, t0);
    			if (if_block0) if_block0.m(div0, null);
    			append(div4, t1);
    			append(div4, div1);
    			append(div1, img);
    			append(div4, t2);
    			append(div4, h1);
    			append(h1, t3);
    			append(div4, t4);
    			append(div4, h3);
    			append(h3, t5);
    			append(div4, t6);
    			append(div4, p);
    			append(p, t7);
    			append(div4, t8);
    			append(div4, ul);
    			append(ul, li0);
    			append(li0, span0);
    			append(li0, t10);
    			append(li0, span1);
    			append(span1, t11);
    			append(ul, t12);
    			append(ul, li1);
    			append(li1, span2);
    			append(li1, t14);
    			append(li1, span3);
    			append(span3, t15);
    			append(div4, t16);
    			append(div4, div2);
    			if (if_block1) if_block1.m(div2, null);
    			append(div2, t17);
    			append(div2, button0);
    			append(button0, t18);
    			append(button0, t19);
    			append(button0, t20);
    			append(div4, t21);
    			append(div4, div3);
    			append(div3, button1);
    			append(div3, t23);
    			append(div3, button2);

    			if (!mounted) {
    				dispose = [
    					listen(button0, "click", /*click_handler_1*/ ctx[14]),
    					listen(button1, "click", /*click_handler_2*/ ctx[15]),
    					listen(button2, "click", /*click_handler_3*/ ctx[16])
    				];

    				mounted = true;
    			}
    		},
    		p(ctx, [dirty]) {
    			if (/*show_timeout*/ ctx[7]) {
    				if (if_block0) {
    					if_block0.p(ctx, dirty);
    				} else {
    					if_block0 = create_if_block_1(ctx);
    					if_block0.c();
    					if_block0.m(div0, null);
    				}
    			} else if (if_block0) {
    				if_block0.d(1);
    				if_block0 = null;
    			}

    			if (dirty & /*tenant_id, user_id*/ 6 && !src_url_equal(img.src, img_src_value = `${apiURL(/*tenant_id*/ ctx[1])}/user_profile_image/${/*user_id*/ ctx[2]}`)) {
    				attr(img, "src", img_src_value);
    			}

    			if (dirty & /*full_name*/ 8) set_data(t3, /*full_name*/ ctx[3]);
    			if (dirty & /*group_name*/ 16) set_data(t5, /*group_name*/ ctx[4]);
    			if (dirty & /*bio*/ 32) set_data(t7, /*bio*/ ctx[5]);
    			if (dirty & /*user_id*/ 4) set_data(t11, /*user_id*/ ctx[2]);
    			if (dirty & /*tenant_name, tenant_id*/ 3 && t15_value !== (t15_value = /*tenant_name*/ ctx[0] + " [" + /*tenant_id*/ ctx[1] + "]" + "")) set_data(t15, t15_value);

    			if (/*return_url*/ ctx[6]) {
    				if (if_block1) {
    					if_block1.p(ctx, dirty);
    				} else {
    					if_block1 = create_if_block$1(ctx);
    					if_block1.c();
    					if_block1.m(div2, t17);
    				}
    			} else if (if_block1) {
    				if_block1.d(1);
    				if_block1 = null;
    			}

    			if (dirty & /*seconds*/ 256) set_data(t19, /*seconds*/ ctx[8]);
    		},
    		i: noop,
    		o: noop,
    		d(detaching) {
    			if (detaching) detach(div6);
    			if (if_block0) if_block0.d();
    			if (if_block1) if_block1.d();
    			mounted = false;
    			run_all(dispose);
    		}
    	};
    }

    function instance$2($$self, $$props, $$invalidate) {
    	let { tenant_name } = $$props;
    	let { tenant_id } = $$props;
    	let { user_id } = $$props;
    	let { full_name } = $$props;
    	let { group_name } = $$props;
    	let { bio } = $$props;
    	let { return_url = undefined } = $$props;
    	let portal_url = `${window.location.origin}/z/portal`;
    	let home_url = window.location.origin;
    	let show_timeout = true;
    	let seconds = 5;

    	const it = setInterval(
    		() => {
    			if (seconds <= 0) {
    				clearInterval(it);

    				if (return_url) {
    					window.location.href = return_url;
    				} else {
    					window.location.href = portal_url;
    				}

    				return;
    			}

    			$$invalidate(8, seconds -= 1);
    		},
    		1000
    	);

    	const cancel = () => {
    		clearInterval(it);
    		$$invalidate(7, show_timeout = false);
    	};

    	const app = getContext("_auth_app_");

    	const click_handler = () => {
    		cancel();
    		window.location.href = return_url;
    	};

    	const click_handler_1 = () => {
    		cancel();
    		window.location.href = portal_url;
    	};

    	const click_handler_2 = () => {
    		cancel();
    		window.location.href = home_url;
    	};

    	const click_handler_3 = () => {
    		cancel();
    		app.clear_authed_data();
    		app.nav.goto_login_page();
    	};

    	$$self.$$set = $$props => {
    		if ('tenant_name' in $$props) $$invalidate(0, tenant_name = $$props.tenant_name);
    		if ('tenant_id' in $$props) $$invalidate(1, tenant_id = $$props.tenant_id);
    		if ('user_id' in $$props) $$invalidate(2, user_id = $$props.user_id);
    		if ('full_name' in $$props) $$invalidate(3, full_name = $$props.full_name);
    		if ('group_name' in $$props) $$invalidate(4, group_name = $$props.group_name);
    		if ('bio' in $$props) $$invalidate(5, bio = $$props.bio);
    		if ('return_url' in $$props) $$invalidate(6, return_url = $$props.return_url);
    	};

    	return [
    		tenant_name,
    		tenant_id,
    		user_id,
    		full_name,
    		group_name,
    		bio,
    		return_url,
    		show_timeout,
    		seconds,
    		portal_url,
    		home_url,
    		cancel,
    		app,
    		click_handler,
    		click_handler_1,
    		click_handler_2,
    		click_handler_3
    	];
    }

    class User_card extends SvelteComponent {
    	constructor(options) {
    		super();

    		init(this, options, instance$2, create_fragment$3, safe_not_equal, {
    			tenant_name: 0,
    			tenant_id: 1,
    			user_id: 2,
    			full_name: 3,
    			group_name: 4,
    			bio: 5,
    			return_url: 6
    		});
    	}
    }

    /* src/entries/auth/pages/final/index.svelte generated by Svelte v3.59.2 */

    function create_else_block(ctx) {
    	let usercard;
    	let current;

    	usercard = new User_card({
    			props: {
    				bio: /*userdata*/ ctx[0]["bio"] || "",
    				full_name: /*userdata*/ ctx[0]["full_name"] || "",
    				group_name: /*userdata*/ ctx[0]["group_name"] || "",
    				tenant_id: /*orgdata*/ ctx[1]["slug"] || "",
    				tenant_name: /*orgdata*/ ctx[1]["name"] || "",
    				user_id: /*userdata*/ ctx[0]["user_id"] || ""
    			}
    		});

    	return {
    		c() {
    			create_component(usercard.$$.fragment);
    		},
    		m(target, anchor) {
    			mount_component(usercard, target, anchor);
    			current = true;
    		},
    		p(ctx, dirty) {
    			const usercard_changes = {};
    			if (dirty & /*userdata*/ 1) usercard_changes.bio = /*userdata*/ ctx[0]["bio"] || "";
    			if (dirty & /*userdata*/ 1) usercard_changes.full_name = /*userdata*/ ctx[0]["full_name"] || "";
    			if (dirty & /*userdata*/ 1) usercard_changes.group_name = /*userdata*/ ctx[0]["group_name"] || "";
    			if (dirty & /*orgdata*/ 2) usercard_changes.tenant_id = /*orgdata*/ ctx[1]["slug"] || "";
    			if (dirty & /*orgdata*/ 2) usercard_changes.tenant_name = /*orgdata*/ ctx[1]["name"] || "";
    			if (dirty & /*userdata*/ 1) usercard_changes.user_id = /*userdata*/ ctx[0]["user_id"] || "";
    			usercard.$set(usercard_changes);
    		},
    		i(local) {
    			if (current) return;
    			transition_in(usercard.$$.fragment, local);
    			current = true;
    		},
    		o(local) {
    			transition_out(usercard.$$.fragment, local);
    			current = false;
    		},
    		d(detaching) {
    			destroy_component(usercard, detaching);
    		}
    	};
    }

    // (18:0) {#if loading}
    function create_if_block(ctx) {
    	let div;

    	return {
    		c() {
    			div = element("div");
    			div.textContent = "loading...";
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d(detaching) {
    			if (detaching) detach(div);
    		}
    	};
    }

    function create_fragment$2(ctx) {
    	let current_block_type_index;
    	let if_block;
    	let if_block_anchor;
    	let current;
    	const if_block_creators = [create_if_block, create_else_block];
    	const if_blocks = [];

    	function select_block_type(ctx, dirty) {
    		if (/*loading*/ ctx[2]) return 0;
    		return 1;
    	}

    	current_block_type_index = select_block_type(ctx);
    	if_block = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx);

    	return {
    		c() {
    			if_block.c();
    			if_block_anchor = empty();
    		},
    		m(target, anchor) {
    			if_blocks[current_block_type_index].m(target, anchor);
    			insert(target, if_block_anchor, anchor);
    			current = true;
    		},
    		p(ctx, [dirty]) {
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
    		i(local) {
    			if (current) return;
    			transition_in(if_block);
    			current = true;
    		},
    		o(local) {
    			transition_out(if_block);
    			current = false;
    		},
    		d(detaching) {
    			if_blocks[current_block_type_index].d(detaching);
    			if (detaching) detach(if_block_anchor);
    		}
    	};
    }

    function instance$1($$self, $$props, $$invalidate) {
    	const app = getContext("_auth_app_");
    	let userdata;
    	let orgdata;
    	let loading = true;
    	app.clear_preauthed_data();

    	(async () => {
    		const resp = await app.about();
    		console.log(resp);
    		$$invalidate(0, userdata = resp["user_info"]);
    		$$invalidate(1, orgdata = resp["org_info"]);
    		$$invalidate(2, loading = false);
    	})();

    	return [userdata, orgdata, loading];
    }

    class Final extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, instance$1, create_fragment$2, safe_not_equal, {});
    	}
    }

    /* src/entries/auth/pages/error/index.svelte generated by Svelte v3.59.2 */

    function create_fragment$1(ctx) {
    	let div;

    	return {
    		c() {
    			div = element("div");
    			div.textContent = "Error Page";
    		},
    		m(target, anchor) {
    			insert(target, div, anchor);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d(detaching) {
    			if (detaching) detach(div);
    		}
    	};
    }

    class Error$1 extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, null, create_fragment$1, safe_not_equal, {});
    	}
    }

    var auth_routes = {
        "/": Start,
        "/login": {
            "/": Login,
            "/next_stage": Nextstage$1,
        },
        "/signup": {
            "/": Signup,
            "/next_stage": Nextstage,
        },
        "/final": Final,
        "/prehook": Prehook,
        "/reset": {
            "/": Reset,
            "/finish": Finish,
        },
        "/alt": {
            "/first_stage": Firststage,
            "/second_stage": Secondstage,
        },
        "/error": Error$1,
    };

    /* src/entries/xcompo/common/_tailwind.svelte generated by Svelte v3.59.2 */

    class Tailwind extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, null, null, safe_not_equal, {});
    	}
    }

    /* src/entries/auth/index.svelte generated by Svelte v3.59.2 */

    function create_fragment(ctx) {
    	let router;
    	let t;
    	let tailwind;
    	let current;
    	router = new Router({});
    	tailwind = new Tailwind({});

    	return {
    		c() {
    			create_component(router.$$.fragment);
    			t = space();
    			create_component(tailwind.$$.fragment);
    		},
    		m(target, anchor) {
    			mount_component(router, target, anchor);
    			insert(target, t, anchor);
    			mount_component(tailwind, target, anchor);
    			current = true;
    		},
    		p: noop,
    		i(local) {
    			if (current) return;
    			transition_in(router.$$.fragment, local);
    			transition_in(tailwind.$$.fragment, local);
    			current = true;
    		},
    		o(local) {
    			transition_out(router.$$.fragment, local);
    			transition_out(tailwind.$$.fragment, local);
    			current = false;
    		},
    		d(detaching) {
    			destroy_component(router, detaching);
    			if (detaching) detach(t);
    			destroy_component(tailwind, detaching);
    		}
    	};
    }

    function instance($$self) {
    	const app = new AuthService();
    	setContext("_auth_app_", app);
    	routes.set(auth_routes);
    	return [];
    }

    class Auth extends SvelteComponent {
    	constructor(options) {
    		super();
    		init(this, options, instance, create_fragment, safe_not_equal, {});
    	}
    }

    const __svelte_app__ = new Auth({
        target: document.body,
        props: {}
    });

    return __svelte_app__;

})();
//# sourceMappingURL=auth.js.map

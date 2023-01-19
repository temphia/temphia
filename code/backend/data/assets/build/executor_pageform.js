(function () {
    'use strict';

    // register factor helper functions 
    const registerExecLoaderFactory = (name, factory) => registerFactory("loader.factory", name, factory);
    const registerFactory = (ftype, name, factory) => {
        const pf = window["__register_factory__"];
        if (!pf) {
            console.warn("factory registry not found");
            return;
        }
        pf(ftype, name, factory);
    };

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
    function set_attributes(node, attributes) {
        // @ts-ignore
        const descriptors = Object.getOwnPropertyDescriptors(node.__proto__);
        for (const key in attributes) {
            if (attributes[key] == null) {
                node.removeAttribute(key);
            }
            else if (key === 'style') {
                node.style.cssText = attributes[key];
            }
            else if (key === '__value') {
                node.value = node[key] = attributes[key];
            }
            else if (descriptors[key] && descriptors[key].set) {
                node[key] = attributes[key];
            }
            else {
                attr(node, key, attributes[key]);
            }
        }
    }
    function children(element) {
        return Array.from(element.childNodes);
    }
    function select_option(select, value) {
        for (let i = 0; i < select.options.length; i += 1) {
            const option = select.options[i];
            if (option.__value === value) {
                option.selected = true;
                return;
            }
        }
        select.selectedIndex = -1; // no option should be selected
    }
    function select_options(select, value) {
        for (let i = 0; i < select.options.length; i += 1) {
            const option = select.options[i];
            option.selected = ~value.indexOf(option.__value);
        }
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
    function prop_dev(node, property, value) {
        node[property] = value;
        dispatch_dev('SvelteDOMSetProperty', { node, property, value });
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

    const FieldShortText = "shorttext";
    const FieldLongText = "longtext";
    const FieldEmail = "email";
    const FieldPhone = "phone";
    const FieldCheckbox = "checkbox";
    const FieldSelect = "select";
    const FieldNumber = "number";
    const FieldRange = "range";
    const FieldColor = "color";
    /*

      executor
        plug.default.main
        plug.default.extension1

    */
    const data = {
        items: [
            {
                name: "title",
                type: "shorttext",
                options: [],
                html_attr: {},
            },
            {
                name: "info",
                type: "longtext",
                info: "what its about?"
            },
            {
                name: "done",
                type: "checkbox",
            },
        ],
        data: {},
        on_load: "",
        on_submit: "",
    };

    /* entries/executor_pageform/field/field.svelte generated by Svelte v3.48.0 */

    const file$2 = "entries/executor_pageform/field/field.svelte";

    function get_each_context$1(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[7] = list[i];
    	return child_ctx;
    }

    // (115:0) {:else}
    function create_else_block(ctx) {
    	let div;

    	const block = {
    		c: function create() {
    			div = element("div");
    			div.textContent = "Not Implemented";
    			add_location(div, file$2, 115, 2, 2906);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    		},
    		p: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_else_block.name,
    		type: "else",
    		source: "(115:0) {:else}",
    		ctx
    	});

    	return block;
    }

    // (104:30) 
    function create_if_block_8(ctx) {
    	let div;
    	let input;
    	let mounted;
    	let dispose;

    	let input_levels = [
    		{ type: "color" },
    		{ id: /*name*/ ctx[0] },
    		{ value: /*value*/ ctx[5] },
    		{
    			class: "p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    		},
    		/*html_attr*/ ctx[4]
    	];

    	let input_data = {};

    	for (let i = 0; i < input_levels.length; i += 1) {
    		input_data = assign(input_data, input_levels[i]);
    	}

    	const block = {
    		c: function create() {
    			div = element("div");
    			input = element("input");
    			set_attributes(input, input_data);
    			add_location(input, file$2, 105, 4, 2688);
    			attr_dev(div, "class", "flex w-full");
    			add_location(div, file$2, 104, 2, 2658);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, input);
    			if (input.autofocus) input.focus();

    			if (!mounted) {
    				dispose = listen_dev(input, "change", /*change*/ ctx[6], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			set_attributes(input, input_data = get_spread_update(input_levels, [
    				{ type: "color" },
    				dirty & /*name*/ 1 && { id: /*name*/ ctx[0] },
    				dirty & /*value*/ 32 && { value: /*value*/ ctx[5] },
    				{
    					class: "p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    				},
    				dirty & /*html_attr*/ 16 && /*html_attr*/ ctx[4]
    			]));
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_8.name,
    		type: "if",
    		source: "(104:30) ",
    		ctx
    	});

    	return block;
    }

    // (93:30) 
    function create_if_block_7(ctx) {
    	let div;
    	let input;
    	let mounted;
    	let dispose;

    	let input_levels = [
    		{ type: "range" },
    		{ id: /*name*/ ctx[0] },
    		{ value: /*value*/ ctx[5] },
    		{
    			class: "p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    		},
    		/*html_attr*/ ctx[4]
    	];

    	let input_data = {};

    	for (let i = 0; i < input_levels.length; i += 1) {
    		input_data = assign(input_data, input_levels[i]);
    	}

    	const block = {
    		c: function create() {
    			div = element("div");
    			input = element("input");
    			set_attributes(input, input_data);
    			add_location(input, file$2, 94, 4, 2417);
    			attr_dev(div, "class", "flex w-full");
    			add_location(div, file$2, 93, 2, 2387);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, input);
    			if (input.autofocus) input.focus();

    			if (!mounted) {
    				dispose = listen_dev(input, "change", /*change*/ ctx[6], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			set_attributes(input, input_data = get_spread_update(input_levels, [
    				{ type: "range" },
    				dirty & /*name*/ 1 && { id: /*name*/ ctx[0] },
    				dirty & /*value*/ 32 && { value: /*value*/ ctx[5] },
    				{
    					class: "p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    				},
    				dirty & /*html_attr*/ 16 && /*html_attr*/ ctx[4]
    			]));
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_7.name,
    		type: "if",
    		source: "(93:30) ",
    		ctx
    	});

    	return block;
    }

    // (81:30) 
    function create_if_block_6(ctx) {
    	let div;
    	let input;
    	let mounted;
    	let dispose;

    	let input_levels = [
    		{ type: "email" },
    		{ id: /*name*/ ctx[0] },
    		{ value: /*value*/ ctx[5] },
    		{
    			class: "p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    		},
    		{ placeholder: "mail@example.com" },
    		/*html_attr*/ ctx[4]
    	];

    	let input_data = {};

    	for (let i = 0; i < input_levels.length; i += 1) {
    		input_data = assign(input_data, input_levels[i]);
    	}

    	const block = {
    		c: function create() {
    			div = element("div");
    			input = element("input");
    			set_attributes(input, input_data);
    			add_location(input, file$2, 82, 4, 2109);
    			attr_dev(div, "class", "flex w-full");
    			add_location(div, file$2, 81, 2, 2079);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, input);
    			input.value = input_data.value;
    			if (input.autofocus) input.focus();

    			if (!mounted) {
    				dispose = listen_dev(input, "change", /*change*/ ctx[6], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			set_attributes(input, input_data = get_spread_update(input_levels, [
    				{ type: "email" },
    				dirty & /*name*/ 1 && { id: /*name*/ ctx[0] },
    				dirty & /*value*/ 32 && input.value !== /*value*/ ctx[5] && { value: /*value*/ ctx[5] },
    				{
    					class: "p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    				},
    				{ placeholder: "mail@example.com" },
    				dirty & /*html_attr*/ 16 && /*html_attr*/ ctx[4]
    			]));

    			if ('value' in input_data) {
    				input.value = input_data.value;
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
    		id: create_if_block_6.name,
    		type: "if",
    		source: "(81:30) ",
    		ctx
    	});

    	return block;
    }

    // (70:33) 
    function create_if_block_5(ctx) {
    	let div;
    	let textarea;
    	let mounted;
    	let dispose;

    	let textarea_levels = [
    		{ id: /*name*/ ctx[0] },
    		{ value: /*value*/ ctx[5] },
    		{
    			class: "p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    		},
    		{ placeholder: "write something..." },
    		/*html_attr*/ ctx[4]
    	];

    	let textarea_data = {};

    	for (let i = 0; i < textarea_levels.length; i += 1) {
    		textarea_data = assign(textarea_data, textarea_levels[i]);
    	}

    	const block = {
    		c: function create() {
    			div = element("div");
    			textarea = element("textarea");
    			set_attributes(textarea, textarea_data);
    			add_location(textarea, file$2, 71, 4, 1815);
    			attr_dev(div, "class", "flex w-full");
    			add_location(div, file$2, 70, 2, 1785);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, textarea);
    			if (textarea.autofocus) textarea.focus();

    			if (!mounted) {
    				dispose = listen_dev(textarea, "change", /*change*/ ctx[6], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			set_attributes(textarea, textarea_data = get_spread_update(textarea_levels, [
    				dirty & /*name*/ 1 && { id: /*name*/ ctx[0] },
    				dirty & /*value*/ 32 && { value: /*value*/ ctx[5] },
    				{
    					class: "p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    				},
    				{ placeholder: "write something..." },
    				dirty & /*html_attr*/ 16 && /*html_attr*/ ctx[4]
    			]));
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_5.name,
    		type: "if",
    		source: "(70:33) ",
    		ctx
    	});

    	return block;
    }

    // (59:31) 
    function create_if_block_4(ctx) {
    	let div;
    	let input;
    	let mounted;
    	let dispose;

    	let input_levels = [
    		{ type: "number" },
    		{ id: FieldNumber },
    		{ value: /*value*/ ctx[5] },
    		{
    			class: "p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    		},
    		/*html_attr*/ ctx[4]
    	];

    	let input_data = {};

    	for (let i = 0; i < input_levels.length; i += 1) {
    		input_data = assign(input_data, input_levels[i]);
    	}

    	const block = {
    		c: function create() {
    			div = element("div");
    			input = element("input");
    			set_attributes(input, input_data);
    			add_location(input, file$2, 60, 4, 1531);
    			attr_dev(div, "class", "flex w-full");
    			add_location(div, file$2, 59, 2, 1501);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, input);
    			if (input.autofocus) input.focus();

    			if (!mounted) {
    				dispose = listen_dev(input, "change", change_handler, false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			set_attributes(input, input_data = get_spread_update(input_levels, [
    				{ type: "number" },
    				{ id: FieldNumber },
    				dirty & /*value*/ 32 && input.value !== /*value*/ ctx[5] && { value: /*value*/ ctx[5] },
    				{
    					class: "p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    				},
    				dirty & /*html_attr*/ 16 && /*html_attr*/ ctx[4]
    			]));
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_4.name,
    		type: "if",
    		source: "(59:31) ",
    		ctx
    	});

    	return block;
    }

    // (48:33) 
    function create_if_block_3(ctx) {
    	let div;
    	let input;
    	let input_checked_value;
    	let mounted;
    	let dispose;

    	let input_levels = [
    		{ type: "checkbox" },
    		{ id: /*name*/ ctx[0] },
    		{ checked: input_checked_value = false },
    		{
    			class: "form-checkbox h-5 w-5 text-gray-600"
    		},
    		/*html_attr*/ ctx[4]
    	];

    	let input_data = {};

    	for (let i = 0; i < input_levels.length; i += 1) {
    		input_data = assign(input_data, input_levels[i]);
    	}

    	const block = {
    		c: function create() {
    			div = element("div");
    			input = element("input");
    			set_attributes(input, input_data);
    			add_location(input, file$2, 49, 4, 1290);
    			attr_dev(div, "class", "flex w-full");
    			add_location(div, file$2, 48, 2, 1260);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, input);
    			if (input.autofocus) input.focus();

    			if (!mounted) {
    				dispose = listen_dev(input, "change", null, false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			set_attributes(input, input_data = get_spread_update(input_levels, [
    				{ type: "checkbox" },
    				dirty & /*name*/ 1 && { id: /*name*/ ctx[0] },
    				{ checked: input_checked_value },
    				{
    					class: "form-checkbox h-5 w-5 text-gray-600"
    				},
    				dirty & /*html_attr*/ 16 && /*html_attr*/ ctx[4]
    			]));
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_3.name,
    		type: "if",
    		source: "(48:33) ",
    		ctx
    	});

    	return block;
    }

    // (35:31) 
    function create_if_block_2(ctx) {
    	let div;
    	let select;
    	let mounted;
    	let dispose;
    	let each_value = /*options*/ ctx[3] || [];
    	validate_each_argument(each_value);
    	let each_blocks = [];

    	for (let i = 0; i < each_value.length; i += 1) {
    		each_blocks[i] = create_each_block$1(get_each_context$1(ctx, each_value, i));
    	}

    	let select_levels = [
    		{ class: "w-full p-2 bg-gray-50 border" },
    		{ value: /*value*/ ctx[5] },
    		/*html_attr*/ ctx[4]
    	];

    	let select_data = {};

    	for (let i = 0; i < select_levels.length; i += 1) {
    		select_data = assign(select_data, select_levels[i]);
    	}

    	const block = {
    		c: function create() {
    			div = element("div");
    			select = element("select");

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].c();
    			}

    			set_attributes(select, select_data);
    			add_location(select, file$2, 36, 4, 992);
    			attr_dev(div, "class", "flex w-full");
    			add_location(div, file$2, 35, 2, 962);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, select);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].m(select, null);
    			}

    			(select_data.multiple ? select_options : select_option)(select, select_data.value);
    			if (select.autofocus) select.focus();

    			if (!mounted) {
    				dispose = listen_dev(select, "change", /*change*/ ctx[6], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*options*/ 8) {
    				each_value = /*options*/ ctx[3] || [];
    				validate_each_argument(each_value);
    				let i;

    				for (i = 0; i < each_value.length; i += 1) {
    					const child_ctx = get_each_context$1(ctx, each_value, i);

    					if (each_blocks[i]) {
    						each_blocks[i].p(child_ctx, dirty);
    					} else {
    						each_blocks[i] = create_each_block$1(child_ctx);
    						each_blocks[i].c();
    						each_blocks[i].m(select, null);
    					}
    				}

    				for (; i < each_blocks.length; i += 1) {
    					each_blocks[i].d(1);
    				}

    				each_blocks.length = each_value.length;
    			}

    			set_attributes(select, select_data = get_spread_update(select_levels, [
    				{ class: "w-full p-2 bg-gray-50 border" },
    				dirty & /*value*/ 32 && { value: /*value*/ ctx[5] },
    				dirty & /*html_attr*/ 16 && /*html_attr*/ ctx[4]
    			]));

    			if (dirty & /*value, html_attr*/ 48 && 'value' in select_data) (select_data.multiple ? select_options : select_option)(select, select_data.value);
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			destroy_each(each_blocks, detaching);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_2.name,
    		type: "if",
    		source: "(35:31) ",
    		ctx
    	});

    	return block;
    }

    // (24:30) 
    function create_if_block_1(ctx) {
    	let div;
    	let input;
    	let mounted;
    	let dispose;

    	let input_levels = [
    		{ type: "tel" },
    		{ id: /*name*/ ctx[0] },
    		{ value: /*value*/ ctx[5] },
    		{
    			class: "p-2 w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    		},
    		/*html_attr*/ ctx[4]
    	];

    	let input_data = {};

    	for (let i = 0; i < input_levels.length; i += 1) {
    		input_data = assign(input_data, input_levels[i]);
    	}

    	const block = {
    		c: function create() {
    			div = element("div");
    			input = element("input");
    			set_attributes(input, input_data);
    			add_location(input, file$2, 25, 4, 729);
    			attr_dev(div, "class", "flex w-full");
    			add_location(div, file$2, 24, 2, 699);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, input);
    			if (input.autofocus) input.focus();

    			if (!mounted) {
    				dispose = listen_dev(input, "change", /*change*/ ctx[6], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			set_attributes(input, input_data = get_spread_update(input_levels, [
    				{ type: "tel" },
    				dirty & /*name*/ 1 && { id: /*name*/ ctx[0] },
    				dirty & /*value*/ 32 && input.value !== /*value*/ ctx[5] && { value: /*value*/ ctx[5] },
    				{
    					class: "p-2 w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    				},
    				dirty & /*html_attr*/ 16 && /*html_attr*/ ctx[4]
    			]));
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
    		source: "(24:30) ",
    		ctx
    	});

    	return block;
    }

    // (13:0) {#if type === FieldShortText}
    function create_if_block(ctx) {
    	let div;
    	let input;
    	let mounted;
    	let dispose;

    	let input_levels = [
    		{ type: "text" },
    		{ id: /*name*/ ctx[0] },
    		{ value: /*value*/ ctx[5] },
    		{
    			class: "p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    		},
    		/*html_attr*/ ctx[4]
    	];

    	let input_data = {};

    	for (let i = 0; i < input_levels.length; i += 1) {
    		input_data = assign(input_data, input_levels[i]);
    	}

    	const block = {
    		c: function create() {
    			div = element("div");
    			input = element("input");
    			set_attributes(input, input_data);
    			add_location(input, file$2, 14, 4, 459);
    			attr_dev(div, "class", "flex w-full");
    			add_location(div, file$2, 13, 2, 429);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, input);
    			input.value = input_data.value;
    			if (input.autofocus) input.focus();

    			if (!mounted) {
    				dispose = listen_dev(input, "change", /*change*/ ctx[6], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			set_attributes(input, input_data = get_spread_update(input_levels, [
    				{ type: "text" },
    				dirty & /*name*/ 1 && { id: /*name*/ ctx[0] },
    				dirty & /*value*/ 32 && input.value !== /*value*/ ctx[5] && { value: /*value*/ ctx[5] },
    				{
    					class: "p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
    				},
    				dirty & /*html_attr*/ 16 && /*html_attr*/ ctx[4]
    			]));

    			if ('value' in input_data) {
    				input.value = input_data.value;
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
    		id: create_if_block.name,
    		type: "if",
    		source: "(13:0) {#if type === FieldShortText}",
    		ctx
    	});

    	return block;
    }

    // (43:6) {#each options || [] as opt}
    function create_each_block$1(ctx) {
    	let option;
    	let t_value = /*opt*/ ctx[7] + "";
    	let t;
    	let option_value_value;

    	const block = {
    		c: function create() {
    			option = element("option");
    			t = text(t_value);
    			option.__value = option_value_value = /*opt*/ ctx[7];
    			option.value = option.__value;
    			add_location(option, file$2, 43, 8, 1152);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, option, anchor);
    			append_dev(option, t);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*options*/ 8 && t_value !== (t_value = /*opt*/ ctx[7] + "")) set_data_dev(t, t_value);

    			if (dirty & /*options*/ 8 && option_value_value !== (option_value_value = /*opt*/ ctx[7])) {
    				prop_dev(option, "__value", option_value_value);
    				option.value = option.__value;
    			}
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(option);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block$1.name,
    		type: "each",
    		source: "(43:6) {#each options || [] as opt}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$3(ctx) {
    	let label;
    	let t0;
    	let t1;
    	let t2;
    	let p;
    	let t3;

    	function select_block_type(ctx, dirty) {
    		if (/*type*/ ctx[2] === FieldShortText) return create_if_block;
    		if (/*type*/ ctx[2] === FieldPhone) return create_if_block_1;
    		if (/*type*/ ctx[2] === FieldSelect) return create_if_block_2;
    		if (/*type*/ ctx[2] === FieldCheckbox) return create_if_block_3;
    		if (/*type*/ ctx[2] === FieldNumber) return create_if_block_4;
    		if (/*type*/ ctx[2] === FieldLongText) return create_if_block_5;
    		if (/*type*/ ctx[2] === FieldEmail) return create_if_block_6;
    		if (/*type*/ ctx[2] === FieldRange) return create_if_block_7;
    		if (/*type*/ ctx[2] === FieldColor) return create_if_block_8;
    		return create_else_block;
    	}

    	let current_block_type = select_block_type(ctx);
    	let if_block = current_block_type(ctx);

    	const block = {
    		c: function create() {
    			label = element("label");
    			t0 = text(/*name*/ ctx[0]);
    			t1 = space();
    			if_block.c();
    			t2 = space();
    			p = element("p");
    			t3 = text(/*info*/ ctx[1]);
    			attr_dev(label, "for", /*name*/ ctx[0]);
    			attr_dev(label, "class", "pb-2 text-gray-700 flex uppercase");
    			add_location(label, file$2, 10, 0, 320);
    			attr_dev(p, "class", "text-sm italic mb-4");
    			add_location(p, file$2, 118, 0, 2940);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, label, anchor);
    			append_dev(label, t0);
    			insert_dev(target, t1, anchor);
    			if_block.m(target, anchor);
    			insert_dev(target, t2, anchor);
    			insert_dev(target, p, anchor);
    			append_dev(p, t3);
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*name*/ 1) set_data_dev(t0, /*name*/ ctx[0]);

    			if (dirty & /*name*/ 1) {
    				attr_dev(label, "for", /*name*/ ctx[0]);
    			}

    			if (current_block_type === (current_block_type = select_block_type(ctx)) && if_block) {
    				if_block.p(ctx, dirty);
    			} else {
    				if_block.d(1);
    				if_block = current_block_type(ctx);

    				if (if_block) {
    					if_block.c();
    					if_block.m(t2.parentNode, t2);
    				}
    			}

    			if (dirty & /*info*/ 2) set_data_dev(t3, /*info*/ ctx[1]);
    		},
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(label);
    			if (detaching) detach_dev(t1);
    			if_block.d(detaching);
    			if (detaching) detach_dev(t2);
    			if (detaching) detach_dev(p);
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

    const change_handler = () => {
    	
    };

    function instance$3($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('Field', slots, []);
    	let { name } = $$props;
    	let { info } = $$props;
    	let { type } = $$props;
    	let { options } = $$props;
    	let { html_attr = {} } = $$props;
    	let { value } = $$props;

    	const change = ev => {
    		
    	};

    	const writable_props = ['name', 'info', 'type', 'options', 'html_attr', 'value'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Field> was created with unknown prop '${key}'`);
    	});

    	$$self.$$set = $$props => {
    		if ('name' in $$props) $$invalidate(0, name = $$props.name);
    		if ('info' in $$props) $$invalidate(1, info = $$props.info);
    		if ('type' in $$props) $$invalidate(2, type = $$props.type);
    		if ('options' in $$props) $$invalidate(3, options = $$props.options);
    		if ('html_attr' in $$props) $$invalidate(4, html_attr = $$props.html_attr);
    		if ('value' in $$props) $$invalidate(5, value = $$props.value);
    	};

    	$$self.$capture_state = () => ({
    		FieldShortText,
    		FieldLongText,
    		FieldEmail,
    		FieldCheckbox,
    		FieldNumber,
    		FieldPhone,
    		FieldSelect,
    		FieldRange,
    		FieldColor,
    		name,
    		info,
    		type,
    		options,
    		html_attr,
    		value,
    		change
    	});

    	$$self.$inject_state = $$props => {
    		if ('name' in $$props) $$invalidate(0, name = $$props.name);
    		if ('info' in $$props) $$invalidate(1, info = $$props.info);
    		if ('type' in $$props) $$invalidate(2, type = $$props.type);
    		if ('options' in $$props) $$invalidate(3, options = $$props.options);
    		if ('html_attr' in $$props) $$invalidate(4, html_attr = $$props.html_attr);
    		if ('value' in $$props) $$invalidate(5, value = $$props.value);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [name, info, type, options, html_attr, value, change];
    }

    class Field extends SvelteComponentDev {
    	constructor(options) {
    		super(options);

    		init(this, options, instance$3, create_fragment$3, safe_not_equal, {
    			name: 0,
    			info: 1,
    			type: 2,
    			options: 3,
    			html_attr: 4,
    			value: 5
    		});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Field",
    			options,
    			id: create_fragment$3.name
    		});

    		const { ctx } = this.$$;
    		const props = options.props || {};

    		if (/*name*/ ctx[0] === undefined && !('name' in props)) {
    			console.warn("<Field> was created without expected prop 'name'");
    		}

    		if (/*info*/ ctx[1] === undefined && !('info' in props)) {
    			console.warn("<Field> was created without expected prop 'info'");
    		}

    		if (/*type*/ ctx[2] === undefined && !('type' in props)) {
    			console.warn("<Field> was created without expected prop 'type'");
    		}

    		if (/*options*/ ctx[3] === undefined && !('options' in props)) {
    			console.warn("<Field> was created without expected prop 'options'");
    		}

    		if (/*value*/ ctx[5] === undefined && !('value' in props)) {
    			console.warn("<Field> was created without expected prop 'value'");
    		}
    	}

    	get name() {
    		throw new Error("<Field>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set name(value) {
    		throw new Error("<Field>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get info() {
    		throw new Error("<Field>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set info(value) {
    		throw new Error("<Field>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get type() {
    		throw new Error("<Field>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set type(value) {
    		throw new Error("<Field>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get options() {
    		throw new Error("<Field>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set options(value) {
    		throw new Error("<Field>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get html_attr() {
    		throw new Error("<Field>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set html_attr(value) {
    		throw new Error("<Field>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get value() {
    		throw new Error("<Field>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set value(value) {
    		throw new Error("<Field>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* entries/executor_pageform/pages/start/start.svelte generated by Svelte v3.48.0 */
    const file$1 = "entries/executor_pageform/pages/start/start.svelte";

    function get_each_context(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[1] = list[i];
    	return child_ctx;
    }

    // (5:0) {#each data.items as item}
    function create_each_block(ctx) {
    	let field;
    	let current;

    	field = new Field({
    			props: {
    				html_attr: /*item*/ ctx[1].html_attr,
    				info: "",
    				name: /*item*/ ctx[1].name,
    				options: /*item*/ ctx[1].options,
    				type: /*item*/ ctx[1].type,
    				value: null
    			},
    			$$inline: true
    		});

    	const block = {
    		c: function create() {
    			create_component(field.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(field, target, anchor);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			const field_changes = {};
    			if (dirty & /*data*/ 1) field_changes.html_attr = /*item*/ ctx[1].html_attr;
    			if (dirty & /*data*/ 1) field_changes.name = /*item*/ ctx[1].name;
    			if (dirty & /*data*/ 1) field_changes.options = /*item*/ ctx[1].options;
    			if (dirty & /*data*/ 1) field_changes.type = /*item*/ ctx[1].type;
    			field.$set(field_changes);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(field.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(field.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(field, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block.name,
    		type: "each",
    		source: "(5:0) {#each data.items as item}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$2(ctx) {
    	let t0;
    	let div;
    	let button;
    	let svg;
    	let path;
    	let t1;
    	let current;
    	let each_value = /*data*/ ctx[0].items;
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

    			t0 = space();
    			div = element("div");
    			button = element("button");
    			svg = svg_element("svg");
    			path = svg_element("path");
    			t1 = text("\n    Submit");
    			attr_dev(path, "fill-rule", "evenodd");
    			attr_dev(path, "d", "M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z");
    			attr_dev(path, "clip-rule", "evenodd");
    			add_location(path, file$1, 25, 6, 548);
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg, "class", "h-5 w-5");
    			attr_dev(svg, "viewBox", "0 0 20 20");
    			attr_dev(svg, "fill", "currentColor");
    			add_location(svg, file$1, 19, 4, 416);
    			attr_dev(button, "class", "p-1 rounded bg-green-500 shadow hover:bg-green-900 flex text-white");
    			add_location(button, file$1, 16, 2, 321);
    			attr_dev(div, "class", "flex justify-end items-center p-1");
    			add_location(div, file$1, 15, 0, 271);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].m(target, anchor);
    			}

    			insert_dev(target, t0, anchor);
    			insert_dev(target, div, anchor);
    			append_dev(div, button);
    			append_dev(button, svg);
    			append_dev(svg, path);
    			append_dev(button, t1);
    			current = true;
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*data*/ 1) {
    				each_value = /*data*/ ctx[0].items;
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
    						each_blocks[i].m(t0.parentNode, t0);
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
    			if (detaching) detach_dev(t0);
    			if (detaching) detach_dev(div);
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
    	validate_slots('Start', slots, []);
    	let { data } = $$props;
    	const writable_props = ['data'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Start> was created with unknown prop '${key}'`);
    	});

    	$$self.$$set = $$props => {
    		if ('data' in $$props) $$invalidate(0, data = $$props.data);
    	};

    	$$self.$capture_state = () => ({ Field, data });

    	$$self.$inject_state = $$props => {
    		if ('data' in $$props) $$invalidate(0, data = $$props.data);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [data];
    }

    class Start extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$2, create_fragment$2, safe_not_equal, { data: 0 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Start",
    			options,
    			id: create_fragment$2.name
    		});

    		const { ctx } = this.$$;
    		const props = options.props || {};

    		if (/*data*/ ctx[0] === undefined && !('data' in props)) {
    			console.warn("<Start> was created without expected prop 'data'");
    		}
    	}

    	get data() {
    		throw new Error("<Start>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set data(value) {
    		throw new Error("<Start>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
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

    /* entries/executor_pageform/index.svelte generated by Svelte v3.48.0 */
    const file = "entries/executor_pageform/index.svelte";

    function create_fragment(ctx) {
    	let div1;
    	let div0;
    	let h1;
    	let t1;
    	let startpage;
    	let t2;
    	let tailwind;
    	let current;
    	startpage = new Start({ props: { data }, $$inline: true });
    	tailwind = new Tailwind({ $$inline: true });

    	const block = {
    		c: function create() {
    			div1 = element("div");
    			div0 = element("div");
    			h1 = element("h1");
    			h1.textContent = "Form 1";
    			t1 = space();
    			create_component(startpage.$$.fragment);
    			t2 = space();
    			create_component(tailwind.$$.fragment);
    			attr_dev(h1, "class", "text-gray-800 font-lg font-bold tracking-normal leading-tight mb-4");
    			add_location(h1, file, 8, 4, 277);
    			attr_dev(div0, "class", "p-4 bg-white rounded");
    			add_location(div0, file, 7, 2, 238);
    			attr_dev(div1, "class", "h-full w-full p-4 bg-blue-50");
    			add_location(div1, file, 6, 0, 193);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div1, anchor);
    			append_dev(div1, div0);
    			append_dev(div0, h1);
    			append_dev(div0, t1);
    			mount_component(startpage, div0, null);
    			insert_dev(target, t2, anchor);
    			mount_component(tailwind, target, anchor);
    			current = true;
    		},
    		p: noop,
    		i: function intro(local) {
    			if (current) return;
    			transition_in(startpage.$$.fragment, local);
    			transition_in(tailwind.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(startpage.$$.fragment, local);
    			transition_out(tailwind.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div1);
    			destroy_component(startpage);
    			if (detaching) detach_dev(t2);
    			destroy_component(tailwind, detaching);
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
    	validate_slots('Executor_pageform', slots, []);
    	let { env } = $$props;
    	const writable_props = ['env'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<Executor_pageform> was created with unknown prop '${key}'`);
    	});

    	$$self.$$set = $$props => {
    		if ('env' in $$props) $$invalidate(0, env = $$props.env);
    	};

    	$$self.$capture_state = () => ({ StartPage: Start, data, Tailwind, env });

    	$$self.$inject_state = $$props => {
    		if ('env' in $$props) $$invalidate(0, env = $$props.env);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [env];
    }

    class Executor_pageform extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance, create_fragment, safe_not_equal, { env: 0 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Executor_pageform",
    			options,
    			id: create_fragment.name
    		});

    		const { ctx } = this.$$;
    		const props = options.props || {};

    		if (/*env*/ ctx[0] === undefined && !('env' in props)) {
    			console.warn("<Executor_pageform> was created without expected prop 'env'");
    		}
    	}

    	get env() {
    		throw new Error("<Executor_pageform>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set env(value) {
    		throw new Error("<Executor_pageform>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    registerExecLoaderFactory("pageform.loader", (opts) => {
        console.log("@@pagefrom.loader", opts);
        new Executor_pageform({
            target: opts.target,
            props: {
                env: opts.env,
            },
        });
    });

})();
//# sourceMappingURL=executor_pageform.js.map

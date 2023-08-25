import type { SvelteComponent } from 'svelte';

declare module '*.svelte' {
    export { SvelteComponent as default };
}
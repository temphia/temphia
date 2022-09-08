import svelte from "rollup-plugin-svelte";
import commonjs from "@rollup/plugin-commonjs";
import resolve from "@rollup/plugin-node-resolve";
import sveltePreprocess from "svelte-preprocess";
import typescript from "@rollup/plugin-typescript";
import css from "rollup-plugin-css-only";
import json from "rollup-plugin-json";
import alias from '@rollup/plugin-alias';

const production = !!process.env.PRODUCTION;
const entryFile = process.env.ENTRY_FILE;

const _config = (entryFile, prod) => ({
  input: `src/entries/${entryFile}.ts`,
  output: {
    sourcemap: true,
    format: "iife",
    name: "app",
    file: `public/build/${entryFile}.js`,
  },
  plugins: [
    alias({
      resolve: ['.svelte', '.js', 'ts'],
      entries: [
        { find: '@lib', replacement: 'src/lib' },
      ]
    }),

    svelte({
      preprocess: sveltePreprocess({ postcss: true }),
      compilerOptions: {
        // enable run-time checks when not in production
        dev: !prod,
      },
    }),
    // we'll extract any component CSS out into
    // a separate file - better for performance
    css({ output: `${entryFile}.css` }),

    // If you have external dependencies installed from
    // npm, you'll most likely need these plugins. In
    // some cases you'll need additional configuration -
    // consult the documentation for details:
    // https://github.com/rollup/plugins/tree/master/packages/commonjs
    resolve({
      browser: true,
      dedupe: ["svelte"],
    }),
    commonjs(),
    json(),
    typescript({
      sourceMap: !prod,
      inlineSources: !prod,
    }),
  ],
});

export default _config(entryFile, production)




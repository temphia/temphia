import svelte from "rollup-plugin-svelte";
import commonjs from "@rollup/plugin-commonjs";
import resolve from "@rollup/plugin-node-resolve";
import sveltePreprocess from "svelte-preprocess";
import typescript from "@rollup/plugin-typescript";
import css from "rollup-plugin-css-only";
import json from "rollup-plugin-json";

const entryFile = process.env.ENTRY_FILE;

export const config = (prod) => ({
  input: `src/entry/${entryFile}.ts`,
  output: {
    sourcemap: true,
    format: "iife",
    name: "app",
    file: `public/build/${entryFile}.js`,
  },
  plugins: [
    svelte({
      preprocess: sveltePreprocess({ postcss: true }),
      compilerOptions: {
        dev: !prod,
      },
    }),
    css({ output: `${entryFile}.css` }),

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

export default config(false);

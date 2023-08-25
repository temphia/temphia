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
const outFile = process.env.OUT_FILE || entryFile

let outdir = "build_dev"
if (production) {
  outdir = "build_prod"
}

const _config = (entryFile, prod) => ({
  input: `src/entries/${entryFile}/index.ts`,
  output: {
    sourcemap: true,
    format: "iife",
    name: "app",
    file: `${outdir}/${outFile}.js`,

  },
  plugins: [
    alias({
      resolve: ['.svelte', '.js', 'ts'],
    }),

    svelte({
      preprocess: sveltePreprocess({ postcss: true }),
      compilerOptions: {
        dev: !prod,
      },
    }),

    css({ output: `${outFile}.css` }),
    resolve({
      browser: true,
      dedupe: ["svelte"],
    }),
    commonjs(),
    json(),
    typescript({
      sourceMap: true,
      inlineSources: true,
      rootDir: './src',
    }),
  ],
});

export default _config(entryFile, production)




import { defineConfig, UserConfig } from 'vite'
import { svelte, vitePreprocess } from '@sveltejs/vite-plugin-svelte'

const sp = svelte({
  preprocess: vitePreprocess(),
})

const isAuth = process.env.AUTH_UI

let config: UserConfig = {
  plugins: [sp],
  build: {
    outDir: "dist",
    lib: {
      entry: ["src/entries/portal/index.ts"],
      formats: ["iife"],
      name: "portal"
    },
  },
}

if (isAuth) {
  config = {
    plugins: [sp],
    build: {
      outDir: "dist",
      lib: {
        entry: ["src/entries/auth/index.ts"],
        formats: ["iife"],
        name: "auth"
      },
    },
  }
}

export default defineConfig(config)

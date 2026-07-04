import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  // Relative asset URLs so they resolve under the injected <base> when the
  // panel is served from a secret sub-path.
  base: './',
  server: {
    port: 5173,
    proxy: {
      '/api': 'http://localhost:8686',
    },
  },
  build: {
    // The Go binary embeds this directory via go:embed.
    outDir: '../internal/webfs/dist',
    emptyOutDir: true,
  },
})

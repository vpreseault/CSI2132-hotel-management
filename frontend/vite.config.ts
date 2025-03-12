import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue';
import Components from 'unplugin-vue-components/vite';
import {PrimeVueResolver} from '@primevue/auto-import-resolver';
import tailwindcss from '@tailwindcss/vite'


export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  return {
    plugins: [
      vue(),
      Components({
        resolvers: [
          PrimeVueResolver()
        ]
      }),
      tailwindcss(),
    ],
    server: {
      proxy: {
        '/api': {
          target: env.VITE_BACKEND_HOST,
          changeOrigin: true,
          secure: false, // If your backend uses self-signed certificates, set to false
        },
      },
      port: Number(env.PORT),
    },
  };
})
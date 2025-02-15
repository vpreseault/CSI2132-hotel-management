import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue';

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  return {
    plugins: [vue()],
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
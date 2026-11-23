import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    // 开发时把接口请求转发到 Go 后端，避免跨域
    proxy: {
      '/api': 'http://localhost:8080',
    },
  },
})

import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import VueDevTools from 'vite-plugin-vue-devtools'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue(), vueJsx(), VueDevTools()],
    resolve: {
        alias: {
            vue: 'vue/dist/vue.esm-bundler.js',
            '@': fileURLToPath(new URL('./src', import.meta.url)),
            '@/models': fileURLToPath(new URL('./src/models', import.meta.url)),
            '@/components': fileURLToPath(new URL('./src/components', import.meta.url)),
            '@/database': fileURLToPath(new URL('./src/database', import.meta.url)),
            '@/stores': fileURLToPath(new URL('./src/stores', import.meta.url)),
            '@/utils': fileURLToPath(new URL('./src/utils', import.meta.url)),
            '@/handlers': fileURLToPath(new URL('./src/handlers', import.meta.url)),
            '@/icons': fileURLToPath(new URL('./src/assets/icons', import.meta.url)),
            '@/style': fileURLToPath(new URL('./src/assets/style', import.meta.url)),
            '@images': fileURLToPath(new URL('./src/assets/images', import.meta.url)),
            '@/plugins': fileURLToPath(new URL('./src/plugins', import.meta.url))
        }
    },
    css: {
        // css预处理器

        preprocessorOptions: {
            less: {
                // 将你需要引入的less文件路径添加到prependData中
                prependData: `@import "@/node_modules/vditor/src/assets/less/index.less";`
            }
        }
    }
})

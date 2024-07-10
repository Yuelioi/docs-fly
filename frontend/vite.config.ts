import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import VueDevTools from 'vite-plugin-vue-devtools'

import dynamicThemePlugin from './tools/vite-plugin-dynamic-theme'

import Components from 'unplugin-vue-components/vite'

import AutoImport from 'unplugin-auto-import/vite'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        vueJsx(),
        VueDevTools(),
        dynamicThemePlugin(),
        AutoImport({
            dts: 'auto-imports.d.ts',

            imports: [
                'vue',
                'vue-router',
                'pinia',
                { '@vueuse/core': ['useDark', 'useToggle'] },
                {
                    from: 'vue-router',
                    imports: ['RouteParams', 'RouteLocationNormalizedLoaded'],
                    type: true
                }
            ],
            dirs: [
                './src/stores/',
                './src/services/',
                './src/utils/**',
                './src/hooks/**',
                './src/plugins/*',
                './src/database/'
            ]
        }),

        Components({
            dirs: ['src/components'],
            extensions: ['vue', 'md'],
            include: [/\.vue$/, /\.vue\?vue/, /\.md$/],
            exclude: [/[\\/]node_modules[\\/]/, /[\\/]\.git[\\/]/],
            dts: 'components.d.ts'
        })
    ],
    resolve: {
        alias: {
            vue: 'vue/dist/vue.esm-bundler.js',

            '@': fileURLToPath(new URL('./src', import.meta.url)),
            '@/models': fileURLToPath(new URL('./src/models', import.meta.url)),
            '@/components': fileURLToPath(new URL('./src/components', import.meta.url)),
            '@/utils': fileURLToPath(new URL('./src/utils', import.meta.url))
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
    },
    publicDir: 'public'

    // assetsInclude: ['**/*.riv']
})

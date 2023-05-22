// https://nuxt.com/docs/api/configuration/nuxt-config
import {loadEnv} from 'vite'

interface VITE_ENV_CONFIG {
    VITE_API_ENV: string,
    VITE_API_URL: string,
    VITE_OSS_URL: string,
}

const envScript = process.env.npm_lifecycle_script?.split(' ')
const envName = envScript[envScript.length - 1] // 通过启动命令区分环境
const envData = loadEnv(envName, 'env') as unknown as VITE_ENV_CONFIG
console.debug(envScript)
console.debug(envName)
console.debug(envData)
export default defineNuxtConfig({
    name: 'Share-CV',
    runtimeConfig: {
        public: envData
    }, // 把env放入这个里面，通过useRuntimeConfig获取
    css: [
        "@/assets/styles/main.scss",
    ],
    modules: [
        'nuxt-windicss',
        '@element-plus/nuxt',
        'nuxt-vue3-google-signin',
        "@pinia/nuxt",
        '@pinia-plugin-persistedstate/nuxt'
    ],
    googleSignIn: {
        clientId: '764151478265-dd5pm3jojqstquo1c3u3n9qam7vbv04q.apps.googleusercontent.com',
    },
    piniaPersistedstate: {
        cookieOptions: {
            sameSite: 'strict',
        },
        storage: 'localStorage'
    },
    build: {
        vendor: ['vue-cropper'],
        plugins: [
            {src: '~/plugins/vue-cropper', ssr: false},
        ],
        transpile:
            process.env.NODE_ENV === 'production'
                ? [
                    'naive-ui',
                    'vueuc',
                    '@css-render/vue3-ssr',
                    '@juggle/resize-observer'
                ]
                : ['@juggle/resize-observer']
    },
    vite: {
        optimizeDeps: {
            include:
                process.env.NODE_ENV === 'development'
                    ? ['naive-ui', 'vueuc', 'date-fns-tz/esm/formatInTimeZone']
                    : []
        },
        css: {
            preprocessorOptions: {
                scss: {
                    additionalData: '@use "@/assets/styles/variable.scss" as *;'
                }
            }
        },
    },
})

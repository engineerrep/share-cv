import { createI18n } from 'vue-i18n'
import { useState, useCookie } from "nuxt/app"
import { i18n } from "@/locals/config"

// import en from '@/locals/en-US'
// import cn from '@/locals/zh-CN'

// const message = {
//   cn,
//   en
// }

// export const i18n = createI18n({
//   legacy: false,
//   globalInjection: true,
//   fallbackLocale: 'en',
//   locale: 'en', //获取cookie值，默认为cn
//   warnHtmlMessage: false,
//   messages: langMsg
// })

export default defineNuxtPlugin(nuxtApp => {
  const languages = useCookie('lang')
  i18n.global.locale.value = languages.value || 'en'
  nuxtApp.vueApp.use(i18n)
})

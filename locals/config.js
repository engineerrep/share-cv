import { createI18n } from 'vue-i18n'

import cn from '@/locals/zh-CN'
import en from '@/locals/en-US'

const langMsg = {
  cn,
  en
}

export const langConfig = [
  {label: "中文", key: "cn"},
  {label: "English", key: "en"}
]

export const i18n = createI18n({
  legacy: false,
  globalInjection: true,
  fallbackLocale: 'en',
  locale: 'en', //获取cookie值，默认为en
  warnHtmlMessage: false,
  messages: langMsg
})
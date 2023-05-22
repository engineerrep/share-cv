import { useState, useCookie } from "nuxt/app";
// import { useState, useCookie } from '#app' 这两种方式都可以导入useState,方法
//示例：（使用命名导出）
export const userLanguages = () => {
  const languages = useCookie('lang')
  return useState('userLang', () => languages.value ?? 'en')
}
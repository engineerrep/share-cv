import { useState, useCookie } from "nuxt/app"

export const useImgUrl = () => {
  const config = useRuntimeConfig()
  const imgUrl = config.public.VITE_OSS_URL
  return useState("imgUrl", () => imgUrl)
}
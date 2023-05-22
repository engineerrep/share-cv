import {useToken, useUserInfo} from "@/composables/user"

// 可以直接访问的页面
const allowList = ["/", "/register", '/forgot-password', "/legal/privacy-policy", "/legal/cookie-policy", "/legal/user-agreement"]

export default defineNuxtRouteMiddleware((to, from) => {
    const token = useToken().value

    // 判断登录情况
    if (token) {
        if (to.path == "/") {
            return navigateTo('/personal-wditor')
        } else {
            const user = useUserInfo()
            if (!user.value) {
                // getUserInfo()
            }
        }
    } else {
        if (!allowList.includes(to.path)) {
            return navigateTo('/')
        }
    }
})
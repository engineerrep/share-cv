import {useState, useCookie} from "nuxt/app"
import {logout} from "@/api/login"
import {userInfo} from "@/api/user"
import {createDiscreteApi} from "naive-ui"
import {i18n} from "@/locals/config"

const {t} = i18n.global

export const useToken = () => {
    const cookieToken = useCookie("token")
    return useCookie('token', () => cookieToken.value ?? '')
}

/**
 * @name  注销登录
 */
export const toLogout = async () => {
    const {message} = createDiscreteApi(["message"])
    const cookieToken = useCookie("token")
    const token = useToken()
    const user = useUserInfo()
    let res = await logout()

    if (res.data.status === 1) {
        token.value = ""
        cookieToken.value = ""
        user.value = {}
        setTimeout(() => {
            message.warning(t("login.您已退出登录"))
        }, 500)
    }
}

/**
 * @name 获取用户信息
 */

export const useUserInfo = () => {
    return useState("userInfo", () => {
    })
}

export const getUserInfo = async () => {
    const user = useUserInfo()
    let res = await userInfo()
    // console.log(res);
    user.value = res.data
}
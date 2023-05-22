import Api from "~/utils/api"

// 获取谷歌授权地址
export const googleAuthUrl = (params) => {
    return Api.post('/api/web/google/auth/url', params)
}

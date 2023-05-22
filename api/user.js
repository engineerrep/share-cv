import Api from "~/utils/api"

// 用户信息
export const userInfo = (params) => {
    return Api.get('/api/web/user/info')
}
// 用户信息
export const apUserInfo = (params) => {
    return Api.post('/api/userbase/info')
}

// // 修改密码
// export const apiAccountChangePassword = (params) => {
//   return Api.post('/api/account/change/password')
// }
// 修改密码
export const apiAccountForgetPassword = (params) => {
    return Api.post('/api/account/forget/password', params)
}

// 登录密码改邮箱
export const apiAccountChangeEmailWithPassword = (params) => {
    return Api.post('/api/account/change/email/with/password')
}

// 邮箱验证码验证
export const apiAccountChangeEmailCode = (params) => {
    return Api.post('/api/account/check/email/code', params)
}
// 反馈
export const apiFeedback = (params) => {
    return Api.post('/api/userbase/feedback', params)
}
